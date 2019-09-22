package models

import (
	"erpweb/logs"
	//"erpweb/util"
)

// 参数: type=1获取员工编号(人员信息表),type=2维修合同编号(采购合同表),
// type=3销售合同编号(采购合同表),type=4 到货单号(物料入库表)
func QueryTimeStamp(queryType string) interface{} {
	var retCode []string
	if queryType == "1" { //获取员工编号
		emps := GetAllEmployees("", "")
		for _, emp := range emps {
			retCode = append(retCode, emp.Cardid)
		}
	} else if queryType == "2" { //维修合同编号
		marks := GetMarketcontractByType(1)
		for _, mark := range marks {
			retCode = append(retCode, mark.Contractcode)
		}
	} else if queryType == "3" { //销售合同编号
		marks := GetMarketcontractByType(2)
		for _, mark := range marks {
			retCode = append(retCode, mark.Contractcode)
		}
	} else if queryType == "4" { //到货单号
		puts := GetAllPutinstore()
		for _, put := range puts {
			retCode = append(retCode, put.Incode)
		}
	} else {
		logs.FileLogs.Error("unkonw this type=", queryType)
	}

	return retCode
}
