package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type SpamCopy struct {
	Id *int32 `gorm:"column:id;type:int(11)" json:"id"` //
}

var (
	SpamCopyFieldId = "id"
)

func (receiver *SpamCopy) TableName() string {
	return "spam_copy"
}

type OrmSpamCopy struct {
	db *gorm.DB
}

func (orm *OrmSpamCopy) TableName() string {
	return "spam_copy"
}

func NewOrmSpamCopy(txs ...*gorm.DB) *OrmSpamCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&SpamCopy{})
	} else {
		tx = txs[0].Model(&SpamCopy{})
	}
	return &OrmSpamCopy{db: tx}
}

func (orm *OrmSpamCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmSpamCopy) GetTableInfo() interface{} {
	return &SpamCopy{}
}

// Create insert the value into database
func (orm *OrmSpamCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmSpamCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmSpamCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmSpamCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmSpamCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmSpamCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmSpamCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmSpamCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmSpamCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmSpamCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmSpamCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmSpamCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmSpamCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmSpamCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmSpamCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmSpamCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmSpamCopy) Unscoped() *OrmSpamCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmSpamCopy) Insert(row *SpamCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmSpamCopy) Inserts(rows []*SpamCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmSpamCopy) Order(value interface{}) *OrmSpamCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmSpamCopy) Group(name string) *OrmSpamCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmSpamCopy) Limit(limit int) *OrmSpamCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmSpamCopy) Offset(offset int) *OrmSpamCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmSpamCopy) Get() SpamCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmSpamCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmSpamCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&SpamCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmSpamCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM spam_copy")
}

func (orm *OrmSpamCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmSpamCopy) First(conds ...interface{}) (*SpamCopy, bool) {
	dest := &SpamCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmSpamCopy) Take(conds ...interface{}) (*SpamCopy, int64) {
	dest := &SpamCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmSpamCopy) Last(conds ...interface{}) (*SpamCopy, int64) {
	dest := &SpamCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmSpamCopy) Find(conds ...interface{}) (SpamCopyList, int64) {
	list := make([]*SpamCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmSpamCopy) Paginate(page int, limit int) (SpamCopyList, int64) {
	var total int64
	list := make([]*SpamCopy, 0)
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
func (orm *OrmSpamCopy) SimplePaginate(page int, limit int) SpamCopyList {
	list := make([]*SpamCopy, 0)
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
func (orm *OrmSpamCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmSpamCopy) FirstOrInit(dest *SpamCopy, conds ...interface{}) (*SpamCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmSpamCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSpamCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSpamCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmSpamCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmSpamCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmSpamCopy) Where(query interface{}, args ...interface{}) *OrmSpamCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmSpamCopy) Select(query interface{}, args ...interface{}) *OrmSpamCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmSpamCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmSpamCopy) And(fuc func(orm *OrmSpamCopy)) *OrmSpamCopy {
	ormAnd := NewOrmSpamCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmSpamCopy) Or(fuc func(orm *OrmSpamCopy)) *OrmSpamCopy {
	ormOr := NewOrmSpamCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmSpamCopy) Preload(query string, args ...interface{}) *OrmSpamCopy {
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
func (orm *OrmSpamCopy) Joins(query string, args ...interface{}) *OrmSpamCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmSpamCopy) WhereId(val int32) *OrmSpamCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmSpamCopy) WhereIdIn(val []int32) *OrmSpamCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmSpamCopy) WhereIdNe(val int32) *OrmSpamCopy {
	orm.db.Where("`id` <> ?", val)
	return orm
}

type SpamCopyList []*SpamCopy

func (l SpamCopyList) GetIdList() []*int32 {
	got := make([]*int32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l SpamCopyList) GetIdMap() map[*int32]*SpamCopy {
	got := make(map[*int32]*SpamCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
