package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmRuleCriterias struct {
	Criteriaid              int32  `gorm:"column:criteriaid;autoIncrement;type:int(11);primaryKey;index:criteriaid,class:BTREE,unique" json:"criteriaid"` //
	Criteriaruleid          int32  `gorm:"column:criteriaruleid;type:int(11);index:idx_rules_criterias,class:BTREE" json:"criteriaruleid"`                //
	Criteriausepredefined   int32  `gorm:"column:criteriausepredefined;type:tinyint(4)" json:"criteriausepredefined"`                                     //
	Criteriapredefinedfield int32  `gorm:"column:criteriapredefinedfield;type:tinyint(4)" json:"criteriapredefinedfield"`                                 //
	Criteriaheadername      string `gorm:"column:criteriaheadername;type:varchar(255)" json:"criteriaheadername"`                                         //
	Criteriamatchtype       int32  `gorm:"column:criteriamatchtype;type:tinyint(4)" json:"criteriamatchtype"`                                             //
	Criteriamatchvalue      string `gorm:"column:criteriamatchvalue;type:varchar(255)" json:"criteriamatchvalue"`                                         //
}

var (
	HmRuleCriteriasFieldCriteriaid              = "criteriaid"
	HmRuleCriteriasFieldCriteriaruleid          = "criteriaruleid"
	HmRuleCriteriasFieldCriteriausepredefined   = "criteriausepredefined"
	HmRuleCriteriasFieldCriteriapredefinedfield = "criteriapredefinedfield"
	HmRuleCriteriasFieldCriteriaheadername      = "criteriaheadername"
	HmRuleCriteriasFieldCriteriamatchtype       = "criteriamatchtype"
	HmRuleCriteriasFieldCriteriamatchvalue      = "criteriamatchvalue"
)

func (receiver *HmRuleCriterias) TableName() string {
	return "hm_rule_criterias"
}

type OrmHmRuleCriterias struct {
	db *gorm.DB
}

func (orm *OrmHmRuleCriterias) TableName() string {
	return "hm_rule_criterias"
}

func NewOrmHmRuleCriterias(txs ...*gorm.DB) *OrmHmRuleCriterias {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmRuleCriterias{})
	} else {
		tx = txs[0].Model(&HmRuleCriterias{})
	}
	return &OrmHmRuleCriterias{db: tx}
}

func (orm *OrmHmRuleCriterias) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmRuleCriterias) GetTableInfo() interface{} {
	return &HmRuleCriterias{}
}

// Create insert the value into database
func (orm *OrmHmRuleCriterias) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmRuleCriterias) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmRuleCriterias) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmRuleCriterias) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmRuleCriterias) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmRuleCriterias) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmRuleCriterias) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmRuleCriterias) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmRuleCriterias) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmRuleCriterias) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmRuleCriterias) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmRuleCriterias) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmRuleCriterias) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmRuleCriterias) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmRuleCriterias) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmRuleCriterias) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmRuleCriterias) Unscoped() *OrmHmRuleCriterias {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmRuleCriterias) Insert(row *HmRuleCriterias) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmRuleCriterias) Inserts(rows []*HmRuleCriterias) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmRuleCriterias) Order(value interface{}) *OrmHmRuleCriterias {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmRuleCriterias) Group(name string) *OrmHmRuleCriterias {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmRuleCriterias) Limit(limit int) *OrmHmRuleCriterias {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmRuleCriterias) Offset(offset int) *OrmHmRuleCriterias {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmRuleCriterias) Get() HmRuleCriteriasList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmRuleCriterias) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmRuleCriterias) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmRuleCriterias{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmRuleCriterias) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_rule_criterias")
}

