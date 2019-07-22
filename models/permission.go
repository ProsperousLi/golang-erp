package models

import (
	"strconv"
	"strings"

	"erpweb/logs"
	"erpweb/util"
)

type Permission struct {
	Id     int64  `json:"id" orm:"column(id)"`
	UserID int64  `json:"userID" orm:"column(userID)"` //employee的id字段
	Read   string `json:"read" orm:"column(read)"`     //只读菜单列表(为逗号分隔的数字)
	Write  string `json:"write" orm:"column(write)"`   //可写菜单列表(为逗号分隔的数字)
}

type WebPermission struct {
	Id     int64   `json:"id" orm:"column(id)"`
	UserID int64   `json:"userID" orm:"column(userID)"` //employee的id字段
	Read   []int64 `json:"read" orm:"column(read)"`     //只读菜单列表(为逗号分隔的数字)
	Write  []int64 `json:"write" orm:"column(write)"`   //可写菜单列表(为逗号分隔的数字)
}

func convert2Web(param Permission) (ret WebPermission) {
	ret.Id = param.Id
	ret.UserID = param.UserID

	reads := strings.Split(param.Read, ",")
	for _, v := range reads {
		read64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ret.Read = append(ret.Read, read64)
		}
	}

	writes := strings.Split(param.Write, ",")
	for _, v := range writes {
		write64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ret.Write = append(ret.Write, write64)
		}
	}

	return ret
}

func GetPermissionBypage(pageNum, pageSize int64) []WebPermission {
	var (
		pers []Permission
		rets []WebPermission
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.PERMISSION_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&pers)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}

	for _, v := range pers {
		rets = append(rets, convert2Web(v))
	}

	return rets
}

func GetPermissionByID(id int64) (ret WebPermission, err error) {
	var per Permission
	per.Id = id
	err = OSQL.Read(&per, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ret, err
	}

	ret = convert2Web(per)

	return ret, nil
}

func EditPermissionById(per Permission) (errorCode int64) {
	var (
		temp Permission
	)
	errorCode = util.SUCESSFUL
	temp.Id = per.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.PERMISSION_EDIT_FAILED
		return errorCode
	}

	args := edit_permission(per)
	num, err2 := OSQL.Update(&per, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.PERMISSION_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)

	return errorCode
}

func edit_permission(per Permission) []string {
	var (
		args []string
	)

	if per.Read != "" {
		args = append(args, "read")
	}
	if per.UserID != 0 {
		args = append(args, "userID")
	}
	if per.Write != "" {
		args = append(args, "write")
	}

	return args
}

func AddPermission(per Permission) (errorCode int64) {
	var (
		temp Permission
	)
	errorCode = util.SUCESSFUL
	temp.Id = per.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("permission have this id=%v", per.UserID)
		errorCode = util.PERMISSION_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&per)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.PERMISSION_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)

	return errorCode
}

func DeletePermission(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Permission
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.PERMISSION_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
