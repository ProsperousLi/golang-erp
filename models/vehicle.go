package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `vehicle`;
// CREATE TABLE `vehicle` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `custcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '客户编码',
//   `vehiclecode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '车辆编码',
//   `name` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '车辆名称',
//   `type` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '车型',
//   `line` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '线路',
//   `productdate` datetime DEFAULT NULL COMMENT '出厂日期',
//   `manufacturer` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '生产厂家',
//   `remark` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='车辆信息表';

//车辆信息表
type Vehicle struct {
	Id           int64  `json:"id" orm:"column(id)"`
	Custcode     string `json:"custcode" orm:"column(custcode)"`
	Vehiclecode  string `json:"vehiclecode" orm:"column(vehiclecode)"`
	Name         string `json:"name" orm:"column(name)"`
	Type         string `json:"type" orm:"column(type)"`
	Line         string `json:"line" orm:"column(line)"`
	Productdate  string `json:"productdate" orm:"column(productdate)"`
	Manufacturer string `json:"manufacturer" orm:"column(manufacturer)"`
	Remark       string `json:"remark" orm:"column(remark)"`
}

func GetVehicleByCustcode(custcode string) []Vehicle {
	var (
		pas []Vehicle
	)
	err := OSQL.Raw("select * from "+util.VEHICLE_TABLE_NAME+" where custcode=? order by id asc", custcode).QueryRow(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas
}

func GetVehicleBypage(pageNum, pageSize int64) []Vehicle {
	var (
		pas []Vehicle
	)
	err := OSQL.Raw("select * from "+util.VEHICLE_TABLE_NAME+" order by id asc limit ?,?",
		pageNum, pageSize).QueryRow(&pas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return pas
}

func GetVehicleById(id int64) (pa Vehicle, err error) {
	pa.Id = id
	err = OSQL.Read(&pa, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return pa, err
	}
	return pa, nil
}

func EditVehicleById(pa Vehicle) (errorCode int64) {
	var (
		temp Vehicle
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.VEHICLE_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.VEHICLE_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func Edit_args_vehicle(param Vehicle) (args []string) {
	if param.Type != "" {
		args = append(args, "type")
	}

	if param.Custcode != "" {
		args = append(args, "custcode")
	}

	if param.Line != "" {
		args = append(args, "line")
	}

	if param.Manufacturer != "" {
		args = append(args, "manufacturer")
	}

	if param.Name != "" {
		args = append(args, "name")
	}

	if param.Productdate != "" {
		args = append(args, "productdate")
	}

	if param.Remark != "" {
		args = append(args, "remark")
	}

	if param.Vehiclecode != "" {
		args = append(args, "vehiclecode")
	}
	return args
}

func AddVehicle(pa Vehicle) (errorCode int64) {
	var (
		temp Vehicle
	)
	errorCode = util.SUCESSFUL
	temp.Id = pa.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("Vehicle have this id=%v", pa.Id)
		errorCode = util.VEHICLE_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&pa)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.VEHICLE_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteVehicle(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Vehicle
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.VEHICLE_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
