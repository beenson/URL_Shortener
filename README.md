# URL Shortener

## 環境依賴
Golang 1.17  
MariaDB 10.6.5  
Redis(預計使用)

## 建置及執行
```go
go mod download
go build -o API.exe
.\API.exe
```

## 使用套件
- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/index.html)
- [GoDotEnv](https://github.com/joho/godotenv)
- [Testify](https://github.com/stretchr/testify)
- [Validator](https://github.com/go-playground/validator)

## 問題與解決
Q:若未來有Load Balance需求，可能會在Create時有race condition情況發生  
A:目前尚未解決(打算使用[GORM Create Hooks](https://gorm.io/docs/create.html#Create-Hooks)來避免)