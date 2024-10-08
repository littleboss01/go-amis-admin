package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type HmDomains struct {
	Domainid                        int64   `gorm:"column:domainid;autoIncrement;type:bigint(11);primaryKey;index:domainid,class:BTREE,unique" json:"domainid"`                             //
	Domainname                      string  `gorm:"column:domainname;type:varchar(80);index:domainname,class:BTREE,unique;index:idx_hm_domains,class:BTREE;comment:'域名'" json:"domainname"` // 域名
	Domainactive                    int32   `gorm:"column:domainactive;type:tinyint(4);default:1;comment:'是否启用'" json:"domainactive"`                                                       // 是否启用
	Domainpostmaster                string  `gorm:"column:domainpostmaster;type:varchar(255);default:;comment:'邮件管理员'" json:"domainpostmaster"`                                             // 邮件管理员
	Domainmaxsize                   int32   `gorm:"column:domainmaxsize;type:int(11);default:0;comment:'最大邮件大小'" json:"domainmaxsize"`                                                      // 最大邮件大小
	Domainaddomain                  string  `gorm:"column:domainaddomain;type:varchar(255);default:;comment:'域名别名'" json:"domainaddomain"`                                                  // 域名别名
	Domainmaxmessagesize            int32   `gorm:"column:domainmaxmessagesize;type:int(11);default:0;comment:'最大邮件大小'" json:"domainmaxmessagesize"`                                        // 最大邮件大小
	Domainuseplusaddressing         int32   `gorm:"column:domainuseplusaddressing;type:tinyint(4);default:0;comment:'是否启用'" json:"domainuseplusaddressing"`                                 // 是否启用
	Domainplusaddressingchar        string  `gorm:"column:domainplusaddressingchar;type:varchar(255);default:;comment:'别名分隔符'" json:"domainplusaddressingchar"`                             // 别名分隔符
	Domainantispamoptions           int32   `gorm:"column:domainantispamoptions;type:tinyint(4);default:0;comment:'反垃圾邮件'" json:"domainantispamoptions"`                                    // 反垃圾邮件
	Domainenablesignature           int32   `gorm:"column:domainenablesignature;type:tinyint(4);default:0;comment:'是否启用'" json:"domainenablesignature"`                                     // 是否启用
	Domainsignaturemethod           int32   `gorm:"column:domainsignaturemethod;type:tinyint(4);default:1;comment:'签名方式'" json:"domainsignaturemethod"`                                     // 签名方式
	Domainsignatureplaintext        *string `gorm:"column:domainsignatureplaintext;type:varchar(255);default:;comment:'签名内容'" json:"domainsignatureplaintext"`                              // 签名内容
	Domainsignaturehtml             *string `gorm:"column:domainsignaturehtml;type:varchar(255);default:;comment:'签名内容'" json:"domainsignaturehtml"`                                        // 签名内容
	Domainaddsignaturestoreplies    int32   `gorm:"column:domainaddsignaturestoreplies;type:tinyint(4);default:0;comment:'是否启用'" json:"domainaddsignaturestoreplies"`                       // 是否启用
	Domainaddsignaturestolocalemail int32   `gorm:"column:domainaddsignaturestolocalemail;type:tinyint(4);default:0;comment:'是否启用'" json:"domainaddsignaturestolocalemail"`                 // 是否启用
	Domainmaxnoofaccounts           int32   `gorm:"column:domainmaxnoofaccounts;type:int(11);default:0;comment:'最大账户数'" json:"domainmaxnoofaccounts"`                                       // 最大账户数
	Domainmaxnoofaliases            int32   `gorm:"column:domainmaxnoofaliases;type:int(11);default:0;comment:'最大别名数'" json:"domainmaxnoofaliases"`                                         // 最大别名数
	Domainmaxnoofdistributionlists  int32   `gorm:"column:domainmaxnoofdistributionlists;type:int(11);default:0;comment:'最大分组数'" json:"domainmaxnoofdistributionlists"`                     // 最大分组数
	Domainlimitationsenabled        int32   `gorm:"column:domainlimitationsenabled;type:tinyint(4);default:0;comment:'是否启用'" json:"domainlimitationsenabled"`                               // 是否启用
	Domainmaxaccountsize            int32   `gorm:"column:domainmaxaccountsize;type:int(11);default:0;comment:'最大账户大小'" json:"domainmaxaccountsize"`                                        // 最大账户大小
	Domaindkimselector              string  `gorm:"column:domaindkimselector;type:varchar(255);default:;comment:'DKIM选择器'" json:"domaindkimselector"`                                       // DKIM选择器
	Domaindkimprivatekeyfile        string  `gorm:"column:domaindkimprivatekeyfile;type:varchar(255);default:;comment:'DKIM私钥文件'" json:"domaindkimprivatekeyfile"`                          // DKIM私钥文件
	Creater                         *string `gorm:"column:creater;type:varchar(255);default:;comment:'创建人'" json:"creater"`                                                                 // 创建人
}

