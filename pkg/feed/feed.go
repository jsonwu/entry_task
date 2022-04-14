package feed

import (
	"entry_task/database"
	"entry_task/model"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var productIDPool []string
var productMap map[string]model.Product

func Init(db *database.MyDB) error {
	p, err := db.LoadProducts(0, 1000)
	if err != nil {
		return err
	}
	productMap = make(map[string]model.Product)
	for _, v := range p {
		productIDPool = append(productIDPool, v.ProductID)
		productMap[v.ProductID] = v
	}
	logrus.Infof("load product num %d", len(productIDPool))
	return nil
}

func GetFeed(userName string, num int) ([]model.Product, error) {
	resultIDPool := productIDPool[:]
	if len(productIDPool) < num {
		resultIDPool = productIDPool
	} else {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(resultIDPool), func(i, j int) {
			resultIDPool[i], resultIDPool[j] = resultIDPool[j], resultIDPool[i]
		})
	}
	result := make([]model.Product, 0, num)
	for i := 0; i < len(resultIDPool) && i < num; i++ {
		result = append(result, productMap[resultIDPool[i]])
	}
	return result, nil
}
