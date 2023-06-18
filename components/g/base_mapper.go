package g

import (
	"gorm.io/gorm"

	gstru "github.com/khaosles/gtools2/utils/struct"
)

/*
   @File: abstract_mapper.go
   @Author: khaosles
   @Time: 2023/6/11 19:28
   @Desc: mapper结构体继承该结构体
*/

type BaseMapper[T any] struct {
	DB *gorm.DB
}

func (mpr BaseMapper[T]) Save(record *T) error {
	return mpr.DB.Save(record).Error
}

func (mpr BaseMapper[T]) Insert(record *T) error {
	return mpr.DB.Create(record).Error
}

func (mpr BaseMapper[T]) InsertList(records []*T) error {
	return mpr.DB.CreateInBatches(records, len(records)).Error
}

func (mpr BaseMapper[T]) InsertBatch(records []*T, batch int) error {
	if batch < 1 {
		batch = len(records)
	}
	return mpr.DB.CreateInBatches(records, batch).Error
}

func (mpr BaseMapper[T]) InsertOrSelect(record *T) error {
	return mpr.DB.FirstOrCreate(record).Error
}

func (mpr BaseMapper[T]) Delete(record *T) error {
	return mpr.DB.Delete(record).Error
}

func (mpr BaseMapper[T]) DeleteHard(record *T) error {
	return mpr.DB.Unscoped().Delete(record).Error
}

func (mpr BaseMapper[T]) DeleteById(id string) error {
	return mpr.DB.Delete(new(T), "id = ?", id).Error
}

func (mpr BaseMapper[T]) DeleteHardById(id string) error {
	return mpr.DB.Unscoped().Delete(new(T), "id = ?", id).Error
}

func (mpr BaseMapper[T]) DeleteByIds(ids ...string) error {
	return mpr.DB.Delete(new(T), "id in (?)", ids).Error
}

func (mpr BaseMapper[T]) DeleteHardByIds(ids ...string) error {
	return mpr.DB.Unscoped().Delete(new(T), "id in (?)", ids).Error
}

func (mpr BaseMapper[T]) DeleteByCondition(conditions *Conditions) error {
	return conditions.To(mpr.DB).Delete(new(T)).Error
}

func (mpr BaseMapper[T]) DeleteHardByCondition(conditions *Conditions) error {
	return conditions.To(mpr.DB).Unscoped().Delete(new(T)).Error
}

func (mpr BaseMapper[T]) Update(record *T) error {
	return mpr.DB.Save(record).Error
}

func (mpr BaseMapper[T]) UpdateSelective(record *T, values any) error {
	return mpr.DB.Model(record).Updates(values).Error
}

func (mpr BaseMapper[T]) UpdateByCondition(record *T, conditions *Conditions) error {
	return conditions.To(mpr.DB).Model(new(T)).Updates(gstru.StructToMapInterface(record)).Error
}

func (mpr BaseMapper[T]) UpdateSelectiveByCondition(record *T, conditions *Conditions) error {
	return conditions.To(mpr.DB).Model(new(T)).Updates(record).Error
}

func (mpr BaseMapper[T]) SelectById(id string) (*T, error) {
	var record T
	err := mpr.DB.Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (mpr BaseMapper[T]) SelectByIds(ids ...string) ([]*T, error) {
	var records []*T
	err := mpr.DB.Where("id = (?)", ids).First(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (mpr BaseMapper[T]) SelectOne(record *T) (*T, error) {
	var entity T
	err := mpr.DB.Where(record).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (mpr BaseMapper[T]) SelectOneByConditions(record *T, conditions *Conditions) (*T, error) {
	var entity T
	err := conditions.To(mpr.DB).Where(record).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (mpr BaseMapper[T]) Select(record *T) ([]*T, error) {
	var entities []*T
	err := mpr.DB.Where(record).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mpr BaseMapper[T]) SelectAll() ([]*T, error) {
	var entities []*T
	err := mpr.DB.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mpr BaseMapper[T]) SelectCount(record *T) (int64, error) {
	var count int64
	err := mpr.DB.Model(new(T)).Where(record).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (mpr BaseMapper[T]) SelectByCondition(conditions *Conditions) ([]*T, error) {
	var entities []*T
	err := conditions.To(mpr.DB).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mpr BaseMapper[T]) SelectCountByCondition(conditions *Conditions) (int64, error) {
	var count int64
	err := conditions.To(mpr.DB).Model(new(T)).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (mpr BaseMapper[T]) SelectDistinct(conditions *Conditions) ([]*T, error) {
	var entities []*T
	err := conditions.To(mpr.DB).Distinct().Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mpr BaseMapper[T]) SelectPage(currentPage, pageSize int, sort string) (*Pagination[T], error) {
	return mpr.SelectPageByCondition(currentPage, pageSize, sort, NewConditions())
}

func (mpr BaseMapper[T]) SelectPageByCondition(currentPage, pageSize int, sort string, conditions *Conditions) (*Pagination[T], error) {
	db := conditions.To(mpr.DB)
	var pagination Pagination[T]
	var entities []*T
	var totalCount int64
	var totalPages int64
	// 计算总记录数
	if err := db.Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 获取总页数
	totalPages = totalCount / int64(pageSize)
	if totalCount%int64(pageSize) > 0 {
		totalPages++
	}
	// 当前页
	pageIndex := (currentPage - 1) * pageSize
	err := db.Order(sort).
		Offset(pageIndex).
		Limit(pageSize).
		Find(&entities).
		Error
	if err != nil {
		return nil, err
	}
	pagination.CurrentPage = currentPage
	pagination.TotalCount = totalCount
	pagination.PageSize = pageSize
	pagination.TotalPages = totalPages
	pagination.DataCollection = entities
	return &pagination, nil
}

func (mpr BaseMapper[T]) Exist(record *T) (bool, error) {
	count, err := mpr.SelectCount(record)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
