/**
 * @Author: lenovo
 * @Description:
 * @File:  emai;
 * @Version: 1.0.0
 * @Date: 2023/03/20 22:43
 */

package request

type SendEmail struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckEmailExist struct {
	Email string `json:"email" binding:"required,email"`
}
