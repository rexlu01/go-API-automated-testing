# go-API-automated-testing

## 1, 项目说明
### 1.1 一句话描述

go 实现针对Http(后续加grpc协议)接口的自动化测试和压测工具，把接口功能测试与接口压测结合起来

### 数据流转过程
![image](https://github.com/rexlu01/go-API-automated-testing/blob/main/dataflowprocess.png)


### 细节思路
* 利用Go-micro框架把用例生成、发送请求、产生并发、解析结果、整理压测数据及指标、结果展示等拆分成单独微服务
* 准备做一个UI界面，主要的想法是后续压测数据可以实时展示，形成曲线图
* 暴露给前端的API用Gin框架实现
* 压测逻辑里会加上sentinel-go的限流功能已实现比如：在指定QPS下执行压测的场景
* 过程数据用Redis，Mysql存储


### 过程中备注
* 发送请求的方法按照协议http or grpc分开写，写个接口统一起来，方便以后扩展。
* 过程数据用Redis and  Mysql存储
* 前端页面用micro web+gin, gin主要和前端HTML交互， micro web指定端口号（感觉不需要再暴露一层micro api --handler=http）
* 需要解决两个问题：1 压测要按时长跑（https://www.jb51.net/article/128532.htm） 2 html要实时展示性能指标折线图 （http://www.360doc.com/content/20/0419/09/33093582_906988529.shtml）
* web写定时任务，主要作用1，定时去拉标志位决定是否要拉数据，2，如果标志位是true就收集数据且展示在前端
* 生成接口文件普通：protoc --proto_path=. --go_out=. --micro_out=. */*.proto
* 需要研究一下TPS的计算方式，用总次数/总时间这样算不准（todo）

### 未完待续ing...

