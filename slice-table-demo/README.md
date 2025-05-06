Go + Web 的通用分页 + 泛型过滤 + 可视化表格

* ✅ 任意结构体类型的数据切片

* ✅ 泛型过滤（支持传入条件函数）

* ✅ 分页（页码 + 每页大小）

* ✅ Web UI 表格展示（用 Go Serve HTML）

### 目录结构

```go
slice-table-demo/
├── main.go
├── data.go       // 模拟数据和数据结构
├── pagination.go // 泛型分页 + 过滤
├── templates/
│   └── index.html
```

