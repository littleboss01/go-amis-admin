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

type HmFetchaccountsUids struct {
	Uidid    int32         `gorm:"column:uidid;autoIncrement;type:int(11);primaryKey;index:uidid,class:BTREE,unique" json:"uidid"` //
	Uidfaid  int32         `gorm:"column:uidfaid;type:int(11);index:idx_hm_fetchaccounts_uids,class:BTREE" json:"uidfaid"`         //
	Uidvalue string        `gorm:"column:uidvalue;type:varchar(255)" json:"uidvalue"`                                              //
	Uidtime  database.Time `gorm:"column:uidtime;type:datetime" json:"uidtime"`                                                    //
}

var (
	HmFetchaccountsUidsFieldUidid    = "uidid"
	HmFetchaccountsUidsFieldUidfaid  = "uidfaid"
	HmFetchaccountsUidsFieldUidvalue = "uidvalue"
	HmFetchaccountsUidsFieldUidtime  = "uidtime"
)

func (receiver *HmFetchaccountsUids) TableName() string {
	return "hm_fetchaccounts_uids"
}

type OrmHmFetchaccountsUids struct {
	db *gorm.DB
}

func (orm *OrmHmFetchaccountsUids) TableName() string {
	return "hm_fetchaccounts_uids"
}

func NewOrmHmFetchaccountsUids(txs ...*gorm.DB) *OrmHmFetchaccountsUids {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmFetchaccountsUids{})
	} else {
		tx = txs[0].Model(&HmFetchaccountsUids{})
	}
	return &OrmHmFetchaccountsUids{db: tx}
}

func (orm *OrmHmFetchaccountsUids) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmFetchaccountsUids) GetTableInfo() interface{} {
	return &HmFetchaccountsUids{}
}

// Create insert the value into database
func (orm *OrmHmFetchaccountsUids) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmFetchaccountsUids) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmFetchaccountsUids) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmFetchaccountsUids) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmFetchaccountsUids) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmFetchaccountsUids) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmFetchaccountsUids) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmFetchaccountsUids) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmFetchaccountsUids) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmFetchaccountsUids) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmFetchaccountsUids) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmFetchaccountsUids) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmFetchaccountsUids) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmFetchaccountsUids) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmFetchaccountsUids) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmFetchaccountsUids) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmFetchaccountsUids) Unscoped() *OrmHmFetchaccountsUids {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmFetchaccountsUids) Insert(row *HmFetchaccountsUids) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmFetchaccountsUids) Inserts(rows []*HmFetchaccountsUids) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmFetchaccountsUids) Order(value interface{}) *OrmHmFetchaccountsUids {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmFetchaccountsUids) Group(name string) *OrmHmFetchaccountsUids {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmFetchaccountsUids) Limit(limit int) *OrmHmFetchaccountsUids {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmFetchaccountsUids) Offset(offset int) *OrmHmFetchaccountsUids {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmFetchaccountsUids) Get() HmFetchaccountsUidsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmFetchaccountsUids) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmFetchaccountsUids) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmFetchaccountsUids{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmFetchaccountsUids) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_fetchaccounts_uids")
}

