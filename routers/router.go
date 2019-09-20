package routers

import (
	"erpweb/controllers"
	"github.com/astaxie/beego"
)

const (
	PROJECT = "api/basedata"
	GET     = "get"
	POST    = "post"
)

func init() {
	beego.Router("/", &controllers.MainController{}, GET+":Index")
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

func Login() {
	beego.Router("/"+PROJECT+"/Login",
		&controllers.LoginController{}, POST+":Login")

	beego.Router("/"+PROJECT+"/Loginout",
		&controllers.LoginController{}, GET+":Loginout")
}

func Account() {
	//api/employee/getAccountList
	beego.Router("/"+PROJECT+"/account/getAccountList",
		&controllers.AccountController{}, GET+":GetAccountList")

	beego.Router("/"+PROJECT+"/account/getAccounts",
		&controllers.AccountController{}, POST+":GetAccounts")

	beego.Router("/"+PROJECT+"/account/getAccountById",
		&controllers.AccountController{}, POST+":GetAccountById")

	beego.Router("/"+PROJECT+"/account/editAccountById",
		&controllers.AccountController{}, POST+":EditAccountById")

	//16.修改账号
	//url: api/employee/updateAccountStatus
	beego.Router("/"+PROJECT+"/account/updateAccountStatus",
		&controllers.AccountController{}, POST+":EditAccountStatusById")

	//开通ERP账号
	//api/employee/openAccount
	beego.Router("/"+PROJECT+"/account/openAccount",
		&controllers.AccountController{}, POST+":AddAccount")

	beego.Router("/"+PROJECT+"/account/deleteAccount",
		&controllers.AccountController{}, POST+":DeleteAccount")
}

func Arrivalbill() {
	beego.Router("/"+PROJECT+"/arrivalbill/getArrivalbills",
		&controllers.ArrivalbillController{}, POST+":GetArrivalbills")

	beego.Router("/"+PROJECT+"/arrivalbill/getArrivalbillById",
		&controllers.ArrivalbillController{}, POST+":GetArrivalbillById")

	beego.Router("/"+PROJECT+"/arrivalbill/editArrivalbillById",
		&controllers.ArrivalbillController{}, POST+":EditArrivalbillById")

	beego.Router("/"+PROJECT+"/arrivalbill/addArrivalbill",
		&controllers.ArrivalbillController{}, POST+":AddArrivalbill")

	beego.Router("/"+PROJECT+"/arrivalbill/deleteArrivalbill",
		&controllers.ArrivalbillController{}, POST+":DeleteArrivalbill")
}

func Arrivaldetail() {
	preRoute := "arrivaldetail"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getArrivaldetails",
		&controllers.ArrivaldetailController{}, POST+":GetArrivaldetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getArrivaldetailById",
		&controllers.ArrivaldetailController{}, POST+":GetArrivaldetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editArrivaldetailById",
		&controllers.ArrivaldetailController{}, POST+":EditArrivaldetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addArrivaldetail",
		&controllers.ArrivaldetailController{}, POST+":AddArrivaldetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteArrivaldetail",
		&controllers.ArrivaldetailController{}, POST+":DeleteArrivaldetail")
}

func Customer() {
	beego.Router("/"+PROJECT+"/queryCustomer",
		&controllers.CustomerController{}, GET+":QueryCustomer")

	beego.Router("/"+PROJECT+"/customer/getCustomers",
		&controllers.CustomerController{}, POST+":GetCustomers")

	beego.Router("/"+PROJECT+"/customer/getCustomerById",
		&controllers.CustomerController{}, POST+":GetCustomerById")

	beego.Router("/"+PROJECT+"/customer/editCustomerById",
		&controllers.CustomerController{}, POST+":EditCustomerById")

	beego.Router("/"+PROJECT+"/customer/addCustomer",
		&controllers.CustomerController{}, POST+":AddCustomer")

	beego.Router("/"+PROJECT+"/customer/deleteCustomer",
		&controllers.CustomerController{}, POST+":DeleteCustomer")
}

