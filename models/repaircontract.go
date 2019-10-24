package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

// DROP TABLE IF EXISTS `repaircontract`;
// CREATE TABLE `repaircontract` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `contractcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '合同编号',
//   `custcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '客户编号',
//   `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '负责人',
//   `execstatus` tinyint(5) NOT NULL COMMENT '执行状态(1:制单；2:审核；3:执行中；4:完工；5:封账;6:结束)',
//   `signdate` date DEFAULT NULL COMMENT '合同签订日期',
//   `deadline` date DEFAULT NULL COMMENT '合同结束日期',
//   `createat` datetime NOT NULL COMMENT '金额',
//   `amount` bigint(20) NOT NULL COMMENT '金额',
//   `settlestatus` tinyint(5) NOT NULL COMMENT '结算状态(1:未结算；2:已结算；3:部分结算)',
//   `settleamount` bigint(20) DEFAULT '0' COMMENT '结算金额',
//   `vehicles` text CHARACTER SET utf8mb4 NOT NULL COMMENT '车辆列表(车辆编号的json数组)',
//   `attachment` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '附件',
//   `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
//   `currentreviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '当前审核人',
//   `currentreviewindex` tinyint(5) DEFAULT '-1' COMMENT '当前审核序号',
//   `relatedcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联合同编号',
//   PRIMARY KEY (`id`),
//   UNIQUE KEY `contractcode` (`contractcode`) USING BTREE
// ) ENGINE=Inn oDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='维修合同表';

//维修合同表
type Repaircontract struct {
	Id                 int64  `json:"id" orm:"column(id)"`
	Contractcode       string `json:"contractcode" orm:"column(contractcode)"`             //合同编号
	Custcode           string `json:"custcode" orm:"column(custcode)"`                     //客户编号
	Handler            string `json:"handler" orm:"column(handler)"`                       //负责人
	Execstatus         string `json:"execstatus" orm:"column(execstatus)"`                 //执行状态(1:制单；2:审核；3:执行中；4:完工；5:封账;6:结束)
	Signdate           string `json:"signdate" orm:"column(signdate)"`                     //合同签订日期
	Deadline           string `json:"deadline" orm:"column(deadline)"`                     //合同结束日期
	Createat           string `json:"createat" orm:"column(createat)"`                     //金额
	Amount             int64  `json:"amount" orm:"column(amount)"`                         //金额
	Settlestatus       int8   `json:"settlestatus" orm:"column(settlestatus)"`             //结算状态(1:未结算；2:已结算；3:部分结算)
	Settleamount       int64  `json:"settleamount" orm:"column(settleamount)"`             //结算金额
	Vehicles           string `json:"vehicles" orm:"column(vehicles)"`                     //车辆列表(车辆编号的json数组)
	Attachment         string `json:"attachment" orm:"column(attachment)"`                 //附件
	Remark             string `json:"remark" orm:"column(remark)"`                         //备注
	Currentreviewer    int64  `json:"currentreviewer" orm:"column(currentreviewer)"`       //当前审核人
	Currentreviewindex int8   `json:"currentreviewindex" orm:"column(currentreviewindex)"` //当前审核序号
	Relatedcode        string `json:"relatedcode" orm:"column(relatedcode)"`               //关联合同编号
}

func GetRepaircontractBypage(pageNum, pageSize int64) []Repaircontract {
	var (
		params []Repaircontract
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Repaircontract_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&params)
	if err != nil {
		beego.Error(err)
	}
	return params
}

func GetRepaircontractById(id int64) (ret Repaircontract, err error) {
	ret.Id = id
	err = OSQL.Read(&ret, "id")
	if err != nil {
		beego.Error(err)
		return ret, err
	}
	return ret, nil
}

func EditRepaircontractById(param Repaircontract) (errorCode int64) {
	var (
		temp Repaircontract
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.Repaircontract_EDIT_FAILED
		return errorCode
	}

	args := edit_repaircontract(param)
	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Repaircontract_EDIT_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func edit_repaircontract(param Repaircontract) (args []string) {
	if param.Amount != 0 {
		args = append(args, "amount")
	}

	if param.Attachment != "" {
		args = append(args, "attachment")
	}

	if param.Contractcode != "" {
		args = append(args, "contractcode")
	}

	if param.Createat != "" {
		args = append(args, "createat")
	}

	if param.Currentreviewer != 0 {
		args = append(args, "currentreviewer")
	}

	if param.Currentreviewindex != 0 {
		args = append(args, "currentreviewindex")
	}

	if param.Custcode != "" {
		args = append(args, "custcode")
	}

	if param.Deadline != "" {
		args = append(args, "deadline")
	}

	if param.Execstatus != "" {
		args = append(args, "execstatus")
	}

	if param.Handler != "" {
		args = append(args, "handler")
	}

	if param.Relatedcode != "" {
		args = append(args, "relatedcode")
	}

	if param.Remark != "" {
		args = append(args, "remark")
	}

	if param.Settleamount != 0 {
		args = append(args, "settleamount")
	}

	if param.Settlestatus != 0 {
		args = append(args, "settlestatus")
	}

	if param.Signdate != "" {
		args = append(args, "signdate")
	}

	if param.Vehicles != "" {
		args = append(args, "vehicles")
	}
	return args
}

func AddRepaircontract(param Repaircontract) (errorCode int64) {
	var (
		temp Repaircontract
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		beego.Error("Repaircontract have this id=", param.Id)
		errorCode = util.Repaircontract_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Repaircontract_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func DeleteRepaircontract(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Repaircontract
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.Repaircontract_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
