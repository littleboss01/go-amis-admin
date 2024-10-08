package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type AccountListCopy struct {
	Id          uint32 `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey;comment:'id'" json:"id"` // id
	Mailuser    string `gorm:"column:mailuser;type:varchar(50)" json:"mailuser"`                                //
	Mark        string `gorm:"column:mark;type:varchar(100)" json:"mark"`                                       //
	MsgMaxCount int64  `gorm:"column:msgMaxCount;type:bigint(10)" json:"msgMaxCount"`                           //
	Language    string `gorm:"column:language;type:varchar(20)" json:"language"`                                //
}

var (
	AccountListCopyFieldId          = "id"
	AccountListCopyFieldMailuser    = "mailuser"
	AccountListCopyFieldMark        = "mark"
	AccountListCopyFieldMsgMaxCount = "msgMaxCount"
	AccountListCopyFieldLanguage    = "language"
)

func (receiver *AccountListCopy) TableName() string {
	return "account_list_copy"
}

type OrmAccountListCopy struct {
	db *gorm.DB
}

func (orm *OrmAccountListCopy) TableName() string {
	return "account_list_copy"
}

func NewOrmAccountListCopy(txs ...*gorm.DB) *OrmAccountListCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AccountListCopy{})
	} else {
		tx = txs[0].Model(&AccountListCopy{})
	}
	return &OrmAccountListCopy{db: tx}
}

func (orm *OrmAccountListCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAccountListCopy) GetTableInfo() interface{} {
	return &AccountListCopy{}
}

// Create insert the value into database
func (orm *OrmAccountListCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAccountListCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAccountListCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAccountListCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAccountListCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAccountListCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAccountListCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAccountListCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAccountListCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAccountListCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAccountListCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAccountListCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAccountListCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAccountListCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAccountListCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAccountListCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAccountListCopy) Unscoped() *OrmAccountListCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAccountListCopy) Insert(row *AccountListCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAccountListCopy) Inserts(rows []*AccountListCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAccountListCopy) Order(value interface{}) *OrmAccountListCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAccountListCopy) Group(name string) *OrmAccountListCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAccountListCopy) Limit(limit int) *OrmAccountListCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAccountListCopy) Offset(offset int) *OrmAccountListCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAccountListCopy) Get() AccountListCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAccountListCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAccountListCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AccountListCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAccountListCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM account_list_copy")
}

func (orm *OrmAccountListCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAccountListCopy) First(conds ...interface{}) (*AccountListCopy, bool) {
	dest := &AccountListCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAccountListCopy) Take(conds ...interface{}) (*AccountListCopy, int64) {
	dest := &AccountListCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAccountListCopy) Last(conds ...interface{}) (*AccountListCopy, int64) {
	dest := &AccountListCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAccountListCopy) Find(conds ...interface{}) (AccountListCopyList, int64) {
	list := make([]*AccountListCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAccountListCopy) Paginate(page int, limit int) (AccountListCopyList, int64) {
	var total int64
	list := make([]*AccountListCopy, 0)
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
func (orm *OrmAccountListCopy) SimplePaginate(page int, limit int) AccountListCopyList {
	list := make([]*AccountListCopy, 0)
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
func (orm *OrmAccountListCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAccountListCopy) FirstOrInit(dest *AccountListCopy, conds ...interface{}) (*AccountListCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAccountListCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAccountListCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAccountListCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAccountListCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAccountListCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAccountListCopy) Where(query interface{}, args ...interface{}) *OrmAccountListCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAccountListCopy) Select(query interface{}, args ...interface{}) *OrmAccountListCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAccountListCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAccountListCopy) And(fuc func(orm *OrmAccountListCopy)) *OrmAccountListCopy {
	ormAnd := NewOrmAccountListCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAccountListCopy) Or(fuc func(orm *OrmAccountListCopy)) *OrmAccountListCopy {
	ormOr := NewOrmAccountListCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAccountListCopy) Preload(query string, args ...interface{}) *OrmAccountListCopy {
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
func (orm *OrmAccountListCopy) Joins(query string, args ...interface{}) *OrmAccountListCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAccountListCopy) WhereId(val uint32) *OrmAccountListCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAccountListCopy) InsertGetId(row *AccountListCopy) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAccountListCopy) WhereIdIn(val []uint32) *OrmAccountListCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereIdGt(val uint32) *OrmAccountListCopy {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereIdGte(val uint32) *OrmAccountListCopy {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereIdLt(val uint32) *OrmAccountListCopy {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereIdLte(val uint32) *OrmAccountListCopy {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereMailuser(val string) *OrmAccountListCopy {
	orm.db.Where("`mailuser` = ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereMark(val string) *OrmAccountListCopy {
	orm.db.Where("`mark` = ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereMsgMaxCount(val int64) *OrmAccountListCopy {
	orm.db.Where("`msgMaxCount` = ?", val)
	return orm
}
func (orm *OrmAccountListCopy) WhereLanguage(val string) *OrmAccountListCopy {
	orm.db.Where("`language` = ?", val)
	return orm
}

type AccountListCopyList []*AccountListCopy

func (l AccountListCopyList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AccountListCopyList) GetIdMap() map[uint32]*AccountListCopy {
	got := make(map[uint32]*AccountListCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
