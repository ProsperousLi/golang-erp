package models

import (
	"strconv"

	"erpweb/util"

	"github.com/astaxie/beego"
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

func QueryStock(warehouseid int64, mattercode string) []Stock {
	var (
		stocks []Stock
	)
	sql := "select * from " + util.Stock_TABLE_NAME + " where 1=1"
	sql += " and warehouseid=" + strconv.FormatInt(warehouseid, 10)
	if mattercode != "" {
		sql += " and mattercode='" + mattercode + "'"
	}
	_, err := OSQL.Raw(sql + " order by warehouseid asc").QueryRows(&stocks)
	if err != nil {
		beego.Error(err)
	}
	return stocks
}

func GetStockBypage(pageNum, pageSize int64) []Stock {
	var (
		stocks []Stock
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Stock_TABLE_NAME+" limit ?,?",
		begin, pageSize).QueryRows(&stocks)
	if err != nil {
		beego.Error(err)
	}
	return stocks
}

func GetStockById(warehouseid int64) (stock Stock, err error) {
	stock.Warehouseid = warehouseid
	err = OSQL.Read(&stock, "warehouseid")
	if err != nil {
		beego.Error(err)
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
		beego.Error(err)
		errorCode = util.Stock_EDIT_FAILED
		return errorCode
	}

	args := edit_stock(stock)

	num, err2 := OSQL.Update(&stock, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Stock_EDIT_FAILED
		return errorCode
	}

	beego.Info("num= err=", num, err2)

	return errorCode
}

func edit_stock(param Stock) (args []string) {
	if param.Averageprice != 0 {
		args = append(args, "averageprice")
	}

	if param.Mattercode != "" {
		args = append(args, "mattercode")
	}

	if param.Num != 0 {
		args = append(args, "num")
	}
	return args
}

func AddStock(stock Stock) (errorCode int64) {
	var (
		temp Stock
	)
	temp.Warehouseid = stock.Warehouseid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "warehouseid")
	if err == nil {
		beego.Info("stock have asixt")
		errorCode = util.Stock_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&stock)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Stock_ADD_FAILED
	}

	beego.Info("num=", id)

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
		beego.Error(err)
		errorCode = util.Stock_DELETE_FAILED
	}

	beego.Info("num=", num)

	return errorCode
}
