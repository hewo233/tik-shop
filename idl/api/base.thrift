namespace go hewo.tikshop.route.base

struct BaseResponse {
    1: i64 code;
    2: string message;  // 描述信息
}

struct MessageResponse {
    1: string message, // Return status description
}

struct NilResponse {}

// cart begin::
struct CartItem {
    1: i64 productId
    2: i64 quantity
}

// cart end::

// product begin::
struct Product {
  1: i64 productId
  2: string name
  3: i64 price
  4: i64 stock
  5: string description
}
// product end::

// order begin::
struct OrderItem {
    1: i64 productId
    2: i64 quantity
    3: i64 price
}

struct Order { 
    1: i64 orderId
    2: i64 userId
    3: string status
    4: i64 totalAmount
    5: string createdAt
    6: list<OrderItem> items
}

struct Address {
    1: string street
    2: string city
    3: string postalCode
    4: string country
}

struct PaymentDetails {
    1: string cardNumber 
    2: string expiryDate 
    3: string cvv 
}
// order end::