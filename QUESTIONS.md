# 遇到的问题



1. `toc`插件生成的目录无法跳转

   在`_book/gitbook/theme.js`中，查找

   ```javascript
   if(m)for(n.handler&&
   ```

   找到后，将`if(m)`改为`if(false)`

