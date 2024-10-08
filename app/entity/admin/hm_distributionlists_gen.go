package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmDistributionlists struct {
	Distributionlistid             int32  `gorm:"column:distributionlistid;autoIncrement;type:int(11);primaryKey;index:distributionlistid,class:BTREE,unique" json:"distributionlistid"` //
	Distributionlistdomainid       int32  `gorm:"column:distributionlistdomainid;type:int(11);index:idx_hm_distributionlists,class:BTREE" json:"distributionlistdomainid"`               //
	Distributionlistaddress        string `gorm:"column:distributionlistaddress;type:varchar(255);index:distributionlistaddress,class:BTREE,unique" json:"distributionlistaddress"`      //
	Distributionlistenabled        int32  `gorm:"column:distributionlistenabled;type:tinyint(4)" json:"distributionlistenabled"`                                                         //
	Distributionlistrequireauth    int32  `gorm:"column:distributionlistrequireauth;type:tinyint(4)" json:"distributionlistrequireauth"`                                                 //
	Distributionlistrequireaddress string `gorm:"column:distributionlistrequireaddress;type:varchar(255)" json:"distributionlistrequireaddress"`                                         //
	Distributionlistmode           int32  `gorm:"column:distributionlistmode;type:tinyint(4)" json:"distributionlistmode"`                                                               //
}

var (
	HmDistributionlistsFieldDistributionlistid             = "distributionlistid"
	HmDistributionlistsFieldDistributionlistdomainid       = "distributionlistdomainid"
	HmDistributionlistsFieldDistributionlistaddress        = "distributionlistaddress"
	HmDistributionlistsFieldDistributionlistenabled        = "distributionlistenabled"
	HmDistributionlistsFieldDistributionlistrequireauth    = "distributionlistrequireauth"
	HmDistributionlistsFieldDistributionlistrequireaddress = "distributionlistrequireaddress"
	HmDistributionlistsFieldDistributionlistmode           = "distributionlistmode"
)

func (receiver *HmDistributionlists) TableName() string {
	return "hm_distributionlists"
}

type OrmHmDistributionlists struct {
	db *gorm.DB
}

func (orm *OrmHmDistributionlists) TableName() string {
	return "hm_distributionlists"
}

func NewOrmHmDistributionlists(txs ...*gorm.DB) *OrmHmDistributionlists {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmDistributionlists{})
	} else {
		tx = txs[0].Model(&HmDistributionlists{})
	}
	return &OrmHmDistributionlists{db: tx}
}

func (orm *OrmHmDistributionlists) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmDistributionlists) GetTableInfo() interface{} {
	return &HmDistributionlists{}
}

// Create insert the value into database
func (orm *OrmHmDistributionlists) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmDistributionlists) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmDistributionlists) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmDistributionlists) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmDistributionlists) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmDistributionlists) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmDistributionlists) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmDistributionlists) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmDistributionlists) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmDistributionlists) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmDistributionlists) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmDistributionlists) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmDistributionlists) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmDistributionlists) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmDistributionlists) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmDistributionlists) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmDistributionlists) Unscoped() *OrmHmDistributionlists {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmDistributionlists) Insert(row *HmDistributionlists) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmDistributionlists) Inserts(rows []*HmDistributionlists) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmDistributionlists) Order(value interface{}) *OrmHmDistributionlists {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmDistributionlists) Group(name string) *OrmHmDistributionlists {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmDistributionlists) Limit(limit int) *OrmHmDistributionlists {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmDistributionlists) Offset(offset int) *OrmHmDistributionlists {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmDistributionlists) Get() HmDistributionlistsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmDistributionlists) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmDistributionlists) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmDistributionlists{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmDistributionlists) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_distributionlists")
}

func (orm *OrmHmDistributionlists) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmDistributionlists) First(conds ...interface{}) (*HmDistributionlists, bool) {
	dest := &HmDistributionlists{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmDistributionlists) Take(conds ...interface{}) (*HmDistributionlists, int64) {
	dest := &HmDistributionlists{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmDistributionlists) Last(conds ...interface{}) (*HmDistributionlists, int64) {
	dest := &HmDistributionlists{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmDistributionlists) Find(conds ...interface{}) (HmDistributionlistsList, int64) {
	list := make([]*HmDistributionlists, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmDistributionlists) Paginate(page int, limit int) (HmDistributionlistsList, int64) {
	var total int64
	list := make([]*HmDistributionlists, 0)
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
func (orm *OrmHmDistributionlists) SimplePaginate(page int, limit int) HmDistributionlistsList {
	list := make([]*HmDistributionlists, 0)
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
func (orm *OrmHmDistributionlists) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmDistributionlists) FirstOrInit(dest *HmDistributionlists, conds ...interface{}) (*HmDistributionlists, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmDistributionlists) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDistributionlists) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDistributionlists) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmDistributionlists) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmDistributionlists) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmDistributionlists) Where(query interface{}, args ...interface{}) *OrmHmDistributionlists {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmDistributionlists) Select(query interface{}, args ...interface{}) *OrmHmDistributionlists {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmDistributionlists) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmDistributionlists) And(fuc func(orm *OrmHmDistributionlists)) *OrmHmDistributionlists {
	ormAnd := NewOrmHmDistributionlists()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmDistributionlists) Or(fuc func(orm *OrmHmDistributionlists)) *OrmHmDistributionlists {
	ormOr := NewOrmHmDistributionlists()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmDistributionlists) Preload(query string, args ...interface{}) *OrmHmDistributionlists {
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
func (orm *OrmHmDistributionlists) Joins(query string, args ...interface{}) *OrmHmDistributionlists {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmDistributionlists) WhereDistributionlistid(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistid` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) InsertGetDistributionlistid(row *HmDistributionlists) int32 {
	orm.db.Create(row)
	return row.Distributionlistid
}
func (orm *OrmHmDistributionlists) WhereDistributionlistidIn(val []int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistid` IN ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistidGt(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistid` > ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistidGte(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistid` >= ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistidLt(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistid` < ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistidLte(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistid` <= ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistdomainid(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistdomainid` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistdomainidIn(val []int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistdomainid` IN ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistdomainidGt(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistdomainid` > ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistdomainidGte(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistdomainid` >= ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistdomainidLt(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistdomainid` < ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistdomainidLte(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistdomainid` <= ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistaddress(val string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistaddress` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistaddressIn(val []string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistaddress` IN ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistaddressGt(val string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistaddress` > ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistaddressGte(val string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistaddress` >= ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistaddressLt(val string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistaddress` < ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistaddressLte(val string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistaddress` <= ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistenabled(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistenabled` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistrequireauth(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistrequireauth` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistrequireaddress(val string) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistrequireaddress` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlists) WhereDistributionlistmode(val int32) *OrmHmDistributionlists {
	orm.db.Where("`distributionlistmode` = ?", val)
	return orm
}

type HmDistributionlistsList []*HmDistributionlists

func (l HmDistributionlistsList) GetDistributionlistidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Distributionlistid)
	}
	return got
}
func (l HmDistributionlistsList) GetDistributionlistidMap() map[int32]*HmDistributionlists {
	got := make(map[int32]*HmDistributionlists)
	for _, val := range l {
		got[val.Distributionlistid] = val
	}
	return got
}
func (l HmDistributionlistsList) GetDistributionlistdomainidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Distributionlistdomainid)
	}
	return got
}
func (l HmDistributionlistsList) GetDistributionlistdomainidMap() map[int32]*HmDistributionlists {
	got := make(map[int32]*HmDistributionlists)
	for _, val := range l {
		got[val.Distributionlistdomainid] = val
	}
	return got
}
