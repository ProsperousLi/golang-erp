package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `purchasecontract`;
// CREATE TABLE `purchasecontract` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `contractcode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '采购合同编号',
//   `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购员',
//   `currentreviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '当前审核人',
//   `currentreviewindex` tinyint(5) DEFAULT NULL COMMENT '当前审核人序号',
//   `status` tinyint(5) NOT NULL COMMENT '状态(1:制单;2:审核；3:执行中；4:执行完；5:已结算)',
//   `suppcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '供应商编号',
//   `type` tinyint(5) NOT NULL COMMENT '源类型(1：配件销售合同；2：维修合同；3：消耗品)',
//   `relatedcode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联合同号(消耗品时可以为空)',
//   `receiveaddress` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '收货地址',
//   `receiver` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '收货人',
//   `receiverphone` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '收货人电话',
//   `taxsign` tinyint(5) NOT NULL COMMENT '含税标志(1：含税价；2：不含税价)',
//   `taxrate` tinyint(5) NOT NULL COMMENT '税率',
//   `ralatedinquirycode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联询价单号',
//   `settlestatus` tinyint(5) NOT NULL COMMENT '结算状态(1:未结算；2:已结算；3:部分结算)',
//   `settleamount` bigint(20) DEFAULT '0' COMMENT '已结算金额',
//   `amount` bigint(20) NOT NULL COMMENT '金额',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='采购合同表';

type Purchasecontract struct {
	Id                 int64  `json:"id" orm:"column(id)"`
	Contractcode       string `json:"contractcode" orm:"column(contractcode)"`             //采购合同编号
	Handler            string `json:"handler" orm:"column(handler)"`                       //采购员
	Currentreviewer    string `json:"type" orm:"column(type)"`                             //当前审核人
	Currentreviewindex int    `json:"price" orm:"column(price)"`                           //当前审核人序号
	Status             int    `json:"outdate" orm:"column(outdate)"`                       //状态(1:制单;2:审核；3:执行中；4:执行完；5:已结算)',
	Suppcode           int    `json:"suppcode" orm:"column(suppcode)"`                     //供应商编号
	Type               string `json:"type" orm:"column(type)"`                             //源类型(1：配件销售合同；2：维修合同；3：消耗品)
	Relatedcode        string `json:"relatedcode" orm:"column(relatedcode)"`               //关联合同号(消耗品时可以为空)
	Receiveaddress     string `json:"receiveaddress" orm:"column(receiveaddress)"`         //收货地址
	Receiver           string `json:"receiver" orm:"column(receiver)"`                     //收货人
	Receiverphone      string `json:"receiverphone" orm:"column(receiverphone)"`           //收货人电话
	Taxsign            int    `json:"taxsign" orm:"column(taxsign)"`                       //含税标志(1：含税价；2：不含税价)
	Taxrate            int    `json:"taxrate" orm:"column(taxrate)"`                       //税率
	Ralatedinquirycode string `json:"ralatedinquirycode" orm:"column(ralatedinquirycode)"` //关联询价单号
	Settlestatus       int    `json:"settlestatus" orm:"column(settlestatus)"`             //结算状态(1:未结算；2:已结算；3:部分结算)
	Settleamount       int64  `json:"settleamount" orm:"column(settleamount)"`             //已结算金额
	Amount             int64  `json:"amount" orm:"column(amount)"`                         //金额

}

func GetPurchasecontractBypage(pageNum, pageSize int64) []Purchasecontract {
	var (
		purchasecontracts []Purchasecontract
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Purchasecontract_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&purchasecontracts)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return purchasecontracts
}

func GetPurchasecontractById(id int64) (purchasecontract Purchasecontract, err error) {
	purchasecontract.Id = id
	err = OSQL.Read(&purchasecontract, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return purchasecontract, err
	}
	return purchasecontract, nil
}

func EditPurchasecontractById(purchasecontract Purchasecontract) (errorCode int64) {
	var (
		temp Purchasecontract
	)
	errorCode = util.SUCESSFUL
	temp.Id = purchasecontract.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Purchasecontract_EDIT_FAILED
		return errorCode
	}

	args := edit_purchasecontract(purchasecontract)

	num, err2 := OSQL.Update(&purchasecontract, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Purchasecontract_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_purchasecontract(param Purchasecontract) (args []string) {
	if param.Amount != 0 {
		args = append(args, "amount")
	}

	if param.Contractcode != "" {
		args = append(args, "contractcode")
	}

	if param.Currentreviewer != "" {
		args = append(args, "currentreviewer")
	}

	if param.Currentreviewindex != 0 {
		args = append(args, "currentreviewindex")
	}

	if param.Handler != "" {
		args = append(args, "handler")
	}

	if param.Ralatedinquirycode != "" {
		args = append(args, "ralatedinquirycode")
	}

	if param.Receiveaddress != "" {
		args = append(args, "receiveaddress")
	}

	if param.Receiver != "" {
		args = append(args, "receiver")
	}

	if param.Receiverphone != "" {
		args = append(args, "receiverphone")
	}

	if param.Relatedcode != "" {
		args = append(args, "relatedcode")
	}

	if param.Settleamount != 0 {
		args = append(args, "settleamount")
	}

	if param.Settlestatus != 0 {
		args = append(args, "settlestatus")
	}

	if param.Status != 0 {
		args = append(args, "status")
	}

	if param.Suppcode != 0 {
		args = append(args, "suppcode")
	}

	if param.Taxrate != 0 {
		args = append(args, "taxrate")
	}

	if param.Taxsign != 0 {
		args = append(args, "taxsign")
	}

	if param.Type != "" {
		args = append(args, "type")
	}

	return args
}

func AddPurchasecontract(purchasecontract Purchasecontract) (errorCode int64) {
	var (
		temp Purchasecontract
	)
	errorCode = util.SUCESSFUL
	temp.Id = purchasecontract.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("purchasecontract have this id=%v", purchasecontract.Id)
		errorCode = util.Purchasecontract_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&purchasecontract)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Purchasecontract_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeletePurchasecontract(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Purchasecontract
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Purchasecontract_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
