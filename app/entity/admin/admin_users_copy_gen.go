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

type AdminUsersCopy struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey;comment:'id'" json:"id"`                                             // id
	Username  string         `gorm:"column:username;type:varchar(190);index:laravel_admin_users_username_unique,class:BTREE,unique;comment:'账户'" json:"username"` // 账户
	Password  string         `gorm:"column:password;type:varchar(60);comment:'密码'" json:"password"`                                                               // 密码
	Name      string         `gorm:"column:name;type:varchar(255);comment:'显示名称'" json:"name"`                                                                    // 显示名称
	Avatar    *string        `gorm:"column:avatar;type:varchar(255);comment:'头像'" json:"avatar"`                                                                  // 头像
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:'软删除'" json:"deleted_at"`                                                            // 软删除
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`                                                     // created_at
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`                                                     // updated_at
}

var (
	AdminUsersCopyFieldId        = "id"
	AdminUsersCopyFieldUsername  = "username"
	AdminUsersCopyFieldPassword  = "password"
	AdminUsersCopyFieldName      = "name"
	AdminUsersCopyFieldAvatar    = "avatar"
	AdminUsersCopyFieldDeletedAt = "deleted_at"
	AdminUsersCopyFieldCreatedAt = "created_at"
	AdminUsersCopyFieldUpdatedAt = "updated_at"
)

func (receiver *AdminUsersCopy) TableName() string {
	return "admin_users_copy"
}

type OrmAdminUsersCopy struct {
	db *gorm.DB
}

func (orm *OrmAdminUsersCopy) TableName() string {
	return "admin_users_copy"
}

func NewOrmAdminUsersCopy(txs ...*gorm.DB) *OrmAdminUsersCopy {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminUsersCopy{})
	} else {
		tx = txs[0].Model(&AdminUsersCopy{})
	}
	return &OrmAdminUsersCopy{db: tx}
}

func (orm *OrmAdminUsersCopy) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminUsersCopy) GetTableInfo() interface{} {
	return &AdminUsersCopy{}
}

// Create insert the value into database
func (orm *OrmAdminUsersCopy) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminUsersCopy) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminUsersCopy) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminUsersCopy) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminUsersCopy) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminUsersCopy) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminUsersCopy) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminUsersCopy) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminUsersCopy) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminUsersCopy) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminUsersCopy) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminUsersCopy) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminUsersCopy) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminUsersCopy) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminUsersCopy) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminUsersCopy) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminUsersCopy) Unscoped() *OrmAdminUsersCopy {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminUsersCopy) Insert(row *AdminUsersCopy) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminUsersCopy) Inserts(rows []*AdminUsersCopy) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminUsersCopy) Order(value interface{}) *OrmAdminUsersCopy {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminUsersCopy) Group(name string) *OrmAdminUsersCopy {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminUsersCopy) Limit(limit int) *OrmAdminUsersCopy {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminUsersCopy) Offset(offset int) *OrmAdminUsersCopy {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminUsersCopy) Get() AdminUsersCopyList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminUsersCopy) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminUsersCopy) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminUsersCopy{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminUsersCopy) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_users_copy")
}

func (orm *OrmAdminUsersCopy) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminUsersCopy) First(conds ...interface{}) (*AdminUsersCopy, bool) {
	dest := &AdminUsersCopy{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminUsersCopy) Take(conds ...interface{}) (*AdminUsersCopy, int64) {
	dest := &AdminUsersCopy{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminUsersCopy) Last(conds ...interface{}) (*AdminUsersCopy, int64) {
	dest := &AdminUsersCopy{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminUsersCopy) Find(conds ...interface{}) (AdminUsersCopyList, int64) {
	list := make([]*AdminUsersCopy, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminUsersCopy) Paginate(page int, limit int) (AdminUsersCopyList, int64) {
	var total int64
	list := make([]*AdminUsersCopy, 0)
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
func (orm *OrmAdminUsersCopy) SimplePaginate(page int, limit int) AdminUsersCopyList {
	list := make([]*AdminUsersCopy, 0)
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
func (orm *OrmAdminUsersCopy) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminUsersCopy) FirstOrInit(dest *AdminUsersCopy, conds ...interface{}) (*AdminUsersCopy, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminUsersCopy) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminUsersCopy) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminUsersCopy) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminUsersCopy) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminUsersCopy) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminUsersCopy) Where(query interface{}, args ...interface{}) *OrmAdminUsersCopy {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminUsersCopy) Select(query interface{}, args ...interface{}) *OrmAdminUsersCopy {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminUsersCopy) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminUsersCopy) And(fuc func(orm *OrmAdminUsersCopy)) *OrmAdminUsersCopy {
	ormAnd := NewOrmAdminUsersCopy()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminUsersCopy) Or(fuc func(orm *OrmAdminUsersCopy)) *OrmAdminUsersCopy {
	ormOr := NewOrmAdminUsersCopy()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminUsersCopy) Preload(query string, args ...interface{}) *OrmAdminUsersCopy {
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
func (orm *OrmAdminUsersCopy) Joins(query string, args ...interface{}) *OrmAdminUsersCopy {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminUsersCopy) WhereId(val uint32) *OrmAdminUsersCopy {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) InsertGetId(row *AdminUsersCopy) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAdminUsersCopy) WhereIdIn(val []uint32) *OrmAdminUsersCopy {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereIdGt(val uint32) *OrmAdminUsersCopy {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereIdGte(val uint32) *OrmAdminUsersCopy {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereIdLt(val uint32) *OrmAdminUsersCopy {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereIdLte(val uint32) *OrmAdminUsersCopy {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUsername(val string) *OrmAdminUsersCopy {
	orm.db.Where("`username` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUsernameIn(val []string) *OrmAdminUsersCopy {
	orm.db.Where("`username` IN ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUsernameGt(val string) *OrmAdminUsersCopy {
	orm.db.Where("`username` > ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUsernameGte(val string) *OrmAdminUsersCopy {
	orm.db.Where("`username` >= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUsernameLt(val string) *OrmAdminUsersCopy {
	orm.db.Where("`username` < ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUsernameLte(val string) *OrmAdminUsersCopy {
	orm.db.Where("`username` <= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WherePassword(val string) *OrmAdminUsersCopy {
	orm.db.Where("`password` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereName(val string) *OrmAdminUsersCopy {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereAvatar(val string) *OrmAdminUsersCopy {
	orm.db.Where("`avatar` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereDeletedAt(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`deleted_at` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereDeletedAtBetween(begin database.Time, end database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`deleted_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereDeletedAtLte(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`deleted_at` <= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereDeletedAtGte(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`deleted_at` >= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereCreatedAt(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereCreatedAtLte(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereCreatedAtGte(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUpdatedAt(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUpdatedAtLte(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminUsersCopy) WhereUpdatedAtGte(val database.Time) *OrmAdminUsersCopy {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminUsersCopyList []*AdminUsersCopy

func (l AdminUsersCopyList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminUsersCopyList) GetIdMap() map[uint32]*AdminUsersCopy {
	got := make(map[uint32]*AdminUsersCopy)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
