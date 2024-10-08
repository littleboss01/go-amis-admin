package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmLogonFailures struct {
	Ipaddress1  int64         `gorm:"column:ipaddress1;type:bigint(20);index:idx_hm_logon_failures_ipaddress,class:BTREE" json:"ipaddress1"`   //
	Ipaddress2  *int64        `gorm:"column:ipaddress2;type:bigint(20);index:idx_hm_logon_failures_ipaddress,class:BTREE" json:"ipaddress2"`   //
	Failuretime database.Time `gorm:"column:failuretime;type:datetime;index:idx_hm_logon_failures_failuretime,class:BTREE" json:"failuretime"` //
}

var (
	HmLogonFailuresFieldIpaddress1  = "ipaddress1"
	HmLogonFailuresFieldIpaddress2  = "ipaddress2"
	HmLogonFailuresFieldFailuretime = "failuretime"
)

func (receiver *HmLogonFailures) TableName() string {
	return "hm_logon_failures"
}

type OrmHmLogonFailures struct {
	db *gorm.DB
}

func (orm *OrmHmLogonFailures) TableName() string {
	return "hm_logon_failures"
}

func NewOrmHmLogonFailures(txs ...*gorm.DB) *OrmHmLogonFailures {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmLogonFailures{})
	} else {
		tx = txs[0].Model(&HmLogonFailures{})
	}
	return &OrmHmLogonFailures{db: tx}
}

func (orm *OrmHmLogonFailures) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmLogonFailures) GetTableInfo() interface{} {
	return &HmLogonFailures{}
}

// Create insert the value into database
func (orm *OrmHmLogonFailures) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmLogonFailures) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmLogonFailures) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmLogonFailures) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmLogonFailures) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmLogonFailures) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmLogonFailures) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmLogonFailures) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmLogonFailures) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmLogonFailures) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmLogonFailures) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmLogonFailures) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmLogonFailures) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmLogonFailures) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmLogonFailures) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmLogonFailures) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmLogonFailures) Unscoped() *OrmHmLogonFailures {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmLogonFailures) Insert(row *HmLogonFailures) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmLogonFailures) Inserts(rows []*HmLogonFailures) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmLogonFailures) Order(value interface{}) *OrmHmLogonFailures {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmLogonFailures) Group(name string) *OrmHmLogonFailures {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmLogonFailures) Limit(limit int) *OrmHmLogonFailures {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmLogonFailures) Offset(offset int) *OrmHmLogonFailures {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmLogonFailures) Get() HmLogonFailuresList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmLogonFailures) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmLogonFailures) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmLogonFailures{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmLogonFailures) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_logon_failures")
}

func (orm *OrmHmLogonFailures) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmLogonFailures) First(conds ...interface{}) (*HmLogonFailures, bool) {
	dest := &HmLogonFailures{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmLogonFailures) Take(conds ...interface{}) (*HmLogonFailures, int64) {
	dest := &HmLogonFailures{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmLogonFailures) Last(conds ...interface{}) (*HmLogonFailures, int64) {
	dest := &HmLogonFailures{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmLogonFailures) Find(conds ...interface{}) (HmLogonFailuresList, int64) {
	list := make([]*HmLogonFailures, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmLogonFailures) Paginate(page int, limit int) (HmLogonFailuresList, int64) {
	var total int64
	list := make([]*HmLogonFailures, 0)
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
func (orm *OrmHmLogonFailures) SimplePaginate(page int, limit int) HmLogonFailuresList {
	list := make([]*HmLogonFailures, 0)
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
func (orm *OrmHmLogonFailures) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmLogonFailures) FirstOrInit(dest *HmLogonFailures, conds ...interface{}) (*HmLogonFailures, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmLogonFailures) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmLogonFailures) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmLogonFailures) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmLogonFailures) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmLogonFailures) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmLogonFailures) Where(query interface{}, args ...interface{}) *OrmHmLogonFailures {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmLogonFailures) Select(query interface{}, args ...interface{}) *OrmHmLogonFailures {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmLogonFailures) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmLogonFailures) And(fuc func(orm *OrmHmLogonFailures)) *OrmHmLogonFailures {
	ormAnd := NewOrmHmLogonFailures()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmLogonFailures) Or(fuc func(orm *OrmHmLogonFailures)) *OrmHmLogonFailures {
	ormOr := NewOrmHmLogonFailures()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmLogonFailures) Preload(query string, args ...interface{}) *OrmHmLogonFailures {
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
func (orm *OrmHmLogonFailures) Joins(query string, args ...interface{}) *OrmHmLogonFailures {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmLogonFailures) WhereIpaddress1(val int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress1` = ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereIpaddress1In(val []int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress1` IN ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereIpaddress1Gt(val int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress1` > ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereIpaddress1Gte(val int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress1` >= ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereIpaddress1Lt(val int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress1` < ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereIpaddress1Lte(val int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress1` <= ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereIpaddress2(val int64) *OrmHmLogonFailures {
	orm.db.Where("`ipaddress2` = ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereFailuretime(val database.Time) *OrmHmLogonFailures {
	orm.db.Where("`failuretime` = ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereFailuretimeIn(val []database.Time) *OrmHmLogonFailures {
	orm.db.Where("`failuretime` IN ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereFailuretimeGt(val database.Time) *OrmHmLogonFailures {
	orm.db.Where("`failuretime` > ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereFailuretimeGte(val database.Time) *OrmHmLogonFailures {
	orm.db.Where("`failuretime` >= ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereFailuretimeLt(val database.Time) *OrmHmLogonFailures {
	orm.db.Where("`failuretime` < ?", val)
	return orm
}
func (orm *OrmHmLogonFailures) WhereFailuretimeLte(val database.Time) *OrmHmLogonFailures {
	orm.db.Where("`failuretime` <= ?", val)
	return orm
}

type HmLogonFailuresList []*HmLogonFailures
