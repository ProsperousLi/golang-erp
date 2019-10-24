package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

// DROP TABLE IF EXISTS `saledetail`;
// CREATE TABLE `saledetail` (
//   `contractcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '合同编号',
//   `mattercode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
//   `num` bigint(20) NOT NULL,
//   `price` bigint(20) NOT NULL COMMENT '单价',
//   PRIMARY KEY (`contractid`,`mattercode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='销售详情表';

//销售详情表
type Saledetail struct {
	Contractcode string `json:"contractcode" orm:"pk;column(contractcode)"` //合同编号
	Mattercode   string `json:"mattercode" orm:"column(mattercode)"`        //物料编码
	Num          int64  `json:"num" orm:"column(num)"`                      //数量
	Price        int64  `json:"price" orm:"column(price)"`                  //单价
}

type AddAndUpdateSaledetailStruct struct {
	Contractcode string
	Detail       []Saledetail
}

func AddOrUpdateSaleDetail(param AddAndUpdateSaledetailStruct) (errorCode int64, msg string) {
	errorCode = util.SUCESSFUL
	details := GetSaledetailByContractcode(param.Contractcode)
	if len(details) > 0 {
		errorCode = DeleteSaledetail(param.Contractcode)
	} else {
		errorCode = util.FAILED
		msg = "获取数据失败"
		return
	}

	if errorCode != util.SUCESSFUL {
		errorCode = util.FAILED
		msg = "删除数据失败"
		return
	}

	//add
	if len(param.Detail) > 0 {
		for _, detail := range param.Detail {
			errorCode = AddSaledetail(detail)
			if errorCode != util.SUCESSFUL {
				msg = "删除数据失败"
				return
			}
		}

	}

	return
}

func GetSaledetailByContractcode(contractcode string) []Saledetail {
	var (
		saledetails []Saledetail
	)
	_, err := OSQL.Raw("select * from "+util.Saledetail_TABLE_NAME+
		" where contractcode=? order by id asc", contractcode).QueryRows(&saledetails)
	if err != nil {
		beego.Error(err)
	}
	return saledetails
}

func GetSaledetailById(Contractcode string) (saledetail Saledetail, err error) {
	saledetail.Contractcode = Contractcode
	err = OSQL.Read(&saledetail, "contractcode")
	if err != nil {
		beego.Error(err)
	}
	return saledetail, nil
}

func EditSaledetailById(saledetail Saledetail) (errorCode int64) {
	var (
		temp Saledetail
	)
	temp.Contractcode = saledetail.Contractcode
	errorCode = util.SUCESSFUL

	err := OSQL.Read(&temp, "contractcode")
	if err != nil {
		beego.Error(err)
		errorCode = util.FAILED
		return errorCode
	}

	args := edit_saledetail(saledetail)
	num, err2 := OSQL.Update(&saledetail, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.FAILED
		return errorCode
	}

	beego.Info("num= err=", num, err2)

	return errorCode
}

func edit_saledetail(param Saledetail) (args []string) {

	if param.Mattercode != "" {
		args = append(args, "mattercode")
	}

	if param.Num != 0 {
		args = append(args, "num")
	}

	if param.Price != 0 {
		args = append(args, "price")
	}
	return args
}

func AddSaledetail(saledetail Saledetail) (errorCode int64) {
	var (
		temp Saledetail
	)
	temp.Contractcode = saledetail.Contractcode
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "contractcode")
	if err == nil {
		beego.Info("saledetail have asixt")
		errorCode = util.FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&saledetail)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.FAILED
	}

	beego.Info("num=", id)

	return errorCode
}

func DeleteSaledetail(contractcode string) (errorCode int64) {
	var (
		saledetail Saledetail
	)
	errorCode = util.SUCESSFUL
	saledetail.Contractcode = contractcode
	num, err := OSQL.Delete(&saledetail, "contractcode")
	if err != nil {
		beego.Error(err)
		errorCode = util.FAILED
	}

	beego.Info("num=", num)

	return errorCode
}
