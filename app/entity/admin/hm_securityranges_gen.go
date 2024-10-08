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

type HmSecurityranges struct {
	Rangeid          int32         `gorm:"column:rangeid;autoIncrement;type:int(11);primaryKey;index:rangeid,class:BTREE,unique" json:"rangeid"` //
	Rangepriorityid  int32         `gorm:"column:rangepriorityid;type:int(11)" json:"rangepriorityid"`                                           //
	Rangelowerip1    int64         `gorm:"column:rangelowerip1;type:bigint(20)" json:"rangelowerip1"`                                            //
	Rangelowerip2    *int64        `gorm:"column:rangelowerip2;type:bigint(20)" json:"rangelowerip2"`                                            //
	Rangeupperip1    int64         `gorm:"column:rangeupperip1;type:bigint(20)" json:"rangeupperip1"`                                            //
	Rangeupperip2    *int64        `gorm:"column:rangeupperip2;type:bigint(20)" json:"rangeupperip2"`                                            //
	Rangeoptions     int32         `gorm:"column:rangeoptions;type:int(11)" json:"rangeoptions"`                                                 //
	Rangename        string        `gorm:"column:rangename;type:varchar(100);index:rangename,class:BTREE,unique" json:"rangename"`               //
	Rangeexpires     int32         `gorm:"column:rangeexpires;type:tinyint(4)" json:"rangeexpires"`                                              //
	Rangeexpirestime database.Time `gorm:"column:rangeexpirestime;type:datetime" json:"rangeexpirestime"`                                        //
}

var (
	HmSecurityrangesFieldRangeid          = "rangeid"
	HmSecurityrangesFieldRangepriorityid  = "rangepriorityid"
	HmSecurityrangesFieldRangelowerip1    = "rangelowerip1"
	HmSecurityrangesFieldRangelowerip2    = "rangelowerip2"
	HmSecurityrangesFieldRangeupperip1    = "rangeupperip1"
	HmSecurityrangesFieldRangeupperip2    = "rangeupperip2"
	HmSecurityrangesFieldRangeoptions     = "rangeoptions"
	HmSecurityrangesFieldRangename        = "rangename"
	HmSecurityrangesFieldRangeexpires     = "rangeexpires"
	HmSecurityrangesFieldRangeexpirestime = "rangeexpirestime"
)

func (receiver *HmSecurityranges) TableName() string {
	return "hm_securityranges"
}

type OrmHmSecurityranges struct {
	db *gorm.DB
}

func (orm *OrmHmSecurityranges) TableName() string {
	return "hm_securityranges"
}

func NewOrmHmSecurityranges(txs ...*gorm.DB) *OrmHmSecurityranges {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmSecurityranges{})
	} else {
		tx = txs[0].Model(&HmSecurityranges{})
	}
	return &OrmHmSecurityranges{db: tx}
}

func (orm *OrmHmSecurityranges) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmSecurityranges) GetTableInfo() interface{} {
	return &HmSecurityranges{}
}

// Create insert the value into database
func (orm *OrmHmSecurityranges) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmSecurityranges) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmSecurityranges) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmSecurityranges) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmSecurityranges) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmSecurityranges) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmSecurityranges) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmSecurityranges) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmSecurityranges) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmSecurityranges) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmSecurityranges) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmSecurityranges) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmSecurityranges) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmSecurityranges) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmSecurityranges) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmSecurityranges) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmSecurityranges) Unscoped() *OrmHmSecurityranges {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmSecurityranges) Insert(row *HmSecurityranges) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmSecurityranges) Inserts(rows []*HmSecurityranges) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmSecurityranges) Order(value interface{}) *OrmHmSecurityranges {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmSecurityranges) Group(name string) *OrmHmSecurityranges {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmSecurityranges) Limit(limit int) *OrmHmSecurityranges {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmSecurityranges) Offset(offset int) *OrmHmSecurityranges {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmSecurityranges) Get() HmSecurityrangesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmSecurityranges) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmSecurityranges) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmSecurityranges{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmSecurityranges) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_securityranges")
}

