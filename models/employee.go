package models

import (
	//"time"

	"erpweb/logs"
	"erpweb/util"
	//"github.com/astaxie/beego/orm"
)

// DROP TABLE IF EXISTS `employee`;
// CREATE TABLE `employee` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '人员id',
//   `cardid` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '工号',
//   `name` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '姓名',
//   `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0男 1女',
//   `compID` tinyint(4) NOT NULL DEFAULT '0' COMMENT '所属分公司ID  0总部1北京2杭州',
//   `deptID` int(4) NOT NULL COMMENT '部门ID',
//   `dutyID` int(4) DEFAULT NULL COMMENT '岗位ID',
//   `health` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '身体状况',
//   `height` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '身高',
//   `nativeplace` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '籍贯',
//   `nation` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '民族',
//   `maritalstatus` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '婚姻状况',
//   `education` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '学历',
//   `university` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '毕业院校',
//   `major` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '专业',
//   `qualification` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '专业资格',
//   `trialsalary` bigint(20) DEFAULT NULL COMMENT '试用期工资',
//   `salary` bigint(20) DEFAULT NULL COMMENT '转正后工资',
//   `idnumber` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '身份证号',
//   `address1` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '户口地址',
//   `postcode1` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮编',
//   `address2` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '现住地址',
//   `postcode2` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮编',
//   `contactnumber` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '联系电话',
//   `phonenumber` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '手机',
//   `email` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'email',
//   `emergencycontact` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '紧急情况联系人',
//   `contactnumber1` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '联系电话',
//   `address3` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '现在何处',
//   `trialexpired` date NOT NULL COMMENT '试用到期',
//   `entrydate` date NOT NULL COMMENT '入职日期',
//   `birthday` date DEFAULT NULL COMMENT '出生日期',
//   `contractbegindate` date NOT NULL COMMENT '合同开始日期',
//   `contractenddate` date NOT NULL COMMENT '合同结束日期',
//   PRIMARY KEY (`id`),
//   UNIQUE KEY `cardid` (`cardid`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='人员信息表';

type Employee struct {
	Id                int64  `json:"id"`                //人员id
	Cardid            string `json:"cardid"`            //工号
	Name              string `json:"name"`              //姓名
	Sex               int8   `json:"sex"`               //性别 0男 1女
	CompID            int8   `orm:"column(compID)"`     //所属分公司ID  0总部1北京2杭州
	DeptID            int    `orm:"column(deptID)"`     //部门ID
	DutyID            int    `orm:"column(dutyID)"`     //岗位ID
	Health            string `json:"health"`            //身体状况
	Height            string `json:"height"`            //身高
	Nativeplace       string `json:"nativeplace"`       //籍贯
	Nation            string `json:"nation"`            //民族
	Maritalstatus     string `json:"maritalstatus"`     //婚姻状况
	Education         string `json:"education"`         //学历
	University        string `json:"university"`        //毕业院校
	Major             string `json:"major"`             //专业
	Qualification     string `json:"qualification"`     //专业资格
	Trialsalary       int64  `json:"trialsalary"`       //试用期工资
	Salary            int64  `json:"salary"`            //转正后工资
	Idnumber          string `json:"idnumber"`          //身份证号
	Address1          string `json:"address1"`          //户口地址
	Postcode1         string `json:"postcode1"`         //邮编
	Address2          string `json:"address2"`          //现住地址
	Postcode2         string `json:"postcode2"`         //邮编
	Contactnumber     string `json:"contactnumber"`     //联系电话
	Phonenumber       string `json:"phonenumber"`       //手机
	Email             string `json:"email"`             //email
	Emergencycontact  string `json:"emergencycontact"`  //紧急情况联系人
	Contactnumber1    string `json:"contactnumber1"`    //联系电话
	Address3          string `json:"address3"`          //现在何处
	Trialexpired      string `json:"trialexpired"`      //试用到期
	Entrydate         string `json:"entrydate"`         //入职日期
	Birthday          string `json:"birthday"`          //出生日期
	Contractbegindate string `json:"contractbegindate"` //合同开始日期
	Contractenddate   string `json:"contractenddate"`   //合同结束日期
}

