package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `reviewresult`;
// CREATE TABLE `reviewresult` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '类型(1:采购合同审核)',
//   `type` tinyint(5) DEFAULT NULL,
//   `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联单号',
//   `reviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '审核人工号',
//   `opinion` text CHARACTER SET utf8mb4 COMMENT '审核意见',
//   `result` tinyint(5) NOT NULL COMMENT '1:审核通过;2:驳回',
//   `reviewtime` datetime DEFAULT NULL COMMENT '审核时间',
//   PRIMARY KEY (`id`),
//   UNIQUE KEY `type` (`type`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='审核结果表';

//审核结果表
type Reviewresult struct {
	Id          int64  `json:"id" orm:"column(id)"`                   //id
	Type        int8   `json:"type" orm:"column(type)"`               //类型(1:采购合同审核)
	Relatedcode string `json:"relatedcode" orm:"column(relatedcode)"` //关联单号
	Reviewer    string `json:"reviewer" orm:"column(reviewer)"`       //审核人工号
	Opinion     string `json:"opinion" orm:"column(opinion)"`         //审核意见
	Result      int8   `json:"result" orm:"column(result)"`           //1:审核通过;2:驳回
	Reviewtime  string `json:"reviewtime" orm:"column(reviewtime)"`   //审核时间
}

type ReviewresultParam struct {
	Type        int8
	Relatedcode string
}

func GetReviewresultBypage(param ReviewresultParam) []Reviewresult {
	var (
		rets []Reviewresult
	)
	_, err := OSQL.Raw("select * from "+util.Reviewresult_TABLE_NAME+
		" where type=? and relatedcode='?' order by id asc", param.Type, param.Relatedcode).QueryRows(&rets)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return rets
}

func GetReviewresultById(id int64) (result Reviewresult, err error) {
	result.Id = id
	err = OSQL.Read(&result, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return result, nil
}

func EditReviewresultById(param Reviewresult) (errorCode int64) {
	var (
		temp Reviewresult
	)
	temp.Id = param.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Reviewresult_EDIT_FAILED
		return errorCode
	}

	args := edit_reviewresult(param)

	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Reviewresult_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func edit_reviewresult(param Reviewresult) (args []string) {
	if param.Opinion != "" {
		args = append(args, "opinion")
	}

	if param.Relatedcode != "" {
		args = append(args, "relatedcode")
	}

	if param.Result != 0 {
		args = append(args, "result")
	}

	if param.Reviewer != "" {
		args = append(args, "reviewer")
	}

	if param.Reviewtime != "" {
		args = append(args, "reviewtime")
	}

	if param.Type != 0 {
		args = append(args, "type")
	}
	return args
}

func AddReviewresult(param Reviewresult) (errorCode int64) {
	var (
		temp Reviewresult
	)
	temp.Id = param.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Info("Reviewresult have asixt")
		errorCode = util.Reviewresult_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Reviewresult_ADD_FAILED
	}

	logs.FileLogs.Info("num=%v", id)

	return errorCode
}

func DeleteReviewresult(id int64) (errorCode int64) {
	var (
		temp Reviewresult
	)
	errorCode = util.SUCESSFUL
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Reviewresult_DELETE_FAILED
	}

	logs.FileLogs.Info("num=%v", num)

	return errorCode
}
