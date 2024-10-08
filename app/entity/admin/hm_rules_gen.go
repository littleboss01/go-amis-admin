package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmRules struct {
	Ruleid        int32  `gorm:"column:ruleid;autoIncrement;type:int(11);primaryKey;index:ruleid,class:BTREE,unique" json:"ruleid"` //
	Ruleaccountid int32  `gorm:"column:ruleaccountid;type:int(11);index:idx_rules,class:BTREE" json:"ruleaccountid"`                //
	Rulename      string `gorm:"column:rulename;type:varchar(100)" json:"rulename"`                                                 //
	Ruleactive    int32  `gorm:"column:ruleactive;type:tinyint(4)" json:"ruleactive"`                                               //
	Ruleuseand    int32  `gorm:"column:ruleuseand;type:tinyint(4)" json:"ruleuseand"`                                               //
	Rulesortorder int32  `gorm:"column:rulesortorder;type:int(11)" json:"rulesortorder"`                                            //
}

var (
	HmRulesFieldRuleid        = "ruleid"
	HmRulesFieldRuleaccountid = "ruleaccountid"
	HmRulesFieldRulename      = "rulename"
	HmRulesFieldRuleactive    = "ruleactive"
	HmRulesFieldRuleuseand    = "ruleuseand"
	HmRulesFieldRulesortorder = "rulesortorder"
)

func (receiver *HmRules) TableName() string {
	return "hm_rules"
}

type OrmHmRules struct {
	db *gorm.DB
}

func (orm *OrmHmRules) TableName() string {
	return "hm_rules"
}

func NewOrmHmRules(txs ...*gorm.DB) *OrmHmRules {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmRules{})
	} else {
		tx = txs[0].Model(&HmRules{})
	}
	return &OrmHmRules{db: tx}
}

func (orm *OrmHmRules) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmRules) GetTableInfo() interface{} {
	return &HmRules{}
}

// Create insert the value into database
func (orm *OrmHmRules) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmRules) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmRules) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmRules) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmRules) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmRules) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmRules) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmRules) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmRules) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmRules) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmRules) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmRules) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmRules) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmRules) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmRules) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmRules) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmRules) Unscoped() *OrmHmRules {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmRules) Insert(row *HmRules) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmRules) Inserts(rows []*HmRules) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmRules) Order(value interface{}) *OrmHmRules {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmRules) Group(name string) *OrmHmRules {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmRules) Limit(limit int) *OrmHmRules {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmRules) Offset(offset int) *OrmHmRules {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmRules) Get() HmRulesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmRules) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmRules) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmRules{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmRules) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_rules")
}

func (orm *OrmHmRules) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmRules) First(conds ...interface{}) (*HmRules, bool) {
	dest := &HmRules{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmRules) Take(conds ...interface{}) (*HmRules, int64) {
	dest := &HmRules{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmRules) Last(conds ...interface{}) (*HmRules, int64) {
	dest := &HmRules{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmRules) Find(conds ...interface{}) (HmRulesList, int64) {
	list := make([]*HmRules, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmRules) Paginate(page int, limit int) (HmRulesList, int64) {
	var total int64
	list := make([]*HmRules, 0)
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
func (orm *OrmHmRules) SimplePaginate(page int, limit int) HmRulesList {
	list := make([]*HmRules, 0)
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
func (orm *OrmHmRules) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmRules) FirstOrInit(dest *HmRules, conds ...interface{}) (*HmRules, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmRules) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRules) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRules) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmRules) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmRules) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmRules) Where(query interface{}, args ...interface{}) *OrmHmRules {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmRules) Select(query interface{}, args ...interface{}) *OrmHmRules {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmRules) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmRules) And(fuc func(orm *OrmHmRules)) *OrmHmRules {
	ormAnd := NewOrmHmRules()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmRules) Or(fuc func(orm *OrmHmRules)) *OrmHmRules {
	ormOr := NewOrmHmRules()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmRules) Preload(query string, args ...interface{}) *OrmHmRules {
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
func (orm *OrmHmRules) Joins(query string, args ...interface{}) *OrmHmRules {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmRules) WhereRuleid(val int32) *OrmHmRules {
	orm.db.Where("`ruleid` = ?", val)
	return orm
}
func (orm *OrmHmRules) InsertGetRuleid(row *HmRules) int32 {
	orm.db.Create(row)
	return row.Ruleid
}
func (orm *OrmHmRules) WhereRuleidIn(val []int32) *OrmHmRules {
	orm.db.Where("`ruleid` IN ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleidGt(val int32) *OrmHmRules {
	orm.db.Where("`ruleid` > ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleidGte(val int32) *OrmHmRules {
	orm.db.Where("`ruleid` >= ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleidLt(val int32) *OrmHmRules {
	orm.db.Where("`ruleid` < ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleidLte(val int32) *OrmHmRules {
	orm.db.Where("`ruleid` <= ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleaccountid(val int32) *OrmHmRules {
	orm.db.Where("`ruleaccountid` = ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleaccountidIn(val []int32) *OrmHmRules {
	orm.db.Where("`ruleaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleaccountidGt(val int32) *OrmHmRules {
	orm.db.Where("`ruleaccountid` > ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleaccountidGte(val int32) *OrmHmRules {
	orm.db.Where("`ruleaccountid` >= ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleaccountidLt(val int32) *OrmHmRules {
	orm.db.Where("`ruleaccountid` < ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleaccountidLte(val int32) *OrmHmRules {
	orm.db.Where("`ruleaccountid` <= ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRulename(val string) *OrmHmRules {
	orm.db.Where("`rulename` = ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleactive(val int32) *OrmHmRules {
	orm.db.Where("`ruleactive` = ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRuleuseand(val int32) *OrmHmRules {
	orm.db.Where("`ruleuseand` = ?", val)
	return orm
}
func (orm *OrmHmRules) WhereRulesortorder(val int32) *OrmHmRules {
	orm.db.Where("`rulesortorder` = ?", val)
	return orm
}

type HmRulesList []*HmRules

func (l HmRulesList) GetRuleidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Ruleid)
	}
	return got
}
func (l HmRulesList) GetRuleidMap() map[int32]*HmRules {
	got := make(map[int32]*HmRules)
	for _, val := range l {
		got[val.Ruleid] = val
	}
	return got
}
func (l HmRulesList) GetRuleaccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Ruleaccountid)
	}
	return got
}
func (l HmRulesList) GetRuleaccountidMap() map[int32]*HmRules {
	got := make(map[int32]*HmRules)
	for _, val := range l {
		got[val.Ruleaccountid] = val
	}
	return got
}
