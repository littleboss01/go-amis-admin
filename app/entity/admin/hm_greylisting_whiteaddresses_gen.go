package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmGreylistingWhiteaddresses struct {
	Whiteid            int64  `gorm:"column:whiteid;autoIncrement;type:bigint(20);primaryKey;index:whiteid,class:BTREE,unique" json:"whiteid"` //
	Whiteipaddress     string `gorm:"column:whiteipaddress;type:varchar(255)" json:"whiteipaddress"`                                           //
	Whiteipdescription string `gorm:"column:whiteipdescription;type:varchar(255)" json:"whiteipdescription"`                                   //
}

var (
	HmGreylistingWhiteaddressesFieldWhiteid            = "whiteid"
	HmGreylistingWhiteaddressesFieldWhiteipaddress     = "whiteipaddress"
	HmGreylistingWhiteaddressesFieldWhiteipdescription = "whiteipdescription"
)

func (receiver *HmGreylistingWhiteaddresses) TableName() string {
	return "hm_greylisting_whiteaddresses"
}

type OrmHmGreylistingWhiteaddresses struct {
	db *gorm.DB
}

func (orm *OrmHmGreylistingWhiteaddresses) TableName() string {
	return "hm_greylisting_whiteaddresses"
}

func NewOrmHmGreylistingWhiteaddresses(txs ...*gorm.DB) *OrmHmGreylistingWhiteaddresses {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmGreylistingWhiteaddresses{})
	} else {
		tx = txs[0].Model(&HmGreylistingWhiteaddresses{})
	}
	return &OrmHmGreylistingWhiteaddresses{db: tx}
}

func (orm *OrmHmGreylistingWhiteaddresses) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmGreylistingWhiteaddresses) GetTableInfo() interface{} {
	return &HmGreylistingWhiteaddresses{}
}

// Create insert the value into database
func (orm *OrmHmGreylistingWhiteaddresses) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmGreylistingWhiteaddresses) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmGreylistingWhiteaddresses) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmGreylistingWhiteaddresses) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmGreylistingWhiteaddresses) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmGreylistingWhiteaddresses) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmGreylistingWhiteaddresses) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmGreylistingWhiteaddresses) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmGreylistingWhiteaddresses) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmGreylistingWhiteaddresses) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmGreylistingWhiteaddresses) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmGreylistingWhiteaddresses) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmGreylistingWhiteaddresses) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmGreylistingWhiteaddresses) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmGreylistingWhiteaddresses) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmGreylistingWhiteaddresses) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmGreylistingWhiteaddresses) Unscoped() *OrmHmGreylistingWhiteaddresses {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmGreylistingWhiteaddresses) Insert(row *HmGreylistingWhiteaddresses) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmGreylistingWhiteaddresses) Inserts(rows []*HmGreylistingWhiteaddresses) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmGreylistingWhiteaddresses) Order(value interface{}) *OrmHmGreylistingWhiteaddresses {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) Group(name string) *OrmHmGreylistingWhiteaddresses {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) Limit(limit int) *OrmHmGreylistingWhiteaddresses {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) Offset(offset int) *OrmHmGreylistingWhiteaddresses {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmGreylistingWhiteaddresses) Get() HmGreylistingWhiteaddressesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmGreylistingWhiteaddresses) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmGreylistingWhiteaddresses) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmGreylistingWhiteaddresses{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmGreylistingWhiteaddresses) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_greylisting_whiteaddresses")
}

func (orm *OrmHmGreylistingWhiteaddresses) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmGreylistingWhiteaddresses) First(conds ...interface{}) (*HmGreylistingWhiteaddresses, bool) {
	dest := &HmGreylistingWhiteaddresses{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmGreylistingWhiteaddresses) Take(conds ...interface{}) (*HmGreylistingWhiteaddresses, int64) {
	dest := &HmGreylistingWhiteaddresses{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmGreylistingWhiteaddresses) Last(conds ...interface{}) (*HmGreylistingWhiteaddresses, int64) {
	dest := &HmGreylistingWhiteaddresses{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmGreylistingWhiteaddresses) Find(conds ...interface{}) (HmGreylistingWhiteaddressesList, int64) {
	list := make([]*HmGreylistingWhiteaddresses, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmGreylistingWhiteaddresses) Paginate(page int, limit int) (HmGreylistingWhiteaddressesList, int64) {
	var total int64
	list := make([]*HmGreylistingWhiteaddresses, 0)
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
func (orm *OrmHmGreylistingWhiteaddresses) SimplePaginate(page int, limit int) HmGreylistingWhiteaddressesList {
	list := make([]*HmGreylistingWhiteaddresses, 0)
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
func (orm *OrmHmGreylistingWhiteaddresses) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmGreylistingWhiteaddresses) FirstOrInit(dest *HmGreylistingWhiteaddresses, conds ...interface{}) (*HmGreylistingWhiteaddresses, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmGreylistingWhiteaddresses) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGreylistingWhiteaddresses) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGreylistingWhiteaddresses) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmGreylistingWhiteaddresses) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmGreylistingWhiteaddresses) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmGreylistingWhiteaddresses) Where(query interface{}, args ...interface{}) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) Select(query interface{}, args ...interface{}) *OrmHmGreylistingWhiteaddresses {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmGreylistingWhiteaddresses) And(fuc func(orm *OrmHmGreylistingWhiteaddresses)) *OrmHmGreylistingWhiteaddresses {
	ormAnd := NewOrmHmGreylistingWhiteaddresses()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) Or(fuc func(orm *OrmHmGreylistingWhiteaddresses)) *OrmHmGreylistingWhiteaddresses {
	ormOr := NewOrmHmGreylistingWhiteaddresses()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmGreylistingWhiteaddresses) Preload(query string, args ...interface{}) *OrmHmGreylistingWhiteaddresses {
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
func (orm *OrmHmGreylistingWhiteaddresses) Joins(query string, args ...interface{}) *OrmHmGreylistingWhiteaddresses {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteid(val int64) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteid` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) InsertGetWhiteid(row *HmGreylistingWhiteaddresses) int64 {
	orm.db.Create(row)
	return row.Whiteid
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteidIn(val []int64) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteid` IN ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteidGt(val int64) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteid` > ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteidGte(val int64) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteid` >= ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteidLt(val int64) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteid` < ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteidLte(val int64) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteid` <= ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteipaddress(val string) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteipaddress` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingWhiteaddresses) WhereWhiteipdescription(val string) *OrmHmGreylistingWhiteaddresses {
	orm.db.Where("`whiteipdescription` = ?", val)
	return orm
}

type HmGreylistingWhiteaddressesList []*HmGreylistingWhiteaddresses

func (l HmGreylistingWhiteaddressesList) GetWhiteidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Whiteid)
	}
	return got
}
func (l HmGreylistingWhiteaddressesList) GetWhiteidMap() map[int64]*HmGreylistingWhiteaddresses {
	got := make(map[int64]*HmGreylistingWhiteaddresses)
	for _, val := range l {
		got[val.Whiteid] = val
	}
	return got
}
