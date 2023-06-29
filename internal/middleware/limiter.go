package middleware

import (
    "github.com/gin-gonic/gin"

    "github.com/hd2yao/blog/pkg/app"
    "github.com/hd2yao/blog/pkg/err_code"
    "github.com/hd2yao/blog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
    return func(context *gin.Context) {
        key := l.Key(context)
        if bucket, ok := l.GetBucket(key); ok {
            count := bucket.TakeAvailable(1)
            if count == 0 {
                response := app.NewResponse(context)
                response.ToErrorResponse(err_code.TooManyRequests)
                context.Abort()
                return
            }
        }
        context.Next()
    }
}
