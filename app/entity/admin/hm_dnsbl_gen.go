package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmDnsbl struct {
	Sblid            int32  `gorm:"column:sblid;autoIncrement;type:int(11);primaryKey;index:sblid,class:BTREE,unique" json:"sblid"` //
	Sblactive        int32  `gorm:"column:sblactive;type:tinyint(4)" json:"sblactive"`                                              //
	Sbldnshost       string `gorm:"column:sbldnshost;type:varchar(255)" json:"sbldnshost"`                                          //
	Sblresult        string `gorm:"column:sblresult;type:varchar(255)" json:"sblresult"`                                            //
	Sblrejectmessage string `gorm:"column:sblrejectmessage;type:varchar(255)" json:"sblrejectmessage"`                              //
	Sblscore         int32  `gorm:"column:sblscore;type:int(11)" json:"sblscore"`                                                   //
}

var (
	HmDnsblFieldSblid            = "sblid"
	HmDnsblFieldSblactive        = "sblactive"
	HmDnsblFieldSbldnshost       = "sbldnshost"
	HmDnsblFieldSblresult        = "sblresult"
	HmDnsblFieldSblrejectmessage = "sblrejectmessage"
	HmDnsblFieldSblscore         = "sblscore"
)

func (receiver *HmDnsbl) TableName() string {
	return "hm_dnsbl"
}

type OrmHmDnsbl struct {
	db *gorm.DB
}

func (orm *OrmHmDnsbl) TableName() string {
	return "hm_dnsbl"
}

func NewOrmHmDnsbl(txs ...*gorm.DB) *OrmHmDnsbl {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmDnsbl{})
	} else {
		tx = txs[0].Model(&HmDnsbl{})
	}
	return &OrmHmDnsbl{db: tx}
}

func (orm *OrmHmDnsbl) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmDnsbl) GetTableInfo() interface{} {
	return &HmDnsbl{}
}

// Create insert the value into database
func (orm *OrmHmDnsbl) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmDnsbl) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmDnsbl) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmDnsbl) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmDnsbl) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmDnsbl) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmDnsbl) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmDnsbl) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmDnsbl) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmDnsbl) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmDnsbl) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmDnsbl) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmDnsbl) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmDnsbl) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmDnsbl) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmDnsbl) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmDnsbl) Unscoped() *OrmHmDnsbl {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmDnsbl) Insert(row *HmDnsbl) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmDnsbl) Inserts(rows []*HmDnsbl) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmDnsbl) Order(value interface{}) *OrmHmDnsbl {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmDnsbl) Group(name string) *OrmHmDnsbl {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmDnsbl) Limit(limit int) *OrmHmDnsbl {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmDnsbl) Offset(offset int) *OrmHmDnsbl {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmDnsbl) Get() HmDnsblList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmDnsbl) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmDnsbl) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmDnsbl{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmDnsbl) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_dnsbl")
}

func (orm *OrmHmDnsbl) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmDnsbl) First(conds ...interface{}) (*HmDnsbl, bool) {
	dest := &HmDnsbl{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmDnsbl) Take(conds ...interface{}) (*HmDnsbl, int64) {
	dest := &HmDnsbl{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmDnsbl) Last(conds ...interface{}) (*HmDnsbl, int64) {
	dest := &HmDnsbl{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmDnsbl) Find(conds ...interface{}) (HmDnsblList, int64) {
	list := make([]*HmDnsbl, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmDnsbl) Paginate(page int, limit int) (HmDnsblList, int64) {
	var total int64
	list := make([]*HmDnsbl, 0)
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
func (orm *OrmHmDnsbl) SimplePaginate(page int, limit int) HmDnsblList {
	list := make([]*HmDnsbl, 0)
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
func (orm *OrmHmDnsbl) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmDnsbl) FirstOrInit(dest *HmDnsbl, conds ...interface{}) (*HmDnsbl, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmDnsbl) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDnsbl) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDnsbl) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmDnsbl) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmDnsbl) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmDnsbl) Where(query interface{}, args ...interface{}) *OrmHmDnsbl {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmDnsbl) Select(query interface{}, args ...interface{}) *OrmHmDnsbl {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmDnsbl) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmDnsbl) And(fuc func(orm *OrmHmDnsbl)) *OrmHmDnsbl {
	ormAnd := NewOrmHmDnsbl()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmDnsbl) Or(fuc func(orm *OrmHmDnsbl)) *OrmHmDnsbl {
	ormOr := NewOrmHmDnsbl()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmDnsbl) Preload(query string, args ...interface{}) *OrmHmDnsbl {
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
func (orm *OrmHmDnsbl) Joins(query string, args ...interface{}) *OrmHmDnsbl {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmDnsbl) WhereSblid(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblid` = ?", val)
	return orm
}
func (orm *OrmHmDnsbl) InsertGetSblid(row *HmDnsbl) int32 {
	orm.db.Create(row)
	return row.Sblid
}
func (orm *OrmHmDnsbl) WhereSblidIn(val []int32) *OrmHmDnsbl {
	orm.db.Where("`sblid` IN ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblidGt(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblid` > ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblidGte(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblid` >= ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblidLt(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblid` < ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblidLte(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblid` <= ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblactive(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblactive` = ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSbldnshost(val string) *OrmHmDnsbl {
	orm.db.Where("`sbldnshost` = ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblresult(val string) *OrmHmDnsbl {
	orm.db.Where("`sblresult` = ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblrejectmessage(val string) *OrmHmDnsbl {
	orm.db.Where("`sblrejectmessage` = ?", val)
	return orm
}
func (orm *OrmHmDnsbl) WhereSblscore(val int32) *OrmHmDnsbl {
	orm.db.Where("`sblscore` = ?", val)
	return orm
}

type HmDnsblList []*HmDnsbl

func (l HmDnsblList) GetSblidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Sblid)
	}
	return got
}
func (l HmDnsblList) GetSblidMap() map[int32]*HmDnsbl {
	got := make(map[int32]*HmDnsbl)
	for _, val := range l {
		got[val.Sblid] = val
	}
	return got
}
