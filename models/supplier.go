package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//供应商信息表
type Supplier struct {
	Id          int64  `json:"id" orm:"column(id)"`
	Suppcode    string `json:"suppcode" orm:"column(suppcode)"`       //供应商编号
	Name        string `json:"name" orm:"column(name)"`               //供应商名称
	Address     string `json:"address" orm:"column(address)"`         //地址
	Zipcode     string `json:"zipcode" orm:"column(zipcode)"`         //邮编
	Fax         string `json:"fax" orm:"column(fax)"`                 //传真
	Website     string `json:"website" orm:"column(website)"`         //官网
	Depositbank string `json:"depositbank" orm:"column(depositbank)"` //开户行
	Bankaccount string `json:"bankaccount" orm:"column(bankaccount)"` //银行账号
	Taxrate     string `json:"taxrate" orm:"column(taxrate)"`         //税率
	Paymethod   string `json:"paymethod" orm:"column(paymethod)"`     //付款方式
	Taxnum      string `json:"taxnum" orm:"column(taxnum)"`           //税号
}

func QuerySupplier(querystr string) []Supplier {
	var (
		custs1, custs2 []Supplier
	)

	if querystr != "" {
		num, err := OSQL.Raw("select * from "+
			util.SUPPLIER_TABLE_NAME+
			" where name like '%?%' order by id asc", querystr).QueryRows(&custs1)
		if err != nil {
			logs.FileLogs.Error("%s", err)
			return custs1
		}
		logs.FileLogs.Info("num1=%v", num)

		num, err = OSQL.Raw("select * from "+
			util.SUPPLIER_TABLE_NAME+
			" where suppcode like '%?%' order by id asc", querystr).QueryRows(&custs2)
		if err != nil {
			logs.FileLogs.Error("%s", err)
			return custs1
		}
		logs.FileLogs.Info("num2=%v", num)
	} else {
		num, err := OSQL.Raw("select * from " +
			util.SUPPLIER_TABLE_NAME +
			" order by id asc",
		).QueryRows(&custs1)

		if err != nil {
			logs.FileLogs.Error("%s", err)
		}
		logs.FileLogs.Info("num=%v", num)
	}

	if len(custs2) > 0 {
		custs1 = append(custs1, custs2...)
	}

	return custs1
}

func GetSupplierBypage(pageNum, pageSize int64) []Supplier {
	var (
		pas []Supplier
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.SUPPLIER_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas
}

func GetSupplierById(id int64) (pa Supplier, err error) {
	pa.Id = id
	err = OSQL.Read(&pa, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return pa, err
	}
	return pa, nil
}

func EditSupplierById(pa Supplier) (errorCode int64) {
	var (
		temp Supplier
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.SUPPLIER_EDIT_FAILED
		return errorCode
	}

	args := editArgs_supp(pa)

	num, err2 := OSQL.Update(&pa, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.SUPPLIER_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func editArgs_supp(pa Supplier) []string {
	var (
		args []string
	)

	if pa.Address != "" {
		args = append(args, "address")
	}

	if pa.Bankaccount != "" {
		args = append(args, "bankaccount")
	}

	if pa.Depositbank != "" {
		args = append(args, "depositbank")
	}

	if pa.Fax != "" {
		args = append(args, "fax")
	}

	if pa.Name != "" {
		args = append(args, "name")
	}

	if pa.Paymethod != "" {
		args = append(args, "paymethod")
	}

	if pa.Suppcode != "" {
		args = append(args, "suppcode")
	}

	if pa.Taxnum != "" {
		args = append(args, "taxnum")
	}

	if pa.Taxrate != "" {
		args = append(args, "taxrate")
	}

	if pa.Website != "" {
		args = append(args, "website")
	}

	if pa.Zipcode != "" {
		args = append(args, "zipcode")
	}

	return args
}

func AddSupplier(pa Supplier) (errorCode int64) {
	var (
		temp Supplier
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", pa.Id)
		errorCode = util.SUPPLIER_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.SUPPLIER_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteSupplier(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Supplier
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.SUPPLIER_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
