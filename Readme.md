# Gin+React Blog

## BL 环境搭建

### 1.配置 docker 容器

```bash
docker run -itd --name mysql -p 3307:3306 -e  MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=blog_db -d mysql

docker run --name es -p 127.0.0.1:9200:9200 -e "discovery.type=single-node" -e "xpack.security.http.ssl.enabled=false" -e "xpack.license.self_generated.type=trial" -e "xpack.security.enabled=false" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -d elasticsearch:8.17.0

docker run --name redis -p 6379:6379 -d redis
```

### 2. Go 环境配置

```bash
// todo
```
