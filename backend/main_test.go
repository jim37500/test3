package main

import (
	"log"
	"os"
	"reflect"
	"sync"
	"testing"

	"test3/database"
	"test3/router"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Main(t *testing.T) {
	Convey("main", t, func() {
		Convey("main 呼叫時 應監聽網頁服務", func() {
			os.Args = []string{}
			waitGroup := sync.WaitGroup{}
			waitGroup.Add(1)
			myPatches := NewPatches()
			defer myPatches.Reset()
			isRun := false
			myPatches.
				ApplyFunc(database.Open, func() {}).
				ApplyFunc(router.Run, func() {
					isRun = true
					waitGroup.Done()
				})

			go main()
			waitGroup.Wait()

			So(isRun, ShouldBeTrue)
		})

		Convey("main 當服務執行失敗時 會 印出錯誤訊息", func() {
			os.Args = []string{"test", "test"}
			waitGroup := sync.WaitGroup{}
			waitGroup.Add(1)
			myPatches := NewPatches()
			defer myPatches.Reset()
			isRun := false
			isPrintRun := false
			myPatches.
				ApplyFunc(database.Open, func() {}).
				ApplyFunc(log.Println, func(v ...any) {
					isPrintRun = true
					waitGroup.Done()
				}).
				ApplyFunc(router.Run, func() {
					isRun = true
					waitGroup.Done()
				})

			go main()
			waitGroup.Wait()

			So(isRun, ShouldBeFalse)
			So(isPrintRun, ShouldBeTrue)
		})

		Convey("main 當停止時 會 停止", func() {
			router.HttpApplication = fiber.New(fiber.Config{})
			waitGroup := sync.WaitGroup{}
			waitGroup.Add(1)
			myPatches := NewPatches()
			defer myPatches.Reset()
			isShutdown := false
			myPatches.
				ApplyMethod(reflect.TypeOf(router.HttpApplication), "Shutdown", func() error {
					isShutdown = true
					waitGroup.Done()
					return nil
				})

			myProgram := program{}
			_ = myProgram.Stop(nil)
			waitGroup.Wait()

			So(isShutdown, ShouldBeTrue)
		})
	})
}
