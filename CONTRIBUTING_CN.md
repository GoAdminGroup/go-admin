# 贡献

如果你想要对项目作出贡献，却不知道怎么做，下面有一些帮助：

* 翻译
    * README.md
    * [docs](https://github.com/GoAdminGroup/docs/issues/1)
* 寻找BUG
* 寻找安全问题
* 主题和插件

在这里：[功能规划](https://github.com/GoAdminGroup/go-admin/projects/3) 可以获得更多信息。

你也可以看一下所有开放的issues，从这里去入手：

https://github.com/search?q=org%3AGoAdminGroup+is%3Aopen+is%3Aissue+archived%3Afalse&type=Issues

## 如何做贡献

GoAdmin 使用 GitHub 来管理项目代码：

- 如果你发现一些微不足道的fix或者功能增加，直接提pr即可；
- 如果你有一些提议，那么你可以先开一个issue进行讨论；

然后，你需要fork远程的master分支到你本地 **\$GOPATH**/src/github.com/GoAdminGroup/go-admin ：

```bash

$ go get github.com/GoAdminGroup/go-admin
$ cd $GOPATH/src/github.com/GoAdminGroup/go-admin # GOPATH is $HOME/go by default.

$ git remote add <FORK_NAME> <FORK_URL>
```

在你提交代码之前，记得执行下面这个命令： 

```
make test
```

看根目录下的```Makefile```获得更多信息。

注意了： `go get` 返回 `package github.com/GoAdminGroup/go-admin: no Go files in /go/src/github.com/GoAdminGroup/go-admin` 是正常的。

### 依赖管理

我们使用 [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) 来管理依赖。
这需要 golang 版本大于1.11，以及安装了 git

要增加或更新依赖，就使用 `go get` 命令：

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```

整理好 `go.mod` 和 `go.sum`：

```bash
go mod tidy
go mod vendor
git add go.mod go.sum vendor
git commit
```

直接提交 `go.mod` and `go.sum` 的修改。

# 赞助

你可以捐助或参与众筹来帮助我们维护服务器费用，以及提供一些奖金资助项目发展。

* **Support on [Open Collective](https://opencollective.com/go-admin)**
* Donate via [PayPal](https://paypal.me/cg80333)

# 成员

如果你已经是GoAdmin的官方开发组成员：

* [Discussions](http://forum.go-admin.cn)
* [Tasks](https://github.com/GoAdminGroup/go-admin/projects)
* [Chat](https://t.me/joinchat/NlyH6Bch2QARZkArithKvg)