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

type HmMessageMetadata struct {
	MetadataId        int64          `gorm:"column:metadata_id;autoIncrement;type:bigint(20);primaryKey" json:"metadata_id"`                                                                                     //
	MetadataAccountid int32          `gorm:"column:metadata_accountid;type:int(11);index:idx_message_metadata_unique,class:BTREE,unique" json:"metadata_accountid"`                                              //
	MetadataFolderid  int32          `gorm:"column:metadata_folderid;type:int(11);index:idx_message_metadata_unique,class:BTREE,unique" json:"metadata_folderid"`                                                //
	MetadataMessageid int64          `gorm:"column:metadata_messageid;type:bigint(20);index:idx_message_metadata_unique,class:BTREE,unique;index:idx_message_metadata_id,class:BTREE" json:"metadata_messageid"` //
	MetadataDateutc   *database.Time `gorm:"column:metadata_dateutc;type:datetime" json:"metadata_dateutc"`                                                                                                      //
	MetadataFrom      string         `gorm:"column:metadata_from;type:varchar(255)" json:"metadata_from"`                                                                                                        //
	MetadataSubject   string         `gorm:"column:metadata_subject;type:varchar(255)" json:"metadata_subject"`                                                                                                  //
	MetadataTo        string         `gorm:"column:metadata_to;type:varchar(255)" json:"metadata_to"`                                                                                                            //
	MetadataCc        string         `gorm:"column:metadata_cc;type:varchar(255)" json:"metadata_cc"`                                                                                                            //
}

var (
	HmMessageMetadataFieldMetadataId        = "metadata_id"
	HmMessageMetadataFieldMetadataAccountid = "metadata_accountid"
	HmMessageMetadataFieldMetadataFolderid  = "metadata_folderid"
	HmMessageMetadataFieldMetadataMessageid = "metadata_messageid"
	HmMessageMetadataFieldMetadataDateutc   = "metadata_dateutc"
	HmMessageMetadataFieldMetadataFrom      = "metadata_from"
	HmMessageMetadataFieldMetadataSubject   = "metadata_subject"
	HmMessageMetadataFieldMetadataTo        = "metadata_to"
	HmMessageMetadataFieldMetadataCc        = "metadata_cc"
)

func (receiver *HmMessageMetadata) TableName() string {
	return "hm_message_metadata"
}

type OrmHmMessageMetadata struct {
	db *gorm.DB
}

func (orm *OrmHmMessageMetadata) TableName() string {
	return "hm_message_metadata"
}

func NewOrmHmMessageMetadata(txs ...*gorm.DB) *OrmHmMessageMetadata {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmMessageMetadata{})
	} else {
		tx = txs[0].Model(&HmMessageMetadata{})
	}
	return &OrmHmMessageMetadata{db: tx}
}

func (orm *OrmHmMessageMetadata) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmMessageMetadata) GetTableInfo() interface{} {
	return &HmMessageMetadata{}
}

// Create insert the value into database
func (orm *OrmHmMessageMetadata) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmMessageMetadata) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmMessageMetadata) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmMessageMetadata) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmMessageMetadata) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmMessageMetadata) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmMessageMetadata) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmMessageMetadata) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmMessageMetadata) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmMessageMetadata) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmMessageMetadata) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmMessageMetadata) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmMessageMetadata) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmMessageMetadata) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmMessageMetadata) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmMessageMetadata) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmMessageMetadata) Unscoped() *OrmHmMessageMetadata {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmMessageMetadata) Insert(row *HmMessageMetadata) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmMessageMetadata) Inserts(rows []*HmMessageMetadata) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmMessageMetadata) Order(value interface{}) *OrmHmMessageMetadata {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmMessageMetadata) Group(name string) *OrmHmMessageMetadata {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmMessageMetadata) Limit(limit int) *OrmHmMessageMetadata {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmMessageMetadata) Offset(offset int) *OrmHmMessageMetadata {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmMessageMetadata) Get() HmMessageMetadataList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmMessageMetadata) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmMessageMetadata) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmMessageMetadata{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmMessageMetadata) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_message_metadata")
}

