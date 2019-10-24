package models

import (
	"strconv"

	"erpweb/util"

	"github.com/astaxie/beego"
)

// DROP TABLE IF EXISTS `review`;
// CREATE TABLE `review` (
//   `type` tinyint(5) DEFAULT NULL COMMENT '类型(1:采购合同审核)',
//   `detail` varchar(1000) CHARACTER SET utf8mb4 NOT NULL COMMENT '详情([{end:0, cardids:['''']}])',
//   PRIMARY KEY (`type`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='审核流程表';

//审核流程表
type Review struct {
	Type   int64  `json:"type" orm:"column(type)"`     //类型(1:采购合同审核)
	Detail string `json:"detail" orm:"column(detail)"` //详情([{end:0, cardids:['''']}],[[],[],[]])
}

func GetReviewBypage(Type int64) []Review {
	var (
		params []Review
	)

	sql := "select * from " + util.Review_TABLE_NAME + " where 1=1 "

	if Type != 0 {
		sql += " and where type=" + strconv.FormatInt(Type, 10)
	}
	_, err := OSQL.Raw(sql).QueryRows(&params)
	if err != nil {
		beego.Error(err)
	}
	return params
}

func GetReviewById(id int64) (ret Review, err error) {
	ret.Type = id
	err = OSQL.Read(&ret, "type")
	if err != nil {
		beego.Error(err)
		return ret, err
	}
	return ret, nil
}

func EditReviewById(param Review) (errorCode int64) {
	var (
		temp Review
	)
	errorCode = util.SUCESSFUL
	temp.Type = param.Type
	err := OSQL.Read(&temp, "type")
	if err != nil {
		// beego.Error("%s", err)
		// errorCode = util.Review_EDIT_FAILED
		// return errorCode
		code := AddReview(param)
		if code != util.SUCESSFUL {
			return code
		}

		return code
	}

	code := DeleteReview(temp.Type)
	if code != util.SUCESSFUL {
		return code
	}

	args := edit_review(param)

	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Review_EDIT_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func edit_review(param Review) (args []string) {
	if param.Detail != "" {
		args = append(args, "detail")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}
	return args
}

func AddReview(param Review) (errorCode int64) {
	var (
		temp Review
	)
	errorCode = util.SUCESSFUL
	temp.Type = param.Type
	err := OSQL.Read(&temp, "type")
	if err == nil {
		beego.Error("Review have this id=", param.Type)
		errorCode = util.Review_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Review_ADD_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}

func DeleteReview(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Review
	)
	temp.Type = id
	num, err := OSQL.Delete(&temp, "type")
	if err != nil {
		beego.Error(err)
		errorCode = util.Review_DELETE_FAILED
		return errorCode
	}
	beego.Info("num=", num)
	return errorCode
}
