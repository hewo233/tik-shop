package superquery

import (
	"fmt"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/cart"
)

// GetCart 获取所有的购物车项
func GetCart() ([]*cart.CartItem, error) {

	cartItems, err := query.Q.CartItem.Find() // 使用 Find() 方法查询所有购物车项
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}
	// 直接将 dbmodel.CartItem 类型的数据转换为 cart.CartItem 类型
	items := make([]*cart.CartItem, len(cartItems)) // 创建一个与 cartItems 长度相同的切片
	for i, cartItem := range cartItems {
		// 将 dbmodel.CartItem 类型转换为 cart.CartItem 类型
		items[i] = (*cart.CartItem)(cartItem) // 这里假设结构体完全相同，直接赋值即可
	}
	return items, nil
}
