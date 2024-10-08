package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmAcl struct {
	Aclid                  int64 `gorm:"column:aclid;autoIncrement;type:bigint(20);primaryKey;index:aclid,class:BTREE,unique" json:"aclid"`                     //
	Aclsharefolderid       int64 `gorm:"column:aclsharefolderid;type:bigint(20);index:aclsharefolderid,class:BTREE,unique" json:"aclsharefolderid"`             //
	Aclpermissiontype      int32 `gorm:"column:aclpermissiontype;type:tinyint(4);index:aclsharefolderid,class:BTREE,unique" json:"aclpermissiontype"`           //
	Aclpermissiongroupid   int64 `gorm:"column:aclpermissiongroupid;type:bigint(20);index:aclsharefolderid,class:BTREE,unique" json:"aclpermissiongroupid"`     //
	Aclpermissionaccountid int64 `gorm:"column:aclpermissionaccountid;type:bigint(20);index:aclsharefolderid,class:BTREE,unique" json:"aclpermissionaccountid"` //
	Aclvalue               int64 `gorm:"column:aclvalue;type:bigint(20)" json:"aclvalue"`                                                                       //
}

var (
	HmAclFieldAclid                  = "aclid"
	HmAclFieldAclsharefolderid       = "aclsharefolderid"
	HmAclFieldAclpermissiontype      = "aclpermissiontype"
	HmAclFieldAclpermissiongroupid   = "aclpermissiongroupid"
	HmAclFieldAclpermissionaccountid = "aclpermissionaccountid"
	HmAclFieldAclvalue               = "aclvalue"
)

func (receiver *HmAcl) TableName() string {
	return "hm_acl"
}

type OrmHmAcl struct {
	db *gorm.DB
}

func (orm *OrmHmAcl) TableName() string {
	return "hm_acl"
}

func NewOrmHmAcl(txs ...*gorm.DB) *OrmHmAcl {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmAcl{})
	} else {
		tx = txs[0].Model(&HmAcl{})
	}
	return &OrmHmAcl{db: tx}
}

func (orm *OrmHmAcl) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmAcl) GetTableInfo() interface{} {
	return &HmAcl{}
}

// Create insert the value into database
func (orm *OrmHmAcl) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmAcl) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmAcl) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmAcl) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmAcl) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmAcl) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmAcl) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmAcl) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmAcl) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmAcl) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmAcl) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmAcl) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmAcl) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmAcl) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmAcl) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmAcl) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmAcl) Unscoped() *OrmHmAcl {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmAcl) Insert(row *HmAcl) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmAcl) Inserts(rows []*HmAcl) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmAcl) Order(value interface{}) *OrmHmAcl {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmAcl) Group(name string) *OrmHmAcl {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmAcl) Limit(limit int) *OrmHmAcl {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmAcl) Offset(offset int) *OrmHmAcl {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmAcl) Get() HmAclList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmAcl) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmAcl) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmAcl{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmAcl) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_acl")
}

