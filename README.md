# gin-blog

## 安装
```
git clone https://github.com/qiuapeng921/gin-blog $GOPATH/src/gin-blog
```

## 如何运行

### 必须

- Mysql
- Redis

### 运行
```
cd $GOPATH/src/gin-blog
cp .env.example .env

###修改.env配置
DB_HOST=127.0.0.1:3306
DB_DATABASE=test
DB_USERNAME=root
DB_PASSWORD=
DB_CHARSET=utf8mb4
DB_MAX_IDLE=20
DB_MAX_OPEN=100

REDIS_HOST=127.0.0.1:6379
REDIS_PASSWORD=null
REDIS_MAX_IDLE=50
REDIS_MAX_ACTIVE=1200

MONGO_HOST=127.0.0.1:27017
MONGO_USERNAME=null
MONGO_PASSWORD=null
MONGO_MAX_POOL=20

ELASTIC_HOST=http://127.0.0.1:9200

go run main.go 
```

## 特性
- Gin
- Gorm
- Mongo
- Redis
- ElasticSearch
- RabbitMq
- Cron
- Socket
- Templates
