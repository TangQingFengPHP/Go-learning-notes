### 

写一个「模拟表格分页的 CLI 程序」，核心功能就是：

* 模拟有很多行的表格数据

* 支持分页查看（指定页码、每页条数）

* 支持翻页（上一页、下一页）

* 支持跳转页数

### 示例效果（CLI）

```shell
$ go run main.go
[Page 1/10] Showing 10 of 100 rows:
1. Name: Alice       Age: 25
2. Name: Bob         Age: 30
...

[n] Next Page | [p] Previous Page | [q] Quit | [g 5] Go to page 5
> 
```

### 项目结构

```go
table-paginator/
├── main.go            // 启动入口
├── paginator/
│   └── paginator.go   // 分页逻辑
├── data/
│   └── mock.go        // 模拟表格数据
```