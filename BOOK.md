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

理论上来说`book.json`可以放在你电脑上的任何位置，对于你本地运行或打包没有任何问题；但如果你通过`GitBook.com`来发布，还是老老实实放到根目录下吧

| 配置项        | 描述                                                         |
| ------------- | ------------------------------------------------------------ |
| `root`        | 包含图书所有文件的根目录，`book.json`可以不在根目录中        |
| `structure`   | 指定自述文件、摘要、词汇表等的路径                           |
| `title`       | 书名；默认值是从README中提取出来的；在gitbook.com上，这个字段是预填的 |
| `description` | 书籍描述；默认值是从README中提取出来的；在gitbook.com上，这个字段是预填的 |
| `author`      | 作者；在gitbook.com上，这个字段是预填的                      |
| `isbn`        | 国际标准书号 ISBN                                            |
| `language`    | 语言类型，[ISO code](https://www.html.am/reference/iso-language-codes.cfm) |
| `direction`   | 文本阅读顺序，`rtl`                                          |
| `gitbook`     | 使用的GitBook版本，使用[SemVer](https://semver.org/lang/zh-CN/)规范，接受`><=`条件判断 |
| `links`       | 链接信息                                                     |
| styles        | 自定义页面样式                                               |
| plugins       | 要加载的插件列表                                             |
| pluginsConfig | 插件的配置                                                   |

**root**

明确指明书籍根目录

```json
{
	"root": "./mybook"
}
```

**structure**

指定`readme`、`summary`、`glossary`、`languages`对应的文件名

```json
{
	"structure": {
		"readme": "./mybook/README.md",
        "summary": "./mybook/SUMMARY.md",
        "glossary": "./mybook/GLOSSARY.md",
        "languages": "./mybook/LANGS.md"
	}
}
```

**title**

```json
{
	"title": "GitBook使用指南"
}
```

**description**

```json
{
	"description": "本书是关于GitBook使用教程"
}
```

**author**

```json
{
	"author": "张三"
}
```

**isbn**

```json
{
	"isbn": "ISBN 7-309-04547-5"
}
```

**language**

大多数国家语言都支持，常用的也就几个

`en, zh-hans, zh-tw`

```json
{
	"language": "zh-hans"
}
```

**direction**

一般来说默认就行，只有两种取值`ltr`或`rtl`

```json
{
	"direction": "ltr"
}
```

**gitbook**

```json
{
	"gitbook": ">=3.2.0"
}
```

**links**

添加链接信息

好像只能在左侧导航栏添加

```json
{
	"links": {
		"sidebar": {
			"GitHub": "https://github.com"
		}
	}
}
```

**styles**

```json
"styles": {
    "website": "styles/website.css",
    "ebook": "styles/ebook.css",
    "pdf": "styles/pdf.css",
    "mobi": "styles/mobi.css",
    "epub": "styles/epub.css"
}
```

**plugins**

1. 新增插件

   将插件名添加到`plugins`数组里，然后执行`gitbook install`来安装新插件

   ```json
   {
       "plugins": [
       	"splitter"
   	]
   }
   ```

2. 去除自带插件

   GitBook默认带有以下插件:

   - highlight
   - search
   - sharing
   - font-settings
   - livereload

   ```json
   {
   	"plugins": [
   		"-search"
   	]
   }
   ```

**pluginsConfig**

配置插件的配置项

```json
{
	"plugins": [
		"github"
	],
	"pluginsConfig": {
		"github": {
			"url": "https://github.com"
		}
	}
}
```



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

## GLOSSARY

添加术语及其定义，GitBook会自动构建索引并突出显示这些术语

`GLOSSARY.md`的格式为`h2`及描述段落的列表

```markdown
## Term
Definition for this term

## Another term
With it's definition, this can contain bold text
and all other kinds of inline markup ...
```



## 页面

**前言**

页面可以包含一个可选前言，用于定义页面的描述，必需放在页面开头，使用上下两条虚线包裹，采用`YAML`格式

```markdown
---
description: This is a short description of my page
---

# The content of my page
...
```

**正文**

GitBook推荐使用`Markdown`语法，就不要浪费时间搞别的语法了

```markdown
# Title of the chapter

This is a great introduction.

## Section 1

Markdown will dictates _most_ of your **book's structure**

## Section 2

...
```