func GetEmployees(pageNum, pageSize int64) []Employee {
	var (
		emps []Employee
		//err  error
	)
	begin := pageSize * pageNum
	logs.FileLogs.Info("begin=%v", begin, ", end =%v", pageSize)
	_, err := OSQL.Raw("select * from "+util.EMPLOYEE_TABLE_NAME+" order by id asc limit ?,?",
		begin, pageSize).QueryRows(&emps)
	if err != nil {
		logs.FileLogs.Error("%s", err) //logs.Error(err)
	}

	return emps
}

func GetEmployeeById(id int64) (emp Employee, errorCode int64) {
	errorCode = util.SUCESSFUL
	emp.Id = id
	err := OSQL.Raw("select * from "+util.EMPLOYEE_TABLE_NAME+" where id=?", emp.Id).QueryRow(&emp)
	if err != nil {
		logs.FileLogs.Error("%s", err)
	}

	return emp, errorCode
}

func EditEmployeeById(emp Employee) (errorCode int64) {
	errorCode = util.SUCESSFUL
	var (
		temp Employee
	)
	errorCode = util.SUCESSFUL
	temp.Id = emp.Id
	err := OSQL.Read(&temp, "id")
	if err != nil {
		logs.FileLogs.Error("%s", err)
		errorCode = util.ACCOUNT_EDIT_FAILED
		return errorCode
	}
	args := edit_employee(emp)
	if len(args) > 0 {
		num, err2 := OSQL.Update(&emp, args...)
		if err2 != nil {
			logs.FileLogs.Error("%s", err)
			errorCode = util.ACCOUNT_EDIT_FAILED
			return errorCode
		}
		logs.FileLogs.Info("num=%v", num)
	} else {
		logs.FileLogs.Info("no data update")
	}

	return errorCode
}

func edit_employee(param Employee) []string {
	var (
		args []string
	)

	if param.Cardid != "" {
		args = append(args, "cardid")
	}

	if param.Name != "" {
		args = append(args, "name")
	}

	if param.Sex != 0 {
		args = append(args, "sex")
	}

	if param.CompID != 0 {
		args = append(args, "compID")
	}

	if param.DeptID != 0 {
		args = append(args, "deptID")
	}

	if param.DutyID != 0 {
		args = append(args, "dutyID")
	}

	if param.Health != "" {
		args = append(args, "health")
	}

	if param.Height != "" {
		args = append(args, "height")
	}

	if param.Nativeplace != "" {
		args = append(args, "nativeplace")
	}

	if param.Nation != "" {
		args = append(args, "nation")
	}

	if param.Maritalstatus != "" {
		args = append(args, "maritalstatus")
	}

	if param.Education != "" {
		args = append(args, "education")
	}

	if param.University != "" {
		args = append(args, "university")
	}

	if param.Major != "" {
		args = append(args, "major")
	}

	if param.Qualification != "" {
		args = append(args, "qualification")
	}

	if param.Trialsalary != 0 {
		args = append(args, "trialsalary")
	}
	if param.Salary != 0 {
		args = append(args, "salary")
	}
	if param.Idnumber != "" {
		args = append(args, "idnumber")
	}
	if param.Address1 != "" {
		args = append(args, "address1")
	}

	if param.Postcode1 != "" {
		args = append(args, "postcode1")
	}

	if param.Address2 != "" {
		args = append(args, "address2")
	}

	if param.Postcode2 != "" {
		args = append(args, "postcode2")
	}

	if param.Contactnumber != "" {
		args = append(args, "contactnumber")
	}

	if param.Phonenumber != "" {
		args = append(args, "phonenumber")
	}

	if param.Email != "" {
		args = append(args, "cardid")
	}
	if param.Emergencycontact != "" {
		args = append(args, "emergencycontact")
	}
	if param.Contactnumber1 != "" {
		args = append(args, "contactnumber1")
	}
	if param.Address3 != "" {
		args = append(args, "address3")
	}
	if param.Trialexpired != "" {
		args = append(args, "trialexpired")
	}
	if param.Entrydate != "" {
		args = append(args, "entrydate")
	}
	if param.Birthday != "" {
		args = append(args, "birthday")
	}
	if param.Contractbegindate != "" {
		args = append(args, "contractbegindate")
	}
	if param.Contractenddate != "" {
		args = append(args, "contractenddate")
	}

	logs.FileLogs.Info("args=%v", args)

	return args
}

