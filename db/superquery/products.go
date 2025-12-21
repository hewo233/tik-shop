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
	productRet, err = p.Preload(p.Merchant).Where(p.ID.Eq(id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	return productRet, nil
}

func (m *ProductSqlManageImpl) ListProducts(merchantID int64, offset int, limit int) (products []*model.Product, err error) {
	productsRet, err := p.Preload(p.Merchant).Offset(offset).Limit(limit).Where(p.ID.Eq(merchantID)).Find()
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return productsRet, nil
}

func UpdateProduct(product *model.Product) error {
	err := p.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int64) (err error) {
	ref, err := p.Where(p.Id.Eq(id)).First()
	if err != nil {
		return err
	}
	_, err = p.Delete(ref)
	if err != nil {
		return err
	}
	return nil
}

*/
