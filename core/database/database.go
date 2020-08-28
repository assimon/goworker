package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goworker/core/config"
	"log"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
	DeletedAt int `json:"deleted_at"`
}

func DatabaseInit()  {
	var err error
	DB, err = gorm.Open(config.DatabaseConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Name))
	if err != nil{
		log.Fatalf("Database connection failed, err: %v", err)
	}
	// 设置默认表前缀.
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DatabaseConfig.TablePrefix + defaultTableName
	}
	DB.SingularTable(true)
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	DB.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DB.DB().SetMaxIdleConns(config.DatabaseConfig.MaxIdleConns)
	DB.DB().SetMaxOpenConns(config.DatabaseConfig.MaxOpenConns)
}

func CloseDB()  {
	defer DB.Close()
}

/**
创建数据时回调修改时间
 */
func updateTimeStampForCreateCallback(scope *gorm.Scope)  {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 当字段为新增的时候.
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok{
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if updateTimeField, ok := scope.FieldByName("UpdatedAt"); ok{
			if updateTimeField.IsBlank {
				updateTimeField.Set(nowTime)
			}
		}
	}
}

/**
修改数据时回调修改时间
 */
func updateTimeStampForUpdateCallback(scope *gorm.Scope)  {
	if _, ok := scope.Get("gorm:update_column"); !ok{
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}

/**
删除数据时回调修改时间
 */
func deleteCallback(scope *gorm.Scope)  {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
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