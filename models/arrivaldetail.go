package models

import (
	"erpweb/logs"
	"erpweb/util"
)

// -- ----------------------------
// -- Table structure for arrivaldetail
// -- ----------------------------
// DROP TABLE IF EXISTS `arrivaldetail`;
// CREATE TABLE `arrivaldetail` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT,
//   `arrivalbillcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '到货单编码',
//   `mattercode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
//   `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料名称',
//   `unit` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '单位',
//   `arrivalnum` bigint(20) NOT NULL COMMENT '到货数量',
//   `putinnum` bigint(20) NOT NULL DEFAULT '0' COMMENT '入库数量',
//   `price` bigint(20) NOT NULL COMMENT '单价',
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='到货详情表';

//到货详情表
type Arrivaldetail struct {
	Id              int64  `json:"id" orm:"column(id)"`                           //id
	Arrivalbillcode string `json:"arrivalbillcode" orm:"column(arrivalbillcode)"` //到货单编码
	Mattercode      string `json:"mattercode" orm:"column(mattercode)"`           //物料编码
	Name            string `json:"name" orm:"column(name)"`                       //物料名称
	Unit            string `json:"unit" orm:"column(unit)"`                       //单位
	Arrivalnum      int64  `json:"arrivalnum" orm:"column(arrivalnum)"`           //到货数量
	Putinnum        int64  `json:"putinnum" orm:"column(putinnum)"`               //入库数量
	Price           int64  `json:"price" orm:"column(price)"`                     //单价
}

func GetArrivaldetailBypage(arrivalbillcode string) []Arrivaldetail {
	var (
		arrivalbills []Arrivaldetail
	)
	_, err := OSQL.Raw("select * from " + util.Arrivaldetail_TABLE_NAME +
		" and where arrivalbillcode='" + arrivalbillcode + "' order by id desc").QueryRows(&arrivalbills)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return arrivalbills
}

func GetArrivaldetailByUserID(id int64) (ret Arrivaldetail, err error) {
	ret.Id = id
	err = OSQL.Read(&ret, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}
	return ret, nil
}

func EditArrivaldetailById(param Arrivaldetail) (errorCode int64) {
	var (
		temp Arrivalbill
	)
	temp.Arrivalbillcode = param.Arrivalbillcode
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "arrivalbillcode")
	if err != nil {
		// logs.FileLogs.Error("%s", err)
		// errorCode = util.Arrivaldetail_EDIT_FAILED
		// return errorCode
		//add
		code := AddArrivaldetail(param)
		if code != util.SUCESSFUL {
			return errorCode
		}

		return code
	}

	//delete
	code := DeleteArrivaldetail(temp.Arrivalbillcode)
	if code != util.SUCESSFUL {
		return errorCode
	}

	args := edit_Arrivaldetail(param)
	num, err2 := OSQL.Update(&param, args...)
	if err2 != nil {
		logs.FileLogs.Error("%s", err2)
		errorCode = util.Arrivaldetail_EDIT_FAILED
		return errorCode
	}

	logs.FileLogs.Info("num=%v err=%v", num, err2)

	return errorCode
}

func edit_Arrivaldetail(param Arrivaldetail) (args []string) {
	if param.Arrivalbillcode != "" {
		args = append(args, "arrivalbillcode")
	}

	if param.Arrivalnum != 0 {
		args = append(args, "arrivalnum")
	}

	if param.Mattercode != "" {
		args = append(args, "mattercode")
	}

	if param.Name != "" {
		args = append(args, "name")
	}

	if param.Price != 0 {
		args = append(args, "price")
	}

	if param.Putinnum != 0 {
		args = append(args, "putinnum")
	}

	if param.Unit != "" {
		args = append(args, "unit")
	}

	return args
}

func AddArrivaldetail(param Arrivaldetail) (errorCode int64) {
	var (
		temp Arrivaldetail
	)
	temp.Id = param.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "arrivalbillcode")
	if err == nil {
		logs.FileLogs.Info("Arrivaldetail have asixt")
		errorCode = util.Arrivaldetail_ADD_FAILED
		return errorCode
	}

	_, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.Arrivaldetail_ADD_FAILED
	}

	return errorCode
}

func DeleteArrivaldetail(arrivalbillcode string) (errorCode int64) {
	var (
		temp Arrivalbill
	)
	errorCode = util.SUCESSFUL
	temp.Arrivalbillcode = arrivalbillcode
	_, err := OSQL.Delete(&temp, "arrivalbillcode")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.Arrivalbill_DELETE_FAILED
	}

	return errorCode
}
