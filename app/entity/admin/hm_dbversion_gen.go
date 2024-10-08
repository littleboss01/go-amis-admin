package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmDbversion struct {
	Value int32 `gorm:"column:value;type:int(11)" json:"value"` //
}

var (
	HmDbversionFieldValue = "value"
)

func (receiver *HmDbversion) TableName() string {
	return "hm_dbversion"
}

type OrmHmDbversion struct {
	db *gorm.DB
}

func (orm *OrmHmDbversion) TableName() string {
	return "hm_dbversion"
}

func NewOrmHmDbversion(txs ...*gorm.DB) *OrmHmDbversion {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmDbversion{})
	} else {
		tx = txs[0].Model(&HmDbversion{})
	}
	return &OrmHmDbversion{db: tx}
}

func (orm *OrmHmDbversion) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmDbversion) GetTableInfo() interface{} {
	return &HmDbversion{}
}

// Create insert the value into database
func (orm *OrmHmDbversion) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmDbversion) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmDbversion) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmDbversion) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmDbversion) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmDbversion) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmDbversion) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmDbversion) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmDbversion) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmDbversion) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmDbversion) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmDbversion) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmDbversion) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmDbversion) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmDbversion) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmDbversion) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmDbversion) Unscoped() *OrmHmDbversion {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmDbversion) Insert(row *HmDbversion) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmDbversion) Inserts(rows []*HmDbversion) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmDbversion) Order(value interface{}) *OrmHmDbversion {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmDbversion) Group(name string) *OrmHmDbversion {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmDbversion) Limit(limit int) *OrmHmDbversion {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmDbversion) Offset(offset int) *OrmHmDbversion {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmDbversion) Get() HmDbversionList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmDbversion) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmDbversion) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmDbversion{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmDbversion) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_dbversion")
}

func (orm *OrmHmDbversion) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmDbversion) First(conds ...interface{}) (*HmDbversion, bool) {
	dest := &HmDbversion{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmDbversion) Take(conds ...interface{}) (*HmDbversion, int64) {
	dest := &HmDbversion{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmDbversion) Last(conds ...interface{}) (*HmDbversion, int64) {
	dest := &HmDbversion{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmDbversion) Find(conds ...interface{}) (HmDbversionList, int64) {
	list := make([]*HmDbversion, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmDbversion) Paginate(page int, limit int) (HmDbversionList, int64) {
	var total int64
	list := make([]*HmDbversion, 0)
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
func (orm *OrmHmDbversion) SimplePaginate(page int, limit int) HmDbversionList {
	list := make([]*HmDbversion, 0)
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
func (orm *OrmHmDbversion) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmDbversion) FirstOrInit(dest *HmDbversion, conds ...interface{}) (*HmDbversion, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmDbversion) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDbversion) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDbversion) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmDbversion) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmDbversion) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmDbversion) Where(query interface{}, args ...interface{}) *OrmHmDbversion {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmDbversion) Select(query interface{}, args ...interface{}) *OrmHmDbversion {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmDbversion) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmDbversion) And(fuc func(orm *OrmHmDbversion)) *OrmHmDbversion {
	ormAnd := NewOrmHmDbversion()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmDbversion) Or(fuc func(orm *OrmHmDbversion)) *OrmHmDbversion {
	ormOr := NewOrmHmDbversion()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmDbversion) Preload(query string, args ...interface{}) *OrmHmDbversion {
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
func (orm *OrmHmDbversion) Joins(query string, args ...interface{}) *OrmHmDbversion {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmDbversion) WhereValue(val int32) *OrmHmDbversion {
	orm.db.Where("`value` = ?", val)
	return orm
}

type HmDbversionList []*HmDbversion
