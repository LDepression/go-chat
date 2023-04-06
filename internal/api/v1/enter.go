/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/03/20 22:54
 */

package v1

type group struct {
	User        user
	Email       email
	Account     account
	Application application
}

var Group = new(group)
