package gserver

import (
	"gorm.io/gorm"
)

/*
   @File: base_dao.go
   @Author: khaosles
   @Time: 2023/4/29 22:33
   @Desc:
*/

type BaseDao[T IModel] interface {
	Gen(*T) error
	GenOrGet(*T) error
	FindByID(string) (*T, error)
	FindByIDs(...string) ([]*T, error)
	FindOne(string, any) (*T, error)
	FindBy(string, ...any) ([]*T, error)
	FindAll() ([]*T, error)
}

type BaseDaoImpl[T IModel] struct {
	db *gorm.DB
}

func (dao *BaseDaoImpl[T]) SetDB(db *gorm.DB) *BaseDaoImpl[T] {
	dao.db = db
	return dao
}

// Create 创建
func (dao BaseDaoImpl[T]) Create(obj *T) (err error) {
	err = dao.db.Create(obj).Error
	return
}

// GetByID 根据id查询
func (dao BaseDaoImpl[T]) GetByID(id string) (obj T, err error) {
	err = dao.db.Where("id = ?", id).First(&obj).Error
	return
}

// GetByIDs 根据多个id查询
func (dao BaseDaoImpl[T]) GetByIDs(ids ...string) (objs []T, err error) {
	err = dao.db.Where("id in (?)", ids).Find(&objs).Error
	return
}

// GetOne 根据字段查询
func (dao BaseDaoImpl[T]) GetOne(filedName string, value any) (obj T, err error) {
	err = dao.db.Where(filedName+" = ?", value).First(&obj).Error
	return
}

// GetBy 获取多条数据
func (dao BaseDaoImpl[T]) GetBy(filedName string, value ...any) (objs []T, err error) {
	err = dao.db.Where(filedName+" = (?)", value).Find(&objs).Error
	return
}

// GetAll 获取全部数据
func (dao BaseDaoImpl[T]) GetAll() (objs []T, err error) {
	err = dao.db.Find(&objs).Error
	return
}

// GetPageBy 分页查询
func (dao BaseDaoImpl[T]) GetPageBy(pagination Pagination, filter string, args ...any) (PagingResult[T], error) {
	var objs []T
	var totalCount int64
	var totalPages int64
	// 计算总记录数
	if err := dao.db.Model(objs).Count(&totalCount).Error; err != nil {
		return PagingResult[T]{}, err
	}
	// 获取总页数
	totalPages = totalCount / int64(pagination.PageSize)
	if totalCount%int64(pagination.PageSize) > 0 {
		totalPages++
	}
	// 当前页
	pageIndex := (pagination.Page - 1) * pagination.PageSize
	err := dao.db.Order(pagination.Sort).
		Offset(pageIndex).
		Limit(pagination.PageSize).
		Where(filter, args).
		Find(&objs).
		Error
	pagingResult := PagingResult[T]{TotalPages: totalPages, TotalCount: totalCount, Objs: objs}
	return pagingResult, err
}

// GetPage 分页查询全部
func (dao BaseDaoImpl[T]) GetPage(pagination Pagination) (PagingResult[T], error) {
	var objs []T
	var totalCount int64
	var totalPages int64
	// 计算总记录数
	if err := dao.db.Model(objs).Count(&totalCount).Error; err != nil {
		return PagingResult[T]{}, err
	}
	// 获取总页数
	totalPages = totalCount / int64(pagination.PageSize)
	if totalCount%int64(pagination.PageSize) > 0 {
		totalPages++
	}
	// 当前页
	pageIndex := (pagination.Page - 1) * pagination.PageSize
	err := dao.db.Order(pagination.Sort).
		Offset(pageIndex).
		Limit(pagination.PageSize).
		Find(&objs).
		Error
	pagingResult := PagingResult[T]{TotalPages: totalPages, TotalCount: totalCount, Objs: objs}
	return pagingResult, err
}

// Delete 根据对象删除
func (dao BaseDaoImpl[T]) Delete(objs ...T) (err error) {
	err = dao.db.Delete(&objs).Error
	return
}

// DeleteByIDs 通过一个或者多个id删除
func (dao BaseDaoImpl[T]) DeleteByIDs(ids ...string) (err error) {
	var objs []T
	err = dao.db.Where("id in (?)", ids).Find(&objs).Error
	if err != nil {
		return
	}
	err = dao.Delete(objs...)
	return
}

// DeleteBy 通过一个或者多个id删除
func (dao BaseDaoImpl[T]) DeleteBy(filedName string, value ...string) (err error) {
	var objs []T
	err = dao.db.Where(filedName+" in (?)", value).Find(&objs).Error
	if err != nil {
		return
	}
	err = dao.Delete(objs...)
	return
}

// Update 通过对象更新
func (dao BaseDaoImpl[T]) Update(obj *T) (err error) {
	err = dao.db.Save(obj).Error
	return err
}

// UpdateBy 通过列更新
func (dao BaseDaoImpl[T]) UpdateBy(obj *T, value map[string]any) (err error) {
	err = dao.db.Model(obj).Updates(value).Error
	return err
}
