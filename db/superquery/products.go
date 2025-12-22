package superquery

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
)

var p = &query.Q.Product

type ProductSqlManageImpl struct{}

func NewProductSqlManageImpl() *ProductSqlManageImpl {
	return &ProductSqlManageImpl{}
}

func (m *ProductSqlManageImpl) CreateProduct(product *model.Product) (productID int64, err error) {
	err = p.Create(product)
	if err != nil {
		return -1, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return product.ID, nil
}

func (m *ProductSqlManageImpl) GetProductByID(id int64) (productRet *model.Product, err error) {
	productRet, err = p.Preload(p.Merchant).Where(p.ID.Eq(id)).Where(p.Status.Neq(0)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	return productRet, nil
}

func (m *ProductSqlManageImpl) ListProducts(merchantID int64, offset int, limit int) (products []*model.Product, err error) {
	productsRet, err := p.Preload(p.Merchant).Offset(offset).Limit(limit).Where(p.ID.Eq(merchantID)).Where(p.Status.Neq(0)).Find()
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return productsRet, nil
}

func (m *ProductSqlManageImpl) CheckAndGetProduct(productID int64, merchantID int64) (*model.Product, error) {
	existed, err := p.Preload(p.Merchant).Where(p.ID.Eq(productID)).Where(p.Status.Neq(0)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	if existed.MerchantID != merchantID {
		return nil, &base.ErrorResponse{Code: consts.StatusForbidden, Message: "permission denied: not the owner of the product"}
	}
	return existed, nil
}

func (m *ProductSqlManageImpl) UpdateProductByID(product *model.Product) error {
	result, err := p.Where(p.ID.Eq(product.ID)).Updates(product)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return &base.ErrorResponse{Code: consts.StatusNotFound, Message: "product not found"}
	}
	return nil
}

func (m *ProductSqlManageImpl) DeleteProductByID(productID int64) (err error) {

	result, err := p.Where(p.ID.Eq(productID)).Delete()
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return &base.ErrorResponse{Code: consts.StatusNotFound, Message: "product not found"}
	}

	return nil
}

func (m *ProductSqlManageImpl) ModifyStockByID(product *model.Product) error {
	// 乐观锁
	oldUpdatedAt := product.UpdatedAt
	result, err := p.Where(p.ID.Eq(product.ID)).Where(p.UpdatedAt.Eq(oldUpdatedAt)).Updates(product)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return &base.ErrorResponse{Code: consts.StatusConflict, Message: "update stock failed due to concurrent modification"}
	}
	return nil

}
