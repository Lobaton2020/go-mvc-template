1. `go mod init github.com/Lobaton2020/go-mvc-template`
2. `touch main.go`
3. Install el nodemon de GO
    `go get github.com/githubnemo/CompileDaemon`
    `go install -mod=mod github.com/githubnemo/CompileDaemon`
    `CompileDaemon --command="./go-mvc-template"`
4. Install orm GORM
    `go get -u gorm.io/gorm`
    `go get -u gorm.io/driver/postgres`
5. Instar dotenv lib
    `go get github.com/joho/godotenv`
6. Install Fiber
    `go get -u github.com/gofiber/fiber/v2`