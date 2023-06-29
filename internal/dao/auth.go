package dao

import "github.com/hd2yao/blog/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
    auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
    return auth.Get(d.engine)
}
