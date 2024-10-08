package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmAliases struct {
	Aliasid       int32  `gorm:"column:aliasid;autoIncrement;type:int(11);primaryKey;index:aliasid,class:BTREE,unique" json:"aliasid"`                    //
	Aliasdomainid int32  `gorm:"column:aliasdomainid;type:int(11);index:idx_hm_aliases,class:BTREE" json:"aliasdomainid"`                                 //
	Aliasname     string `gorm:"column:aliasname;type:varchar(255);index:aliasname,class:BTREE,unique;index:idx_hm_aliases,class:BTREE" json:"aliasname"` //
	Aliasvalue    string `gorm:"column:aliasvalue;type:varchar(255)" json:"aliasvalue"`                                                                   //
	Aliasactive   int32  `gorm:"column:aliasactive;type:tinyint(4)" json:"aliasactive"`                                                                   //
}

var (
	HmAliasesFieldAliasid       = "aliasid"
	HmAliasesFieldAliasdomainid = "aliasdomainid"
	HmAliasesFieldAliasname     = "aliasname"
	HmAliasesFieldAliasvalue    = "aliasvalue"
	HmAliasesFieldAliasactive   = "aliasactive"
)

func (receiver *HmAliases) TableName() string {
	return "hm_aliases"
}

type OrmHmAliases struct {
	db *gorm.DB
}

func (orm *OrmHmAliases) TableName() string {
	return "hm_aliases"
}

func NewOrmHmAliases(txs ...*gorm.DB) *OrmHmAliases {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmAliases{})
	} else {
		tx = txs[0].Model(&HmAliases{})
	}
	return &OrmHmAliases{db: tx}
}

func (orm *OrmHmAliases) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmAliases) GetTableInfo() interface{} {
	return &HmAliases{}
}

// Create insert the value into database
func (orm *OrmHmAliases) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmAliases) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmAliases) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmAliases) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmAliases) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmAliases) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmAliases) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmAliases) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmAliases) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmAliases) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmAliases) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmAliases) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmAliases) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmAliases) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmAliases) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmAliases) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmAliases) Unscoped() *OrmHmAliases {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmAliases) Insert(row *HmAliases) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmAliases) Inserts(rows []*HmAliases) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmAliases) Order(value interface{}) *OrmHmAliases {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmAliases) Group(name string) *OrmHmAliases {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmAliases) Limit(limit int) *OrmHmAliases {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmAliases) Offset(offset int) *OrmHmAliases {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmAliases) Get() HmAliasesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmAliases) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmAliases) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmAliases{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmAliases) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_aliases")
}

func (orm *OrmHmAliases) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmAliases) First(conds ...interface{}) (*HmAliases, bool) {
	dest := &HmAliases{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmAliases) Take(conds ...interface{}) (*HmAliases, int64) {
	dest := &HmAliases{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmAliases) Last(conds ...interface{}) (*HmAliases, int64) {
	dest := &HmAliases{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmAliases) Find(conds ...interface{}) (HmAliasesList, int64) {
	list := make([]*HmAliases, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmAliases) Paginate(page int, limit int) (HmAliasesList, int64) {
	var total int64
	list := make([]*HmAliases, 0)
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
func (orm *OrmHmAliases) SimplePaginate(page int, limit int) HmAliasesList {
	list := make([]*HmAliases, 0)
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
func (orm *OrmHmAliases) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmAliases) FirstOrInit(dest *HmAliases, conds ...interface{}) (*HmAliases, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmAliases) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmAliases) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmAliases) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmAliases) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmAliases) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmAliases) Where(query interface{}, args ...interface{}) *OrmHmAliases {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmAliases) Select(query interface{}, args ...interface{}) *OrmHmAliases {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmAliases) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmAliases) And(fuc func(orm *OrmHmAliases)) *OrmHmAliases {
	ormAnd := NewOrmHmAliases()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmAliases) Or(fuc func(orm *OrmHmAliases)) *OrmHmAliases {
	ormOr := NewOrmHmAliases()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmAliases) Preload(query string, args ...interface{}) *OrmHmAliases {
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
func (orm *OrmHmAliases) Joins(query string, args ...interface{}) *OrmHmAliases {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmAliases) WhereAliasid(val int32) *OrmHmAliases {
	orm.db.Where("`aliasid` = ?", val)
	return orm
}
func (orm *OrmHmAliases) InsertGetAliasid(row *HmAliases) int32 {
	orm.db.Create(row)
	return row.Aliasid
}
func (orm *OrmHmAliases) WhereAliasidIn(val []int32) *OrmHmAliases {
	orm.db.Where("`aliasid` IN ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasidGt(val int32) *OrmHmAliases {
	orm.db.Where("`aliasid` > ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasidGte(val int32) *OrmHmAliases {
	orm.db.Where("`aliasid` >= ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasidLt(val int32) *OrmHmAliases {
	orm.db.Where("`aliasid` < ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasidLte(val int32) *OrmHmAliases {
	orm.db.Where("`aliasid` <= ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasdomainid(val int32) *OrmHmAliases {
	orm.db.Where("`aliasdomainid` = ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasdomainidIn(val []int32) *OrmHmAliases {
	orm.db.Where("`aliasdomainid` IN ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasdomainidGt(val int32) *OrmHmAliases {
	orm.db.Where("`aliasdomainid` > ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasdomainidGte(val int32) *OrmHmAliases {
	orm.db.Where("`aliasdomainid` >= ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasdomainidLt(val int32) *OrmHmAliases {
	orm.db.Where("`aliasdomainid` < ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasdomainidLte(val int32) *OrmHmAliases {
	orm.db.Where("`aliasdomainid` <= ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasname(val string) *OrmHmAliases {
	orm.db.Where("`aliasname` = ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasnameIn(val []string) *OrmHmAliases {
	orm.db.Where("`aliasname` IN ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasnameGt(val string) *OrmHmAliases {
	orm.db.Where("`aliasname` > ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasnameGte(val string) *OrmHmAliases {
	orm.db.Where("`aliasname` >= ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasnameLt(val string) *OrmHmAliases {
	orm.db.Where("`aliasname` < ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasnameLte(val string) *OrmHmAliases {
	orm.db.Where("`aliasname` <= ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasvalue(val string) *OrmHmAliases {
	orm.db.Where("`aliasvalue` = ?", val)
	return orm
}
func (orm *OrmHmAliases) WhereAliasactive(val int32) *OrmHmAliases {
	orm.db.Where("`aliasactive` = ?", val)
	return orm
}

type HmAliasesList []*HmAliases

func (l HmAliasesList) GetAliasidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Aliasid)
	}
	return got
}
func (l HmAliasesList) GetAliasidMap() map[int32]*HmAliases {
	got := make(map[int32]*HmAliases)
	for _, val := range l {
		got[val.Aliasid] = val
	}
	return got
}
func (l HmAliasesList) GetAliasdomainidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Aliasdomainid)
	}
	return got
}
func (l HmAliasesList) GetAliasdomainidMap() map[int32]*HmAliases {
	got := make(map[int32]*HmAliases)
	for _, val := range l {
		got[val.Aliasdomainid] = val
	}
	return got
}
