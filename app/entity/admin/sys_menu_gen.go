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

type SysMenu struct {
	Id        int64         `gorm:"column:id;autoIncrement;type:bigint(20);primaryKey" json:"id"`                      //
	ParentId  int64         `gorm:"column:parent_id;type:bigint(20);default:0;comment:'父级ID'" json:"parent_id"`        // 父级ID
	Name      string        `gorm:"column:name;type:varchar(32);comment:'菜单名称'" json:"name"`                           // 菜单名称
	Icon      string        `gorm:"column:icon;type:varchar(32);default:;comment:'图标'" json:"icon"`                    // 图标
	Path      string        `gorm:"column:path;type:varchar(64);comment:'菜单路径L'" json:"path"`                          // 菜单路径L
	Redirect  string        `gorm:"column:redirect;type:varchar(32);default:;comment:'重定向地址'" json:"redirect"`         // 重定向地址
	SchemaApi string        `gorm:"column:schema_api;type:varchar(64);comment:'schema json url 地址'" json:"schema_api"` // schema json url 地址
	SortNo    int32         `gorm:"column:sort_no;type:int(11);default:100;comment:'排序序号'" json:"sort_no"`             // 排序序号
	Visible   int32         `gorm:"column:visible;type:int(11);default:1;comment:'是否显示'" json:"visible"`               // 是否显示
	IsSys     int32         `gorm:"column:is_sys;type:int(11);default:0;comment:'是否是系统菜单'" json:"is_sys"`              // 是否是系统菜单
	Status    string        `gorm:"column:status;type:varchar(16);comment:'状态'" json:"status"`                         // 状态
	Remark    string        `gorm:"column:remark;type:varchar(512);default:;comment:'备注'" json:"remark"`               // 备注
	Utime     database.Time `gorm:"column:utime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'更新时间'" json:"utime"` // 更新时间
	Ctime     database.Time `gorm:"column:ctime;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"ctime"` // 创建时间
}

var (
	SysMenuFieldId        = "id"
	SysMenuFieldParentId  = "parent_id"
	SysMenuFieldName      = "name"
	SysMenuFieldIcon      = "icon"
	SysMenuFieldPath      = "path"
	SysMenuFieldRedirect  = "redirect"
	SysMenuFieldSchemaApi = "schema_api"
	SysMenuFieldSortNo    = "sort_no"
	SysMenuFieldVisible   = "visible"
	SysMenuFieldIsSys     = "is_sys"
	SysMenuFieldStatus    = "status"
	SysMenuFieldRemark    = "remark"
	SysMenuFieldUtime     = "utime"
	SysMenuFieldCtime     = "ctime"
)

func (receiver *SysMenu) TableName() string {
	return "sys_menu"
}

type OrmSysMenu struct {
	db *gorm.DB
}

func (orm *OrmSysMenu) TableName() string {
	return "sys_menu"
}

func NewOrmSysMenu(txs ...*gorm.DB) *OrmSysMenu {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&SysMenu{})
	} else {
		tx = txs[0].Model(&SysMenu{})
	}
	return &OrmSysMenu{db: tx}
}

func (orm *OrmSysMenu) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmSysMenu) GetTableInfo() interface{} {
	return &SysMenu{}
}

// Create insert the value into database
func (orm *OrmSysMenu) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmSysMenu) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmSysMenu) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmSysMenu) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmSysMenu) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmSysMenu) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmSysMenu) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmSysMenu) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmSysMenu) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmSysMenu) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmSysMenu) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmSysMenu) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmSysMenu) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmSysMenu) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmSysMenu) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmSysMenu) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmSysMenu) Unscoped() *OrmSysMenu {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmSysMenu) Insert(row *SysMenu) error {
	return orm.db.Create(row).Error
}

func (orm *OrmSysMenu) Inserts(rows []*SysMenu) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmSysMenu) Order(value interface{}) *OrmSysMenu {
	orm.db.Order(value)
	return orm
}

func (orm *OrmSysMenu) Group(name string) *OrmSysMenu {
	orm.db.Group(name)
	return orm
}

func (orm *OrmSysMenu) Limit(limit int) *OrmSysMenu {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmSysMenu) Offset(offset int) *OrmSysMenu {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmSysMenu) Get() SysMenuList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmSysMenu) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmSysMenu) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&SysMenu{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmSysMenu) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM sys_menu")
}

