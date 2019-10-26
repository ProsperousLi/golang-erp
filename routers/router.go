package routers

import (
	"erpweb/controllers"
	"github.com/astaxie/beego"
)

const (
	PROJECT = "api"
	GET     = "get"
	POST    = "post"
)

func init() {
	beego.Router("/", &controllers.MainController{}, GET+":Index")
	Allinterfaces()
	Login()
	Account()
	Arrivalbill()
	Arrivaldetail()
	Customer()
	Department()
	Duty()
	Employee()
	Financeflow()
	Inquiry()
	Inquirydetail()
	Leave()
	Marketcontract()
	Matter()
	Matterplan()
	Menu()
	Operlog()
	Outofdetail()
	Outofstore()
	Permission()
	Purchasecontract()
	Purchasedetail()
	Putindetail()
	Putinstore()
	Repaircontract()
	Repaircost()
	Repairitem()
	Review()
	Reviewresult()
	Saledetail()
	Stock()
	Supplier()
	Supplyrelation()
	Vehicle()
	Warehouse()
}

func Allinterfaces() {
	//获取单据序号
	beego.Router("/"+PROJECT+"/ts/queryTimeStamp",
		&controllers.AllinterfacesController{}, GET+":QueryTimeStamp")

	//93.查询时间戳
	beego.Router("/"+PROJECT+"/review/queryTimeStamp",
		&controllers.AllinterfacesController{}, GET+":QueryTimeStampDays")

	beego.Router("/"+PROJECT+"/basedata/queryMattersOfSupplier",
		&controllers.AllinterfacesController{}, GET+":QueryMattersOfSupplier")

	beego.Router("/"+PROJECT+"/basedata/querySuppliersOfMatter",
		&controllers.AllinterfacesController{}, GET+":QuerySuppliersOfMatter")

	beego.Router("/"+PROJECT+"/basedata/updateMatterListOfSupplier",
		&controllers.AllinterfacesController{}, POST+":UpdateMatterListOfSupplier")

	beego.Router("/"+PROJECT+"/basedata/updateSupplierListOfMatter",
		&controllers.AllinterfacesController{}, POST+":UpdateSupplierListOfMatter")

	beego.Router("/"+PROJECT+"/basedata/queryNumsOfPurchasecontract",
		&controllers.AllinterfacesController{}, POST+":QueryNumsOfPurchasecontract")
}

func Login() {
	beego.Router("/"+PROJECT+"/login/login",
		&controllers.LoginController{}, POST+":Login")

	beego.Router("/"+PROJECT+"/login/logout",
		&controllers.LoginController{}, POST+":Loginout")
	//api/user/info
	beego.Router("/"+PROJECT+"/user/info",
		&controllers.LoginController{}, GET+":UserInfo")

	//刷新验证码
	beego.Router("/"+PROJECT+"/login/refreshVerifyCode",
		&controllers.LoginController{}, POST+":RefreshVerifyCode")
}

func Account() {
	//api/employee/getAccountList
	beego.Router("/"+PROJECT+"/employee/getAccountList",
		&controllers.AccountController{}, GET+":GetAccountList")

	// beego.Router("/"+PROJECT+"/account/getAccounts",
	// 	&controllers.AccountController{}, POST+":GetAccounts")

	// beego.Router("/"+PROJECT+"/account/getAccountById",
	// 	&controllers.AccountController{}, POST+":GetAccountById")

	// beego.Router("/"+PROJECT+"/account/editAccountById",
	// 	&controllers.AccountController{}, POST+":EditAccountById")

	//16.修改账号
	//url: api/employee/updateAccountStatus
	beego.Router("/"+PROJECT+"/employee/updateAccountStatus",
		&controllers.AccountController{}, POST+":EditAccountStatusById")

	//开通ERP账号
	//api/employee/openAccount
	beego.Router("/"+PROJECT+"/employee/openAccount",
		&controllers.AccountController{}, POST+":AddAccount")

	// beego.Router("/"+PROJECT+"/account/deleteAccount",
	// 	&controllers.AccountController{}, POST+":DeleteAccount")

	//17.修改密码
	//api/employee/modifyPwd
	//{cardid: “xxx”, oldpwd: “xxx”, newpwd: “xxx”}
	beego.Router("/"+PROJECT+"/employee/modifyPwd",
		&controllers.AccountController{}, POST+":ModifyPwd")

	// 	18.重置密码
	// url: api/employee/resetAccount
	//{cardid: “xxx”}
	beego.Router("/"+PROJECT+"/employee/resetAccount",
		&controllers.AccountController{}, POST+":ResetAccount")
}

