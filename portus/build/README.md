镜像：

```bash
go build ../start.go
docker build -t portusapi:latest .

```

启动：
```bash
docker-compose up -d
```

配置:

启动时添加如下环境变量可切换数据库:

DB_HOST          string
DB_PORT          string
DB_DATABASE      string
DB_USERNAME      string
DB_PASSWORD      string
