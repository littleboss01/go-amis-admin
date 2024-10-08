package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmMessagerecipients struct {
	Recipientid              int64  `gorm:"column:recipientid;autoIncrement;type:bigint(20);primaryKey;index:recipientid,class:BTREE,unique" json:"recipientid"` //
	Recipientmessageid       int64  `gorm:"column:recipientmessageid;type:bigint(20);index:idx_hm_messagerecipients,class:BTREE" json:"recipientmessageid"`      //
	Recipientaddress         string `gorm:"column:recipientaddress;type:varchar(255)" json:"recipientaddress"`                                                   //
	Recipientlocalaccountid  int32  `gorm:"column:recipientlocalaccountid;type:int(11)" json:"recipientlocalaccountid"`                                          //
	Recipientoriginaladdress string `gorm:"column:recipientoriginaladdress;type:varchar(255)" json:"recipientoriginaladdress"`                                   //
}

var (
	HmMessagerecipientsFieldRecipientid              = "recipientid"
	HmMessagerecipientsFieldRecipientmessageid       = "recipientmessageid"
	HmMessagerecipientsFieldRecipientaddress         = "recipientaddress"
	HmMessagerecipientsFieldRecipientlocalaccountid  = "recipientlocalaccountid"
	HmMessagerecipientsFieldRecipientoriginaladdress = "recipientoriginaladdress"
)

func (receiver *HmMessagerecipients) TableName() string {
	return "hm_messagerecipients"
}

type OrmHmMessagerecipients struct {
	db *gorm.DB
}

func (orm *OrmHmMessagerecipients) TableName() string {
	return "hm_messagerecipients"
}

func NewOrmHmMessagerecipients(txs ...*gorm.DB) *OrmHmMessagerecipients {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmMessagerecipients{})
	} else {
		tx = txs[0].Model(&HmMessagerecipients{})
	}
	return &OrmHmMessagerecipients{db: tx}
}

func (orm *OrmHmMessagerecipients) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmMessagerecipients) GetTableInfo() interface{} {
	return &HmMessagerecipients{}
}

// Create insert the value into database
func (orm *OrmHmMessagerecipients) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmMessagerecipients) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmMessagerecipients) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmMessagerecipients) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmMessagerecipients) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmMessagerecipients) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmMessagerecipients) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmMessagerecipients) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmMessagerecipients) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmMessagerecipients) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmMessagerecipients) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmMessagerecipients) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmMessagerecipients) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmMessagerecipients) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmMessagerecipients) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmMessagerecipients) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmMessagerecipients) Unscoped() *OrmHmMessagerecipients {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmMessagerecipients) Insert(row *HmMessagerecipients) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmMessagerecipients) Inserts(rows []*HmMessagerecipients) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmMessagerecipients) Order(value interface{}) *OrmHmMessagerecipients {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmMessagerecipients) Group(name string) *OrmHmMessagerecipients {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmMessagerecipients) Limit(limit int) *OrmHmMessagerecipients {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmMessagerecipients) Offset(offset int) *OrmHmMessagerecipients {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmMessagerecipients) Get() HmMessagerecipientsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmMessagerecipients) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmMessagerecipients) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmMessagerecipients{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmMessagerecipients) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_messagerecipients")
}

