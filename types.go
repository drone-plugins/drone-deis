package main

type Params struct {
	Controller  string `json:"controller"`
	Application string `json:"app"`
	Force       bool   `json:"force"`
	Commit      bool   `json:"commit"`
}
