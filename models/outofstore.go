package models

import (
	"strconv"

	"erpweb/util"

	"github.com/astaxie/beego"
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

type QueryOutofstoreStruct struct {
	Outcode     string
	Warehouseid int64
	Datebegin   string
	Dateend     string
	Pageno      int64
	Pagesize    int64
}

func QueryOutofStore(param QueryOutofstoreStruct) ([]Outofstore, int64) {
	var (
		rets []Outofstore
	)

	sql := "select * from " + util.Outofstore_TABLE_NAME + "where 1=1 "

	if param.Outcode != "" {
		sql += " and outcode='" + param.Outcode + "' "
	}

	sql += " and warehouseid=" + strconv.FormatInt(param.Warehouseid, 10)
	if param.Datebegin != "" {
		sql += " and outdate>=" + param.Datebegin
	}

	if param.Dateend != "" {
		sql += " and outdate<=" + param.Dateend
	}

	begin := param.Pageno * param.Pagesize
	_, err := OSQL.Raw(sql+" order by id desc limit ?,?",
		begin, param.Pagesize).QueryRows(&rets)
	if err != nil {
		beego.Error(err)
	}

	allNums, err := OSQL.QueryTable(util.Marketcontract_TABLE_NAME).Count()
	if err != nil {
		beego.Error(err)
	}
	return rets, allNums
}

func GetOutofstoreBypage(pageNum, pageSize int64) []Outofstore {
	var (
		outofstores []Outofstore
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Outofstore_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&outofstores)
	if err != nil {
		beego.Error(err)
	}
	return outofstores
}

func GetOutofstoreById(id int64) (outofstore Outofstore, err error) {
	outofstore.Id = id
	err = OSQL.Read(&outofstore, "id")
	if err != nil {
		beego.Error(err)
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
		beego.Error(err)
		errorCode = util.Outofstore_EDIT_FAILED
		return errorCode
	}

	args := edit_outofstore(outofstore)

	num, err2 := OSQL.Update(&outofstore, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Outofstore_EDIT_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func edit_outofstore(param Outofstore) (args []string) {
	if param.Contractcode != "" {
		args = append(args, "contractcode")
	}

	if param.Itemname != "" {
		args = append(args, "itemname")
	}

	if param.Outcode != "" {
		args = append(args, "outcode")
	}

	if param.Outdate != "" {
		args = append(args, "outdate")
	}

	if param.Pickhandler != "" {
		args = append(args, "pickhandler")
	}

	if param.Relatedcode != "" {
		args = append(args, "relatedcode")
	}

	if param.Storehandler != "" {
		args = append(args, "storehandler")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}

	if param.Vehiclecode != "" {
		args = append(args, "vehiclecode")
	}

	if param.Warehouseid != 0 {
		args = append(args, "warehouseid")
	}

	return args
}

func AddOutofstore(outofstore Outofstore) (errorCode, id int64) {
	var (
		temp Outofstore
	)
	errorCode = util.SUCESSFUL
	temp.Id = outofstore.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		beego.Error("outofstore have this id=", outofstore.Id)
		errorCode = util.Outofstore_ADD_FAILED
		return errorCode, id
	}

	num, err2 := OSQL.Insert(&outofstore)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Outofstore_ADD_FAILED
		return errorCode, num
	}
	beego.Info("num=", num)
	return errorCode, num
}

func DeleteOutofstore(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Outofstore
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.Outofstore_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
