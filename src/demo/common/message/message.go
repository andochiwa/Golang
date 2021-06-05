package message

type (
	Message struct {
		Type string `json:"type"` // 消息类型
		Data string `json:"data"` // 消息内容
	}

	LoginMessage struct {
		UserId   int    `json:"user_id"`   // 用户id
		UserPwd  string `json:"user_pwd"`  // 用户密码
		UserName string `json:"user_name"` // 用户名
	}

	LoginResult struct {
		Code  int    `json:"code"`  // 返回状态码 200成功 444未注册
		Error string `json:"error"` // 返回错误信息
	}
)

// 定义一些消息类型

const (
	LoginMessageType = "LoginMessage"
	LoginResultType  = "LoginResultType"
)
