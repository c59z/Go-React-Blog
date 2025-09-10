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
go get "gopkg.in/yaml.v3"

go get "go.uber.org/zap"
go get "github.com/natefinch/lumberjack"

go get "github.com/fvbock/endless"

go get "github.com/gin-gonic/gin"

go get "gorm.io/gorm"
go get "gorm.io/driver/mysql"

go get "github.com/go-redis/redis"

go get "github.com/elastic/go-elasticsearch/v8"

go get "github.com/robfig/cron/v3"

go get "github.com/songzhibin97/gkit/cache/local_cache"

go get "github.com/gofrs/uuid"
```
