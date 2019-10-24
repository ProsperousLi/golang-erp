package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

type Account struct {
	Id       int64  `json:"id" orm:"column(id)"`             //账号信息表
	Cardid   string `json:"cardid" orm:"column(cardid)"`     //cardid
	Status   int    `json:"status" orm:"column(status)"`     //账号状态 1：正常；2:冻结; 3:停用；4：锁定
	Password string `json:"password" orm:"column(password)"` //密码
}

const (
	RefreshTime = 1
)

func GetAccountsNotPwd() (RetAccounts []map[string]interface{}) {
	var (
		accounts []Account
	)

	_, err := OSQL.Raw("select * from " + util.ACCOUNT_TABLE_NAME +
		" order by id asc").QueryRows(&accounts)
	if err != nil {
		beego.Error(err)

		return RetAccounts
	}

	for _, acc := range accounts {
		tempMsp := util.StructToMap(acc)
		delete(tempMsp, "password")
		delete(tempMsp, "Password")
		RetAccounts = append(RetAccounts, tempMsp)
	}

	return RetAccounts
}

func GetAccountBypage(pageNum, pageSize int64) []Account {
	var (
		accounts []Account
	)

	begin := pageSize * pageNum
	beego.Info("begin=", begin, ", end =", pageSize)
	_, err := OSQL.Raw("select * from "+util.ACCOUNT_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&accounts)
	if err != nil {
		beego.Error(err)
	}
	return accounts
}

func GetAccountByUserID(cardid string) (account Account, err error) {
	account.Cardid = cardid
	err = OSQL.Read(&account, "cardid")
	if err != nil {
		beego.Error(err)
	}
	return account, nil
}

func EditAccountById(account Account) (errorCode int64, msg string) {
	var (
		temp Account
	)
	temp.Cardid = account.Cardid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "cardid")
	if err != nil {
		beego.Error(err)
		errorCode = 20001
		msg = "未找到该用户"
		return
	}

	account.Id = temp.Id

	args := editArgs(account)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&account, args...)

		if err2 != nil {
			beego.Error(err2)
			errorCode = 20001
			msg = "数据更新失败"
			return
		}

		beego.Info("num= err=", num, err2)
	} else {
		beego.Info("no data update")
		errorCode = 20001
		msg = "没有数据要更新"
	}
	return
}

func EditAccountStatusById(cardid string, status int8) (errorCode int64) {
	var (
		temp Account
	)
	temp.Cardid = cardid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "cardid")
	if err != nil {
		beego.Error(err)
		errorCode = util.ACCOUNT_EDIT_FAILED
		return errorCode
	}

	temp.Status = int(status)

	args := editArgs(temp)
	if len(args) > 0 {
		var updateTemp Account
		updateTemp.Cardid = temp.Cardid
		updateTemp.Status = temp.Status

		num, err2 := OSQL.Update(&updateTemp, args...)

		if err2 != nil {
			beego.Error(err2)
			errorCode = util.ACCOUNT_EDIT_FAILED
			return errorCode
		}

		beego.Info("num= err=", num, err2)
	} else {
		beego.Info("no data update")
	}
	return errorCode
}

func editArgs(temp Account) []string {
	var (
		args []string
	)
	if temp.Status != 0 {
		args = append(args, "status")
	}
	// if temp.UserID != 0 {
	// 	args = append(args, "userID")
	// }
	if temp.Password != "" {
		args = append(args, "password")
	}
	beego.Info("args=", args)
	return args
}

func AddAccountment(account Account) (errorCode int64) {
	var (
		temp Account
	)
	temp.Cardid = account.Cardid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "Cardid")
	if err == nil {
		beego.Info("account have asixt")
		errorCode = util.ACCOUNT_ADD_FAILED
		return errorCode
	}

	if account.Password == "" {
		account.Password = util.GETMd5(util.DEFUAL_PWD_PRE + util.DEFUAL_PWD)
	} else {
		account.Password = util.GETMd5(util.DEFUAL_PWD_PRE + account.Password)
	}

	id, err2 := OSQL.Insert(&account)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.ACCOUNT_ADD_FAILED
	}

	beego.Info("num=", id)

	return errorCode
}

func DeleteAccount(cardid string) (errorCode int64) {
	var (
		account Account
	)
	errorCode = util.SUCESSFUL
	account.Cardid = cardid
	num, err := OSQL.Delete(&account, "cardid")
	if err != nil {
		beego.Error(err)
		errorCode = util.ACCOUNT_DELETE_FAILED
	}

	beego.Info("num=", num)

	return errorCode
}

type ModifyPwdStruct struct {
	Cardid string
	Oldpwd string
	Newpwd string
}

func ModifyPwd(param ModifyPwdStruct, token string) (errorCode int64, errorMessage string) {

	var (
		temp Account
	)
	temp.Cardid = param.Cardid
	errorCode = 20001
	err := OSQL.Read(&temp, "cardid")
	if err != nil {
		beego.Error(err)
		errorMessage = "未找到该用户"
		return
	}

	// 	Cardid string
	// Oldpwd string
	// Newpwd string

	if param.Oldpwd == "" {
		errorMessage = "旧密码为空"
		return
	}

	if param.Newpwd == "" {
		errorMessage = "新密码为空"
		return
	}

	if /*len(param.Oldpwd) <= 8 ||*/ len(param.Newpwd) < 8 {
		errorMessage = "密码至少8位"
		return
	}

	if /*!util.PanddingPwd(param.Oldpwd) ||*/ !util.PanddingPwd(param.Newpwd) {
		errorMessage = "密码必须至少包含字母和数字"
		return
	}
	//util.GETMd5(util.DEFUAL_PWD_PRE + util.DEFUAL_PWD)

	if temp.Password != util.GETMd5(util.DEFUAL_PWD_PRE+param.Oldpwd) {
		errorMessage = "旧密码不正确"
		return
	}

	temp.Password = util.GETMd5(util.DEFUAL_PWD_PRE + param.Newpwd)

	args := editArgs(temp)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&temp, args...)

		if err2 != nil {
			beego.Error(err2)
			errorMessage = "数据库更新密码失败"
			//errorCode = util.ACCOUNT_EDIT_FAILED
			return
		}

		beego.Info("num= err=", num, err2)
	} else {
		beego.Info("no data update")
		errorMessage = "没有数据要更新"
		return
	}

	//TODO 删除此人token
	if cardid, ok := AccsMap[token]; ok {
		delete(LimitMap, cardid)
		delete(TimeMap, token)
		delete(AccsMap, token)
		delete(TokenMap, cardid)
	}

	errorCode = util.SUCESSFUL
	return

}
