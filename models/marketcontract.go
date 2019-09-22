package models

import (
	"strconv"

	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `marketcontract`;
// CREATE TABLE `marketcontract` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `contractcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '合同编号',
//   `type` tinyint(5) NOT NULL COMMENT '类型 1:维修合同;2:销售合同',
//   `custcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '客户编号',
//   `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '制单人',
//   `execstatus` tinyint(5) DEFAULT NULL COMMENT '执行状态(1:制单；2:审核；3:执行中；6:结束)',
//   `signdate` date DEFAULT NULL COMMENT '签订日期',
//   `deadline` date DEFAULT NULL COMMENT '结束日期',
//   `createat` datetime DEFAULT NULL COMMENT '制单日期',
//   `amount` bigint(20) DEFAULT NULL COMMENT '金额',
//   `settlestatus` tinyint(5) DEFAULT NULL COMMENT '结算状态(1:未结算；2:已结算；3:部分结算)',
//   `settleamount` bigint(20) DEFAULT NULL COMMENT '结算金额',
//   `attachment` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '附件',
//   `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
//   `currentreviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '当前审核人',
//   `currentreviewindex` tinyint(5) DEFAULT '-1' COMMENT '当前审核序号',
//   `relatedcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联合同编号',
//   `vehicles` text CHARACTER SET utf8mb4 NOT NULL COMMENT '车辆列表(车辆编号的json数组)',
//   PRIMARY KEY (`id`),
//   UNIQUE KEY `contractcode` (`contractcode`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='市场合同表';

//市场合同表
type Marketcontract struct {
	Id                 int64  `json:"id" orm:"column(id)"`                                 //id
	Contractcode       string `json:"contractcode" orm:"column(contractcode)"`             //合同编号
	Type               int8   `json:"type" orm:"column(type)"`                             //类型 1:维修合同;2:销售合同
	Custcode           string `json:"custcode" orm:"column(custcode)"`                     //客户编号
	Handler            string `json:"handler" orm:"column(handler)"`                       //制单人
	Execstatus         int8   `json:"execstatus" orm:"column(execstatus)"`                 //执行状态(1:制单；2:审核；3:执行中；6:结束)
	Signdate           string `json:"signdate" orm:"column(signdate)"`                     //签订日期
	Deadline           string `json:"deadline" orm:"column(deadline)"`                     //结束日期
	Createat           string `json:"createat" orm:"column(createat)"`                     //制单日期
	Amount             int64  `json:"amount" orm:"column(amount)"`                         //金额
	Settleamount       int64  `json:"settleamount" orm:"column(settleamount)"`             //结算金额
	Attachment         string `json:"attachment" orm:"column(attachment)"`                 //附件
	Remark             string `json:"remark" orm:"column(remark)"`                         //备注
	Currentreviewer    string `json:"currentreviewer" orm:"column(currentreviewer)"`       //当前审核人
	Currentreviewindex int8   `json:"currentreviewindex" orm:"column(currentreviewindex)"` //当前审核序号
	Relatedcode        string `json:"relatedcode" orm:"column(relatedcode)"`               //关联合同编号
	Vehicles           string `json:"vehicles" orm:"column(vehicles)"`                     //车辆列表(车辆编号的json数组)
}

func GetMarketcontractByType(codeType int8) []Marketcontract {
	var (
		params []Marketcontract
	)
	_, err := OSQL.Raw("select * from "+util.Marketcontract_TABLE_NAME+
		"where type=? order by id asc", codeType).QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

func GetMarketcontractBypage(marketType, execstatus,
	contractcode, custcode, handler string, pageNum, pageSize int64) []Marketcontract {
	var (
		params []Marketcontract
	)

	sql := "select * from " + util.Marketcontract_TABLE_NAME + " where 1=1"
	if marketType != "" {
		_, err := strconv.ParseInt(marketType, 10, 64)
		if err != nil {
			sql += " and type=" + marketType
		}
	}

	if execstatus != "" {
		_, err := strconv.ParseInt(execstatus, 10, 64)
		if err != nil {
			sql += " and execstatus=" + execstatus
		}
	}

	if contractcode != "" {
		sql += " and contractcode like '%" + contractcode + "%'"
	}

	if custcode != "" {
		sql += " and custcode='" + custcode + "'"
	}

	if handler != "" {
		sql += " and handler='" + handler + "'"
	}

	begin := pageSize * pageNum
	_, err := OSQL.Raw(sql+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

func GetMarketcontractById(id int64) (result Marketcontract, err error) {
	result.Id = id
	err = OSQL.Read(&result, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return result, nil
}

func EditMarketcontractById(param Marketcontract) (errorCode int64) {
	var (
		temp Marketcontract
	)
	temp.Id = param.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Marketcontract_EDIT_FAILED
		return errorCode
	}
	args := editArgs_market(param)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&param, args...)
		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.Marketcontract_EDIT_FAILED
			return errorCode
		}

		logs.FileLogs.Info("num=%v err=%v", num, err2)
	} else {
		logs.FileLogs.Info("no data update")
	}

	return errorCode
}

func editArgs_market(param Marketcontract) []string {
	var (
		args []string
	)

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
	if param.Currentreviewer != "" {
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
	if param.Execstatus != 0 {
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
	if param.Signdate != "" {
		args = append(args, "signdate")
	}
	if param.Type != 0 {
		args = append(args, "type")
	}

	if param.Vehicles != "" { //TODO pandding json type
		args = append(args, "vehicles")
	}

	return args
}

func AddMarketcontract(param Marketcontract) (errorCode int64, msg string) {
	var (
		temp Marketcontract
	)
	temp.Id = param.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Info("Marketcontract have asixt")
		errorCode = util.FAILED
		msg = "查询到已经有该数据"
		return
	}

	id, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.FAILED
		msg = "数据更新失败"
		return
	}

	logs.FileLogs.Info("num=%v", id)

	return
}

func DeleteMarketcontract(id int64) (errorCode int64, msg string) {
	var (
		temp Marketcontract
	)
	errorCode = util.SUCESSFUL
	temp.Id = id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Info("Marketcontract have asixt")
		errorCode = util.FAILED
		msg = "未查询到该数据"
		return
	}

	if temp.Execstatus != 1 {
		logs.FileLogs.Error("不能删除状态不为1的市场合同,status = %s", temp.Execstatus)
		errorCode = util.FAILED
		msg = "不能删除状态不为1的市场合同"
		return
	}

	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.FAILED
		msg = "删除失败"
		return
	}

	logs.FileLogs.Info("num=%v", num)

	return
}
