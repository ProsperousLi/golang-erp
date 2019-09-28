package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `purchasedetail`;
// CREATE TABLE `purchasedetail` (
//   `contractcode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购合同编号',
//   `mattercode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
//   `type` tinyint(5) NOT NULL COMMENT '源类型,1：配件销售合同；2：维修合同；3：消耗品',
//   `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联合同号,消耗品时可以为空',
//   `num` bigint(20) DEFAULT NULL COMMENT '采购数量',
//   `price` bigint(20) DEFAULT NULL COMMENT '单价',
//   `value` bigint(20) DEFAULT NULL COMMENT '总价',
//   `pureprice` bigint(20) DEFAULT NULL COMMENT '不含税价',
//   `purevalue` bigint(20) DEFAULT NULL COMMENT '不含税额',
//   `deadline` datetime DEFAULT NULL COMMENT '要求到货截止日期',
//   PRIMARY KEY (`contractcode`,`mattercode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='采购合同详情表';

type Purchasedetail struct {
	Contractcode string `json:"contractcode" orm:"column(contractcode)"` //采购合同编号
	Mattercode   string `json:"mattercode" orm:"column(mattercode)"`     //物料编码
	Type         int8   `json:"type" orm:"column(type)"`                 //源类型,1：配件销售合同；2：维修合同；3：消耗品
	Relatedcode  string `json:"relatedcode" orm:"column(relatedcode)"`   //关联合同号,消耗品时可以为空
	Num          int64  `json:"num" orm:"column(num)"`                   //采购数量
	Price        int64  `json:"price" orm:"column(price)"`               //单价
	Value        int64  `json:"value" orm:"column(value)"`               //总价
	Pureprice    int64  `json:"pureprice" orm:"column(pureprice)"`       //不含税价
	Purevalue    int64  `json:"purevalue" orm:"column(purevalue)"`       //不含税额
	Deadline     string `json:"deadline" orm:"column(deadline)"`         //要求到货截止日期
}

type PurchasedetailWeb struct {
	Contractcode string `json:"contractcode" orm:"column(contractcode)"` //采购合同编号
	Mattercode   string `json:"mattercode" orm:"column(mattercode)"`     //物料编码
	Type         int8   `json:"type" orm:"column(type)"`                 //源类型,1：配件销售合同；2：维修合同；3：消耗品
	Relatedcode  string `json:"relatedcode" orm:"column(relatedcode)"`   //关联合同号,消耗品时可以为空
	Num          int64  `json:"num" orm:"column(num)"`                   //采购数量
	Price        int64  `json:"price" orm:"column(price)"`               //单价
	Value        int64  `json:"value" orm:"column(value)"`               //总价
	Pureprice    int64  `json:"pureprice" orm:"column(pureprice)"`       //不含税价
	Purevalue    int64  `json:"purevalue" orm:"column(purevalue)"`       //不含税额
	Deadline     string `json:"deadline" orm:"column(deadline)"`         //要求到货截止日期
	Unit         string `json:"unit"`
	Param        string `json:"param"`
	Name         string `json:"name"`
}

func GetPurchasedetailBypage(contractcode string) (rets []PurchasedetailWeb) {
	var (
		purchasedetails []Purchasedetail
	)
	_, err := OSQL.Raw("select * from " + util.Purchasedetail_TABLE_NAME +
		" and contractcode='" + contractcode + "' order by contractcode desc").QueryRows(&purchasedetails)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}

	for _, temp := range purchasedetails {
		var tempRet PurchasedetailWeb
		tempRet.Contractcode = temp.Contractcode
		tempRet.Mattercode = temp.Mattercode
		tempRet.Type = temp.Type
		tempRet.Num = temp.Num
		tempRet.Price = temp.Price
		tempRet.Value = temp.Value
		tempRet.Pureprice = temp.Pureprice
		tempRet.Purevalue = temp.Purevalue
		tempRet.Deadline = temp.Deadline

		mat, err := GetMatterByMattercode(temp.Mattercode)
		if err != nil {
			continue
		}

		tempRet.Unit = mat.Unit
		tempRet.Param = mat.Param
		tempRet.Name = mat.Name

		rets = append(rets, tempRet)
	}

	return rets
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
	if err != nil { //add
		// logs.FileLogs.Error("%s", err)
		// errorCode = util.Purchasedetail_EDIT_FAILED
		// return errorCode
		code := AddPurchasedetail(purchasedetail)
		if code != util.SUCESSFUL {
			return errorCode
		}
	}

	//delete
	code := DeletePurchasedetail(temp.Contractcode)
	if code != util.SUCESSFUL {
		return errorCode
	}

	args := edit_purchasedetail(purchasedetail)

	num, err2 := OSQL.Update(&purchasedetail, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Purchasedetail_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_purchasedetail(param Purchasedetail) (args []string) {
	if param.Mattercode != "" {
		args = append(args, "mattercode")
	}

	if param.Num != 0 {
		args = append(args, "num")
	}

	if param.Price != 0 {
		args = append(args, "price")
	}

	if param.Value != 0 {
		args = append(args, "value")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}

	if param.Relatedcode != "" {
		args = append(args, "relatedcode")
	}

	if param.Pureprice != 0 {
		args = append(args, "pureprice")
	}

	if param.Purevalue != 0 {
		args = append(args, "purevalue")
	}

	if param.Deadline != "" {
		args = append(args, "deadline")
	}
	return args
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
