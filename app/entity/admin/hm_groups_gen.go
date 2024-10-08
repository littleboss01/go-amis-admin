package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmGroups struct {
	Groupid   int64   `gorm:"column:groupid;autoIncrement;type:bigint(20);primaryKey;index:groupid,class:BTREE,unique" json:"groupid"` //
	Groupname *string `gorm:"column:groupname;type:varchar(255)" json:"groupname"`                                                     //
}

var (
	HmGroupsFieldGroupid   = "groupid"
	HmGroupsFieldGroupname = "groupname"
)

func (receiver *HmGroups) TableName() string {
	return "hm_groups"
}

type OrmHmGroups struct {
	db *gorm.DB
}

func (orm *OrmHmGroups) TableName() string {
	return "hm_groups"
}

func NewOrmHmGroups(txs ...*gorm.DB) *OrmHmGroups {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmGroups{})
	} else {
		tx = txs[0].Model(&HmGroups{})
	}
	return &OrmHmGroups{db: tx}
}

func (orm *OrmHmGroups) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmGroups) GetTableInfo() interface{} {
	return &HmGroups{}
}

// Create insert the value into database
func (orm *OrmHmGroups) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmGroups) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmGroups) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmGroups) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmGroups) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmGroups) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmGroups) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmGroups) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmGroups) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmGroups) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmGroups) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmGroups) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmGroups) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmGroups) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmGroups) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmGroups) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmGroups) Unscoped() *OrmHmGroups {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmGroups) Insert(row *HmGroups) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmGroups) Inserts(rows []*HmGroups) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmGroups) Order(value interface{}) *OrmHmGroups {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmGroups) Group(name string) *OrmHmGroups {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmGroups) Limit(limit int) *OrmHmGroups {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmGroups) Offset(offset int) *OrmHmGroups {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmGroups) Get() HmGroupsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmGroups) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmGroups) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmGroups{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmGroups) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_groups")
}

func (orm *OrmHmGroups) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmGroups) First(conds ...interface{}) (*HmGroups, bool) {
	dest := &HmGroups{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmGroups) Take(conds ...interface{}) (*HmGroups, int64) {
	dest := &HmGroups{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmGroups) Last(conds ...interface{}) (*HmGroups, int64) {
	dest := &HmGroups{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmGroups) Find(conds ...interface{}) (HmGroupsList, int64) {
	list := make([]*HmGroups, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmGroups) Paginate(page int, limit int) (HmGroupsList, int64) {
	var total int64
	list := make([]*HmGroups, 0)
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
func (orm *OrmHmGroups) SimplePaginate(page int, limit int) HmGroupsList {
	list := make([]*HmGroups, 0)
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
func (orm *OrmHmGroups) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmGroups) FirstOrInit(dest *HmGroups, conds ...interface{}) (*HmGroups, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmGroups) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGroups) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGroups) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmGroups) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmGroups) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmGroups) Where(query interface{}, args ...interface{}) *OrmHmGroups {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmGroups) Select(query interface{}, args ...interface{}) *OrmHmGroups {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmGroups) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmGroups) And(fuc func(orm *OrmHmGroups)) *OrmHmGroups {
	ormAnd := NewOrmHmGroups()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmGroups) Or(fuc func(orm *OrmHmGroups)) *OrmHmGroups {
	ormOr := NewOrmHmGroups()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmGroups) Preload(query string, args ...interface{}) *OrmHmGroups {
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
func (orm *OrmHmGroups) Joins(query string, args ...interface{}) *OrmHmGroups {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmGroups) WhereGroupid(val int64) *OrmHmGroups {
	orm.db.Where("`groupid` = ?", val)
	return orm
}
func (orm *OrmHmGroups) InsertGetGroupid(row *HmGroups) int64 {
	orm.db.Create(row)
	return row.Groupid
}
func (orm *OrmHmGroups) WhereGroupidIn(val []int64) *OrmHmGroups {
	orm.db.Where("`groupid` IN ?", val)
	return orm
}
func (orm *OrmHmGroups) WhereGroupidGt(val int64) *OrmHmGroups {
	orm.db.Where("`groupid` > ?", val)
	return orm
}
func (orm *OrmHmGroups) WhereGroupidGte(val int64) *OrmHmGroups {
	orm.db.Where("`groupid` >= ?", val)
	return orm
}
func (orm *OrmHmGroups) WhereGroupidLt(val int64) *OrmHmGroups {
	orm.db.Where("`groupid` < ?", val)
	return orm
}
func (orm *OrmHmGroups) WhereGroupidLte(val int64) *OrmHmGroups {
	orm.db.Where("`groupid` <= ?", val)
	return orm
}
func (orm *OrmHmGroups) WhereGroupname(val string) *OrmHmGroups {
	orm.db.Where("`groupname` = ?", val)
	return orm
}

type HmGroupsList []*HmGroups

func (l HmGroupsList) GetGroupidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Groupid)
	}
	return got
}
func (l HmGroupsList) GetGroupidMap() map[int64]*HmGroups {
	got := make(map[int64]*HmGroups)
	for _, val := range l {
		got[val.Groupid] = val
	}
	return got
}
