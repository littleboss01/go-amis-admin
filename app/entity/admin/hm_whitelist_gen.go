package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmWhitelist struct {
	Whiteid              int64  `gorm:"column:whiteid;autoIncrement;type:bigint(20);primaryKey;index:whiteid,class:BTREE,unique" json:"whiteid"` //
	Whiteloweripaddress1 int64  `gorm:"column:whiteloweripaddress1;type:bigint(20)" json:"whiteloweripaddress1"`                                 //
	Whiteloweripaddress2 *int64 `gorm:"column:whiteloweripaddress2;type:bigint(20)" json:"whiteloweripaddress2"`                                 //
	Whiteupperipaddress1 int64  `gorm:"column:whiteupperipaddress1;type:bigint(20)" json:"whiteupperipaddress1"`                                 //
	Whiteupperipaddress2 *int64 `gorm:"column:whiteupperipaddress2;type:bigint(20)" json:"whiteupperipaddress2"`                                 //
	Whiteemailaddress    string `gorm:"column:whiteemailaddress;type:varchar(255)" json:"whiteemailaddress"`                                     //
	Whitedescription     string `gorm:"column:whitedescription;type:varchar(255)" json:"whitedescription"`                                       //
}

var (
	HmWhitelistFieldWhiteid              = "whiteid"
	HmWhitelistFieldWhiteloweripaddress1 = "whiteloweripaddress1"
	HmWhitelistFieldWhiteloweripaddress2 = "whiteloweripaddress2"
	HmWhitelistFieldWhiteupperipaddress1 = "whiteupperipaddress1"
	HmWhitelistFieldWhiteupperipaddress2 = "whiteupperipaddress2"
	HmWhitelistFieldWhiteemailaddress    = "whiteemailaddress"
	HmWhitelistFieldWhitedescription     = "whitedescription"
)

func (receiver *HmWhitelist) TableName() string {
	return "hm_whitelist"
}

type OrmHmWhitelist struct {
	db *gorm.DB
}

func (orm *OrmHmWhitelist) TableName() string {
	return "hm_whitelist"
}

func NewOrmHmWhitelist(txs ...*gorm.DB) *OrmHmWhitelist {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmWhitelist{})
	} else {
		tx = txs[0].Model(&HmWhitelist{})
	}
	return &OrmHmWhitelist{db: tx}
}

func (orm *OrmHmWhitelist) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmWhitelist) GetTableInfo() interface{} {
	return &HmWhitelist{}
}

// Create insert the value into database
func (orm *OrmHmWhitelist) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmWhitelist) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmWhitelist) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmWhitelist) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmWhitelist) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmWhitelist) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmWhitelist) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmWhitelist) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmWhitelist) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmWhitelist) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmWhitelist) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmWhitelist) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmWhitelist) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmWhitelist) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmWhitelist) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmWhitelist) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmWhitelist) Unscoped() *OrmHmWhitelist {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmWhitelist) Insert(row *HmWhitelist) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmWhitelist) Inserts(rows []*HmWhitelist) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmWhitelist) Order(value interface{}) *OrmHmWhitelist {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmWhitelist) Group(name string) *OrmHmWhitelist {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmWhitelist) Limit(limit int) *OrmHmWhitelist {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmWhitelist) Offset(offset int) *OrmHmWhitelist {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmWhitelist) Get() HmWhitelistList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmWhitelist) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmWhitelist) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmWhitelist{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmWhitelist) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_whitelist")
}

func (orm *OrmHmWhitelist) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmWhitelist) First(conds ...interface{}) (*HmWhitelist, bool) {
	dest := &HmWhitelist{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmWhitelist) Take(conds ...interface{}) (*HmWhitelist, int64) {
	dest := &HmWhitelist{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmWhitelist) Last(conds ...interface{}) (*HmWhitelist, int64) {
	dest := &HmWhitelist{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmWhitelist) Find(conds ...interface{}) (HmWhitelistList, int64) {
	list := make([]*HmWhitelist, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmWhitelist) Paginate(page int, limit int) (HmWhitelistList, int64) {
	var total int64
	list := make([]*HmWhitelist, 0)
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
func (orm *OrmHmWhitelist) SimplePaginate(page int, limit int) HmWhitelistList {
	list := make([]*HmWhitelist, 0)
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
func (orm *OrmHmWhitelist) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmWhitelist) FirstOrInit(dest *HmWhitelist, conds ...interface{}) (*HmWhitelist, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmWhitelist) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmWhitelist) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmWhitelist) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmWhitelist) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmWhitelist) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmWhitelist) Where(query interface{}, args ...interface{}) *OrmHmWhitelist {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmWhitelist) Select(query interface{}, args ...interface{}) *OrmHmWhitelist {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmWhitelist) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmWhitelist) And(fuc func(orm *OrmHmWhitelist)) *OrmHmWhitelist {
	ormAnd := NewOrmHmWhitelist()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmWhitelist) Or(fuc func(orm *OrmHmWhitelist)) *OrmHmWhitelist {
	ormOr := NewOrmHmWhitelist()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmWhitelist) Preload(query string, args ...interface{}) *OrmHmWhitelist {
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
func (orm *OrmHmWhitelist) Joins(query string, args ...interface{}) *OrmHmWhitelist {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmWhitelist) WhereWhiteid(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteid` = ?", val)
	return orm
}
func (orm *OrmHmWhitelist) InsertGetWhiteid(row *HmWhitelist) int64 {
	orm.db.Create(row)
	return row.Whiteid
}
func (orm *OrmHmWhitelist) WhereWhiteidIn(val []int64) *OrmHmWhitelist {
	orm.db.Where("`whiteid` IN ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteidGt(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteid` > ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteidGte(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteid` >= ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteidLt(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteid` < ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteidLte(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteid` <= ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteloweripaddress1(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteloweripaddress1` = ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteloweripaddress2(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteloweripaddress2` = ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteupperipaddress1(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteupperipaddress1` = ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteupperipaddress2(val int64) *OrmHmWhitelist {
	orm.db.Where("`whiteupperipaddress2` = ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhiteemailaddress(val string) *OrmHmWhitelist {
	orm.db.Where("`whiteemailaddress` = ?", val)
	return orm
}
func (orm *OrmHmWhitelist) WhereWhitedescription(val string) *OrmHmWhitelist {
	orm.db.Where("`whitedescription` = ?", val)
	return orm
}

type HmWhitelistList []*HmWhitelist

func (l HmWhitelistList) GetWhiteidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Whiteid)
	}
	return got
}
func (l HmWhitelistList) GetWhiteidMap() map[int64]*HmWhitelist {
	got := make(map[int64]*HmWhitelist)
	for _, val := range l {
		got[val.Whiteid] = val
	}
	return got
}
