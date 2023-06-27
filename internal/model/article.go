package model

import "github.com/hd2yao/blog/pkg/app"

type Article struct {
    *Model
    Title         string `json:"title"`
    Desc          string `json:"desc"`
    Content       string `json:"content"`
    CoverImageUrl string `json:"cover_image_url"`
    State         uint8  `json:"state"`
}

func (a Article) TableName() string {
    return "blog_article"
}

type ArticleSwagger struct {
    List  []*Article
    Pager *app.Pager
}

type CountArticleRequest struct {
    Title string `json:"title" form:"title" binding:"max=100"`
    State uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
    Title string `json:"title" form:"title" binding:"max=100"`
    State uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
    Title         string `json:"title" form:"title" binding:"max=100"`
    Desc          string `json:"desc" form:"desc" binding:"required"`
    Content       string `json:"content" form:"content" binding:"required,min=0"`
    CoverImageUrl string `json:"cover_image_url" form:"cover_image_url" binding:"required,min=0"`
    CreatedBy     string `json:"created_by" form:"created_by" binding:"required,min=3,max=100"`
    State         uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
    ID         uint32 `json:"id" form:"id" binding:"required,gte=1"`
    Title      string `json:"title" form:"title" binding:"max=100"`
    State      uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
    ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
    ID uint32 `json:"id" form:"id" binding:"required,gte=1"`
}
