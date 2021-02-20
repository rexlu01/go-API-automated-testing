# go-API-automated-testing

## 1, 项目说明
### 1.1 一句话描述

go 实现针对Http接口的自动化测试和压测工具，把接口功能测试与接口压测结合起来

### 细节思路
* 利用Go-micro框架把用例生成、发送请求、产生并发、解析结果、整理压测数据及指标、结果展示等拆分成单独微服务
* 准备做一个UI界面，主要的想法是后续压测数据可以实时展示，形成曲线图
* 暴露给前端的API用Gin框架实现
* 压测逻辑里会加上sentinel-go的限流功能已实现比如：在指定QPS下执行压测的场景

### 未完待续ing...