var (
	HmDomainsFieldDomainid                        = "domainid"
	HmDomainsFieldDomainname                      = "domainname"
	HmDomainsFieldDomainactive                    = "domainactive"
	HmDomainsFieldDomainpostmaster                = "domainpostmaster"
	HmDomainsFieldDomainmaxsize                   = "domainmaxsize"
	HmDomainsFieldDomainaddomain                  = "domainaddomain"
	HmDomainsFieldDomainmaxmessagesize            = "domainmaxmessagesize"
	HmDomainsFieldDomainuseplusaddressing         = "domainuseplusaddressing"
	HmDomainsFieldDomainplusaddressingchar        = "domainplusaddressingchar"
	HmDomainsFieldDomainantispamoptions           = "domainantispamoptions"
	HmDomainsFieldDomainenablesignature           = "domainenablesignature"
	HmDomainsFieldDomainsignaturemethod           = "domainsignaturemethod"
	HmDomainsFieldDomainsignatureplaintext        = "domainsignatureplaintext"
	HmDomainsFieldDomainsignaturehtml             = "domainsignaturehtml"
	HmDomainsFieldDomainaddsignaturestoreplies    = "domainaddsignaturestoreplies"
	HmDomainsFieldDomainaddsignaturestolocalemail = "domainaddsignaturestolocalemail"
	HmDomainsFieldDomainmaxnoofaccounts           = "domainmaxnoofaccounts"
	HmDomainsFieldDomainmaxnoofaliases            = "domainmaxnoofaliases"
	HmDomainsFieldDomainmaxnoofdistributionlists  = "domainmaxnoofdistributionlists"
	HmDomainsFieldDomainlimitationsenabled        = "domainlimitationsenabled"
	HmDomainsFieldDomainmaxaccountsize            = "domainmaxaccountsize"
	HmDomainsFieldDomaindkimselector              = "domaindkimselector"
	HmDomainsFieldDomaindkimprivatekeyfile        = "domaindkimprivatekeyfile"
	HmDomainsFieldCreater                         = "creater"
)

func (receiver *HmDomains) TableName() string {
	return "hm_domains"
}

type OrmHmDomains struct {
	db *gorm.DB
}

func (orm *OrmHmDomains) TableName() string {
	return "hm_domains"
}

func NewOrmHmDomains(txs ...*gorm.DB) *OrmHmDomains {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmDomains{})
	} else {
		tx = txs[0].Model(&HmDomains{})
	}
	return &OrmHmDomains{db: tx}
}

func (orm *OrmHmDomains) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmDomains) GetTableInfo() interface{} {
	return &HmDomains{}
}

// Create insert the value into database
func (orm *OrmHmDomains) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmDomains) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmDomains) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmDomains) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmDomains) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmDomains) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmDomains) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmDomains) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmDomains) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmDomains) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmDomains) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmDomains) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmDomains) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmDomains) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmDomains) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmDomains) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmDomains) Unscoped() *OrmHmDomains {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmDomains) Insert(row *HmDomains) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmDomains) Inserts(rows []*HmDomains) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmDomains) Order(value interface{}) *OrmHmDomains {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmDomains) Group(name string) *OrmHmDomains {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmDomains) Limit(limit int) *OrmHmDomains {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmDomains) Offset(offset int) *OrmHmDomains {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmDomains) Get() HmDomainsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmDomains) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmDomains) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmDomains{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmDomains) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_domains")
}