func Department() {
	beego.Router("/"+PROJECT+"/queryDept",
		&controllers.DepartmentController{}, GET+":QueryDept")

	beego.Router("/"+PROJECT+"/department/getDepartments",
		&controllers.DepartmentController{}, POST+":GetDepartments")

	beego.Router("/"+PROJECT+"/department/getDepartmentById",
		&controllers.DepartmentController{}, POST+":GetDepartmentById")

	beego.Router("/"+PROJECT+"/department/editDepartmentById",
		&controllers.DepartmentController{}, POST+":EditDepartmentById")

	beego.Router("/"+PROJECT+"/department/addDepartment",
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
	preRoute := "financeflow"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getFinanceflows",
		&controllers.FinanceflowController{}, POST+":GetFinanceflows")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getFinanceflowById",
		&controllers.FinanceflowController{}, POST+":GetFinanceflowById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editFinanceflowById",
		&controllers.FinanceflowController{}, POST+":EditFinanceflowById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addFinanceflow",
		&controllers.FinanceflowController{}, POST+":AddFinanceflow")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteFinanceflow",
		&controllers.FinanceflowController{}, POST+":DeleteFinanceflow")
}

func Inquiry() {
	preRoute := "inquiry"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquirys",
		&controllers.InquiryController{}, POST+":GetInquirys")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquiryById",
		&controllers.InquiryController{}, POST+":GetInquiryById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editInquiryById",
		&controllers.InquiryController{}, POST+":EditInquiryById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addInquiry",
		&controllers.InquiryController{}, POST+":AddInquiry")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteInquiry",
		&controllers.InquiryController{}, POST+":DeleteInquiry")
}

func Inquirydetail() {
	preRoute := "inquirydetail"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquirydetails",
		&controllers.InquirydetailController{}, POST+":GetInquirydetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getInquirydetailById",
		&controllers.InquirydetailController{}, POST+":GetInquirydetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editInquirydetailById",
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
	beego.Router("/"+PROJECT+"/employee/getPersons",
		&controllers.EmployeeController{}, POST+":GetEmployees")

	//get person by id
	// /erp/employee/getPerson
	beego.Router("/"+PROJECT+"/employee/getPerson",
		&controllers.EmployeeController{}, POST+":GetEmployeeById")

	//edit person info by id
	// /erp/employee/updateEmployee
	beego.Router("/"+PROJECT+"/employee/updateEmployee",
		&controllers.EmployeeController{}, POST+":EditEmployeeById")

	//add person by id
	// /erp/employee/newEmployee
	beego.Router("/"+PROJECT+"/employee/newEmployee",
		&controllers.EmployeeController{}, POST+":AddEmployee")

	beego.Router("/"+PROJECT+"/employee/deletePerson",
		&controllers.EmployeeController{}, POST+":DeleteEmployee")
}

func Leave() {
	beego.Router("/"+PROJECT+"/leave/getLeaves",
		&controllers.LeaveController{}, POST+":GetLeaves")

	beego.Router("/"+PROJECT+"/leave/getLeaveById",
		&controllers.LeaveController{}, POST+":GetLeaveById")

	beego.Router("/"+PROJECT+"/leave/editLeaveById",
		&controllers.LeaveController{}, POST+":EditLeaveById")

	beego.Router("/"+PROJECT+"/leave/addLeave",
		&controllers.LeaveController{}, POST+":AddLeave")

	beego.Router("/"+PROJECT+"/leave/deleteLeave",
		&controllers.LeaveController{}, POST+":DeleteLeave")
}

func Marketcontract() {
	beego.Router("/"+PROJECT+"/marketcontract/getMarketcontracts",
		&controllers.MarketcontractController{}, POST+":GetMarketcontracts")

	beego.Router("/"+PROJECT+"/marketcontract/getMarketcontractById",
		&controllers.MarketcontractController{}, POST+":GetMarketcontractById")

	beego.Router("/"+PROJECT+"/marketcontract/editMarketcontractById",
		&controllers.MarketcontractController{}, POST+":EditMarketcontractById")

	beego.Router("/"+PROJECT+"/marketcontract/addMarketcontract",
		&controllers.MarketcontractController{}, POST+":AddMarketcontract")

	beego.Router("/"+PROJECT+"/marketcontract/deleteMarketcontract",
		&controllers.MarketcontractController{}, POST+":DeleteMarketcontract")
}