func Arrivalbill() {
	beego.Router("/"+PROJECT+"/purchase/queryArrivalBill",
		&controllers.ArrivalbillController{}, GET+":QueryArrivalBill")

	beego.Router("/"+PROJECT+"/purchase/getArrivalbills",
		&controllers.ArrivalbillController{}, GET+":GetArrivalbills")

	beego.Router("/"+PROJECT+"/purchase/getArrivalbillById",
		&controllers.ArrivalbillController{}, POST+":GetArrivalbillById")

	beego.Router("/"+PROJECT+"/purchase/updateArrivalBill",
		&controllers.ArrivalbillController{}, POST+":EditArrivalbillById")

	beego.Router("/"+PROJECT+"/purchase/addArrivalBill",
		&controllers.ArrivalbillController{}, POST+":AddArrivalbill")

	beego.Router("/"+PROJECT+"/purchase/delArrivalBill",
		&controllers.ArrivalbillController{}, POST+":DeleteArrivalbill")
}

func Arrivaldetail() {
	preRoute := "purchase"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryArrivalDetail",
		&controllers.ArrivaldetailController{}, GET+":GetArrivaldetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getArrivaldetailById",
		&controllers.ArrivaldetailController{}, POST+":GetArrivaldetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateArrivalDetail",
		&controllers.ArrivaldetailController{}, POST+":EditArrivaldetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addArrivaldetail",
		&controllers.ArrivaldetailController{}, POST+":AddArrivaldetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteArrivaldetail",
		&controllers.ArrivaldetailController{}, POST+":DeleteArrivaldetail")
}

func Customer() {
	beego.Router("/"+PROJECT+"/basedata/queryCustomer",
		&controllers.CustomerController{}, GET+":QueryCustomer")

	beego.Router("/"+PROJECT+"/customer/getCustomers",
		&controllers.CustomerController{}, POST+":GetCustomers")

	beego.Router("/"+PROJECT+"/customer/getCustomerById",
		&controllers.CustomerController{}, POST+":GetCustomerById")

	beego.Router("/"+PROJECT+"/basedata/updateCustomer",
		&controllers.CustomerController{}, POST+":EditCustomerById")

	beego.Router("/"+PROJECT+"/basedata/newCustomer",
		&controllers.CustomerController{}, POST+":AddCustomer")

	beego.Router("/"+PROJECT+"/customer/deleteCustomer",
		&controllers.CustomerController{}, POST+":DeleteCustomer")
}

func Department() {
	beego.Router("/"+PROJECT+"/basedata/queryDept",
		&controllers.DepartmentController{}, GET+":QueryDept")

	beego.Router("/"+PROJECT+"/department/getDepartments",
		&controllers.DepartmentController{}, POST+":GetDepartments")

	beego.Router("/"+PROJECT+"/department/getDepartmentById",
		&controllers.DepartmentController{}, POST+":GetDepartmentById")

	beego.Router("/"+PROJECT+"/basedata/updateDept",
		&controllers.DepartmentController{}, POST+":EditDepartmentById")

	beego.Router("/"+PROJECT+"/basedata/newDept",
		&controllers.DepartmentController{}, POST+":AddDepartment")

	beego.Router("/"+PROJECT+"/department/deleteDepartment",
		&controllers.DepartmentController{}, POST+":DeleteDepartment")
}

