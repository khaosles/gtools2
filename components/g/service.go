package g

/*
   @File: base_service.go
   @Author: khaosles
   @Time: 2023/6/12 01:28
   @Desc:
*/

type Service[T any] interface {
	Save(entity *T)
	DeleteByID(id string)
	deleteByIds(ids string)
	update(entity *T)
	findById(id string) *T
	findBy(colName string, value any) *T
	findByIds(ids ...string) []*T
	findByCondition(conditions *Conditions) []*T
	findAll() []*T
}
