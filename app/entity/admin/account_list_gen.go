package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type AccountList struct {
	Id          uint32 `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey;comment:'id'" json:"id"` // id
	Mailuser    string `gorm:"column:mailuser;type:varchar(50)" json:"mailuser"`                                //
	Mark        string `gorm:"column:mark;type:varchar(100)" json:"mark"`                                       //
	MsgMaxCount int64  `gorm:"column:msgMaxCount;type:bigint(10)" json:"msgMaxCount"`                           //
	Language    string `gorm:"column:language;type:varchar(20)" json:"language"`                                //
}

var (
	AccountListFieldId          = "id"
	AccountListFieldMailuser    = "mailuser"
	AccountListFieldMark        = "mark"
	AccountListFieldMsgMaxCount = "msgMaxCount"
	AccountListFieldLanguage    = "language"
)

func (receiver *AccountList) TableName() string {
	return "account_list"
}

type OrmAccountList struct {
	db *gorm.DB
}

func (orm *OrmAccountList) TableName() string {
	return "account_list"
}

func NewOrmAccountList(txs ...*gorm.DB) *OrmAccountList {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AccountList{})
	} else {
		tx = txs[0].Model(&AccountList{})
	}
	return &OrmAccountList{db: tx}
}

func (orm *OrmAccountList) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAccountList) GetTableInfo() interface{} {
	return &AccountList{}
}

// Create insert the value into database
func (orm *OrmAccountList) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAccountList) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAccountList) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAccountList) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAccountList) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAccountList) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAccountList) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAccountList) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAccountList) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAccountList) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAccountList) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAccountList) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAccountList) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAccountList) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAccountList) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAccountList) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAccountList) Unscoped() *OrmAccountList {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAccountList) Insert(row *AccountList) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAccountList) Inserts(rows []*AccountList) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAccountList) Order(value interface{}) *OrmAccountList {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAccountList) Group(name string) *OrmAccountList {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAccountList) Limit(limit int) *OrmAccountList {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAccountList) Offset(offset int) *OrmAccountList {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAccountList) Get() AccountListList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAccountList) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAccountList) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AccountList{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAccountList) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM account_list")
}

func (orm *OrmAccountList) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAccountList) First(conds ...interface{}) (*AccountList, bool) {
	dest := &AccountList{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAccountList) Take(conds ...interface{}) (*AccountList, int64) {
	dest := &AccountList{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAccountList) Last(conds ...interface{}) (*AccountList, int64) {
	dest := &AccountList{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAccountList) Find(conds ...interface{}) (AccountListList, int64) {
	list := make([]*AccountList, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAccountList) Paginate(page int, limit int) (AccountListList, int64) {
	var total int64
	list := make([]*AccountList, 0)
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
func (orm *OrmAccountList) SimplePaginate(page int, limit int) AccountListList {
	list := make([]*AccountList, 0)
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
func (orm *OrmAccountList) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAccountList) FirstOrInit(dest *AccountList, conds ...interface{}) (*AccountList, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAccountList) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAccountList) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAccountList) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAccountList) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAccountList) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAccountList) Where(query interface{}, args ...interface{}) *OrmAccountList {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAccountList) Select(query interface{}, args ...interface{}) *OrmAccountList {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAccountList) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAccountList) And(fuc func(orm *OrmAccountList)) *OrmAccountList {
	ormAnd := NewOrmAccountList()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAccountList) Or(fuc func(orm *OrmAccountList)) *OrmAccountList {
	ormOr := NewOrmAccountList()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAccountList) Preload(query string, args ...interface{}) *OrmAccountList {
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
func (orm *OrmAccountList) Joins(query string, args ...interface{}) *OrmAccountList {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAccountList) WhereId(val uint32) *OrmAccountList {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAccountList) InsertGetId(row *AccountList) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAccountList) WhereIdIn(val []uint32) *OrmAccountList {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAccountList) WhereIdGt(val uint32) *OrmAccountList {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAccountList) WhereIdGte(val uint32) *OrmAccountList {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAccountList) WhereIdLt(val uint32) *OrmAccountList {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAccountList) WhereIdLte(val uint32) *OrmAccountList {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAccountList) WhereMailuser(val string) *OrmAccountList {
	orm.db.Where("`mailuser` = ?", val)
	return orm
}
func (orm *OrmAccountList) WhereMark(val string) *OrmAccountList {
	orm.db.Where("`mark` = ?", val)
	return orm
}
func (orm *OrmAccountList) WhereMsgMaxCount(val int64) *OrmAccountList {
	orm.db.Where("`msgMaxCount` = ?", val)
	return orm
}
func (orm *OrmAccountList) WhereLanguage(val string) *OrmAccountList {
	orm.db.Where("`language` = ?", val)
	return orm
}

type AccountListList []*AccountList

func (l AccountListList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AccountListList) GetIdMap() map[uint32]*AccountList {
	got := make(map[uint32]*AccountList)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
