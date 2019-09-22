package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `repairitem`;
// CREATE TABLE `repairitem` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `contractcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '合同编号',
//   `itemname` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '项目名称',
//   `status` tinyint(5) NOT NULL COMMENT '状态(3:进行中；4：完工)',
//   `fault` text CHARACTER SET utf8mb4 COMMENT '故障现象',
//   `causeanalysis` text CHARACTER SET utf8mb4 COMMENT '原因分析',
//   `measures` text CHARACTER SET utf8mb4 COMMENT '修复措施',
//   `vehiclecode` varchar(20) COLLATE utf8_bin NOT NULL,
//   PRIMARY KEY (`contractcode`,`itemname`,`vehiclecode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='车辆维修表';

//维修合同表
type Repairitem struct {
	Id            int64  `json:"id" orm:"column(id)"`
	Contractcode  string `json:"contractcode" orm:"column(contractcode)"`
	Itemname      string `json:"itemname" orm:"column(itemname)"`           //合同编号
	Status        int8   `json:"status" orm:"column(status)"`               //项目名称
	Fault         string `json:"fault" orm:"column(fault)"`                 //状态(3:进行中；4：完工)
	Causeanalysis string `json:"causeanalysis" orm:"column(causeanalysis)"` //故障现象
	Measures      string `json:"measures" orm:"column(measures)"`           //原因分析
	Vehiclecode   string `json:"vehiclecode" orm:"column(vehiclecode)"`     //修复措施
}

func GetRepairitemBCode(contractcode, vehiclecode string) []Repairitem {
	var (
		params []Repairitem
	)
	_, err := OSQL.Raw("select * from "+util.Repairitem_TABLE_NAME+
		"where contractcode='?' and vehiclecode='?' order by id desc",
		contractcode, vehiclecode).QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

func GetRepairitemById(contractcode string) (ret Repairitem, err error) {
	ret.Contractcode = contractcode
	err = OSQL.Read(&ret, "contractcode")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ret, err
	}
	return ret, nil
}

func EditRepairitemById(param Repairitem) (errorCode int64) {
	var (
		temp Repairitem
	)
	errorCode = util.SUCESSFUL
	temp.Contractcode = param.Contractcode
	temp.Vehiclecode = param.Vehiclecode
	temp.Itemname = param.Itemname
	err := OSQL.Read(&temp, "contractcode", "vehiclecode", "itemname")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Repairitem_EDIT_FAILED
		return errorCode
	}

	args := edit_repairitem(param)

	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Repairitem_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_repairitem(param Repairitem) (args []string) {
	if param.Causeanalysis != "" {
		args = append(args, "causeanalysis")
	}

	if param.Fault != "" {
		args = append(args, "fault")
	}

	if param.Itemname != "" {
		args = append(args, "itemname")
	}

	if param.Measures != "" {
		args = append(args, "measures")
	}

	if param.Status != 0 {
		args = append(args, "status")
	}

	if param.Vehiclecode != "" {
		args = append(args, "vehiclecode")
	}
	return args
}

func AddRepairitem(param Repairitem) (errorCode int64) {
	var (
		temp Repairitem
	)
	errorCode = util.SUCESSFUL
	temp.Contractcode = param.Contractcode
	err := OSQL.Read(&temp, "contractcode")
	if err == nil {
		logs.FileLogs.Error("Repairitem have this contractcode=%v", param.Contractcode)
		errorCode = util.FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteRepairitem(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Repairitem
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Repairitem_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
