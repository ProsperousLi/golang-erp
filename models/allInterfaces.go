package models

import (
	"erpweb/logs"
	//"erpweb/util"
)

type RetNums struct {
	Nums    int64
	Matters []WebMatterplanAndMatter
}

func QueryNumsOfPurchasecontract(contractcode string) (ret RetNums, err error) {
	var matters []WebMatterplanAndMatter
	purch, err1 := GetPurchasecontractByContractcode(contractcode)
	if err1 != nil {
		return ret, err1
	}

	purchDetail, err1 := GetPurchasedetailById(contractcode)
	if err1 != nil {
		return ret, err1
	}

	if purch.Type == "1" { //销售合同

		ret.Nums = purchDetail.Num
	} else if purch.Type == "2" { //维修合同
		items, err1 := GetRepairitemByItemname(purchDetail.Relatedcode)
		if err1 != nil {
			return ret, err1
		}

		for _, item := range items {
			tempMatters := GetMatterplansByItemid(item.Id)
			if len(tempMatters) > 0 {
				matters = append(matters, tempMatters...)
			}

			for _, tempMatter := range tempMatters {
				ret.Nums += tempMatter.Plannum
			}
		}

		ret.Matters = append(ret.Matters, matters...)
	}

	return ret, nil

}

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

func QueryTimeStampDays(queryType string) int64 {
	var (
		retTimes int64
	)
	if tempTimes, ok := TimeStampMaps[queryType]; ok {
		if tempTimes >= 10000 || tempTimes+1 >= 10000 {
			return 10000
		}
		TimeStampMaps[queryType] = tempTimes + 1
		retTimes = tempTimes + 1
	} else {
		TimeStampMaps[queryType] = TimeStamp
		if TimeStamp >= 10000 {
			return 10000
		}
		retTimes = TimeStamp
	}

	return retTimes
}

//隔天清零
