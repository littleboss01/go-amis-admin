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

type HmAccounts struct {
	Accountid                  int32         `gorm:"column:accountid;autoIncrement;type:int(11);primaryKey;index:accountid,class:BTREE,unique" json:"accountid"`                                           //
	Accountdomainid            int32         `gorm:"column:accountdomainid;type:int(11)" json:"accountdomainid"`                                                                                           //
	Accountadminlevel          int32         `gorm:"column:accountadminlevel;type:tinyint(4);default:0;comment:'管理级别'" json:"accountadminlevel"`                                                           // 管理级别
	Accountaddress             string        `gorm:"column:accountaddress;type:varchar(255);index:accountaddress,class:BTREE,unique;index:idx_hm_accounts,class:BTREE;comment:'地址'" json:"accountaddress"` // 地址
	Accountpassword            string        `gorm:"column:accountpassword;type:varchar(255);comment:'密码'" json:"accountpassword"`                                                                         // 密码
	Accountactive              int32         `gorm:"column:accountactive;type:tinyint(4);default:1;comment:'启用'" json:"accountactive"`                                                                     // 启用
	Accountisad                int32         `gorm:"column:accountisad;type:tinyint(4);comment:'是否是AD账户'" json:"accountisad"`                                                                              // 是否是AD账户
	Accountaddomain            string        `gorm:"column:accountaddomain;type:varchar(255)" json:"accountaddomain"`                                                                                      //
	Accountadusername          string        `gorm:"column:accountadusername;type:varchar(255)" json:"accountadusername"`                                                                                  //
	Accountmaxsize             int32         `gorm:"column:accountmaxsize;type:int(11);default:0;comment:'最大容量'" json:"accountmaxsize"`                                                                    // 最大容量
	Accountvacationmessageon   int32         `gorm:"column:accountvacationmessageon;type:tinyint(4);default:0" json:"accountvacationmessageon"`                                                            //
	Accountvacationmessage     string        `gorm:"column:accountvacationmessage;type:text" json:"accountvacationmessage"`                                                                                //
	Accountvacationsubject     string        `gorm:"column:accountvacationsubject;type:varchar(200)" json:"accountvacationsubject"`                                                                        //
	Accountpwencryption        int32         `gorm:"column:accountpwencryption;type:tinyint(4);default:0;comment:'加密方式'" json:"accountpwencryption"`                                                       // 加密方式
	Accountforwardenabled      int32         `gorm:"column:accountforwardenabled;type:tinyint(4);default:0" json:"accountforwardenabled"`                                                                  //
	Accountforwardaddress      string        `gorm:"column:accountforwardaddress;type:varchar(255)" json:"accountforwardaddress"`                                                                          //
	Accountforwardkeeporiginal int32         `gorm:"column:accountforwardkeeporiginal;type:tinyint(4)" json:"accountforwardkeeporiginal"`                                                                  //
	Accountenablesignature     int32         `gorm:"column:accountenablesignature;type:tinyint(4)" json:"accountenablesignature"`                                                                          //
	Accountsignatureplaintext  string        `gorm:"column:accountsignatureplaintext;type:text" json:"accountsignatureplaintext"`                                                                          //
	Accountsignaturehtml       string        `gorm:"column:accountsignaturehtml;type:text" json:"accountsignaturehtml"`                                                                                    //
	Accountlastlogontime       database.Time `gorm:"column:accountlastlogontime;type:datetime;default:2023-02-19 19:12:50;comment:'最后登录时间'" json:"accountlastlogontime"`                                   // 最后登录时间
	Accountvacationexpires     uint32        `gorm:"column:accountvacationexpires;type:tinyint(3) unsigned" json:"accountvacationexpires"`                                                                 //
	Accountvacationexpiredate  database.Time `gorm:"column:accountvacationexpiredate;type:datetime;default:2023-02-19 19:12:50" json:"accountvacationexpiredate"`                                          //
	Accountpersonfirstname     string        `gorm:"column:accountpersonfirstname;type:varchar(60)" json:"accountpersonfirstname"`                                                                         //
	Accountpersonlastname      string        `gorm:"column:accountpersonlastname;type:varchar(60)" json:"accountpersonlastname"`                                                                           //
	Creater                    *string       `gorm:"column:creater;type:varchar(255);default:;comment:'创建人'" json:"creater"`                                                                               // 创建人
}

