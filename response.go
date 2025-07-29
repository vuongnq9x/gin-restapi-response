package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response đại diện cho cấu trúc phản hồi chuẩn của HTTP API
type Response struct {
	// Code là mã trạng thái HTTP
	Code int `json:"code"`
	// Success cho biết yêu cầu có thành công hay không
	Success bool `json:"success"`
	// Message là thông báo mô tả kết quả
	Message string `json:"message"`
	// Error chứa thông tin lỗi (nếu có)
	Error any `json:"error,omitempty"`
	// Data chứa dữ liệu trả về (nếu có)
	Data any `json:"data,omitempty"`
}

// Success trả về response thành công với HTTP status 200 OK
func Success(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Ok(http.StatusOK, message).WithData(data))
}

// Created trả về response thành công với HTTP status 201 Created
func Created(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, Ok(http.StatusCreated, message).WithData(data))
}

// NoContent trả về response thành công với HTTP status 204 No Content
func NoContent(c *gin.Context, message string) {
	c.JSON(http.StatusNoContent, Ok(http.StatusNoContent, message))
}

// BadRequest trả về response lỗi với HTTP status 400 Bad Request
func BadRequest(c *gin.Context, message string, error any, data any) {
	c.JSON(http.StatusBadGateway, Err(http.StatusBadGateway, message).WithError(error).WithData(data))
}

// Unauthorized trả về response lỗi với HTTP status 401 Unauthorized
func Unauthorized(c *gin.Context, message string, data any) {
	c.JSON(http.StatusUnauthorized, Err(http.StatusUnauthorized, message).WithData(data))
}

// Forbidden trả về response lỗi với HTTP status 403 Forbidden
func Forbidden(c *gin.Context, message string, data any) {
	c.JSON(http.StatusForbidden, Err(http.StatusForbidden, message).WithData(data))
}

// NotFound trả về response lỗi với HTTP status 404 Not Found
func NotFound(c *gin.Context, message string, data any) {
	c.JSON(http.StatusNotFound, Err(http.StatusNotFound, message).WithData(data))
}

// Conflict trả về response lỗi với HTTP status 409 Conflict
func Conflict(c *gin.Context, message string, data any) {
	c.JSON(http.StatusConflict, Err(http.StatusConflict, message).WithData(data))
}

// UnprocessableEntity trả về response lỗi với HTTP status 422 Unprocessable Entity
func UnprocessableEntity(c *gin.Context, message string, error any, data any) {
	c.JSON(http.StatusUnprocessableEntity, Err(http.StatusUnprocessableEntity, message).WithError(error).WithData(data))
}

// ServiceUnavailable trả về response lỗi với HTTP status 503 Service Unavailable
func ServiceUnavailable(c *gin.Context, message string, error any, data any) {
	c.JSON(http.StatusServiceUnavailable, Err(http.StatusServiceUnavailable, message).WithError(error).WithData(data))
}

// InternalServerError trả về response lỗi với HTTP status 500 Internal Server Error
func InternalServerError(c *gin.Context, message string, error any, data any) {
	c.JSON(http.StatusInternalServerError, Err(http.StatusInternalServerError, message).WithError(error).WithData(data))
}

// New tạo một đối tượng Response mới với giá trị mặc định
func New() *Response {
	return &Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "OK",
		Data:    nil,
	}
}

// WithCode thiết lập mã HTTP cho response
func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

// WithSuccess thiết lập trạng thái thành công cho response
func (r *Response) WithSuccess(success bool) *Response {
	r.Success = success
	return r
}

// WithMessage thiết lập thông báo cho response
func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

// WithData thiết lập dữ liệu cho response
func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

// WithError thiết lập thông tin lỗi cho response
func (r *Response) WithError(error any) *Response {
	r.Error = error
	return r
}

// Ok tạo một response thành công với mã và thông báo được chỉ định
func Ok(code int, message string) *Response {
	return New().
		WithSuccess(true).
		WithCode(code).
		WithMessage(msg(message, code))
}

// Err tạo một response lỗi với mã và thông báo được chỉ định
func Err(code int, message string) *Response {
	return New().
		WithSuccess(false).
		WithCode(code).
		WithMessage(msg(message, code))
}

// msg trả về thông báo tùy chỉnh hoặc thông báo mặc định dựa trên mã HTTP
func msg(message string, code int) string {
	if message != "" {
		return message
	}

	return http.StatusText(code)
}
