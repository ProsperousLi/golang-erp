package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

const (
	sqlUser = "root"
	sqlPwd  = "123456"
)

var (
	OSQL orm.Ormer
)

func Init() {
	//log.Info("init sql...")
	// set default database
	orm.RegisterDataBase("default", "mysql", sqlUser+":"+sqlPwd+"@/erp?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(Account), new(Customer), new(Department) /*new(Duty),*/, new(Employee),
		new(Leaves), new(Matter) /*new(Matterpackage),*/, new(Menu), new(Operlog), /*new(Packagerelation),*/
		new(Permission), new(Reviewresult), new(Marketcontract), new(Saledetail), new(Stock),
		new(Supplier), new(Supplyrelation), new(Vehicle), new(Warehouse))

	// create table
	orm.RunSyncdb("default", false, true)

	OSQL = orm.NewOrm()

}
