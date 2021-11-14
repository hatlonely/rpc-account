package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/pkg/errors"
)

func GenerateCaptcha() string {
	buf := make([]byte, 8)
	_, _ = rand.Read(buf)
	return fmt.Sprintf("%06d", binary.LittleEndian.Uint64(buf)%1000000)
}

func (s *AccountService) GetCaptcha(ctx context.Context, req *api.GetCaptchaReq) (*api.Empty, error) {
	captcha, err := s.cache.GetOrSetCaptcha(ctx, req.Email)
	if err != nil {
		return nil, errors.WithMessagef(err, "cache.GetOrSetCaptcha failed")
	}

	buf := &bytes.Buffer{}
	if err := s.captchaEmailTpl.Execute(buf, map[string]interface{}{
		"name":    req.Name,
		"captcha": captcha,
	}); err != nil {
		return nil, err
	}

	if err := s.emailCli.Send(req.Email, "验证码", buf.String()); err != nil {
		return nil, errors.Wrapf(err, "email [%v] send failed", req.Email)
	}

	return &api.Empty{}, nil
}

var captchaTpl = `<html>
<style>
    body {
        background-color: #fafafa;
    }
    .paper {
        background-color: #fff;
        width: 400px;
        border: 1px solid rgba(0, 0, 0, 0.12);
        border-radius: 8px;
        padding: 20px;
        margin: auto;
    }
    .captcha {
        font-weight: bold
    }
</style>
<body>
    <div class="paper">
        <p>
            您好，如果 {{.name}} 不是您的 hpifu 账户，请不要点击此邮件中的任何内容！
        </p>
        <p>
            以下是您的验证码：
        </p>
        <p class="captcha">
            {{.captcha}}
        </p>
        <p>
            {{.name}}，您好！
        </p>
        <p>
            我们收到了来自您的 hpifu 账户的安全请求。请使用上面的验证码验证您的账号归属。
        </p>
        <p>
            请注意：该验证码将在10分钟后过期，请尽快验证！
        </p>
        <p>
            享受您的历险！
        </p>
        <p>
            hpifu 客服团队
        </p>
    </div>
</body>
</html>`
