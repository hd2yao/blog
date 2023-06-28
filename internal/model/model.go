package model

import (
    "fmt"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/hd2yao/blog/global"
    "github.com/hd2yao/blog/pkg/setting"
)

type Model struct {
    ID         uint32 `gorm:"primary_key" json:"id"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    CreatedOn  uint32 `json:"created_on"`
    ModifiedOn uint32 `json:"modified_on"`
    DeletedOn  uint32 `json:"deleted_on"`
    IsDel      uint32 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
    s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
    db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
        databaseSetting.UserName,
        databaseSetting.Password,
        databaseSetting.Host,
        databaseSetting.DBName,
        databaseSetting.Charset,
        databaseSetting.ParseTime,
    ))
    if err != nil {
        return nil, err
    }

    if global.ServerSetting.RunMode == "debug" {
        db.LogMode(true)
    }
    db.SingularTable(true)
    // 注册回调行为
    db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
    db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
    db.Callback().Delete().Replace("gorm:delete", deleteCallback)

    db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
    db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

    return db, nil
}

/*
   model 的回调方法
*/

// 新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
    if !scope.HasError() {
        nowTime := time.Now().Unix()
        // 通过调用 scope.FieldByName 方法，获取当前是否包含所需的字段
        if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
            if createTimeField.IsBlank {
                _ = createTimeField.Set(nowTime)
            }
        }

        if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
            if modifyTimeField.IsBlank {
                _ = modifyTimeField.Set(nowTime)
            }
        }
    }
}

// 更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
    // 获取当前设置了标识 gorm:update_column 的字段属性
    if _, ok := scope.Get("gorm:update_column"); !ok {
        // 若不存在，也就是没有自定义设置 update_column，
        // 那么将会在更新回调内设置默认字段 ModifiedOn 的值为当前的时间戳
        _ = scope.SetColumn("ModifiedOn", time.Now().Unix())
    }
}

// 删除行为的回调
func deleteCallback(scope *gorm.Scope) {
    if !scope.HasError() {
        var extraOption string
        if str, ok := scope.Get("gorm:delete_option"); ok {
            extraOption = fmt.Sprint(str)
        }

        // 判断是否存在 DeletedOn 和 IsDel 字段
        deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
        isDelField, hasIdDelField := scope.FieldByName("IsDel")
        // 若存在则调整为执行 UPDATE 操作进行软删除（修改 DeletedOn 和 IsDel 的值）
        if !scope.Search.Unscoped && hasDeletedOnField && hasIdDelField {
            now := time.Now().Unix()
            scope.Raw(fmt.Sprintf(
                "UPDATE %v SET %v=%v,%v=%v%v%v",
                scope.QuotedTableName(),
                scope.Quote(deletedOnField.DBName),
                scope.AddToVars(now),
                scope.Quote(isDelField.DBName),
                scope.AddToVars(1),
                // 调用 scope.CombinedConditionSql 方法完成 SQL 语句的组装
                addExtraSpaceIfExist(scope.CombinedConditionSql()),
                addExtraSpaceIfExist(extraOption),
            )).Exec()
        } else { // 否则执行 DELETE 进行硬删除
            scope.Raw(fmt.Sprintf(
                "DELETE FROM %v%v%v",
                scope.QuotedTableName(),
                addExtraSpaceIfExist(scope.CombinedConditionSql()),
                addExtraSpaceIfExist(extraOption),
            )).Exec()
        }
    }
}

func addExtraSpaceIfExist(str string) string {
    if str != "" {
        return " " + str
    }
    return ""
}