var (
	HmAccountsFieldAccountid                  = "accountid"
	HmAccountsFieldAccountdomainid            = "accountdomainid"
	HmAccountsFieldAccountadminlevel          = "accountadminlevel"
	HmAccountsFieldAccountaddress             = "accountaddress"
	HmAccountsFieldAccountpassword            = "accountpassword"
	HmAccountsFieldAccountactive              = "accountactive"
	HmAccountsFieldAccountisad                = "accountisad"
	HmAccountsFieldAccountaddomain            = "accountaddomain"
	HmAccountsFieldAccountadusername          = "accountadusername"
	HmAccountsFieldAccountmaxsize             = "accountmaxsize"
	HmAccountsFieldAccountvacationmessageon   = "accountvacationmessageon"
	HmAccountsFieldAccountvacationmessage     = "accountvacationmessage"
	HmAccountsFieldAccountvacationsubject     = "accountvacationsubject"
	HmAccountsFieldAccountpwencryption        = "accountpwencryption"
	HmAccountsFieldAccountforwardenabled      = "accountforwardenabled"
	HmAccountsFieldAccountforwardaddress      = "accountforwardaddress"
	HmAccountsFieldAccountforwardkeeporiginal = "accountforwardkeeporiginal"
	HmAccountsFieldAccountenablesignature     = "accountenablesignature"
	HmAccountsFieldAccountsignatureplaintext  = "accountsignatureplaintext"
	HmAccountsFieldAccountsignaturehtml       = "accountsignaturehtml"
	HmAccountsFieldAccountlastlogontime       = "accountlastlogontime"
	HmAccountsFieldAccountvacationexpires     = "accountvacationexpires"
	HmAccountsFieldAccountvacationexpiredate  = "accountvacationexpiredate"
	HmAccountsFieldAccountpersonfirstname     = "accountpersonfirstname"
	HmAccountsFieldAccountpersonlastname      = "accountpersonlastname"
	HmAccountsFieldCreater                    = "creater"
)

func (receiver *HmAccounts) TableName() string {
	return "hm_accounts"
}

type OrmHmAccounts struct {
	db *gorm.DB
}

func (orm *OrmHmAccounts) TableName() string {
	return "hm_accounts"
}

func NewOrmHmAccounts(txs ...*gorm.DB) *OrmHmAccounts {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmAccounts{})
	} else {
		tx = txs[0].Model(&HmAccounts{})
	}
	return &OrmHmAccounts{db: tx}
}

func (orm *OrmHmAccounts) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmAccounts) GetTableInfo() interface{} {
	return &HmAccounts{}
}

// Create insert the value into database
func (orm *OrmHmAccounts) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmAccounts) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmAccounts) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmAccounts) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmAccounts) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmAccounts) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmAccounts) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmAccounts) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmAccounts) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmAccounts) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmAccounts) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmAccounts) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmAccounts) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmAccounts) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmAccounts) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmAccounts) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmAccounts) Unscoped() *OrmHmAccounts {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmAccounts) Insert(row *HmAccounts) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmAccounts) Inserts(rows []*HmAccounts) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmAccounts) Order(value interface{}) *OrmHmAccounts {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmAccounts) Group(name string) *OrmHmAccounts {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmAccounts) Limit(limit int) *OrmHmAccounts {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmAccounts) Offset(offset int) *OrmHmAccounts {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmAccounts) Get() HmAccountsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmAccounts) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmAccounts) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmAccounts{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmAccounts) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_accounts")
}

