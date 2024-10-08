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

type SysConfig struct {
	Id     int64         `gorm:"column:id;autoIncrement;type:bigint(20);primaryKey" json:"id"`                                     //
	Scope  string        `gorm:"column:scope;type:varchar(32);index:uk_s_k,class:BTREE,unique;default:;comment:'范围'" json:"scope"` // 范围
	Key    string        `gorm:"column:key;type:varchar(64);index:uk_s_k,class:BTREE,unique;comment:'配置键'" json:"key"`             // 配置键
	Value  string        `gorm:"column:value;type:text;comment:'配置值'" json:"value"`                                                // 配置值
	Status string        `gorm:"column:status;type:varchar(16);comment:'配置状态'" json:"status"`                                      // 配置状态
	Remark string        `gorm:"column:remark;type:varchar(512);comment:'备注'" json:"remark"`                                       // 备注
	Utime  database.Time `gorm:"column:utime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'更新时间'" json:"utime"`                // 更新时间
	Ctime  database.Time `gorm:"column:ctime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"ctime"`                // 创建时间
}

var (
	SysConfigFieldId     = "id"
	SysConfigFieldScope  = "scope"
	SysConfigFieldKey    = "key"
	SysConfigFieldValue  = "value"
	SysConfigFieldStatus = "status"
	SysConfigFieldRemark = "remark"
	SysConfigFieldUtime  = "utime"
	SysConfigFieldCtime  = "ctime"
)

func (receiver *SysConfig) TableName() string {
	return "sys_config"
}

type OrmSysConfig struct {
	db *gorm.DB
}

func (orm *OrmSysConfig) TableName() string {
	return "sys_config"
}

func NewOrmSysConfig(txs ...*gorm.DB) *OrmSysConfig {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&SysConfig{})
	} else {
		tx = txs[0].Model(&SysConfig{})
	}
	return &OrmSysConfig{db: tx}
}

func (orm *OrmSysConfig) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmSysConfig) GetTableInfo() interface{} {
	return &SysConfig{}
}

// Create insert the value into database
func (orm *OrmSysConfig) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmSysConfig) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmSysConfig) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmSysConfig) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmSysConfig) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmSysConfig) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmSysConfig) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmSysConfig) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmSysConfig) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmSysConfig) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmSysConfig) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmSysConfig) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmSysConfig) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmSysConfig) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmSysConfig) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmSysConfig) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmSysConfig) Unscoped() *OrmSysConfig {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmSysConfig) Insert(row *SysConfig) error {
	return orm.db.Create(row).Error
}

func (orm *OrmSysConfig) Inserts(rows []*SysConfig) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmSysConfig) Order(value interface{}) *OrmSysConfig {
	orm.db.Order(value)
	return orm
}

func (orm *OrmSysConfig) Group(name string) *OrmSysConfig {
	orm.db.Group(name)
	return orm
}

func (orm *OrmSysConfig) Limit(limit int) *OrmSysConfig {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmSysConfig) Offset(offset int) *OrmSysConfig {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmSysConfig) Get() SysConfigList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmSysConfig) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmSysConfig) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&SysConfig{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmSysConfig) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM sys_config")
}

func (orm *OrmSysConfig) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmSysConfig) First(conds ...interface{}) (*SysConfig, bool) {
	dest := &SysConfig{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmSysConfig) Take(conds ...interface{}) (*SysConfig, int64) {
	dest := &SysConfig{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmSysConfig) Last(conds ...interface{}) (*SysConfig, int64) {
	dest := &SysConfig{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmSysConfig) Find(conds ...interface{}) (SysConfigList, int64) {
	list := make([]*SysConfig, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmSysConfig) Paginate(page int, limit int) (SysConfigList, int64) {
	var total int64
	list := make([]*SysConfig, 0)
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
func (orm *OrmSysConfig) SimplePaginate(page int, limit int) SysConfigList {
	list := make([]*SysConfig, 0)
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
func (orm *OrmSysConfig) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmSysConfig) FirstOrInit(dest *SysConfig, conds ...interface{}) (*SysConfig, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmSysConfig) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysConfig) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysConfig) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmSysConfig) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmSysConfig) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmSysConfig) Where(query interface{}, args ...interface{}) *OrmSysConfig {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmSysConfig) Select(query interface{}, args ...interface{}) *OrmSysConfig {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmSysConfig) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmSysConfig) And(fuc func(orm *OrmSysConfig)) *OrmSysConfig {
	ormAnd := NewOrmSysConfig()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmSysConfig) Or(fuc func(orm *OrmSysConfig)) *OrmSysConfig {
	ormOr := NewOrmSysConfig()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmSysConfig) Preload(query string, args ...interface{}) *OrmSysConfig {
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
func (orm *OrmSysConfig) Joins(query string, args ...interface{}) *OrmSysConfig {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmSysConfig) WhereId(val int64) *OrmSysConfig {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmSysConfig) InsertGetId(row *SysConfig) int64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmSysConfig) WhereIdIn(val []int64) *OrmSysConfig {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereIdGt(val int64) *OrmSysConfig {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereIdGte(val int64) *OrmSysConfig {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereIdLt(val int64) *OrmSysConfig {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereIdLte(val int64) *OrmSysConfig {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereScope(val string) *OrmSysConfig {
	orm.db.Where("`scope` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereKey(val string) *OrmSysConfig {
	orm.db.Where("`key` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereKeyIn(val []string) *OrmSysConfig {
	orm.db.Where("`key` IN ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereKeyGt(val string) *OrmSysConfig {
	orm.db.Where("`key` > ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereKeyGte(val string) *OrmSysConfig {
	orm.db.Where("`key` >= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereKeyLt(val string) *OrmSysConfig {
	orm.db.Where("`key` < ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereKeyLte(val string) *OrmSysConfig {
	orm.db.Where("`key` <= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereValue(val string) *OrmSysConfig {
	orm.db.Where("`value` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereStatus(val string) *OrmSysConfig {
	orm.db.Where("`status` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereStatusIn(val []string) *OrmSysConfig {
	orm.db.Where("`status` IN ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereStatusNe(val string) *OrmSysConfig {
	orm.db.Where("`status` <> ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereRemark(val string) *OrmSysConfig {
	orm.db.Where("`remark` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereUtime(val database.Time) *OrmSysConfig {
	orm.db.Where("`utime` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereUtimeBetween(begin database.Time, end database.Time) *OrmSysConfig {
	orm.db.Where("`utime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysConfig) WhereUtimeLte(val database.Time) *OrmSysConfig {
	orm.db.Where("`utime` <= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereUtimeGte(val database.Time) *OrmSysConfig {
	orm.db.Where("`utime` >= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereCtime(val database.Time) *OrmSysConfig {
	orm.db.Where("`ctime` = ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereCtimeBetween(begin database.Time, end database.Time) *OrmSysConfig {
	orm.db.Where("`ctime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysConfig) WhereCtimeLte(val database.Time) *OrmSysConfig {
	orm.db.Where("`ctime` <= ?", val)
	return orm
}
func (orm *OrmSysConfig) WhereCtimeGte(val database.Time) *OrmSysConfig {
	orm.db.Where("`ctime` >= ?", val)
	return orm
}

type SysConfigList []*SysConfig

func (l SysConfigList) GetIdList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l SysConfigList) GetIdMap() map[int64]*SysConfig {
	got := make(map[int64]*SysConfig)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
