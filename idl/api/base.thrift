namespace go hewo.tikshop.route.base

struct BaseErrorResponse {
  1: string code       // Error code, e.g., "BAD_REQUEST", "NOT_FOUND"
  2: string message    // Human-readable error message
  3: optional string details // Additional error details, e.g., JSON string
}
