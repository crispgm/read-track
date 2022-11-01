// Package app serves
package app

// Response .
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse creates a new response body
func NewResponse() *Response {
	return &Response{
		Code:    0,
		Message: "success",
	}
}

// SetCode .
func (r *Response) SetCode(code int) {
	r.Code = code
}

// SetMessage .
func (r *Response) SetMessage(msg string) {
	r.Message = msg
}

// SetData .
func (r *Response) SetData(data interface{}) {
	r.Data = data
}
