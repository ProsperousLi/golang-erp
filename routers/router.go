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
	Employee()
	Account()
	Customer()
	Department()
	Duty()
	Leave()
	Marketcontract()
	Matter()
	Menu()
	//test
	Operlog()
	Permission()
	//controller
	Reviewresult()
	Saledetail()
	Stock()
	Supplier()
	//end
	Warehouse()
}

func Account() {
	beego.Router("/"+PROJECT+"/account/getAccounts",
		&controllers.AccountController{}, POST+":GetAccounts")

	beego.Router("/"+PROJECT+"/account/getAccountById",
		&controllers.AccountController{}, POST+":GetAccountById")

	beego.Router("/"+PROJECT+"/account/editAccountById",
		&controllers.AccountController{}, POST+":EditAccountById")

	beego.Router("/"+PROJECT+"/account/addAccount",
		&controllers.AccountController{}, POST+":AddAccount")

	beego.Router("/"+PROJECT+"/account/deleteAccount",
		&controllers.AccountController{}, POST+":DeleteAccount")
}

func Customer() {
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

//人员信息 employee
func Employee() {
	//get persons by page
	// /erp/employee/getPersons
	beego.Router("/"+PROJECT+"/employee/getPersons",
		&controllers.EmployeeController{}, POST+":GetEmployees")

	//get person by id
	// /erp/employee/getPerson
	beego.Router("/"+PROJECT+"/employee/getPerson",
		&controllers.EmployeeController{}, POST+":GetEmployeeById")

	//edit person info by id
	// /erp/employee/editPerson
	beego.Router("/"+PROJECT+"/employee/editPerson",
		&controllers.EmployeeController{}, POST+":EditEmployeeById")

	//add person by id
	// /erp/employee/addPerson
	beego.Router("/"+PROJECT+"/employee/addPerson",
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

func Reviewresult() {
	// beego.Router("/"+PROJECT+"/reviewresult/getReviewresults",
	// 	&controllers.ReviewresultController{}, POST+":GetReviewresults")

	// beego.Router("/"+PROJECT+"/reviewresult/getReviewresultById",
	// 	&controllers.ReviewresultController{}, POST+":GetReviewresultById")

	// beego.Router("/"+PROJECT+"/reviewresult/editReviewresultById",
	// 	&controllers.ReviewresultController{}, POST+":EditReviewresultById")

	// beego.Router("/"+PROJECT+"/reviewresult/addReviewresult",
	// 	&controllers.ReviewresultController{}, POST+":AddReviewresult")

	// beego.Router("/"+PROJECT+"/reviewresult/deleteReviewresult",
	// 	&controllers.ReviewresultController{}, POST+":DeleteReviewresult")
}

func Saledetail() {
	// beego.Router("/"+PROJECT+"/saledetail/getSaledetails",
	// 	&controllers.SaledetailController{}, POST+":GetSaledetails")

	// beego.Router("/"+PROJECT+"/saledetail/getSaledetailById",
	// 	&controllers.SaledetailController{}, POST+":GetSaledetailById")

	// beego.Router("/"+PROJECT+"/saledetail/editSaledetailById",
	// 	&controllers.SaledetailController{}, POST+":EditSaledetailById")

	// beego.Router("/"+PROJECT+"/saledetail/addSaledetail",
	// 	&controllers.SaledetailController{}, POST+":AddSaledetail")

	// beego.Router("/"+PROJECT+"/saledetail/deleteSaledetail",
	// 	&controllers.SaledetailController{}, POST+":DeleteSaledetail")
}

func Stock() {
	// beego.Router("/"+PROJECT+"/stock/getStocks",
	// 	&controllers.StockController{}, POST+":GetStocks")

	// beego.Router("/"+PROJECT+"/stock/getStockById",
	// 	&controllers.StockController{}, POST+":GetStockById")

	// beego.Router("/"+PROJECT+"/stock/editStockById",
	// 	&controllers.StockController{}, POST+":EditStockById")

	// beego.Router("/"+PROJECT+"/stock/addStock",
	// 	&controllers.StockController{}, POST+":AddStock")

	// beego.Router("/"+PROJECT+"/stock/deleteStock",
	// 	&controllers.StockController{}, POST+":DeleteStock")
}

func Supplier() {
	// beego.Router("/"+PROJECT+"/supplier/getSuppliers",
	// 	&controllers.SupplierController{}, POST+":GetSuppliers")

	// beego.Router("/"+PROJECT+"/supplier/getSupplierById",
	// 	&controllers.SupplierController{}, POST+":GetSupplierById")

	// beego.Router("/"+PROJECT+"/supplier/editSupplierById",
	// 	&controllers.SupplierController{}, POST+":EditSupplierById")

	// beego.Router("/"+PROJECT+"/supplier/addSupplier",
	// 	&controllers.SupplierController{}, POST+":AddSupplier")

	// beego.Router("/"+PROJECT+"/supplier/deleteSupplier",
	// 	&controllers.SupplierController{}, POST+":DeleteSupplier")
}

func Supplyrelation() {
	// beego.Router("/"+PROJECT+"/supplyrelation/getSupplyrelations",
	// 	&controllers.SupplyrelationController{}, POST+":GetSupplyrelations")

	// beego.Router("/"+PROJECT+"/supplyrelation/getSupplyrelationById",
	// 	&controllers.SupplyrelationController{}, POST+":GetSupplyrelationById")

	// beego.Router("/"+PROJECT+"/supplyrelation/editSupplyrelationById",
	// 	&controllers.SupplyrelationController{}, POST+":EditSupplyrelationById")

	// beego.Router("/"+PROJECT+"/supplyrelation/addSupplyrelation",
	// 	&controllers.SupplyrelationController{}, POST+":AddSupplyrelation")

	// beego.Router("/"+PROJECT+"/supplyrelation/deleteSupplyrelation",
	// 	&controllers.SupplyrelationController{}, POST+":DeleteSupplyrelation")
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
