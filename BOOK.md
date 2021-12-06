# 如何创建第一本书？



<!-- toc -->



## 初始化

使用这个命令将当前目录初始化为样板书，你当然也可以指定别的目录

```bash
$ gitbook init .
```



## GitBook项目结构

```bash
mybook/
  |-- book.json
  |-- README.md
  |-- SUMMARY.md
  |-- chapter-1/
  |     |-- README.md
  |     |-- something.md
  |-- chapter-2/
  |     |-- README.md
  |     |-- something.md
```

**基础功能文件**

- `book.json` 配置数据 (O)
- `README.md` 书籍前言 (O)
- `SUMMARY.md` 书籍目录 (R)
- `GLOSSARY.md` 词汇/注释术语列表 (O)



## 静态文件与图片

静态文件是指未在`SUMMARY.md`中列出的文件与未被忽略的文件。

构建时，所有的静态文件都会被复制到输出路径



## 忽略文件与目录

`GitBook`会读取`.gitignore`、`.bookignore`和`.ignore`文件

这些文件都遵循gitignore规则





## SUMMARY

