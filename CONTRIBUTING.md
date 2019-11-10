# Contributing

If you want to contribute, but not sure what to do, here's a list of things that I always need help with:

* Translations
    * README.md
    * [docs](https://github.com/GoAdminGroup/docs/issues/1)
* Bug-hunting
* Finding security problems
* Themes and Plugins

See [manual](https://github.com/GoAdminGroup/go-admin/projects/3) for more information.

You can view all open issues on github, which is usually a good starting point if you want to start contributing:

https://github.com/search?q=org%3AGoAdminGroup+is%3Aopen+is%3Aissue+archived%3Afalse&type=Issues

## how to

GoAdmin uses GitHub to manage reviews of pull requests:

- If you have a trivial fix or improvement, go ahead and create a pull request.
- If you plan to do something more involved, discuss your ideas on the relevant GitHub issue.

For now, you need to add your fork as a remote on the original **\$GOPATH**/src/github.com/GoAdminGroup/go-admin clone, so:

```bash

$ go get github.com/GoAdminGroup/go-admin
$ cd $GOPATH/src/github.com/GoAdminGroup/go-admin # GOPATH is $HOME/go by default.

$ git remote add <FORK_NAME> <FORK_URL>
```

And before you commit, remember to execute the command: 

```
make test
```

See the Makefile for more details.

Notice: `go get` return `package github.com/GoAdminGroup/go-admin: no Go files in /go/src/github.com/GoAdminGroup/go-admin` is normal.

### Dependency management

We uses [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies on external packages.
This requires a working Go environment with version 1.13 or greater and git installed.

To add or update a new dependency, use the `go get` command:

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```

Tidy up the `go.mod` and `go.sum` files:

```bash
go mod tidy
go mod vendor
git add go.mod go.sum vendor
git commit
```

You have to commit the changes to `go.mod` and `go.sum` before submitting the pull request.

# Support

You can also donate or become a patreon, which helps out covering server costs and potentially make it possible to put out bounties:

* **Support on [Open Collective](https://opencollective.com/go-admin)**
* Donate via [PayPal](https://paypal.me/cg80333)

# Members

If you are a member of the official GoAdmin developer Team:

* [Discussions](http://forum.go-admin.cn)
* [Tasks](https://github.com/GoAdminGroup/go-admin/projects)
* [Chat](https://t.me/joinchat/NlyH6Bch2QARZkArithKvg)