/**
 * @Author: lenovo
 * @Description:
 * @File:  account
 * @Version: 1.0.0
 * @Date: 2023/03/27 22:30
 */

package automigrate

type Account struct {
	BaseModel
	UserID    uint
	User      User   `gorm:"foreignKey:UserID;references:ID"`
	Name      string `gorm:"type:varchar(255);not null"`
	Signature string `gorm:"type:varchar(255);not null"`
	Avatar    string `gorm:"type:varchar(255);not null" default:"https://cn.bing.com/images/search?view=detailV2&ccid=72WHMpOP&id=56DC04F7B5E8AE1D66E4443C5BABD0ABC06F96C2&thid=OIP.72WHMpOPTPC6jEvW628seQAAAA&mediaurl=https%3a%2f%2ftupian.qqw21.com%2farticle%2fUploadPic%2f2021-1%2f20211722215388977.jpg&exph=400&expw=400&q=%e5%a4%b4%e5%83%8f%e5%9b%be%e7%89%87&simid=608011543894899998&FORM=IRPRST&ck=7C7431204FDC289C05044CF4352EA643&selectedIndex=0"`
	Gender    int    `gorm:"type:int" default:"0" comment:"0表示男性,1表示女性"`
}
