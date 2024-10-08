package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type MailRules struct {
	Id     int32   `gorm:"column:id;type:int(10)" json:"id"`              //
	Str    string  `gorm:"column:str;type:varchar(100)" json:"str"`       //
	Mailto string  `gorm:"column:mailto;type:varchar(100)" json:"mailto"` //
	Switch int32   `gorm:"column:switch;type:int(1)" json:"switch"`       //
	Mark   *string `gorm:"column:mark;type:varchar(100)" json:"mark"`     //
}

var (
	MailRulesFieldId     = "id"
	MailRulesFieldStr    = "str"
	MailRulesFieldMailto = "mailto"
	MailRulesFieldSwitch = "switch"
	MailRulesFieldMark   = "mark"
)

func (receiver *MailRules) TableName() string {
	return "mail_rules"
}

type OrmMailRules struct {
	db *gorm.DB
}

func (orm *OrmMailRules) TableName() string {
	return "mail_rules"
}

func NewOrmMailRules(txs ...*gorm.DB) *OrmMailRules {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&MailRules{})
	} else {
		tx = txs[0].Model(&MailRules{})
	}
	return &OrmMailRules{db: tx}
}

func (orm *OrmMailRules) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmMailRules) GetTableInfo() interface{} {
	return &MailRules{}
}

// Create insert the value into database
func (orm *OrmMailRules) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmMailRules) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmMailRules) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmMailRules) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmMailRules) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmMailRules) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmMailRules) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmMailRules) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmMailRules) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmMailRules) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmMailRules) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmMailRules) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmMailRules) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmMailRules) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmMailRules) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmMailRules) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmMailRules) Unscoped() *OrmMailRules {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmMailRules) Insert(row *MailRules) error {
	return orm.db.Create(row).Error
}

func (orm *OrmMailRules) Inserts(rows []*MailRules) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmMailRules) Order(value interface{}) *OrmMailRules {
	orm.db.Order(value)
	return orm
}

func (orm *OrmMailRules) Group(name string) *OrmMailRules {
	orm.db.Group(name)
	return orm
}

func (orm *OrmMailRules) Limit(limit int) *OrmMailRules {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmMailRules) Offset(offset int) *OrmMailRules {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmMailRules) Get() MailRulesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmMailRules) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmMailRules) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&MailRules{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmMailRules) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM mail_rules")
}

func (orm *OrmMailRules) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmMailRules) First(conds ...interface{}) (*MailRules, bool) {
	dest := &MailRules{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmMailRules) Take(conds ...interface{}) (*MailRules, int64) {
	dest := &MailRules{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmMailRules) Last(conds ...interface{}) (*MailRules, int64) {
	dest := &MailRules{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmMailRules) Find(conds ...interface{}) (MailRulesList, int64) {
	list := make([]*MailRules, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmMailRules) Paginate(page int, limit int) (MailRulesList, int64) {
	var total int64
	list := make([]*MailRules, 0)
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
func (orm *OrmMailRules) SimplePaginate(page int, limit int) MailRulesList {
	list := make([]*MailRules, 0)
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
func (orm *OrmMailRules) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmMailRules) FirstOrInit(dest *MailRules, conds ...interface{}) (*MailRules, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmMailRules) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMailRules) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmMailRules) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmMailRules) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmMailRules) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmMailRules) Where(query interface{}, args ...interface{}) *OrmMailRules {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmMailRules) Select(query interface{}, args ...interface{}) *OrmMailRules {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmMailRules) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmMailRules) And(fuc func(orm *OrmMailRules)) *OrmMailRules {
	ormAnd := NewOrmMailRules()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmMailRules) Or(fuc func(orm *OrmMailRules)) *OrmMailRules {
	ormOr := NewOrmMailRules()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmMailRules) Preload(query string, args ...interface{}) *OrmMailRules {
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
func (orm *OrmMailRules) Joins(query string, args ...interface{}) *OrmMailRules {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmMailRules) WhereId(val int32) *OrmMailRules {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmMailRules) WhereIdIn(val []int32) *OrmMailRules {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmMailRules) WhereIdNe(val int32) *OrmMailRules {
	orm.db.Where("`id` <> ?", val)
	return orm
}
func (orm *OrmMailRules) WhereStr(val string) *OrmMailRules {
	orm.db.Where("`str` = ?", val)
	return orm
}
func (orm *OrmMailRules) WhereMailto(val string) *OrmMailRules {
	orm.db.Where("`mailto` = ?", val)
	return orm
}
func (orm *OrmMailRules) WhereSwitch(val int32) *OrmMailRules {
	orm.db.Where("`switch` = ?", val)
	return orm
}
func (orm *OrmMailRules) WhereMark(val string) *OrmMailRules {
	orm.db.Where("`mark` = ?", val)
	return orm
}

type MailRulesList []*MailRules

func (l MailRulesList) GetIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l MailRulesList) GetIdMap() map[int32]*MailRules {
	got := make(map[int32]*MailRules)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
