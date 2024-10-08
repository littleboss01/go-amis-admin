package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type AdminAdsfCopy struct {
	Id   *int32 `gorm:"column:id;type:int(11)" json:"id"`     //
	Adsf *int32 `gorm:"column:adsf;type:int(11)" json:"adsf"` //
}

var (
	AdminAdsfCopyFieldId   = "id"
	AdminAdsfCopyFieldAdsf = "adsf"
)

func (receiver *AdminAdsfCopy) TableName() string {
	return "admin_adsf_copy"
}

type OrmAdminAdsfCopy struct {
	db *gorm.DB
}

func (orm *OrmAdminAdsfCopy) TableName() string {
	return "admin_adsf_copy"
}

func NewOrmAdminAdsfCopy(txs ...*gorm.DB) *OrmAdminAdsfCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminAdsfCopy{})
	} else {
		tx = txs[0].Model(&AdminAdsfCopy{})
	}
	return &OrmAdminAdsfCopy{db: tx}
}

func (orm *OrmAdminAdsfCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminAdsfCopy) GetTableInfo() interface{} {
	return &AdminAdsfCopy{}
}

// Create insert the value into database
func (orm *OrmAdminAdsfCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminAdsfCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminAdsfCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminAdsfCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminAdsfCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminAdsfCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminAdsfCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminAdsfCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminAdsfCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminAdsfCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminAdsfCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminAdsfCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminAdsfCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminAdsfCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminAdsfCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminAdsfCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminAdsfCopy) Unscoped() *OrmAdminAdsfCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminAdsfCopy) Insert(row *AdminAdsfCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminAdsfCopy) Inserts(rows []*AdminAdsfCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminAdsfCopy) Order(value interface{}) *OrmAdminAdsfCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminAdsfCopy) Group(name string) *OrmAdminAdsfCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminAdsfCopy) Limit(limit int) *OrmAdminAdsfCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminAdsfCopy) Offset(offset int) *OrmAdminAdsfCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminAdsfCopy) Get() AdminAdsfCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminAdsfCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminAdsfCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminAdsfCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminAdsfCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_adsf_copy")
}

func (orm *OrmAdminAdsfCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminAdsfCopy) First(conds ...interface{}) (*AdminAdsfCopy, bool) {
	dest := &AdminAdsfCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminAdsfCopy) Take(conds ...interface{}) (*AdminAdsfCopy, int64) {
	dest := &AdminAdsfCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminAdsfCopy) Last(conds ...interface{}) (*AdminAdsfCopy, int64) {
	dest := &AdminAdsfCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminAdsfCopy) Find(conds ...interface{}) (AdminAdsfCopyList, int64) {
	list := make([]*AdminAdsfCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminAdsfCopy) Paginate(page int, limit int) (AdminAdsfCopyList, int64) {
	var total int64
	list := make([]*AdminAdsfCopy, 0)
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
func (orm *OrmAdminAdsfCopy) SimplePaginate(page int, limit int) AdminAdsfCopyList {
	list := make([]*AdminAdsfCopy, 0)
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
func (orm *OrmAdminAdsfCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminAdsfCopy) FirstOrInit(dest *AdminAdsfCopy, conds ...interface{}) (*AdminAdsfCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminAdsfCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminAdsfCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminAdsfCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminAdsfCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminAdsfCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminAdsfCopy) Where(query interface{}, args ...interface{}) *OrmAdminAdsfCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminAdsfCopy) Select(query interface{}, args ...interface{}) *OrmAdminAdsfCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminAdsfCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminAdsfCopy) And(fuc func(orm *OrmAdminAdsfCopy)) *OrmAdminAdsfCopy {
	ormAnd := NewOrmAdminAdsfCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminAdsfCopy) Or(fuc func(orm *OrmAdminAdsfCopy)) *OrmAdminAdsfCopy {
	ormOr := NewOrmAdminAdsfCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminAdsfCopy) Preload(query string, args ...interface{}) *OrmAdminAdsfCopy {
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
func (orm *OrmAdminAdsfCopy) Joins(query string, args ...interface{}) *OrmAdminAdsfCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminAdsfCopy) WhereId(val int32) *OrmAdminAdsfCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminAdsfCopy) WhereIdIn(val []int32) *OrmAdminAdsfCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminAdsfCopy) WhereIdNe(val int32) *OrmAdminAdsfCopy {
	orm.db.Where("`id` <> ?", val)
	return orm
}
func (orm *OrmAdminAdsfCopy) WhereAdsf(val int32) *OrmAdminAdsfCopy {
	orm.db.Where("`adsf` = ?", val)
	return orm
}

type AdminAdsfCopyList []*AdminAdsfCopy

func (l AdminAdsfCopyList) GetIdList() []*int32 {
	got := make([]*int32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminAdsfCopyList) GetIdMap() map[*int32]*AdminAdsfCopy {
	got := make(map[*int32]*AdminAdsfCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
