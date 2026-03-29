/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package model

type SvipPrivilege struct {
	ID        uint      `json:"id"`
	Vid       int       `json:"sid"`
	Pid       int       `json:"pid"`
	Privilege Privilege `gorm:"foreignKey:Pid;" json:"privilege"`
}
