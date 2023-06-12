package g

/*
   @File: service.go
   @Author: khaosles
   @Time: 2023/6/12 01:28
   @Desc:  service 接口继承该接口
*/

type Service[T any] interface {
	Save(entity *T)                              // 保存
	Saves(entities []*T)                         // 批量保存
	DeleteByID(id string)                        // 根据id删除单个
	DeleteByIds(ids ...string)                   // 根据多个id删除
	Update(entity *T)                            // 更新
	FindById(id string) *T                       // 根据id查找
	FindBy(colName string, value any) *T         // 根据某个字段查找唯一值
	FindByIds(ids ...string) []*T                // 根据id查找多个
	FindByCondition(conditions *Conditions) []*T // 根据条件查找
	FindAll() []*T                               // 查找全部
}
