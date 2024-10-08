package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmBlockedAttachments struct {
	Baid          int64  `gorm:"column:baid;autoIncrement;type:bigint(20);primaryKey;index:baid,class:BTREE,unique" json:"baid"` //
	Bawildcard    string `gorm:"column:bawildcard;type:varchar(255)" json:"bawildcard"`                                          //
	Badescription string `gorm:"column:badescription;type:varchar(255)" json:"badescription"`                                    //
}

var (
	HmBlockedAttachmentsFieldBaid          = "baid"
	HmBlockedAttachmentsFieldBawildcard    = "bawildcard"
	HmBlockedAttachmentsFieldBadescription = "badescription"
)

func (receiver *HmBlockedAttachments) TableName() string {
	return "hm_blocked_attachments"
}

type OrmHmBlockedAttachments struct {
	db *gorm.DB
}

func (orm *OrmHmBlockedAttachments) TableName() string {
	return "hm_blocked_attachments"
}

func NewOrmHmBlockedAttachments(txs ...*gorm.DB) *OrmHmBlockedAttachments {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmBlockedAttachments{})
	} else {
		tx = txs[0].Model(&HmBlockedAttachments{})
	}
	return &OrmHmBlockedAttachments{db: tx}
}

func (orm *OrmHmBlockedAttachments) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmBlockedAttachments) GetTableInfo() interface{} {
	return &HmBlockedAttachments{}
}

// Create insert the value into database
func (orm *OrmHmBlockedAttachments) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmBlockedAttachments) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmBlockedAttachments) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmBlockedAttachments) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmBlockedAttachments) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmBlockedAttachments) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmBlockedAttachments) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmBlockedAttachments) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmBlockedAttachments) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmBlockedAttachments) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmBlockedAttachments) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmBlockedAttachments) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmBlockedAttachments) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmBlockedAttachments) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmBlockedAttachments) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmBlockedAttachments) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmBlockedAttachments) Unscoped() *OrmHmBlockedAttachments {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmBlockedAttachments) Insert(row *HmBlockedAttachments) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmBlockedAttachments) Inserts(rows []*HmBlockedAttachments) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmBlockedAttachments) Order(value interface{}) *OrmHmBlockedAttachments {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmBlockedAttachments) Group(name string) *OrmHmBlockedAttachments {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmBlockedAttachments) Limit(limit int) *OrmHmBlockedAttachments {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmBlockedAttachments) Offset(offset int) *OrmHmBlockedAttachments {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmBlockedAttachments) Get() HmBlockedAttachmentsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmBlockedAttachments) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmBlockedAttachments) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmBlockedAttachments{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmBlockedAttachments) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_blocked_attachments")
}

func (orm *OrmHmBlockedAttachments) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmBlockedAttachments) First(conds ...interface{}) (*HmBlockedAttachments, bool) {
	dest := &HmBlockedAttachments{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmBlockedAttachments) Take(conds ...interface{}) (*HmBlockedAttachments, int64) {
	dest := &HmBlockedAttachments{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmBlockedAttachments) Last(conds ...interface{}) (*HmBlockedAttachments, int64) {
	dest := &HmBlockedAttachments{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmBlockedAttachments) Find(conds ...interface{}) (HmBlockedAttachmentsList, int64) {
	list := make([]*HmBlockedAttachments, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmBlockedAttachments) Paginate(page int, limit int) (HmBlockedAttachmentsList, int64) {
	var total int64
	list := make([]*HmBlockedAttachments, 0)
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
func (orm *OrmHmBlockedAttachments) SimplePaginate(page int, limit int) HmBlockedAttachmentsList {
	list := make([]*HmBlockedAttachments, 0)
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
func (orm *OrmHmBlockedAttachments) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmBlockedAttachments) FirstOrInit(dest *HmBlockedAttachments, conds ...interface{}) (*HmBlockedAttachments, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmBlockedAttachments) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmBlockedAttachments) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmBlockedAttachments) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmBlockedAttachments) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmBlockedAttachments) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmBlockedAttachments) Where(query interface{}, args ...interface{}) *OrmHmBlockedAttachments {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmBlockedAttachments) Select(query interface{}, args ...interface{}) *OrmHmBlockedAttachments {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmBlockedAttachments) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmBlockedAttachments) And(fuc func(orm *OrmHmBlockedAttachments)) *OrmHmBlockedAttachments {
	ormAnd := NewOrmHmBlockedAttachments()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmBlockedAttachments) Or(fuc func(orm *OrmHmBlockedAttachments)) *OrmHmBlockedAttachments {
	ormOr := NewOrmHmBlockedAttachments()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmBlockedAttachments) Preload(query string, args ...interface{}) *OrmHmBlockedAttachments {
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
func (orm *OrmHmBlockedAttachments) Joins(query string, args ...interface{}) *OrmHmBlockedAttachments {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmBlockedAttachments) WhereBaid(val int64) *OrmHmBlockedAttachments {
	orm.db.Where("`baid` = ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) InsertGetBaid(row *HmBlockedAttachments) int64 {
	orm.db.Create(row)
	return row.Baid
}
func (orm *OrmHmBlockedAttachments) WhereBaidIn(val []int64) *OrmHmBlockedAttachments {
	orm.db.Where("`baid` IN ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) WhereBaidGt(val int64) *OrmHmBlockedAttachments {
	orm.db.Where("`baid` > ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) WhereBaidGte(val int64) *OrmHmBlockedAttachments {
	orm.db.Where("`baid` >= ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) WhereBaidLt(val int64) *OrmHmBlockedAttachments {
	orm.db.Where("`baid` < ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) WhereBaidLte(val int64) *OrmHmBlockedAttachments {
	orm.db.Where("`baid` <= ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) WhereBawildcard(val string) *OrmHmBlockedAttachments {
	orm.db.Where("`bawildcard` = ?", val)
	return orm
}
func (orm *OrmHmBlockedAttachments) WhereBadescription(val string) *OrmHmBlockedAttachments {
	orm.db.Where("`badescription` = ?", val)
	return orm
}

type HmBlockedAttachmentsList []*HmBlockedAttachments

func (l HmBlockedAttachmentsList) GetBaidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Baid)
	}
	return got
}
func (l HmBlockedAttachmentsList) GetBaidMap() map[int64]*HmBlockedAttachments {
	got := make(map[int64]*HmBlockedAttachments)
	for _, val := range l {
		got[val.Baid] = val
	}
	return got
}
