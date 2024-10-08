package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmRoutes struct {
	Routeid                       int32  `gorm:"column:routeid;autoIncrement;type:int(11);primaryKey;index:routeid,class:BTREE,unique" json:"routeid"` //
	Routedomainname               string `gorm:"column:routedomainname;type:varchar(255)" json:"routedomainname"`                                      //
	Routedescription              string `gorm:"column:routedescription;type:varchar(255)" json:"routedescription"`                                    //
	Routetargetsmthost            string `gorm:"column:routetargetsmthost;type:varchar(255)" json:"routetargetsmthost"`                                //
	Routetargetsmtport            int32  `gorm:"column:routetargetsmtport;type:int(11)" json:"routetargetsmtport"`                                     //
	Routenooftries                int32  `gorm:"column:routenooftries;type:int(11)" json:"routenooftries"`                                             //
	Routeminutesbetweentry        int32  `gorm:"column:routeminutesbetweentry;type:int(11)" json:"routeminutesbetweentry"`                             //
	Routealladdresses             uint32 `gorm:"column:routealladdresses;type:tinyint(3) unsigned" json:"routealladdresses"`                           //
	Routeuseauthentication        int32  `gorm:"column:routeuseauthentication;type:tinyint(4)" json:"routeuseauthentication"`                          //
	Routeauthenticationusername   string `gorm:"column:routeauthenticationusername;type:varchar(255)" json:"routeauthenticationusername"`              //
	Routeauthenticationpassword   string `gorm:"column:routeauthenticationpassword;type:varchar(255)" json:"routeauthenticationpassword"`              //
	Routetreatsecurityaslocal     int32  `gorm:"column:routetreatsecurityaslocal;type:tinyint(4)" json:"routetreatsecurityaslocal"`                    //
	Routeconnectionsecurity       int32  `gorm:"column:routeconnectionsecurity;type:tinyint(4)" json:"routeconnectionsecurity"`                        //
	Routetreatsenderaslocaldomain int32  `gorm:"column:routetreatsenderaslocaldomain;type:tinyint(4)" json:"routetreatsenderaslocaldomain"`            //
}

var (
	HmRoutesFieldRouteid                       = "routeid"
	HmRoutesFieldRoutedomainname               = "routedomainname"
	HmRoutesFieldRoutedescription              = "routedescription"
	HmRoutesFieldRoutetargetsmthost            = "routetargetsmthost"
	HmRoutesFieldRoutetargetsmtport            = "routetargetsmtport"
	HmRoutesFieldRoutenooftries                = "routenooftries"
	HmRoutesFieldRouteminutesbetweentry        = "routeminutesbetweentry"
	HmRoutesFieldRoutealladdresses             = "routealladdresses"
	HmRoutesFieldRouteuseauthentication        = "routeuseauthentication"
	HmRoutesFieldRouteauthenticationusername   = "routeauthenticationusername"
	HmRoutesFieldRouteauthenticationpassword   = "routeauthenticationpassword"
	HmRoutesFieldRoutetreatsecurityaslocal     = "routetreatsecurityaslocal"
	HmRoutesFieldRouteconnectionsecurity       = "routeconnectionsecurity"
	HmRoutesFieldRoutetreatsenderaslocaldomain = "routetreatsenderaslocaldomain"
)

func (receiver *HmRoutes) TableName() string {
	return "hm_routes"
}

type OrmHmRoutes struct {
	db *gorm.DB
}

func (orm *OrmHmRoutes) TableName() string {
	return "hm_routes"
}

func NewOrmHmRoutes(txs ...*gorm.DB) *OrmHmRoutes {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmRoutes{})
	} else {
		tx = txs[0].Model(&HmRoutes{})
	}
	return &OrmHmRoutes{db: tx}
}

func (orm *OrmHmRoutes) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmRoutes) GetTableInfo() interface{} {
	return &HmRoutes{}
}

// Create insert the value into database
func (orm *OrmHmRoutes) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmRoutes) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmRoutes) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmRoutes) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmRoutes) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmRoutes) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmRoutes) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmRoutes) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmRoutes) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmRoutes) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmRoutes) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmRoutes) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmRoutes) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmRoutes) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmRoutes) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmRoutes) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmRoutes) Unscoped() *OrmHmRoutes {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmRoutes) Insert(row *HmRoutes) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmRoutes) Inserts(rows []*HmRoutes) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmRoutes) Order(value interface{}) *OrmHmRoutes {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmRoutes) Group(name string) *OrmHmRoutes {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmRoutes) Limit(limit int) *OrmHmRoutes {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmRoutes) Offset(offset int) *OrmHmRoutes {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmRoutes) Get() HmRoutesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmRoutes) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmRoutes) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmRoutes{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmRoutes) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_routes")
}

