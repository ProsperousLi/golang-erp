package models

import (
	"erpweb/logs"
	"erpweb/util"
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
		logs.FileLogs.Error("%s", err)

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
	logs.FileLogs.Info("begin=%v", begin, ", end =%v", pageSize)
	_, err := OSQL.Raw("select * from "+util.ACCOUNT_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&accounts)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return accounts
}

func GetAccountByUserID(cardid string) (account Account, err error) {
	account.Cardid = cardid
	err = OSQL.Read(&account, "cardid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return account, nil
}

func EditAccountById(account Account) (errorCode int64) {
	var (
		temp Account
	)
	temp.Id = account.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.ACCOUNT_EDIT_FAILED
		return errorCode
	}

	args := editArgs(account)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&account, args...)

		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.ACCOUNT_EDIT_FAILED
			return errorCode
		}

		logs.FileLogs.Info("num=%v err=%v", num, err2)
	} else {
		logs.FileLogs.Info("no data update")
	}
	return errorCode
}

func EditAccountStatusById(cardid string, status int8) (errorCode int64) {
	var (
		temp Account
	)
	temp.Cardid = cardid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "cardid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.ACCOUNT_EDIT_FAILED
		return errorCode
	}

	temp.Status = int(status)

	args := editArgs(temp)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&temp, args...)

		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.ACCOUNT_EDIT_FAILED
			return errorCode
		}

		logs.FileLogs.Info("num=%v err=%v", num, err2)
	} else {
		logs.FileLogs.Info("no data update")
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
	logs.FileLogs.Info("args=%v", args)
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
		logs.FileLogs.Info("account have asixt")
		errorCode = util.ACCOUNT_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&account)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.ACCOUNT_ADD_FAILED
	}

	logs.FileLogs.Info("num=%v", id)

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
		logs.FileLogs.Error("%v", err)
		errorCode = util.ACCOUNT_DELETE_FAILED
	}

	logs.FileLogs.Info("num=%v", num)

	return errorCode
}
