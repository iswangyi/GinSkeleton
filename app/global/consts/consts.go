package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (
	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	//服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误,检测文件mime类型发生错误。"

	// token相关
	JwtTokenSignKey         string = "goskeleton"
	JwtTokenCreatedExpireAt int64  = 3600    // 创建时token默认有效期3600秒
	JwtTokenRefreshExpireAt int64  = 7200    // 刷新token时，延长7200秒
	JwtTokenOK              int    = 200100  //token有效
	JwtTokenInvalid         int    = -400100 //无效的token
	JwtTokenExpired         int    = -400101 //过期的token
	JwtTokenOnlineUsers     int    = 10      // 设置一个账号最大允许几个用户同时在线，默认为10

	//snowflake错误
	SnowFlakeMachineId      int16  = 1024
	SnowFlakeMachineIllegal string = "SnowFlake数据越界，大于65535"

	// CURD 常用业务状态码
	CurdStatusOkCode         int    = 200
	CurdStatusOkMsg          string = "Success"
	CurdCreatFailCode        int    = -400200
	CurdCreatFailMsg         string = "新增失败"
	CurdUpdateFailCode       int    = -400201
	CurdUpdateFailMsg        string = "更新失败"
	CurdDeleteFailCode       int    = -400202
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = -400203
	CurdSelectFailMsg        string = "查询无数据"
	CurdRegisterFailCode     int    = -400204
	CurdRegisterFailMsg      string = "注册失败"
	CurdLoginFailCode        int    = -400205
	CurdLoginFailMsg         string = "登录失败"
	CurdRefreshTokenFailCode int    = -400206
	CurdRefreshTokenFailMsg  string = "刷新Token失败"

	//文件上传
	FilesUploadFailCode            int    = -400250
	FilesUploadFailMsg             string = "文件上传失败, 获取上传文件发生错误!"
	FilesUploadMoreThanMaxSizeCode int    = -400251
	FilesUploadMoreThanMaxSizeMsg  string = "长传文件超过系统设定的最大值,系统允许的最大值（M）："
	FilesUploadMimeTypeFailCode    int    = -400252
	FilesUploadMimeTypeFailMsg     string = "文件mime类型不允许"

	//websocket
	WsServerNotStartCode int    = -400300
	WsServerNotStartMsg  string = "websocket 服务没有开启，请在配置文件开启，相关路径：config/config.yml"
	WsOpenFailCode       int    = -400301
	WsOpenFailMsg        string = "websocket open阶段初始化基本参数失败"

	//验证码
	CaptchaGetParamsInvalidMsg    string = "获取验证码：提交的验证码参数无效,请检查验证码ID以及文件名后缀是否完整"
	CaptchaGetParamsInvalidCode   int    = -400350
	CaptchaCheckParamsInvalidMsg  string = "校验验证码：提交的参数无效，请确保提交的验证码ID和值有效"
	CaptchaCheckParamsInvalidCode int    = -400351
	CaptchaCheckOkMsg             string = "验证码校验通过"
	CaptchaCheckOkCode            int    = 200
	CaptchaCheckFailCode          int    = -400355
	CaptchaCheckFailMsg           string = "验证码校验失败"
)