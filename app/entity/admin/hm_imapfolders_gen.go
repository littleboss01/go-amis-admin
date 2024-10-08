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

type HmImapfolders struct {
	Folderid           int64          `gorm:"column:folderid;autoIncrement;type:bigint(11);primaryKey;index:folderid,class:BTREE,unique" json:"folderid"`                                                        //
	Folderaccountid    int32          `gorm:"column:folderaccountid;type:int(10);index:idx_hm_imapfolders_unique,class:BTREE,unique;index:idx_hm_imapfolders,class:BTREE;comment:'账号id'" json:"folderaccountid"` // 账号id
	Folderparentid     int32          `gorm:"column:folderparentid;type:int(11);index:idx_hm_imapfolders_unique,class:BTREE,unique;default:-1" json:"folderparentid"`                                            //
	Foldername         string         `gorm:"column:foldername;type:varchar(255);index:idx_hm_imapfolders_unique,class:BTREE,unique;default:INBOX" json:"foldername"`                                            //
	Folderissubscribed int32          `gorm:"column:folderissubscribed;type:tinyint(3);default:1" json:"folderissubscribed"`                                                                                     //
	Foldercreationtime *database.Time `gorm:"column:foldercreationtime;type:datetime;default:2021-01-01 00:00:00" json:"foldercreationtime"`                                                                     //
	Foldercurrentuid   int64          `gorm:"column:foldercurrentuid;type:bigint(20);default:0" json:"foldercurrentuid"`                                                                                         //
}

var (
	HmImapfoldersFieldFolderid           = "folderid"
	HmImapfoldersFieldFolderaccountid    = "folderaccountid"
	HmImapfoldersFieldFolderparentid     = "folderparentid"
	HmImapfoldersFieldFoldername         = "foldername"
	HmImapfoldersFieldFolderissubscribed = "folderissubscribed"
	HmImapfoldersFieldFoldercreationtime = "foldercreationtime"
	HmImapfoldersFieldFoldercurrentuid   = "foldercurrentuid"
)

func (receiver *HmImapfolders) TableName() string {
	return "hm_imapfolders"
}

type OrmHmImapfolders struct {
	db *gorm.DB
}

func (orm *OrmHmImapfolders) TableName() string {
	return "hm_imapfolders"
}

func NewOrmHmImapfolders(txs ...*gorm.DB) *OrmHmImapfolders {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmImapfolders{})
	} else {
		tx = txs[0].Model(&HmImapfolders{})
	}
	return &OrmHmImapfolders{db: tx}
}

func (orm *OrmHmImapfolders) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmImapfolders) GetTableInfo() interface{} {
	return &HmImapfolders{}
}

// Create insert the value into database
func (orm *OrmHmImapfolders) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmImapfolders) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmImapfolders) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmImapfolders) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmImapfolders) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmImapfolders) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmImapfolders) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmImapfolders) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmImapfolders) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmImapfolders) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmImapfolders) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmImapfolders) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmImapfolders) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmImapfolders) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmImapfolders) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmImapfolders) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmImapfolders) Unscoped() *OrmHmImapfolders {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmImapfolders) Insert(row *HmImapfolders) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmImapfolders) Inserts(rows []*HmImapfolders) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmImapfolders) Order(value interface{}) *OrmHmImapfolders {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmImapfolders) Group(name string) *OrmHmImapfolders {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmImapfolders) Limit(limit int) *OrmHmImapfolders {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmImapfolders) Offset(offset int) *OrmHmImapfolders {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmImapfolders) Get() HmImapfoldersList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmImapfolders) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmImapfolders) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmImapfolders{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmImapfolders) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_imapfolders")
}

