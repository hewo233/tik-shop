// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q              = new(Query)
	Address        *address
	CartItem       *cartItem
	Order          *order
	OrderItem      *orderItem
	PaymentDetails *paymentDetails
	Product        *product
	Users          *users
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Address = &Q.Address
	CartItem = &Q.CartItem
	Order = &Q.Order
	OrderItem = &Q.OrderItem
	PaymentDetails = &Q.PaymentDetails
	Product = &Q.Product
	Users = &Q.Users
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:             db,
		Address:        newAddress(db, opts...),
		CartItem:       newCartItem(db, opts...),
		Order:          newOrder(db, opts...),
		OrderItem:      newOrderItem(db, opts...),
		PaymentDetails: newPaymentDetails(db, opts...),
		Product:        newProduct(db, opts...),
		Users:          newUsers(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Address        address
	CartItem       cartItem
	Order          order
	OrderItem      orderItem
	PaymentDetails paymentDetails
	Product        product
	Users          users
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		Address:        q.Address.clone(db),
		CartItem:       q.CartItem.clone(db),
		Order:          q.Order.clone(db),
		OrderItem:      q.OrderItem.clone(db),
		PaymentDetails: q.PaymentDetails.clone(db),
		Product:        q.Product.clone(db),
		Users:          q.Users.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		Address:        q.Address.replaceDB(db),
		CartItem:       q.CartItem.replaceDB(db),
		Order:          q.Order.replaceDB(db),
		OrderItem:      q.OrderItem.replaceDB(db),
		PaymentDetails: q.PaymentDetails.replaceDB(db),
		Product:        q.Product.replaceDB(db),
		Users:          q.Users.replaceDB(db),
	}
}

type queryCtx struct {
	Address        IAddressDo
	CartItem       ICartItemDo
	Order          IOrderDo
	OrderItem      IOrderItemDo
	PaymentDetails IPaymentDetailsDo
	Product        IProductDo
	Users          IUsersDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Address:        q.Address.WithContext(ctx),
		CartItem:       q.CartItem.WithContext(ctx),
		Order:          q.Order.WithContext(ctx),
		OrderItem:      q.OrderItem.WithContext(ctx),
		PaymentDetails: q.PaymentDetails.WithContext(ctx),
		Product:        q.Product.WithContext(ctx),
		Users:          q.Users.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
