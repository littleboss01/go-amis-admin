package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmDomainAliases struct {
	Daid       int32  `gorm:"column:daid;autoIncrement;type:int(11);primaryKey;index:daid,class:BTREE,unique" json:"daid"` //
	Dadomainid int32  `gorm:"column:dadomainid;type:int(11)" json:"dadomainid"`                                            //
	Daalias    string `gorm:"column:daalias;type:varchar(255)" json:"daalias"`                                             //
}

var (
	HmDomainAliasesFieldDaid       = "daid"
	HmDomainAliasesFieldDadomainid = "dadomainid"
	HmDomainAliasesFieldDaalias    = "daalias"
)

func (receiver *HmDomainAliases) TableName() string {
	return "hm_domain_aliases"
}

type OrmHmDomainAliases struct {
	db *gorm.DB
}

func (orm *OrmHmDomainAliases) TableName() string {
	return "hm_domain_aliases"
}

func NewOrmHmDomainAliases(txs ...*gorm.DB) *OrmHmDomainAliases {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmDomainAliases{})
	} else {
		tx = txs[0].Model(&HmDomainAliases{})
	}
	return &OrmHmDomainAliases{db: tx}
}

func (orm *OrmHmDomainAliases) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmDomainAliases) GetTableInfo() interface{} {
	return &HmDomainAliases{}
}

// Create insert the value into database
func (orm *OrmHmDomainAliases) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmDomainAliases) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmDomainAliases) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmDomainAliases) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmDomainAliases) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmDomainAliases) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmDomainAliases) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmDomainAliases) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmDomainAliases) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmDomainAliases) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmDomainAliases) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmDomainAliases) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmDomainAliases) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmDomainAliases) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmDomainAliases) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmDomainAliases) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmDomainAliases) Unscoped() *OrmHmDomainAliases {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmDomainAliases) Insert(row *HmDomainAliases) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmDomainAliases) Inserts(rows []*HmDomainAliases) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmDomainAliases) Order(value interface{}) *OrmHmDomainAliases {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmDomainAliases) Group(name string) *OrmHmDomainAliases {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmDomainAliases) Limit(limit int) *OrmHmDomainAliases {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmDomainAliases) Offset(offset int) *OrmHmDomainAliases {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmDomainAliases) Get() HmDomainAliasesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmDomainAliases) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmDomainAliases) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmDomainAliases{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmDomainAliases) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_domain_aliases")
}

func (orm *OrmHmDomainAliases) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmDomainAliases) First(conds ...interface{}) (*HmDomainAliases, bool) {
	dest := &HmDomainAliases{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmDomainAliases) Take(conds ...interface{}) (*HmDomainAliases, int64) {
	dest := &HmDomainAliases{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmDomainAliases) Last(conds ...interface{}) (*HmDomainAliases, int64) {
	dest := &HmDomainAliases{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmDomainAliases) Find(conds ...interface{}) (HmDomainAliasesList, int64) {
	list := make([]*HmDomainAliases, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmDomainAliases) Paginate(page int, limit int) (HmDomainAliasesList, int64) {
	var total int64
	list := make([]*HmDomainAliases, 0)
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
func (orm *OrmHmDomainAliases) SimplePaginate(page int, limit int) HmDomainAliasesList {
	list := make([]*HmDomainAliases, 0)
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
func (orm *OrmHmDomainAliases) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmDomainAliases) FirstOrInit(dest *HmDomainAliases, conds ...interface{}) (*HmDomainAliases, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmDomainAliases) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDomainAliases) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDomainAliases) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmDomainAliases) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmDomainAliases) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmDomainAliases) Where(query interface{}, args ...interface{}) *OrmHmDomainAliases {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmDomainAliases) Select(query interface{}, args ...interface{}) *OrmHmDomainAliases {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmDomainAliases) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmDomainAliases) And(fuc func(orm *OrmHmDomainAliases)) *OrmHmDomainAliases {
	ormAnd := NewOrmHmDomainAliases()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmDomainAliases) Or(fuc func(orm *OrmHmDomainAliases)) *OrmHmDomainAliases {
	ormOr := NewOrmHmDomainAliases()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmDomainAliases) Preload(query string, args ...interface{}) *OrmHmDomainAliases {
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
func (orm *OrmHmDomainAliases) Joins(query string, args ...interface{}) *OrmHmDomainAliases {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmDomainAliases) WhereDaid(val int32) *OrmHmDomainAliases {
	orm.db.Where("`daid` = ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) InsertGetDaid(row *HmDomainAliases) int32 {
	orm.db.Create(row)
	return row.Daid
}
func (orm *OrmHmDomainAliases) WhereDaidIn(val []int32) *OrmHmDomainAliases {
	orm.db.Where("`daid` IN ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDaidGt(val int32) *OrmHmDomainAliases {
	orm.db.Where("`daid` > ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDaidGte(val int32) *OrmHmDomainAliases {
	orm.db.Where("`daid` >= ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDaidLt(val int32) *OrmHmDomainAliases {
	orm.db.Where("`daid` < ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDaidLte(val int32) *OrmHmDomainAliases {
	orm.db.Where("`daid` <= ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDadomainid(val int32) *OrmHmDomainAliases {
	orm.db.Where("`dadomainid` = ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDadomainidIn(val []int32) *OrmHmDomainAliases {
	orm.db.Where("`dadomainid` IN ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDadomainidNe(val int32) *OrmHmDomainAliases {
	orm.db.Where("`dadomainid` <> ?", val)
	return orm
}
func (orm *OrmHmDomainAliases) WhereDaalias(val string) *OrmHmDomainAliases {
	orm.db.Where("`daalias` = ?", val)
	return orm
}

type HmDomainAliasesList []*HmDomainAliases

func (l HmDomainAliasesList) GetDaidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Daid)
	}
	return got
}
func (l HmDomainAliasesList) GetDaidMap() map[int32]*HmDomainAliases {
	got := make(map[int32]*HmDomainAliases)
	for _, val := range l {
		got[val.Daid] = val
	}
	return got
}
func (l HmDomainAliasesList) GetDadomainidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Dadomainid)
	}
	return got
}
func (l HmDomainAliasesList) GetDadomainidMap() map[int32]*HmDomainAliases {
	got := make(map[int32]*HmDomainAliases)
	for _, val := range l {
		got[val.Dadomainid] = val
	}
	return got
}
