package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `review`;
// CREATE TABLE `review` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `type` tinyint(5) DEFAULT NULL COMMENT '类型(1:采购合同审核)',
//   `detail` varchar(1000) CHARACTER SET utf8mb4 NOT NULL COMMENT '详情([{end:0, cardids:['''']}])',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='审核流程表';

//审核流程表
type Review struct {
	Id     int64  `json:"id" orm:"column(id)"`
	Type   string `json:"type" orm:"column(type)"`     //类型(1:采购合同审核)
	Detail int8   `json:"detail" orm:"column(detail)"` //详情([{end:0, cardids:['''']}])
}

func GetReviewBypage(pageNum, pageSize int64) []Review {
	var (
		params []Review
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Review_TABLE_NAME+" order by id desc limit ?,?",
		begin, pageSize).QueryRows(&params)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return params
}

func GetReviewById(id int64) (ret Review, err error) {
	ret.Id = id
	err = OSQL.Read(&ret, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ret, err
	}
	return ret, nil
}

func EditReviewById(param Review) (errorCode int64) {
	var (
		temp Review
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Review_EDIT_FAILED
		return errorCode
	}

	args := edit_review(param)

	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Review_EDIT_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func edit_review(param Review) (args []string) {
	if param.Detail != 0 {
		args = append(args, "detail")
	}

	if param.Type != "" {
		args = append(args, "type")
	}
	return args
}

func AddReview(param Review) (errorCode int64) {
	var (
		temp Review
	)
	errorCode = util.SUCESSFUL
	temp.Id = param.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("Review have this id=%v", param.Id)
		errorCode = util.Review_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Review_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteReview(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Review
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Review_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
