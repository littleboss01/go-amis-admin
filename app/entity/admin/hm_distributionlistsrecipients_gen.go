package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmDistributionlistsrecipients struct {
	Distributionlistrecipientid      int32   `gorm:"column:distributionlistrecipientid;autoIncrement;type:int(11);primaryKey;index:distributionlistrecipientid,class:BTREE,unique" json:"distributionlistrecipientid"` //
	Distributionlistrecipientlistid  int32   `gorm:"column:distributionlistrecipientlistid;type:int(11);index:idx_hm_distributionlistsrecipients,class:BTREE" json:"distributionlistrecipientlistid"`                  //
	Distributionlistrecipientaddress *string `gorm:"column:distributionlistrecipientaddress;type:varchar(255)" json:"distributionlistrecipientaddress"`                                                                //
}

var (
	HmDistributionlistsrecipientsFieldDistributionlistrecipientid      = "distributionlistrecipientid"
	HmDistributionlistsrecipientsFieldDistributionlistrecipientlistid  = "distributionlistrecipientlistid"
	HmDistributionlistsrecipientsFieldDistributionlistrecipientaddress = "distributionlistrecipientaddress"
)

func (receiver *HmDistributionlistsrecipients) TableName() string {
	return "hm_distributionlistsrecipients"
}

type OrmHmDistributionlistsrecipients struct {
	db *gorm.DB
}

func (orm *OrmHmDistributionlistsrecipients) TableName() string {
	return "hm_distributionlistsrecipients"
}

func NewOrmHmDistributionlistsrecipients(txs ...*gorm.DB) *OrmHmDistributionlistsrecipients {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmDistributionlistsrecipients{})
	} else {
		tx = txs[0].Model(&HmDistributionlistsrecipients{})
	}
	return &OrmHmDistributionlistsrecipients{db: tx}
}

func (orm *OrmHmDistributionlistsrecipients) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmDistributionlistsrecipients) GetTableInfo() interface{} {
	return &HmDistributionlistsrecipients{}
}

// Create insert the value into database
func (orm *OrmHmDistributionlistsrecipients) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmDistributionlistsrecipients) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmDistributionlistsrecipients) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmDistributionlistsrecipients) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmDistributionlistsrecipients) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmDistributionlistsrecipients) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmDistributionlistsrecipients) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmDistributionlistsrecipients) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmDistributionlistsrecipients) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmDistributionlistsrecipients) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmDistributionlistsrecipients) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmDistributionlistsrecipients) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmDistributionlistsrecipients) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmDistributionlistsrecipients) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmDistributionlistsrecipients) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmDistributionlistsrecipients) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmDistributionlistsrecipients) Unscoped() *OrmHmDistributionlistsrecipients {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmDistributionlistsrecipients) Insert(row *HmDistributionlistsrecipients) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmDistributionlistsrecipients) Inserts(rows []*HmDistributionlistsrecipients) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmDistributionlistsrecipients) Order(value interface{}) *OrmHmDistributionlistsrecipients {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) Group(name string) *OrmHmDistributionlistsrecipients {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) Limit(limit int) *OrmHmDistributionlistsrecipients {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) Offset(offset int) *OrmHmDistributionlistsrecipients {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmDistributionlistsrecipients) Get() HmDistributionlistsrecipientsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmDistributionlistsrecipients) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmDistributionlistsrecipients) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmDistributionlistsrecipients{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmDistributionlistsrecipients) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_distributionlistsrecipients")
}