func (orm *OrmHmDomains) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmDomains) First(conds ...interface{}) (*HmDomains, bool) {
	dest := &HmDomains{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmDomains) Take(conds ...interface{}) (*HmDomains, int64) {
	dest := &HmDomains{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmDomains) Last(conds ...interface{}) (*HmDomains, int64) {
	dest := &HmDomains{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmDomains) Find(conds ...interface{}) (HmDomainsList, int64) {
	list := make([]*HmDomains, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmDomains) Paginate(page int, limit int) (HmDomainsList, int64) {
	var total int64
	list := make([]*HmDomains, 0)
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
func (orm *OrmHmDomains) SimplePaginate(page int, limit int) HmDomainsList {
	list := make([]*HmDomains, 0)
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
func (orm *OrmHmDomains) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmDomains) FirstOrInit(dest *HmDomains, conds ...interface{}) (*HmDomains, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmDomains) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDomains) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmDomains) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmDomains) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmDomains) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmDomains) Where(query interface{}, args ...interface{}) *OrmHmDomains {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmDomains) Select(query interface{}, args ...interface{}) *OrmHmDomains {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmDomains) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmDomains) And(fuc func(orm *OrmHmDomains)) *OrmHmDomains {
	ormAnd := NewOrmHmDomains()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmDomains) Or(fuc func(orm *OrmHmDomains)) *OrmHmDomains {
	ormOr := NewOrmHmDomains()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmDomains) Preload(query string, args ...interface{}) *OrmHmDomains {
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
func (orm *OrmHmDomains) Joins(query string, args ...interface{}) *OrmHmDomains {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmDomains) WhereDomainid(val int64) *OrmHmDomains {
	orm.db.Where("`domainid` = ?", val)
	return orm
}
func (orm *OrmHmDomains) InsertGetDomainid(row *HmDomains) int64 {
	orm.db.Create(row)
	return row.Domainid
}
func (orm *OrmHmDomains) WhereDomainidIn(val []int64) *OrmHmDomains {
	orm.db.Where("`domainid` IN ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainidGt(val int64) *OrmHmDomains {
	orm.db.Where("`domainid` > ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainidGte(val int64) *OrmHmDomains {
	orm.db.Where("`domainid` >= ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainidLt(val int64) *OrmHmDomains {
	orm.db.Where("`domainid` < ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainidLte(val int64) *OrmHmDomains {
	orm.db.Where("`domainid` <= ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainname(val string) *OrmHmDomains {
	orm.db.Where("`domainname` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainnameIn(val []string) *OrmHmDomains {
	orm.db.Where("`domainname` IN ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainnameGt(val string) *OrmHmDomains {
	orm.db.Where("`domainname` > ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainnameGte(val string) *OrmHmDomains {
	orm.db.Where("`domainname` >= ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainnameLt(val string) *OrmHmDomains {
	orm.db.Where("`domainname` < ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainnameLte(val string) *OrmHmDomains {
	orm.db.Where("`domainname` <= ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainactive(val int32) *OrmHmDomains {
	orm.db.Where("`domainactive` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainpostmaster(val string) *OrmHmDomains {
	orm.db.Where("`domainpostmaster` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainmaxsize(val int32) *OrmHmDomains {
	orm.db.Where("`domainmaxsize` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainaddomain(val string) *OrmHmDomains {
	orm.db.Where("`domainaddomain` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainmaxmessagesize(val int32) *OrmHmDomains {
	orm.db.Where("`domainmaxmessagesize` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainuseplusaddressing(val int32) *OrmHmDomains {
	orm.db.Where("`domainuseplusaddressing` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainplusaddressingchar(val string) *OrmHmDomains {
	orm.db.Where("`domainplusaddressingchar` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainantispamoptions(val int32) *OrmHmDomains {
	orm.db.Where("`domainantispamoptions` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainenablesignature(val int32) *OrmHmDomains {
	orm.db.Where("`domainenablesignature` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainsignaturemethod(val int32) *OrmHmDomains {
	orm.db.Where("`domainsignaturemethod` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainsignatureplaintext(val string) *OrmHmDomains {
	orm.db.Where("`domainsignatureplaintext` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainsignaturehtml(val string) *OrmHmDomains {
	orm.db.Where("`domainsignaturehtml` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainaddsignaturestoreplies(val int32) *OrmHmDomains {
	orm.db.Where("`domainaddsignaturestoreplies` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainaddsignaturestolocalemail(val int32) *OrmHmDomains {
	orm.db.Where("`domainaddsignaturestolocalemail` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainmaxnoofaccounts(val int32) *OrmHmDomains {
	orm.db.Where("`domainmaxnoofaccounts` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainmaxnoofaliases(val int32) *OrmHmDomains {
	orm.db.Where("`domainmaxnoofaliases` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainmaxnoofdistributionlists(val int32) *OrmHmDomains {
	orm.db.Where("`domainmaxnoofdistributionlists` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainlimitationsenabled(val int32) *OrmHmDomains {
	orm.db.Where("`domainlimitationsenabled` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomainmaxaccountsize(val int32) *OrmHmDomains {
	orm.db.Where("`domainmaxaccountsize` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomaindkimselector(val string) *OrmHmDomains {
	orm.db.Where("`domaindkimselector` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereDomaindkimprivatekeyfile(val string) *OrmHmDomains {
	orm.db.Where("`domaindkimprivatekeyfile` = ?", val)
	return orm
}
func (orm *OrmHmDomains) WhereCreater(val string) *OrmHmDomains {
	orm.db.Where("`creater` = ?", val)
	return orm
}

type HmDomainsList []*HmDomains

func (l HmDomainsList) GetDomainidList() []int64 {
	got := make([]int64, 0)
	for _, val := range l {
		got = append(got, val.Domainid)
	}
	return got
}
func (l HmDomainsList) GetDomainidMap() map[int64]*HmDomains {
	got := make(map[int64]*HmDomains)
	for _, val := range l {
		got[val.Domainid] = val
	}
	return got
}
