package g

import (
	glog "github.com/khaosles/gtools2/core/log"
	gstru "github.com/khaosles/gtools2/g/struct"
)

/*
   @File: abstract_service.go
   @Author: khaosles
   @Time: 2023/6/12 01:29
   @Desc:
*/

type AbstractService[T any] struct {
	mpr Mapper[T]
}

func (srv AbstractService[T]) Save(entity *T) {
	err := srv.mpr.Insert(entity)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) DeleteByID(id string) {
	err := srv.mpr.DeleteByID(id)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) deleteByIds(ids ...string) {
	err := srv.mpr.DeleteByIDs(ids...)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) update(entity *T) {
	err := srv.mpr.Update(entity)
	if err != nil {
		glog.Error(err)
	}
}

func (srv AbstractService[T]) findById(id string) *T {
	entity, err := srv.mpr.SelectByID(id)
	if err != nil {
		glog.Error(err)
	}
	return entity
}

func (srv AbstractService[T]) findBy(colName string, value any) *T {
	var entity T
	gstru.SetField(&entity, colName, value)
	obj, err := srv.mpr.SelectOne(&entity)
	if err != nil {
		glog.Error(err)
	}
	return obj
}

func (srv AbstractService[T]) findByIds(ids ...string) []*T {
	entities, err := srv.mpr.SelectByIDs(ids...)
	if err != nil {
		glog.Error(err)
	}
	return entities
}

func (srv AbstractService[T]) findByCondition(conditions *Conditions) []*T {
	entities, err := srv.mpr.SelectByCondition(conditions)
	if err != nil {
		glog.Error(err)
	}
	return entities
}

func (srv AbstractService[T]) findAll() []*T {
	entities, err := srv.mpr.SelectAll()
	if err != nil {
		glog.Error(err)
	}
	return entities
}
