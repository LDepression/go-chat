/**
 * @Author: lenovo
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2023/05/05 20:51
 */

package serve

type SettingType string

const (
	SettingPin        SettingType = "pin"
	SettingShow       SettingType = "show"
	SettingNotDisturb SettingType = "not-disturb"
)

type UpdateSettingType struct {
	Encode     string
	RelationID int64
	SType      SettingType
}

type UpdateNickName struct {
	Encode     string
	RelationID int64
}

type UpdateIsDisturbState struct {
	Encode         string
	RelationID     int64
	SType          SettingType
	IsDisturbState bool
}

type UpdateShowState struct {
	Encode     string
	RelationID int64
	SType      SettingType
	IsShow     bool
}
