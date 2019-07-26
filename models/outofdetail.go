package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `outofdetail`;
// CREATE TABLE `outofdetail` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `outcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '出库单编号',
//   `mattercode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '物料编码',
//   `num` bigint(20) NOT NULL COMMENT '出库数量',
//   `price` bigint(20) NOT NULL COMMENT '单价',
//   `value` bigint(20) NOT NULL COMMENT '总价',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='出库详情表';

type Outofdetail struct {
	Id         int64  `json:"id" orm:"column(id)"`
	Outcode    string `json:"outcode" orm:"column(outcode)"`       //出库单编号
	Mattercode string `json:"mattercode" orm:"column(mattercode)"` //物料编码
	Num        int64  `json:"num" orm:"column(num)"`               //出库数量
	Price      int64  `json:"price" orm:"column(price)"`           //单价
	Value      int64  `json:"value" orm:"column(value)"`           //总价
}

func GetOutofdetailBypage(pageNum, pageSize int64) []Outofdetail {
	var (
		outofdetails []Outofdetail
	)
	err := OSQL.Raw("select * from "+util.Outofdetail_TABLE_NAME+" order by id desc limit ?,?",
		pageNum, pageSize).QueryRow(&outofdetails)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return outofdetails
}

func GetOutofdetailById(id int64) (outofdetail Outofdetail, err error) {
	outofdetail.Id = id
	err = OSQL.Read(&outofdetail, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return outofdetail, err
	}
	return outofdetail, nil
}

func EditOutofdetailById(outofdetail Outofdetail) (errorCode int64) {
	var (
		temp Outofdetail
	)
	errorCode = util.SUCESSFUL
	temp.Id = outofdetail.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Outofdetail_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&outofdetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Outofdetail_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddOutofdetail(outofdetail Outofdetail) (errorCode int64) {
	var (
		temp Outofdetail
	)
	errorCode = util.SUCESSFUL
	temp.Id = outofdetail.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("outofdetail have this id=%v", outofdetail.Id)
		errorCode = util.Outofdetail_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&outofdetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Outofdetail_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteOutofdetail(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Outofdetail
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Outofdetail_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