func Duty() {
	beego.Router("/"+PROJECT+"/duty/getDutys",
		&controllers.DutyController{}, POST+":GetDutys")

	beego.Router("/"+PROJECT+"/duty/getDutyById",
		&controllers.DutyController{}, POST+":GetDutyById")

	beego.Router("/"+PROJECT+"/duty/editDutyById",
		&controllers.DutyController{}, POST+":EditDutyById")

	beego.Router("/"+PROJECT+"/duty/addDuty",
		&controllers.DutyController{}, POST+":AddDuty")

	beego.Router("/"+PROJECT+"/duty/deleteDuty",
		&controllers.DutyController{}, POST+":DeleteDuty")
}

func Financeflow() {
	preRoute := "finance"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryFinanceFlow",
		&controllers.FinanceflowController{}, GET+":QueryFinanceFlow")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getFinanceflows",
		&controllers.FinanceflowController{}, POST+":GetFinanceflows")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getFinanceflowById",
		&controllers.FinanceflowController{}, POST+":GetFinanceflowById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateFinanceFlow",
		&controllers.FinanceflowController{}, POST+":EditFinanceflowById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addFinanceflow",
		&controllers.FinanceflowController{}, POST+":AddFinanceflow")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delFinanceFlow",
		&controllers.FinanceflowController{}, POST+":DeleteFinanceflow")
}

func Inquiry() {
	preRoute := "sale"

	beego.Router("/"+PROJECT+"/"+preRoute+"/queryInquiry",
		&controllers.InquiryController{}, GET+":QueryInquiry")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquirys",
		&controllers.InquiryController{}, POST+":GetInquirys")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquiryById",
		&controllers.InquiryController{}, POST+":GetInquiryById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateInquiry",
		&controllers.InquiryController{}, POST+":EditInquiryById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addInquiry",
		&controllers.InquiryController{}, POST+":AddInquiry")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delInquiry",
		&controllers.InquiryController{}, POST+":DeleteInquiry")
}

func Inquirydetail() {
	preRoute := "sale"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryInquiryDetail",
		&controllers.InquirydetailController{}, GET+":GetInquirydetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquirydetailById",
		&controllers.InquirydetailController{}, POST+":GetInquirydetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateInquiryDetail",
		&controllers.InquirydetailController{}, POST+":EditInquirydetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addInquirydetail",
		&controllers.InquirydetailController{}, POST+":AddInquirydetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteInquirydetail",
		&controllers.InquirydetailController{}, POST+":DeleteInquirydetail")
}

//人员信息 employee
func Employee() {
	//13.获取员工列表(不分页)
	//api/employee/getEmployeeList
	beego.Router("/"+PROJECT+"/employee/getEmployeeList",
		&controllers.EmployeeController{}, GET+":GetAllEmployees")

	//get persons by page
	// /erp/employee/getPersons
	// beego.Router("/"+PROJECT+"/employee/getPersons",
	// 	&controllers.EmployeeController{}, POST+":GetEmployees")

	// //get person by id
	// // /erp/employee/getPerson
	// beego.Router("/"+PROJECT+"/employee/getPerson",
	// 	&controllers.EmployeeController{}, POST+":GetEmployeeById")

	//edit person info by id
	// /erp/employee/updateEmployee
	beego.Router("/"+PROJECT+"/employee/updateEmployee",
		&controllers.EmployeeController{}, POST+":EditEmployeeById")

	//add person by id
	// /erp/employee/newEmployee
	beego.Router("/"+PROJECT+"/employee/newEmployee",
		&controllers.EmployeeController{}, POST+":AddEmployee")

	// beego.Router("/"+PROJECT+"/employee/deletePerson",
	// 	&controllers.EmployeeController{}, POST+":DeleteEmployee")
}

