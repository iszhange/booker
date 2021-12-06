# gitbook-cli

列表出所有可用命令

```shell
$ gitbook help
```



**运行GitBook**

打包成静态网页，会在根目录下生成`_book`目录，将这个目录上传到你网站上即可，或者放到CDN上

```bash
$ gitbook build ./demo-book
```

运行本地服务，它会在`http://localhost:4000`起个web服务，可以让你打开书籍的预览

```bash
$ gitbook serve
```



**指定版本**

`gitbook-cli`默认是从配置文件读取版本，你可以通过`--gitbook`参数显式的指定版本

比如，生成静态页面时指定版本

```bash
$ gitbook build ./demo-book --gitbook=4.0.0
```

比如，查看某版本GitBook命令列表

```bash
$ gitbook help --gitbook=3.0.0
```



**管理版本**

查看本地已安装的版本

```bash
$ gitbook ls
```

查看`NPM`上可供安装的版本

```bash
$ gitbook ls-remote
```

安装指定版本

```bash
$ gitbook fetch 2.1.0
```

更新至最新版本

```bash
$ gitbook update
```

卸载指定版本

```bash
$ gitbook uninstall 2.0.1
```

