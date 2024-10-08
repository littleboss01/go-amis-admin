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

type HmFetchaccounts struct {
	Faid                    int32         `gorm:"column:faid;autoIncrement;type:int(11);primaryKey;index:faid,class:BTREE,unique" json:"faid"` //
	Faactive                int32         `gorm:"column:faactive;type:tinyint(4)" json:"faactive"`                                             //
	Faaccountid             int32         `gorm:"column:faaccountid;type:int(11)" json:"faaccountid"`                                          //
	Faaccountname           string        `gorm:"column:faaccountname;type:varchar(255)" json:"faaccountname"`                                 //
	Faserveraddress         string        `gorm:"column:faserveraddress;type:varchar(255)" json:"faserveraddress"`                             //
	Faserverport            int32         `gorm:"column:faserverport;type:int(11)" json:"faserverport"`                                        //
	Faservertype            int32         `gorm:"column:faservertype;type:tinyint(4)" json:"faservertype"`                                     //
	Fausername              string        `gorm:"column:fausername;type:varchar(255)" json:"fausername"`                                       //
	Fapassword              string        `gorm:"column:fapassword;type:varchar(255)" json:"fapassword"`                                       //
	Faminutes               int32         `gorm:"column:faminutes;type:int(11)" json:"faminutes"`                                              //
	Fanexttry               database.Time `gorm:"column:fanexttry;type:datetime" json:"fanexttry"`                                             //
	Fadaystokeep            int32         `gorm:"column:fadaystokeep;type:int(11)" json:"fadaystokeep"`                                        //
	Falocked                int32         `gorm:"column:falocked;type:tinyint(4)" json:"falocked"`                                             //
	Faprocessmimerecipients int32         `gorm:"column:faprocessmimerecipients;type:tinyint(4)" json:"faprocessmimerecipients"`               //
	Faprocessmimedate       int32         `gorm:"column:faprocessmimedate;type:tinyint(4)" json:"faprocessmimedate"`                           //
	Faconnectionsecurity    int32         `gorm:"column:faconnectionsecurity;type:tinyint(4)" json:"faconnectionsecurity"`                     //
	Fauseantispam           int32         `gorm:"column:fauseantispam;type:tinyint(4)" json:"fauseantispam"`                                   //
	Fauseantivirus          int32         `gorm:"column:fauseantivirus;type:tinyint(4)" json:"fauseantivirus"`                                 //
	Faenablerouterecipients int32         `gorm:"column:faenablerouterecipients;type:tinyint(4)" json:"faenablerouterecipients"`               //
}

var (
	HmFetchaccountsFieldFaid                    = "faid"
	HmFetchaccountsFieldFaactive                = "faactive"
	HmFetchaccountsFieldFaaccountid             = "faaccountid"
	HmFetchaccountsFieldFaaccountname           = "faaccountname"
	HmFetchaccountsFieldFaserveraddress         = "faserveraddress"
	HmFetchaccountsFieldFaserverport            = "faserverport"
	HmFetchaccountsFieldFaservertype            = "faservertype"
	HmFetchaccountsFieldFausername              = "fausername"
	HmFetchaccountsFieldFapassword              = "fapassword"
	HmFetchaccountsFieldFaminutes               = "faminutes"
	HmFetchaccountsFieldFanexttry               = "fanexttry"
	HmFetchaccountsFieldFadaystokeep            = "fadaystokeep"
	HmFetchaccountsFieldFalocked                = "falocked"
	HmFetchaccountsFieldFaprocessmimerecipients = "faprocessmimerecipients"
	HmFetchaccountsFieldFaprocessmimedate       = "faprocessmimedate"
	HmFetchaccountsFieldFaconnectionsecurity    = "faconnectionsecurity"
	HmFetchaccountsFieldFauseantispam           = "fauseantispam"
	HmFetchaccountsFieldFauseantivirus          = "fauseantivirus"
	HmFetchaccountsFieldFaenablerouterecipients = "faenablerouterecipients"
)

func (receiver *HmFetchaccounts) TableName() string {
	return "hm_fetchaccounts"
}

type OrmHmFetchaccounts struct {
	db *gorm.DB
}

func (orm *OrmHmFetchaccounts) TableName() string {
	return "hm_fetchaccounts"
}

func NewOrmHmFetchaccounts(txs ...*gorm.DB) *OrmHmFetchaccounts {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&HmFetchaccounts{})
	} else {
		tx = txs[0].Model(&HmFetchaccounts{})
	}
	return &OrmHmFetchaccounts{db: tx}
}

func (orm *OrmHmFetchaccounts) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmHmFetchaccounts) GetTableInfo() interface{} {
	return &HmFetchaccounts{}
}

// Create insert the value into database
func (orm *OrmHmFetchaccounts) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmHmFetchaccounts) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmHmFetchaccounts) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmHmFetchaccounts) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmHmFetchaccounts) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmHmFetchaccounts) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmHmFetchaccounts) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmHmFetchaccounts) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmHmFetchaccounts) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmHmFetchaccounts) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmHmFetchaccounts) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmHmFetchaccounts) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmHmFetchaccounts) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmHmFetchaccounts) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmHmFetchaccounts) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmHmFetchaccounts) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

func (orm *OrmHmFetchaccounts) Unscoped() *OrmHmFetchaccounts {
	orm.db.Unscoped()
	return orm
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmHmFetchaccounts) Insert(row *HmFetchaccounts) error {
	return orm.db.Create(row).Error
}

func (orm *OrmHmFetchaccounts) Inserts(rows []*HmFetchaccounts) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmHmFetchaccounts) Order(value interface{}) *OrmHmFetchaccounts {
	orm.db.Order(value)
	return orm
}