func Matter() {

	beego.Router("/"+PROJECT+"/queryMatter",
		&controllers.MatterController{}, GET+":QueryMatter")

	beego.Router("/"+PROJECT+"/matter/getMatters",
		&controllers.MatterController{}, POST+":GetMatters")

	beego.Router("/"+PROJECT+"/matter/getMatterById",
		&controllers.MatterController{}, POST+":GetMatterById")

	beego.Router("/"+PROJECT+"/matter/editMatterById",
		&controllers.MatterController{}, POST+":EditMatterById")

	beego.Router("/"+PROJECT+"/matter/addMatter",
		&controllers.MatterController{}, POST+":AddMatter")

	beego.Router("/"+PROJECT+"/matter/deleteMatter",
		&controllers.MatterController{}, POST+":DeleteMatter")
}

func Matterplan() {
	preRoute := "matterplan"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getMatterplans",
		&controllers.MatterplanController{}, POST+":GetMatterplans")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getMatterplanById",
		&controllers.MatterplanController{}, POST+":GetMatterplanById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editMatterplanById",
		&controllers.MatterplanController{}, POST+":EditMatterplanById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addMatterplan",
		&controllers.MatterplanController{}, POST+":AddMatterplan")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteMatterplan",
		&controllers.MatterplanController{}, POST+":DeleteMatterplan")
}

func Menu() {
	beego.Router("/"+PROJECT+"/queryMenu",
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
	preRoute := "outofdetail"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofdetails",
		&controllers.OutofdetailController{}, POST+":GetOutofdetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofdetailById",
		&controllers.OutofdetailController{}, POST+":GetOutofdetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editOutofdetailById",
		&controllers.OutofdetailController{}, POST+":EditOutofdetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addOutofdetail",
		&controllers.OutofdetailController{}, POST+":AddOutofdetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteOutofdetail",
		&controllers.OutofdetailController{}, POST+":DeleteOutofdetail")
}

func Outofstore() {
	preRoute := "outofstore"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofstores",
		&controllers.OutofstoreController{}, POST+":GetOutofstores")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getOutofstoreById",
		&controllers.OutofstoreController{}, POST+":GetOutofstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editOutofstoreById",
		&controllers.OutofstoreController{}, POST+":EditOutofstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addOutofstore",
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
}

func Purchasecontract() {
	preRoute := "purchasecontract"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasecontracts",
		&controllers.PurchasecontractController{}, POST+":GetPurchasecontracts")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasecontractById",
		&controllers.PurchasecontractController{}, POST+":GetPurchasecontractById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editPurchasecontractById",
		&controllers.PurchasecontractController{}, POST+":EditPurchasecontractById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPurchasecontract",
		&controllers.PurchasecontractController{}, POST+":AddPurchasecontract")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deletePurchasecontract",
		&controllers.PurchasecontractController{}, POST+":DeletePurchasecontract")
}

func Purchasedetail() {
	preRoute := "purchasedetail"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasedetails",
		&controllers.PurchasedetailController{}, POST+":GetPurchasedetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPurchasedetailById",
		&controllers.PurchasedetailController{}, POST+":GetPurchasedetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editPurchasedetailById",
		&controllers.PurchasedetailController{}, POST+":EditPurchasedetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPurchasedetail",
		&controllers.PurchasedetailController{}, POST+":AddPurchasedetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deletePurchasedetail",
		&controllers.PurchasedetailController{}, POST+":DeletePurchasedetail")
}

