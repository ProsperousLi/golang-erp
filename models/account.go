package models

import (
	"errors"
	"time"

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

var (
	TokenMap = make(map[string]string)      //cardid token
	AccsMap  = make(map[string]interface{}) //token employee
	TimeMap  = make(map[string]int64)       //token TimeMap
)

//TODO Timer delete token
func New_Time_count() {
	//frist clear map
	TokenMap = make(map[string]string)
	AccsMap = make(map[string]interface{})
	TimeMap = make(map[string]int64)
}

// 登录接口
func Login(cardid string, password string) (error, string) {
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
		if qurey.Status == 1 {
			preToken := TokenMap[cardid]
			uuid = util.GetToken()
			TokenMap[cardid] = uuid
			delete(TimeMap, preToken)
			TimeMap[uuid] = time.Now().Unix()
			delete(AccsMap, preToken)

			//qurey employee info
			emp, code := GetEmployeeByCardid(cardid)
			if code != util.SUCESSFUL {
				logs.FileLogs.Error("GetEmployeeByCardid failed")
				return errors.New("GetEmployeeByCardid failed"), uuid
			}
			AccsMap[uuid] = emp
		} else {
			logs.FileLogs.Error("status is :%s", qurey.Status)
			return errors.New("status isn't 1"), uuid
		}
	} else {
		logs.FileLogs.Error("cardid or password is invild")
		return errors.New("cardid or password is invild"), uuid
	}

	return nil, uuid
}

// 单点登录
func SSOLogin(token string) (err error, code int64) {
	code = util.SUCESSFUL
	if _, ok := AccsMap[token]; ok {
		//pandding time
		if lastTime, ok := TimeMap[token]; ok {
			if lastTime+1*60*60 < time.Now().Unix() {
				//token过期了
				code = 50014
				logs.FileLogs.Error("token过期了")
				return errors.New("token过期了"), code
			}
		}
	} else {
		code = 50008
		logs.FileLogs.Error("this token not exist :%s", token)

		return errors.New("token not exist"), code
	}

	return nil, code
}

func Loginout(token string) (code int64) {
	code = util.SUCESSFUL
	if emp, ok := AccsMap[token]; ok {
		//delete all map
		TokenMap = make(map[string]string)     //cardid token
		AccsMap = make(map[string]interface{}) //token employee
		TimeMap = make(map[string]int64)       //token TimeMap

		tempEmp := emp.(Employee)
		delete(TokenMap, tempEmp.Cardid)
		delete(AccsMap, token)
		delete(TimeMap, token)
	} else {
		code = 50008
	}

	return code
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
