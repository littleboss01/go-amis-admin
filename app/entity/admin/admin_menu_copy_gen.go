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

type AdminMenuCopy struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey;comment:'ID'" json:"id"`    // ID
	ParentId  uint32         `gorm:"column:parent_id;type:int(10) unsigned;default:0;comment:'父级'" json:"parent_id"`     // 父级
	Name      string         `gorm:"column:name;type:varchar(50);comment:'组件名称'" json:"name"`                            // 组件名称
	Component string         `gorm:"column:component;type:varchar(50);comment:'组件'" json:"component"`                    // 组件
	Path      *string        `gorm:"column:path;type:varchar(255);default:;comment:'地址'" json:"path"`                    // 地址
	Redirect  *string        `gorm:"column:redirect;type:varchar(255);default:;comment:'重定向'" json:"redirect"`           // 重定向
	Meta      database.JSON  `gorm:"column:meta;type:json;comment:'元数据'" json:"meta"`                                    // 元数据
	Hidden    uint32         `gorm:"column:hidden;type:tinyint(3) unsigned;default:2;comment:'0忽略1隐藏2显示'" json:"hidden"` // 0忽略1隐藏2显示
	Sort      uint32         `gorm:"column:sort;type:int(10) unsigned;default:0;comment:'排序'" json:"sort"`               // 排序
	ApiList   *database.JSON `gorm:"column:api_list;type:json;comment:'api'" json:"api_list"`                            // api
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`            // created_at
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`            // updated_at
}

var (
	AdminMenuCopyFieldId        = "id"
	AdminMenuCopyFieldParentId  = "parent_id"
	AdminMenuCopyFieldName      = "name"
	AdminMenuCopyFieldComponent = "component"
	AdminMenuCopyFieldPath      = "path"
	AdminMenuCopyFieldRedirect  = "redirect"
	AdminMenuCopyFieldMeta      = "meta"
	AdminMenuCopyFieldHidden    = "hidden"
	AdminMenuCopyFieldSort      = "sort"
	AdminMenuCopyFieldApiList   = "api_list"
	AdminMenuCopyFieldCreatedAt = "created_at"
	AdminMenuCopyFieldUpdatedAt = "updated_at"
)

func (receiver *AdminMenuCopy) TableName() string {
	return "admin_menu_copy"
}

type OrmAdminMenuCopy struct {
	db *gorm.DB
}

func (orm *OrmAdminMenuCopy) TableName() string {
	return "admin_menu_copy"
}

func NewOrmAdminMenuCopy(txs ...*gorm.DB) *OrmAdminMenuCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminMenuCopy{})
	} else {
		tx = txs[0].Model(&AdminMenuCopy{})
	}
	return &OrmAdminMenuCopy{db: tx}
}

func (orm *OrmAdminMenuCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminMenuCopy) GetTableInfo() interface{} {
	return &AdminMenuCopy{}
}

// Create insert the value into database
func (orm *OrmAdminMenuCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminMenuCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminMenuCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminMenuCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminMenuCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminMenuCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminMenuCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminMenuCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminMenuCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminMenuCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminMenuCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminMenuCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminMenuCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminMenuCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminMenuCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminMenuCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminMenuCopy) Unscoped() *OrmAdminMenuCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminMenuCopy) Insert(row *AdminMenuCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminMenuCopy) Inserts(rows []*AdminMenuCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminMenuCopy) Order(value interface{}) *OrmAdminMenuCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminMenuCopy) Group(name string) *OrmAdminMenuCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminMenuCopy) Limit(limit int) *OrmAdminMenuCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminMenuCopy) Offset(offset int) *OrmAdminMenuCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminMenuCopy) Get() AdminMenuCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminMenuCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminMenuCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminMenuCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminMenuCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_menu_copy")
}