func Leave() {
	beego.Router("/"+PROJECT+"/employee/queryAllLeave",
		&controllers.LeaveController{}, GET+":QueryAllLeave")

	beego.Router("/"+PROJECT+"/leave/getLeaves",
		&controllers.LeaveController{}, POST+":GetLeaves")

	beego.Router("/"+PROJECT+"/leave/getLeaveById",
		&controllers.LeaveController{}, POST+":GetLeaveById")

	beego.Router("/"+PROJECT+"/leave/editLeaveById",
		&controllers.LeaveController{}, POST+":EditLeaveById")

	beego.Router("/"+PROJECT+"/employee/newLeave",
		&controllers.LeaveController{}, POST+":AddLeave")

	beego.Router("/"+PROJECT+"/leave/deleteLeave",
		&controllers.LeaveController{}, POST+":DeleteLeave")
}

func Marketcontract() {
	beego.Router("/"+PROJECT+"/contract/queryMarketContract",
		&controllers.MarketcontractController{}, GET+":GetMarketcontracts")

	beego.Router("/"+PROJECT+"/marketcontract/getMarketcontractById",
		&controllers.MarketcontractController{}, POST+":GetMarketcontractById")

	beego.Router("/"+PROJECT+"/contract/updateMarketContract",
		&controllers.MarketcontractController{}, POST+":EditMarketcontractById")

	beego.Router("/"+PROJECT+"/contract/addMarketContract",
		&controllers.MarketcontractController{}, POST+":AddMarketcontract")

	beego.Router("/"+PROJECT+"/contract/delMarketContract",
		&controllers.MarketcontractController{}, POST+":DeleteMarketcontract")
}

func Matter() {

	beego.Router("/"+PROJECT+"/basedata/queryMatter",
		&controllers.MatterController{}, GET+":QueryMatter")

	beego.Router("/"+PROJECT+"/matter/getMatters",
		&controllers.MatterController{}, POST+":GetMatters")

	beego.Router("/"+PROJECT+"/basedata/updateMatter",
		&controllers.MatterController{}, POST+":EditMatterById")

	beego.Router("/"+PROJECT+"/basedata/newMatter",
		&controllers.MatterController{}, POST+":AddMatter")

	beego.Router("/"+PROJECT+"/matter/deleteMatter",
		&controllers.MatterController{}, POST+":DeleteMatter")
}

func Matterplan() {
	preRoute := "repair"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryMatterPlan",
		&controllers.MatterplanController{}, POST+":GetMatterplans")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getMatterplanById",
		&controllers.MatterplanController{}, POST+":GetMatterplanById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateMatterPlan",
		&controllers.MatterplanController{}, POST+":EditMatterplanById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addMatterPlan",
		&controllers.MatterplanController{}, POST+":AddMatterplan")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delMatterPlan",
		&controllers.MatterplanController{}, POST+":DeleteMatterplan")
}

func Menu() {
	//api/permission/allmenu
	beego.Router("/"+PROJECT+"/permission/allmenu",
		&controllers.MenuController{}, GET+":GetMenus")
}

func Operlog() {
	beego.Router("/"+PROJECT+"/operlog/getOperlogs",
		&controllers.OperlogController{}, POST+":GetOperlogs")

	beego.Router("/"+PROJECT+"/operlog/getOperlogById",
		&controllers.OperlogController{}, POST+":GetOperlogById")

	beego.Router("/"+PROJECT+"/operlog/editOperlogById",
		&controllers.OperlogController{}, POST+":EditOperlogById")

	beego.Router("/"+PROJECT+"/operlog/addOperlog",
		&controllers.OperlogController{}, POST+":AddOperlog")

	beego.Router("/"+PROJECT+"/operlog/deleteOperlog",
		&controllers.OperlogController{}, POST+":DeleteOperlog")
}

func Outofdetail() {
	preRoute := "store"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryOutofDetail",
		&controllers.OutofdetailController{}, GET+":GetOutofdetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofdetailById",
		&controllers.OutofdetailController{}, POST+":GetOutofdetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editOutofdetailById",
		&controllers.OutofdetailController{}, POST+":EditOutofdetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addOutnDetail",
		&controllers.OutofdetailController{}, POST+":AddOutofdetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteOutofdetail",
		&controllers.OutofdetailController{}, POST+":DeleteOutofdetail")
}

