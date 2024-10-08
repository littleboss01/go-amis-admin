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

type AdminRoleUsersCopy struct {
	RoleId    int32          `gorm:"column:role_id;type:int(11);index:laravel_admin_role_users_role_id_user_id_index,class:BTREE;comment:'role_id'" json:"role_id"` // role_id
	UserId    int32          `gorm:"column:user_id;type:int(11);index:laravel_admin_role_users_role_id_user_id_index,class:BTREE;comment:'user_id'" json:"user_id"` // user_id
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`                                                       // created_at
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`                                                       // updated_at
}

var (
	AdminRoleUsersCopyFieldRoleId    = "role_id"
	AdminRoleUsersCopyFieldUserId    = "user_id"
	AdminRoleUsersCopyFieldCreatedAt = "created_at"
	AdminRoleUsersCopyFieldUpdatedAt = "updated_at"
)

func (receiver *AdminRoleUsersCopy) TableName() string {
	return "admin_role_users_copy"
}

type OrmAdminRoleUsersCopy struct {
	db *gorm.DB
}

func (orm *OrmAdminRoleUsersCopy) TableName() string {
	return "admin_role_users_copy"
}

func NewOrmAdminRoleUsersCopy(txs ...*gorm.DB) *OrmAdminRoleUsersCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminRoleUsersCopy{})
	} else {
		tx = txs[0].Model(&AdminRoleUsersCopy{})
	}
	return &OrmAdminRoleUsersCopy{db: tx}
}

func (orm *OrmAdminRoleUsersCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminRoleUsersCopy) GetTableInfo() interface{} {
	return &AdminRoleUsersCopy{}
}

// Create insert the value into database
func (orm *OrmAdminRoleUsersCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminRoleUsersCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminRoleUsersCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminRoleUsersCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminRoleUsersCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminRoleUsersCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminRoleUsersCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminRoleUsersCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminRoleUsersCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminRoleUsersCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminRoleUsersCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminRoleUsersCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminRoleUsersCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminRoleUsersCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminRoleUsersCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminRoleUsersCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminRoleUsersCopy) Unscoped() *OrmAdminRoleUsersCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminRoleUsersCopy) Insert(row *AdminRoleUsersCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminRoleUsersCopy) Inserts(rows []*AdminRoleUsersCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminRoleUsersCopy) Order(value interface{}) *OrmAdminRoleUsersCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) Group(name string) *OrmAdminRoleUsersCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) Limit(limit int) *OrmAdminRoleUsersCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) Offset(offset int) *OrmAdminRoleUsersCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminRoleUsersCopy) Get() AdminRoleUsersCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminRoleUsersCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminRoleUsersCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminRoleUsersCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminRoleUsersCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_role_users_copy")
}

func (orm *OrmAdminRoleUsersCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminRoleUsersCopy) First(conds ...interface{}) (*AdminRoleUsersCopy, bool) {
	dest := &AdminRoleUsersCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminRoleUsersCopy) Take(conds ...interface{}) (*AdminRoleUsersCopy, int64) {
	dest := &AdminRoleUsersCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminRoleUsersCopy) Last(conds ...interface{}) (*AdminRoleUsersCopy, int64) {
	dest := &AdminRoleUsersCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminRoleUsersCopy) Find(conds ...interface{}) (AdminRoleUsersCopyList, int64) {
	list := make([]*AdminRoleUsersCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminRoleUsersCopy) Paginate(page int, limit int) (AdminRoleUsersCopyList, int64) {
	var total int64
	list := make([]*AdminRoleUsersCopy, 0)
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
func (orm *OrmAdminRoleUsersCopy) SimplePaginate(page int, limit int) AdminRoleUsersCopyList {
	list := make([]*AdminRoleUsersCopy, 0)
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
func (orm *OrmAdminRoleUsersCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminRoleUsersCopy) FirstOrInit(dest *AdminRoleUsersCopy, conds ...interface{}) (*AdminRoleUsersCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminRoleUsersCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleUsersCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminRoleUsersCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminRoleUsersCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminRoleUsersCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminRoleUsersCopy) Where(query interface{}, args ...interface{}) *OrmAdminRoleUsersCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) Select(query interface{}, args ...interface{}) *OrmAdminRoleUsersCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminRoleUsersCopy) And(fuc func(orm *OrmAdminRoleUsersCopy)) *OrmAdminRoleUsersCopy {
	ormAnd := NewOrmAdminRoleUsersCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) Or(fuc func(orm *OrmAdminRoleUsersCopy)) *OrmAdminRoleUsersCopy {
	ormOr := NewOrmAdminRoleUsersCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminRoleUsersCopy) Preload(query string, args ...interface{}) *OrmAdminRoleUsersCopy {
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
func (orm *OrmAdminRoleUsersCopy) Joins(query string, args ...interface{}) *OrmAdminRoleUsersCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminRoleUsersCopy) WhereRoleId(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`role_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereRoleIdIn(val []int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`role_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereRoleIdGt(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`role_id` > ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereRoleIdGte(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`role_id` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereRoleIdLt(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`role_id` < ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereRoleIdLte(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`role_id` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUserId(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`user_id` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUserIdIn(val []int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`user_id` IN ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUserIdNe(val int32) *OrmAdminRoleUsersCopy {
	orm.db.Where("`user_id` <> ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereCreatedAt(val database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereCreatedAtLte(val database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereCreatedAtGte(val database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUpdatedAt(val database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUpdatedAtLte(val database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminRoleUsersCopy) WhereUpdatedAtGte(val database.Time) *OrmAdminRoleUsersCopy {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminRoleUsersCopyList []*AdminRoleUsersCopy

func (l AdminRoleUsersCopyList) GetRoleIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.RoleId)
	}
	return got
}
func (l AdminRoleUsersCopyList) GetRoleIdMap() map[int32]*AdminRoleUsersCopy {
	got := make(map[int32]*AdminRoleUsersCopy)
	for _, val := range l {
		got[val.RoleId] = val
	}
	return got
}
func (l AdminRoleUsersCopyList) GetUserIdList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.UserId)
	}
	return got
}
func (l AdminRoleUsersCopyList) GetUserIdMap() map[int32]*AdminRoleUsersCopy {
	got := make(map[int32]*AdminRoleUsersCopy)
	for _, val := range l {
		got[val.UserId] = val
	}
	return got
}
