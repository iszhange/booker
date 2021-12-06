# 安装



## 环境需求

- NodeJS(v12.*)

  对于gitbook(3.2.3)，Node高于12会报错

## 通过npm安装

```bash
$ npm install gitbook-cli -g
```

`gitbook-cli`是GitBook的命令行工具

执行下面命令，查看GitBook版本(会同时安装GitBook)，以验证是否安装成功

```bash
$ gitbook -V
```

## 安装指定版本

列出可用版本

```bash
$ gitbook ls-remote
```

安装指定版本

```bash
$ gitbook fetch 4.0.0
```

更新至最新版本

```bash
$ gitbook update
```

卸载指定版本

```bash
$ gitbook uninstall 4.0.0
```







## 参考

1. [gitbook-cli](https://github.com/GitbookIO/gitbook-cli)