func Outofstore() {
	preRoute := "store"

	beego.Router("/"+PROJECT+"/"+preRoute+"/queryOutofStore",
		&controllers.OutofstoreController{}, GET+":QueryOutofStore")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofstores",
		&controllers.OutofstoreController{}, POST+":GetOutofstores")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofstoreById",
		&controllers.OutofstoreController{}, POST+":GetOutofstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editOutofstoreById",
		&controllers.OutofstoreController{}, POST+":EditOutofstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addOutofStore",
		&controllers.OutofstoreController{}, POST+":AddOutofstore")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteOutofstore",
		&controllers.OutofstoreController{}, POST+":DeleteOutofstore")
}

func Permission() {
	beego.Router("/"+PROJECT+"/permission/getPermissions",
		&controllers.PermissionController{}, POST+":GetPermissions")

	beego.Router("/"+PROJECT+"/permission/getPermissionById",
		&controllers.PermissionController{}, POST+":GetPermissionById")

	beego.Router("/"+PROJECT+"/permission/editPermissionById",
		&controllers.PermissionController{}, POST+":EditPermissionById")

	beego.Router("/"+PROJECT+"/permission/addPermission",
		&controllers.PermissionController{}, POST+":AddPermission")

	beego.Router("/"+PROJECT+"/permission/deletePermission",
		&controllers.PermissionController{}, POST+":DeletePermission")

	//19.设置账号权限
	beego.Router("/"+PROJECT+"/permission/setPermission",
		&controllers.PermissionController{}, POST+":SetPermission")

	//20.获取账号权限
	beego.Router("/"+PROJECT+"/permission/queryPermission",
		&controllers.PermissionController{}, GET+":QueryPermission")
}

func Purchasecontract() {
	preRoute := "contract"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryPurchaseContract",
		&controllers.PurchasecontractController{}, GET+":QueryPurchaseContract")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasecontracts",
		&controllers.PurchasecontractController{}, POST+":GetPurchasecontracts")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasecontractById",
		&controllers.PurchasecontractController{}, POST+":GetPurchasecontractById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updatePurchaseContract",
		&controllers.PurchasecontractController{}, POST+":EditPurchasecontractById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPurchasecontract",
		&controllers.PurchasecontractController{}, POST+":AddPurchasecontract")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delPurchaseContract",
		&controllers.PurchasecontractController{}, POST+":DeletePurchasecontract")
}

func Purchasedetail() {
	preRoute := "contract"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryPurchaseDetail",
		&controllers.PurchasedetailController{}, GET+":GetPurchasedetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasedetailById",
		&controllers.PurchasedetailController{}, POST+":GetPurchasedetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updatePurchaseDetail",
		&controllers.PurchasedetailController{}, POST+":EditPurchasedetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPurchasedetail",
		&controllers.PurchasedetailController{}, POST+":AddPurchasedetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deletePurchasedetail",
		&controllers.PurchasedetailController{}, POST+":DeletePurchasedetail")
}

