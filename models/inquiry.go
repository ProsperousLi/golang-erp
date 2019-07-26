package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `inquiry`;
// CREATE TABLE `inquiry` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `inquirycode` bigint(20) DEFAULT NULL COMMENT '询价单号',
//   `type` tinyint(5) NOT NULL COMMENT '类型(1：维修；2：销售)',
//   `custcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '客户编码',
//   `createat` datetime NOT NULL COMMENT '制单日期',
//   `deadline` date DEFAULT NULL COMMENT '截止日期',
//   `status` tinyint(5) NOT NULL COMMENT '状态(1：未回复；2：已回复)',
//   `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '制单人',
//   `replyhandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '回复人',
//   `replydate` date DEFAULT NULL COMMENT '回复日期',
//   `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
//   `attachment` text CHARACTER SET utf8mb4 COMMENT '附件',
//   `validity` int(5) DEFAULT NULL COMMENT '价格有效期(单位：月)',
//   PRIMARY KEY (`id`),
//   UNIQUE KEY `inquirycode` (`inquirycode`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='询价表';

//财务流水表
type Inquiry struct {
	Id           int64  `json:"id" orm:"column(id)"`                     //id
	Inquirycode  int64  `json:"inquirycode" orm:"column(inquirycode)"`   //询价单号
	Type         int    `json:"type" orm:"column(type)"`                 //类型(1：维修；2：销售)
	Custcode     string `json:"custcode" orm:"column(custcode)"`         //客户编码
	Createat     string `json:"createat" orm:"column(createat)"`         //制单日期
	Deadline     string `json:"deadline" orm:"column(deadline)"`         //截止日期
	Status       int    `json:"handler" orm:"column(handler)"`           //状态(1：未回复；2：已回复)
	Handler      string `json:"remark" orm:"column(remark)"`             //制单人
	Replyhandler string `json:"replyhandler" orm:"column(replyhandler)"` //回复人
	Replydate    string `json:"replydate" orm:"column(replydate)"`       //回复日期
	Remark       string `json:"remark" orm:"column(remark)"`             //备注
	Attachment   string `json:"attachment" orm:"column(attachment)"`     //附件
	Validity     int    `json:"validity" orm:"column(validity)"`         //价格有效期(单位：月)
}

func GetInquiryBypage(pageNum, pageSize int64) []Inquiry {
	var (
		inquirys []Inquiry
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Inquiry_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&inquirys)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return inquirys
}

func GetInquiryByUserID(id int64) (inquiry Inquiry, err error) {
	inquiry.Id = id
	err = OSQL.Read(&inquiry, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return inquiry, nil
}

func EditInquiryById(inquiry Inquiry) (errorCode int64) {
	var (
		temp Inquiry
	)
	temp.Id = inquiry.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Inquiry_EDIT_FAILED
		return errorCode
	}

	args := edit_InquiryArgs(inquiry)
	num, err2 := OSQL.Update(&inquiry, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Inquiry_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func edit_InquiryArgs(param Inquiry) (args []string) {
	if param.Attachment != "" {
		args = append(args, "attachment")
	}
	if param.Createat != "" {
		args = append(args, "createat")
	}

	if param.Custcode != "" {
		args = append(args, "custcode")
	}

	if param.Deadline != "" {
		args = append(args, "deadline")
	}

	if param.Handler != "" {
		args = append(args, "handler")
	}

	if param.Inquirycode != 0 {
		args = append(args, "inquirycode")
	}

	if param.Remark != "" {
		args = append(args, "remark")
	}

	if param.Replydate != "" {
		args = append(args, "replydate")
	}

	if param.Replyhandler != "" {
		args = append(args, "replyhandler")
	}

	if param.Status != 0 {
		args = append(args, "status")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}

	if param.Validity != 0 {
		args = append(args, "validity")
	}
	return args
}

func AddInquiry(inquiry Inquiry) (errorCode int64) {
	var (
		temp Inquiry
	)
	temp.Id = inquiry.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Info("inquiry have asixt")
		errorCode = util.Inquiry_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&inquiry)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Inquiry_ADD_FAILED
	}

	return errorCode
}

func DeleteInquiry(id int64) (errorCode int64) {
	var (
		inquiry Inquiry
	)
	errorCode = util.SUCESSFUL
	inquiry.Id = id
	_, err := OSQL.Delete(&inquiry, "id")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Inquiry_DELETE_FAILED
	}

	return errorCode
}
