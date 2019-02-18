package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/appleboy/drone-git-push/repo"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	version = "unknown"
)

func main() {
	fmt.Printf("Drone Deis Plugin built from %s\n", version)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Controller) == 0 {
		fmt.Println("Please provide a controller")

		os.Exit(1)
		return
	}

	if len(vargs.Application) == 0 {
		vargs.Application = repo.Name
	}

	err := run(&workspace, &build, &vargs)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

func run(workspace *drone.Workspace, build *drone.Build, vargs *Params) error {
	repo.GlobalName(build.Email).Run()
	repo.GlobalUser(build.Author).Run()

	if err := repo.WriteKey(workspace.Keys.Private); err != nil {
		return err
	}

	defer func() {
		execute(
			repo.RemoteRemove(
				"deploy"),
			workspace)
	}()

	cmd := repo.RemoteAdd(
		"deploy",
		remote(vargs))

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	if vargs.Commit {
		if err := execute(repo.ForceAdd(), workspace); err != nil {
			return err
		}

		if err := execute(repo.ForceCommit(), workspace); err != nil {
			return err
		}
	}

	cmd = repo.RemotePush(
		"deploy",
		"master",
		vargs.Force)

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	return nil
}

func remote(vargs *Params) string {
	return fmt.Sprintf(
		"ssh://git@%s:2222/%s.git",
		vargs.Controller,
		vargs.Application)
}

func execute(cmd *exec.Cmd, workspace *drone.Workspace) error {
	trace(cmd)

	cmd.Dir = workspace.Path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
