Use this plugin for deplying an application to Deis. You can override the
default configuration with the following parameters:

* `controller` - Deis controller, e.g. deis.deisapps.com
* `app` - Application name, defaults to repo name
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  deis:
    controller: deis.deisapps.com
    app: awesomeness
    force: true
```
