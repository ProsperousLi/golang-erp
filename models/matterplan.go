package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// DROP TABLE IF EXISTS `matterplan`;
// CREATE TABLE `matterplan` (
//   `itemid` bigint(20) NOT NULL,
//   `mattercode` varchar(20) COLLATE utf8_bin NOT NULL,
//   `plannum` bigint(20) NOT NULL,
//   PRIMARY KEY (`itemid`,`mattercode`,`plannum`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='备料计划表';

//备料计划表
type Matterplan struct {
	Itemid     int64  `json:"itemid" orm:"column(itemid)"`         //
	Mattercode string `json:"mattercode" orm:"column(mattercode)"` //
	Plannum    int64  `json:"plannum" orm:"column(plannum)"`       //
}

func GetMatterplanBypage(pageNum, pageSize int64) []Matterplan {
	var (
		matterplans []Matterplan
	)
	err := OSQL.Raw("select * from "+util.Matterplan_TABLE_NAME+" order by itemid desc limit ?,?",
		pageNum, pageSize).QueryRow(&matterplans)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return matterplans
}

func GetMatterplanByID(itemid int64) (matterplan Matterplan, err error) {
	matterplan.Itemid = itemid
	err = OSQL.Read(&matterplan, "itemid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return matterplan, nil
}

func EditMatterplanById(matterplan Matterplan) (errorCode int64) {
	var (
		temp Matterplan
	)
	temp.Itemid = matterplan.Itemid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "itemid")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.Matterplan_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&matterplan)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Matterplan_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func AddMatterplan(matterplan Matterplan) (errorCode int64) {
	var (
		temp Matterplan
	)
	temp.Itemid = matterplan.Itemid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "itemid")
	if err == nil {
		logs.FileLogs.Info("matterplan have asixt")
		errorCode = util.Matterplan_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&matterplan)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Matterplan_ADD_FAILED
	}

	return errorCode
}

func DeleteMatterplan(itemid int64) (errorCode int64) {
	var (
		temp Matterplan
	)
	errorCode = util.SUCESSFUL
	temp.Itemid = itemid
	_, err := OSQL.Delete(&temp, "itemid")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Inquirydetail_DELETE_FAILED
	}

	return errorCode
}
