package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type MailRulesCopy struct {
	Id     int32   `gorm:"column:id;type:int(10)" json:"id"`              //
	Str    string  `gorm:"column:str;type:varchar(100)" json:"str"`       //
	Mailto string  `gorm:"column:mailto;type:varchar(100)" json:"mailto"` //
	Switch int32   `gorm:"column:switch;type:int(1)" json:"switch"`       //
	Mark   *string `gorm:"column:mark;type:varchar(100)" json:"mark"`     //
}

var (
	MailRulesCopyFieldId     = "id"
	MailRulesCopyFieldStr    = "str"
	MailRulesCopyFieldMailto = "mailto"
	MailRulesCopyFieldSwitch = "switch"
	MailRulesCopyFieldMark   = "mark"
)

func (receiver *MailRulesCopy) TableName() string {
	return "mail_rules_copy"
}

type OrmMailRulesCopy struct {
	db *gorm.DB
}

func (orm *OrmMailRulesCopy) TableName() string {
	return "mail_rules_copy"
}

func NewOrmMailRulesCopy(txs ...*gorm.DB) *OrmMailRulesCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&MailRulesCopy{})
	} else {
		tx = txs[0].Model(&MailRulesCopy{})
	}
	return &OrmMailRulesCopy{db: tx}
}

func (orm *OrmMailRulesCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmMailRulesCopy) GetTableInfo() interface{} {
	return &MailRulesCopy{}
}

// Create insert the value into database
func (orm *OrmMailRulesCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmMailRulesCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmMailRulesCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmMailRulesCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmMailRulesCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmMailRulesCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmMailRulesCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmMailRulesCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmMailRulesCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmMailRulesCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmMailRulesCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmMailRulesCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmMailRulesCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmMailRulesCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmMailRulesCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmMailRulesCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmMailRulesCopy) Unscoped() *OrmMailRulesCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmMailRulesCopy) Insert(row *MailRulesCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmMailRulesCopy) Inserts(rows []*MailRulesCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmMailRulesCopy) Order(value interface{}) *OrmMailRulesCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmMailRulesCopy) Group(name string) *OrmMailRulesCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmMailRulesCopy) Limit(limit int) *OrmMailRulesCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmMailRulesCopy) Offset(offset int) *OrmMailRulesCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmMailRulesCopy) Get() MailRulesCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmMailRulesCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmMailRulesCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&MailRulesCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmMailRulesCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM mail_rules_copy")
}

func (orm *OrmMailRulesCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmMailRulesCopy) First(conds ...interface{}) (*MailRulesCopy, bool) {
	dest := &MailRulesCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmMailRulesCopy) Take(conds ...interface{}) (*MailRulesCopy, int64) {
	dest := &MailRulesCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmMailRulesCopy) Last(conds ...interface{}) (*MailRulesCopy, int64) {
	dest := &MailRulesCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmMailRulesCopy) Find(conds ...interface{}) (MailRulesCopyList, int64) {
	list := make([]*MailRulesCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmMailRulesCopy) Paginate(page int, limit int) (MailRulesCopyList, int64) {
	var total int64
	list := make([]*MailRulesCopy, 0)
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
func (orm *OrmMailRulesCopy) SimplePaginate(page int, limit int) MailRulesCopyList {
	list := make([]*MailRulesCopy, 0)
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
func (orm *OrmMailRulesCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmMailRulesCopy) FirstOrInit(dest *MailRulesCopy, conds ...interface{}) (*MailRulesCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmMailRulesCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMailRulesCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMailRulesCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmMailRulesCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmMailRulesCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmMailRulesCopy) Where(query interface{}, args ...interface{}) *OrmMailRulesCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmMailRulesCopy) Select(query interface{}, args ...interface{}) *OrmMailRulesCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmMailRulesCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmMailRulesCopy) And(fuc func(orm *OrmMailRulesCopy)) *OrmMailRulesCopy {
	ormAnd := NewOrmMailRulesCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmMailRulesCopy) Or(fuc func(orm *OrmMailRulesCopy)) *OrmMailRulesCopy {
	ormOr := NewOrmMailRulesCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmMailRulesCopy) Preload(query string, args ...interface{}) *OrmMailRulesCopy {
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
func (orm *OrmMailRulesCopy) Joins(query string, args ...interface{}) *OrmMailRulesCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmMailRulesCopy) WhereId(val int32) *OrmMailRulesCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmMailRulesCopy) WhereIdIn(val []int32) *OrmMailRulesCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmMailRulesCopy) WhereIdNe(val int32) *OrmMailRulesCopy {
	orm.db.Where("`id` <> ?", val)
	return orm
}
func (orm *OrmMailRulesCopy) WhereStr(val string) *OrmMailRulesCopy {
	orm.db.Where("`str` = ?", val)
	return orm
}
func (orm *OrmMailRulesCopy) WhereMailto(val string) *OrmMailRulesCopy {
	orm.db.Where("`mailto` = ?", val)
	return orm
}
func (orm *OrmMailRulesCopy) WhereSwitch(val int32) *OrmMailRulesCopy {
	orm.db.Where("`switch` = ?", val)
	return orm
}
func (orm *OrmMailRulesCopy) WhereMark(val string) *OrmMailRulesCopy {
	orm.db.Where("`mark` = ?", val)
	return orm
}

type MailRulesCopyList []*MailRulesCopy

func (l MailRulesCopyList) GetIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l MailRulesCopyList) GetIdMap() map[int32]*MailRulesCopy {
	got := make(map[int32]*MailRulesCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
