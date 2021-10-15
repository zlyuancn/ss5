# socks5 服务

# 自行构建

1. 要求golang1.16及以上环境
2. 克隆代码后, 运行 `go build -o ss5 .`
3. 运行 `./ss5 [-addr :1080] [-user 用户名 -pwd 密码]`

# docker部署

```docker
docker run -d -p 1080:1080 zlyuan/ss5
```

```docker
docker run -d -p 1080:1080 zlyuan/ss5 -addr :1080 -user admin -pwd 12345
```