func Putindetail() {
	preRoute := "store"

	beego.Router("/"+PROJECT+"/"+preRoute+"/queryPutinDetail",
		&controllers.PutindetailController{}, GET+":QueryPutinDetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutindetails",
		&controllers.PutindetailController{}, POST+":GetPutindetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutindetailById",
		&controllers.PutindetailController{}, POST+":GetPutindetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editPutindetailById",
		&controllers.PutindetailController{}, POST+":EditPutindetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPutinDetail",
		&controllers.PutindetailController{}, POST+":AddPutindetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deletePutindetail",
		&controllers.PutindetailController{}, POST+":DeletePutindetail")
}
func Putinstore() {
	preRoute := "store"
	beego.Router("/"+PROJECT+"/"+preRoute+"/putinStore",
		&controllers.PutinstoreController{}, POST+":PutinStore")

	beego.Router("/"+PROJECT+"/"+preRoute+"/queryPutinStore",
		&controllers.PutinstoreController{}, GET+":GetPutinstores")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutinstoreById",
		&controllers.PutinstoreController{}, POST+":GetPutinstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editPutinstoreById",
		&controllers.PutinstoreController{}, POST+":EditPutinstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPutinstore",
		&controllers.PutinstoreController{}, POST+":AddPutinstore")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delPutinStore",
		&controllers.PutinstoreController{}, POST+":DeletePutinstore")
}
func Repaircontract() {
	preRoute := "repaircontract"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepaircontracts",
		&controllers.RepaircontractController{}, POST+":GetRepaircontracts")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepaircontractById",
		&controllers.RepaircontractController{}, POST+":GetRepaircontractById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editRepaircontractById",
		&controllers.RepaircontractController{}, POST+":EditRepaircontractById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addRepaircontract",
		&controllers.RepaircontractController{}, POST+":AddRepaircontract")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteRepaircontract",
		&controllers.RepaircontractController{}, POST+":DeleteRepaircontract")
}
func Repaircost() {
	preRoute := "repair"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryRepairCost",
		&controllers.RepaircostController{}, POST+":QueryRepairCost")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepaircostById",
		&controllers.RepaircostController{}, POST+":GetRepaircostById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateRepairCost",
		&controllers.RepaircostController{}, POST+":EditRepaircostById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addRepairCost",
		&controllers.RepaircostController{}, POST+":AddRepaircost")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delRepairCost",
		&controllers.RepaircostController{}, POST+":DeleteRepaircost")
}
func Repairitem() {
	preRoute := "repair"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryRepairItem",
		&controllers.RepairitemController{}, GET+":GetRepairitems")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepairitemById",
		&controllers.RepairitemController{}, POST+":GetRepairitemById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateRepairItem",
		&controllers.RepairitemController{}, POST+":EditRepairitemById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addRepairItem",
		&controllers.RepairitemController{}, POST+":AddRepairitem")

	beego.Router("/"+PROJECT+"/"+preRoute+"/delRepairItem",
		&controllers.RepairitemController{}, POST+":DeleteRepairitem")
}
func Review() {
	preRoute := "review"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryReviewFlow",
		&controllers.ReviewController{}, GET+":GetReviews")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getReviewById",
		&controllers.ReviewController{}, POST+":GetReviewById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/updateReviewFlow",
		&controllers.ReviewController{}, POST+":EditReviewById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addReview",
		&controllers.ReviewController{}, POST+":AddReview")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteReview",
		&controllers.ReviewController{}, POST+":DeleteReview")
}

func Reviewresult() {
	preRoute := "reviewresult"
	beego.Router("/"+PROJECT+"/"+preRoute+"/queryReviewResult",
		&controllers.ReviewresultController{}, GET+":GetReviewresults")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getReviewresultById",
		&controllers.ReviewresultController{}, POST+":GetReviewresultById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editReviewresultById",
		&controllers.ReviewresultController{}, POST+":EditReviewresultById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addReviewResult",
		&controllers.ReviewresultController{}, POST+":AddReviewresult")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteReviewresult",
		&controllers.ReviewresultController{}, POST+":DeleteReviewresult")
}

func Saledetail() {
	preRoute := "contract"
	beego.Router("/"+PROJECT+"/"+preRoute+"/querySaleDetail",
		&controllers.SaledetailController{}, GET+":GetSaledetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getSaledetailById",
		&controllers.SaledetailController{}, POST+":GetSaledetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editSaledetailById",
		&controllers.SaledetailController{}, POST+":EditSaledetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addSaledetail",
		&controllers.SaledetailController{}, POST+":AddSaledetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteSaledetail",
		&controllers.SaledetailController{}, POST+":DeleteSaledetail")

	//45.新增或修改销售详情(先删除，再新增)
	beego.Router("/"+PROJECT+"/"+preRoute+"/addOrUpdateSaleDetail",
		&controllers.SaledetailController{}, POST+":AddOrUpdateSaleDetail")
}

