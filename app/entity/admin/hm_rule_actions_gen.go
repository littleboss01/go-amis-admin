package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmRuleActions struct {
	Actionid             int32  `gorm:"column:actionid;autoIncrement;type:int(11);primaryKey;index:actionid,class:BTREE,unique" json:"actionid"` //
	Actionruleid         int32  `gorm:"column:actionruleid;type:int(11);index:idx_rules_actions,class:BTREE" json:"actionruleid"`                //
	Actiontype           int32  `gorm:"column:actiontype;type:tinyint(4)" json:"actiontype"`                                                     //
	Actionimapfolder     string `gorm:"column:actionimapfolder;type:varchar(255)" json:"actionimapfolder"`                                       //
	Actionsubject        string `gorm:"column:actionsubject;type:varchar(255)" json:"actionsubject"`                                             //
	Actionfromname       string `gorm:"column:actionfromname;type:varchar(255)" json:"actionfromname"`                                           //
	Actionfromaddress    string `gorm:"column:actionfromaddress;type:varchar(255)" json:"actionfromaddress"`                                     //
	Actionto             string `gorm:"column:actionto;type:varchar(255)" json:"actionto"`                                                       //
	Actionbody           string `gorm:"column:actionbody;type:text" json:"actionbody"`                                                           //
	Actionfilename       string `gorm:"column:actionfilename;type:varchar(255)" json:"actionfilename"`                                           //
	Actionsortorder      int32  `gorm:"column:actionsortorder;type:int(11)" json:"actionsortorder"`                                              //
	Actionscriptfunction string `gorm:"column:actionscriptfunction;type:varchar(255)" json:"actionscriptfunction"`                               //
	Actionheader         string `gorm:"column:actionheader;type:varchar(80)" json:"actionheader"`                                                //
	Actionvalue          string `gorm:"column:actionvalue;type:varchar(255)" json:"actionvalue"`                                                 //
	Actionrouteid        int32  `gorm:"column:actionrouteid;type:int(11)" json:"actionrouteid"`                                                  //
}

var (
	HmRuleActionsFieldActionid             = "actionid"
	HmRuleActionsFieldActionruleid         = "actionruleid"
	HmRuleActionsFieldActiontype           = "actiontype"
	HmRuleActionsFieldActionimapfolder     = "actionimapfolder"
	HmRuleActionsFieldActionsubject        = "actionsubject"
	HmRuleActionsFieldActionfromname       = "actionfromname"
	HmRuleActionsFieldActionfromaddress    = "actionfromaddress"
	HmRuleActionsFieldActionto             = "actionto"
	HmRuleActionsFieldActionbody           = "actionbody"
	HmRuleActionsFieldActionfilename       = "actionfilename"
	HmRuleActionsFieldActionsortorder      = "actionsortorder"
	HmRuleActionsFieldActionscriptfunction = "actionscriptfunction"
	HmRuleActionsFieldActionheader         = "actionheader"
	HmRuleActionsFieldActionvalue          = "actionvalue"
	HmRuleActionsFieldActionrouteid        = "actionrouteid"
)

func (receiver *HmRuleActions) TableName() string {
	return "hm_rule_actions"
}

type OrmHmRuleActions struct {
	db *gorm.DB
}

func (orm *OrmHmRuleActions) TableName() string {
	return "hm_rule_actions"
}

func NewOrmHmRuleActions(txs ...*gorm.DB) *OrmHmRuleActions {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmRuleActions{})
	} else {
		tx = txs[0].Model(&HmRuleActions{})
	}
	return &OrmHmRuleActions{db: tx}
}

func (orm *OrmHmRuleActions) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmRuleActions) GetTableInfo() interface{} {
	return &HmRuleActions{}
}

// Create insert the value into database
func (orm *OrmHmRuleActions) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmRuleActions) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmRuleActions) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmRuleActions) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmRuleActions) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmRuleActions) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmRuleActions) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmRuleActions) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmRuleActions) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmRuleActions) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmRuleActions) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmRuleActions) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmRuleActions) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmRuleActions) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmRuleActions) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmRuleActions) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmRuleActions) Unscoped() *OrmHmRuleActions {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmRuleActions) Insert(row *HmRuleActions) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmRuleActions) Inserts(rows []*HmRuleActions) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmRuleActions) Order(value interface{}) *OrmHmRuleActions {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmRuleActions) Group(name string) *OrmHmRuleActions {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmRuleActions) Limit(limit int) *OrmHmRuleActions {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmRuleActions) Offset(offset int) *OrmHmRuleActions {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmRuleActions) Get() HmRuleActionsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmRuleActions) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmRuleActions) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmRuleActions{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmRuleActions) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_rule_actions")
}