func (orm *OrmHmFetchaccountsUids) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmFetchaccountsUids) First(conds ...interface{}) (*HmFetchaccountsUids, bool) {
	dest := &HmFetchaccountsUids{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmFetchaccountsUids) Take(conds ...interface{}) (*HmFetchaccountsUids, int64) {
	dest := &HmFetchaccountsUids{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmFetchaccountsUids) Last(conds ...interface{}) (*HmFetchaccountsUids, int64) {
	dest := &HmFetchaccountsUids{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmFetchaccountsUids) Find(conds ...interface{}) (HmFetchaccountsUidsList, int64) {
	list := make([]*HmFetchaccountsUids, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmFetchaccountsUids) Paginate(page int, limit int) (HmFetchaccountsUidsList, int64) {
	var total int64
	list := make([]*HmFetchaccountsUids, 0)
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
func (orm *OrmHmFetchaccountsUids) SimplePaginate(page int, limit int) HmFetchaccountsUidsList {
	list := make([]*HmFetchaccountsUids, 0)
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
func (orm *OrmHmFetchaccountsUids) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmFetchaccountsUids) FirstOrInit(dest *HmFetchaccountsUids, conds ...interface{}) (*HmFetchaccountsUids, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmFetchaccountsUids) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmFetchaccountsUids) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmFetchaccountsUids) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmFetchaccountsUids) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmFetchaccountsUids) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmFetchaccountsUids) Where(query interface{}, args ...interface{}) *OrmHmFetchaccountsUids {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmFetchaccountsUids) Select(query interface{}, args ...interface{}) *OrmHmFetchaccountsUids {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmFetchaccountsUids) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmFetchaccountsUids) And(fuc func(orm *OrmHmFetchaccountsUids)) *OrmHmFetchaccountsUids {
	ormAnd := NewOrmHmFetchaccountsUids()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmFetchaccountsUids) Or(fuc func(orm *OrmHmFetchaccountsUids)) *OrmHmFetchaccountsUids {
	ormOr := NewOrmHmFetchaccountsUids()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmFetchaccountsUids) Preload(query string, args ...interface{}) *OrmHmFetchaccountsUids {
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
func (orm *OrmHmFetchaccountsUids) Joins(query string, args ...interface{}) *OrmHmFetchaccountsUids {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmFetchaccountsUids) WhereUidid(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidid` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) InsertGetUidid(row *HmFetchaccountsUids) int32 {
	orm.db.Create(row)
	return row.Uidid
}
func (orm *OrmHmFetchaccountsUids) WhereUididIn(val []int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidid` IN ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUididGt(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidid` > ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUididGte(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidid` >= ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUididLt(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidid` < ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUididLte(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidid` <= ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidfaid(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidfaid` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidfaidIn(val []int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidfaid` IN ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidfaidGt(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidfaid` > ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidfaidGte(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidfaid` >= ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidfaidLt(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidfaid` < ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidfaidLte(val int32) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidfaid` <= ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidvalue(val string) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidvalue` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidvalueIn(val []string) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidvalue` IN ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidvalueNe(val string) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidvalue` <> ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidtime(val database.Time) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidtime` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidtimeIn(val []database.Time) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidtime` IN ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidtimeNe(val database.Time) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidtime` <> ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidtimeBetween(begin database.Time, end database.Time) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidtime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidtimeLte(val database.Time) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidtime` <= ?", val)
	return orm
}
func (orm *OrmHmFetchaccountsUids) WhereUidtimeGte(val database.Time) *OrmHmFetchaccountsUids {
	orm.db.Where("`uidtime` >= ?", val)
	return orm
}

type HmFetchaccountsUidsList []*HmFetchaccountsUids

func (l HmFetchaccountsUidsList) GetUididList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Uidid)
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUididMap() map[int32]*HmFetchaccountsUids {
	got := make(map[int32]*HmFetchaccountsUids)
	for _, val := range l {
		got[val.Uidid] = val
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUidfaidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Uidfaid)
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUidfaidMap() map[int32]*HmFetchaccountsUids {
	got := make(map[int32]*HmFetchaccountsUids)
	for _, val := range l {
		got[val.Uidfaid] = val
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUidvalueList() []string {
	got := make([]string, 0)
	for _, val := range l {
		got = append(got, val.Uidvalue)
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUidvalueMap() map[string]*HmFetchaccountsUids {
	got := make(map[string]*HmFetchaccountsUids)
	for _, val := range l {
		got[val.Uidvalue] = val
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUidtimeList() []database.Time {
	got := make([]database.Time, 0)
	for _, val := range l {
		got = append(got, val.Uidtime)
	}
	return got
}
func (l HmFetchaccountsUidsList) GetUidtimeMap() map[database.Time]*HmFetchaccountsUids {
	got := make(map[database.Time]*HmFetchaccountsUids)
	for _, val := range l {
		got[val.Uidtime] = val
	}
	return got
}
