# Campop

##### 概述
练手 Go, 抓取器材到 Elasticsearch (6) 然后做全文检索, 用 CLI 的方式展示, 数据来源 [xitek](http://production.xitek.com/), 功能不通用. 本来想着批量索引, 不过数据量就一点也不需要了. 后续可以测试一下.

##### 主要依赖

- goquery
- Elasticsearch([olivere/elastic](https://github.com/olivere/elastic))
- `db/` 测试 [standalone_migrations](https://github.com/thuss/standalone-migrations) 不影响项目运行
- `go.mod`, `go.sum` 是测试 vgo

##### 运行

Elasticsearch - http://127.0.0.1:9200

```
go get github.com/PuerkitoBio/goquery
go get github.com/olivere/elastic
go build
./campop
./campop setup
./campop search -q "24 70"
...
```

![](http://wx4.sinaimg.cn/large/62fdd4d5gy1ftwwf5ejvqj227y18on5n.jpg)

![](http://wx2.sinaimg.cn/large/62fdd4d5gy1ftwwf4prk7j227y17ugz6.jpg)

##### TODO

- 写 test
- 优化搜索
- 更多 cli 操作
- api