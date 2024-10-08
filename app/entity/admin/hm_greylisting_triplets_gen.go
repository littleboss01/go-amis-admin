package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmGreylistingTriplets struct {
	Glid               int64         `gorm:"column:glid;autoIncrement;type:bigint(20);primaryKey;index:glid,class:BTREE,unique" json:"glid"`                   //
	Glcreatetime       database.Time `gorm:"column:glcreatetime;type:datetime" json:"glcreatetime"`                                                            //
	Glblockendtime     database.Time `gorm:"column:glblockendtime;type:datetime" json:"glblockendtime"`                                                        //
	Gldeletetime       database.Time `gorm:"column:gldeletetime;type:datetime" json:"gldeletetime"`                                                            //
	Glipaddress1       int64         `gorm:"column:glipaddress1;type:bigint(20);index:idx_greylisting_triplets,class:BTREE" json:"glipaddress1"`               //
	Glipaddress2       *int64        `gorm:"column:glipaddress2;type:bigint(20);index:idx_greylisting_triplets,class:BTREE" json:"glipaddress2"`               //
	Glsenderaddress    string        `gorm:"column:glsenderaddress;type:varchar(255);index:idx_greylisting_triplets,class:BTREE" json:"glsenderaddress"`       //
	Glrecipientaddress string        `gorm:"column:glrecipientaddress;type:varchar(255);index:idx_greylisting_triplets,class:BTREE" json:"glrecipientaddress"` //
	Glblockedcount     int32         `gorm:"column:glblockedcount;type:int(11)" json:"glblockedcount"`                                                         //
	Glpassedcount      int32         `gorm:"column:glpassedcount;type:int(11)" json:"glpassedcount"`                                                           //
}

var (
	HmGreylistingTripletsFieldGlid               = "glid"
	HmGreylistingTripletsFieldGlcreatetime       = "glcreatetime"
	HmGreylistingTripletsFieldGlblockendtime     = "glblockendtime"
	HmGreylistingTripletsFieldGldeletetime       = "gldeletetime"
	HmGreylistingTripletsFieldGlipaddress1       = "glipaddress1"
	HmGreylistingTripletsFieldGlipaddress2       = "glipaddress2"
	HmGreylistingTripletsFieldGlsenderaddress    = "glsenderaddress"
	HmGreylistingTripletsFieldGlrecipientaddress = "glrecipientaddress"
	HmGreylistingTripletsFieldGlblockedcount     = "glblockedcount"
	HmGreylistingTripletsFieldGlpassedcount      = "glpassedcount"
)

func (receiver *HmGreylistingTriplets) TableName() string {
	return "hm_greylisting_triplets"
}

type OrmHmGreylistingTriplets struct {
	db *gorm.DB
}

func (orm *OrmHmGreylistingTriplets) TableName() string {
	return "hm_greylisting_triplets"
}

func NewOrmHmGreylistingTriplets(txs ...*gorm.DB) *OrmHmGreylistingTriplets {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmGreylistingTriplets{})
	} else {
		tx = txs[0].Model(&HmGreylistingTriplets{})
	}
	return &OrmHmGreylistingTriplets{db: tx}
}

func (orm *OrmHmGreylistingTriplets) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmGreylistingTriplets) GetTableInfo() interface{} {
	return &HmGreylistingTriplets{}
}

// Create insert the value into database
func (orm *OrmHmGreylistingTriplets) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmGreylistingTriplets) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmGreylistingTriplets) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmGreylistingTriplets) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmGreylistingTriplets) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmGreylistingTriplets) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmGreylistingTriplets) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmGreylistingTriplets) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmGreylistingTriplets) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmGreylistingTriplets) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmGreylistingTriplets) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmGreylistingTriplets) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmGreylistingTriplets) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmGreylistingTriplets) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmGreylistingTriplets) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmGreylistingTriplets) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmGreylistingTriplets) Unscoped() *OrmHmGreylistingTriplets {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmGreylistingTriplets) Insert(row *HmGreylistingTriplets) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmGreylistingTriplets) Inserts(rows []*HmGreylistingTriplets) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmGreylistingTriplets) Order(value interface{}) *OrmHmGreylistingTriplets {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmGreylistingTriplets) Group(name string) *OrmHmGreylistingTriplets {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmGreylistingTriplets) Limit(limit int) *OrmHmGreylistingTriplets {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmGreylistingTriplets) Offset(offset int) *OrmHmGreylistingTriplets {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmGreylistingTriplets) Get() HmGreylistingTripletsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmGreylistingTriplets) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmGreylistingTriplets) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmGreylistingTriplets{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmGreylistingTriplets) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_greylisting_triplets")
}

