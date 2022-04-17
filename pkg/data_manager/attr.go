package data_manager

import (
	"entry_task/database"
	"entry_task/model"
	"errors"
	logs "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var attrMux sync.Mutex
var attrs map[string]*model.Attr

func startAttrManager(db *database.MyDB) error {
	am, err := LoadAttr(db)
	if err != nil {
		return err
	}
	attrs = am
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for {
			<-ticker.C
			am, err := LoadAttr(db)
			if err != nil {
				logs.Errorf("LoadBrand err %s", err.Error())
				continue
			}
			logs.Infof("reload attr len %d", len(attrs))
			attrMux.Lock()
			attrs = am
		}
		ticker.Stop()
	}()
	logs.Infof("start attr success and load attr %+v", attrs)
	return nil
}

func LoadAttr(db *database.MyDB) (map[string]*model.Attr, error) {
	attrs, err := db.LoadAttr()
	if err != nil {
		return nil, err
	}
	am := make(map[string]*model.Attr)
	for _, v := range attrs {
		am[v.Name] = v
	}
	return am, nil
}

func GetAttr(attrName string) (*model.Attr, error) {
	if v, ok := attrs[attrName]; ok {
		return v, nil
	}
	return nil, errors.New("no exist")
}
