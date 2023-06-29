package middleware

import (
    "fmt"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/hd2yao/blog/global"
    "github.com/hd2yao/blog/pkg/app"
    "github.com/hd2yao/blog/pkg/email"
    "github.com/hd2yao/blog/pkg/err_code"
)

func Recovery() gin.HandlerFunc {
    defaultMailer := email.NewEmail(&email.SMTPInfo{
        Host:     global.EmailSetting.Host,
        Port:     global.EmailSetting.Port,
        IsSSL:    global.EmailSetting.IsSSL,
        UserName: global.EmailSetting.UserName,
        Password: global.EmailSetting.Password,
        From:     global.EmailSetting.From,
    })
    return func(context *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)

                err := defaultMailer.SendMail(
                    global.EmailSetting.To,
                    fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
                    fmt.Sprintf("错误信息: %v", err),
                )
                if err != nil {
                    global.Logger.Panicf("mail.SendMail err: %v", err)
                }

                app.NewResponse(context).ToErrorResponse(err_code.ServerError)
                context.Abort()
            }
        }()
        context.Next()
    }
}
