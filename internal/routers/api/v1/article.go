package v1

import (
    "github.com/gin-gonic/gin"

    "github.com/hd2yao/blog/pkg/app"
    "github.com/hd2yao/blog/pkg/err_code"
)

func NewArticle() Article {
    return Article{}
}

type Article struct{}

func (a Article) Get(c *gin.Context) {
    app.NewResponse(c).ToErrorResponse(err_code.ServerError)
    return
}

// @Summary 获取多个文章
// @Produce  json
// @Param title query string false "文章标题" maxlength(100)
// @Param desc query string false "简述" maxlength(255)
// @Param content query string false "文章内容" maxlength(65535)
// @Param cover_image_url query string false "封面图片地址" maxlength(255)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} err_code.Error "请求错误"
// @Failure 500 {object} err_code.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary 新增文章
// @Produce  json
// @Param title body string false "文章标题" maxlength(100)
// @Param desc body string false "简述" maxlength(255)
// @Param content body string false "文章内容" maxlength(65535)
// @Param cover_image_url body string false "封面图片地址" maxlength(255)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} err_code.Error "请求错误"
// @Failure 500 {object} err_code.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Param title body string false "文章标题" maxlength(100)
// @Param desc body string false "简述" maxlength(255)
// @Param content body string false "文章内容" maxlength(65535)
// @Param cover_image_url body string false "封面图片地址" maxlength(255)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} err_code.Error "请求错误"
// @Failure 500 {object} err_code.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} err_code.Error "请求错误"
// @Failure 500 {object} err_code.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
