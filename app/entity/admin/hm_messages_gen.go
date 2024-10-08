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

type HmMessages struct {
	Messageid           int64         `gorm:"column:messageid;autoIncrement;type:bigint(20);primaryKey;index:messageid,class:BTREE,unique" json:"messageid"` //
	Messageaccountid    int32         `gorm:"column:messageaccountid;type:int(11);index:idx_hm_messages,class:BTREE" json:"messageaccountid"`                //
	Messagefolderid     int32         `gorm:"column:messagefolderid;type:int(11);index:idx_hm_messages,class:BTREE;default:0" json:"messagefolderid"`        //
	Messagefilename     string        `gorm:"column:messagefilename;type:varchar(255)" json:"messagefilename"`                                               //
	Messagetype         int32         `gorm:"column:messagetype;type:tinyint(4);index:idx_hm_messages_type,class:BTREE" json:"messagetype"`                  //
	Messagefrom         string        `gorm:"column:messagefrom;type:varchar(255)" json:"messagefrom"`                                                       //
	Messagesize         int64         `gorm:"column:messagesize;type:bigint(20)" json:"messagesize"`                                                         //
	Messagecurnooftries int32         `gorm:"column:messagecurnooftries;type:int(11)" json:"messagecurnooftries"`                                            //
	Messagenexttrytime  database.Time `gorm:"column:messagenexttrytime;type:datetime" json:"messagenexttrytime"`                                             //
	Messageflags        int32         `gorm:"column:messageflags;type:tinyint(4)" json:"messageflags"`                                                       //
	Messagecreatetime   database.Time `gorm:"column:messagecreatetime;type:datetime" json:"messagecreatetime"`                                               //
	Messagelocked       int32         `gorm:"column:messagelocked;type:tinyint(4)" json:"messagelocked"`                                                     //
	Messageuid          int64         `gorm:"column:messageuid;type:bigint(20)" json:"messageuid"`                                                           //
}

var (
	HmMessagesFieldMessageid           = "messageid"
	HmMessagesFieldMessageaccountid    = "messageaccountid"
	HmMessagesFieldMessagefolderid     = "messagefolderid"
	HmMessagesFieldMessagefilename     = "messagefilename"
	HmMessagesFieldMessagetype         = "messagetype"
	HmMessagesFieldMessagefrom         = "messagefrom"
	HmMessagesFieldMessagesize         = "messagesize"
	HmMessagesFieldMessagecurnooftries = "messagecurnooftries"
	HmMessagesFieldMessagenexttrytime  = "messagenexttrytime"
	HmMessagesFieldMessageflags        = "messageflags"
	HmMessagesFieldMessagecreatetime   = "messagecreatetime"
	HmMessagesFieldMessagelocked       = "messagelocked"
	HmMessagesFieldMessageuid          = "messageuid"
)

func (receiver *HmMessages) TableName() string {
	return "hm_messages"
}

type OrmHmMessages struct {
	db *gorm.DB
}

func (orm *OrmHmMessages) TableName() string {
	return "hm_messages"
}

func NewOrmHmMessages(txs ...*gorm.DB) *OrmHmMessages {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmMessages{})
	} else {
		tx = txs[0].Model(&HmMessages{})
	}
	return &OrmHmMessages{db: tx}
}

func (orm *OrmHmMessages) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmMessages) GetTableInfo() interface{} {
	return &HmMessages{}
}

// Create insert the value into database
func (orm *OrmHmMessages) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmMessages) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmMessages) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmMessages) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmMessages) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmMessages) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmMessages) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmMessages) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmMessages) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmMessages) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmMessages) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmMessages) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmMessages) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmMessages) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmMessages) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmMessages) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmMessages) Unscoped() *OrmHmMessages {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmMessages) Insert(row *HmMessages) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmMessages) Inserts(rows []*HmMessages) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmMessages) Order(value interface{}) *OrmHmMessages {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmMessages) Group(name string) *OrmHmMessages {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmMessages) Limit(limit int) *OrmHmMessages {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmMessages) Offset(offset int) *OrmHmMessages {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmMessages) Get() HmMessagesList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmMessages) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmMessages) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmMessages{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmMessages) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_messages")
}

