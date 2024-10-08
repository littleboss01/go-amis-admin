package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmGroupMembers struct {
	Memberid        int64 `gorm:"column:memberid;autoIncrement;type:bigint(20);primaryKey;index:memberid,class:BTREE,unique" json:"memberid"` //
	Membergroupid   int64 `gorm:"column:membergroupid;type:bigint(20)" json:"membergroupid"`                                                  //
	Memberaccountid int64 `gorm:"column:memberaccountid;type:bigint(20)" json:"memberaccountid"`                                              //
}

var (
	HmGroupMembersFieldMemberid        = "memberid"
	HmGroupMembersFieldMembergroupid   = "membergroupid"
	HmGroupMembersFieldMemberaccountid = "memberaccountid"
)

func (receiver *HmGroupMembers) TableName() string {
	return "hm_group_members"
}

type OrmHmGroupMembers struct {
	db *gorm.DB
}

func (orm *OrmHmGroupMembers) TableName() string {
	return "hm_group_members"
}

func NewOrmHmGroupMembers(txs ...*gorm.DB) *OrmHmGroupMembers {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmGroupMembers{})
	} else {
		tx = txs[0].Model(&HmGroupMembers{})
	}
	return &OrmHmGroupMembers{db: tx}
}

func (orm *OrmHmGroupMembers) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmGroupMembers) GetTableInfo() interface{} {
	return &HmGroupMembers{}
}

// Create insert the value into database
func (orm *OrmHmGroupMembers) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmGroupMembers) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmGroupMembers) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmGroupMembers) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmGroupMembers) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmGroupMembers) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmGroupMembers) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmGroupMembers) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmGroupMembers) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmGroupMembers) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmGroupMembers) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmGroupMembers) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmGroupMembers) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmGroupMembers) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmGroupMembers) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmGroupMembers) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmGroupMembers) Unscoped() *OrmHmGroupMembers {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmGroupMembers) Insert(row *HmGroupMembers) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmGroupMembers) Inserts(rows []*HmGroupMembers) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmGroupMembers) Order(value interface{}) *OrmHmGroupMembers {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmGroupMembers) Group(name string) *OrmHmGroupMembers {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmGroupMembers) Limit(limit int) *OrmHmGroupMembers {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmGroupMembers) Offset(offset int) *OrmHmGroupMembers {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmGroupMembers) Get() HmGroupMembersList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmGroupMembers) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmGroupMembers) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmGroupMembers{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmGroupMembers) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_group_members")
}

func (orm *OrmHmGroupMembers) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmGroupMembers) First(conds ...interface{}) (*HmGroupMembers, bool) {
	dest := &HmGroupMembers{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmGroupMembers) Take(conds ...interface{}) (*HmGroupMembers, int64) {
	dest := &HmGroupMembers{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmGroupMembers) Last(conds ...interface{}) (*HmGroupMembers, int64) {
	dest := &HmGroupMembers{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmGroupMembers) Find(conds ...interface{}) (HmGroupMembersList, int64) {
	list := make([]*HmGroupMembers, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmGroupMembers) Paginate(page int, limit int) (HmGroupMembersList, int64) {
	var total int64
	list := make([]*HmGroupMembers, 0)
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
func (orm *OrmHmGroupMembers) SimplePaginate(page int, limit int) HmGroupMembersList {
	list := make([]*HmGroupMembers, 0)
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
func (orm *OrmHmGroupMembers) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmGroupMembers) FirstOrInit(dest *HmGroupMembers, conds ...interface{}) (*HmGroupMembers, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmGroupMembers) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGroupMembers) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmGroupMembers) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmGroupMembers) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmGroupMembers) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmGroupMembers) Where(query interface{}, args ...interface{}) *OrmHmGroupMembers {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmGroupMembers) Select(query interface{}, args ...interface{}) *OrmHmGroupMembers {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmGroupMembers) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmGroupMembers) And(fuc func(orm *OrmHmGroupMembers)) *OrmHmGroupMembers {
	ormAnd := NewOrmHmGroupMembers()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmGroupMembers) Or(fuc func(orm *OrmHmGroupMembers)) *OrmHmGroupMembers {
	ormOr := NewOrmHmGroupMembers()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmGroupMembers) Preload(query string, args ...interface{}) *OrmHmGroupMembers {
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
func (orm *OrmHmGroupMembers) Joins(query string, args ...interface{}) *OrmHmGroupMembers {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmGroupMembers) WhereMemberid(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberid` = ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) InsertGetMemberid(row *HmGroupMembers) int64 {
	orm.db.Create(row)
	return row.Memberid
}
func (orm *OrmHmGroupMembers) WhereMemberidIn(val []int64) *OrmHmGroupMembers {
	orm.db.Where("`memberid` IN ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberidGt(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberid` > ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberidGte(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberid` >= ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberidLt(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberid` < ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberidLte(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberid` <= ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMembergroupid(val int64) *OrmHmGroupMembers {
	orm.db.Where("`membergroupid` = ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMembergroupidIn(val []int64) *OrmHmGroupMembers {
	orm.db.Where("`membergroupid` IN ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMembergroupidNe(val int64) *OrmHmGroupMembers {
	orm.db.Where("`membergroupid` <> ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberaccountid(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberaccountid` = ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberaccountidIn(val []int64) *OrmHmGroupMembers {
	orm.db.Where("`memberaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmGroupMembers) WhereMemberaccountidNe(val int64) *OrmHmGroupMembers {
	orm.db.Where("`memberaccountid` <> ?", val)
	return orm
}

type HmGroupMembersList []*HmGroupMembers

func (l HmGroupMembersList) GetMemberidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Memberid)
	}
	return got
}
func (l HmGroupMembersList) GetMemberidMap() map[int64]*HmGroupMembers {
	got := make(map[int64]*HmGroupMembers)
	for _, val := range l {
		got[val.Memberid] = val
	}
	return got
}
func (l HmGroupMembersList) GetMembergroupidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Membergroupid)
	}
	return got
}
func (l HmGroupMembersList) GetMembergroupidMap() map[int64]*HmGroupMembers {
	got := make(map[int64]*HmGroupMembers)
	for _, val := range l {
		got[val.Membergroupid] = val
	}
	return got
}
func (l HmGroupMembersList) GetMemberaccountidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Memberaccountid)
	}
	return got
}
func (l HmGroupMembersList) GetMemberaccountidMap() map[int64]*HmGroupMembers {
	got := make(map[int64]*HmGroupMembers)
	for _, val := range l {
		got[val.Memberaccountid] = val
	}
	return got
}
