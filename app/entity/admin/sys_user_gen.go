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

type SysUser struct {
	Id       int64         `gorm:"column:id;autoIncrement;type:bigint(20);primaryKey" json:"id"`                                                           //
	Username string        `gorm:"column:username;type:varchar(32);index:uk_u,class:BTREE,unique;index:idx_u_s,class:BTREE;comment:'用户名'" json:"username"` // 用户名
	Pwd      string        `gorm:"column:pwd;type:varchar(64);comment:'密码'" json:"pwd"`                                                                    // 密码
	Salt     string        `gorm:"column:salt;type:varchar(64);comment:'密码盐'" json:"salt"`                                                                 // 密码盐
	Email    string        `gorm:"column:email;type:varchar(64);default:;comment:'邮箱'" json:"email"`                                                       // 邮箱
	Nickname string        `gorm:"column:nickname;type:varchar(32);comment:'昵称'" json:"nickname"`                                                          // 昵称
	Avatar   string        `gorm:"column:avatar;type:varchar(256);default:;comment:'头像'" json:"avatar"`                                                    // 头像
	Phone    string        `gorm:"column:phone;type:varchar(32);default:;comment:'手机号'" json:"phone"`                                                      // 手机号
	Status   string        `gorm:"column:status;type:varchar(32);index:idx_u_s,class:BTREE;default:normal;comment:'用户状态'" json:"status"`                   // 用户状态
	Remark   string        `gorm:"column:remark;type:varchar(512);default:;comment:'备注'" json:"remark"`                                                    // 备注
	Utime    database.Time `gorm:"column:utime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'更新时间'" json:"utime"`                                      // 更新时间
	Ctime    database.Time `gorm:"column:ctime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"ctime"`                                      // 创建时间
}

var (
	SysUserFieldId       = "id"
	SysUserFieldUsername = "username"
	SysUserFieldPwd      = "pwd"
	SysUserFieldSalt     = "salt"
	SysUserFieldEmail    = "email"
	SysUserFieldNickname = "nickname"
	SysUserFieldAvatar   = "avatar"
	SysUserFieldPhone    = "phone"
	SysUserFieldStatus   = "status"
	SysUserFieldRemark   = "remark"
	SysUserFieldUtime    = "utime"
	SysUserFieldCtime    = "ctime"
)

func (receiver *SysUser) TableName() string {
	return "sys_user"
}

type OrmSysUser struct {
	db *gorm.DB
}

func (orm *OrmSysUser) TableName() string {
	return "sys_user"
}

func NewOrmSysUser(txs ...*gorm.DB) *OrmSysUser {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&SysUser{})
	} else {
		tx = txs[0].Model(&SysUser{})
	}
	return &OrmSysUser{db: tx}
}

func (orm *OrmSysUser) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmSysUser) GetTableInfo() interface{} {
	return &SysUser{}
}

// Create insert the value into database
func (orm *OrmSysUser) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmSysUser) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmSysUser) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmSysUser) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmSysUser) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmSysUser) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmSysUser) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmSysUser) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmSysUser) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmSysUser) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmSysUser) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmSysUser) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmSysUser) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmSysUser) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmSysUser) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmSysUser) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmSysUser) Unscoped() *OrmSysUser {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmSysUser) Insert(row *SysUser) error {
	return orm.db.Create(row).Error
}

func (orm *OrmSysUser) Inserts(rows []*SysUser) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmSysUser) Order(value interface{}) *OrmSysUser {
	orm.db.Order(value)
	return orm
}

func (orm *OrmSysUser) Group(name string) *OrmSysUser {
	orm.db.Group(name)
	return orm
}

func (orm *OrmSysUser) Limit(limit int) *OrmSysUser {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmSysUser) Offset(offset int) *OrmSysUser {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmSysUser) Get() SysUserList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmSysUser) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmSysUser) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&SysUser{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmSysUser) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM sys_user")
}