func Stock() {
	preRoute := "stock"

	beego.Router("/"+PROJECT+"/"+preRoute+"/queryStock",
		&controllers.StockController{}, GET+":QueryStock")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getStocks",
		&controllers.StockController{}, POST+":GetStocks")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getStockById",
		&controllers.StockController{}, POST+":GetStockById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editStockById",
		&controllers.StockController{}, POST+":EditStockById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addStock",
		&controllers.StockController{}, POST+":AddStock")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteStock",
		&controllers.StockController{}, POST+":DeleteStock")
}

func Supplier() {
	beego.Router("/"+PROJECT+"/basedata/querySupplier",
		&controllers.SupplierController{}, GET+":QuerySupplier")

	beego.Router("/"+PROJECT+"/supplier/getSuppliers",
		&controllers.SupplierController{}, POST+":GetSuppliers")

	beego.Router("/"+PROJECT+"/supplier/getSupplierById",
		&controllers.SupplierController{}, POST+":GetSupplierById")

	beego.Router("/"+PROJECT+"/basedata/updateSupplier",
		&controllers.SupplierController{}, POST+":EditSupplierById")

	beego.Router("/"+PROJECT+"/basedata/newSupplier",
		&controllers.SupplierController{}, POST+":AddSupplier")

	beego.Router("/"+PROJECT+"/supplier/deleteSupplier",
		&controllers.SupplierController{}, POST+":DeleteSupplier")
}

func Supplyrelation() {
	preRoute := "supplyrelation"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getSupplyrelations",
		&controllers.SupplyrelationController{}, POST+":GetSupplyrelations")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getSupplyrelationById",
		&controllers.SupplyrelationController{}, POST+":GetSupplyrelationById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editSupplyrelationById",
		&controllers.SupplyrelationController{}, POST+":EditSupplyrelationById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addSupplyrelation",
		&controllers.SupplyrelationController{}, POST+":AddSupplyrelation")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteSupplyrelation",
		&controllers.SupplyrelationController{}, POST+":DeleteSupplyrelation")
}

func Vehicle() {
	beego.Router("/"+PROJECT+"/basedata/queryVehicle",
		&controllers.VehicleController{}, GET+":QueryVehicle")

	// beego.Router("/"+PROJECT+"/vehicle/getVehicleById",
	// 	&controllers.VehicleController{}, POST+":GetVehicleById")

	beego.Router("/"+PROJECT+"/basedata/updateVehicle",
		&controllers.VehicleController{}, POST+":UpdateVehicle")

	beego.Router("/"+PROJECT+"/basedata/newVehicle",
		&controllers.VehicleController{}, POST+":NewVehicle")

	beego.Router("/"+PROJECT+"/basedata/delVehicle",
		&controllers.VehicleController{}, POST+":DelVehicle")
}

func Warehouse() {
	beego.Router("/"+PROJECT+"/basedata/queryWarehouse",
		&controllers.WarehouseController{}, GET+":GetAllWarehouses")

	beego.Router("/"+PROJECT+"/warehouse/getWarehouses",
		&controllers.WarehouseController{}, POST+":GetWarehouses")

	beego.Router("/"+PROJECT+"/warehouse/getWarehouseById",
		&controllers.WarehouseController{}, POST+":GetWarehouseById")

	beego.Router("/"+PROJECT+"/warehouse/editWarehouseById",
		&controllers.WarehouseController{}, POST+":EditWarehouseById")

	beego.Router("/"+PROJECT+"/warehouse/addWarehouse",
		&controllers.WarehouseController{}, POST+":AddWarehouse")

	beego.Router("/"+PROJECT+"/warehouse/deleteWarehouse",
		&controllers.WarehouseController{}, POST+":DeleteWarehouse")
}
