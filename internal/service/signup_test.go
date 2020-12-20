package service

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	account "github.com/hatlonely/go-rpc/rpc-account/api/gen/go/api"
	"github.com/hatlonely/go-rpc/rpc-account/internal/storage"
)

func TestAccountService_SignUp(t *testing.T) {
	Convey("TestAccountService_SignUp", t, func() {
		mysqlCli.Delete(&storage.Account{Email: "hatlonely@foxmail.com"})
		redisCli.Set("captcha_hatlonely@foxmail.com", "041736", 5*time.Second)

		_, err := service.SignUp(context.Background(), &account.SignUpReq{
			Email:    "hatlonely@foxmail.com",
			Phone:    "13810242048",
			Name:     "hatlonely",
			Password: "12345678",
			Birthday: "1992-01-01",
			Gender:   account.Gender_Male,
			Avatar:   "http://avatar/hatlonlely.png",
			Captcha:  "041736",
		})
		So(err, ShouldBeNil)
	})
}