func (orm *OrmHmRuleCriterias) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmRuleCriterias) First(conds ...interface{}) (*HmRuleCriterias, bool) {
	dest := &HmRuleCriterias{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmRuleCriterias) Take(conds ...interface{}) (*HmRuleCriterias, int64) {
	dest := &HmRuleCriterias{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmRuleCriterias) Last(conds ...interface{}) (*HmRuleCriterias, int64) {
	dest := &HmRuleCriterias{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmRuleCriterias) Find(conds ...interface{}) (HmRuleCriteriasList, int64) {
	list := make([]*HmRuleCriterias, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmRuleCriterias) Paginate(page int, limit int) (HmRuleCriteriasList, int64) {
	var total int64
	list := make([]*HmRuleCriterias, 0)
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
func (orm *OrmHmRuleCriterias) SimplePaginate(page int, limit int) HmRuleCriteriasList {
	list := make([]*HmRuleCriterias, 0)
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
func (orm *OrmHmRuleCriterias) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmRuleCriterias) FirstOrInit(dest *HmRuleCriterias, conds ...interface{}) (*HmRuleCriterias, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmRuleCriterias) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRuleCriterias) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRuleCriterias) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmRuleCriterias) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmRuleCriterias) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmRuleCriterias) Where(query interface{}, args ...interface{}) *OrmHmRuleCriterias {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmRuleCriterias) Select(query interface{}, args ...interface{}) *OrmHmRuleCriterias {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmRuleCriterias) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmRuleCriterias) And(fuc func(orm *OrmHmRuleCriterias)) *OrmHmRuleCriterias {
	ormAnd := NewOrmHmRuleCriterias()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmRuleCriterias) Or(fuc func(orm *OrmHmRuleCriterias)) *OrmHmRuleCriterias {
	ormOr := NewOrmHmRuleCriterias()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmRuleCriterias) Preload(query string, args ...interface{}) *OrmHmRuleCriterias {
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
func (orm *OrmHmRuleCriterias) Joins(query string, args ...interface{}) *OrmHmRuleCriterias {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmRuleCriterias) WhereCriteriaid(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaid` = ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) InsertGetCriteriaid(row *HmRuleCriterias) int32 {
	orm.db.Create(row)
	return row.Criteriaid
}
func (orm *OrmHmRuleCriterias) WhereCriteriaidIn(val []int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaid` IN ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaidGt(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaid` > ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaidGte(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaid` >= ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaidLt(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaid` < ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaidLte(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaid` <= ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaruleid(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaruleid` = ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaruleidIn(val []int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaruleid` IN ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaruleidGt(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaruleid` > ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaruleidGte(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaruleid` >= ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaruleidLt(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaruleid` < ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaruleidLte(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaruleid` <= ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriausepredefined(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriausepredefined` = ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriapredefinedfield(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriapredefinedfield` = ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriaheadername(val string) *OrmHmRuleCriterias {
	orm.db.Where("`criteriaheadername` = ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriamatchtype(val int32) *OrmHmRuleCriterias {
	orm.db.Where("`criteriamatchtype` = ?", val)
	return orm
}
func (orm *OrmHmRuleCriterias) WhereCriteriamatchvalue(val string) *OrmHmRuleCriterias {
	orm.db.Where("`criteriamatchvalue` = ?", val)
	return orm
}

type HmRuleCriteriasList []*HmRuleCriterias

func (l HmRuleCriteriasList) GetCriteriaidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Criteriaid)
	}
	return got
}
func (l HmRuleCriteriasList) GetCriteriaidMap() map[int32]*HmRuleCriterias {
	got := make(map[int32]*HmRuleCriterias)
	for _, val := range l {
		got[val.Criteriaid] = val
	}
	return got
}
func (l HmRuleCriteriasList) GetCriteriaruleidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Criteriaruleid)
	}
	return got
}
func (l HmRuleCriteriasList) GetCriteriaruleidMap() map[int32]*HmRuleCriterias {
	got := make(map[int32]*HmRuleCriterias)
	for _, val := range l {
		got[val.Criteriaruleid] = val
	}
	return got
}
