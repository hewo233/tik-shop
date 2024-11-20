namespace go hewo.tikshop.base

// 公共的错误结构
exception ErrorResponse {
    1: i64    code,      // 错误码
    2: string message    // 错误信息
}
