/**
 * @Author: lenovo
 * @Description:
 * @File:  models
 * @Version: 1.0.0
 * @Date: 2023/04/06 19:55
 */

package query

type ApplicationStatus string

const (
	WAITING  ApplicationStatus = "对方验证中"
	ACCEPTED ApplicationStatus = "对方已同意"
	REFUSED  ApplicationStatus = "对方已拒绝"
)
