package g

import (
	gstru "github.com/khaosles/gtools2/utils/struct"
)

/*
   @File: base_service.go
   @Author: khaosles
   @Time: 2023/6/12 01:29
   @Desc: service结构体继承该结构体
*/

type BaseService[T any] struct {
	Mpr Mapper[T]
}

func (srv BaseService[T]) Save(entity *T) error {
	return srv.Mpr.Save(entity)
}

func (srv BaseService[T]) Saves(entities []*T) error {
	return srv.Mpr.InsertList(entities)
}

func (srv BaseService[T]) DeleteById(id string) error {
	return srv.Mpr.DeleteById(id)
}

func (srv BaseService[T]) DeleteByIds(ids ...string) error {
	return srv.Mpr.DeleteByIds(ids...)
}

func (srv BaseService[T]) Update(entity *T) error {
	return srv.Mpr.Update(entity)
}

func (srv BaseService[T]) FindById(id string) (*T, error) {
	return srv.Mpr.SelectById(id)
}

func (srv BaseService[T]) FindBy(colName string, value any) (*T, error) {
	var record T
	gstru.SetField(&record, colName, value)
	return srv.Mpr.SelectOne(&record)
}

func (srv BaseService[T]) FindByIds(ids ...string) ([]*T, error) {
	return srv.Mpr.SelectByIds(ids...)
}

func (srv BaseService[T]) FindByCondition(conditions *Conditions) ([]*T, error) {
	return srv.Mpr.SelectByCondition(conditions)
}

func (srv BaseService[T]) FindAll() ([]*T, error) {
	return srv.Mpr.SelectAll()
}
