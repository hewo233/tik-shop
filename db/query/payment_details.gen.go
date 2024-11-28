// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/hewo/tik-shop/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPaymentDetails(db *gorm.DB, opts ...gen.DOOption) paymentDetails {
	_paymentDetails := paymentDetails{}

	_paymentDetails.paymentDetailsDo.UseDB(db, opts...)
	_paymentDetails.paymentDetailsDo.UseModel(&model.PaymentDetails{})

	tableName := _paymentDetails.paymentDetailsDo.TableName()
	_paymentDetails.ALL = field.NewAsterisk(tableName)
	_paymentDetails.ID = field.NewInt64(tableName, "Id")
	_paymentDetails.CardNumber = field.NewString(tableName, "CardNumber")
	_paymentDetails.ExpiryDate = field.NewString(tableName, "ExpiryDate")
	_paymentDetails.Cvv = field.NewString(tableName, "Cvv")

	_paymentDetails.fillFieldMap()

	return _paymentDetails
}

type paymentDetails struct {
	paymentDetailsDo

	ALL        field.Asterisk
	ID         field.Int64
	CardNumber field.String
	ExpiryDate field.String
	Cvv        field.String

	fieldMap map[string]field.Expr
}

func (p paymentDetails) Table(newTableName string) *paymentDetails {
	p.paymentDetailsDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p paymentDetails) As(alias string) *paymentDetails {
	p.paymentDetailsDo.DO = *(p.paymentDetailsDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *paymentDetails) updateTableName(table string) *paymentDetails {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "Id")
	p.CardNumber = field.NewString(table, "CardNumber")
	p.ExpiryDate = field.NewString(table, "ExpiryDate")
	p.Cvv = field.NewString(table, "Cvv")

	p.fillFieldMap()

	return p
}

func (p *paymentDetails) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *paymentDetails) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["Id"] = p.ID
	p.fieldMap["CardNumber"] = p.CardNumber
	p.fieldMap["ExpiryDate"] = p.ExpiryDate
	p.fieldMap["Cvv"] = p.Cvv
}

func (p paymentDetails) clone(db *gorm.DB) paymentDetails {
	p.paymentDetailsDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p paymentDetails) replaceDB(db *gorm.DB) paymentDetails {
	p.paymentDetailsDo.ReplaceDB(db)
	return p
}

type paymentDetailsDo struct{ gen.DO }

type IPaymentDetailsDo interface {
	gen.SubQuery
	Debug() IPaymentDetailsDo
	WithContext(ctx context.Context) IPaymentDetailsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPaymentDetailsDo
	WriteDB() IPaymentDetailsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPaymentDetailsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPaymentDetailsDo
	Not(conds ...gen.Condition) IPaymentDetailsDo
	Or(conds ...gen.Condition) IPaymentDetailsDo
	Select(conds ...field.Expr) IPaymentDetailsDo
	Where(conds ...gen.Condition) IPaymentDetailsDo
	Order(conds ...field.Expr) IPaymentDetailsDo
	Distinct(cols ...field.Expr) IPaymentDetailsDo
	Omit(cols ...field.Expr) IPaymentDetailsDo
	Join(table schema.Tabler, on ...field.Expr) IPaymentDetailsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPaymentDetailsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPaymentDetailsDo
	Group(cols ...field.Expr) IPaymentDetailsDo
	Having(conds ...gen.Condition) IPaymentDetailsDo
	Limit(limit int) IPaymentDetailsDo
	Offset(offset int) IPaymentDetailsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPaymentDetailsDo
	Unscoped() IPaymentDetailsDo
	Create(values ...*model.PaymentDetails) error
	CreateInBatches(values []*model.PaymentDetails, batchSize int) error
	Save(values ...*model.PaymentDetails) error
	First() (*model.PaymentDetails, error)
	Take() (*model.PaymentDetails, error)
	Last() (*model.PaymentDetails, error)
	Find() ([]*model.PaymentDetails, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PaymentDetails, err error)
	FindInBatches(result *[]*model.PaymentDetails, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PaymentDetails) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPaymentDetailsDo
	Assign(attrs ...field.AssignExpr) IPaymentDetailsDo
	Joins(fields ...field.RelationField) IPaymentDetailsDo
	Preload(fields ...field.RelationField) IPaymentDetailsDo
	FirstOrInit() (*model.PaymentDetails, error)
	FirstOrCreate() (*model.PaymentDetails, error)
	FindByPage(offset int, limit int) (result []*model.PaymentDetails, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPaymentDetailsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p paymentDetailsDo) Debug() IPaymentDetailsDo {
	return p.withDO(p.DO.Debug())
}

func (p paymentDetailsDo) WithContext(ctx context.Context) IPaymentDetailsDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p paymentDetailsDo) ReadDB() IPaymentDetailsDo {
	return p.Clauses(dbresolver.Read)
}