func (orm *OrmHmFetchaccounts) Group(name string) *OrmHmFetchaccounts {
	orm.db.Group(name)
	return orm
}

func (orm *OrmHmFetchaccounts) Limit(limit int) *OrmHmFetchaccounts {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmHmFetchaccounts) Offset(offset int) *OrmHmFetchaccounts {
	orm.db.Offset(offset)
	return orm
}

// Get 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmHmFetchaccounts) Get() HmFetchaccountsList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmHmFetchaccounts) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmHmFetchaccounts) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&HmFetchaccounts{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmHmFetchaccounts) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM hm_fetchaccounts")
}

func (orm *OrmHmFetchaccounts) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmHmFetchaccounts) First(conds ...interface{}) (*HmFetchaccounts, bool) {
	dest := &HmFetchaccounts{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmHmFetchaccounts) Take(conds ...interface{}) (*HmFetchaccounts, int64) {
	dest := &HmFetchaccounts{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmHmFetchaccounts) Last(conds ...interface{}) (*HmFetchaccounts, int64) {
	dest := &HmFetchaccounts{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmHmFetchaccounts) Find(conds ...interface{}) (HmFetchaccountsList, int64) {
	list := make([]*HmFetchaccounts, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmHmFetchaccounts) Paginate(page int, limit int) (HmFetchaccountsList, int64) {
	var total int64
	list := make([]*HmFetchaccounts, 0)
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
func (orm *OrmHmFetchaccounts) SimplePaginate(page int, limit int) HmFetchaccountsList {
	list := make([]*HmFetchaccounts, 0)
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
func (orm *OrmHmFetchaccounts) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmHmFetchaccounts) FirstOrInit(dest *HmFetchaccounts, conds ...interface{}) (*HmFetchaccounts, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmHmFetchaccounts) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmFetchaccounts) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmHmFetchaccounts) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmHmFetchaccounts) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmHmFetchaccounts) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmHmFetchaccounts) Where(query interface{}, args ...interface{}) *OrmHmFetchaccounts {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmHmFetchaccounts) Select(query interface{}, args ...interface{}) *OrmHmFetchaccounts {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmHmFetchaccounts) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmHmFetchaccounts) And(fuc func(orm *OrmHmFetchaccounts)) *OrmHmFetchaccounts {
	ormAnd := NewOrmHmFetchaccounts()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmHmFetchaccounts) Or(fuc func(orm *OrmHmFetchaccounts)) *OrmHmFetchaccounts {
	ormOr := NewOrmHmFetchaccounts()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmHmFetchaccounts) Preload(query string, args ...interface{}) *OrmHmFetchaccounts {
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
func (orm *OrmHmFetchaccounts) Joins(query string, args ...interface{}) *OrmHmFetchaccounts {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmHmFetchaccounts) WhereFaid(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faid` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) InsertGetFaid(row *HmFetchaccounts) int32 {
	orm.db.Create(row)
	return row.Faid
}
func (orm *OrmHmFetchaccounts) WhereFaidIn(val []int32) *OrmHmFetchaccounts {
	orm.db.Where("`faid` IN ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaidGt(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faid` > ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaidGte(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faid` >= ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaidLt(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faid` < ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaidLte(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faid` <= ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaactive(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faactive` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaaccountid(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faaccountid` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaaccountidIn(val []int32) *OrmHmFetchaccounts {
	orm.db.Where("`faaccountid` IN ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaaccountidNe(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faaccountid` <> ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaaccountname(val string) *OrmHmFetchaccounts {
	orm.db.Where("`faaccountname` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaserveraddress(val string) *OrmHmFetchaccounts {
	orm.db.Where("`faserveraddress` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaserverport(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faserverport` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaservertype(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faservertype` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFausername(val string) *OrmHmFetchaccounts {
	orm.db.Where("`fausername` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFapassword(val string) *OrmHmFetchaccounts {
	orm.db.Where("`fapassword` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaminutes(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faminutes` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFanexttry(val database.Time) *OrmHmFetchaccounts {
	orm.db.Where("`fanexttry` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFanexttryBetween(begin database.Time, end database.Time) *OrmHmFetchaccounts {
	orm.db.Where("`fanexttry` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFanexttryLte(val database.Time) *OrmHmFetchaccounts {
	orm.db.Where("`fanexttry` <= ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFanexttryGte(val database.Time) *OrmHmFetchaccounts {
	orm.db.Where("`fanexttry` >= ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFadaystokeep(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`fadaystokeep` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFalocked(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`falocked` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaprocessmimerecipients(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faprocessmimerecipients` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaprocessmimedate(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faprocessmimedate` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaconnectionsecurity(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faconnectionsecurity` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFauseantispam(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`fauseantispam` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFauseantivirus(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`fauseantivirus` = ?", val)
	return orm
}
func (orm *OrmHmFetchaccounts) WhereFaenablerouterecipients(val int32) *OrmHmFetchaccounts {
	orm.db.Where("`faenablerouterecipients` = ?", val)
	return orm
}

type HmFetchaccountsList []*HmFetchaccounts

func (l HmFetchaccountsList) GetFaidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Faid)
	}
	return got
}
func (l HmFetchaccountsList) GetFaidMap() map[int32]*HmFetchaccounts {
	got := make(map[int32]*HmFetchaccounts)
	for _, val := range l {
		got[val.Faid] = val
	}
	return got
}
func (l HmFetchaccountsList) GetFaaccountidList() []int32 {
	got := make([]int32, 0)
	for _, val := range l {
		got = append(got, val.Faaccountid)
	}
	return got
}
func (l HmFetchaccountsList) GetFaaccountidMap() map[int32]*HmFetchaccounts {
	got := make(map[int32]*HmFetchaccounts)
	for _, val := range l {
		got[val.Faaccountid] = val
	}
	return got
}
