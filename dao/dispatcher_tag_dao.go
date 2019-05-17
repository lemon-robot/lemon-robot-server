package dao

import (
	"github.com/jinzhu/gorm"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
)

type DispatcherTagDao struct {
}

func NewDispatcherTagDao() *DispatcherTagDao {
	return &DispatcherTagDao{}
}

func (i *DispatcherTagDao) Save(entity *entity.DispatcherTag) *gorm.DB {
	return db.Db().Save(entity)
}

func (i *DispatcherTagDao) Delete(query interface{}, where ...interface{}) {
	db.Db().Where(query, where).Delete(&entity.DispatcherTag{})
}

func (i *DispatcherTagDao) FirstByExample(example *entity.DispatcherTag) entity.DispatcherTag {
	result := entity.DispatcherTag{}
	db.Db().First(&result, example)
	return result
}

func (i *DispatcherTagDao) FindAllByExample(example *entity.DispatcherTag) []entity.DispatcherTag {
	result := make([]entity.DispatcherTag, 3)
	db.Db().Find(&result, example)
	return result
}
