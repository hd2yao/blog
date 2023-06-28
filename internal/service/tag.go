package service

import (
	"github.com/hd2yao/blog/internal/model"
	"github.com/hd2yao/blog/pkg/app"
)

type CountTagRequest struct {
	Name  string `json:"name" form:"name" binding:"max=100"`
	State uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `json:"name" form:"name" binding:"max=100"`
	State uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `json:"name" form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `json:"created_by" form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `json:"id" form:"id" binding:"required,gte=1"`
	Name       string `json:"name" form:"name" binding:"min=3,max=100"`
	State      *uint8 `json:"state" form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, *param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
