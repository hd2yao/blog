package model

import (
	"github.com/jinzhu/gorm"

	"github.com/hd2yao/blog/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State *uint8 `json:"state"`
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

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(t).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
