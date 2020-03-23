package defination

//程序error code
type ErrorCode int32

//http 业务部分ErrorCode
const (
	ErrorCodeSuccess          ErrorCode = 0
	ErrorCodeBindJSONFailed   ErrorCode = -1
	ErrorCodeUserExists       ErrorCode = -2
	ErrorCodeInsertUserFailed ErrorCode = -3
)
