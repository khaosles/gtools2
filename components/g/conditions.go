package g

import (
	"container/list"

	"gorm.io/gorm"
)

/*
   @File: conditions.go
   @Author: khaosles
   @Time: 2023/6/11 19:53
   @Desc:
*/

type Conditions struct {
	queue *list.List
}

type condition struct {
	query string
	join  int
	value []any
}

const (
	AND = iota
	OR
	ORDER
	SELECT
	GROUP
	JOIN
)

func NewConditions() *Conditions {
	queue := list.New()
	return &Conditions{queue: queue}
}

func (c *Conditions) To(db *gorm.DB) *gorm.DB {
	if c.queue.Len() == 0 {
		return db
	}
	for e := c.queue.Front(); e != nil; e = e.Next() {
		cd := e.Value.(*condition)
		switch cd.join {
		case AND:
			db = db.Where(cd.query, cd.value...)
		case ORDER:
			db = db.Order(cd.query)
		case SELECT:
			db = db.Select(cd.query)
		case GROUP:
			db = db.Group(cd.query)
		case OR:
			db = db.Or(cd.query, cd.value...)
		case JOIN:
			db = db.Joins(cd.query)
		}
	}
	return db
}

// join 连接方式 or and，op 运算符 field 字段  value 字段值
func (c *Conditions) sql(join int, query string, values ...any) *Conditions {
	c.queue.PushBack(&condition{join: join, query: query, value: values})
	return c
}

func (c *Conditions) Order(fields string) *Conditions {
	return c.sql(ORDER, fields)
}

func (c *Conditions) Select(fields string) *Conditions {
	return c.sql(SELECT, fields)
}

func (c *Conditions) Group(fields string) *Conditions {
	return c.sql(GROUP, fields)
}

func (c *Conditions) Joins(joinCondition string) *Conditions {
	return c.sql(JOIN, joinCondition)
}

func (c *Conditions) AndIn(field string, value any) *Conditions {
	return c.sql(AND, "? in (?)", field, value)
}

func (c *Conditions) AndNotIn(field string, value any) *Conditions {
	return c.sql(AND, "? not in (?)", field, value)
}

func (c *Conditions) AndEqualTo(field string, value any) *Conditions {
	return c.sql(AND, "? = ?", field, value)
}

func (c *Conditions) AndLessThan(field string, value any) *Conditions {
	return c.sql(AND, "? < ?", field, value)
}

func (c *Conditions) AndLessThanOrEqualTo(field string, value any) *Conditions {
	return c.sql(AND, "? <= ?", field, value)
}

func (c *Conditions) AndGreaterThan(field string, value any) *Conditions {
	return c.sql(AND, "? > ?", field, value)
}

func (c *Conditions) AndGreaterThanOrEqualTo(field string, value any) *Conditions {
	return c.sql(AND, "? >= ?", field, value)
}

func (c *Conditions) AndNotEqualTo(field string, value any) *Conditions {
	return c.sql(AND, "? <> ?", field, value)
}

func (c *Conditions) AndLike(field string, value any) *Conditions {
	return c.sql(AND, "? like ?", field, value)
}

func (c *Conditions) AndNotLike(field string, value any) *Conditions {
	return c.sql(AND, "? not like ?", field, value)
}

func (c *Conditions) AndILike(field string, value any) *Conditions {
	return c.sql(AND, "? ilike ?", field, value)
}

func (c *Conditions) AndBetween(field string, value1, value2 any) *Conditions {
	return c.sql(AND, "? between ? and ?", field, value1, value2)
}

func (c *Conditions) AndNotBetween(field string, value1, value2 any) *Conditions {
	return c.sql(AND, "? not between ? and ?", field, value1, value2)
}

func (c *Conditions) AndIsNull(field string) *Conditions {
	return c.sql(AND, "? is null", field, nil)
}

func (c *Conditions) AndIsNotNull(field string) *Conditions {
	return c.sql(AND, "? is not null", field)
}

func (c *Conditions) OrIn(field string, value any) *Conditions {
	return c.sql(OR, "? in (?)", field, value)
}

func (c *Conditions) OtNotIn(field string, value any) *Conditions {
	return c.sql(OR, "? not in (?)", field, value)
}

func (c *Conditions) OrEqualTo(field string, value any) *Conditions {
	return c.sql(OR, "? = ?", field, value)
}

func (c *Conditions) OrLessThan(field string, value any) *Conditions {
	return c.sql(OR, "? < ?", field, value)
}

func (c *Conditions) OrLessThanOrEqualTo(field string, value any) *Conditions {
	return c.sql(OR, "? <= ?", field, value)
}

func (c *Conditions) OrGreaterThan(field string, value any) *Conditions {
	return c.sql(OR, "? > ?", field, value)
}

func (c *Conditions) OrGreaterThanOrEqualTo(field string, value any) *Conditions {
	return c.sql(OR, "? >= ?", field, value)
}

func (c *Conditions) OrNotEqualTo(field string, value any) *Conditions {
	return c.sql(OR, "? <> ?", field, value)
}

func (c *Conditions) OrLike(field string, value any) *Conditions {
	return c.sql(OR, "? like ?", field, value)
}

func (c *Conditions) OrNotLike(field string, value any) *Conditions {
	return c.sql(OR, "? not like ?", field, value)
}

func (c *Conditions) OrILike(field string, value any) *Conditions {
	return c.sql(OR, "? ilike ?", field, value)
}

func (c *Conditions) OrIsNull(field string) *Conditions {
	return c.sql(OR, "? is null", field)
}

func (c *Conditions) OrIsNotNull(field string) *Conditions {
	return c.sql(OR, "? is not null", field)
}

func (c *Conditions) OrBetween(field string, value1, value2 any) *Conditions {
	return c.sql(OR, "? between ? and ?", field, value1, value2)
}

func (c *Conditions) OrNotBetween(field string, value1, value2 any) *Conditions {
	return c.sql(OR, "? not between ? and ?", field, value1, value2)
}
