package models

import (
	"erpweb/logs"
	"erpweb/util"
)

type Menu struct {
	Id        int64  `json:"id" orm:"column(id)"`               //菜单id
	Title     string `json:"title" orm:"column(title)"`         //标题
	Icon      string `json:"icon" orm:"column(icon)"`           //图标
	ParentID  int64  `json:"parentID" orm:"column(parentID)"`   //父ID
	Path      string `json:"path" orm:"column(path)"`           //路径
	Component string `json:"component" orm:"column(component)"` //组件名
}

func GetMenus() []Menu {
	var (
		menus []Menu
	)
	_, err := OSQL.Raw("select * from " + util.EMPLOYEE_TABLE_NAME).QueryRows(&menus)
	if err != nil {
		logs.FileLogs.Error("%s", err) //logs.Error(err)
	}
	return menus
}

// func GetMenuByUserID(userID int64) (menu Menu, err error) {
// 	menu.UserID = userID

// 	return menu, nil
// }

// func EditMenuById(id int64) (errorCode int64) {
// 	return errorCode
// }

// func AddMenument(menu Menu) (errorCode int64) {
// 	return errorCode
// }

// func DeleteMenu(id int64) (errorCode int64) {
// 	return errorCode
// }
