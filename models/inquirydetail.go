package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

// DROP TABLE IF EXISTS `inquirydetail`;
// CREATE TABLE `inquirydetail` (
//   `inquirycode` bigint(20) NOT NULL COMMENT '询价单号',
//   `mattercode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编号',
//   `num` bigint(20) NOT NULL COMMENT '数量',
//   `price` bigint(20) DEFAULT NULL COMMENT '价格',
//   `validity` int(5) DEFAULT NULL COMMENT '价格有效期(单位：月)',
//   PRIMARY KEY (`inquirycode`,`mattercode`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='询价详情表';

//财务流水表
type Inquirydetail struct {
	Inquirycode int64  `json:"inquirycode" orm:"column(inquirycode)"` //询价单号
	Mattercode  string `json:"mattercode" orm:"column(mattercode)"`   //物料编号
	Num         int64  `json:"num" orm:"column(num)"`                 //数量
	Price       int64  `json:"price" orm:"column(price)"`             //价格
	Validity    int    `json:"validity" orm:"column(validity)"`       //价格有效期(单位：月)

}

type InquirydetailWeb struct {
	Inquirycode int64  `json:"inquirycode" orm:"column(inquirycode)"` //询价单号
	Mattercode  string `json:"mattercode" orm:"column(mattercode)"`   //物料编号
	Num         int64  `json:"num" orm:"column(num)"`                 //数量
	Price       int64  `json:"price" orm:"column(price)"`             //价格
	Validity    int    `json:"validity" orm:"column(validity)"`       //价格有效期(单位：月)
	Unit        string `json:"unit"`
	Param       string `json:"param"`
	Name        string `json:"name"`
}

func GetInquirydetailBypage(inquirycode string) (rets []InquirydetailWeb) {
	var (
		inquirydetails []Inquirydetail
	)
	_, err := OSQL.Raw("select * from " + util.Inquirydetail_TABLE_NAME +
		" where inquirycode='" + inquirycode + "' order by id desc").QueryRows(&inquirydetails)
	if err != nil {
		beego.Error(err)
	}

	for _, temp := range inquirydetails {
		var tempRet InquirydetailWeb
		tempRet.Inquirycode = temp.Inquirycode
		tempRet.Mattercode = temp.Mattercode
		tempRet.Validity = temp.Validity
		tempRet.Num = temp.Num
		tempRet.Price = temp.Price

		mat, err := GetMatterByMattercode(temp.Mattercode)
		if err != nil {
			continue
		}

		tempRet.Unit = mat.Unit
		tempRet.Param = mat.Param
		tempRet.Name = mat.Name

		rets = append(rets, tempRet)
	}

	return rets
}

func GetInquirydetailId(inquirycode int64) (inquirydetail Inquirydetail, err error) {
	inquirydetail.Inquirycode = inquirycode
	err = OSQL.Read(&inquirydetail, "inquirycode")
	if err != nil {
		beego.Error(err)
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
		// beego.Error("%s", err)
		// errorCode = util.Inquirydetail_EDIT_FAILED
		// return errorCode
		code := AddInquirydetail(inquirydetail)
		if code != util.SUCESSFUL {
			return errorCode
		}
		return code
	}

	//delete
	code := DeleteInquirydetail(temp.Inquirycode)
	if code != util.SUCESSFUL {
		return errorCode
	}

	args := edit_Inquirydetail(inquirydetail)

	num, err2 := OSQL.Update(&inquirydetail, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Inquirydetail_EDIT_FAILED
		return errorCode
	}

	beego.Info("num= err=", num, err2)

	return errorCode
}

func edit_Inquirydetail(param Inquirydetail) (args []string) {
	if param.Mattercode != "" {
		args = append(args, "mattercode")
	}
	if param.Num != 0 {
		args = append(args, "num")
	}
	if param.Price != 0 {
		args = append(args, "price")
	}
	if param.Validity != 0 {
		args = append(args, "validity")
	}
	return args
}

func AddInquirydetail(inquirydetail Inquirydetail) (errorCode int64) {
	var (
		temp Inquirydetail
	)
	temp.Inquirycode = inquirydetail.Inquirycode
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "inquirycode")
	if err == nil {
		beego.Info("inquirydetail have asixt")
		errorCode = util.Inquirydetail_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&inquirydetail)
	if err2 != nil {
		beego.Error(err2)
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
		beego.Error(err)
		errorCode = util.Inquirydetail_DELETE_FAILED
	}

	return errorCode
}
