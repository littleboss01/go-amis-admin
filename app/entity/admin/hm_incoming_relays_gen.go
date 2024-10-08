package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmIncomingRelays struct {
	Relayid       int32  `gorm:"column:relayid;autoIncrement;type:int(11);primaryKey;index:relayid,class:BTREE,unique" json:"relayid"` //
	Relayname     string `gorm:"column:relayname;type:varchar(100)" json:"relayname"`                                                  //
	Relaylowerip1 int64  `gorm:"column:relaylowerip1;type:bigint(20)" json:"relaylowerip1"`                                            //
	Relaylowerip2 *int64 `gorm:"column:relaylowerip2;type:bigint(20)" json:"relaylowerip2"`                                            //
	Relayupperip1 int64  `gorm:"column:relayupperip1;type:bigint(20)" json:"relayupperip1"`                                            //
	Relayupperip2 *int64 `gorm:"column:relayupperip2;type:bigint(20)" json:"relayupperip2"`                                            //
}

var (
	HmIncomingRelaysFieldRelayid       = "relayid"
	HmIncomingRelaysFieldRelayname     = "relayname"
	HmIncomingRelaysFieldRelaylowerip1 = "relaylowerip1"
	HmIncomingRelaysFieldRelaylowerip2 = "relaylowerip2"
	HmIncomingRelaysFieldRelayupperip1 = "relayupperip1"
	HmIncomingRelaysFieldRelayupperip2 = "relayupperip2"
)

func (receiver *HmIncomingRelays) TableName() string {
	return "hm_incoming_relays"
}

type OrmHmIncomingRelays struct {
	db *gorm.DB
}

func (orm *OrmHmIncomingRelays) TableName() string {
	return "hm_incoming_relays"
}

func NewOrmHmIncomingRelays(txs ...*gorm.DB) *OrmHmIncomingRelays {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmIncomingRelays{})
	} else {
		tx = txs[0].Model(&HmIncomingRelays{})
	}
	return &OrmHmIncomingRelays{db: tx}
}

func (orm *OrmHmIncomingRelays) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmIncomingRelays) GetTableInfo() interface{} {
	return &HmIncomingRelays{}
}

// Create insert the value into database
func (orm *OrmHmIncomingRelays) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmIncomingRelays) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmIncomingRelays) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmIncomingRelays) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmIncomingRelays) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmIncomingRelays) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmIncomingRelays) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmIncomingRelays) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmIncomingRelays) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmIncomingRelays) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmIncomingRelays) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmIncomingRelays) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmIncomingRelays) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmIncomingRelays) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmIncomingRelays) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmIncomingRelays) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmIncomingRelays) Unscoped() *OrmHmIncomingRelays {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmIncomingRelays) Insert(row *HmIncomingRelays) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmIncomingRelays) Inserts(rows []*HmIncomingRelays) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmIncomingRelays) Order(value interface{}) *OrmHmIncomingRelays {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmIncomingRelays) Group(name string) *OrmHmIncomingRelays {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmIncomingRelays) Limit(limit int) *OrmHmIncomingRelays {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmIncomingRelays) Offset(offset int) *OrmHmIncomingRelays {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmIncomingRelays) Get() HmIncomingRelaysList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmIncomingRelays) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmIncomingRelays) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmIncomingRelays{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmIncomingRelays) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_incoming_relays")
}

func (orm *OrmHmIncomingRelays) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmIncomingRelays) First(conds ...interface{}) (*HmIncomingRelays, bool) {
	dest := &HmIncomingRelays{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmIncomingRelays) Take(conds ...interface{}) (*HmIncomingRelays, int64) {
	dest := &HmIncomingRelays{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmIncomingRelays) Last(conds ...interface{}) (*HmIncomingRelays, int64) {
	dest := &HmIncomingRelays{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmIncomingRelays) Find(conds ...interface{}) (HmIncomingRelaysList, int64) {
	list := make([]*HmIncomingRelays, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmIncomingRelays) Paginate(page int, limit int) (HmIncomingRelaysList, int64) {
	var total int64
	list := make([]*HmIncomingRelays, 0)
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
func (orm *OrmHmIncomingRelays) SimplePaginate(page int, limit int) HmIncomingRelaysList {
	list := make([]*HmIncomingRelays, 0)
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
func (orm *OrmHmIncomingRelays) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmIncomingRelays) FirstOrInit(dest *HmIncomingRelays, conds ...interface{}) (*HmIncomingRelays, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmIncomingRelays) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmIncomingRelays) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmIncomingRelays) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmIncomingRelays) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmIncomingRelays) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmIncomingRelays) Where(query interface{}, args ...interface{}) *OrmHmIncomingRelays {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmIncomingRelays) Select(query interface{}, args ...interface{}) *OrmHmIncomingRelays {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmIncomingRelays) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmIncomingRelays) And(fuc func(orm *OrmHmIncomingRelays)) *OrmHmIncomingRelays {
	ormAnd := NewOrmHmIncomingRelays()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmIncomingRelays) Or(fuc func(orm *OrmHmIncomingRelays)) *OrmHmIncomingRelays {
	ormOr := NewOrmHmIncomingRelays()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmIncomingRelays) Preload(query string, args ...interface{}) *OrmHmIncomingRelays {
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
func (orm *OrmHmIncomingRelays) Joins(query string, args ...interface{}) *OrmHmIncomingRelays {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmIncomingRelays) WhereRelayid(val int32) *OrmHmIncomingRelays {
	orm.db.Where("`relayid` = ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) InsertGetRelayid(row *HmIncomingRelays) int32 {
	orm.db.Create(row)
	return row.Relayid
}
func (orm *OrmHmIncomingRelays) WhereRelayidIn(val []int32) *OrmHmIncomingRelays {
	orm.db.Where("`relayid` IN ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayidGt(val int32) *OrmHmIncomingRelays {
	orm.db.Where("`relayid` > ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayidGte(val int32) *OrmHmIncomingRelays {
	orm.db.Where("`relayid` >= ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayidLt(val int32) *OrmHmIncomingRelays {
	orm.db.Where("`relayid` < ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayidLte(val int32) *OrmHmIncomingRelays {
	orm.db.Where("`relayid` <= ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayname(val string) *OrmHmIncomingRelays {
	orm.db.Where("`relayname` = ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelaylowerip1(val int64) *OrmHmIncomingRelays {
	orm.db.Where("`relaylowerip1` = ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelaylowerip2(val int64) *OrmHmIncomingRelays {
	orm.db.Where("`relaylowerip2` = ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayupperip1(val int64) *OrmHmIncomingRelays {
	orm.db.Where("`relayupperip1` = ?", val)
	return orm
}
func (orm *OrmHmIncomingRelays) WhereRelayupperip2(val int64) *OrmHmIncomingRelays {
	orm.db.Where("`relayupperip2` = ?", val)
	return orm
}

type HmIncomingRelaysList []*HmIncomingRelays

func (l HmIncomingRelaysList) GetRelayidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Relayid)
	}
	return got
}
func (l HmIncomingRelaysList) GetRelayidMap() map[int32]*HmIncomingRelays {
	got := make(map[int32]*HmIncomingRelays)
	for _, val := range l {
		got[val.Relayid] = val
	}
	return got
}