func (orm *OrmHmAcl) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmAcl) First(conds ...interface{}) (*HmAcl, bool) {
	dest := &HmAcl{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmAcl) Take(conds ...interface{}) (*HmAcl, int64) {
	dest := &HmAcl{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmAcl) Last(conds ...interface{}) (*HmAcl, int64) {
	dest := &HmAcl{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmAcl) Find(conds ...interface{}) (HmAclList, int64) {
	list := make([]*HmAcl, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmAcl) Paginate(page int, limit int) (HmAclList, int64) {
	var total int64
	list := make([]*HmAcl, 0)
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
func (orm *OrmHmAcl) SimplePaginate(page int, limit int) HmAclList {
	list := make([]*HmAcl, 0)
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
func (orm *OrmHmAcl) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmAcl) FirstOrInit(dest *HmAcl, conds ...interface{}) (*HmAcl, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmAcl) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmAcl) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmAcl) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmAcl) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmAcl) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmAcl) Where(query interface{}, args ...interface{}) *OrmHmAcl {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmAcl) Select(query interface{}, args ...interface{}) *OrmHmAcl {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmAcl) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmAcl) And(fuc func(orm *OrmHmAcl)) *OrmHmAcl {
	ormAnd := NewOrmHmAcl()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmAcl) Or(fuc func(orm *OrmHmAcl)) *OrmHmAcl {
	ormOr := NewOrmHmAcl()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmAcl) Preload(query string, args ...interface{}) *OrmHmAcl {
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
func (orm *OrmHmAcl) Joins(query string, args ...interface{}) *OrmHmAcl {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmAcl) WhereAclid(val int64) *OrmHmAcl {
	orm.db.Where("`aclid` = ?", val)
	return orm
}
func (orm *OrmHmAcl) InsertGetAclid(row *HmAcl) int64 {
	orm.db.Create(row)
	return row.Aclid
}
func (orm *OrmHmAcl) WhereAclidIn(val []int64) *OrmHmAcl {
	orm.db.Where("`aclid` IN ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclidGt(val int64) *OrmHmAcl {
	orm.db.Where("`aclid` > ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclidGte(val int64) *OrmHmAcl {
	orm.db.Where("`aclid` >= ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclidLt(val int64) *OrmHmAcl {
	orm.db.Where("`aclid` < ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclidLte(val int64) *OrmHmAcl {
	orm.db.Where("`aclid` <= ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclsharefolderid(val int64) *OrmHmAcl {
	orm.db.Where("`aclsharefolderid` = ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclsharefolderidIn(val []int64) *OrmHmAcl {
	orm.db.Where("`aclsharefolderid` IN ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclsharefolderidGt(val int64) *OrmHmAcl {
	orm.db.Where("`aclsharefolderid` > ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclsharefolderidGte(val int64) *OrmHmAcl {
	orm.db.Where("`aclsharefolderid` >= ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclsharefolderidLt(val int64) *OrmHmAcl {
	orm.db.Where("`aclsharefolderid` < ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclsharefolderidLte(val int64) *OrmHmAcl {
	orm.db.Where("`aclsharefolderid` <= ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissiontype(val int32) *OrmHmAcl {
	orm.db.Where("`aclpermissiontype` = ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissiongroupid(val int64) *OrmHmAcl {
	orm.db.Where("`aclpermissiongroupid` = ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissiongroupidIn(val []int64) *OrmHmAcl {
	orm.db.Where("`aclpermissiongroupid` IN ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissiongroupidNe(val int64) *OrmHmAcl {
	orm.db.Where("`aclpermissiongroupid` <> ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissionaccountid(val int64) *OrmHmAcl {
	orm.db.Where("`aclpermissionaccountid` = ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissionaccountidIn(val []int64) *OrmHmAcl {
	orm.db.Where("`aclpermissionaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclpermissionaccountidNe(val int64) *OrmHmAcl {
	orm.db.Where("`aclpermissionaccountid` <> ?", val)
	return orm
}
func (orm *OrmHmAcl) WhereAclvalue(val int64) *OrmHmAcl {
	orm.db.Where("`aclvalue` = ?", val)
	return orm
}

type HmAclList []*HmAcl

func (l HmAclList) GetAclidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Aclid)
	}
	return got
}
func (l HmAclList) GetAclidMap() map[int64]*HmAcl {
	got := make(map[int64]*HmAcl)
	for _, val := range l {
		got[val.Aclid] = val
	}
	return got
}
func (l HmAclList) GetAclsharefolderidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Aclsharefolderid)
	}
	return got
}
func (l HmAclList) GetAclsharefolderidMap() map[int64]*HmAcl {
	got := make(map[int64]*HmAcl)
	for _, val := range l {
		got[val.Aclsharefolderid] = val
	}
	return got
}
func (l HmAclList) GetAclpermissiongroupidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Aclpermissiongroupid)
	}
	return got
}
func (l HmAclList) GetAclpermissiongroupidMap() map[int64]*HmAcl {
	got := make(map[int64]*HmAcl)
	for _, val := range l {
		got[val.Aclpermissiongroupid] = val
	}
	return got
}
func (l HmAclList) GetAclpermissionaccountidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Aclpermissionaccountid)
	}
	return got
}
func (l HmAclList) GetAclpermissionaccountidMap() map[int64]*HmAcl {
	got := make(map[int64]*HmAcl)
	for _, val := range l {
		got[val.Aclpermissionaccountid] = val
	}
	return got
}