func (orm *OrmHmMessagerecipients) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmMessagerecipients) First(conds ...interface{}) (*HmMessagerecipients, bool) {
	dest := &HmMessagerecipients{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmMessagerecipients) Take(conds ...interface{}) (*HmMessagerecipients, int64) {
	dest := &HmMessagerecipients{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmMessagerecipients) Last(conds ...interface{}) (*HmMessagerecipients, int64) {
	dest := &HmMessagerecipients{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmMessagerecipients) Find(conds ...interface{}) (HmMessagerecipientsList, int64) {
	list := make([]*HmMessagerecipients, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmMessagerecipients) Paginate(page int, limit int) (HmMessagerecipientsList, int64) {
	var total int64
	list := make([]*HmMessagerecipients, 0)
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
func (orm *OrmHmMessagerecipients) SimplePaginate(page int, limit int) HmMessagerecipientsList {
	list := make([]*HmMessagerecipients, 0)
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
func (orm *OrmHmMessagerecipients) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmMessagerecipients) FirstOrInit(dest *HmMessagerecipients, conds ...interface{}) (*HmMessagerecipients, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmMessagerecipients) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmMessagerecipients) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmMessagerecipients) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmMessagerecipients) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmMessagerecipients) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmMessagerecipients) Where(query interface{}, args ...interface{}) *OrmHmMessagerecipients {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmMessagerecipients) Select(query interface{}, args ...interface{}) *OrmHmMessagerecipients {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmMessagerecipients) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmMessagerecipients) And(fuc func(orm *OrmHmMessagerecipients)) *OrmHmMessagerecipients {
	ormAnd := NewOrmHmMessagerecipients()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmMessagerecipients) Or(fuc func(orm *OrmHmMessagerecipients)) *OrmHmMessagerecipients {
	ormOr := NewOrmHmMessagerecipients()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmMessagerecipients) Preload(query string, args ...interface{}) *OrmHmMessagerecipients {
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
func (orm *OrmHmMessagerecipients) Joins(query string, args ...interface{}) *OrmHmMessagerecipients {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmMessagerecipients) WhereRecipientid(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientid` = ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) InsertGetRecipientid(row *HmMessagerecipients) int64 {
	orm.db.Create(row)
	return row.Recipientid
}
func (orm *OrmHmMessagerecipients) WhereRecipientidIn(val []int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientid` IN ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientidGt(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientid` > ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientidGte(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientid` >= ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientidLt(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientid` < ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientidLte(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientid` <= ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientmessageid(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientmessageid` = ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientmessageidIn(val []int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientmessageid` IN ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientmessageidGt(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientmessageid` > ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientmessageidGte(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientmessageid` >= ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientmessageidLt(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientmessageid` < ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientmessageidLte(val int64) *OrmHmMessagerecipients {
	orm.db.Where("`recipientmessageid` <= ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientaddress(val string) *OrmHmMessagerecipients {
	orm.db.Where("`recipientaddress` = ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientlocalaccountid(val int32) *OrmHmMessagerecipients {
	orm.db.Where("`recipientlocalaccountid` = ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientlocalaccountidIn(val []int32) *OrmHmMessagerecipients {
	orm.db.Where("`recipientlocalaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientlocalaccountidNe(val int32) *OrmHmMessagerecipients {
	orm.db.Where("`recipientlocalaccountid` <> ?", val)
	return orm
}
func (orm *OrmHmMessagerecipients) WhereRecipientoriginaladdress(val string) *OrmHmMessagerecipients {
	orm.db.Where("`recipientoriginaladdress` = ?", val)
	return orm
}

type HmMessagerecipientsList []*HmMessagerecipients

func (l HmMessagerecipientsList) GetRecipientidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Recipientid)
	}
	return got
}
func (l HmMessagerecipientsList) GetRecipientidMap() map[int64]*HmMessagerecipients {
	got := make(map[int64]*HmMessagerecipients)
	for _, val := range l {
		got[val.Recipientid] = val
	}
	return got
}
func (l HmMessagerecipientsList) GetRecipientmessageidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Recipientmessageid)
	}
	return got
}
func (l HmMessagerecipientsList) GetRecipientmessageidMap() map[int64]*HmMessagerecipients {
	got := make(map[int64]*HmMessagerecipients)
	for _, val := range l {
		got[val.Recipientmessageid] = val
	}
	return got
}
func (l HmMessagerecipientsList) GetRecipientlocalaccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Recipientlocalaccountid)
	}
	return got
}
func (l HmMessagerecipientsList) GetRecipientlocalaccountidMap() map[int32]*HmMessagerecipients {
	got := make(map[int32]*HmMessagerecipients)
	for _, val := range l {
		got[val.Recipientlocalaccountid] = val
	}
	return got
}
