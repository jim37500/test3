package database

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"test3/configuration"
	"test3/model"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Open(t *testing.T) {
	Convey("open", t, func() {
		configuration.Connectionstring = "FakeConnectionString"

		Convey("Open_呼叫時_應回依設定檔開啟資料庫", func() {
			mydb := &gorm.DB{}
			db := &sql.DB{}
			isFirst := true

			myPatches := NewPatches()
			defer myPatches.Reset()

			myPatches.ApplyFunc(gorm.Open, func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
				So(dialector.(*mysql.Dialector).Config.DSN, ShouldEqual, configuration.Connectionstring)
				if isFirst {
					isFirst = false
					return nil, errors.New("Fake Error")
				}
				return mydb, nil
			})
			myPatches.ApplyMethod(reflect.TypeOf(mydb), "DB", func(*gorm.DB) (*sql.DB, error) {
				return db, nil
			})
			myPatches.ApplyMethod(reflect.TypeOf(db), "SetConnMaxLifetime", func(_ *sql.DB, duration time.Duration) { So(duration, ShouldEqual, time.Hour) })
			myPatches.ApplyFunc(model.AutoMigrate, func(*gorm.DB) { So(db, ShouldEqual, db) })

			Open()
		})
	})
}
