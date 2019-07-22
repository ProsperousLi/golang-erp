package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//客户信息表
type Customer struct {
	Id              int64  `json:"id" orm:"column(id)"`
	Custcode        string `json:"custcode" orm:"column(custcode)"`               //客户编码
	Name            string `json:"name" orm:"column(name)"`                       //客户名称
	Zipcode         string `json:"zipcode" orm:"column(zipcode)"`                 //邮编
	Postaddress     string `json:"postaddress" orm:"column(postaddress)"`         //通讯地址
	Taxnum          string `json:"taxnum" orm:"column(taxnum)"`                   //税号
	Depositbank     string `json:"depositbank" orm:"column(depositbank)"`         //开户银行
	Bankaccount     string `json:"bankaccount" orm:"column(bankaccount)"`         //银行账号
	Railwayadmin    string `json:"railwayadmin" orm:"column(railwayadmin)"`       //铁路局
	Maintainsection string `json:"maintainsection" orm:"column(maintainsection)"` //机务段
	Remark          string `json:"remark" orm:"column(remark)"`                   //备注
}

func GetCustomerBypage(pageNum, pageSize int64) []Customer {
	var (
		custs []Customer
	)
	begin := pageSize * pageNum
	logs.FileLogs.Info("begin=%v", begin, ", end =%v", pageSize)
	num, err := OSQL.Raw("select * from "+util.CUSTOMER_TABLE_NAME+" order by id asc limit ?,?",
		pageNum, pageSize).QueryRows(&custs)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	logs.FileLogs.Info("num=%v", num)
	return custs
}

func GetCustomerById(id int64) (cust Customer, err error) {
	cust.Id = id
	err = OSQL.Read(&cust, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return cust, err
	}
	return cust, nil
}

func EditCustomerById(cust Customer) (errorCode int64) {
	var (
		temp Customer
	)
	errorCode = util.SUCESSFUL
	temp.Id = cust.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.CUSTOMER_EDIT_FAILED
		return errorCode
	}

	args := editArgs_cu(cust)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&cust, args...)
		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.CUSTOMER_EDIT_FAILED
			return errorCode
		}
		logs.FileLogs.Info("num=%v", num)
	} else {
		logs.FileLogs.Info("no data update")
	}

	return errorCode
}

func editArgs_cu(temp Customer) []string {
	var (
		args []string
	)
	if temp.Custcode != "" {
		args = append(args, "custcode")
	}

	if temp.Name != "" {
		args = append(args, "name")
	}

	if temp.Zipcode != "" {
		args = append(args, "zipcode")
	}

	if temp.Postaddress != "" {
		args = append(args, "postaddress")
	}

	if temp.Taxnum != "" {
		args = append(args, "taxnum")
	}

	if temp.Depositbank != "" {
		args = append(args, "depositbank")
	}

	if temp.Bankaccount != "" {
		args = append(args, "bankaccount")
	}

	if temp.Railwayadmin != "" {
		args = append(args, "railwayadmin")
	}

	if temp.Maintainsection != "" {
		args = append(args, "maintainsection")
	}

	if temp.Remark != "" {
		args = append(args, "remark")
	}

	logs.FileLogs.Info("args=%v", args)
	return args
}

func AddCustomer(cust Customer) (errorCode int64) {
	var (
		temp Customer
	)
	errorCode = util.SUCESSFUL
	temp.Id = cust.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", cust.Id)
		errorCode = util.CUSTOMER_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&cust)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.CUSTOMER_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteCustomer(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Customer
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.CUSTOMER_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
