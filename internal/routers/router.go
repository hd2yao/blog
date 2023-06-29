package routers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/hd2yao/blog/global"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"

    _ "github.com/hd2yao/blog/docs"
    "github.com/hd2yao/blog/internal/middleware"
    "github.com/hd2yao/blog/internal/routers/api"
    "github.com/hd2yao/blog/internal/routers/api/v1"
)

func NewPouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(middleware.Translations())

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    upload := api.NewUpload()
    r.POST("/upload/file", upload.UploadFile)
    r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

    article := v1.NewArticle()
    tag := v1.NewTag()

    apiv1 := r.Group("/api/v1")
    {
        apiv1.POST("/tags", tag.Create)
        apiv1.DELETE("/tags/:id", tag.Delete)
        apiv1.PUT("/tags/:id", tag.Update)
        apiv1.PATCH("/tags/:id/state", tag.Update)
        apiv1.GET("/tags", tag.List)

        apiv1.POST("/articles", article.Create)
        apiv1.DELETE("/articles/:id", article.Delete)
        apiv1.PUT("/articles/:id", article.Update)
        apiv1.PATCH("/articles/:id/state", article.Update)
        apiv1.GET("/articles/:id", article.Get)
        apiv1.GET("/articles", article.List)
    }

    return r
}
