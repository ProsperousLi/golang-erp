package models

import (
	"strconv"
	"strings"

	"erpweb/logs"
	"erpweb/util"
)

type Permission struct {
	Id     int64  `json:"id" orm:"column(id)"`
	Cardid string `json:"cardid" orm:"column(cardid)"` //工号
	Read   string `json:"read" orm:"column(read)"`     //只读菜单列表(为逗号分隔的数字)
	Write  string `json:"write" orm:"column(write)"`   //可写菜单列表(为逗号分隔的数字)
}

type WebPermission struct {
	Id     int64   `json:"id" orm:"column(id)"`
	Cardid string  `json:"cardid" orm:"column(cardid)"` //工号
	Read   []int64 `json:"read" orm:"column(read)"`     //只读菜单列表(为逗号分隔的数字)
	Write  []int64 `json:"write" orm:"column(write)"`   //可写菜单列表(为逗号分隔的数字)
}

func convert2Web(param Permission) (ret WebPermission) {
	ret.Id = param.Id
	ret.Cardid = param.Cardid

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
	if per.Cardid != "" {
		args = append(args, "cardid")
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
	temp.Cardid = per.Cardid
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("permission have this id=%v", per.Cardid)
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

//{cardid: 1, read:[1,2,3], write:[4,5,6]}
type AddPersionStruct struct {
	Cardid int64
	Read   []int64
	Write  []int64
}

// 参数：{cardid: 1, read:[1,2,3], write:[4,5,6]}
// 返回：{code:20000}或者{code:20001, message:”没有权限设置用户账户权限”}
func SetPermission(param AddPersionStruct) (errorCode int64, msg string) {
	var (
		temp, updatePer Permission
	)
	errorCode = util.SUCESSFUL
	temp.Cardid = strconv.FormatInt(param.Cardid, 10)
	updatePer.Cardid = temp.Cardid

	for i, re := range param.Read {
		if i == 0 {
			updatePer.Read += strconv.FormatInt(re, 10)
		} else {
			updatePer.Read += "," + strconv.FormatInt(re, 10)
		}
	}

	for i, wr := range param.Write {
		if i == 0 {
			updatePer.Write += strconv.FormatInt(wr, 10)
		} else {
			updatePer.Write += "," + strconv.FormatInt(wr, 10)
		}
	}
	err := OSQL.Read(&temp, "cardid")
	if err == nil { //find update
		args := edit_permission(updatePer)
		num, err2 := OSQL.Update(&updatePer, args...)
		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.FAILED
			msg = "更新失败"
			return
		}
		logs.FileLogs.Info("num=%v", num)
	} else { //add

		num, err2 := OSQL.Insert(&updatePer)
		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.FAILED
			msg = "添加失败"
			return
		}
		logs.FileLogs.Info("num=%v", num)
	}

	return
}

func QueryPermission(cardid string) (ret WebPermission, err error) {
	var per Permission
	per.Cardid = cardid
	err = OSQL.Read(&per, "cardid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ret, err
	}

	ret = convert2Web(per)

	return ret, nil
}
