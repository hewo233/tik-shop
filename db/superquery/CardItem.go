package superquery

import (
	"errors"
	"fmt"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	cart "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart"
	"gorm.io/gorm"
	"log"
)

var c = query.Q.CartItem

func GetCart(request *cart.GetCartRequest) ([]*cart.CartItem, error) {

	cartItems, err := c.Where(c.UserId.Eq(uint(request.UserId))).Find()
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}
	items := make([]*cart.CartItem, len(cartItems))
	for i, cartItem := range cartItems {
		// 将 dbmodel.CartItem 类型转换为 cart.CartItem 类型
		items[i] = &cart.CartItem{
			cartItem.ProductId,
			cartItem.Quantity,
		}
	}
	return items, nil
}

func AddToCart(request *cart.AddToCartRequest) (*cart.AddToCartResponse, error) {
	userID := uint(request.UserId)
	productID := request.ProductId
	quantity := request.Quantity

	// 查询商品的库存
	product, err := query.Q.Product.Where(query.Product.Id.Eq(productID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product not found")
		}
		log.Println("Error querying product:", err)
		return nil, fmt.Errorf("failed to retrieve product details: %w", err)
	}

	// 检查库存是否足够
	if product.Stock < quantity {
		return nil, fmt.Errorf("insufficient stock. Available stock: %d", product.Stock)
	}

	// 查询用户购物车中是否已经有这个商品
	_, err = c.Where(c.UserId.Eq(userID), c.ProductId.Eq(productID)).First()
	if err != nil {
		// 如果没有找到，说明购物车中没有该商品，插入新商品
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCartItem := &model.CartItem{
				UserId:    userID,
				ProductId: productID,
				Quantity:  quantity,
			}
			err := c.Create(newCartItem)
			if err != nil {
				log.Println("Error adding product to cart:", err)
				return nil, fmt.Errorf("failed to add product to cart: %w", err)
			}
			return &cart.AddToCartResponse{
				Message: "Product added to cart successfully",
			}, nil
		}
		// 处理其他查询错误
		log.Println("Error querying cart item:", err)
		return nil, fmt.Errorf("failed to check cart item: %w", err)
	}

	_, err = c.Where(c.UserId.Eq(userID), c.ProductId.Eq(productID)).Update(c.Quantity, quantity)
	if err != nil {
		log.Println("Error updating cart item:", err)
		return nil, fmt.Errorf("failed to update cart item: %w", err)
	}

	return &cart.AddToCartResponse{
		Message: "Product quantity updated in cart successfully",
	}, nil
}

// UpdateCart implements the CartServiceImpl interface.
func UpdateCart(request *cart.UpdateCartRequest) (resp *cart.UpdateCartResponse, err error) {
	userID := uint(request.UserId)
	productID := request.ProductId
	quantity := request.Quantity

	_, err = c.Where(c.UserId.Eq(userID), c.ProductId.Eq(productID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cart item not found for user %d and product %d", userID, productID)
		}
		// 其他数据库错误
		return nil, fmt.Errorf("failed to query cart item: %w", err)
	}

	_, err = c.Where(c.UserId.Eq(userID), c.ProductId.Eq(productID)).Update(c.Quantity, quantity)
	if err != nil {
		// 如果保存失败，返回错误
		return nil, fmt.Errorf("failed to update cart item: %w", err)
	}
	resp = &cart.UpdateCartResponse{
		Message: "Cart updated successfully",
	}
	return resp, nil
}

// RemoveFromCart implements the CartServiceImpl interface.
func RemoveFromCart(request *cart.RemoveFromCartRequest) (resp *cart.RemoveFromCartResponse, err error) {
	// 从请求中提取用户 ID 和商品 ID
	userID := uint(request.UserId)
	productID := request.ProductId

	_, err = c.Where(c.UserId.Eq(userID), c.ProductId.Eq(productID)).First()
	if err != nil {
		// 如果购物车项不存在，返回错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cart item not found for user %d and product %d", userID, productID)
		}
		// 其他数据库错误
		return nil, fmt.Errorf("failed to query cart item: %w", err)
	}

	// 执行删除操作
	_, err = c.Where(c.UserId.Eq(userID), c.ProductId.Eq(productID)).Delete()
	if err != nil {
		// 如果删除失败，返回错误
		return nil, fmt.Errorf("failed to remove product from cart: %w", err)
	}

	resp = &cart.RemoveFromCartResponse{
		Message: "Product removed from cart successfully",
	}
	return resp, nil
}

// ClearCart implements the CartServiceImpl interface.
func ClearCart(request *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {
	// 从请求中提取用户 ID
	userID := uint(request.UserId)

	cartItems, err := c.Where(c.UserId.Eq(userID)).Delete()
	// 错误处理
	if err != nil {
		// 如果发生删除错误，返回错误信息
		return nil, fmt.Errorf("failed to delete cart items for user %d: %w", request.UserId, err)
	}

	// 检查是否删除了任何记录
	if cartItems.RowsAffected == 0 {
		// 如果没有记录被删除，返回提示信息
		return &cart.ClearCartResponse{
			Message: fmt.Sprintf("No cart items found for user %d", request.UserId),
		}, nil
	}

	// 删除成功，返回成功消息
	return &cart.ClearCartResponse{
		Message: fmt.Sprintf("Cart items for user %d cleared successfully", request.UserId),
	}, nil
}
