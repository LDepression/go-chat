/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/03/20 18:49
 */

package request

type Register struct {
	Email      string `json:"email" binding:"required,email"`
	Mobile     string `json:"mobile" binding:"required,mobile"`
	EmailCode  string `json:"email_code" binding:"required"`
	Password   string `json:"password" binding:"gte=3,lte=12"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type LoginType int

var (
	LoginByEmail    LoginType = 1
	LoginByPassword LoginType = 2
)

type Login struct {
	Email     string    `json:"email" binding:"required"`     //邮箱
	Password  string    `json:"password"`                     //密码
	EmailCode string    `json:"email_code"`                   //之间发送的邮箱验证码
	LoginType LoginType `json:"loginType" binding:"required"` //1表示使用邮箱登录, 2表示使用密码登录
}

type ReqModifyPassword struct {
	Password  string `json:"password" binding:"required"`  //修改成的密码
	EmailCode string `json:"emailCode" binding:"required"` //身份认证的验证码
}