func (orm *OrmHmRuleActions) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmRuleActions) First(conds ...interface{}) (*HmRuleActions, bool) {
	dest := &HmRuleActions{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmRuleActions) Take(conds ...interface{}) (*HmRuleActions, int64) {
	dest := &HmRuleActions{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmRuleActions) Last(conds ...interface{}) (*HmRuleActions, int64) {
	dest := &HmRuleActions{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmRuleActions) Find(conds ...interface{}) (HmRuleActionsList, int64) {
	list := make([]*HmRuleActions, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmRuleActions) Paginate(page int, limit int) (HmRuleActionsList, int64) {
	var total int64
	list := make([]*HmRuleActions, 0)
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
func (orm *OrmHmRuleActions) SimplePaginate(page int, limit int) HmRuleActionsList {
	list := make([]*HmRuleActions, 0)
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
func (orm *OrmHmRuleActions) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmRuleActions) FirstOrInit(dest *HmRuleActions, conds ...interface{}) (*HmRuleActions, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmRuleActions) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRuleActions) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmRuleActions) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmRuleActions) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmRuleActions) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmRuleActions) Where(query interface{}, args ...interface{}) *OrmHmRuleActions {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmRuleActions) Select(query interface{}, args ...interface{}) *OrmHmRuleActions {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmRuleActions) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmRuleActions) And(fuc func(orm *OrmHmRuleActions)) *OrmHmRuleActions {
	ormAnd := NewOrmHmRuleActions()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmRuleActions) Or(fuc func(orm *OrmHmRuleActions)) *OrmHmRuleActions {
	ormOr := NewOrmHmRuleActions()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmRuleActions) Preload(query string, args ...interface{}) *OrmHmRuleActions {
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
func (orm *OrmHmRuleActions) Joins(query string, args ...interface{}) *OrmHmRuleActions {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmRuleActions) WhereActionid(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionid` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) InsertGetActionid(row *HmRuleActions) int32 {
	orm.db.Create(row)
	return row.Actionid
}
func (orm *OrmHmRuleActions) WhereActionidIn(val []int32) *OrmHmRuleActions {
	orm.db.Where("`actionid` IN ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionidGt(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionid` > ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionidGte(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionid` >= ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionidLt(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionid` < ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionidLte(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionid` <= ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionruleid(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionruleid` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionruleidIn(val []int32) *OrmHmRuleActions {
	orm.db.Where("`actionruleid` IN ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionruleidGt(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionruleid` > ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionruleidGte(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionruleid` >= ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionruleidLt(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionruleid` < ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionruleidLte(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionruleid` <= ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActiontype(val int32) *OrmHmRuleActions {
	orm.db.Where("`actiontype` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionimapfolder(val string) *OrmHmRuleActions {
	orm.db.Where("`actionimapfolder` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionsubject(val string) *OrmHmRuleActions {
	orm.db.Where("`actionsubject` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionfromname(val string) *OrmHmRuleActions {
	orm.db.Where("`actionfromname` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionfromaddress(val string) *OrmHmRuleActions {
	orm.db.Where("`actionfromaddress` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionto(val string) *OrmHmRuleActions {
	orm.db.Where("`actionto` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionbody(val string) *OrmHmRuleActions {
	orm.db.Where("`actionbody` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionfilename(val string) *OrmHmRuleActions {
	orm.db.Where("`actionfilename` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionsortorder(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionsortorder` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionscriptfunction(val string) *OrmHmRuleActions {
	orm.db.Where("`actionscriptfunction` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionheader(val string) *OrmHmRuleActions {
	orm.db.Where("`actionheader` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionvalue(val string) *OrmHmRuleActions {
	orm.db.Where("`actionvalue` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionrouteid(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionrouteid` = ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionrouteidIn(val []int32) *OrmHmRuleActions {
	orm.db.Where("`actionrouteid` IN ?", val)
	return orm
}
func (orm *OrmHmRuleActions) WhereActionrouteidNe(val int32) *OrmHmRuleActions {
	orm.db.Where("`actionrouteid` <> ?", val)
	return orm
}

type HmRuleActionsList []*HmRuleActions

func (l HmRuleActionsList) GetActionidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Actionid)
	}
	return got
}
func (l HmRuleActionsList) GetActionidMap() map[int32]*HmRuleActions {
	got := make(map[int32]*HmRuleActions)
	for _, val := range l {
		got[val.Actionid] = val
	}
	return got
}
func (l HmRuleActionsList) GetActionruleidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Actionruleid)
	}
	return got
}
func (l HmRuleActionsList) GetActionruleidMap() map[int32]*HmRuleActions {
	got := make(map[int32]*HmRuleActions)
	for _, val := range l {
		got[val.Actionruleid] = val
	}
	return got
}
func (l HmRuleActionsList) GetActionrouteidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Actionrouteid)
	}
	return got
}
func (l HmRuleActionsList) GetActionrouteidMap() map[int32]*HmRuleActions {
	got := make(map[int32]*HmRuleActions)
	for _, val := range l {
		got[val.Actionrouteid] = val
	}
	return got
}
