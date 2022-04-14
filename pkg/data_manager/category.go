package data_manager

import (
	"entry_task/database"
	"entry_task/model"
	"errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var categoryMux sync.Mutex
var categorys map[int64]model.Category

func StartCategoryManager(db *database.MyDB) error {
	cm, err := LoadCategory(db)
	if err != nil {
		return err
	}
	categorys = cm
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for {
			<-ticker.C
			cMap, err := LoadCategory(db)
			if err != nil {
				logrus.Errorf("LoadBrand err %s", err.Error())
				continue
			}
			categoryMux.Lock()
			categorys = cMap
		}
		ticker.Stop()
	}()
	return nil
}

func LoadCategory(db *database.MyDB) (map[int64]model.Category, error) {
	category, err := db.LoadCategory()
	if err != nil {
		return nil, err
	}
	bMap := make(map[int64]model.Category)
	for _, v := range category {
		bMap[v.ID] = v
	}
	return bMap, nil
}
func GetCategory(id int64) (model.Category, error) {
	if v, ok := categorys[id]; ok {
		return v, nil
	}
	return model.Category{}, errors.New("no exist")
}
