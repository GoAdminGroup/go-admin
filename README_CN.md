<p align="center">
  <a href="https://github.com/chenhg5/go-admin">
    <img width="50%" alt="go-admin" src="https://ws2.sinaimg.cn/large/006tNc79ly1ftvqf8qeddj31bz07e40e.jpg">
  </a>
</p>
<p align="center">
    遗失的Golang语言编写的Web管理平台构建框架
</p>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/chenhg5/go-admin"><img alt="Go Report Card" src="https://camo.githubusercontent.com/59eed852617e19c272a4a4764fd09c669957fe75/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f6368656e6867352f676f2d61646d696e"></a>
</p> 

## 前言

对于一个管理平台来说，有几个东西是重要的：

- 安全性和易于使用
- 独立性，独立于业务系统

![](https://ws3.sinaimg.cn/large/006tNc79ly1ft048byoafj31kw0w847v.jpg)

## 特征

- 使用adminlte构建的漂亮的管理界面
- 可配置的，易于管理数据库数据
- 完善的认证系统
- 使用Go编写
- 可移植的
- 部署简单

## 环境要求

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## 安装

见 [wiki](https://github.com/chenhg5/go-admin/wiki)

## make 命令

- build
- test
- clean
- run
- restart
- deps : 安装依赖
- cross : 跨平台编译
- pages : 将html文件编译为go文件
- assets : 将静态文件编译为go文件
- fmt

## 技术支持

- [fasthttp](https://github.com/valyala/fasthttp)
- [adminlte](https://adminlte.io/themes/AdminLTE/index2.html)
- [hero](https://github.com/shiyanhui/hero)

## todo

- [x] 增加 [go-bindata](https://github.com/go-bindata/go-bindata) 支持
- [X] 增加更多表格表单组件
- [X] 菜单结构
- [ ] rcba认证
- [ ] 自定义页面
- [ ] 合并优化静态资源
- [ ] 自动安装引擎
- [ ] demo网站的搭建
- [ ] 性能分析

## 贡献

非常欢迎提pr，<strong>这里可以加入开发小组</strong>

QQ群: 756664859

## 十分感谢

inspired by [laravel-admin](https://github.com/z-song/laravel-admin)