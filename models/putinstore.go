package models

import (
	"strconv"

	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `putinstore`;
// CREATE TABLE `putinstore` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `incode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '入库单编号',
//   `warehouseid` bigint(20) NOT NULL COMMENT '仓库id',
//   `source` int(5) NOT NULL COMMENT '来源(1：采购到货单；2：调拨入库；3：盘亏盘盈)',
//   `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联单据号',
//   `indate` datetime NOT NULL COMMENT '入库时间',
//   `storehandler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '入库操作人(cardid)',
//   `purchasehandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '采购处理人(cardid)',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='物料入库表';

//入库详情表
type Putinstore struct {
	Id              int64  `json:"id" orm:"column(id)"`
	Incode          string `json:"incode" orm:"column(incode)"`                   //入库单编号
	Warehouseid     int64  `json:"warehouseid" orm:"column(warehouseid)"`         //仓库id
	Source          int8   `json:"source" orm:"column(source)"`                   //来源(1：采购到货单；2：调拨入库；3：盘亏盘盈)
	Relatedcode     string `json:"relatedcode" orm:"column(relatedcode)"`         //关联单据号
	Indate          string `json:"indate" orm:"column(indate)"`                   //入库时间
	Storehandler    string `json:"storehandler" orm:"column(storehandler)"`       //入库操作人(cardid)
	Purchasehandler string `json:"purchasehandler" orm:"column(purchasehandler)"` //采购处理人(cardid)
}

func PutinStore(param Putinstore) (errorCode, lastId int64) {
	var (
		temp Putinstore
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("putinstore have this id=%v", param.Id)
		errorCode = util.Putinstore_ADD_FAILED
		return errorCode, lastId
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Putinstore_ADD_FAILED
		return errorCode, num
	}
	logs.FileLogs.Info("num=%v", num)

	lastId = num
	return errorCode, num
}

//SELECT LAST_INSERT_ID() last_id
func GetAllPutinstore() []Putinstore {
	var (
		params []Putinstore
	)
	_, err := OSQL.Raw("select * from " + util.Putinstore_TABLE_NAME + " order by id desc").QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

//incode=xxx&warehouseid=xxx&datebegin=xxx&dateend=xxx&pageno=1&pagesize=10
//参数warehouseid、pageno和pagesize必须传，incode和datebegin至少传一个
type QueryPutistoreStruct struct {
	Incode      string
	Warehouseid int64
	Datebegin   string
	Dateend     string
	Pageno      int64
	Pagesize    int64
}

func GetPutinstoreBypage(param QueryPutistoreStruct) []Putinstore {
	var (
		rets []Putinstore
	)

	sql := "select * from " + util.Putinstore_TABLE_NAME + "where 1=1 "

	if param.Incode != "" {
		sql += " and incode='" + param.Incode + "' "
	}

	sql += " and warehouseid=" + strconv.FormatInt(param.Warehouseid, 10)
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
	return rets
}

func GetPutinstoreById(id int64) (ret Putinstore, err error) {
	ret.Id = id
	err = OSQL.Read(&ret, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ret, err
	}
	return ret, nil
}

func EditPutinstoreById(param Putinstore) (errorCode int64) {
	var (
		temp Putinstore
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Putinstore_EDIT_FAILED
		return errorCode
	}

	args := edit_putinstore(param)

	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Putinstore_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_putinstore(param Putinstore) (args []string) {
	if param.Incode != "" {
		args = append(args, "incode")
	}

	if param.Indate != "" {
		args = append(args, "indate")
	}

	if param.Purchasehandler != "" {
		args = append(args, "purchasehandler")
	}

	if param.Relatedcode != "" {
		args = append(args, "relatedcode")
	}

	if param.Source != 0 {
		args = append(args, "source")
	}

	if param.Storehandler != "" {
		args = append(args, "storehandler")
	}

	if param.Warehouseid != 0 {
		args = append(args, "warehouseid")
	}
	return args
}

func AddPutinstore(param Putinstore) (errorCode, id int64) {
	var (
		temp Putinstore
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("putinstore have this id=%v", param.Id)
		errorCode = util.Putinstore_ADD_FAILED
		return errorCode, id
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Putinstore_ADD_FAILED
		return errorCode, num
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode, num
}

func DeletePutinstore(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Putinstore
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Putinstore_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
