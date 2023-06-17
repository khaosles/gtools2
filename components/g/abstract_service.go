package g

import (
	glog "github.com/khaosles/gtools2/core/log"
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

func (srv AbstractService[T]) Save(entity *T) {
	err := srv.Mpr.Save(entity)
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

func (srv AbstractService[T]) DeleteById(id string) {
	err := srv.Mpr.DeleteById(id)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) DeleteByIds(ids ...string) {
	err := srv.Mpr.DeleteByIds(ids...)
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
	entity, err := srv.Mpr.SelectById(id)
	if err != nil {
		glog.Error(err)
	}
	return entity
}

func (srv AbstractService[T]) FindBy(colName string, value any) *T {
	var record T
	gstru.SetField(&record, colName, value)
	entity, err := srv.Mpr.SelectOne(&record)
	if err != nil {
		glog.Error(err)
	}
	return entity
}

func (srv AbstractService[T]) FindByIds(ids ...string) []*T {
	entities, err := srv.Mpr.SelectByIds(ids...)
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
