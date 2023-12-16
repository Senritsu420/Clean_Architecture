package router

import (
	"app/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	t := e.Group("/tasks")
	// /tasksエンドポイントにミドルウェア追加
	t.Use(echojwt.WithConfig(echojwt.Config{
		// jwtを生成した時と同じSECRET_KEYを指定
		SigningKey: []byte(os.Getenv("SECRET")),
		// Clientから送られてくるjwtトークンの置き場所を指定
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
