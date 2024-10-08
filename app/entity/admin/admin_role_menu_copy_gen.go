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

type AdminRoleMenuCopy struct {
	RoleId    int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_menu_role_id_menu_id_index,class:BTREE;comment:'role_id'" json:"role_id"` // role_id
	MenuId    int32          `gorm:"column:menu_id;type:int(11);index:laravel_admin_role_menu_role_id_menu_id_index,class:BTREE;comment:'menu_id'" json:"menu_id"` // menu_id
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`                                                      // created_at
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`                                                      // updated_at
}

var (
	AdminRoleMenuCopyFieldRoleId    = "role_id"
	AdminRoleMenuCopyFieldMenuId    = "menu_id"
	AdminRoleMenuCopyFieldCreatedAt = "created_at"
	AdminRoleMenuCopyFieldUpdatedAt = "updated_at"
)

func (receiver *AdminRoleMenuCopy) TableName() string {
	return "admin_role_menu_copy"
}

type OrmAdminRoleMenuCopy struct {
	db *gorm.DB
}

func (orm *OrmAdminRoleMenuCopy) TableName() string {
	return "admin_role_menu_copy"
}

func NewOrmAdminRoleMenuCopy(txs ...*gorm.DB) *OrmAdminRoleMenuCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminRoleMenuCopy{})
	} else {
		tx = txs[0].Model(&AdminRoleMenuCopy{})
	}
	return &OrmAdminRoleMenuCopy{db: tx}
}

func (orm *OrmAdminRoleMenuCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminRoleMenuCopy) GetTableInfo() interface{} {
	return &AdminRoleMenuCopy{}
}

// Create insert the value into database
func (orm *OrmAdminRoleMenuCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminRoleMenuCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminRoleMenuCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminRoleMenuCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminRoleMenuCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminRoleMenuCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminRoleMenuCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminRoleMenuCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminRoleMenuCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminRoleMenuCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminRoleMenuCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminRoleMenuCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminRoleMenuCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminRoleMenuCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminRoleMenuCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminRoleMenuCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminRoleMenuCopy) Unscoped() *OrmAdminRoleMenuCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminRoleMenuCopy) Insert(row *AdminRoleMenuCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminRoleMenuCopy) Inserts(rows []*AdminRoleMenuCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminRoleMenuCopy) Order(value interface{}) *OrmAdminRoleMenuCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) Group(name string) *OrmAdminRoleMenuCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) Limit(limit int) *OrmAdminRoleMenuCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) Offset(offset int) *OrmAdminRoleMenuCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminRoleMenuCopy) Get() AdminRoleMenuCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminRoleMenuCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminRoleMenuCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminRoleMenuCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminRoleMenuCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_role_menu_copy")
}

func (orm *OrmAdminRoleMenuCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminRoleMenuCopy) First(conds ...interface{}) (*AdminRoleMenuCopy, bool) {
	dest := &AdminRoleMenuCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminRoleMenuCopy) Take(conds ...interface{}) (*AdminRoleMenuCopy, int64) {
	dest := &AdminRoleMenuCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminRoleMenuCopy) Last(conds ...interface{}) (*AdminRoleMenuCopy, int64) {
	dest := &AdminRoleMenuCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminRoleMenuCopy) Find(conds ...interface{}) (AdminRoleMenuCopyList, int64) {
	list := make([]*AdminRoleMenuCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminRoleMenuCopy) Paginate(page int, limit int) (AdminRoleMenuCopyList, int64) {
	var total int64
	list := make([]*AdminRoleMenuCopy, 0)
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
func (orm *OrmAdminRoleMenuCopy) SimplePaginate(page int, limit int) AdminRoleMenuCopyList {
	list := make([]*AdminRoleMenuCopy, 0)
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
func (orm *OrmAdminRoleMenuCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminRoleMenuCopy) FirstOrInit(dest *AdminRoleMenuCopy, conds ...interface{}) (*AdminRoleMenuCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminRoleMenuCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleMenuCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleMenuCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminRoleMenuCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminRoleMenuCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminRoleMenuCopy) Where(query interface{}, args ...interface{}) *OrmAdminRoleMenuCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) Select(query interface{}, args ...interface{}) *OrmAdminRoleMenuCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminRoleMenuCopy) And(fuc func(orm *OrmAdminRoleMenuCopy)) *OrmAdminRoleMenuCopy {
	ormAnd := NewOrmAdminRoleMenuCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) Or(fuc func(orm *OrmAdminRoleMenuCopy)) *OrmAdminRoleMenuCopy {
	ormOr := NewOrmAdminRoleMenuCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminRoleMenuCopy) Preload(query string, args ...interface{}) *OrmAdminRoleMenuCopy {
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
func (orm *OrmAdminRoleMenuCopy) Joins(query string, args ...interface{}) *OrmAdminRoleMenuCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminRoleMenuCopy) WhereRoleId(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereRoleIdIn(val []int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereRoleIdGt(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereRoleIdGte(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereRoleIdLt(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereRoleIdLte(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereMenuId(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`menu_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereMenuIdIn(val []int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`menu_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereMenuIdNe(val int32) *OrmAdminRoleMenuCopy {
	orm.db.Where("`menu_id` <> ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereCreatedAt(val database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereCreatedAtLte(val database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereCreatedAtGte(val database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereUpdatedAt(val database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereUpdatedAtLte(val database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleMenuCopy) WhereUpdatedAtGte(val database.Time) *OrmAdminRoleMenuCopy {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminRoleMenuCopyList []*AdminRoleMenuCopy

func (l AdminRoleMenuCopyList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l AdminRoleMenuCopyList) GetRoleIdMap() map[int32]*AdminRoleMenuCopy {
	got := make(map[int32]*AdminRoleMenuCopy)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l AdminRoleMenuCopyList) GetMenuIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.MenuId)
	}
	return got
}
func (l AdminRoleMenuCopyList) GetMenuIdMap() map[int32]*AdminRoleMenuCopy {
	got := make(map[int32]*AdminRoleMenuCopy)
	for _, val := range l {
		got[val.MenuId] = val
	}
	return got
}
