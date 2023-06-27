package model

import "github.com/hd2yao/blog/pkg/app"

type Tag struct {
    *Model
    Name  string `json:"name"`
    State uint8  `json:"state"`
}

func (t Tag) TableName() string {
    return "blog_tag"
}

type TagSwagger struct {
    List  []*Tag
    Pager *app.Pager
}

type CountTagRequest struct {
    Name  string `json:"name" form:"name" binding:"max=100"`
    State uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
    Name  string `json:"name" form:"name" binding:"max=100"`
    State uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
    Name      string `json:"name" form:"name" binding:"required,min=3,max=100"`
    CreatedBy string `json:"created_by" form:"created_by" binding:"required,min=3,max=100"`
    State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
    ID         uint32 `json:"id" form:"id" binding:"required,gte=1"`
    Name       string `json:"name" form:"name" binding:"min=3,max=100"`
    State      uint8  `json:"state" form:"state" binding:"required,oneof=0 1"`
    ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
    ID uint32 `json:"id" form:"id" binding:"required,gte=1"`
}
