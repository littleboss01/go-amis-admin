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

type SysDict struct {
	Id        int64         `gorm:"column:id;autoIncrement;type:bigint(20);primaryKey" json:"id"`                             //
	Group     string        `gorm:"column:group;type:varchar(32);index:uk_g_v,class:BTREE,unique;comment:'分组'" json:"group"`  // 分组
	GroupName string        `gorm:"column:group_name;type:varchar(32);comment:'分组名称'" json:"group_name"`                      // 分组名称
	Value     string        `gorm:"column:value;type:varchar(64);index:uk_g_v,class:BTREE,unique;comment:'数据值'" json:"value"` // 数据值
	Label     string        `gorm:"column:label;type:varchar(128);comment:'显示标签'" json:"label"`                               // 显示标签
	Status    string        `gorm:"column:status;type:varchar(16);comment:'数据状态'" json:"status"`                              // 数据状态
	Remark    string        `gorm:"column:remark;type:varchar(512);default:;comment:'备注'" json:"remark"`                      // 备注
	Utime     database.Time `gorm:"column:utime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'更新时间'" json:"utime"`        // 更新时间
	Ctime     database.Time `gorm:"column:ctime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"ctime"`        // 创建时间
}

var (
	SysDictFieldId        = "id"
	SysDictFieldGroup     = "group"
	SysDictFieldGroupName = "group_name"
	SysDictFieldValue     = "value"
	SysDictFieldLabel     = "label"
	SysDictFieldStatus    = "status"
	SysDictFieldRemark    = "remark"
	SysDictFieldUtime     = "utime"
	SysDictFieldCtime     = "ctime"
)

func (receiver *SysDict) TableName() string {
	return "sys_dict"
}

type OrmSysDict struct {
	db *gorm.DB
}

func (orm *OrmSysDict) TableName() string {
	return "sys_dict"
}

func NewOrmSysDict(txs ...*gorm.DB) *OrmSysDict {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&SysDict{})
	} else {
		tx = txs[0].Model(&SysDict{})
	}
	return &OrmSysDict{db: tx}
}

func (orm *OrmSysDict) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmSysDict) GetTableInfo() interface{} {
	return &SysDict{}
}

// Create insert the value into database
func (orm *OrmSysDict) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmSysDict) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmSysDict) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmSysDict) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmSysDict) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmSysDict) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmSysDict) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmSysDict) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmSysDict) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmSysDict) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmSysDict) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmSysDict) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmSysDict) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmSysDict) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmSysDict) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmSysDict) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmSysDict) Unscoped() *OrmSysDict {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmSysDict) Insert(row *SysDict) error {
	return orm.db.Create(row).Error
}

func (orm *OrmSysDict) Inserts(rows []*SysDict) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmSysDict) Order(value interface{}) *OrmSysDict {
	orm.db.Order(value)
	return orm
}

func (orm *OrmSysDict) Group(name string) *OrmSysDict {
	orm.db.Group(name)
	return orm
}

func (orm *OrmSysDict) Limit(limit int) *OrmSysDict {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmSysDict) Offset(offset int) *OrmSysDict {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmSysDict) Get() SysDictList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmSysDict) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmSysDict) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&SysDict{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmSysDict) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM sys_dict")
}

func (orm *OrmSysDict) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmSysDict) First(conds ...interface{}) (*SysDict, bool) {
	dest := &SysDict{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmSysDict) Take(conds ...interface{}) (*SysDict, int64) {
	dest := &SysDict{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmSysDict) Last(conds ...interface{}) (*SysDict, int64) {
	dest := &SysDict{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmSysDict) Find(conds ...interface{}) (SysDictList, int64) {
	list := make([]*SysDict, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmSysDict) Paginate(page int, limit int) (SysDictList, int64) {
	var total int64
	list := make([]*SysDict, 0)
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
func (orm *OrmSysDict) SimplePaginate(page int, limit int) SysDictList {
	list := make([]*SysDict, 0)
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
func (orm *OrmSysDict) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmSysDict) FirstOrInit(dest *SysDict, conds ...interface{}) (*SysDict, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmSysDict) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysDict) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysDict) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmSysDict) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmSysDict) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmSysDict) Where(query interface{}, args ...interface{}) *OrmSysDict {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmSysDict) Select(query interface{}, args ...interface{}) *OrmSysDict {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmSysDict) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmSysDict) And(fuc func(orm *OrmSysDict)) *OrmSysDict {
	ormAnd := NewOrmSysDict()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmSysDict) Or(fuc func(orm *OrmSysDict)) *OrmSysDict {
	ormOr := NewOrmSysDict()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmSysDict) Preload(query string, args ...interface{}) *OrmSysDict {
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
func (orm *OrmSysDict) Joins(query string, args ...interface{}) *OrmSysDict {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmSysDict) WhereId(val int64) *OrmSysDict {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmSysDict) InsertGetId(row *SysDict) int64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmSysDict) WhereIdIn(val []int64) *OrmSysDict {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmSysDict) WhereIdGt(val int64) *OrmSysDict {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmSysDict) WhereIdGte(val int64) *OrmSysDict {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereIdLt(val int64) *OrmSysDict {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmSysDict) WhereIdLte(val int64) *OrmSysDict {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroup(val string) *OrmSysDict {
	orm.db.Where("`group` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroupIn(val []string) *OrmSysDict {
	orm.db.Where("`group` IN ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroupGt(val string) *OrmSysDict {
	orm.db.Where("`group` > ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroupGte(val string) *OrmSysDict {
	orm.db.Where("`group` >= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroupLt(val string) *OrmSysDict {
	orm.db.Where("`group` < ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroupLte(val string) *OrmSysDict {
	orm.db.Where("`group` <= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereGroupName(val string) *OrmSysDict {
	orm.db.Where("`group_name` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereValue(val string) *OrmSysDict {
	orm.db.Where("`value` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereLabel(val string) *OrmSysDict {
	orm.db.Where("`label` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereStatus(val string) *OrmSysDict {
	orm.db.Where("`status` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereStatusIn(val []string) *OrmSysDict {
	orm.db.Where("`status` IN ?", val)
	return orm
}
func (orm *OrmSysDict) WhereStatusNe(val string) *OrmSysDict {
	orm.db.Where("`status` <> ?", val)
	return orm
}
func (orm *OrmSysDict) WhereRemark(val string) *OrmSysDict {
	orm.db.Where("`remark` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereUtime(val database.Time) *OrmSysDict {
	orm.db.Where("`utime` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereUtimeBetween(begin database.Time, end database.Time) *OrmSysDict {
	orm.db.Where("`utime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysDict) WhereUtimeLte(val database.Time) *OrmSysDict {
	orm.db.Where("`utime` <= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereUtimeGte(val database.Time) *OrmSysDict {
	orm.db.Where("`utime` >= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereCtime(val database.Time) *OrmSysDict {
	orm.db.Where("`ctime` = ?", val)
	return orm
}
func (orm *OrmSysDict) WhereCtimeBetween(begin database.Time, end database.Time) *OrmSysDict {
	orm.db.Where("`ctime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysDict) WhereCtimeLte(val database.Time) *OrmSysDict {
	orm.db.Where("`ctime` <= ?", val)
	return orm
}
func (orm *OrmSysDict) WhereCtimeGte(val database.Time) *OrmSysDict {
	orm.db.Where("`ctime` >= ?", val)
	return orm
}

type SysDictList []*SysDict

func (l SysDictList) GetIdList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l SysDictList) GetIdMap() map[int64]*SysDict {
	got := make(map[int64]*SysDict)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
