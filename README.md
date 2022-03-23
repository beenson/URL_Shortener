# URL Shortener

## 環境依賴
Golang 1.17  
MariaDB 10.6.5  
Redis 6.2.6

## 建置及執行
```go
go mod download
go build -o API.exe
.\API.exe
```

## 使用套件
- [Fiber](https://github.com/gofiber/fiber)
- [GoDotEnv](https://github.com/joho/godotenv)
- [GORM](https://gorm.io/index.html)
- [go-redis](https://github.com/go-redis/redis)
- [Testify](https://github.com/stretchr/testify)
- [Validator](https://github.com/go-playground/validator)

## 問題與解決
Q:若未來有Load Balance需求，可能會在Create時有race condition情況發生  
A:使用[GORM Create Hooks](https://gorm.io/docs/create.html#Create-Hooks)避免，不過功能與想像有些落差，可能無法有效解決問題，未來朝向Locking方向進行研究  
  
Q:原設計想法為使用redis TTL機制判斷CODE是否已經過期，不過redis的TTL可能有不準確的問題  
A:為求精確，將過期的時間資訊一併存放到redis上，而TTL設置時間不會因為過期時間接近改變  
  
Q:因前項問題所倒置，CODE有可能無法在過期當下馬上從redis上移除，導致該CODE已經能夠使用，卻無法使用的情況發生  
A:尚未解決(在取得時檢查是否已經過期，並將其移除redis)  