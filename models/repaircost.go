package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `repaircost`;
// CREATE TABLE `repaircost` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `type` tinyint(5) NOT NULL COMMENT '费用类型(1：物料;2:人工；3：其他；4：外协)',
//   `itemid` bigint(20) DEFAULT NULL COMMENT '维修项id',
//   `extend` text CHARACTER SET utf8mb4 COMMENT '扩展信息(json;如人工的出发和返回时间)',
//   `unit` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '单位',
//   `num` bigint(20) NOT NULL COMMENT '数量',
//   `price` bigint(20) NOT NULL COMMENT '单价',
//   `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='维修项费用表';

//维修合同表
type Repaircost struct {
	Id     int64  `json:"id" orm:"column(id)"`         //合同编号
	Type   int8   `json:"type" orm:"column(type)"`     //费用类型(1：物料;2:人工；3：其他；4：外协)
	Itemid int64  `json:"itemid" orm:"column(itemid)"` //维修项id
	Extend string `json:"extend" orm:"column(extend)"` //扩展信息(json;如人工的出发和返回时间)
	Unit   string `json:"unit" orm:"column(unit)"`     //单位
	Num    int64  `json:"num" orm:"column(num)"`       //数量
	Price  int64  `json:"price" orm:"column(price)"`   //单价
	Remark string `json:"remark" orm:"column(remark)"` //备注
}

type QueryRepaircostStruct struct {
	Itemid int64
	Type   int8
}

func QueryRepairCost(param QueryRepaircostStruct) []Repaircost {
	var (
		params []Repaircost
	)
	_, err := OSQL.Raw("select * from "+util.Repaircost_TABLE_NAME+
		" where itemid=? and type=? order by id desc",
		param.Itemid, param.Type).QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

func GetRepaircostBypage(pageNum, pageSize int64) []Repaircost {
	var (
		params []Repaircost
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Repaircost_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

func GetRepaircostById(id int64) (ret Repaircost, err error) {
	ret.Id = id
	err = OSQL.Read(&ret, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ret, err
	}
	return ret, nil
}

func EditRepaircostById(param Repaircost) (errorCode int64) {
	var (
		temp Repaircost
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Repaircost_EDIT_FAILED
		return errorCode
	}

	args := edit_repaircost(param)

	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Repaircost_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_repaircost(param Repaircost) (args []string) {
	if param.Extend != "" {
		args = append(args, "extend")
	}
	if param.Num != 0 {
		args = append(args, "num")
	}

	if param.Price != 0 {
		args = append(args, "price")
	}

	if param.Remark != "" {
		args = append(args, "remark")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}

	if param.Itemid != 0 {
		args = append(args, "itemid")
	}

	if param.Unit != "" {
		args = append(args, "unit")
	}

	return args
}

func AddRepaircost(param Repaircost) (errorCode int64) {
	var (
		temp Repaircost
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("Repaircost have this id=%v", param.Id)
		errorCode = util.Repaircost_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Repaircost_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteRepaircost(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Repaircontract
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Repaircost_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
