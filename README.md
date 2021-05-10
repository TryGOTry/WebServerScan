# WebServerScan
一个简单的c段web服务探测工具
# 说明
根据自己平时的需求写了一个web服务探测工具(主要用于内网探测，但是你要扫公网，我也没意见)，可自定义端口，设置线程，设置超时时间，可返回容器名称，标题，返回页面字符长度等。

1.终端标题显示任务进度
2.可扫http和https,默认扫描http
3.默认端口80,8080
# 参数
```
Usage of Web-Server-Scan_x64.exe:
  -h string
        ip地址(c段)
  -p string
        扫描端口 (default "80,8080")
  -s int
        线程 (default 5)（请根据自己电脑配置来进行设置）
  -t int
        超时时间 (default 2)
  -type string
        扫描类型(http,https)默认http (default "http")
  -u string
        单个地址扫描
```
# 运行截图
## 扫描http
![运行截图](https://github.com/TRYblog/WebServerScan/blob/main/1.png)
## 扫描https
![运行截图](https://github.com/TRYblog/WebServerScan/blob/main/3.png)
## 关于作者
一个菜鸟.
[个人博客](https://www.nctry.com)