func (orm *OrmHmMessageMetadata) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmMessageMetadata) First(conds ...interface{}) (*HmMessageMetadata, bool) {
	dest := &HmMessageMetadata{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmMessageMetadata) Take(conds ...interface{}) (*HmMessageMetadata, int64) {
	dest := &HmMessageMetadata{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmMessageMetadata) Last(conds ...interface{}) (*HmMessageMetadata, int64) {
	dest := &HmMessageMetadata{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmMessageMetadata) Find(conds ...interface{}) (HmMessageMetadataList, int64) {
	list := make([]*HmMessageMetadata, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmMessageMetadata) Paginate(page int, limit int) (HmMessageMetadataList, int64) {
	var total int64
	list := make([]*HmMessageMetadata, 0)
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
func (orm *OrmHmMessageMetadata) SimplePaginate(page int, limit int) HmMessageMetadataList {
	list := make([]*HmMessageMetadata, 0)
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
func (orm *OrmHmMessageMetadata) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmMessageMetadata) FirstOrInit(dest *HmMessageMetadata, conds ...interface{}) (*HmMessageMetadata, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmMessageMetadata) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmMessageMetadata) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmMessageMetadata) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmMessageMetadata) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmMessageMetadata) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmMessageMetadata) Where(query interface{}, args ...interface{}) *OrmHmMessageMetadata {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmMessageMetadata) Select(query interface{}, args ...interface{}) *OrmHmMessageMetadata {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmMessageMetadata) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmMessageMetadata) And(fuc func(orm *OrmHmMessageMetadata)) *OrmHmMessageMetadata {
	ormAnd := NewOrmHmMessageMetadata()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmMessageMetadata) Or(fuc func(orm *OrmHmMessageMetadata)) *OrmHmMessageMetadata {
	ormOr := NewOrmHmMessageMetadata()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmMessageMetadata) Preload(query string, args ...interface{}) *OrmHmMessageMetadata {
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
func (orm *OrmHmMessageMetadata) Joins(query string, args ...interface{}) *OrmHmMessageMetadata {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmMessageMetadata) WhereMetadataId(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_id` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) InsertGetMetadataId(row *HmMessageMetadata) int64 {
	orm.db.Create(row)
	return row.MetadataId
}
func (orm *OrmHmMessageMetadata) WhereMetadataIdIn(val []int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_id` IN ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataIdGt(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_id` > ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataIdGte(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_id` >= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataIdLt(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_id` < ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataIdLte(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_id` <= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataAccountid(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_accountid` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataAccountidIn(val []int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_accountid` IN ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataAccountidGt(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_accountid` > ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataAccountidGte(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_accountid` >= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataAccountidLt(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_accountid` < ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataAccountidLte(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_accountid` <= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataFolderid(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_folderid` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataFolderidIn(val []int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_folderid` IN ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataFolderidNe(val int32) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_folderid` <> ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataMessageid(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_messageid` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataMessageidIn(val []int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_messageid` IN ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataMessageidGt(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_messageid` > ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataMessageidGte(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_messageid` >= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataMessageidLt(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_messageid` < ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataMessageidLte(val int64) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_messageid` <= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataDateutc(val database.Time) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_dateutc` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataDateutcBetween(begin database.Time, end database.Time) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_dateutc` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataDateutcLte(val database.Time) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_dateutc` <= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataDateutcGte(val database.Time) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_dateutc` >= ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataFrom(val string) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_from` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataSubject(val string) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_subject` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataTo(val string) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_to` = ?", val)
	return orm
}
func (orm *OrmHmMessageMetadata) WhereMetadataCc(val string) *OrmHmMessageMetadata {
	orm.db.Where("`metadata_cc` = ?", val)
	return orm
}

type HmMessageMetadataList []*HmMessageMetadata

func (l HmMessageMetadataList) GetMetadataIdList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.MetadataId)
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataIdMap() map[int64]*HmMessageMetadata {
	got := make(map[int64]*HmMessageMetadata)
	for _, val := range l {
		got[val.MetadataId] = val
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataAccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.MetadataAccountid)
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataAccountidMap() map[int32]*HmMessageMetadata {
	got := make(map[int32]*HmMessageMetadata)
	for _, val := range l {
		got[val.MetadataAccountid] = val
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataFolderidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.MetadataFolderid)
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataFolderidMap() map[int32]*HmMessageMetadata {
	got := make(map[int32]*HmMessageMetadata)
	for _, val := range l {
		got[val.MetadataFolderid] = val
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataMessageidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.MetadataMessageid)
	}
	return got
}
func (l HmMessageMetadataList) GetMetadataMessageidMap() map[int64]*HmMessageMetadata {
	got := make(map[int64]*HmMessageMetadata)
	for _, val := range l {
		got[val.MetadataMessageid] = val
	}
	return got
}