func (orm *OrmHmAccounts) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmAccounts) First(conds ...interface{}) (*HmAccounts, bool) {
	dest := &HmAccounts{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmAccounts) Take(conds ...interface{}) (*HmAccounts, int64) {
	dest := &HmAccounts{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmAccounts) Last(conds ...interface{}) (*HmAccounts, int64) {
	dest := &HmAccounts{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmAccounts) Find(conds ...interface{}) (HmAccountsList, int64) {
	list := make([]*HmAccounts, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmAccounts) Paginate(page int, limit int) (HmAccountsList, int64) {
	var total int64
	list := make([]*HmAccounts, 0)
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
func (orm *OrmHmAccounts) SimplePaginate(page int, limit int) HmAccountsList {
	list := make([]*HmAccounts, 0)
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
func (orm *OrmHmAccounts) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmAccounts) FirstOrInit(dest *HmAccounts, conds ...interface{}) (*HmAccounts, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmAccounts) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmAccounts) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmAccounts) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmAccounts) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmAccounts) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmAccounts) Where(query interface{}, args ...interface{}) *OrmHmAccounts {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmAccounts) Select(query interface{}, args ...interface{}) *OrmHmAccounts {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmAccounts) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmAccounts) And(fuc func(orm *OrmHmAccounts)) *OrmHmAccounts {
	ormAnd := NewOrmHmAccounts()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmAccounts) Or(fuc func(orm *OrmHmAccounts)) *OrmHmAccounts {
	ormOr := NewOrmHmAccounts()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmAccounts) Preload(query string, args ...interface{}) *OrmHmAccounts {
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
func (orm *OrmHmAccounts) Joins(query string, args ...interface{}) *OrmHmAccounts {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmAccounts) WhereAccountid(val int32) *OrmHmAccounts {
	orm.db.Where("`accountid` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) InsertGetAccountid(row *HmAccounts) int32 {
	orm.db.Create(row)
	return row.Accountid
}
func (orm *OrmHmAccounts) WhereAccountidIn(val []int32) *OrmHmAccounts {
	orm.db.Where("`accountid` IN ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountidGt(val int32) *OrmHmAccounts {
	orm.db.Where("`accountid` > ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountidGte(val int32) *OrmHmAccounts {
	orm.db.Where("`accountid` >= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountidLt(val int32) *OrmHmAccounts {
	orm.db.Where("`accountid` < ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountidLte(val int32) *OrmHmAccounts {
	orm.db.Where("`accountid` <= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountdomainid(val int32) *OrmHmAccounts {
	orm.db.Where("`accountdomainid` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountdomainidIn(val []int32) *OrmHmAccounts {
	orm.db.Where("`accountdomainid` IN ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountdomainidNe(val int32) *OrmHmAccounts {
	orm.db.Where("`accountdomainid` <> ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountadminlevel(val int32) *OrmHmAccounts {
	orm.db.Where("`accountadminlevel` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddress(val string) *OrmHmAccounts {
	orm.db.Where("`accountaddress` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddressIn(val []string) *OrmHmAccounts {
	orm.db.Where("`accountaddress` IN ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddressGt(val string) *OrmHmAccounts {
	orm.db.Where("`accountaddress` > ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddressGte(val string) *OrmHmAccounts {
	orm.db.Where("`accountaddress` >= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddressLt(val string) *OrmHmAccounts {
	orm.db.Where("`accountaddress` < ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddressLte(val string) *OrmHmAccounts {
	orm.db.Where("`accountaddress` <= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountpassword(val string) *OrmHmAccounts {
	orm.db.Where("`accountpassword` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountactive(val int32) *OrmHmAccounts {
	orm.db.Where("`accountactive` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountisad(val int32) *OrmHmAccounts {
	orm.db.Where("`accountisad` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountaddomain(val string) *OrmHmAccounts {
	orm.db.Where("`accountaddomain` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountadusername(val string) *OrmHmAccounts {
	orm.db.Where("`accountadusername` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountmaxsize(val int32) *OrmHmAccounts {
	orm.db.Where("`accountmaxsize` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationmessageon(val int32) *OrmHmAccounts {
	orm.db.Where("`accountvacationmessageon` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationmessage(val string) *OrmHmAccounts {
	orm.db.Where("`accountvacationmessage` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationsubject(val string) *OrmHmAccounts {
	orm.db.Where("`accountvacationsubject` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountpwencryption(val int32) *OrmHmAccounts {
	orm.db.Where("`accountpwencryption` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountforwardenabled(val int32) *OrmHmAccounts {
	orm.db.Where("`accountforwardenabled` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountforwardaddress(val string) *OrmHmAccounts {
	orm.db.Where("`accountforwardaddress` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountforwardkeeporiginal(val int32) *OrmHmAccounts {
	orm.db.Where("`accountforwardkeeporiginal` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountenablesignature(val int32) *OrmHmAccounts {
	orm.db.Where("`accountenablesignature` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountsignatureplaintext(val string) *OrmHmAccounts {
	orm.db.Where("`accountsignatureplaintext` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountsignaturehtml(val string) *OrmHmAccounts {
	orm.db.Where("`accountsignaturehtml` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountlastlogontime(val database.Time) *OrmHmAccounts {
	orm.db.Where("`accountlastlogontime` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountlastlogontimeBetween(begin database.Time, end database.Time) *OrmHmAccounts {
	orm.db.Where("`accountlastlogontime` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountlastlogontimeLte(val database.Time) *OrmHmAccounts {
	orm.db.Where("`accountlastlogontime` <= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountlastlogontimeGte(val database.Time) *OrmHmAccounts {
	orm.db.Where("`accountlastlogontime` >= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationexpires(val uint32) *OrmHmAccounts {
	orm.db.Where("`accountvacationexpires` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationexpiredate(val database.Time) *OrmHmAccounts {
	orm.db.Where("`accountvacationexpiredate` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationexpiredateBetween(begin database.Time, end database.Time) *OrmHmAccounts {
	orm.db.Where("`accountvacationexpiredate` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationexpiredateLte(val database.Time) *OrmHmAccounts {
	orm.db.Where("`accountvacationexpiredate` <= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountvacationexpiredateGte(val database.Time) *OrmHmAccounts {
	orm.db.Where("`accountvacationexpiredate` >= ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountpersonfirstname(val string) *OrmHmAccounts {
	orm.db.Where("`accountpersonfirstname` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereAccountpersonlastname(val string) *OrmHmAccounts {
	orm.db.Where("`accountpersonlastname` = ?", val)
	return orm
}
func (orm *OrmHmAccounts) WhereCreater(val string) *OrmHmAccounts {
	orm.db.Where("`creater` = ?", val)
	return orm
}

type HmAccountsList []*HmAccounts

func (l HmAccountsList) GetAccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Accountid)
	}
	return got
}
func (l HmAccountsList) GetAccountidMap() map[int32]*HmAccounts {
	got := make(map[int32]*HmAccounts)
	for _, val := range l {
		got[val.Accountid] = val
	}
	return got
}
func (l HmAccountsList) GetAccountdomainidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Accountdomainid)
	}
	return got
}
func (l HmAccountsList) GetAccountdomainidMap() map[int32]*HmAccounts {
	got := make(map[int32]*HmAccounts)
	for _, val := range l {
		got[val.Accountdomainid] = val
	}
	return got
}
