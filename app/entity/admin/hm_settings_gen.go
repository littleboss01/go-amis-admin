package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmSettings struct {
	Settingid      int64  `gorm:"column:settingid;autoIncrement;type:bigint(11);primaryKey;index:settingid,class:BTREE,unique" json:"settingid"` //
	Settingname    string `gorm:"column:settingname;type:varchar(30);index:settingname,class:BTREE,unique" json:"settingname"`                   //
	Settingstring  string `gorm:"column:settingstring;type:varchar(4000)" json:"settingstring"`                                                  //
	Settinginteger int32  `gorm:"column:settinginteger;type:int(11)" json:"settinginteger"`                                                      //
}

var (
	HmSettingsFieldSettingid      = "settingid"
	HmSettingsFieldSettingname    = "settingname"
	HmSettingsFieldSettingstring  = "settingstring"
	HmSettingsFieldSettinginteger = "settinginteger"
)

func (receiver *HmSettings) TableName() string {
	return "hm_settings"
}

type OrmHmSettings struct {
	db *gorm.DB
}

func (orm *OrmHmSettings) TableName() string {
	return "hm_settings"
}

func NewOrmHmSettings(txs ...*gorm.DB) *OrmHmSettings {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmSettings{})
	} else {
		tx = txs[0].Model(&HmSettings{})
	}
	return &OrmHmSettings{db: tx}
}

func (orm *OrmHmSettings) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmSettings) GetTableInfo() interface{} {
	return &HmSettings{}
}

// Create insert the value into database
func (orm *OrmHmSettings) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmSettings) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmSettings) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmSettings) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmSettings) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmSettings) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmSettings) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmSettings) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmSettings) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmSettings) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmSettings) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmSettings) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmSettings) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmSettings) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmSettings) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmSettings) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmSettings) Unscoped() *OrmHmSettings {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmSettings) Insert(row *HmSettings) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmSettings) Inserts(rows []*HmSettings) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmSettings) Order(value interface{}) *OrmHmSettings {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmSettings) Group(name string) *OrmHmSettings {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmSettings) Limit(limit int) *OrmHmSettings {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmSettings) Offset(offset int) *OrmHmSettings {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmSettings) Get() HmSettingsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmSettings) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmSettings) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmSettings{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmSettings) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_settings")
}

func (orm *OrmHmSettings) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmSettings) First(conds ...interface{}) (*HmSettings, bool) {
	dest := &HmSettings{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmSettings) Take(conds ...interface{}) (*HmSettings, int64) {
	dest := &HmSettings{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmSettings) Last(conds ...interface{}) (*HmSettings, int64) {
	dest := &HmSettings{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmSettings) Find(conds ...interface{}) (HmSettingsList, int64) {
	list := make([]*HmSettings, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmSettings) Paginate(page int, limit int) (HmSettingsList, int64) {
	var total int64
	list := make([]*HmSettings, 0)
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
func (orm *OrmHmSettings) SimplePaginate(page int, limit int) HmSettingsList {
	list := make([]*HmSettings, 0)
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
func (orm *OrmHmSettings) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmSettings) FirstOrInit(dest *HmSettings, conds ...interface{}) (*HmSettings, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmSettings) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSettings) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmSettings) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmSettings) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmSettings) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmSettings) Where(query interface{}, args ...interface{}) *OrmHmSettings {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmSettings) Select(query interface{}, args ...interface{}) *OrmHmSettings {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmSettings) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmSettings) And(fuc func(orm *OrmHmSettings)) *OrmHmSettings {
	ormAnd := NewOrmHmSettings()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmSettings) Or(fuc func(orm *OrmHmSettings)) *OrmHmSettings {
	ormOr := NewOrmHmSettings()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmSettings) Preload(query string, args ...interface{}) *OrmHmSettings {
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
func (orm *OrmHmSettings) Joins(query string, args ...interface{}) *OrmHmSettings {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmSettings) WhereSettingid(val int64) *OrmHmSettings {
	orm.db.Where("`settingid` = ?", val)
	return orm
}
func (orm *OrmHmSettings) InsertGetSettingid(row *HmSettings) int64 {
	orm.db.Create(row)
	return row.Settingid
}
func (orm *OrmHmSettings) WhereSettingidIn(val []int64) *OrmHmSettings {
	orm.db.Where("`settingid` IN ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingidGt(val int64) *OrmHmSettings {
	orm.db.Where("`settingid` > ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingidGte(val int64) *OrmHmSettings {
	orm.db.Where("`settingid` >= ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingidLt(val int64) *OrmHmSettings {
	orm.db.Where("`settingid` < ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingidLte(val int64) *OrmHmSettings {
	orm.db.Where("`settingid` <= ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingname(val string) *OrmHmSettings {
	orm.db.Where("`settingname` = ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingnameIn(val []string) *OrmHmSettings {
	orm.db.Where("`settingname` IN ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingnameGt(val string) *OrmHmSettings {
	orm.db.Where("`settingname` > ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingnameGte(val string) *OrmHmSettings {
	orm.db.Where("`settingname` >= ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingnameLt(val string) *OrmHmSettings {
	orm.db.Where("`settingname` < ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingnameLte(val string) *OrmHmSettings {
	orm.db.Where("`settingname` <= ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettingstring(val string) *OrmHmSettings {
	orm.db.Where("`settingstring` = ?", val)
	return orm
}
func (orm *OrmHmSettings) WhereSettinginteger(val int32) *OrmHmSettings {
	orm.db.Where("`settinginteger` = ?", val)
	return orm
}

type HmSettingsList []*HmSettings

func (l HmSettingsList) GetSettingidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Settingid)
	}
	return got
}
func (l HmSettingsList) GetSettingidMap() map[int64]*HmSettings {
	got := make(map[int64]*HmSettings)
	for _, val := range l {
		got[val.Settingid] = val
	}
	return got
}
