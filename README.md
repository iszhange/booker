# Booker

中文名：GitBook管理者，作用是用于自动构建和发布书籍的程序

## 实现原理

![booker-schematic-diagram](https://book-docs.oss-cn-hongkong.aliyuncs.com/20221128/booker-schematic-diagram.png)

图中包含两个流程：一是虚线表示的webhook；二是实现表示的GitBook发布

### WebHook流程

当push后，github会调用webhook地址，nginx在这里会将请求转发到5454端口

### gitbook发布流程

booker会起个进程监听5454，收到请求后，会校验secret是否正确，验证成功后会拉取更新，增量构建书籍

构建完成后，会将更新文件拷贝到共享目录，通过nginx提供web服务

## 使用技术

1. Docker
2. Golang
3. Nginx

## 镜像构建

### 构建参数

|名称|类型|描述|  
| ---- | ---- | ---- | ---- |  
|IS_CHINA|bool| 是否国内 `true`是 `false`否，国内的话会使用镜像 |

### 环境变量

|名称|类型|描述|  
| ---- | ---- | ---- | ---- |  
|REPOSITORY_DIR|string| github仓库目录 |  
|BOOKS_DIR|string| 构建后book目录 |  
|CONFIG_DIR|string| 配置文件目录 |  

### 构建镜像

注：开发环境使用`.devcontainer`下的Dockerfile

```sh
docker build --build-arg IS_CHINA=true -t gitbook:v1 .
```

## 配置文件

booker的配置都在`config.yml`中

|名称|类型|描述|  
| ---- | ---- | ---- | ---- |  
|repository|string| `*` github仓库地址 |  
|secret|string| `*` webhook密钥 |  