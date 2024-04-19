# go-jwt-auth
## Запуск
```
docker-compose up -d
go install
go run main.go
```
## Ручки (localhost:8080)
```
POST("/login", controllers.Login)
POST("/signup", controllers.Signup)
POST("/refresh, controllers.Refresh)
GET("/home", middlewares.IsAuthorized(), controllers.Home)
GET("/premium", controllers.Premium)
GET("/me", middlewares.IsAuthorized(), controllers.Me)
PATCH("/me", middlewares.IsAuthorized(), controllers.Update)
```
