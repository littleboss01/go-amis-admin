package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type Smtpfilter struct {
	Id *int32 `gorm:"column:id;type:int(11)" json:"id"` //
}

var (
	SmtpfilterFieldId = "id"
)

func (receiver *Smtpfilter) TableName() string {
	return "smtpfilter"
}

type OrmSmtpfilter struct {
	db *gorm.DB
}

func (orm *OrmSmtpfilter) TableName() string {
	return "smtpfilter"
}

func NewOrmSmtpfilter(txs ...*gorm.DB) *OrmSmtpfilter {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&Smtpfilter{})
	} else {
		tx = txs[0].Model(&Smtpfilter{})
	}
	return &OrmSmtpfilter{db: tx}
}

func (orm *OrmSmtpfilter) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmSmtpfilter) GetTableInfo() interface{} {
	return &Smtpfilter{}
}

// Create insert the value into database
func (orm *OrmSmtpfilter) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmSmtpfilter) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmSmtpfilter) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmSmtpfilter) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmSmtpfilter) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmSmtpfilter) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmSmtpfilter) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmSmtpfilter) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmSmtpfilter) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmSmtpfilter) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmSmtpfilter) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmSmtpfilter) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmSmtpfilter) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmSmtpfilter) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmSmtpfilter) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmSmtpfilter) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmSmtpfilter) Unscoped() *OrmSmtpfilter {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmSmtpfilter) Insert(row *Smtpfilter) error {
	return orm.db.Create(row).Error
}

func (orm *OrmSmtpfilter) Inserts(rows []*Smtpfilter) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmSmtpfilter) Order(value interface{}) *OrmSmtpfilter {
	orm.db.Order(value)
	return orm
}

func (orm *OrmSmtpfilter) Group(name string) *OrmSmtpfilter {
	orm.db.Group(name)
	return orm
}

func (orm *OrmSmtpfilter) Limit(limit int) *OrmSmtpfilter {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmSmtpfilter) Offset(offset int) *OrmSmtpfilter {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmSmtpfilter) Get() SmtpfilterList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmSmtpfilter) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmSmtpfilter) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&Smtpfilter{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmSmtpfilter) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM smtpfilter")
}

func (orm *OrmSmtpfilter) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmSmtpfilter) First(conds ...interface{}) (*Smtpfilter, bool) {
	dest := &Smtpfilter{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmSmtpfilter) Take(conds ...interface{}) (*Smtpfilter, int64) {
	dest := &Smtpfilter{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmSmtpfilter) Last(conds ...interface{}) (*Smtpfilter, int64) {
	dest := &Smtpfilter{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmSmtpfilter) Find(conds ...interface{}) (SmtpfilterList, int64) {
	list := make([]*Smtpfilter, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmSmtpfilter) Paginate(page int, limit int) (SmtpfilterList, int64) {
	var total int64
	list := make([]*Smtpfilter, 0)
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
func (orm *OrmSmtpfilter) SimplePaginate(page int, limit int) SmtpfilterList {
	list := make([]*Smtpfilter, 0)
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
func (orm *OrmSmtpfilter) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmSmtpfilter) FirstOrInit(dest *Smtpfilter, conds ...interface{}) (*Smtpfilter, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmSmtpfilter) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSmtpfilter) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSmtpfilter) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmSmtpfilter) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmSmtpfilter) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmSmtpfilter) Where(query interface{}, args ...interface{}) *OrmSmtpfilter {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmSmtpfilter) Select(query interface{}, args ...interface{}) *OrmSmtpfilter {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmSmtpfilter) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmSmtpfilter) And(fuc func(orm *OrmSmtpfilter)) *OrmSmtpfilter {
	ormAnd := NewOrmSmtpfilter()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmSmtpfilter) Or(fuc func(orm *OrmSmtpfilter)) *OrmSmtpfilter {
	ormOr := NewOrmSmtpfilter()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmSmtpfilter) Preload(query string, args ...interface{}) *OrmSmtpfilter {
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
func (orm *OrmSmtpfilter) Joins(query string, args ...interface{}) *OrmSmtpfilter {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmSmtpfilter) WhereId(val int32) *OrmSmtpfilter {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmSmtpfilter) WhereIdIn(val []int32) *OrmSmtpfilter {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmSmtpfilter) WhereIdNe(val int32) *OrmSmtpfilter {
	orm.db.Where("`id` <> ?", val)
	return orm
}

type SmtpfilterList []*Smtpfilter

func (l SmtpfilterList) GetIdList() []*int32 {
	got := make([]*int32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l SmtpfilterList) GetIdMap() map[*int32]*Smtpfilter {
	got := make(map[*int32]*Smtpfilter)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
