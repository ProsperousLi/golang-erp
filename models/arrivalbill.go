package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `arrivalbill`;
// CREATE TABLE `arrivalbill` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `arrivalbillcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '到货单编码',
//   `contractcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购合同编号',
//   `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '制单人',
//   `status` tinyint(5) NOT NULL COMMENT '状态(1:制单;2:已入库)',
//   `warehouseid` bigint(20) NOT NULL COMMENT '仓库id(仓库表主键)',
//   `createat` datetime NOT NULL COMMENT '制单时间',
//   `indate` datetime DEFAULT NULL COMMENT '入库日期',
//   `storehandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '入库操作员',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='到货订单表';

type Arrivalbill struct {
	Id              int64  `json:"id" orm:"column(id)"`                           //id
	Arrivalbillcode string `json:"arrivalbillcode" orm:"column(arrivalbillcode)"` //到货单编码
	Contractcode    string `json:"contractcode" orm:"column(contractcode)"`       //采购合同编号
	Handler         string `json:"handler" orm:"column(handler)"`                 //制单人
	Status          int    `json:"status" orm:"column(status)"`                   //状态(1:制单;2:已入库)
	Warehouseid     int64  `json:"warehouseid" orm:"column(warehouseid)"`         //仓库id(仓库表主键)
	Createat        string `json:"createat" orm:"column(createat)"`               //制单时间
	Indate          string `json:"indate" orm:"column(indate)"`                   //入库日期
	Storehandler    string `json:"storehandler" orm:"column(storehandler)"`       //入库操作员
}

type QueryArrivalBillStruct struct {
	Arrivalbillcode string
	Handler         string
	Contractcode    string
	Datebegin       string
	Dateend         string
	Pageno          int64
	Pagesize        int64
}

func QueryArrivalBill(param QueryArrivalBillStruct) ([]Arrivalbill, int64) {
	var (
		rets []Arrivalbill
	)

	sql := "select * from " + util.Arrivalbill_TABLE_NAME + "where 1=1 "

	if param.Arrivalbillcode != "" {
		sql += " and arrivalbillcode='" + param.Arrivalbillcode + "' "
	}

	if param.Handler != "" {
		sql += " and handler='" + param.Handler + "' "
	}

	if param.Contractcode != "" {
		sql += " and contractcode='" + param.Contractcode + "' "
	}

	if param.Datebegin != "" {
		sql += " and indate>=" + param.Datebegin
	}

	if param.Dateend != "" {
		sql += " and indate<=" + param.Dateend
	}

	begin := param.Pageno * param.Pagesize
	_, err := OSQL.Raw(sql+" order by id desc limit ?,?",
		begin, param.Pagesize).QueryRows(&rets)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}

	allNums, err := OSQL.QueryTable(util.Marketcontract_TABLE_NAME).Count()
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return rets, allNums
}

func GetArrivalbillBypage(arrivalbillcode string) []Arrivalbill {
	var (
		arrivalbills []Arrivalbill
	)
	_, err := OSQL.Raw("select * from " + util.Arrivalbill_TABLE_NAME +
		" and where arrivalbillcode='" + arrivalbillcode + "' order by id desc").QueryRows(&arrivalbills)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return arrivalbills
}

func GetArrivalbillByUserID(id int64) (arrivalbill Arrivalbill, err error) {
	arrivalbill.Id = id
	err = OSQL.Read(&arrivalbill, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return arrivalbill, nil
}

func EditArrivalbillById(arrivalbill Arrivalbill) (errorCode int64) {
	var (
		temp Arrivalbill
	)
	temp.Id = arrivalbill.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Arrivalbill_EDIT_FAILED
		return errorCode
	}

	args := edit_arrivalbill(arrivalbill)
	num, err2 := OSQL.Update(&arrivalbill, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Arrivalbill_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func edit_arrivalbill(param Arrivalbill) []string {
	var (
		args []string
	)

	if param.Arrivalbillcode != "" {
		args = append(args, "arrivalbillcode")
	}

	if param.Contractcode != "" {
		args = append(args, "contractcode")
	}

	if param.Createat != "" {
		args = append(args, "createat")
	}

	if param.Handler != "" {
		args = append(args, "handler")
	}

	if param.Status != 0 {
		args = append(args, "status")
	}

	if param.Storehandler != "" {
		args = append(args, "storehandler")
	}

	if param.Warehouseid != 0 {
		args = append(args, "warehouseid")
	}

	if param.Indate != "" {
		args = append(args, "indate")
	}

	return args
}

func AddArrivalbill(arrivalbill Arrivalbill) (errorCode, id int64) {
	var (
		temp Arrivalbill
	)
	temp.Id = arrivalbill.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Info("arrivalbill have asixt")
		errorCode = util.Arrivalbill_ADD_FAILED
		return errorCode, id
	}

	num, err2 := OSQL.Insert(&arrivalbill)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Arrivalbill_ADD_FAILED
		return errorCode, num
	}

	return errorCode, num
}

func DeleteArrivalbill(id int64) (errorCode int64) {
	var (
		arrivalbill Arrivalbill
	)
	errorCode = util.SUCESSFUL
	arrivalbill.Id = id
	_, err := OSQL.Delete(&arrivalbill, "id")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Arrivalbill_DELETE_FAILED
	}

	return errorCode
}