func (orm *OrmHmImapfolders) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmImapfolders) First(conds ...interface{}) (*HmImapfolders, bool) {
	dest := &HmImapfolders{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmImapfolders) Take(conds ...interface{}) (*HmImapfolders, int64) {
	dest := &HmImapfolders{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmImapfolders) Last(conds ...interface{}) (*HmImapfolders, int64) {
	dest := &HmImapfolders{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmImapfolders) Find(conds ...interface{}) (HmImapfoldersList, int64) {
	list := make([]*HmImapfolders, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmImapfolders) Paginate(page int, limit int) (HmImapfoldersList, int64) {
	var total int64
	list := make([]*HmImapfolders, 0)
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
func (orm *OrmHmImapfolders) SimplePaginate(page int, limit int) HmImapfoldersList {
	list := make([]*HmImapfolders, 0)
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
func (orm *OrmHmImapfolders) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmImapfolders) FirstOrInit(dest *HmImapfolders, conds ...interface{}) (*HmImapfolders, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmImapfolders) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmImapfolders) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmImapfolders) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmImapfolders) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmImapfolders) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmImapfolders) Where(query interface{}, args ...interface{}) *OrmHmImapfolders {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmImapfolders) Select(query interface{}, args ...interface{}) *OrmHmImapfolders {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmImapfolders) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmImapfolders) And(fuc func(orm *OrmHmImapfolders)) *OrmHmImapfolders {
	ormAnd := NewOrmHmImapfolders()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmImapfolders) Or(fuc func(orm *OrmHmImapfolders)) *OrmHmImapfolders {
	ormOr := NewOrmHmImapfolders()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmImapfolders) Preload(query string, args ...interface{}) *OrmHmImapfolders {
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
func (orm *OrmHmImapfolders) Joins(query string, args ...interface{}) *OrmHmImapfolders {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmImapfolders) WhereFolderid(val int64) *OrmHmImapfolders {
	orm.db.Where("`folderid` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) InsertGetFolderid(row *HmImapfolders) int64 {
	orm.db.Create(row)
	return row.Folderid
}
func (orm *OrmHmImapfolders) WhereFolderidIn(val []int64) *OrmHmImapfolders {
	orm.db.Where("`folderid` IN ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderidGt(val int64) *OrmHmImapfolders {
	orm.db.Where("`folderid` > ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderidGte(val int64) *OrmHmImapfolders {
	orm.db.Where("`folderid` >= ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderidLt(val int64) *OrmHmImapfolders {
	orm.db.Where("`folderid` < ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderidLte(val int64) *OrmHmImapfolders {
	orm.db.Where("`folderid` <= ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderaccountid(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderaccountid` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderaccountidIn(val []int32) *OrmHmImapfolders {
	orm.db.Where("`folderaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderaccountidGt(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderaccountid` > ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderaccountidGte(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderaccountid` >= ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderaccountidLt(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderaccountid` < ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderaccountidLte(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderaccountid` <= ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderparentid(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderparentid` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderparentidIn(val []int32) *OrmHmImapfolders {
	orm.db.Where("`folderparentid` IN ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderparentidNe(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderparentid` <> ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldername(val string) *OrmHmImapfolders {
	orm.db.Where("`foldername` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFolderissubscribed(val int32) *OrmHmImapfolders {
	orm.db.Where("`folderissubscribed` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercreationtime(val database.Time) *OrmHmImapfolders {
	orm.db.Where("`foldercreationtime` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercreationtimeBetween(begin database.Time, end database.Time) *OrmHmImapfolders {
	orm.db.Where("`foldercreationtime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercreationtimeLte(val database.Time) *OrmHmImapfolders {
	orm.db.Where("`foldercreationtime` <= ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercreationtimeGte(val database.Time) *OrmHmImapfolders {
	orm.db.Where("`foldercreationtime` >= ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercurrentuid(val int64) *OrmHmImapfolders {
	orm.db.Where("`foldercurrentuid` = ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercurrentuidIn(val []int64) *OrmHmImapfolders {
	orm.db.Where("`foldercurrentuid` IN ?", val)
	return orm
}
func (orm *OrmHmImapfolders) WhereFoldercurrentuidNe(val int64) *OrmHmImapfolders {
	orm.db.Where("`foldercurrentuid` <> ?", val)
	return orm
}

type HmImapfoldersList []*HmImapfolders

func (l HmImapfoldersList) GetFolderidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Folderid)
	}
	return got
}
func (l HmImapfoldersList) GetFolderidMap() map[int64]*HmImapfolders {
	got := make(map[int64]*HmImapfolders)
	for _, val := range l {
		got[val.Folderid] = val
	}
	return got
}
func (l HmImapfoldersList) GetFolderaccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Folderaccountid)
	}
	return got
}
func (l HmImapfoldersList) GetFolderaccountidMap() map[int32]*HmImapfolders {
	got := make(map[int32]*HmImapfolders)
	for _, val := range l {
		got[val.Folderaccountid] = val
	}
	return got
}
func (l HmImapfoldersList) GetFolderparentidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Folderparentid)
	}
	return got
}
func (l HmImapfoldersList) GetFolderparentidMap() map[int32]*HmImapfolders {
	got := make(map[int32]*HmImapfolders)
	for _, val := range l {
		got[val.Folderparentid] = val
	}
	return got
}
func (l HmImapfoldersList) GetFoldercurrentuidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Foldercurrentuid)
	}
	return got
}
func (l HmImapfoldersList) GetFoldercurrentuidMap() map[int64]*HmImapfolders {
	got := make(map[int64]*HmImapfolders)
	for _, val := range l {
		got[val.Foldercurrentuid] = val
	}
	return got
}
