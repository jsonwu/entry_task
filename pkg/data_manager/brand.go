package data_manager

import (
	"entry_task/database"
	"entry_task/model"
	"errors"
	logs "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var brandMux sync.Mutex
var brands map[int64]*model.Brand

func startBrandManager(db *database.MyDB) error {
	bm, err := LoadBrand(db)
	if err != nil {
		return err
	}
	brands = bm
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for {
			<-ticker.C
			bMap, err := LoadBrand(db)
			if err != nil {
				logs.Errorf("LoadBrand err %s", err.Error())
				continue
			}
			brandMux.Lock()
			brands = bMap
		}
		ticker.Stop()
	}()
	logs.Infof("start brand success and load brand %+v", brands)
	return nil
}

func LoadBrand(db *database.MyDB) (map[int64]*model.Brand, error) {
	brands, err := db.LoadBrand()
	if err != nil {
		return nil, err
	}
	bm := make(map[int64]*model.Brand)
	for _, v := range brands {
		bm[v.ID] = v
	}
	return bm, nil
}

func GetBrand(id int64) (*model.Brand, error) {
	if v, ok := brands[id]; ok {
		return v, nil
	}
	return nil, errors.New("no exist")
}
