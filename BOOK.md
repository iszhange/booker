# 如何创建第一本书？



<!-- toc -->



## 初始化

使用这个命令将当前目录初始化为样板书，你当然也可以指定别的目录

```bash
$ gitbook init .
```



## GitBook项目结构

```markdown
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



## book.json

| 配置项        | 描述                                                         |
| ------------- | ------------------------------------------------------------ |
| `root`        | 包含图书所有文件的根目录，`book.json`可以不在根目录中        |
| `structure`   | 指定自述文件、摘要、词汇表等的路径                           |
| `title`       | 书名；默认值是从README中提取出来的；在gitbook.com上，这个字段是预填的 |
| `description` | 书籍描述；默认值是从README中提取出来的；在gitbook.com上，这个字段是预填的 |
| `author`      | 作者；在gitbook.com上，这个字段是预填的                      |
| `isbn`        | 国际标准书号 ISBN                                            |
| `language`    | 语言类型，[ISO code](https://www.html.am/reference/iso-language-codes.cfm) |
| `direction`   | 文本阅读顺序，`rtl`|`ltr`，默认值依赖于`language`的值        |
| `gitbook`     | 使用的GitBook版本，使用[SemVer](https://semver.org/lang/zh-CN/)规范，接受`><=`条件判断 |



## SUMMARY

用来定义本书的章节和子章节的结构，生成目录

`SUMMARY.md`的结构是一个超链接的列表，链接的标题为章节标题，链接目标为章节文件的(相对)路径

用列表嵌套来创建子章节

**示例**

```markdown
# Summary

* [Part I](part1/README.md)
    * [Writing is nice](part1/writing.md)
    * [GitBook is nice](part1/gitbook.md)
* [Part II](part2/README.md)
    * [We love feedback](part2/feedback_please.md)
    * [Better tools for authors](part2/better_tools.md)
```

**锚点**

使用锚点可以将目录指向文件中特定部分

```markdown
# Summary

### Part I

* [Part I](part1/README.md)
    * [Writing is nice](part1/README.md#writing)
    * [GitBook is nice](part1/README.md#gitbook)
* [Part II](part2/README.md)
    * [We love feedback](part2/README.md#feedback)
    * [Better tools for authors](part2/README.md#tools)
```

**分隔**

目录可以使用标题或`----`来分隔

```markdown
# Summary

### Part I

* [Writing is nice](part1/writing.md)
* [GitBook is nice](part1/gitbook.md)

### Part II

* [We love feedback](part2/feedback_please.md)
* [Better tools for authors](part2/better_tools.md)

----

* [Last part without title](part3/title.md)
```

