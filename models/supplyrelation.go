package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//供货关系表
type Supplyrelation struct {
	Id         int64 `json:"id" orm:"column(id)"`
	Supplierid int64 `json:"supplierid" orm:"column(supplierid)"` //供应商主键
	Matterid   int64 `json:"matterid" orm:"column(matterid)"`     //物料主键
}

func UpdateSupplierListOfMatter(matterid int64, supplierList []int64) (errorCode int64, msg string) {
	errorCode = util.SUCESSFUL
	//delete
	code := DeleteSupplyrelationByMatterId(matterid)
	if code != util.SUCESSFUL {
		msg = "删除失败"
		return code, msg
	}

	//add
	for _, supplierid := range supplierList {
		var tempRelation Supplyrelation
		tempRelation.Matterid = matterid
		tempRelation.Supplierid = supplierid
		AddSupplyrelation(tempRelation)
	}

	return
}

func UpdateRelationBySupplierid(supplierid int64, matterList []int64) (errorCode int64, msg string) {
	errorCode = util.SUCESSFUL
	//delete
	code := DeleteSupplyrelationBySupplierid(supplierid)
	if code != util.SUCESSFUL {
		msg = "删除失败"
		return code, msg
	}

	//add
	for _, matterId := range matterList {
		var tempRelation Supplyrelation
		tempRelation.Matterid = matterId
		tempRelation.Supplierid = supplierid
		AddSupplyrelation(tempRelation)
	}

	return
}

func GetSupplyrelationByMatterid(matterid int64) (pa []Supplyrelation, err error) {
	var (
		pas []Supplyrelation
	)
	_, err = OSQL.Raw("select * from " + util.SUPPLYRELATION_TABLE_NAME +
		" where matterid=? order by id asc").QueryRows(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas, err
}

func GetSupplyrelationBySupplierid(Supplierid int64) (pa []Supplyrelation, err error) {
	var (
		pas []Supplyrelation
	)
	_, err = OSQL.Raw("select * from " + util.SUPPLYRELATION_TABLE_NAME +
		" where supplierid=? order by id asc").QueryRows(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas, err
}

func GetSupplyrelationBypage(pageNum, pageSize int64) []Supplyrelation {
	var (
		pas []Supplyrelation
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.SUPPLYRELATION_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas
}

func GetSupplyrelationById(id int64) (pa Supplyrelation, err error) {
	pa.Id = id
	err = OSQL.Read(&pa, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return pa, err
	}
	return pa, nil
}

func EditSupplyrelationById(pa Supplyrelation) (errorCode int64) {
	var (
		temp Supplyrelation
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.SUPPLYRELATION_EDIT_FAILED
		return errorCode
	}

	args := edit_supplyrelation(pa)
	num, err2 := OSQL.Update(&pa, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.SUPPLYRELATION_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_supplyrelation(param Supplyrelation) (args []string) {
	if param.Matterid != 0 {
		args = append(args, "matterid")
	}
	if param.Supplierid != 0 {
		args = append(args, "supplierid")
	}
	return args
}

func AddSupplyrelation(pa Supplyrelation) (errorCode int64) {
	var (
		temp Supplyrelation
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", pa.Id)
		errorCode = util.SUPPLYRELATION_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.SUPPLYRELATION_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteSupplyrelation(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Supplyrelation
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.SUPPLYRELATION_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteSupplyrelationBySupplierid(supplierid int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Supplyrelation
	)
	temp.Supplierid = supplierid
	num, err := OSQL.Delete(&temp, "supplierid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteSupplyrelationByMatterId(matterid int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Supplyrelation
	)
	temp.Matterid = matterid
	num, err := OSQL.Delete(&temp, "matterid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
