package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `saledetail`;
// CREATE TABLE `saledetail` (
//   `contractid` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '合同编号',
//   `mattercode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
//   `num` bigint(20) NOT NULL,
//   `price` bigint(20) NOT NULL COMMENT '单价',
//   PRIMARY KEY (`contractid`,`mattercode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='销售详情表';

//销售详情表
type Saledetail struct {
	Contractid string `json:"contractid" orm:"pk;column(contractid)"` //合同编号
	Mattercode string `json:"mattercode" orm:"column(mattercode)"`    //物料编码
	Num        int64  `json:"num" orm:"column(num)"`                  //数量
	Price      int64  `json:"price" orm:"column(price)"`              //单价
}

func GetSaledetailBypage(pageNum, pageSize int64) []Saledetail {
	var (
		saledetails []Saledetail
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Saledetail_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&saledetails)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return saledetails
}

func GetSaledetailById(contractid string) (saledetail Saledetail, err error) {
	saledetail.Contractid = contractid
	err = OSQL.Read(&saledetail, "contractid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return saledetail, nil
}

func EditSaledetailById(saledetail Saledetail) (errorCode int64) {
	var (
		temp Saledetail
	)
	temp.Contractid = saledetail.Contractid
	errorCode = util.SUCESSFUL

	err := OSQL.Read(&temp, "contractid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Saledetail_EDIT_FAILED
		return errorCode
	}

	args := edit_saledetail(saledetail)
	num, err2 := OSQL.Update(&saledetail, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Saledetail_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

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
	temp.Contractid = saledetail.Contractid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "contractid")
	if err == nil {
		logs.FileLogs.Info("saledetail have asixt")
		errorCode = util.Saledetail_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&saledetail)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Saledetail_ADD_FAILED
	}

	logs.FileLogs.Info("num=%v", id)

	return errorCode
}

func DeleteSaledetail(contractid string) (errorCode int64) {
	var (
		saledetail Saledetail
	)
	errorCode = util.SUCESSFUL
	saledetail.Contractid = contractid
	num, err := OSQL.Delete(&saledetail, "contractid")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Saledetail_DELETE_FAILED
	}

	logs.FileLogs.Info("num=%v", num)

	return errorCode
}