func Putindetail() {
	preRoute := "putindetail"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutindetails",
		&controllers.PutindetailController{}, POST+":GetPutindetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutindetailById",
		&controllers.PutindetailController{}, POST+":GetPutindetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editPutindetailById",
		&controllers.PutindetailController{}, POST+":EditPutindetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPutindetail",
		&controllers.PutindetailController{}, POST+":AddPutindetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deletePutindetail",
		&controllers.PutindetailController{}, POST+":DeletePutindetail")
}
func Putinstore() {
	preRoute := "putinstore"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutinstores",
		&controllers.PutinstoreController{}, POST+":GetPutinstores")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getPutinstoreById",
		&controllers.PutinstoreController{}, POST+":GetPutinstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editPutinstoreById",
		&controllers.PutinstoreController{}, POST+":EditPutinstoreById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addPutinstore",
		&controllers.PutinstoreController{}, POST+":AddPutinstore")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deletePutinstore",
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
	preRoute := "repaircost"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepaircosts",
		&controllers.RepaircostController{}, POST+":GetRepaircosts")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepaircostById",
		&controllers.RepaircostController{}, POST+":GetRepaircostById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editRepaircostById",
		&controllers.RepaircostController{}, POST+":EditRepaircostById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addRepaircost",
		&controllers.RepaircostController{}, POST+":AddRepaircost")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteRepaircost",
		&controllers.RepaircostController{}, POST+":DeleteRepaircost")
}
func Repairitem() {
	preRoute := "repairitem"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepairitems",
		&controllers.RepairitemController{}, POST+":GetRepairitems")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getRepairitemById",
		&controllers.RepairitemController{}, POST+":GetRepairitemById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editRepairitemById",
		&controllers.RepairitemController{}, POST+":EditRepairitemById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addRepairitem",
		&controllers.RepairitemController{}, POST+":AddRepairitem")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteRepairitem",
		&controllers.RepairitemController{}, POST+":DeleteRepairitem")
}
func Review() {
	preRoute := "review"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getReviews",
		&controllers.ReviewController{}, POST+":GetReviews")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getReviewById",
		&controllers.ReviewController{}, POST+":GetReviewById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editReviewById",
		&controllers.ReviewController{}, POST+":EditReviewById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addReview",
		&controllers.ReviewController{}, POST+":AddReview")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteReview",
		&controllers.ReviewController{}, POST+":DeleteReview")
}

func Reviewresult() {
	preRoute := "reviewresult"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getReviewresults",
		&controllers.ReviewresultController{}, POST+":GetReviewresults")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getReviewresultById",
		&controllers.ReviewresultController{}, POST+":GetReviewresultById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editReviewresultById",
		&controllers.ReviewresultController{}, POST+":EditReviewresultById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addReviewresult",
		&controllers.ReviewresultController{}, POST+":AddReviewresult")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteReviewresult",
		&controllers.ReviewresultController{}, POST+":DeleteReviewresult")
}

func Saledetail() {
	preRoute := "saledetail"
	beego.Router("/"+PROJECT+"/"+preRoute+"/getSaledetails",
		&controllers.SaledetailController{}, POST+":GetSaledetails")

	beego.Router("/"+PROJECT+"/"+preRoute+"/getSaledetailById",
		&controllers.SaledetailController{}, POST+":GetSaledetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/editSaledetailById",
		&controllers.SaledetailController{}, POST+":EditSaledetailById")

	beego.Router("/"+PROJECT+"/"+preRoute+"/addSaledetail",
		&controllers.SaledetailController{}, POST+":AddSaledetail")

	beego.Router("/"+PROJECT+"/"+preRoute+"/deleteSaledetail",
		&controllers.SaledetailController{}, POST+":DeleteSaledetail")
}

func Stock() {
	preRoute := "stock"
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
	beego.Router("/"+PROJECT+"/querySupplier",
		&controllers.SupplierController{}, GET+":QuerySupplier")

	beego.Router("/"+PROJECT+"/supplier/getSuppliers",
		&controllers.SupplierController{}, POST+":GetSuppliers")

	beego.Router("/"+PROJECT+"/supplier/getSupplierById",
		&controllers.SupplierController{}, POST+":GetSupplierById")

	beego.Router("/"+PROJECT+"/supplier/editSupplierById",
		&controllers.SupplierController{}, POST+":EditSupplierById")

	beego.Router("/"+PROJECT+"/supplier/addSupplier",
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
	// beego.Router("/"+PROJECT+"/vehicle/getVehicles",
	// 	&controllers.VehicleController{}, POST+":GetVehicles")

	// beego.Router("/"+PROJECT+"/vehicle/getVehicleById",
	// 	&controllers.VehicleController{}, POST+":GetVehicleById")

	// beego.Router("/"+PROJECT+"/vehicle/editVehicleById",
	// 	&controllers.VehicleController{}, POST+":EditVehicleById")

	// beego.Router("/"+PROJECT+"/vehicle/addVehicle",
	// 	&controllers.VehicleController{}, POST+":AddVehicle")

	// beego.Router("/"+PROJECT+"/vehicle/deleteVehicle",
	// 	&controllers.VehicleController{}, POST+":DeleteVehicle")
}

func Warehouse() {
	beego.Router("/"+PROJECT+"/queryWarehouse",
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
