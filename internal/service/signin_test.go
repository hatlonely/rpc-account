package service

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
)

func TestGenerateToken(t *testing.T) {
	Convey("TestGenerateToken", t, func() {
		fmt.Println(GenerateToken())
	})
}

func TestAccountService_SignIn(t *testing.T) {
	Convey("TestAccountService_SignIn", t, func() {
		res, err := service.SignIn(context.Background(), &account.SignInReq{
			Username: "hatlonely@foxmail.com",
			Password: "12345678",
		})
		So(err, ShouldBeNil)
		So(len(res.Token), ShouldEqual, 32)
	})
}
