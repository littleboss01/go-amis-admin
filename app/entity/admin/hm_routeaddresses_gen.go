package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmRouteaddresses struct {
	Routeaddressid      int32  `gorm:"column:routeaddressid;autoIncrement;type:mediumint(9);primaryKey;index:routeaddressid,class:BTREE,unique" json:"routeaddressid"` //
	Routeaddressrouteid int32  `gorm:"column:routeaddressrouteid;type:int(11)" json:"routeaddressrouteid"`                                                             //
	Routeaddressaddress string `gorm:"column:routeaddressaddress;type:varchar(255)" json:"routeaddressaddress"`                                                        //
}

var (
	HmRouteaddressesFieldRouteaddressid      = "routeaddressid"
	HmRouteaddressesFieldRouteaddressrouteid = "routeaddressrouteid"
	HmRouteaddressesFieldRouteaddressaddress = "routeaddressaddress"
)

func (receiver *HmRouteaddresses) TableName() string {
	return "hm_routeaddresses"
}

type OrmHmRouteaddresses struct {
	db *gorm.DB
}

func (orm *OrmHmRouteaddresses) TableName() string {
	return "hm_routeaddresses"
}

func NewOrmHmRouteaddresses(txs ...*gorm.DB) *OrmHmRouteaddresses {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmRouteaddresses{})
	} else {
		tx = txs[0].Model(&HmRouteaddresses{})
	}
	return &OrmHmRouteaddresses{db: tx}
}

func (orm *OrmHmRouteaddresses) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmRouteaddresses) GetTableInfo() interface{} {
	return &HmRouteaddresses{}
}

// Create insert the value into database
func (orm *OrmHmRouteaddresses) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmRouteaddresses) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmRouteaddresses) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmRouteaddresses) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmRouteaddresses) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmRouteaddresses) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmRouteaddresses) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmRouteaddresses) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmRouteaddresses) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmRouteaddresses) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmRouteaddresses) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmRouteaddresses) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmRouteaddresses) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmRouteaddresses) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmRouteaddresses) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmRouteaddresses) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmRouteaddresses) Unscoped() *OrmHmRouteaddresses {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmRouteaddresses) Insert(row *HmRouteaddresses) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmRouteaddresses) Inserts(rows []*HmRouteaddresses) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmRouteaddresses) Order(value interface{}) *OrmHmRouteaddresses {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmRouteaddresses) Group(name string) *OrmHmRouteaddresses {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmRouteaddresses) Limit(limit int) *OrmHmRouteaddresses {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmRouteaddresses) Offset(offset int) *OrmHmRouteaddresses {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmRouteaddresses) Get() HmRouteaddressesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmRouteaddresses) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmRouteaddresses) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmRouteaddresses{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmRouteaddresses) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_routeaddresses")
}

func (orm *OrmHmRouteaddresses) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmRouteaddresses) First(conds ...interface{}) (*HmRouteaddresses, bool) {
	dest := &HmRouteaddresses{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmRouteaddresses) Take(conds ...interface{}) (*HmRouteaddresses, int64) {
	dest := &HmRouteaddresses{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmRouteaddresses) Last(conds ...interface{}) (*HmRouteaddresses, int64) {
	dest := &HmRouteaddresses{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmRouteaddresses) Find(conds ...interface{}) (HmRouteaddressesList, int64) {
	list := make([]*HmRouteaddresses, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmRouteaddresses) Paginate(page int, limit int) (HmRouteaddressesList, int64) {
	var total int64
	list := make([]*HmRouteaddresses, 0)
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
func (orm *OrmHmRouteaddresses) SimplePaginate(page int, limit int) HmRouteaddressesList {
	list := make([]*HmRouteaddresses, 0)
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
func (orm *OrmHmRouteaddresses) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmRouteaddresses) FirstOrInit(dest *HmRouteaddresses, conds ...interface{}) (*HmRouteaddresses, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmRouteaddresses) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRouteaddresses) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRouteaddresses) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmRouteaddresses) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmRouteaddresses) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmRouteaddresses) Where(query interface{}, args ...interface{}) *OrmHmRouteaddresses {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmRouteaddresses) Select(query interface{}, args ...interface{}) *OrmHmRouteaddresses {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmRouteaddresses) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmRouteaddresses) And(fuc func(orm *OrmHmRouteaddresses)) *OrmHmRouteaddresses {
	ormAnd := NewOrmHmRouteaddresses()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmRouteaddresses) Or(fuc func(orm *OrmHmRouteaddresses)) *OrmHmRouteaddresses {
	ormOr := NewOrmHmRouteaddresses()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmRouteaddresses) Preload(query string, args ...interface{}) *OrmHmRouteaddresses {
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
func (orm *OrmHmRouteaddresses) Joins(query string, args ...interface{}) *OrmHmRouteaddresses {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmRouteaddresses) WhereRouteaddressid(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressid` = ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) InsertGetRouteaddressid(row *HmRouteaddresses) int32 {
	orm.db.Create(row)
	return row.Routeaddressid
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressidIn(val []int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressid` IN ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressidGt(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressid` > ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressidGte(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressid` >= ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressidLt(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressid` < ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressidLte(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressid` <= ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressrouteid(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressrouteid` = ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressrouteidIn(val []int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressrouteid` IN ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressrouteidNe(val int32) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressrouteid` <> ?", val)
	return orm
}
func (orm *OrmHmRouteaddresses) WhereRouteaddressaddress(val string) *OrmHmRouteaddresses {
	orm.db.Where("`routeaddressaddress` = ?", val)
	return orm
}

type HmRouteaddressesList []*HmRouteaddresses

func (l HmRouteaddressesList) GetRouteaddressidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Routeaddressid)
	}
	return got
}
func (l HmRouteaddressesList) GetRouteaddressidMap() map[int32]*HmRouteaddresses {
	got := make(map[int32]*HmRouteaddresses)
	for _, val := range l {
		got[val.Routeaddressid] = val
	}
	return got
}
func (l HmRouteaddressesList) GetRouteaddressrouteidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Routeaddressrouteid)
	}
	return got
}
func (l HmRouteaddressesList) GetRouteaddressrouteidMap() map[int32]*HmRouteaddresses {
	got := make(map[int32]*HmRouteaddresses)
	for _, val := range l {
		got[val.Routeaddressrouteid] = val
	}
	return got
}