func (orm *OrmHmRoutes) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmRoutes) First(conds ...interface{}) (*HmRoutes, bool) {
	dest := &HmRoutes{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmRoutes) Take(conds ...interface{}) (*HmRoutes, int64) {
	dest := &HmRoutes{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmRoutes) Last(conds ...interface{}) (*HmRoutes, int64) {
	dest := &HmRoutes{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmRoutes) Find(conds ...interface{}) (HmRoutesList, int64) {
	list := make([]*HmRoutes, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmRoutes) Paginate(page int, limit int) (HmRoutesList, int64) {
	var total int64
	list := make([]*HmRoutes, 0)
	orm.db.Count(&total)
	if total > 0 {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * limit
		tx := orm.db.Offset(offset).Limit(limit).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}

	return list, total
}

// SimplePaginate 不统计总数的分页
func (orm *OrmHmRoutes) SimplePaginate(page int, limit int) HmRoutesList {
	list := make([]*HmRoutes, 0)
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * limit
	tx := orm.db.Offset(offset).Limit(limit).Find(&list)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list
}

// FindInBatches find records in batches
func (orm *OrmHmRoutes) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmRoutes) FirstOrInit(dest *HmRoutes, conds ...interface{}) (*HmRoutes, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmRoutes) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRoutes) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRoutes) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmRoutes) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmRoutes) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmRoutes) Where(query interface{}, args ...interface{}) *OrmHmRoutes {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmRoutes) Select(query interface{}, args ...interface{}) *OrmHmRoutes {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmRoutes) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmRoutes) And(fuc func(orm *OrmHmRoutes)) *OrmHmRoutes {
	ormAnd := NewOrmHmRoutes()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmRoutes) Or(fuc func(orm *OrmHmRoutes)) *OrmHmRoutes {
	ormOr := NewOrmHmRoutes()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmRoutes) Preload(query string, args ...interface{}) *OrmHmRoutes {
	arr := strings.Split(query, ".")
	for i, _ := range arr {
		arr[i] = home.StringToHump(arr[i])
	}
	orm.db.Preload(strings.Join(arr, "."), args...)
	return orm
}

// Joins specify Joins conditions
// db.Joins("Account|account").Find(&user)
// db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
// db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))
func (orm *OrmHmRoutes) Joins(query string, args ...interface{}) *OrmHmRoutes {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmRoutes) WhereRouteid(val int32) *OrmHmRoutes {
	orm.db.Where("`routeid` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) InsertGetRouteid(row *HmRoutes) int32 {
	orm.db.Create(row)
	return row.Routeid
}
func (orm *OrmHmRoutes) WhereRouteidIn(val []int32) *OrmHmRoutes {
	orm.db.Where("`routeid` IN ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteidGt(val int32) *OrmHmRoutes {
	orm.db.Where("`routeid` > ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteidGte(val int32) *OrmHmRoutes {
	orm.db.Where("`routeid` >= ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteidLt(val int32) *OrmHmRoutes {
	orm.db.Where("`routeid` < ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteidLte(val int32) *OrmHmRoutes {
	orm.db.Where("`routeid` <= ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutedomainname(val string) *OrmHmRoutes {
	orm.db.Where("`routedomainname` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutedescription(val string) *OrmHmRoutes {
	orm.db.Where("`routedescription` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutetargetsmthost(val string) *OrmHmRoutes {
	orm.db.Where("`routetargetsmthost` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutetargetsmtport(val int32) *OrmHmRoutes {
	orm.db.Where("`routetargetsmtport` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutenooftries(val int32) *OrmHmRoutes {
	orm.db.Where("`routenooftries` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteminutesbetweentry(val int32) *OrmHmRoutes {
	orm.db.Where("`routeminutesbetweentry` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutealladdresses(val uint32) *OrmHmRoutes {
	orm.db.Where("`routealladdresses` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteuseauthentication(val int32) *OrmHmRoutes {
	orm.db.Where("`routeuseauthentication` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteauthenticationusername(val string) *OrmHmRoutes {
	orm.db.Where("`routeauthenticationusername` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteauthenticationpassword(val string) *OrmHmRoutes {
	orm.db.Where("`routeauthenticationpassword` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutetreatsecurityaslocal(val int32) *OrmHmRoutes {
	orm.db.Where("`routetreatsecurityaslocal` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRouteconnectionsecurity(val int32) *OrmHmRoutes {
	orm.db.Where("`routeconnectionsecurity` = ?", val)
	return orm
}
func (orm *OrmHmRoutes) WhereRoutetreatsenderaslocaldomain(val int32) *OrmHmRoutes {
	orm.db.Where("`routetreatsenderaslocaldomain` = ?", val)
	return orm
}

type HmRoutesList []*HmRoutes

func (l HmRoutesList) GetRouteidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Routeid)
	}
	return got
}
func (l HmRoutesList) GetRouteidMap() map[int32]*HmRoutes {
	got := make(map[int32]*HmRoutes)
	for _, val := range l {
		got[val.Routeid] = val
	}
	return got
}
