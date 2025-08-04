# go-exp-demo
使用wails3编写了一个exp的框子，后续写什么利用工具直接添加检测利用代理即可

# 安装wails
```
go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest
```
# 添加exp

## 在exp目录下创建

![](https://changge1001.oss-cn-nanjing.aliyuncs.com/img/25-03/20250728142656.png)

## 函数注册

![](https://changge1001.oss-cn-nanjing.aliyuncs.com/img/25-03/20250728142745.png)

## poc&exp检测编写

![](https://changge1001.oss-cn-nanjing.aliyuncs.com/img/25-03/20250728142949.png)

## 写完调式打包

```
wails3 dev   
wails3 build
```

