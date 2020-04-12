package errorx

import (
	"fmt"
	"testing"

	"github.com/goroom/utils/exit"
	"github.com/smartystreets/goconvey/convey"
)

func TestError(t *testing.T) {
	convey.Convey("TestXYError OK", t, func() {
		err := ErrorMsg("test error inter")
		err = ErrorWithMsg(err, "test error outer")
		err = Error(err)
		if err != nil {
			fmt.Println("show: ", err)
		}
	})

	convey.Convey("TestXYError ExitError", t, func() {
		err := Error(nil)
		exit.ExitError(err)
		fmt.Println("ExitError should not exit")
	})

}
