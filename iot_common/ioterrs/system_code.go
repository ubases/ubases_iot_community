/**
 * @Author: lizefang
 * @Date: 2022/08/14 09:00
 */

package ioterrs

// 系统级错误
var (
	Success                       int32 = 0
	ErrBadRequest                 int32 = 400
	ErrUnauthorized               int32 = 401
	ErrForbidden                  int32 = 403
	ErrNotFound                   int32 = 404
	ErrMethodNotAllowed           int32 = 405
	ErrTimeout                    int32 = 408
	ErrConflict                   int32 = 409
	ErrInternalServer             int32 = 500
	ErrServiceUnavailable         int32 = 503
	ErrSystem                     int32 = 100001
	ErrRemoteService              int32 = 100002
	ErrRpc                        int32 = 100003
	ErrIllegalRequest             int32 = 100004
	ErrInvalidUser                int32 = 100005
	ErrRequestBodyLengthOverLimit int32 = 100006
	ErrRequestApiNotFound         int32 = 100007
	ErrHttpMethodNotSupport       int32 = 100008
	ErrIpRequestOverLimit         int32 = 100009
	ErrUserRequestOverLimit       int32 = 100010
	ErrApiRequestOverLimit        int32 = 100011
)
