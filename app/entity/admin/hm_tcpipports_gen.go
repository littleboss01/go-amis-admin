package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmTcpipports struct {
	Portid                 int64  `gorm:"column:portid;autoIncrement;type:bigint(20);primaryKey;index:portid,class:BTREE,unique" json:"portid"` //
	Portprotocol           int32  `gorm:"column:portprotocol;type:tinyint(4)" json:"portprotocol"`                                              //
	Portnumber             int32  `gorm:"column:portnumber;type:int(11)" json:"portnumber"`                                                     //
	Portaddress1           int64  `gorm:"column:portaddress1;type:bigint(20)" json:"portaddress1"`                                              //
	Portaddress2           *int64 `gorm:"column:portaddress2;type:bigint(20)" json:"portaddress2"`                                              //
	Portconnectionsecurity int32  `gorm:"column:portconnectionsecurity;type:tinyint(4)" json:"portconnectionsecurity"`                          //
	Portsslcertificateid   int64  `gorm:"column:portsslcertificateid;type:bigint(20)" json:"portsslcertificateid"`                              //
}

var (
	HmTcpipportsFieldPortid                 = "portid"
	HmTcpipportsFieldPortprotocol           = "portprotocol"
	HmTcpipportsFieldPortnumber             = "portnumber"
	HmTcpipportsFieldPortaddress1           = "portaddress1"
	HmTcpipportsFieldPortaddress2           = "portaddress2"
	HmTcpipportsFieldPortconnectionsecurity = "portconnectionsecurity"
	HmTcpipportsFieldPortsslcertificateid   = "portsslcertificateid"
)

func (receiver *HmTcpipports) TableName() string {
	return "hm_tcpipports"
}

type OrmHmTcpipports struct {
	db *gorm.DB
}

func (orm *OrmHmTcpipports) TableName() string {
	return "hm_tcpipports"
}

func NewOrmHmTcpipports(txs ...*gorm.DB) *OrmHmTcpipports {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmTcpipports{})
	} else {
		tx = txs[0].Model(&HmTcpipports{})
	}
	return &OrmHmTcpipports{db: tx}
}

func (orm *OrmHmTcpipports) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmTcpipports) GetTableInfo() interface{} {
	return &HmTcpipports{}
}

// Create insert the value into database
func (orm *OrmHmTcpipports) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmTcpipports) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmTcpipports) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmTcpipports) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmTcpipports) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmTcpipports) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmTcpipports) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmTcpipports) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmTcpipports) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmTcpipports) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmTcpipports) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmTcpipports) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmTcpipports) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmTcpipports) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmTcpipports) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmTcpipports) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmTcpipports) Unscoped() *OrmHmTcpipports {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmTcpipports) Insert(row *HmTcpipports) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmTcpipports) Inserts(rows []*HmTcpipports) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmTcpipports) Order(value interface{}) *OrmHmTcpipports {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmTcpipports) Group(name string) *OrmHmTcpipports {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmTcpipports) Limit(limit int) *OrmHmTcpipports {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmTcpipports) Offset(offset int) *OrmHmTcpipports {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmTcpipports) Get() HmTcpipportsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmTcpipports) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmTcpipports) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmTcpipports{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmTcpipports) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_tcpipports")
}

func (orm *OrmHmTcpipports) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmTcpipports) First(conds ...interface{}) (*HmTcpipports, bool) {
	dest := &HmTcpipports{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmTcpipports) Take(conds ...interface{}) (*HmTcpipports, int64) {
	dest := &HmTcpipports{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmTcpipports) Last(conds ...interface{}) (*HmTcpipports, int64) {
	dest := &HmTcpipports{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmTcpipports) Find(conds ...interface{}) (HmTcpipportsList, int64) {
	list := make([]*HmTcpipports, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmTcpipports) Paginate(page int, limit int) (HmTcpipportsList, int64) {
	var total int64
	list := make([]*HmTcpipports, 0)
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
func (orm *OrmHmTcpipports) SimplePaginate(page int, limit int) HmTcpipportsList {
	list := make([]*HmTcpipports, 0)
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
func (orm *OrmHmTcpipports) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmTcpipports) FirstOrInit(dest *HmTcpipports, conds ...interface{}) (*HmTcpipports, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmTcpipports) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmTcpipports) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmTcpipports) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmTcpipports) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmTcpipports) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmTcpipports) Where(query interface{}, args ...interface{}) *OrmHmTcpipports {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmTcpipports) Select(query interface{}, args ...interface{}) *OrmHmTcpipports {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmTcpipports) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmTcpipports) And(fuc func(orm *OrmHmTcpipports)) *OrmHmTcpipports {
	ormAnd := NewOrmHmTcpipports()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmTcpipports) Or(fuc func(orm *OrmHmTcpipports)) *OrmHmTcpipports {
	ormOr := NewOrmHmTcpipports()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmTcpipports) Preload(query string, args ...interface{}) *OrmHmTcpipports {
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
func (orm *OrmHmTcpipports) Joins(query string, args ...interface{}) *OrmHmTcpipports {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmTcpipports) WherePortid(val int64) *OrmHmTcpipports {
	orm.db.Where("`portid` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) InsertGetPortid(row *HmTcpipports) int64 {
	orm.db.Create(row)
	return row.Portid
}
func (orm *OrmHmTcpipports) WherePortidIn(val []int64) *OrmHmTcpipports {
	orm.db.Where("`portid` IN ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortidGt(val int64) *OrmHmTcpipports {
	orm.db.Where("`portid` > ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortidGte(val int64) *OrmHmTcpipports {
	orm.db.Where("`portid` >= ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortidLt(val int64) *OrmHmTcpipports {
	orm.db.Where("`portid` < ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortidLte(val int64) *OrmHmTcpipports {
	orm.db.Where("`portid` <= ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortprotocol(val int32) *OrmHmTcpipports {
	orm.db.Where("`portprotocol` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortnumber(val int32) *OrmHmTcpipports {
	orm.db.Where("`portnumber` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortaddress1(val int64) *OrmHmTcpipports {
	orm.db.Where("`portaddress1` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortaddress2(val int64) *OrmHmTcpipports {
	orm.db.Where("`portaddress2` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortconnectionsecurity(val int32) *OrmHmTcpipports {
	orm.db.Where("`portconnectionsecurity` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortsslcertificateid(val int64) *OrmHmTcpipports {
	orm.db.Where("`portsslcertificateid` = ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortsslcertificateidIn(val []int64) *OrmHmTcpipports {
	orm.db.Where("`portsslcertificateid` IN ?", val)
	return orm
}
func (orm *OrmHmTcpipports) WherePortsslcertificateidNe(val int64) *OrmHmTcpipports {
	orm.db.Where("`portsslcertificateid` <> ?", val)
	return orm
}

type HmTcpipportsList []*HmTcpipports

func (l HmTcpipportsList) GetPortidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Portid)
	}
	return got
}
func (l HmTcpipportsList) GetPortidMap() map[int64]*HmTcpipports {
	got := make(map[int64]*HmTcpipports)
	for _, val := range l {
		got[val.Portid] = val
	}
	return got
}
func (l HmTcpipportsList) GetPortsslcertificateidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Portsslcertificateid)
	}
	return got
}
func (l HmTcpipportsList) GetPortsslcertificateidMap() map[int64]*HmTcpipports {
	got := make(map[int64]*HmTcpipports)
	for _, val := range l {
		got[val.Portsslcertificateid] = val
	}
	return got
}
