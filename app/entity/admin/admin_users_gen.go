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

type AdminUsers struct {
	Id         uint32         `gorm:"column:id;autoIncrement;type:int(10) unsigned;primaryKey;comment:'id'" json:"id"`                                             // id
	Username   string         `gorm:"column:username;type:varchar(190);index:laravel_admin_users_username_unique,class:BTREE,unique;comment:'账户'" json:"username"` // 账户
	Password   string         `gorm:"column:password;type:varchar(60);comment:'密码'" json:"password"`                                                               // 密码
	Name       string         `gorm:"column:name;type:varchar(255);comment:'显示名称'" json:"name"`                                                                    // 显示名称
	Avatar     *string        `gorm:"column:avatar;type:varchar(255);comment:'头像'" json:"avatar"`                                                                  // 头像
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:'软删除'" json:"deleted_at"`                                                            // 软删除
	CreatedAt  *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`                                                     // created_at
	UpdatedAt  *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`                                                     // updated_at
	AdminRoles []*AdminRoles  `gorm:"many2many:admin_role_users;joinForeignKey:user_id;joinReferences:role_id;"`
}

var (
	AdminUsersFieldId        = "id"
	AdminUsersFieldUsername  = "username"
	AdminUsersFieldPassword  = "password"
	AdminUsersFieldName      = "name"
	AdminUsersFieldAvatar    = "avatar"
	AdminUsersFieldDeletedAt = "deleted_at"
	AdminUsersFieldCreatedAt = "created_at"
	AdminUsersFieldUpdatedAt = "updated_at"
)

func (receiver *AdminUsers) TableName() string {
	return "admin_users"
}

type OrmAdminUsers struct {
	db *gorm.DB
}

func (orm *OrmAdminUsers) TableName() string {
	return "admin_users"
}

func NewOrmAdminUsers(txs ...*gorm.DB) *OrmAdminUsers {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminUsers{})
	} else {
		tx = txs[0].Model(&AdminUsers{})
	}
	return &OrmAdminUsers{db: tx}
}

func (orm *OrmAdminUsers) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminUsers) GetTableInfo() interface{} {
	return &AdminUsers{}
}

// Create insert the value into database
func (orm *OrmAdminUsers) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminUsers) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminUsers) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminUsers) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminUsers) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminUsers) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminUsers) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminUsers) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminUsers) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminUsers) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminUsers) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminUsers) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminUsers) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminUsers) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminUsers) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminUsers) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmAdminUsers) Unscoped() *OrmAdminUsers {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminUsers) Insert(row *AdminUsers) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminUsers) Inserts(rows []*AdminUsers) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminUsers) Order(value interface{}) *OrmAdminUsers {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminUsers) Group(name string) *OrmAdminUsers {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminUsers) Limit(limit int) *OrmAdminUsers {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminUsers) Offset(offset int) *OrmAdminUsers {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminUsers) Get() AdminUsersList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminUsers) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminUsers) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminUsers{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminUsers) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_users")
}

func (orm *OrmAdminUsers) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminUsers) First(conds ...interface{}) (*AdminUsers, bool) {
	dest := &AdminUsers{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminUsers) Take(conds ...interface{}) (*AdminUsers, int64) {
	dest := &AdminUsers{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminUsers) Last(conds ...interface{}) (*AdminUsers, int64) {
	dest := &AdminUsers{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminUsers) Find(conds ...interface{}) (AdminUsersList, int64) {
	list := make([]*AdminUsers, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminUsers) Paginate(page int, limit int) (AdminUsersList, int64) {
	var total int64
	list := make([]*AdminUsers, 0)
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
func (orm *OrmAdminUsers) SimplePaginate(page int, limit int) AdminUsersList {
	list := make([]*AdminUsers, 0)
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
func (orm *OrmAdminUsers) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminUsers) FirstOrInit(dest *AdminUsers, conds ...interface{}) (*AdminUsers, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminUsers) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminUsers) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminUsers) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminUsers) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminUsers) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminUsers) Where(query interface{}, args ...interface{}) *OrmAdminUsers {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminUsers) Select(query interface{}, args ...interface{}) *OrmAdminUsers {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminUsers) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminUsers) And(fuc func(orm *OrmAdminUsers)) *OrmAdminUsers {
	ormAnd := NewOrmAdminUsers()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminUsers) Or(fuc func(orm *OrmAdminUsers)) *OrmAdminUsers {
	ormOr := NewOrmAdminUsers()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminUsers) Preload(query string, args ...interface{}) *OrmAdminUsers {
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
func (orm *OrmAdminUsers) Joins(query string, args ...interface{}) *OrmAdminUsers {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminUsers) WhereId(val uint32) *OrmAdminUsers {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) InsertGetId(row *AdminUsers) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAdminUsers) WhereIdIn(val []uint32) *OrmAdminUsers {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereIdGt(val uint32) *OrmAdminUsers {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereIdGte(val uint32) *OrmAdminUsers {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereIdLt(val uint32) *OrmAdminUsers {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereIdLte(val uint32) *OrmAdminUsers {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUsername(val string) *OrmAdminUsers {
	orm.db.Where("`username` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUsernameIn(val []string) *OrmAdminUsers {
	orm.db.Where("`username` IN ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUsernameGt(val string) *OrmAdminUsers {
	orm.db.Where("`username` > ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUsernameGte(val string) *OrmAdminUsers {
	orm.db.Where("`username` >= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUsernameLt(val string) *OrmAdminUsers {
	orm.db.Where("`username` < ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUsernameLte(val string) *OrmAdminUsers {
	orm.db.Where("`username` <= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WherePassword(val string) *OrmAdminUsers {
	orm.db.Where("`password` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereName(val string) *OrmAdminUsers {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereAvatar(val string) *OrmAdminUsers {
	orm.db.Where("`avatar` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereDeletedAt(val database.Time) *OrmAdminUsers {
	orm.db.Where("`deleted_at` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereDeletedAtBetween(begin database.Time, end database.Time) *OrmAdminUsers {
	orm.db.Where("`deleted_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminUsers) WhereDeletedAtLte(val database.Time) *OrmAdminUsers {
	orm.db.Where("`deleted_at` <= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereDeletedAtGte(val database.Time) *OrmAdminUsers {
	orm.db.Where("`deleted_at` >= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereCreatedAt(val database.Time) *OrmAdminUsers {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminUsers {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminUsers) WhereCreatedAtLte(val database.Time) *OrmAdminUsers {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereCreatedAtGte(val database.Time) *OrmAdminUsers {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUpdatedAt(val database.Time) *OrmAdminUsers {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminUsers {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminUsers) WhereUpdatedAtLte(val database.Time) *OrmAdminUsers {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminUsers) WhereUpdatedAtGte(val database.Time) *OrmAdminUsers {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminUsersList []*AdminUsers

func (l AdminUsersList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminUsersList) GetIdMap() map[uint32]*AdminUsers {
	got := make(map[uint32]*AdminUsers)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
