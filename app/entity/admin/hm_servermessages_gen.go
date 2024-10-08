package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmServermessages struct {
	Smid   int32  `gorm:"column:smid;autoIncrement;type:int(11);primaryKey;index:smid,class:BTREE,unique" json:"smid"` //
	Smname string `gorm:"column:smname;type:varchar(255)" json:"smname"`                                               //
	Smtext string `gorm:"column:smtext;type:text" json:"smtext"`                                                       //
}

var (
	HmServermessagesFieldSmid   = "smid"
	HmServermessagesFieldSmname = "smname"
	HmServermessagesFieldSmtext = "smtext"
)

func (receiver *HmServermessages) TableName() string {
	return "hm_servermessages"
}

type OrmHmServermessages struct {
	db *gorm.DB
}

func (orm *OrmHmServermessages) TableName() string {
	return "hm_servermessages"
}

func NewOrmHmServermessages(txs ...*gorm.DB) *OrmHmServermessages {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmServermessages{})
	} else {
		tx = txs[0].Model(&HmServermessages{})
	}
	return &OrmHmServermessages{db: tx}
}

func (orm *OrmHmServermessages) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmServermessages) GetTableInfo() interface{} {
	return &HmServermessages{}
}

// Create insert the value into database
func (orm *OrmHmServermessages) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmServermessages) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmServermessages) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmServermessages) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmServermessages) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmServermessages) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmServermessages) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmServermessages) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmServermessages) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmServermessages) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmServermessages) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmServermessages) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmServermessages) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmServermessages) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmServermessages) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmServermessages) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmServermessages) Unscoped() *OrmHmServermessages {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmServermessages) Insert(row *HmServermessages) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmServermessages) Inserts(rows []*HmServermessages) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmServermessages) Order(value interface{}) *OrmHmServermessages {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmServermessages) Group(name string) *OrmHmServermessages {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmServermessages) Limit(limit int) *OrmHmServermessages {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmServermessages) Offset(offset int) *OrmHmServermessages {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmServermessages) Get() HmServermessagesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmServermessages) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmServermessages) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmServermessages{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmServermessages) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_servermessages")
}

func (orm *OrmHmServermessages) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmServermessages) First(conds ...interface{}) (*HmServermessages, bool) {
	dest := &HmServermessages{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmServermessages) Take(conds ...interface{}) (*HmServermessages, int64) {
	dest := &HmServermessages{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmServermessages) Last(conds ...interface{}) (*HmServermessages, int64) {
	dest := &HmServermessages{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmServermessages) Find(conds ...interface{}) (HmServermessagesList, int64) {
	list := make([]*HmServermessages, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmServermessages) Paginate(page int, limit int) (HmServermessagesList, int64) {
	var total int64
	list := make([]*HmServermessages, 0)
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
func (orm *OrmHmServermessages) SimplePaginate(page int, limit int) HmServermessagesList {
	list := make([]*HmServermessages, 0)
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
func (orm *OrmHmServermessages) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmServermessages) FirstOrInit(dest *HmServermessages, conds ...interface{}) (*HmServermessages, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmServermessages) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmServermessages) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmServermessages) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmServermessages) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmServermessages) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmServermessages) Where(query interface{}, args ...interface{}) *OrmHmServermessages {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmServermessages) Select(query interface{}, args ...interface{}) *OrmHmServermessages {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmServermessages) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmServermessages) And(fuc func(orm *OrmHmServermessages)) *OrmHmServermessages {
	ormAnd := NewOrmHmServermessages()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmServermessages) Or(fuc func(orm *OrmHmServermessages)) *OrmHmServermessages {
	ormOr := NewOrmHmServermessages()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmServermessages) Preload(query string, args ...interface{}) *OrmHmServermessages {
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
func (orm *OrmHmServermessages) Joins(query string, args ...interface{}) *OrmHmServermessages {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmServermessages) WhereSmid(val int32) *OrmHmServermessages {
	orm.db.Where("`smid` = ?", val)
	return orm
}
func (orm *OrmHmServermessages) InsertGetSmid(row *HmServermessages) int32 {
	orm.db.Create(row)
	return row.Smid
}
func (orm *OrmHmServermessages) WhereSmidIn(val []int32) *OrmHmServermessages {
	orm.db.Where("`smid` IN ?", val)
	return orm
}
func (orm *OrmHmServermessages) WhereSmidGt(val int32) *OrmHmServermessages {
	orm.db.Where("`smid` > ?", val)
	return orm
}
func (orm *OrmHmServermessages) WhereSmidGte(val int32) *OrmHmServermessages {
	orm.db.Where("`smid` >= ?", val)
	return orm
}
func (orm *OrmHmServermessages) WhereSmidLt(val int32) *OrmHmServermessages {
	orm.db.Where("`smid` < ?", val)
	return orm
}
func (orm *OrmHmServermessages) WhereSmidLte(val int32) *OrmHmServermessages {
	orm.db.Where("`smid` <= ?", val)
	return orm
}
func (orm *OrmHmServermessages) WhereSmname(val string) *OrmHmServermessages {
	orm.db.Where("`smname` = ?", val)
	return orm
}
func (orm *OrmHmServermessages) WhereSmtext(val string) *OrmHmServermessages {
	orm.db.Where("`smtext` = ?", val)
	return orm
}

type HmServermessagesList []*HmServermessages

func (l HmServermessagesList) GetSmidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Smid)
	}
	return got
}
func (l HmServermessagesList) GetSmidMap() map[int32]*HmServermessages {
	got := make(map[int32]*HmServermessages)
	for _, val := range l {
		got[val.Smid] = val
	}
	return got
}
