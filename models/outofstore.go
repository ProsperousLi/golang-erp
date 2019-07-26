package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `outofstore`;
// CREATE TABLE `outofstore` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `outcode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '出库单编号',
//   `warehouseid` bigint(20) NOT NULL COMMENT '仓库id',
//   `type` tinyint(5) NOT NULL COMMENT '类型(1：维修领料；2：销售出库；3：调拨出库)',
//   `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联单据号(当类型为销售出库时，此为销售合同编号，决定出库列表)',
//   `outdate` datetime NOT NULL COMMENT '出库时间',
//   `storehandler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '出库操作人',
//   `pickhandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '领料人',
//   `contractcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '维修合同编号',
//   `vehiclecode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '车辆编号',
//   `itemname` varchar(1000) CHARACTER SET utf8mb4 NOT NULL COMMENT '维修项目名称',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='物料出库表';

type Outofstore struct {
	Id           int64  `json:"id" orm:"column(id)"`
	Outcode      string `json:"outcode" orm:"column(outcode)"`           //出库单编号
	Warehouseid  int64  `json:"warehouseid" orm:"column(warehouseid)"`   //仓库id
	Type         int64  `json:"type" orm:"column(type)"`                 //类型(1：维修领料；2：销售出库；3：调拨出库)
	Relatedcode  string `json:"price" orm:"column(price)"`               //关联单据号(当类型为销售出库时，此为销售合同编号，决定出库列表)
	Outdate      string `json:"outdate" orm:"column(outdate)"`           //出库时间
	Storehandler string `json:"storehandler" orm:"column(storehandler)"` //出库操作人
	Pickhandler  string `json:"pickhandler" orm:"column(pickhandler)"`   //领料人
	Contractcode string `json:"contractcode" orm:"column(contractcode)"` //维修合同编号
	Vehiclecode  string `json:"vehiclecode" orm:"column(vehiclecode)"`   //车辆编号
	Itemname     string `json:"itemname" orm:"column(itemname)"`         //维修项目名称

}

func GetOutofstoreBypage(pageNum, pageSize int64) []Outofstore {
	var (
		outofstores []Outofstore
	)
	err := OSQL.Raw("select * from "+util.Outofstore_TABLE_NAME+" order by id desc limit ?,?",
		pageNum, pageSize).QueryRow(&outofstores)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return outofstores
}

func GetOutofstoreById(id int64) (outofstore Outofstore, err error) {
	outofstore.Id = id
	err = OSQL.Read(&outofstore, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return outofstore, err
	}
	return outofstore, nil
}

func EditOutofstoreById(outofstore Outofstore) (errorCode int64) {
	var (
		temp Outofstore
	)
	errorCode = util.SUCESSFUL
	temp.Id = outofstore.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Outofstore_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&outofstore)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Outofstore_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func AddOutofstore(outofstore Outofstore) (errorCode int64) {
	var (
		temp Outofstore
	)
	errorCode = util.SUCESSFUL
	temp.Id = outofstore.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("outofstore have this id=%v", outofstore.Id)
		errorCode = util.Outofstore_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&outofstore)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Outofstore_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteOutofstore(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Outofstore
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Outofstore_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}