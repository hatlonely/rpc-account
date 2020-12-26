package service

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
)

func TestAccountService_GetCaptcha(t *testing.T) {
	Convey("TestAccountService_GetCaptcha", t, func() {
		_, err := service.GetCaptcha(context.Background(), &api.GetCaptchaReq{
			Email: "hatlonely@foxmail.com",
			Name:  "hatlonely",
		})
		So(err, ShouldBeNil)
	})
}
