package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `inquirydetail`;
// CREATE TABLE `inquirydetail` (
//   `inquirycode` bigint(20) NOT NULL COMMENT '询价单号',
//   `mattercode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编号',
//   `num` bigint(20) NOT NULL COMMENT '数量',
//   `price` bigint(20) DEFAULT NULL COMMENT '价格',
//   PRIMARY KEY (`inquirycode`,`mattercode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='询价详情表';

//财务流水表
type Inquirydetail struct {
	Inquirycode int64  `json:"inquirycode" orm:"column(inquirycode)"` //询价单号
	Mattercode  string `json:"mattercode" orm:"column(mattercode)"`   //物料编号
	Num         int64  `json:"num" orm:"column(num)"`                 //数量
	Price       int64  `json:"price" orm:"column(price)"`             //价格
}

func GetInquirydetailBypage(pageNum, pageSize int64) []Inquirydetail {
	var (
		inquirydetails []Inquirydetail
	)
	err := OSQL.Raw("select * from "+util.Inquirydetail_TABLE_NAME+" order by id desc limit ?,?",
		pageNum, pageSize).QueryRow(&inquirydetails)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return inquirydetails
}

func GetInquirydetailByUserID(inquirycode int64) (inquirydetail Inquirydetail, err error) {
	inquirydetail.Inquirycode = inquirycode
	err = OSQL.Read(&inquirydetail, "inquirycode")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return inquirydetail, nil
}

func EditInquirydetailById(inquirydetail Inquirydetail) (errorCode int64) {
	var (
		temp Inquirydetail
	)
	temp.Inquirycode = inquirydetail.Inquirycode
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "inquirycode")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Inquirydetail_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&inquirydetail)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Inquirydetail_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func AddInquirydetail(inquirydetail Inquirydetail) (errorCode int64) {
	var (
		temp Inquirydetail
	)
	temp.Inquirycode = inquirydetail.Inquirycode
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "inquirycode")
	if err == nil {
		logs.FileLogs.Info("inquirydetail have asixt")
		errorCode = util.Inquirydetail_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&inquirydetail)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Inquirydetail_ADD_FAILED
	}

	return errorCode
}

func DeleteInquirydetail(inquirycode int64) (errorCode int64) {
	var (
		inquirydetail Inquirydetail
	)
	errorCode = util.SUCESSFUL
	inquirydetail.Inquirycode = inquirycode
	_, err := OSQL.Delete(&inquirydetail, "inquirycode")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Inquirydetail_DELETE_FAILED
	}

	return errorCode
}
