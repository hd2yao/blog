package v1

import (
    "github.com/gin-gonic/gin"

    "blog/pkg/app"
    "blog/pkg/err_code"
)

func NewArticle() Article {
    return Article{}
}

type Article struct{}

func (a Article) Get(c *gin.Context) {
    app.NewResponse(c).ToErrorResponse(err_code.ServerError)
    return
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
