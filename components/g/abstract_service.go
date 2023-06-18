package g

import (
	gstru "github.com/khaosles/gtools2/utils/struct"
)

/*
   @File: abstract_service.go
   @Author: khaosles
   @Time: 2023/6/12 01:29
   @Desc: service结构体继承该结构体
*/

type AbstractService[T any] struct {
	Mpr Mapper[T]
}

func (srv AbstractService[T]) Save(entity *T) error {
	return srv.Mpr.Save(entity)
}

func (srv AbstractService[T]) Saves(entities []*T) error {
	return srv.Mpr.InsertList(entities)
}

func (srv AbstractService[T]) DeleteById(id string) error {
	return srv.Mpr.DeleteById(id)
}

func (srv AbstractService[T]) DeleteByIds(ids ...string) error {
	return srv.Mpr.DeleteByIds(ids...)
}

func (srv AbstractService[T]) Update(entity *T) error {
	return srv.Mpr.Update(entity)
}

func (srv AbstractService[T]) FindById(id string) (*T, error) {
	return srv.Mpr.SelectById(id)
}

func (srv AbstractService[T]) FindBy(colName string, value any) (*T, error) {
	var record T
	gstru.SetField(&record, colName, value)
	return srv.Mpr.SelectOne(&record)
}

func (srv AbstractService[T]) FindByIds(ids ...string) ([]*T, error) {
	return srv.Mpr.SelectByIds(ids...)
}

func (srv AbstractService[T]) FindByCondition(conditions *Conditions) ([]*T, error) {
	return srv.Mpr.SelectByCondition(conditions)
}

func (srv AbstractService[T]) FindAll() ([]*T, error) {
	return srv.Mpr.SelectAll()
}
