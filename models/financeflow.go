package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `financeflow`;
// CREATE TABLE `financeflow` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `type` tinyint(5) NOT NULL COMMENT '类型(1：维修合同；2：销售合同；3：采购合同)',
//   `direction` tinyint(5) NOT NULL COMMENT '资金流向(1：流入；-1：流出)',
//   `account` bigint(20) NOT NULL COMMENT '金额',
//   `paymethod` tinyint(5) NOT NULL COMMENT '支付方式(1：现金；2：银行转账；3：支票)',
//   `billcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '发票编号',
//   `attachment` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '附件',
//   `remark` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
//   `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '操作人',
//   `operdate` datetime NOT NULL COMMENT '操作时间',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='财务流水表';

//财务流水表
type Financeflow struct {
	Id         int64  `json:"id" orm:"column(id)"`                 //id
	Type       int    `json:"type" orm:"column(type)"`             //类型(1：维修合同；2：销售合同；3：采购合同)
	Direction  int    `json:"direction" orm:"column(direction)"`   //资金流向(1：流入；-1：流出)
	Account    int64  `json:"account" orm:"column(account)"`       //金额
	Paymethod  int    `json:"paymethod" orm:"column(paymethod)"`   //支付方式(1：现金；2：银行转账；3：支票)
	Billcode   string `json:"billcode" orm:"column(billcode)"`     //发票编号
	Attachment string `json:"attachment" orm:"column(attachment)"` //附件
	Remark     string `json:"remark" orm:"column(remark)"`         //备注
	Handler    string `json:"handler" orm:"column(handler)"`       //操作人
	Operdate   string `json:"operdate" orm:"column(operdate)"`     //操作时间
}

func GetFinanceflowBypage(pageNum, pageSize int64) []Financeflow {
	var (
		financeflows []Financeflow
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Financeflow_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&financeflows)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return financeflows
}

func GetFinanceflowByUserID(id int64) (financeflow Financeflow, err error) {
	financeflow.Id = id
	err = OSQL.Read(&financeflow, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return financeflow, nil
}

func EditFinanceflowById(financeflow Financeflow) (errorCode int64) {
	var (
		temp Financeflow
	)
	temp.Id = financeflow.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Financeflow_EDIT_FAILED
		return errorCode
	}

	args := edit_Financeflow(financeflow)
	num, err2 := OSQL.Update(&financeflow, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Financeflow_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func edit_Financeflow(param Financeflow) (args []string) {
	if param.Account != 0 {
		args = append(args, "account")
	}

	if param.Attachment != "" {
		args = append(args, "attachment")
	}

	if param.Billcode != "" {
		args = append(args, "billcode")
	}

	if param.Direction != 0 {
		args = append(args, "direction")
	}

	if param.Handler != "" {
		args = append(args, "handler")
	}

	if param.Operdate != "" {
		args = append(args, "operdate")
	}

	if param.Paymethod != 0 {
		args = append(args, "paymethod")
	}

	if param.Remark != "" {
		args = append(args, "remark")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}

	return args
}

func AddFinanceflow(financeflow Financeflow) (errorCode int64) {
	var (
		temp Financeflow
	)
	temp.Id = financeflow.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Info("financeflow have asixt")
		errorCode = util.Financeflow_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&financeflow)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Financeflow_ADD_FAILED
	}

	return errorCode
}

func DeleteFinanceflow(id int64) (errorCode int64) {
	var (
		financeflow Financeflow
	)
	errorCode = util.SUCESSFUL
	financeflow.Id = id
	_, err := OSQL.Delete(&financeflow, "id")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Financeflow_DELETE_FAILED
	}

	return errorCode
}
