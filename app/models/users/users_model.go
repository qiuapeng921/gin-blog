package users

import (
	"gin-blog/helpers/pool/grom"
)

var (
	Table = "users"
	Model = grom.GetConn()
)

func FindOne(query interface{})  (Entity, error) {
	var entity Entity
	if err := Model.Where(query).Find(&entity).Error; err != nil {
		return entity, nil
	}
	return entity, nil
}

func FindAll(where ...interface{}) ([]Entity, error) {
	var entity []Entity
	if err := Model.Where(where).Find(&entity).Error; err != nil {
		return nil, nil
	}
	return entity, nil
}

func Insert(data map[string]interface{}) (int64, error) {
	result := Model.Save(data)
	if result.Error != nil {
		return 0, nil
	}
	return result.RowsAffected, nil
}

func Update(where interface{}, data map[string]interface{}) (int64, error) {
	result := Model.Where(where).Update(data)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}
	return result.RowsAffected, nil
}

func Delete(where interface{}) (int64, error) {
	var entity Entity
	result := Model.Where(where).Delete(&entity)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}
	return result.RowsAffected, nil
}