func (orm *OrmAdminMenuCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminMenuCopy) First(conds ...interface{}) (*AdminMenuCopy, bool) {
	dest := &AdminMenuCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminMenuCopy) Take(conds ...interface{}) (*AdminMenuCopy, int64) {
	dest := &AdminMenuCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminMenuCopy) Last(conds ...interface{}) (*AdminMenuCopy, int64) {
	dest := &AdminMenuCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminMenuCopy) Find(conds ...interface{}) (AdminMenuCopyList, int64) {
	list := make([]*AdminMenuCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminMenuCopy) Paginate(page int, limit int) (AdminMenuCopyList, int64) {
	var total int64
	list := make([]*AdminMenuCopy, 0)
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
func (orm *OrmAdminMenuCopy) SimplePaginate(page int, limit int) AdminMenuCopyList {
	list := make([]*AdminMenuCopy, 0)
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
func (orm *OrmAdminMenuCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminMenuCopy) FirstOrInit(dest *AdminMenuCopy, conds ...interface{}) (*AdminMenuCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminMenuCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminMenuCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminMenuCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminMenuCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminMenuCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminMenuCopy) Where(query interface{}, args ...interface{}) *OrmAdminMenuCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminMenuCopy) Select(query interface{}, args ...interface{}) *OrmAdminMenuCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminMenuCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminMenuCopy) And(fuc func(orm *OrmAdminMenuCopy)) *OrmAdminMenuCopy {
	ormAnd := NewOrmAdminMenuCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminMenuCopy) Or(fuc func(orm *OrmAdminMenuCopy)) *OrmAdminMenuCopy {
	ormOr := NewOrmAdminMenuCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminMenuCopy) Preload(query string, args ...interface{}) *OrmAdminMenuCopy {
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
func (orm *OrmAdminMenuCopy) Joins(query string, args ...interface{}) *OrmAdminMenuCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminMenuCopy) WhereId(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) InsertGetId(row *AdminMenuCopy) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAdminMenuCopy) WhereIdIn(val []uint32) *OrmAdminMenuCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereIdGt(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereIdGte(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereIdLt(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereIdLte(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereParentId(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`parent_id` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereParentIdIn(val []uint32) *OrmAdminMenuCopy {
	orm.db.Where("`parent_id` IN ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereParentIdNe(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`parent_id` <> ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereName(val string) *OrmAdminMenuCopy {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereComponent(val string) *OrmAdminMenuCopy {
	orm.db.Where("`component` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WherePath(val string) *OrmAdminMenuCopy {
	orm.db.Where("`path` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereRedirect(val string) *OrmAdminMenuCopy {
	orm.db.Where("`redirect` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereMeta(val database.JSON) *OrmAdminMenuCopy {
	orm.db.Where("`meta` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereHidden(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`hidden` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereHiddenIn(val []uint32) *OrmAdminMenuCopy {
	orm.db.Where("`hidden` IN ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereHiddenNe(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`hidden` <> ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereSort(val uint32) *OrmAdminMenuCopy {
	orm.db.Where("`sort` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereApiList(val database.JSON) *OrmAdminMenuCopy {
	orm.db.Where("`api_list` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereCreatedAt(val database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereCreatedAtLte(val database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereCreatedAtGte(val database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereUpdatedAt(val database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereUpdatedAtLte(val database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminMenuCopy) WhereUpdatedAtGte(val database.Time) *OrmAdminMenuCopy {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminMenuCopyList []*AdminMenuCopy

func (l AdminMenuCopyList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminMenuCopyList) GetIdMap() map[uint32]*AdminMenuCopy {
	got := make(map[uint32]*AdminMenuCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l AdminMenuCopyList) GetParentIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.ParentId)
	}
	return got
}
func (l AdminMenuCopyList) GetParentIdMap() map[uint32]*AdminMenuCopy {
	got := make(map[uint32]*AdminMenuCopy)
	for _, val := range l {
		got[val.ParentId] = val
	}
	return got
}
func (l AdminMenuCopyList) GetHiddenList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Hidden)
	}
	return got
}
func (l AdminMenuCopyList) GetHiddenMap() map[uint32]*AdminMenuCopy {
	got := make(map[uint32]*AdminMenuCopy)
	for _, val := range l {
		got[val.Hidden] = val
	}
	return got
}
