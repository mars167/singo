# 任务简报


```
# 1. 启动容器
cd devops/dev 
docker-compose up -d 
# 2. 安装这个依赖
go mod init go-crud
# 3. 启动项目
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装  
```


