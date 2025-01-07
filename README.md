# goravel 扩展包测试集合

## 创建一个扩展包
```bash
go run . artisan make:package <package_name>
```
```bash
go run . artisan make:package elastic
go run . artisan make:package spider
go run . artisan make:package rabbitmq
```

使用到第三方的扩展包
```bash

```

## 将扩展包独立发布

初始化一个github仓库
```bash
git init
```

## 发布
```bash
git tag v1.0.0
git push origin v1.0.0
```