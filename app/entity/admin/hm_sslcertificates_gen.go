package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmSslcertificates struct {
	Sslcertificateid   int64  `gorm:"column:sslcertificateid;autoIncrement;type:bigint(20);primaryKey;index:sslcertificateid,class:BTREE,unique" json:"sslcertificateid"` //
	Sslcertificatename string `gorm:"column:sslcertificatename;type:varchar(255)" json:"sslcertificatename"`                                                              //
	Sslcertificatefile string `gorm:"column:sslcertificatefile;type:varchar(255)" json:"sslcertificatefile"`                                                              //
	Sslprivatekeyfile  string `gorm:"column:sslprivatekeyfile;type:varchar(255)" json:"sslprivatekeyfile"`                                                                //
}

var (
	HmSslcertificatesFieldSslcertificateid   = "sslcertificateid"
	HmSslcertificatesFieldSslcertificatename = "sslcertificatename"
	HmSslcertificatesFieldSslcertificatefile = "sslcertificatefile"
	HmSslcertificatesFieldSslprivatekeyfile  = "sslprivatekeyfile"
)

func (receiver *HmSslcertificates) TableName() string {
	return "hm_sslcertificates"
}

type OrmHmSslcertificates struct {
	db *gorm.DB
}

func (orm *OrmHmSslcertificates) TableName() string {
	return "hm_sslcertificates"
}

func NewOrmHmSslcertificates(txs ...*gorm.DB) *OrmHmSslcertificates {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmSslcertificates{})
	} else {
		tx = txs[0].Model(&HmSslcertificates{})
	}
	return &OrmHmSslcertificates{db: tx}
}

func (orm *OrmHmSslcertificates) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmSslcertificates) GetTableInfo() interface{} {
	return &HmSslcertificates{}
}

// Create insert the value into database
func (orm *OrmHmSslcertificates) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmSslcertificates) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmSslcertificates) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmSslcertificates) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmSslcertificates) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmSslcertificates) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmSslcertificates) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmSslcertificates) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmSslcertificates) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmSslcertificates) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmSslcertificates) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmSslcertificates) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmSslcertificates) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmSslcertificates) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmSslcertificates) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmSslcertificates) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmSslcertificates) Unscoped() *OrmHmSslcertificates {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmSslcertificates) Insert(row *HmSslcertificates) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmSslcertificates) Inserts(rows []*HmSslcertificates) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmSslcertificates) Order(value interface{}) *OrmHmSslcertificates {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmSslcertificates) Group(name string) *OrmHmSslcertificates {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmSslcertificates) Limit(limit int) *OrmHmSslcertificates {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmSslcertificates) Offset(offset int) *OrmHmSslcertificates {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmSslcertificates) Get() HmSslcertificatesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmSslcertificates) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmSslcertificates) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmSslcertificates{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmSslcertificates) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_sslcertificates")
}

func (orm *OrmHmSslcertificates) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmSslcertificates) First(conds ...interface{}) (*HmSslcertificates, bool) {
	dest := &HmSslcertificates{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmSslcertificates) Take(conds ...interface{}) (*HmSslcertificates, int64) {
	dest := &HmSslcertificates{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmSslcertificates) Last(conds ...interface{}) (*HmSslcertificates, int64) {
	dest := &HmSslcertificates{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmSslcertificates) Find(conds ...interface{}) (HmSslcertificatesList, int64) {
	list := make([]*HmSslcertificates, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmSslcertificates) Paginate(page int, limit int) (HmSslcertificatesList, int64) {
	var total int64
	list := make([]*HmSslcertificates, 0)
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
func (orm *OrmHmSslcertificates) SimplePaginate(page int, limit int) HmSslcertificatesList {
	list := make([]*HmSslcertificates, 0)
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
func (orm *OrmHmSslcertificates) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmSslcertificates) FirstOrInit(dest *HmSslcertificates, conds ...interface{}) (*HmSslcertificates, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmSslcertificates) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSslcertificates) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSslcertificates) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmSslcertificates) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmSslcertificates) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmSslcertificates) Where(query interface{}, args ...interface{}) *OrmHmSslcertificates {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmSslcertificates) Select(query interface{}, args ...interface{}) *OrmHmSslcertificates {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmSslcertificates) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmSslcertificates) And(fuc func(orm *OrmHmSslcertificates)) *OrmHmSslcertificates {
	ormAnd := NewOrmHmSslcertificates()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmSslcertificates) Or(fuc func(orm *OrmHmSslcertificates)) *OrmHmSslcertificates {
	ormOr := NewOrmHmSslcertificates()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmSslcertificates) Preload(query string, args ...interface{}) *OrmHmSslcertificates {
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
func (orm *OrmHmSslcertificates) Joins(query string, args ...interface{}) *OrmHmSslcertificates {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmSslcertificates) WhereSslcertificateid(val int64) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificateid` = ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) InsertGetSslcertificateid(row *HmSslcertificates) int64 {
	orm.db.Create(row)
	return row.Sslcertificateid
}
func (orm *OrmHmSslcertificates) WhereSslcertificateidIn(val []int64) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificateid` IN ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslcertificateidGt(val int64) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificateid` > ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslcertificateidGte(val int64) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificateid` >= ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslcertificateidLt(val int64) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificateid` < ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslcertificateidLte(val int64) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificateid` <= ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslcertificatename(val string) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificatename` = ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslcertificatefile(val string) *OrmHmSslcertificates {
	orm.db.Where("`sslcertificatefile` = ?", val)
	return orm
}
func (orm *OrmHmSslcertificates) WhereSslprivatekeyfile(val string) *OrmHmSslcertificates {
	orm.db.Where("`sslprivatekeyfile` = ?", val)
	return orm
}

type HmSslcertificatesList []*HmSslcertificates

func (l HmSslcertificatesList) GetSslcertificateidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Sslcertificateid)
	}
	return got
}
func (l HmSslcertificatesList) GetSslcertificateidMap() map[int64]*HmSslcertificates {
	got := make(map[int64]*HmSslcertificates)
	for _, val := range l {
		got[val.Sslcertificateid] = val
	}
	return got
}