func (orm *OrmHmDistributionlistsrecipients) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmDistributionlistsrecipients) First(conds ...interface{}) (*HmDistributionlistsrecipients, bool) {
	dest := &HmDistributionlistsrecipients{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmDistributionlistsrecipients) Take(conds ...interface{}) (*HmDistributionlistsrecipients, int64) {
	dest := &HmDistributionlistsrecipients{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmDistributionlistsrecipients) Last(conds ...interface{}) (*HmDistributionlistsrecipients, int64) {
	dest := &HmDistributionlistsrecipients{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmDistributionlistsrecipients) Find(conds ...interface{}) (HmDistributionlistsrecipientsList, int64) {
	list := make([]*HmDistributionlistsrecipients, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmDistributionlistsrecipients) Paginate(page int, limit int) (HmDistributionlistsrecipientsList, int64) {
	var total int64
	list := make([]*HmDistributionlistsrecipients, 0)
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
func (orm *OrmHmDistributionlistsrecipients) SimplePaginate(page int, limit int) HmDistributionlistsrecipientsList {
	list := make([]*HmDistributionlistsrecipients, 0)
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
func (orm *OrmHmDistributionlistsrecipients) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmDistributionlistsrecipients) FirstOrInit(dest *HmDistributionlistsrecipients, conds ...interface{}) (*HmDistributionlistsrecipients, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmDistributionlistsrecipients) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDistributionlistsrecipients) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDistributionlistsrecipients) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmDistributionlistsrecipients) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmDistributionlistsrecipients) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmDistributionlistsrecipients) Where(query interface{}, args ...interface{}) *OrmHmDistributionlistsrecipients {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) Select(query interface{}, args ...interface{}) *OrmHmDistributionlistsrecipients {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmDistributionlistsrecipients) And(fuc func(orm *OrmHmDistributionlistsrecipients)) *OrmHmDistributionlistsrecipients {
	ormAnd := NewOrmHmDistributionlistsrecipients()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) Or(fuc func(orm *OrmHmDistributionlistsrecipients)) *OrmHmDistributionlistsrecipients {
	ormOr := NewOrmHmDistributionlistsrecipients()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmDistributionlistsrecipients) Preload(query string, args ...interface{}) *OrmHmDistributionlistsrecipients {
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
func (orm *OrmHmDistributionlistsrecipients) Joins(query string, args ...interface{}) *OrmHmDistributionlistsrecipients {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientid(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientid` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) InsertGetDistributionlistrecipientid(row *HmDistributionlistsrecipients) int32 {
	orm.db.Create(row)
	return row.Distributionlistrecipientid
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientidIn(val []int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientid` IN ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientidGt(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientid` > ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientidGte(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientid` >= ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientidLt(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientid` < ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientidLte(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientid` <= ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientlistid(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientlistid` = ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientlistidIn(val []int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientlistid` IN ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientlistidGt(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientlistid` > ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientlistidGte(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientlistid` >= ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientlistidLt(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientlistid` < ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientlistidLte(val int32) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientlistid` <= ?", val)
	return orm
}
func (orm *OrmHmDistributionlistsrecipients) WhereDistributionlistrecipientaddress(val string) *OrmHmDistributionlistsrecipients {
	orm.db.Where("`distributionlistrecipientaddress` = ?", val)
	return orm
}

type HmDistributionlistsrecipientsList []*HmDistributionlistsrecipients

func (l HmDistributionlistsrecipientsList) GetDistributionlistrecipientidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Distributionlistrecipientid)
	}
	return got
}
func (l HmDistributionlistsrecipientsList) GetDistributionlistrecipientidMap() map[int32]*HmDistributionlistsrecipients {
	got := make(map[int32]*HmDistributionlistsrecipients)
	for _, val := range l {
		got[val.Distributionlistrecipientid] = val
	}
	return got
}
func (l HmDistributionlistsrecipientsList) GetDistributionlistrecipientlistidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Distributionlistrecipientlistid)
	}
	return got
}
func (l HmDistributionlistsrecipientsList) GetDistributionlistrecipientlistidMap() map[int32]*HmDistributionlistsrecipients {
	got := make(map[int32]*HmDistributionlistsrecipients)
	for _, val := range l {
		got[val.Distributionlistrecipientlistid] = val
	}
	return got
}
