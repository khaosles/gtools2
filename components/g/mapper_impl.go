package g

import (
	gstru "github.com/khaosles/gtools2/g/struct"
	"gorm.io/gorm"
)

/*
   @File: mapper_impl.go
   @Author: khaosles
   @Time: 2023/6/11 19:28
   @Desc:
*/

type MapperImpl[T any] struct {
	DB *gorm.DB
}

func (mapper MapperImpl[T]) Insert(record *T) error {
	return mapper.DB.Create(record).Error
}

func (mapper MapperImpl[T]) InsertList(records []*T) error {
	return mapper.DB.CreateInBatches(records, len(records)).Error
}

func (mapper MapperImpl[T]) InsertBatch(records []*T, batch int) error {
	if batch < 1 {
		batch = len(records)
	}
	return mapper.DB.CreateInBatches(records, batch).Error
}

func (mapper MapperImpl[T]) InsertOrSelect(record *T) error {
	return mapper.DB.FirstOrCreate(record).Error
}

func (mapper MapperImpl[T]) Delete(record *T) error {
	return mapper.DB.Delete(record).Error
}

func (mapper MapperImpl[T]) DeleteHard(record *T) error {
	return mapper.DB.Unscoped().Delete(record).Error
}

func (mapper MapperImpl[T]) DeleteByID(id string) error {
	return mapper.DB.Delete(new(T), "id = ?", id).Error
}

func (mapper MapperImpl[T]) DeleteHardByID(id string) error {
	return mapper.DB.Unscoped().Delete(new(T), "id = ?", id).Error
}

func (mapper MapperImpl[T]) DeleteByIDs(ids ...string) error {
	return mapper.DB.Delete(new(T), "id in (?)", ids).Error
}

func (mapper MapperImpl[T]) DeleteHardByIDs(ids ...string) error {
	return mapper.DB.Unscoped().Delete(new(T), "id in (?)", ids).Error
}

func (mapper MapperImpl[T]) DeleteByCondition(conditions *Conditions) error {
	return conditions.To(mapper.DB).Delete(new(T)).Error
}

func (mapper MapperImpl[T]) DeleteHardByCondition(conditions *Conditions) error {
	return conditions.To(mapper.DB).Unscoped().Delete(new(T)).Error
}

func (mapper MapperImpl[T]) Update(record *T) error {
	return mapper.DB.Save(record).Error
}

func (mapper MapperImpl[T]) UpdateSelective(record *T, values any) error {
	return mapper.DB.Model(record).Updates(values).Error
}

func (mapper MapperImpl[T]) UpdateByCondition(record *T, conditions *Conditions) error {
	return conditions.To(mapper.DB).Model(new(T)).Updates(gstru.StructToMapInterface(record)).Error
}

func (mapper MapperImpl[T]) UpdateSelectiveByCondition(record *T, conditions *Conditions) error {
	return conditions.To(mapper.DB).Model(new(T)).Updates(record).Error
}

func (mapper MapperImpl[T]) SelectByID(id string) (*T, error) {
	var record T
	err := mapper.DB.Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (mapper MapperImpl[T]) SelectByIDs(ids ...string) ([]*T, error) {
	var records []*T
	err := mapper.DB.Where("id = (?)", ids).First(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (mapper MapperImpl[T]) SelectOne(record *T) (*T, error) {
	var entity T
	err := mapper.DB.Where(record).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (mapper MapperImpl[T]) Select(record *T) ([]*T, error) {
	var entities []*T
	err := mapper.DB.Where(record).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mapper MapperImpl[T]) SelectAll() ([]*T, error) {
	var entities []*T
	err := mapper.DB.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mapper MapperImpl[T]) SelectCount(record *T) (int64, error) {
	var count int64
	err := mapper.DB.Where(record).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (mapper MapperImpl[T]) SelectByCondition(conditions *Conditions) ([]*T, error) {
	var entities []*T
	err := conditions.To(mapper.DB).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mapper MapperImpl[T]) SelectCountByCondition(conditions *Conditions) (int64, error) {
	var count int64
	err := conditions.To(mapper.DB).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (mapper MapperImpl[T]) SelectDistinct(conditions *Conditions) ([]*T, error) {
	var entities []*T
	err := conditions.To(mapper.DB).Distinct().Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mapper MapperImpl[T]) SelectByJoinAndCondition(joinCondition string, conditions *Conditions) ([]*T, error) {
	var entities []*T
	err := conditions.To(mapper.DB.Joins(joinCondition)).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (mapper MapperImpl[T]) SelectPage(currentPage, pageSize int, sort string) (*Pagination[T], error) {
	var pagination Pagination[T]
	var entities []*T
	var totalCount int64
	var totalPages int64
	// 计算总记录数
	if err := mapper.DB.Model(entities).Count(&totalCount).Error; err != nil {
		return nil, err
	}
	// 获取总页数
	totalPages = totalCount / int64(pageSize)
	if totalCount%int64(pageSize) > 0 {
		totalPages++
	}
	// 当前页
	pageIndex := (currentPage - 1) * pageSize
	err := mapper.DB.Order(sort).
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

func (mapper MapperImpl[T]) SelectPageByCondition(currentPage, pageSize int, sort string, conditions Conditions) (*Pagination[T], error) {
	db := conditions.To(mapper.DB)
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

func (mapper MapperImpl[T]) SelectPageByJoinAndCondition(joinCondition string, currentPage, pageSize int, sort string, conditions Conditions) (*Pagination[T], error) {
	db := conditions.To(mapper.DB.Joins(joinCondition))
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

func (mapper MapperImpl[T]) Exist(record *T) (bool, error) {
	count, err := mapper.SelectCount(record)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
