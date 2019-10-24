package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

type Warehouse struct {
	Id   int64  `json:"id" orm:"column(id)"`
	Name string `json:"name" orm:"column(name)"` //仓库名称
}

func GetAllWarehouses() []Warehouse {
	var (
		wares []Warehouse
	)

	_, err := OSQL.Raw("select * from " + util.WAREHOUSE_TABLE_NAME + " order by id asc").QueryRows(&wares)
	if err != nil {
		beego.Error(err)
	}
	return wares
}

func GetWarehouseBypage(pageNum, pageSize int64) []Warehouse {
	var (
		wares []Warehouse
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.WAREHOUSE_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&wares)
	if err != nil {
		beego.Error(err)
	}
	return wares
}

func GetWarehouseById(id int64) (ware Warehouse, err error) {
	ware.Id = id
	err = OSQL.Read(&ware, "id")
	if err != nil {
		beego.Error(err)
		return ware, err
	}
	return ware, nil
}

func EditWarehouseById(ware Warehouse) (errorCode int64) {
	var (
		temp Warehouse
	)
	errorCode = util.SUCESSFUL
	temp.Id = ware.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.WAREHOUSE_EDIT_FAILED
		return errorCode
	}
	args := editArgs_Ware(ware)
	num, err2 := OSQL.Update(&ware, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.WAREHOUSE_EDIT_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func editArgs_Ware(ware Warehouse) []string {
	var (
		args []string
	)
	if ware.Name != "" {
		args = append(args, "name")
	}

	return args
}

func AddWarehouse(ware Warehouse) (errorCode int64) {
	var (
		temp Warehouse
	)
	errorCode = util.SUCESSFUL
	temp.Id = ware.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		beego.Error("ware have this id=", ware.Id)
		errorCode = util.WAREHOUSE_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&ware)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.WAREHOUSE_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func DeleteWarehouse(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Warehouse
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		beego.Error(err)
		errorCode = util.WAREHOUSE_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
