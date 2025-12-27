package cache

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/hewo/tik-shop/db/model"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"log"
)

type ProductCacheImpl struct {
	r *redis.Client
	ProductCacheSqlManage
}

type ProductCacheSqlManage interface {
	GetProductByID(id int64) (productRet *model.Product, err error)
}

func NewProductCacheImpl(r *redis.Client) *ProductCacheImpl {
	return &ProductCacheImpl{r: r}
}

func GetProductInfoKey(productID int64) string {
	return ProductInfoKey + string(rune(productID))
}

func GetProductStockKey(productID int64) string {
	return ProductStockKey + string(rune(productID))
}

func (c *ProductCacheImpl) GetProductInfo(ctx context.Context, productID int64) (error, *model.Product) {
	infoKey := GetProductInfoKey(productID)
	stockKey := GetProductStockKey(productID)

	val, err := c.r.Get(ctx, infoKey).Result()

	if err == nil {
		// cache hit
		log.Println("product info cache hit for productID:", productID)
		cproduct := &CachedProduct{}
		err := json.Unmarshal([]byte(val), cproduct)
		if err != nil {
			return err, nil
		}
		// stock
		stockVal, err := c.r.Get(ctx, stockKey).Result()
		if err == nil {
			var stock int64
			err = json.Unmarshal([]byte(stockVal), &stock)
			if err != nil {
				return err, nil
			}
			product := &model.Product{}
			if err := copier.Copy(product, cproduct); err != nil {
				return err, nil
			}
			product.Stock = stock
			return nil, product
		}
	}

	// cache miss
	product, err := c.ProductCacheSqlManage.GetProductByID(productID)
	if err != nil {
		return err, nil
	}

	cProduct := &CachedProduct{}
	if err := copier.Copy(cProduct, product); err != nil {
		return err, nil
	}

	infoToCache, err := json.Marshal(cProduct)
	if err != nil {
		return err, nil
	}

	pipe := c.r.Pipeline()
	pipe.Set(ctx, infoKey, infoToCache, DefaultExpire)
	pipe.Set(ctx, stockKey, product.Stock, DefaultExpire)
	_, pipeErr := pipe.Exec(ctx)
	if pipeErr != nil {
		return pipeErr, nil
	}
	return nil, product
}

func (c *ProductCacheImpl) UpdateProductStock(ctx context.Context, productID int64) error {

	infoKey := GetProductInfoKey(productID)
	stockKey := GetProductStockKey(productID)

	c.r.Del(ctx, infoKey, stockKey)

	return nil
}

func (c *ProductCacheImpl) ModifyProductStock(ctx context.Context, productID int64, delta int64) error {
	stockKey := GetProductStockKey(productID)

	err := c.r.IncrBy(ctx, stockKey, delta).Err()
	return err
}

func (c *ProductCacheImpl) DeleteProductCache(ctx context.Context, productID int64) error {
	infoKey := GetProductInfoKey(productID)
	stockKey := GetProductStockKey(productID)

	c.r.Del(ctx, infoKey, stockKey)

	return nil
}
