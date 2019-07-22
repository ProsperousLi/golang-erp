package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `stock`;
// CREATE TABLE `stock` (
//   `warehouseid` bigint(20) NOT NULL COMMENT '仓库id',
//   `mattercode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '物料编码',
//   `num` bigint(20) NOT NULL COMMENT '数量',
//   `averageprice` bigint(20) NOT NULL COMMENT '均价',
//   PRIMARY KEY (`warehouseid`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='库存表';

//库存表
type Stock struct {
	Warehouseid  int64  `json:"warehouseid" orm:"pk;column(warehouseid)"` //仓库id
	Mattercode   string `json:"mattercode" orm:"column(mattercode)"`      //物料编码
	Num          int64  `json:"num" orm:"column(num)"`                    //数量
	Averageprice int64  `json:"averageprice" orm:"column(averageprice)"`  //均价
}

func GetStockBypage(pageNum, pageSize int64) []Stock {
	var (
		stocks []Stock
	)
	err := OSQL.Raw("select * from "+util.Stock_TABLE_NAME+" order by id asc limit ?,?",
		pageNum, pageSize).QueryRow(&stocks)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return stocks
}

func GetStockByUserID(warehouseid int64) (stock Stock, err error) {
	stock.Warehouseid = warehouseid
	err = OSQL.Read(&stock, "warehouseid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return stock, nil
}

func EditStockById(stock Stock) (errorCode int64) {
	var (
		temp Stock
	)
	temp.Warehouseid = stock.Warehouseid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "warehouseid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Stock_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&stock)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Stock_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func AddStock(stock Stock) (errorCode int64) {
	var (
		temp Stock
	)
	temp.Warehouseid = stock.Warehouseid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "warehouseid")
	if err == nil {
		logs.FileLogs.Info("stock have asixt")
		errorCode = util.Stock_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&stock)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Stock_ADD_FAILED
	}

	logs.FileLogs.Info("num=%v", id)

	return errorCode
}

func DeleteStock(warehouseid int64) (errorCode int64) {
	var (
		stock Stock
	)
	errorCode = util.SUCESSFUL
	stock.Warehouseid = warehouseid
	num, err := OSQL.Delete(&stock, "warehouseid")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Stock_DELETE_FAILED
	}

	logs.FileLogs.Info("num=%v", num)

	return errorCode
}