func (orm *OrmSysMenu) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmSysMenu) First(conds ...interface{}) (*SysMenu, bool) {
	dest := &SysMenu{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmSysMenu) Take(conds ...interface{}) (*SysMenu, int64) {
	dest := &SysMenu{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmSysMenu) Last(conds ...interface{}) (*SysMenu, int64) {
	dest := &SysMenu{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmSysMenu) Find(conds ...interface{}) (SysMenuList, int64) {
	list := make([]*SysMenu, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmSysMenu) Paginate(page int, limit int) (SysMenuList, int64) {
	var total int64
	list := make([]*SysMenu, 0)
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
func (orm *OrmSysMenu) SimplePaginate(page int, limit int) SysMenuList {
	list := make([]*SysMenu, 0)
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
func (orm *OrmSysMenu) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmSysMenu) FirstOrInit(dest *SysMenu, conds ...interface{}) (*SysMenu, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmSysMenu) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysMenu) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmSysMenu) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmSysMenu) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmSysMenu) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmSysMenu) Where(query interface{}, args ...interface{}) *OrmSysMenu {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmSysMenu) Select(query interface{}, args ...interface{}) *OrmSysMenu {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmSysMenu) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmSysMenu) And(fuc func(orm *OrmSysMenu)) *OrmSysMenu {
	ormAnd := NewOrmSysMenu()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmSysMenu) Or(fuc func(orm *OrmSysMenu)) *OrmSysMenu {
	ormOr := NewOrmSysMenu()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmSysMenu) Preload(query string, args ...interface{}) *OrmSysMenu {
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
func (orm *OrmSysMenu) Joins(query string, args ...interface{}) *OrmSysMenu {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmSysMenu) WhereId(val int64) *OrmSysMenu {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmSysMenu) InsertGetId(row *SysMenu) int64 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmSysMenu) WhereIdIn(val []int64) *OrmSysMenu {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereIdGt(val int64) *OrmSysMenu {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereIdGte(val int64) *OrmSysMenu {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereIdLt(val int64) *OrmSysMenu {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereIdLte(val int64) *OrmSysMenu {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereParentId(val int64) *OrmSysMenu {
	orm.db.Where("`parent_id` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereParentIdIn(val []int64) *OrmSysMenu {
	orm.db.Where("`parent_id` IN ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereParentIdNe(val int64) *OrmSysMenu {
	orm.db.Where("`parent_id` <> ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereName(val string) *OrmSysMenu {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereIcon(val string) *OrmSysMenu {
	orm.db.Where("`icon` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WherePath(val string) *OrmSysMenu {
	orm.db.Where("`path` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereRedirect(val string) *OrmSysMenu {
	orm.db.Where("`redirect` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereSchemaApi(val string) *OrmSysMenu {
	orm.db.Where("`schema_api` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereSortNo(val int32) *OrmSysMenu {
	orm.db.Where("`sort_no` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereVisible(val int32) *OrmSysMenu {
	orm.db.Where("`visible` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereIsSys(val int32) *OrmSysMenu {
	orm.db.Where("`is_sys` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereStatus(val string) *OrmSysMenu {
	orm.db.Where("`status` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereStatusIn(val []string) *OrmSysMenu {
	orm.db.Where("`status` IN ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereStatusNe(val string) *OrmSysMenu {
	orm.db.Where("`status` <> ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereRemark(val string) *OrmSysMenu {
	orm.db.Where("`remark` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereUtime(val database.Time) *OrmSysMenu {
	orm.db.Where("`utime` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereUtimeBetween(begin database.Time, end database.Time) *OrmSysMenu {
	orm.db.Where("`utime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysMenu) WhereUtimeLte(val database.Time) *OrmSysMenu {
	orm.db.Where("`utime` <= ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereUtimeGte(val database.Time) *OrmSysMenu {
	orm.db.Where("`utime` >= ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereCtime(val database.Time) *OrmSysMenu {
	orm.db.Where("`ctime` = ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereCtimeBetween(begin database.Time, end database.Time) *OrmSysMenu {
	orm.db.Where("`ctime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmSysMenu) WhereCtimeLte(val database.Time) *OrmSysMenu {
	orm.db.Where("`ctime` <= ?", val)
	return orm
}
func (orm *OrmSysMenu) WhereCtimeGte(val database.Time) *OrmSysMenu {
	orm.db.Where("`ctime` >= ?", val)
	return orm
}

type SysMenuList []*SysMenu

func (l SysMenuList) GetIdList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l SysMenuList) GetIdMap() map[int64]*SysMenu {
	got := make(map[int64]*SysMenu)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l SysMenuList) GetParentIdList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.ParentId)
	}
	return got
}
func (l SysMenuList) GetParentIdMap() map[int64]*SysMenu {
	got := make(map[int64]*SysMenu)
	for _, val := range l {
		got[val.ParentId] = val
	}
	return got
}
