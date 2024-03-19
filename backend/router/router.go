package router

import (
	// "os"
	// "path/filepath"
	"test3/api"
	// "test3/configuration"
	// "test3/websocket"
	// "strings"

	// jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/compress"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/recover"
	// "github.com/golang-jwt/jwt/v4"
)

const (
	bodyLimit = 1 * 1024 * 1024 // 上傳大小限制1MB
)

var HttpApplication *fiber.App

// 監聽網頁服務
func Run() {
	HttpApplication = fiber.New(fiber.Config{BodyLimit: bodyLimit, UnescapePath: true, ErrorHandler: notFoundHandler})

	// HttpApplication.Use(compress.New()) // 啟用壓縮
	// HttpApplication.Use(recover.New())  // 啟用錯誤處理
	// HttpApplication.Use(logger.New(logger.Config{
	// 	TimeFormat: "2006/01/02 15:04:05",
	// 	TimeZone:   "Asia/Taipei",
	// })) // 啟用日誌

	setupRoute() // 設定路由

	_ = HttpApplication.Listen(":61018")
}

// 未找到路徑轉到首頁
func notFoundHandler(context *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// 若 未找到路徑 則 回到首頁
	if code == fiber.StatusNotFound {
		return context.Redirect("/")
	}

	return context.Status(code).SendString(err.Error())
}

// 設定路由
func setupRoute() {
	HttpApplication.Get("/api/calendar", api.GetCalendar)           // 取得行事曆
	HttpApplication.Post("/api/calendar", api.AddCalendar)          // 新增行事曆
	HttpApplication.Put("/api/calendar/:id", api.UpdateCalendar)    // 更新行事曆
	HttpApplication.Delete("/api/calendar/:id", api.DeleteCalendar) // 刪除行事曆

	// apiGroup := HttpApplication.Group("/api")

	// 要求授權
	// apiGroup.Use(jwtware.New(jwtware.Config{
	// 	SigningKey:     configuration.JWTKey,
	// 	SuccessHandler: jwtSuccessHandler, // 權杖驗證成功後檢查授權
	// }))

	// apiGroup.Get("/product", api.GetTeachers)                // 取得老師
	// apiGroup.Post("/product", api.AddTeacher)                // 新增老師
	// apiGroup.Put("/product/:id", api.UpdateTeacher)          // 修改老師
	// apiGroup.Delete("/product/:id", api.ChangeTeacherStatus) // 刪除老師
}

// 權杖驗證成功後檢查授權
// func jwtSuccessHandler(context *fiber.Ctx) error {
// 	claims := context.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

// 	// 若 沒有老師主鍵或電子郵件 則 回未授權
// 	if claims["id"] == nil || claims["email"] == nil {
// 		return context.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// 若 路徑為管理老師 且 不為管理者 則 回未授權
// 	if strings.HasPrefix(context.Path(), "/api/teacher") && claims["admin"] == nil {
// 		return context.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	context.Locals("id", claims["id"])       // 老師主鍵
// 	context.Locals("email", claims["email"]) // 老師電子郵件

// 	return context.Next()
// }
