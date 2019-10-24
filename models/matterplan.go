package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
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

type WebMatterplanAndMatter struct {
	Itemid     int64  `json:"itemid"`
	Mattercode string `json:"mattercode"`
	Plannum    int64  `json:"plannum"`
	Name       string `json:"name" orm:"column(name)"`   //物料名称
	Param      string `json:"param" orm:"column(param)"` //规格参数
	Unit       string `json:"unit" orm:"column(unit)"`   //单位
}

func GetMatterplansByItemid(itemid int64) (rets []WebMatterplanAndMatter) {
	mapls, _ := GetMatterplanById(itemid)
	if len(mapls) > 0 {
		for _, mp := range mapls {
			var temp WebMatterplanAndMatter
			matter, err := GetMatterByMattercode(mp.Mattercode)
			if err != nil {
				continue
			}
			temp.Itemid = mp.Itemid
			temp.Mattercode = mp.Mattercode
			temp.Plannum = mp.Plannum
			temp.Name = matter.Name
			temp.Param = matter.Param
			temp.Unit = matter.Unit

			rets = append(rets, temp)
		}
	}

	return
}

func GetMatterplanBypage(pageNum, pageSize int64) []Matterplan {
	var (
		matterplans []Matterplan
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.Matterplan_TABLE_NAME+" order by itemid desc limit ?,?",
		begin, pageSize).QueryRows(&matterplans)
	if err != nil {
		beego.Error(err)
	}
	return matterplans
}

func GetMatterplanById(itemid int64) (matterplans []Matterplan, err error) {
	_, err = OSQL.Raw("select * from "+util.Matterplan_TABLE_NAME+
		" where itemid=? order by itemid desc", itemid).QueryRows(&matterplans)
	if err != nil {
		beego.Error(err)
		return matterplans, err
	}
	return matterplans, nil
}

func EditMatterplanById(matterplan Matterplan) (errorCode int64) {
	var (
		temp Matterplan
	)
	temp.Itemid = matterplan.Itemid
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "itemid")
	if err != nil {
		beego.Error(err)
		errorCode = util.Matterplan_EDIT_FAILED
		return errorCode
	}

	num, err2 := OSQL.Update(&matterplan)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Matterplan_EDIT_FAILED
		return errorCode
	}

	beego.Info("num= err=", num, err2)

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
		beego.Info("matterplan have asixt")
		errorCode = util.Matterplan_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&matterplan)
	if err2 != nil {
		beego.Error(err2)
		errorCode = util.Matterplan_ADD_FAILED
	}

	return errorCode
}

type DeleteMatterStruct struct {
	Itemid     int64
	Mattercode string
}

func DeleteMatterplan(param DeleteMatterStruct) (errorCode int64) {
	var (
		temp Matterplan
	)
	errorCode = util.SUCESSFUL
	temp.Itemid = param.Itemid
	temp.Mattercode = param.Mattercode
	_, err := OSQL.Delete(&temp, "itemid", "mattercode")
	if err != nil {
		beego.Error(err)
		errorCode = util.Matterplan_DELETE_FAILED
	}

	return errorCode
}
