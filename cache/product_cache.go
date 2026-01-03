package cache

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/hewo/tik-shop/db/model"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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

func (c *ProductCacheImpl) GetProductInfo(ctx context.Context, productID int64) (*model.Product, error) {
	infoKey := GetProductInfoKey(productID)
	stockKey := GetProductStockKey(productID)

	pipe := c.r.Pipeline()
	infocmd := pipe.Get(ctx, infoKey)
	stockcmd := pipe.Get(ctx, stockKey)
	_, _ = pipe.Exec(ctx)

	var pInfo *CachedProduct
	finalProduct := &model.Product{}

	infostr, err := infocmd.Result()
	StockVal, stockErr := stockcmd.Int64()
	if err == nil {
		// info cache hit
		if infostr == NullPlaceholder {
			return nil, gorm.ErrRecordNotFound
		}
		_ = json.Unmarshal([]byte(infostr), pInfo)
		if err := copier.Copy(finalProduct, pInfo); err != nil {
			return nil, err
		}
		if stockErr == nil {
			finalProduct.Stock = StockVal
		} else {
			// stock cache miss, get from db
			dbStock, err := c.ProductCacheSqlManage.GetProductByID(productID)
			if err != nil {
				return nil, err
			}
			finalProduct.Stock = dbStock.Stock
			// update stock cache
			c.r.SetNX(ctx, stockKey, finalProduct.Stock, 0)
		}
		return finalProduct, nil
	} else {
		val, err := c.ProductCacheSqlManage.GetProductByID(productID)
		if err != nil {
			return nil, err
		}
		finalProduct = val
		// update info cache
		cachedPro := &CachedProduct{}
		_ = copier.Copy(cachedPro, finalProduct)
		bytes, _ := json.Marshal(cachedPro)
		err = c.r.SetNX(ctx, infoKey, bytes, DefaultExpire).Err()
		if err != nil {
			return nil, err
		}

		if stockErr == nil {
			// stock cache hit
			finalProduct.Stock = StockVal
			return finalProduct, nil
		}
		// stock cache miss, update stock cache
		log.Println("product stock cache miss, updating cache for productID:", productID)
		err = c.r.SetNX(ctx, stockKey, finalProduct.Stock, 0).Err()
		if err != nil {
			return nil, err
		}
		return finalProduct, nil

	}

}

func (c *ProductCacheImpl) UpdateProductStock(ctx context.Context, productID int64) error {

	infoKey := GetProductInfoKey(productID)
	stockKey := GetProductStockKey(productID)

	c.r.Del(ctx, infoKey, stockKey)

	return nil
}

// ModifyProductStock 这个是给后台用的，只做增量
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

var deductStockScript = redis.NewScript(`
 -- KEYS: stockKey1, stockKey2, ...
 -- ARGV: quantity1, quantity2, ...
 
 -- 第一阶段：检查所有商品库存
 for i = 1, #KEYS do
  local stock = redis.call('GET', KEYS[i])
  if not stock then
   return {-1, i}  -- 返回 [-1, 商品索引]，表示第i个商品库存不存在
  end
  
  stock = tonumber(stock)
  local quantity = tonumber(ARGV[i])
  
  if stock < quantity then
   return {-2, i}  -- 返回 [-2, 商品索引]，表示第i个商品库存不足
  end
 end
 
 -- 第二阶段：所有检查通过后，批量扣减
 for i = 1, #KEYS do
  redis.call('DECRBY', KEYS[i], ARGV[i])
 end
 
 return {0}  -- 成功
`)

func (c *ProductCacheImpl) DeductStock(ctx context.Context, productIDs []int64, quantities []int64) error {

	if len(productIDs) != len(quantities) {
		return errors.New("productIDs and quantities do not match")
	}

	stockKeys := make([]string, len(productIDs))

	for i, productID := range productIDs {
		stockKeys[i] = GetProductStockKey(productID)
	}

	args := make([]interface{}, len(stockKeys))
	for i, qty := range quantities {
		args[i] = qty
	}

	result, err := deductStockScript.Run(ctx, c.r, stockKeys, args...).Result()
	if err != nil {
		return err
	}

	resultCode, ok := result.([]interface{})
	if !ok {
		return errors.New("DeductStock result type error")
	}

	switch resultCode[0] {
	case -1:
		return errors.New("StockNotFound")
	case -2:
		return errors.New("InsufficientStock")
	default:
		return nil
	}
}

func (c *ProductCacheImpl) RollBackDeductStock(ctx context.Context, productIDs []int64, quantities []int64) error {
	pipeline := c.r.Pipeline()
	for i, productId := range productIDs {
		stockKey := GetProductStockKey(productId)
		pipeline.IncrBy(ctx, stockKey, quantities[i])
	}
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