// func AddEmployee(emp Employee) (errorCode int64) {
// 	var (
// 		trialexpired, entrydate, birthday, contractbegindate, contractenddate time.Time
// 	)
// 	trialexpired = util.Str2Time(emp.Trialexpired)
// 	entrydate = util.Str2Time(emp.Entrydate)
// 	birthday = util.Str2Time(emp.Birthday)
// 	contractbegindate = util.Str2Time(emp.Contractbegindate)
// 	contractenddate = util.Str2Time(emp.Contractenddate)
// 	if emp.Cardid == "" || emp.Name == "" || emp.DeptID == 0 ||
// 		emp.Idnumber == "" {
// 		logs.FileLogs.Error("must need params have null %v", emp)
// 		return util.PARAM_FAILED
// 	}

// 	//INSERT INTO `erpweb`.`employee` (`id`, `cardid`, `name`, `sex`, `compID`, `deptID`, `dutyID`, `health`, `height`, `nativeplace`, `nation`, `maritalstatus`, `education`, `university`, `major`, `qualification`, `trialsalary`, `salary`, `idnumber`, `address1`, `postcode1`, `address2`, `postcode2`, `contactnumber`, `phonenumber`, `email`, `emergencycontact`, `contactnumber1`, `address3`, `trialexpired`, `entrydate`, `birthday`, `contractbegindate`, `contractenddate`) VALUES (NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
// 	ret, err := OSQL.Raw("INSERT INTO `erp`.`employee` (`cardid`, `name`,"+
// 		"`sex`, `compID`, `deptID`, `dutyID`, `health`,"+
// 		"`height`, `nativeplace`, `nation`, `maritalstatus`,"+
// 		"`education`, `university`, `major`, `qualification`,"+
// 		"`trialsalary`, `salary`, `idnumber`, `address1`,"+
// 		"`postcode1`, `address2`, `postcode2`, `contactnumber`,"+
// 		"`phonenumber`, `email`, `emergencycontact`, `contactnumber1`,"+
// 		"`address3`, `trialexpired`, `entrydate`, `birthday`,"+
// 		"`contractbegindate`, `contractenddate`)"+
// 		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,"+
// 		"?, ?, ?, ?, ?, ?, ?, ?, ?)", emp.Cardid, emp.Name, emp.Sex, emp.CompID,
// 		emp.DeptID, emp.DutyID, emp.Health, emp.Height, emp.Nativeplace, emp.Nation,
// 		emp.Maritalstatus, emp.Education, emp.University, emp.Major, emp.Qualification,
// 		emp.Trialsalary, emp.Salary, emp.Idnumber, emp.Address1, emp.Postcode1, emp.Address2,
// 		emp.Postcode2, emp.Contactnumber, emp.Phonenumber, emp.Email, emp.Emergencycontact,
// 		emp.Contactnumber1, emp.Address3, trialexpired, entrydate, birthday, contractbegindate,
// 		contractenddate).Exec()
// 	if err != nil {
// 		logs.FileLogs.Error("%s", err)
// 		return util.PARAM_FAILED
// 	}
// 	affect, _ := ret.RowsAffected()

// 	logs.FileLogs.Info("%v", affect)

// 	return util.SUCESSFUL
// }

func AddEmployee(param Employee) (errorCode int64) {
	var (
		temp Employee
	)
	temp.Id = param.Id
	errorCode = util.SUCESSFUL
	err := OSQL.Read(&temp, "id")
	if err == nil {
		logs.FileLogs.Error("table have this id=%v", param.Id)
		errorCode = util.EMPLOYEE_ADD_FAILED
		return errorCode
	}

	id, err2 := OSQL.Insert(&param)
	if err2 != nil {
		logs.FileLogs.Error("%v", err2)
		errorCode = util.EMPLOYEE_ADD_FAILED
	}

	logs.FileLogs.Info("num=%v", id)
	return errorCode
}

func DeleteEmployee(id int64) (errorCode int64) {
	var (
		param Employee
	)
	errorCode = util.SUCESSFUL
	param.Id = id
	num, err := OSQL.Delete(&param, "id")
	if err != nil {
		logs.FileLogs.Error("%v", err)
		errorCode = util.EMPLOYEE_DELETE_FAILED
	}
	logs.FileLogs.Info("num=%v", num)
	return errorCode
}