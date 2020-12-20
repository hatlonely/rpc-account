package service

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
)

func TestAccountService_GetCaptcha(t *testing.T) {
	Convey("TestAccountService_GetCaptcha", t, func() {
		_, err := service.GetCaptcha(context.Background(), &account.GetCaptchaReq{
			Email: "hatlonely@foxmail.com",
			Name:  "hatlonely",
		})
		So(err, ShouldBeNil)
	})
}
