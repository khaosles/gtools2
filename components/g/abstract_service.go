package g

import (
	glog "github.com/khaosles/gtools2/core/log"
	gstru "github.com/khaosles/gtools2/g/struct"
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

func (srv AbstractService[T]) Save(entity *T) {
	err := srv.Mpr.Insert(entity)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) Saves(entities []*T) {
	err := srv.Mpr.InsertList(entities)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) DeleteByID(id string) {
	err := srv.Mpr.DeleteByID(id)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) DeleteByIds(ids ...string) {
	err := srv.Mpr.DeleteByIDs(ids...)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) Update(entity *T) {
	err := srv.Mpr.Update(entity)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) FindById(id string) *T {
	entity, err := srv.Mpr.SelectByID(id)
	if err != nil {
		glog.Error(err)
	}
	return entity
}

func (srv AbstractService[T]) FindBy(colName string, value any) *T {
	var entity T
	gstru.SetField(&entity, colName, value)
	obj, err := srv.Mpr.SelectOne(&entity)
	if err != nil {
		glog.Error(err)
	}
	return obj
}

func (srv AbstractService[T]) FindByIds(ids ...string) []*T {
	entities, err := srv.Mpr.SelectByIDs(ids...)
	if err != nil {
		glog.Error(err)
	}
	return entities
}

func (srv AbstractService[T]) FindByCondition(conditions *Conditions) []*T {
	entities, err := srv.Mpr.SelectByCondition(conditions)
	if err != nil {
		glog.Error(err)
	}
	return entities
}

func (srv AbstractService[T]) FindAll() []*T {
	entities, err := srv.Mpr.SelectAll()
	if err != nil {
		glog.Error(err)
	}
	return entities
}