func (p paymentDetailsDo) WriteDB() IPaymentDetailsDo {
	return p.Clauses(dbresolver.Write)
}

func (p paymentDetailsDo) Session(config *gorm.Session) IPaymentDetailsDo {
	return p.withDO(p.DO.Session(config))
}

func (p paymentDetailsDo) Clauses(conds ...clause.Expression) IPaymentDetailsDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p paymentDetailsDo) Returning(value interface{}, columns ...string) IPaymentDetailsDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p paymentDetailsDo) Not(conds ...gen.Condition) IPaymentDetailsDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p paymentDetailsDo) Or(conds ...gen.Condition) IPaymentDetailsDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p paymentDetailsDo) Select(conds ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p paymentDetailsDo) Where(conds ...gen.Condition) IPaymentDetailsDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p paymentDetailsDo) Order(conds ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p paymentDetailsDo) Distinct(cols ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p paymentDetailsDo) Omit(cols ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p paymentDetailsDo) Join(table schema.Tabler, on ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p paymentDetailsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p paymentDetailsDo) RightJoin(table schema.Tabler, on ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p paymentDetailsDo) Group(cols ...field.Expr) IPaymentDetailsDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p paymentDetailsDo) Having(conds ...gen.Condition) IPaymentDetailsDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p paymentDetailsDo) Limit(limit int) IPaymentDetailsDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p paymentDetailsDo) Offset(offset int) IPaymentDetailsDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p paymentDetailsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPaymentDetailsDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p paymentDetailsDo) Unscoped() IPaymentDetailsDo {
	return p.withDO(p.DO.Unscoped())
}

func (p paymentDetailsDo) Create(values ...*model.PaymentDetails) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p paymentDetailsDo) CreateInBatches(values []*model.PaymentDetails, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p paymentDetailsDo) Save(values ...*model.PaymentDetails) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p paymentDetailsDo) First() (*model.PaymentDetails, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentDetails), nil
	}
}

func (p paymentDetailsDo) Take() (*model.PaymentDetails, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentDetails), nil
	}
}

func (p paymentDetailsDo) Last() (*model.PaymentDetails, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentDetails), nil
	}
}

func (p paymentDetailsDo) Find() ([]*model.PaymentDetails, error) {
	result, err := p.DO.Find()
	return result.([]*model.PaymentDetails), err
}

func (p paymentDetailsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PaymentDetails, err error) {
	buf := make([]*model.PaymentDetails, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p paymentDetailsDo) FindInBatches(result *[]*model.PaymentDetails, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p paymentDetailsDo) Attrs(attrs ...field.AssignExpr) IPaymentDetailsDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p paymentDetailsDo) Assign(attrs ...field.AssignExpr) IPaymentDetailsDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p paymentDetailsDo) Joins(fields ...field.RelationField) IPaymentDetailsDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p paymentDetailsDo) Preload(fields ...field.RelationField) IPaymentDetailsDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p paymentDetailsDo) FirstOrInit() (*model.PaymentDetails, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentDetails), nil
	}
}

func (p paymentDetailsDo) FirstOrCreate() (*model.PaymentDetails, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PaymentDetails), nil
	}
}

func (p paymentDetailsDo) FindByPage(offset int, limit int) (result []*model.PaymentDetails, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p paymentDetailsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p paymentDetailsDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p paymentDetailsDo) Delete(models ...*model.PaymentDetails) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *paymentDetailsDo) withDO(do gen.Dao) *paymentDetailsDo {
	p.DO = *do.(*gen.DO)
	return p
}
