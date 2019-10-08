package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

const (
	sqlUser = "root"
	sqlPwd  = "lrf64787"
)

var (
	OSQL orm.Ormer
)

var (
	TimeStampMaps = make(map[string]int64) //more 时间戳 = 0
	TimeStamp     int64
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

	InitTime()

	go TimeToClean()

}

func InitTime() {
	hour := time.Now().Hour()
	//(10000/24)*2
	TimeStamp = (10000 / 24) * int64(hour)
}

func TimeToClean() {
	for {
		now := time.Now()               //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(),
			0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now)) //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		//Printf("凌晨清零: %v\n", time.Now())
		//以下为定时执行的操作
		TimeStampMaps = make(map[string]int64)
		TimeStamp = 0
	}
}
