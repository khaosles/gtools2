package g

import (
	"time"

	"gorm.io/gorm"

	"github.com/khaosles/gtools2/g/gen"
)

/*
   @File: model.go
   @Author: khaosles
   @Time: 2023/6/11 11:10
   @Desc: model 继承该类，则拥有常见基础属性，以及自动生成无分隔符的uuid
*/

type Model struct {
	ID         string         `json:"id" gorm:"type:varchar(32);comment:主键"`                   // 主键ID
	CreateTime time.Time      `json:"createTime,omitempty" gorm:"autoCreateTime;comment:创建时间"` // 创建时间
	CreateBy   string         `json:"createBy,omitempty" gorm:"type:varchar(100);comment:创建人"` // 创建人
	UpdateTime time.Time      `json:"updateTime,omitempty" gorm:"autoUpdateTime;comment:更新时间"` // 更新时间
	UpdateBy   string         `json:"updateBy,omitempty" gorm:"type:varchar(100);comment:更新人"` // 更新人
	DeleteTime gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                             // 删除标记
	Remarks    string         `json:"remarks" gorm:"default:null;comment:备注"`                  // 备注
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.ID = gen.UuidNoSeparator()
	return nil
}
