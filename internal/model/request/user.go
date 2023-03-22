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
	Password   string `json:"password" binding:"required, gte=6,lte=12"`
	RePassword string `json:"rePassword" binding:"required,equals=password"`
}
