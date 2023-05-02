/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/03/20 18:40
 */

package routing

type group struct {
	User        user
	Email       email
	Account     account
	Application application
	Setting     setting
}

var Group = new(group)
