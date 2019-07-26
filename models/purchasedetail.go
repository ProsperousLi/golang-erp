package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `purchasedetail`;
// CREATE TABLE `purchasedetail` (
//   `contractcode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购合同编号',
//   `mattercode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
//   `num` bigint(20) DEFAULT NULL COMMENT '采购数量',
//   `price` bigint(20) DEFAULT NULL COMMENT '单价',
//   `value` bigint(20) DEFAULT NULL COMMENT '总价',
//   PRIMARY KEY (`contractcode`,`mattercode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='采购合同详情表';

type Purchasedetail struct {
	Contractcode string `json:"contractcode" orm:"column(contractcode)"` //采购合同编号
	Mattercode   string `json:"mattercode" orm:"column(mattercode)"`     //物料编码
	Num          int64  `json:"num" orm:"column(num)"`                   //采购数量
	Price        int64  `json:"price" orm:"column(price)"`               //单价
	Value        int64  `json:"value" orm:"column(value)"`               //总价
}

func GetPurchasedetailBypage(pageNum, pageSize int64) []Purchasedetail {
	var (
		purchasedetails []Purchasedetail
	)
	err := OSQL.Raw("select * from "+util.Purchasedetail_TABLE_NAME+" order by contractcode desc limit ?,?",
		pageNum, pageSize).QueryRow(&purchasedetails)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return purchasedetails
}

func GetPurchasedetailById(contractcode string) (purchasedetail Purchasedetail, err error) {
	purchasedetail.Contractcode = contractcode
	err = OSQL.Read(&purchasedetail, "contractcode")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return purchasedetail, err
	}
	return purchasedetail, nil
}

func EditPurchasedetailById(purchasedetail Purchasedetail) (errorCode int64) {
	var (
		temp Purchasedetail
	)
	errorCode = util.SUCESSFUL
	temp.Contractcode = purchasedetail.Contractcode
	err := OSQL.Read(&temp, "contractcode")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Purchasedetail_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&purchasedetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Purchasedetail_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddPurchasedetail(purchasedetail Purchasedetail) (errorCode int64) {
	var (
		temp Purchasedetail
	)
	errorCode = util.SUCESSFUL
	temp.Contractcode = purchasedetail.Contractcode
	err := OSQL.Read(&temp, "contractcode")
	if err == nil {
		logs.FileLogs.Error("purchasedetail have this id=%v", purchasedetail.Contractcode)
		errorCode = util.Purchasedetail_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&purchasedetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Purchasedetail_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeletePurchasedetail(contractcode string) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Purchasedetail
	)
	temp.Contractcode = contractcode
	num, err := OSQL.Delete(&temp, "contractcode")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Purchasedetail_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
