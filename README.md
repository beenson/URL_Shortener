# URL Shortener
本專案為產生並管理具時效性的短網址  
使用者傳入欲縮短的網址以及過期時間，將回傳一個可用的短網址，在過期時限內使用該網址即可導向到註冊時所輸入的網址  
詳細API參數與路徑請先執行程式後查看[/swagger/index.html](http://localhost/swagger/index.html)

## 環境依賴
Golang 1.17  
MariaDB 10.6.5  
Redis 6.2.6

## 建置及執行
將.env.example複製為.env並設定參數再執行以下指令  
```go
go mod download
go build -o API.exe
.\API.exe
```

## 使用套件
- [Fiber](https://github.com/gofiber/fiber) : web framework
- [fiber-swagger](https://github.com/arsmn/fiber-swagger) : 製作API說明文件
- [GoDotEnv](https://github.com/joho/godotenv) : 讀取.env檔案
- [GORM](https://gorm.io/index.html) : 與database互動，設定簡單易懂
- [go-redis](https://github.com/go-redis/redis) : 與redis互動
- [Testify](https://github.com/stretchr/testify) : 測試使用
- [Validator](https://github.com/go-playground/validator) : 確認輸入格式正確
- [swag](https://github.com/swaggo/swag) : 製作API說明文件

## 檔案結構
```
│  .env.example             // .env範例檔
│  .env.test                // 測試環境
│  .gitignore
│  go.mod
│  go.sum
│  main.go                  // 主程式
│  main_test.go             // 測試
│  README.md                // readme
│
├─app                       // 存放商業邏輯
│  ├─controllers            // controllers
│  │      url_shorten.go   
│  │
│  └─models                 // models
│          requests.go
│          responses.go
│          shorten.go
│
├─docs                      // swagger文件
│      docs.go
│      swagger.json
│      swagger.yaml
│
├─pkg                       // 存放專案功能
│  ├─migrate                // migrations
│  │      migrate.go
│  │
│  ├─repository             // 存放固定變數
│  │      errors.go
│  │      variables.go
│  │
│  ├─routes                 // routers
│  │      not_found.go
│  │      routes.go
│  │      swagger.go
│  │
│  └─utils                  // utils
│          code_generate.go
│          initialize.go
│          parse.go
│
└─service                   // 使用外部服務
    ├─cache                 // redis相關
    │      connection.go    // redis連線
    │
    └─database              // database相關
            connection.go   // database連線
```

## 問題與解決
Q:若未來有Load Balance需求，可能會在Create時有race condition情況發生  
A:原使用[GORM Create Hooks](https://gorm.io/docs/create.html#Create-Hooks)避免，但功能與想像有些落差，無法解決。  
現在解決方法是在新增後再次確認現階段該CODE在資料庫是否唯一。  
  
Q:原設計想法為使用redis TTL機制判斷CODE是否已經過期，不過redis的TTL可能有不準確的問題  
A:為求精確，將過期的時間資訊一併存放到redis上，而TTL設置時間不會因為過期時間接近改變  
  
Q:因前項問題所倒置，CODE有可能無法在過期當下馬上從redis上移除，導致該CODE已經能夠使用，卻無法使用的情況發生  
A:在取得時檢查是否已經過期，若發現已過期將其從redis移除  
  
Q:有可能有大量存取不存在的shorten URL  
A:目前不打算特別進行處理，若將資訊儲存於redis中，可能會有極大量資訊存放於redis中(如掃描CODE)，進而影響真正有需求的使用者效能  
  
Q:特定IP進行掃描CODE  
A:尚未實做(使用middleware，記錄特定IP時間內的請求次數，暫時不提供服務給短時間大量請求的IP)  

## 心得
很開心有機會參加這一次的活動，原本看到題目時，以為不會有什麼問題，不過等到真正實作時才開始想到有可能出現的問題，經驗沒有很豐富的我，也是花了好多力氣解決大大小小的問題。  
本次實做問題最大的大概就是關於unit test了，之前上課老師雖然有提過相關方面，不過當時是使用視窗程式做教學，雖然了解testing的重要性，不過卻沒有實做API unit test的經驗，倒置我做的test已經算是API testing等級了，希望未來有機會可以再做的更好。