package configuration

import (
 "testing"

 . "github.com/smartystreets/goconvey/convey"
)

func TestConfiguration(t *testing.T) {
 Convey("configuration", t, func() {
  Convey("ReadConfiguration_呼叫時_應讀取設定檔", func() {
   ReadConfiguration()

   So(len(Connectionstring), ShouldBeGreaterThan, 0)
  })
 })
}