package models

import (
	"erpweb/logs"
	"erpweb/util"
)

//客户信息表
type Matter struct {
	Id          int64  `json:"id" orm:"column(id)"`
	Mattercode  string `json:"mattercode" orm:"column(mattercode)"`   //物料编码
	Name        string `json:"name" orm:"column(name)"`               //物料名称
	Brand       string `json:"brand" orm:"column(brand)"`             //品牌
	Unit        string `json:"unit" orm:"column(unit)"`               //单位
	Class       int64  `json:"class" orm:"column(class)"`             //物料分类id
	Param       string `json:"param" orm:"column(param)"`             //规格参数
	Grossweight string `json:"grossweight" orm:"column(grossweight)"` //毛重
	Netweight   string `json:"netweight" orm:"column(netweight)"`     //净重
}

func QueryMatter(mattercode, name string) []Matter {
	var (
		mas []Matter
		sql string
	)
	if mattercode != "" && name != "" {
		sql = "select * from " + util.MATTER_TABLE_NAME + " where mattercode like %?% " +
			"or name like %?% order by id asc"
	} else if mattercode != "" && name == "" {
		sql = "select * from " + util.MATTER_TABLE_NAME + " where mattercode like %?% " +
			"order by id asc"
	} else if mattercode == "" && name != "" {
		sql = "select * from " + util.MATTER_TABLE_NAME + " where name like %?% order by id asc"
	} else {
		sql = "select * from " + util.MATTER_TABLE_NAME + " order by id asc"
	}

	_, err := OSQL.Raw(sql).QueryRows(&mas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return mas
}

func GetMatterBypage(pageNum, pageSize int64) []Matter {
	var (
		mas []Matter
	)
	begin := pageSize * pageNum
	_, err := OSQL.Raw("select * from "+util.MATTER_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&mas)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return mas
}

func GetMatterById(id int64) (ma Matter, err error) {
	ma.Id = id
	err = OSQL.Read(&ma, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		return ma, err
	}
	return ma, nil
}

func EditMatterById(ma Matter) (errorCode int64) {
	var (
		temp Matter
	)
	errorCode = util.SUCESSFUL
	temp.Id = ma.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.MATTER_EDIT_FAILED
		return errorCode
	}

	args := edit_matter(ma)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&ma, args...)
		if err2 != nil {
			logs.FileLogs.Error("%s", err2)
			errorCode = util.MATTER_EDIT_FAILED
			return errorCode
		}
		logs.FileLogs.Info("num=%v", num)
	} else {
		logs.FileLogs.Info("no data update")
	}

	return errorCode
}

func edit_matter(ma Matter) []string {
	var (
		args []string
	)
	logs.FileLogs.Info("%v", ma)
	if ma.Brand != "" {
		args = append(args, "brand")
	}
	if ma.Class != 0 {
		args = append(args, "class")
	}
	if ma.Grossweight != "" {
		args = append(args, "grossweight")
	}
	if ma.Mattercode != "" {
		args = append(args, "mattercode")
	}
	if ma.Name != "" {
		args = append(args, "name")
	}
	if ma.Netweight != "" {
		args = append(args, "netweight")
	}
	if ma.Param != "" {
		args = append(args, "param")
	}
	if ma.Unit != "" {
		args = append(args, "unit")
	}

	logs.FileLogs.Info("%v", args)

	return args
}

func AddMatter(ma Matter) (errorCode int64) {
	var (
		temp Matter
	)
	errorCode = util.SUCESSFUL
	temp.Id = ma.Id
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("ware have this id=%v", ma.Id)
		errorCode = util.MATTER_ADD_FAILED
		return errorCode
	}

	num, err2 := OSQL.Insert(&ma)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.MATTER_ADD_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}

func DeleteMatter(id int64) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Matter
	)
	temp.Id = id
	num, err := OSQL.Delete(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.MATTER_DELETE_FAILED
		return errorCode
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}
