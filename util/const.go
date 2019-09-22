package util

import ()

const (
	DEFUAL_PWD_PRE = "bcxskyc@"
	DEFUAL_PWD     = "bxkc!QAZ888"
)

const (
	EMPLOYEE_TABLE_NAME         = "employee"
	ACCOUNT_TABLE_NAME          = "account"
	DEPARTMENT_TABLE_NAME       = "department"
	DUTY_TABLE_NAME             = "duty"
	MENU_TABLE_NAME             = "menu"
	OPERLOG_TABLE_NAME          = "operlog"
	PERMISSION_TABLE_NAME       = "permission"
	WAREHOUSE_TABLE_NAME        = "warehouse"
	CUSTOMER_TABLE_NAME         = "customer"
	LEAVE_TABLE_NAME            = "leaves"
	MATTER_TABLE_NAME           = "matter"
	MATTERPACKAGE_TABLE_NAME    = "matterpackage"
	PACKAGERALATION_TABLE_NAME  = "packagerelation"
	SUPPLIER_TABLE_NAME         = "supplier"
	SUPPLYRELATION_TABLE_NAME   = "supplyrelation"
	VEHICLE_TABLE_NAME          = "vehicle"
	Stock_TABLE_NAME            = "stock"
	Saledetail_TABLE_NAME       = "saledetail"
	Marketcontract_TABLE_NAME   = "marketcontract"
	Reviewresult_TABLE_NAME     = "reviewresult"
	Arrivalbill_TABLE_NAME      = "arrivalbill"
	Arrivaldetail_TABLE_NAME    = "arrivaldetail"
	Financeflow_TABLE_NAME      = "financeflow"
	Inquiry_TABLE_NAME          = "inquiry"
	Inquirydetail_TABLE_NAME    = "inquirydetail"
	Matterplan_TABLE_NAME       = "matterplan"
	Outofdetail_TABLE_NAME      = "outofdetail"
	Outofstore_TABLE_NAME       = "outofstore"
	Purchasecontract_TABLE_NAME = "purchasecontract"
	Purchasedetail_TABLE_NAME   = "purchasedetail"
	Putindetail_TABLE_NAME      = "putindetail"
	Putinstore_TABLE_NAME       = "putinstore"
	Repaircontract_TABLE_NAME   = "repaircontract"
	Repaircost_TABLE_NAME       = "repaircost"
	Repairitem_TABLE_NAME       = "repairitem"
	Review_TABLE_NAME           = "review"

	SUCESSFUL = 20000
	//20000 begin
	FAILED                 = 20001
	EMPLOYEE_ADD_FAILED    = 20002
	EMPLOYEE_DELETE_FAILED = 30001

	ACCOUNT_EDIT_FAILED   = 20003
	ACCOUNT_ADD_FAILED    = 20004
	ACCOUNT_DELETE_FAILED = 20005

	DEPART_EDIT_FAILED   = 20006
	DEPART_ADD_FAILED    = 20007
	DEPART_DELETE_FAILED = 20008

	DUTY_EDIT_FAILED   = 20009
	DUTY_ADD_FAILED    = 20010
	DUTY_DELETE_FAILED = 20011

	OPERLOG_EDIT_FAILED   = 20012
	OPERLOG_ADD_FAILED    = 20013
	OPERLOG_DELETE_FAILED = 20014

	PERMISSION_EDIT_FAILED   = 20015
	PERMISSION_ADD_FAILED    = 20016
	PERMISSION_DELETE_FAILED = 20017

	WAREHOUSE_EDIT_FAILED   = 20018
	WAREHOUSE_ADD_FAILED    = 20019
	WAREHOUSE_DELETE_FAILED = 20020

	CUSTOMER_EDIT_FAILED   = 20021
	CUSTOMER_ADD_FAILED    = 20022
	CUSTOMER_DELETE_FAILED = 20023

	LEAVE_EDIT_FAILED   = 20024
	LEAVE_ADD_FAILED    = 20025
	LEAVE_DELETE_FAILED = 20026

	MATTER_EDIT_FAILED   = 20027
	MATTER_ADD_FAILED    = 20028
	MATTER_DELETE_FAILED = 20029

	MATTERPACKAGE_EDIT_FAILED   = 20030
	MATTERPACKAGE_ADD_FAILED    = 20031
	MATTERPACKAGE_DELETE_FAILED = 20032

	PACKAGERALATION_EDIT_FAILED   = 20033
	PACKAGERALATION_ADD_FAILED    = 20034
	PACKAGERALATION_DELETE_FAILED = 20035

	//PACKAGERALATION_EDIT_FAILED   = 20036
	//PACKAGERALATION_ADD_FAILED    = 20037
	//PACKAGERALATION_DELETE_FAILED = 20038

	SUPPLIER_EDIT_FAILED   = 20039
	SUPPLIER_ADD_FAILED    = 20040
	SUPPLIER_DELETE_FAILED = 20041

	SUPPLYRELATION_EDIT_FAILED   = 20042
	SUPPLYRELATION_ADD_FAILED    = 20043
	SUPPLYRELATION_DELETE_FAILED = 20044

	VEHICLE_EDIT_FAILED   = 20045
	VEHICLE_ADD_FAILED    = 20046
	VEHICLE_DELETE_FAILED = 20047

	//TODO
	Stock_EDIT_FAILED   = 20045
	Stock_ADD_FAILED    = 20046
	Stock_DELETE_FAILED = 20047

	Saledetail_EDIT_FAILED   = 20045
	Saledetail_ADD_FAILED    = 20046
	Saledetail_DELETE_FAILED = 20047

	Marketcontract_EDIT_FAILED   = 20045
	Marketcontract_ADD_FAILED    = 20046
	Marketcontract_DELETE_FAILED = 20047

	Reviewresult_EDIT_FAILED   = 20045
	Reviewresult_ADD_FAILED    = 20046
	Reviewresult_DELETE_FAILED = 20047

	Arrivalbill_EDIT_FAILED   = 20048
	Arrivalbill_ADD_FAILED    = 20049
	Arrivalbill_DELETE_FAILED = 20050

	Arrivaldetail_EDIT_FAILED   = 20051
	Arrivaldetail_ADD_FAILED    = 20052
	Arrivaldetail_DELETE_FAILED = 20053

	Financeflow_EDIT_FAILED   = 20054
	Financeflow_ADD_FAILED    = 20055
	Financeflow_DELETE_FAILED = 20056

	Inquiry_EDIT_FAILED   = 20057
	Inquiry_ADD_FAILED    = 20058
	Inquiry_DELETE_FAILED = 20059

	Inquirydetail_EDIT_FAILED   = 20060
	Inquirydetail_ADD_FAILED    = 20061
	Inquirydetail_DELETE_FAILED = 20062

	Matterplan_EDIT_FAILED   = 20063
	Matterplan_ADD_FAILED    = 20064
	Matterplan_DELETE_FAILED = 20065

	Outofdetail_EDIT_FAILED   = 20066
	Outofdetail_ADD_FAILED    = 20067
	Outofdetail_DELETE_FAILED = 20068

	Outofstore_EDIT_FAILED   = 20069
	Outofstore_ADD_FAILED    = 20070
	Outofstore_DELETE_FAILED = 20071

	Purchasecontract_EDIT_FAILED   = 20072
	Purchasecontract_ADD_FAILED    = 20073
	Purchasecontract_DELETE_FAILED = 20074

	Purchasedetail_EDIT_FAILED   = 20072
	Purchasedetail_ADD_FAILED    = 20073
	Purchasedetail_DELETE_FAILED = 20074

	Putindetail_EDIT_FAILED   = 20075
	Putindetail_ADD_FAILED    = 20076
	Putindetail_DELETE_FAILED = 20077

	Putinstore_EDIT_FAILED   = 20078
	Putinstore_ADD_FAILED    = 20079
	Putinstore_DELETE_FAILED = 20080

	Repaircontract_EDIT_FAILED   = 20081
	Repaircontract_ADD_FAILED    = 20082
	Repaircontract_DELETE_FAILED = 20083

	Repaircost_EDIT_FAILED   = 20084
	Repaircost_ADD_FAILED    = 20085
	Repaircost_DELETE_FAILED = 20086

	Repairitem_EDIT_FAILED   = 20087
	Repairitem_ADD_FAILED    = 20088
	Repairitem_DELETE_FAILED = 20089

	Review_EDIT_FAILED   = 20090
	Review_ADD_FAILED    = 20091
	Review_DELETE_FAILED = 20092

	//all errorCode
	PARAM_FAILED = 19999
)

var (
	RetContent Result
)

type Result struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
