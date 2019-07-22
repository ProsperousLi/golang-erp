package models

import (
	"errors"

	"erpweb/logs"
	"erpweb/util"
)

type Account struct {
	Id       int64  `json:"id" orm:"column(id)"`             //账号信息表
	Cardid   int64  `json:"cardid" orm:"column(cardid)"`     //cardid
	Status   int    `json:"status" orm:"column(status)"`     //账号状态 1：正常；2:冻结; 3:停用；4：锁定
	Password string `json:"password" orm:"column(password)"` //密码
}

var (
	TokenMap = make(map[int64]string)       //cardid token
	AccsMap  = make(map[string]interface{}) //token account
)

func Login(cardid int64, password string) (error, string) {
	var (
		qurey Account
		uuid  string
	)
	qurey.Cardid = cardid
	//qurey.Password = password
	err := OSQL.Read(&qurey, "cardid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return err, uuid
	}

	if qurey.Cardid == cardid && qurey.Password == password { //login sucess
		uuid = util.GetToken()
		delete(TokenMap, cardid)
	} else {
		return errors.New("cardid or password is invild"), uuid
	}

	return nil, uuid
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

func GetAccountByUserID(cardid int64) (account Account, err error) {
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

func DeleteAccount(cardid int64) (errorCode int64) {
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