func (orm *OrmSysUser) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmSysUser) First(conds ...interface{}) (*SysUser, bool) {
	dest := &SysUser{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmSysUser) Take(conds ...interface{}) (*SysUser, int64) {
	dest := &SysUser{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmSysUser) Last(conds ...interface{}) (*SysUser, int64) {
	dest := &SysUser{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmSysUser) Find(conds ...interface{}) (SysUserList, int64) {
	list := make([]*SysUser, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmSysUser) Paginate(page int, limit int) (SysUserList, int64) {
	var total int64
	list := make([]*SysUser, 0)
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
func (orm *OrmSysUser) SimplePaginate(page int, limit int) SysUserList {
	list := make([]*SysUser, 0)
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
func (orm *OrmSysUser) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmSysUser) FirstOrInit(dest *SysUser, conds ...interface{}) (*SysUser, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmSysUser) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysUser) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysUser) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmSysUser) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmSysUser) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmSysUser) Where(query interface{}, args ...interface{}) *OrmSysUser {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmSysUser) Select(query interface{}, args ...interface{}) *OrmSysUser {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmSysUser) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmSysUser) And(fuc func(orm *OrmSysUser)) *OrmSysUser {
	ormAnd := NewOrmSysUser()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmSysUser) Or(fuc func(orm *OrmSysUser)) *OrmSysUser {
	ormOr := NewOrmSysUser()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmSysUser) Preload(query string, args ...interface{}) *OrmSysUser {
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
func (orm *OrmSysUser) Joins(query string, args ...interface{}) *OrmSysUser {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmSysUser) WhereId(val int64) *OrmSysUser {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmSysUser) InsertGetId(row *SysUser) int64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmSysUser) WhereIdIn(val []int64) *OrmSysUser {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmSysUser) WhereIdGt(val int64) *OrmSysUser {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmSysUser) WhereIdGte(val int64) *OrmSysUser {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmSysUser) WhereIdLt(val int64) *OrmSysUser {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmSysUser) WhereIdLte(val int64) *OrmSysUser {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUsername(val string) *OrmSysUser {
	orm.db.Where("`username` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUsernameIn(val []string) *OrmSysUser {
	orm.db.Where("`username` IN ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUsernameGt(val string) *OrmSysUser {
	orm.db.Where("`username` > ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUsernameGte(val string) *OrmSysUser {
	orm.db.Where("`username` >= ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUsernameLt(val string) *OrmSysUser {
	orm.db.Where("`username` < ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUsernameLte(val string) *OrmSysUser {
	orm.db.Where("`username` <= ?", val)
	return orm
}
func (orm *OrmSysUser) WherePwd(val string) *OrmSysUser {
	orm.db.Where("`pwd` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereSalt(val string) *OrmSysUser {
	orm.db.Where("`salt` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereEmail(val string) *OrmSysUser {
	orm.db.Where("`email` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereNickname(val string) *OrmSysUser {
	orm.db.Where("`nickname` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereAvatar(val string) *OrmSysUser {
	orm.db.Where("`avatar` = ?", val)
	return orm
}
func (orm *OrmSysUser) WherePhone(val string) *OrmSysUser {
	orm.db.Where("`phone` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereStatus(val string) *OrmSysUser {
	orm.db.Where("`status` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereStatusIn(val []string) *OrmSysUser {
	orm.db.Where("`status` IN ?", val)
	return orm
}
func (orm *OrmSysUser) WhereStatusNe(val string) *OrmSysUser {
	orm.db.Where("`status` <> ?", val)
	return orm
}
func (orm *OrmSysUser) WhereRemark(val string) *OrmSysUser {
	orm.db.Where("`remark` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUtime(val database.Time) *OrmSysUser {
	orm.db.Where("`utime` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUtimeBetween(begin database.Time, end database.Time) *OrmSysUser {
	orm.db.Where("`utime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysUser) WhereUtimeLte(val database.Time) *OrmSysUser {
	orm.db.Where("`utime` <= ?", val)
	return orm
}
func (orm *OrmSysUser) WhereUtimeGte(val database.Time) *OrmSysUser {
	orm.db.Where("`utime` >= ?", val)
	return orm
}
func (orm *OrmSysUser) WhereCtime(val database.Time) *OrmSysUser {
	orm.db.Where("`ctime` = ?", val)
	return orm
}
func (orm *OrmSysUser) WhereCtimeBetween(begin database.Time, end database.Time) *OrmSysUser {
	orm.db.Where("`ctime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysUser) WhereCtimeLte(val database.Time) *OrmSysUser {
	orm.db.Where("`ctime` <= ?", val)
	return orm
}
func (orm *OrmSysUser) WhereCtimeGte(val database.Time) *OrmSysUser {
	orm.db.Where("`ctime` >= ?", val)
	return orm
}

type SysUserList []*SysUser

func (l SysUserList) GetIdList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l SysUserList) GetIdMap() map[int64]*SysUser {
	got := make(map[int64]*SysUser)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
