# 发布电子书



<!-- toc -->



## 封面

封面用于所有电子书格式

封面必需提供一大一小，放在根目录下，命名为`cover.jpg`和`cover_small.jpg`

好的封面应该遵守以下准则：

1. `cover.jpg`大小为1800x2360px，`cover_small.jpg`大小为200x262px
2. 没有边界
3. 清晰可见的书名
4. 任何重要的文字应该在小版本中可见

你可以使用以下方式来添加封面:

1. 提花封面图片放于书籍根目录下

2. 使用`autocover`插件生成

   > [!NOTE]
   >
   > 这个插件依赖于`canvas`，安装比较麻烦

   ```json
   {
   	"plugins": ["autocover"],
   	"pluginsConfig": {
   		"autocover": {
   			"font": {
                   "size": null,
                   "family": "Impact",
                   "color": "#FFF"
               },
               "size": {
                   "w": 1800,
                   "h": 2360
               },
               "background": {
                   "color": "#09F"
               }
   		}
   	}
   }
   ```

## 托管到GitBook



## 托管到GitHub



## 托管到GitHub Pages



## 生成电子书

可以将GitBook打包成电子书，支持的格式有：`pdf`、`epub`、`mobi`

**安装`ebook-convert`**

从[Calibre](https://calibre-ebook.com/download)下载适用你平台的版本安装

**打包电子书**

```bash
# Generate a PDF file
$ gitbook pdf . mybook.pdf

# Generate an ePub file
$ gitbook epub . mybook.epub

# Generate a Mobi file
$ gitbook mobi . mybook.mobi
```