func (orm *OrmHmMessages) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmMessages) First(conds ...interface{}) (*HmMessages, bool) {
	dest := &HmMessages{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmMessages) Take(conds ...interface{}) (*HmMessages, int64) {
	dest := &HmMessages{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmMessages) Last(conds ...interface{}) (*HmMessages, int64) {
	dest := &HmMessages{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmMessages) Find(conds ...interface{}) (HmMessagesList, int64) {
	list := make([]*HmMessages, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmMessages) Paginate(page int, limit int) (HmMessagesList, int64) {
	var total int64
	list := make([]*HmMessages, 0)
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
func (orm *OrmHmMessages) SimplePaginate(page int, limit int) HmMessagesList {
	list := make([]*HmMessages, 0)
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
func (orm *OrmHmMessages) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmMessages) FirstOrInit(dest *HmMessages, conds ...interface{}) (*HmMessages, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmMessages) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmMessages) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmMessages) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmMessages) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmMessages) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmMessages) Where(query interface{}, args ...interface{}) *OrmHmMessages {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmMessages) Select(query interface{}, args ...interface{}) *OrmHmMessages {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmMessages) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmMessages) And(fuc func(orm *OrmHmMessages)) *OrmHmMessages {
	ormAnd := NewOrmHmMessages()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmMessages) Or(fuc func(orm *OrmHmMessages)) *OrmHmMessages {
	ormOr := NewOrmHmMessages()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmMessages) Preload(query string, args ...interface{}) *OrmHmMessages {
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
func (orm *OrmHmMessages) Joins(query string, args ...interface{}) *OrmHmMessages {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmMessages) WhereMessageid(val int64) *OrmHmMessages {
	orm.db.Where("`messageid` = ?", val)
	return orm
}
func (orm *OrmHmMessages) InsertGetMessageid(row *HmMessages) int64 {
	orm.db.Create(row)
	return row.Messageid
}
func (orm *OrmHmMessages) WhereMessageidIn(val []int64) *OrmHmMessages {
	orm.db.Where("`messageid` IN ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageidGt(val int64) *OrmHmMessages {
	orm.db.Where("`messageid` > ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageidGte(val int64) *OrmHmMessages {
	orm.db.Where("`messageid` >= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageidLt(val int64) *OrmHmMessages {
	orm.db.Where("`messageid` < ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageidLte(val int64) *OrmHmMessages {
	orm.db.Where("`messageid` <= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageaccountid(val int32) *OrmHmMessages {
	orm.db.Where("`messageaccountid` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageaccountidIn(val []int32) *OrmHmMessages {
	orm.db.Where("`messageaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageaccountidGt(val int32) *OrmHmMessages {
	orm.db.Where("`messageaccountid` > ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageaccountidGte(val int32) *OrmHmMessages {
	orm.db.Where("`messageaccountid` >= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageaccountidLt(val int32) *OrmHmMessages {
	orm.db.Where("`messageaccountid` < ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageaccountidLte(val int32) *OrmHmMessages {
	orm.db.Where("`messageaccountid` <= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagefolderid(val int32) *OrmHmMessages {
	orm.db.Where("`messagefolderid` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagefolderidIn(val []int32) *OrmHmMessages {
	orm.db.Where("`messagefolderid` IN ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagefolderidNe(val int32) *OrmHmMessages {
	orm.db.Where("`messagefolderid` <> ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagefilename(val string) *OrmHmMessages {
	orm.db.Where("`messagefilename` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagetype(val int32) *OrmHmMessages {
	orm.db.Where("`messagetype` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagetypeIn(val []int32) *OrmHmMessages {
	orm.db.Where("`messagetype` IN ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagetypeGt(val int32) *OrmHmMessages {
	orm.db.Where("`messagetype` > ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagetypeGte(val int32) *OrmHmMessages {
	orm.db.Where("`messagetype` >= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagetypeLt(val int32) *OrmHmMessages {
	orm.db.Where("`messagetype` < ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagetypeLte(val int32) *OrmHmMessages {
	orm.db.Where("`messagetype` <= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagefrom(val string) *OrmHmMessages {
	orm.db.Where("`messagefrom` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagesize(val int64) *OrmHmMessages {
	orm.db.Where("`messagesize` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagecurnooftries(val int32) *OrmHmMessages {
	orm.db.Where("`messagecurnooftries` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagenexttrytime(val database.Time) *OrmHmMessages {
	orm.db.Where("`messagenexttrytime` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagenexttrytimeBetween(begin database.Time, end database.Time) *OrmHmMessages {
	orm.db.Where("`messagenexttrytime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmMessages) WhereMessagenexttrytimeLte(val database.Time) *OrmHmMessages {
	orm.db.Where("`messagenexttrytime` <= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagenexttrytimeGte(val database.Time) *OrmHmMessages {
	orm.db.Where("`messagenexttrytime` >= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageflags(val int32) *OrmHmMessages {
	orm.db.Where("`messageflags` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagecreatetime(val database.Time) *OrmHmMessages {
	orm.db.Where("`messagecreatetime` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagecreatetimeBetween(begin database.Time, end database.Time) *OrmHmMessages {
	orm.db.Where("`messagecreatetime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmMessages) WhereMessagecreatetimeLte(val database.Time) *OrmHmMessages {
	orm.db.Where("`messagecreatetime` <= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagecreatetimeGte(val database.Time) *OrmHmMessages {
	orm.db.Where("`messagecreatetime` >= ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessagelocked(val int32) *OrmHmMessages {
	orm.db.Where("`messagelocked` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageuid(val int64) *OrmHmMessages {
	orm.db.Where("`messageuid` = ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageuidIn(val []int64) *OrmHmMessages {
	orm.db.Where("`messageuid` IN ?", val)
	return orm
}
func (orm *OrmHmMessages) WhereMessageuidNe(val int64) *OrmHmMessages {
	orm.db.Where("`messageuid` <> ?", val)
	return orm
}

type HmMessagesList []*HmMessages

func (l HmMessagesList) GetMessageidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Messageid)
	}
	return got
}
func (l HmMessagesList) GetMessageidMap() map[int64]*HmMessages {
	got := make(map[int64]*HmMessages)
	for _, val := range l {
		got[val.Messageid] = val
	}
	return got
}
func (l HmMessagesList) GetMessageaccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Messageaccountid)
	}
	return got
}
func (l HmMessagesList) GetMessageaccountidMap() map[int32]*HmMessages {
	got := make(map[int32]*HmMessages)
	for _, val := range l {
		got[val.Messageaccountid] = val
	}
	return got
}
func (l HmMessagesList) GetMessagefolderidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Messagefolderid)
	}
	return got
}
func (l HmMessagesList) GetMessagefolderidMap() map[int32]*HmMessages {
	got := make(map[int32]*HmMessages)
	for _, val := range l {
		got[val.Messagefolderid] = val
	}
	return got
}
func (l HmMessagesList) GetMessageuidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Messageuid)
	}
	return got
}
func (l HmMessagesList) GetMessageuidMap() map[int64]*HmMessages {
	got := make(map[int64]*HmMessages)
	for _, val := range l {
		got[val.Messageuid] = val
	}
	return got
}
