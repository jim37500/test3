package router

import (
	"errors"
	"net"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Router(t *testing.T) {
	Convey("Router", t, func() {
		Convey("當呼叫Run時 應 監聽61018", func() {
			go Run()

			for {
				time.Sleep(1000 * time.Millisecond)
				if HttpApplication != nil && HttpApplication.Server() != nil {
					break
				}
			}

			connection, err := net.Dial("tcp", ":61018")
			defer func() {
				if connection != nil {
					connection.Close()
				}
				if HttpApplication != nil {
					_ = HttpApplication.Shutdown()
				}
			}()

			So(err, ShouldBeNil)
		})

		Convey("當 未找到路徑時 應 轉到首頁", func() {
			request := httptest.NewRequest("GET", "/notfound", nil)

			response, _ := fiber.New(fiber.Config{ErrorHandler: notFoundHandler}).Test(request)
			defer response.Body.Close()

			So(response.StatusCode, ShouldEqual, fiber.StatusFound)
			So(response.Header.Get("Location"), ShouldEqual, "/")
		})

		Convey("當 未找到路徑且有錯誤訊息時 應 回傳錯誤訊息", func() {
			application := fiber.New(fiber.Config{ErrorHandler: notFoundHandler})
			application.Get("/error", func(c *fiber.Ctx) error { return errors.New("error") })

			request := httptest.NewRequest("GET", "/error", nil)
			response, _ := application.Test(request)
			defer response.Body.Close()

			So(response.StatusCode, ShouldEqual, fiber.StatusInternalServerError)
		})
	})
}
