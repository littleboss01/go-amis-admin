package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmSurblservers struct {
	Surblid            int32  `gorm:"column:surblid;autoIncrement;type:int(11);primaryKey;index:surblid,class:BTREE,unique" json:"surblid"` //
	Surblactive        int32  `gorm:"column:surblactive;type:tinyint(4)" json:"surblactive"`                                                //
	Surblhost          string `gorm:"column:surblhost;type:varchar(255)" json:"surblhost"`                                                  //
	Surblrejectmessage string `gorm:"column:surblrejectmessage;type:varchar(255)" json:"surblrejectmessage"`                                //
	Surblscore         int32  `gorm:"column:surblscore;type:int(11)" json:"surblscore"`                                                     //
}

var (
	HmSurblserversFieldSurblid            = "surblid"
	HmSurblserversFieldSurblactive        = "surblactive"
	HmSurblserversFieldSurblhost          = "surblhost"
	HmSurblserversFieldSurblrejectmessage = "surblrejectmessage"
	HmSurblserversFieldSurblscore         = "surblscore"
)

func (receiver *HmSurblservers) TableName() string {
	return "hm_surblservers"
}

type OrmHmSurblservers struct {
	db *gorm.DB
}

func (orm *OrmHmSurblservers) TableName() string {
	return "hm_surblservers"
}

func NewOrmHmSurblservers(txs ...*gorm.DB) *OrmHmSurblservers {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmSurblservers{})
	} else {
		tx = txs[0].Model(&HmSurblservers{})
	}
	return &OrmHmSurblservers{db: tx}
}

func (orm *OrmHmSurblservers) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmSurblservers) GetTableInfo() interface{} {
	return &HmSurblservers{}
}

// Create insert the value into database
func (orm *OrmHmSurblservers) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmSurblservers) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmSurblservers) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmSurblservers) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmSurblservers) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmSurblservers) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmSurblservers) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmSurblservers) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmSurblservers) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmSurblservers) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmSurblservers) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmSurblservers) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmSurblservers) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmSurblservers) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmSurblservers) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmSurblservers) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmSurblservers) Unscoped() *OrmHmSurblservers {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmSurblservers) Insert(row *HmSurblservers) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmSurblservers) Inserts(rows []*HmSurblservers) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmSurblservers) Order(value interface{}) *OrmHmSurblservers {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmSurblservers) Group(name string) *OrmHmSurblservers {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmSurblservers) Limit(limit int) *OrmHmSurblservers {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmSurblservers) Offset(offset int) *OrmHmSurblservers {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmSurblservers) Get() HmSurblserversList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmSurblservers) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmSurblservers) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmSurblservers{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmSurblservers) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_surblservers")
}

func (orm *OrmHmSurblservers) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmSurblservers) First(conds ...interface{}) (*HmSurblservers, bool) {
	dest := &HmSurblservers{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmSurblservers) Take(conds ...interface{}) (*HmSurblservers, int64) {
	dest := &HmSurblservers{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmSurblservers) Last(conds ...interface{}) (*HmSurblservers, int64) {
	dest := &HmSurblservers{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmSurblservers) Find(conds ...interface{}) (HmSurblserversList, int64) {
	list := make([]*HmSurblservers, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmSurblservers) Paginate(page int, limit int) (HmSurblserversList, int64) {
	var total int64
	list := make([]*HmSurblservers, 0)
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
func (orm *OrmHmSurblservers) SimplePaginate(page int, limit int) HmSurblserversList {
	list := make([]*HmSurblservers, 0)
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
func (orm *OrmHmSurblservers) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmSurblservers) FirstOrInit(dest *HmSurblservers, conds ...interface{}) (*HmSurblservers, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmSurblservers) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSurblservers) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSurblservers) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmSurblservers) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmSurblservers) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmSurblservers) Where(query interface{}, args ...interface{}) *OrmHmSurblservers {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmSurblservers) Select(query interface{}, args ...interface{}) *OrmHmSurblservers {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmSurblservers) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmSurblservers) And(fuc func(orm *OrmHmSurblservers)) *OrmHmSurblservers {
	ormAnd := NewOrmHmSurblservers()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmSurblservers) Or(fuc func(orm *OrmHmSurblservers)) *OrmHmSurblservers {
	ormOr := NewOrmHmSurblservers()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmSurblservers) Preload(query string, args ...interface{}) *OrmHmSurblservers {
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
func (orm *OrmHmSurblservers) Joins(query string, args ...interface{}) *OrmHmSurblservers {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmSurblservers) WhereSurblid(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblid` = ?", val)
	return orm
}
func (orm *OrmHmSurblservers) InsertGetSurblid(row *HmSurblservers) int32 {
	orm.db.Create(row)
	return row.Surblid
}
func (orm *OrmHmSurblservers) WhereSurblidIn(val []int32) *OrmHmSurblservers {
	orm.db.Where("`surblid` IN ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblidGt(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblid` > ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblidGte(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblid` >= ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblidLt(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblid` < ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblidLte(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblid` <= ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblactive(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblactive` = ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblhost(val string) *OrmHmSurblservers {
	orm.db.Where("`surblhost` = ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblrejectmessage(val string) *OrmHmSurblservers {
	orm.db.Where("`surblrejectmessage` = ?", val)
	return orm
}
func (orm *OrmHmSurblservers) WhereSurblscore(val int32) *OrmHmSurblservers {
	orm.db.Where("`surblscore` = ?", val)
	return orm
}

type HmSurblserversList []*HmSurblservers

func (l HmSurblserversList) GetSurblidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Surblid)
	}
	return got
}
func (l HmSurblserversList) GetSurblidMap() map[int32]*HmSurblservers {
	got := make(map[int32]*HmSurblservers)
	for _, val := range l {
		got[val.Surblid] = val
	}
	return got
}