func (orm *OrmHmSecurityranges) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmSecurityranges) First(conds ...interface{}) (*HmSecurityranges, bool) {
	dest := &HmSecurityranges{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmSecurityranges) Take(conds ...interface{}) (*HmSecurityranges, int64) {
	dest := &HmSecurityranges{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmSecurityranges) Last(conds ...interface{}) (*HmSecurityranges, int64) {
	dest := &HmSecurityranges{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmSecurityranges) Find(conds ...interface{}) (HmSecurityrangesList, int64) {
	list := make([]*HmSecurityranges, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmSecurityranges) Paginate(page int, limit int) (HmSecurityrangesList, int64) {
	var total int64
	list := make([]*HmSecurityranges, 0)
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
func (orm *OrmHmSecurityranges) SimplePaginate(page int, limit int) HmSecurityrangesList {
	list := make([]*HmSecurityranges, 0)
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
func (orm *OrmHmSecurityranges) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmSecurityranges) FirstOrInit(dest *HmSecurityranges, conds ...interface{}) (*HmSecurityranges, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmSecurityranges) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSecurityranges) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSecurityranges) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmSecurityranges) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmSecurityranges) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmSecurityranges) Where(query interface{}, args ...interface{}) *OrmHmSecurityranges {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmSecurityranges) Select(query interface{}, args ...interface{}) *OrmHmSecurityranges {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmSecurityranges) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmSecurityranges) And(fuc func(orm *OrmHmSecurityranges)) *OrmHmSecurityranges {
	ormAnd := NewOrmHmSecurityranges()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmSecurityranges) Or(fuc func(orm *OrmHmSecurityranges)) *OrmHmSecurityranges {
	ormOr := NewOrmHmSecurityranges()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmSecurityranges) Preload(query string, args ...interface{}) *OrmHmSecurityranges {
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
func (orm *OrmHmSecurityranges) Joins(query string, args ...interface{}) *OrmHmSecurityranges {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmSecurityranges) WhereRangeid(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeid` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) InsertGetRangeid(row *HmSecurityranges) int32 {
	orm.db.Create(row)
	return row.Rangeid
}
func (orm *OrmHmSecurityranges) WhereRangeidIn(val []int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeid` IN ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeidGt(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeid` > ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeidGte(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeid` >= ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeidLt(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeid` < ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeidLte(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeid` <= ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangepriorityid(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangepriorityid` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangepriorityidIn(val []int32) *OrmHmSecurityranges {
	orm.db.Where("`rangepriorityid` IN ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangepriorityidNe(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangepriorityid` <> ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangelowerip1(val int64) *OrmHmSecurityranges {
	orm.db.Where("`rangelowerip1` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangelowerip2(val int64) *OrmHmSecurityranges {
	orm.db.Where("`rangelowerip2` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeupperip1(val int64) *OrmHmSecurityranges {
	orm.db.Where("`rangeupperip1` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeupperip2(val int64) *OrmHmSecurityranges {
	orm.db.Where("`rangeupperip2` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeoptions(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeoptions` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangename(val string) *OrmHmSecurityranges {
	orm.db.Where("`rangename` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangenameIn(val []string) *OrmHmSecurityranges {
	orm.db.Where("`rangename` IN ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangenameGt(val string) *OrmHmSecurityranges {
	orm.db.Where("`rangename` > ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangenameGte(val string) *OrmHmSecurityranges {
	orm.db.Where("`rangename` >= ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangenameLt(val string) *OrmHmSecurityranges {
	orm.db.Where("`rangename` < ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangenameLte(val string) *OrmHmSecurityranges {
	orm.db.Where("`rangename` <= ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeexpires(val int32) *OrmHmSecurityranges {
	orm.db.Where("`rangeexpires` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeexpirestime(val database.Time) *OrmHmSecurityranges {
	orm.db.Where("`rangeexpirestime` = ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeexpirestimeBetween(begin database.Time, end database.Time) *OrmHmSecurityranges {
	orm.db.Where("`rangeexpirestime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeexpirestimeLte(val database.Time) *OrmHmSecurityranges {
	orm.db.Where("`rangeexpirestime` <= ?", val)
	return orm
}
func (orm *OrmHmSecurityranges) WhereRangeexpirestimeGte(val database.Time) *OrmHmSecurityranges {
	orm.db.Where("`rangeexpirestime` >= ?", val)
	return orm
}

type HmSecurityrangesList []*HmSecurityranges

func (l HmSecurityrangesList) GetRangeidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Rangeid)
	}
	return got
}
func (l HmSecurityrangesList) GetRangeidMap() map[int32]*HmSecurityranges {
	got := make(map[int32]*HmSecurityranges)
	for _, val := range l {
		got[val.Rangeid] = val
	}
	return got
}
func (l HmSecurityrangesList) GetRangepriorityidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Rangepriorityid)
	}
	return got
}
func (l HmSecurityrangesList) GetRangepriorityidMap() map[int32]*HmSecurityranges {
	got := make(map[int32]*HmSecurityranges)
	for _, val := range l {
		got[val.Rangepriorityid] = val
	}
	return got
}
