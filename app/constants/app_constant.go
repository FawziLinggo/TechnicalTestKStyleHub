package constants

const (
	ConstantLogFormat   = `{"host":"${host}","pid":"${pid}","time":"${time}","request_id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}","user-agent":"${ua}","bytes_in":"${bytesReceived}","bytes_out":"${bytesSent}"}` + "\n"
	ConstantLogTimeZone = "Asia/Jakarta"
	ConstantMethodCORS  = "GET,POST,HEAD,PUT,DELETE,PATCH"
	ConstantHeaderCORS  = "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Cache-Control, X-Requested-With, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods, Access-Control-Allow-Credentials"
)

const (
	Authorization       = "Authorization"
	AuthorizationBearer = "Bearer"
	AuthorizationBasic  = "Basic"
)

const (
	ContextUserData = "user_data"
)