func (orm *OrmHmGreylistingTriplets) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmGreylistingTriplets) First(conds ...interface{}) (*HmGreylistingTriplets, bool) {
	dest := &HmGreylistingTriplets{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmGreylistingTriplets) Take(conds ...interface{}) (*HmGreylistingTriplets, int64) {
	dest := &HmGreylistingTriplets{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmGreylistingTriplets) Last(conds ...interface{}) (*HmGreylistingTriplets, int64) {
	dest := &HmGreylistingTriplets{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmGreylistingTriplets) Find(conds ...interface{}) (HmGreylistingTripletsList, int64) {
	list := make([]*HmGreylistingTriplets, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmGreylistingTriplets) Paginate(page int, limit int) (HmGreylistingTripletsList, int64) {
	var total int64
	list := make([]*HmGreylistingTriplets, 0)
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
func (orm *OrmHmGreylistingTriplets) SimplePaginate(page int, limit int) HmGreylistingTripletsList {
	list := make([]*HmGreylistingTriplets, 0)
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
func (orm *OrmHmGreylistingTriplets) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmGreylistingTriplets) FirstOrInit(dest *HmGreylistingTriplets, conds ...interface{}) (*HmGreylistingTriplets, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmGreylistingTriplets) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGreylistingTriplets) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGreylistingTriplets) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmGreylistingTriplets) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmGreylistingTriplets) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmGreylistingTriplets) Where(query interface{}, args ...interface{}) *OrmHmGreylistingTriplets {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmGreylistingTriplets) Select(query interface{}, args ...interface{}) *OrmHmGreylistingTriplets {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmGreylistingTriplets) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmGreylistingTriplets) And(fuc func(orm *OrmHmGreylistingTriplets)) *OrmHmGreylistingTriplets {
	ormAnd := NewOrmHmGreylistingTriplets()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmGreylistingTriplets) Or(fuc func(orm *OrmHmGreylistingTriplets)) *OrmHmGreylistingTriplets {
	ormOr := NewOrmHmGreylistingTriplets()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmGreylistingTriplets) Preload(query string, args ...interface{}) *OrmHmGreylistingTriplets {
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
func (orm *OrmHmGreylistingTriplets) Joins(query string, args ...interface{}) *OrmHmGreylistingTriplets {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmGreylistingTriplets) WhereGlid(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glid` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) InsertGetGlid(row *HmGreylistingTriplets) int64 {
	orm.db.Create(row)
	return row.Glid
}
func (orm *OrmHmGreylistingTriplets) WhereGlidIn(val []int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glid` IN ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlidGt(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glid` > ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlidGte(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glid` >= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlidLt(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glid` < ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlidLte(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glid` <= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlcreatetime(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glcreatetime` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlcreatetimeBetween(begin database.Time, end database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glcreatetime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlcreatetimeLte(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glcreatetime` <= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlcreatetimeGte(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glcreatetime` >= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlblockendtime(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glblockendtime` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlblockendtimeBetween(begin database.Time, end database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glblockendtime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlblockendtimeLte(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glblockendtime` <= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlblockendtimeGte(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`glblockendtime` >= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGldeletetime(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`gldeletetime` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGldeletetimeBetween(begin database.Time, end database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`gldeletetime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGldeletetimeLte(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`gldeletetime` <= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGldeletetimeGte(val database.Time) *OrmHmGreylistingTriplets {
	orm.db.Where("`gldeletetime` >= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress1(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress1` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress1In(val []int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress1` IN ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress1Gt(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress1` > ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress1Gte(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress1` >= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress1Lt(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress1` < ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress1Lte(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress1` <= ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlipaddress2(val int64) *OrmHmGreylistingTriplets {
	orm.db.Where("`glipaddress2` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlsenderaddress(val string) *OrmHmGreylistingTriplets {
	orm.db.Where("`glsenderaddress` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlrecipientaddress(val string) *OrmHmGreylistingTriplets {
	orm.db.Where("`glrecipientaddress` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlblockedcount(val int32) *OrmHmGreylistingTriplets {
	orm.db.Where("`glblockedcount` = ?", val)
	return orm
}
func (orm *OrmHmGreylistingTriplets) WhereGlpassedcount(val int32) *OrmHmGreylistingTriplets {
	orm.db.Where("`glpassedcount` = ?", val)
	return orm
}

type HmGreylistingTripletsList []*HmGreylistingTriplets

func (l HmGreylistingTripletsList) GetGlidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Glid)
	}
	return got
}
func (l HmGreylistingTripletsList) GetGlidMap() map[int64]*HmGreylistingTriplets {
	got := make(map[int64]*HmGreylistingTriplets)
	for _, val := range l {
		got[val.Glid] = val
	}
	return got
}
