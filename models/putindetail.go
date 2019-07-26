package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `putindetail`;
// CREATE TABLE `putindetail` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `incode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '入库单编号(同putinstore的incode字段)',
//   `mattercode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
//   `realnum` bigint(20) NOT NULL COMMENT '入库数量',
//   `num` bigint(20) NOT NULL COMMENT '到货单数量',
//   `price` bigint(20) NOT NULL COMMENT '单价',
//   `value` bigint(20) NOT NULL COMMENT '总价',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='入库详情表';

type Putindetail struct {
	Id         int64  `json:"id" orm:"column(id)"`
	Incode     string `json:"incode" orm:"column(incode)"`         //入库单编号(同putinstore的incode字段)
	Mattercode string `json:"mattercode" orm:"column(mattercode)"` //物料编码
	Realnum    int64  `json:"realnum" orm:"column(realnum)"`       //采购数量
	Num        int64  `json:"num" orm:"column(num)"`               //入库数量
	Price      int64  `json:"price" orm:"column(price)"`           //到货单数量
	Value      int64  `json:"value" orm:"column(value)"`           //总价
}

func GetPutindetailBypage(pageNum, pageSize int64) []Putindetail {
	var (
		putindetails []Putindetail
	)
	err := OSQL.Raw("select * from "+util.Putindetail_TABLE_NAME+" order by id desc limit ?,?",
		pageNum, pageSize).QueryRow(&putindetails)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return putindetails
}

func GetPutindetailById(id int64) (putindetail Putindetail, err error) {
	putindetail.Id = id
	err = OSQL.Read(&putindetail, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return putindetail, err
	}
	return putindetail, nil
}

func EditPutindetailById(putindetail Putindetail) (errorCode int64) {
	var (
		temp Putindetail
	)
	errorCode = util.SUCESSFUL
	temp.Id = putindetail.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Putindetail_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&putindetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Putindetail_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddPutindetail(putindetail Putindetail) (errorCode int64) {
	var (
		temp Putindetail
	)
	errorCode = util.SUCESSFUL
	temp.Id = putindetail.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("putindetail have this id=%v", putindetail.Id)
		errorCode = util.Putindetail_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&putindetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Putindetail_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeletePutindetail(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Putindetail
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Putindetail_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
