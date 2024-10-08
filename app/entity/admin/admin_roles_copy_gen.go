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

type AdminRolesCopy struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey;comment:'id'" json:"id"`                                 // id
	Name      string         `gorm:"column:name;type:varchar(50);index:laravel_admin_roles_name_unique,class:BTREE,unique;comment:'角色名'" json:"name"` // 角色名
	Slug      string         `gorm:"column:slug;type:varchar(50);default:;comment:'默认权限'" json:"slug"`                                                // 默认权限
	Remark    string         `gorm:"column:remark;type:varchar(255);default:;comment:'备注'" json:"remark"`                                             // 备注
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`                                         // created_at
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`                                         // updated_at
}

var (
	AdminRolesCopyFieldId        = "id"
	AdminRolesCopyFieldName      = "name"
	AdminRolesCopyFieldSlug      = "slug"
	AdminRolesCopyFieldRemark    = "remark"
	AdminRolesCopyFieldCreatedAt = "created_at"
	AdminRolesCopyFieldUpdatedAt = "updated_at"
)

func (receiver *AdminRolesCopy) TableName() string {
	return "admin_roles_copy"
}

type OrmAdminRolesCopy struct {
	db *gorm.DB
}

func (orm *OrmAdminRolesCopy) TableName() string {
	return "admin_roles_copy"
}

func NewOrmAdminRolesCopy(txs ...*gorm.DB) *OrmAdminRolesCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminRolesCopy{})
	} else {
		tx = txs[0].Model(&AdminRolesCopy{})
	}
	return &OrmAdminRolesCopy{db: tx}
}

func (orm *OrmAdminRolesCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminRolesCopy) GetTableInfo() interface{} {
	return &AdminRolesCopy{}
}

// Create insert the value into database
func (orm *OrmAdminRolesCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminRolesCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminRolesCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminRolesCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminRolesCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminRolesCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminRolesCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminRolesCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminRolesCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminRolesCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminRolesCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminRolesCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminRolesCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminRolesCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminRolesCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminRolesCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminRolesCopy) Unscoped() *OrmAdminRolesCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminRolesCopy) Insert(row *AdminRolesCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminRolesCopy) Inserts(rows []*AdminRolesCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminRolesCopy) Order(value interface{}) *OrmAdminRolesCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminRolesCopy) Group(name string) *OrmAdminRolesCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminRolesCopy) Limit(limit int) *OrmAdminRolesCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminRolesCopy) Offset(offset int) *OrmAdminRolesCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminRolesCopy) Get() AdminRolesCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminRolesCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminRolesCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminRolesCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminRolesCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_roles_copy")
}

func (orm *OrmAdminRolesCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminRolesCopy) First(conds ...interface{}) (*AdminRolesCopy, bool) {
	dest := &AdminRolesCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminRolesCopy) Take(conds ...interface{}) (*AdminRolesCopy, int64) {
	dest := &AdminRolesCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminRolesCopy) Last(conds ...interface{}) (*AdminRolesCopy, int64) {
	dest := &AdminRolesCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminRolesCopy) Find(conds ...interface{}) (AdminRolesCopyList, int64) {
	list := make([]*AdminRolesCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminRolesCopy) Paginate(page int, limit int) (AdminRolesCopyList, int64) {
	var total int64
	list := make([]*AdminRolesCopy, 0)
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
func (orm *OrmAdminRolesCopy) SimplePaginate(page int, limit int) AdminRolesCopyList {
	list := make([]*AdminRolesCopy, 0)
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
func (orm *OrmAdminRolesCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminRolesCopy) FirstOrInit(dest *AdminRolesCopy, conds ...interface{}) (*AdminRolesCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminRolesCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRolesCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRolesCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminRolesCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminRolesCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminRolesCopy) Where(query interface{}, args ...interface{}) *OrmAdminRolesCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminRolesCopy) Select(query interface{}, args ...interface{}) *OrmAdminRolesCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminRolesCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminRolesCopy) And(fuc func(orm *OrmAdminRolesCopy)) *OrmAdminRolesCopy {
	ormAnd := NewOrmAdminRolesCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminRolesCopy) Or(fuc func(orm *OrmAdminRolesCopy)) *OrmAdminRolesCopy {
	ormOr := NewOrmAdminRolesCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminRolesCopy) Preload(query string, args ...interface{}) *OrmAdminRolesCopy {
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
func (orm *OrmAdminRolesCopy) Joins(query string, args ...interface{}) *OrmAdminRolesCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminRolesCopy) WhereId(val uint32) *OrmAdminRolesCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) InsertGetId(row *AdminRolesCopy) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAdminRolesCopy) WhereIdIn(val []uint32) *OrmAdminRolesCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereIdGt(val uint32) *OrmAdminRolesCopy {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereIdGte(val uint32) *OrmAdminRolesCopy {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereIdLt(val uint32) *OrmAdminRolesCopy {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereIdLte(val uint32) *OrmAdminRolesCopy {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereName(val string) *OrmAdminRolesCopy {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereNameIn(val []string) *OrmAdminRolesCopy {
	orm.db.Where("`name` IN ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereNameGt(val string) *OrmAdminRolesCopy {
	orm.db.Where("`name` > ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereNameGte(val string) *OrmAdminRolesCopy {
	orm.db.Where("`name` >= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereNameLt(val string) *OrmAdminRolesCopy {
	orm.db.Where("`name` < ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereNameLte(val string) *OrmAdminRolesCopy {
	orm.db.Where("`name` <= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereSlug(val string) *OrmAdminRolesCopy {
	orm.db.Where("`slug` = ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereRemark(val string) *OrmAdminRolesCopy {
	orm.db.Where("`remark` = ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereCreatedAt(val database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereCreatedAtLte(val database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereCreatedAtGte(val database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereUpdatedAt(val database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereUpdatedAtLte(val database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRolesCopy) WhereUpdatedAtGte(val database.Time) *OrmAdminRolesCopy {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminRolesCopyList []*AdminRolesCopy

func (l AdminRolesCopyList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminRolesCopyList) GetIdMap() map[uint32]*AdminRolesCopy {
	got := make(map[uint32]*AdminRolesCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
