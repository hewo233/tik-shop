package superquery

import (
	"fmt"

	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/product"
	"github.com/jinzhu/copier"
)

var p = &query.Q.Product

func GetProducts(page, limit int64) (products []*product.Product, err error) {
	ps, err := p.Limit(int(limit)).Offset(int(limit) * int(page)).Find()
	if err != nil {
		return nil, fmt.Errorf("products can't be fetched from db: %v", err)
	}
	products = make([]*product.Product, len(ps))
	err = copier.Copy(&products, ps)
	if err != nil {
		return nil, fmt.Errorf("products fetched from db can be copied: %v", err)
	}
	return
}

func GetProductById(id int64) (product *product.Product, err error) {
	tmp, err := p.Where(p.Id.Eq(id)).Find()
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&product, tmp)
	if err != nil {
		return nil, err
	}
	return
}

func CreateProduct(product *model.Product) error {
	err := p.Create(product)
	if err != nil {
		return err
	}
	return nil
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
