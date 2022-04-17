package feed

import (
	"entry_task/database"
	"entry_task/model"
	"errors"
	"github.com/sirupsen/logrus"
	"math/rand"
)

var productIDPool []string
var products map[string]*model.Product

type strategy interface {
	GetFeed(uerName string, num int) ([]string, error)
}

var feedStategy strategy

func Init(db *database.MyDB) error {
	p, err := db.LoadProducts(0, 1000)
	if err != nil {
		return err
	}
	products = make(map[string]*model.Product)
	for _, v := range p {
		productIDPool = append(productIDPool, v.ProductID)
		products[v.ProductID] = v
	}
	logrus.Infof("load product num %d", len(productIDPool))
	feedStategy = RandomShuffle{}
	//feedStategy = No{}
	return nil
}

const recommedPoolSize int = 100

func GetFeed(userName string, num int) ([]*model.Product, error) {
	if num > recommedPoolSize {
		return nil, errors.New("over size")
	}
	ids, err := feedStategy.GetFeed(userName, num)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Product, 0, num)
	for i := 0; i < len(ids) && i < num; i++ {
		p := products[ids[i]]
		result = append(result, p)
	}
	return result, nil
}

type RandomShuffle struct {
}

func (r RandomShuffle) GetFeed(uerName string, num int) ([]string, error) {
	var recommedIDPool []string
	if len(productIDPool) <= recommedPoolSize {
		recommedIDPool = productIDPool[:]
	} else {
		begin := int(rand.Int31n(int32(len(productIDPool) - recommedPoolSize)))
		recommedIDPool = productIDPool[begin : begin+recommedPoolSize+1]
	}
	rand.Shuffle(len(recommedIDPool), func(i, j int) {
		recommedIDPool[i], recommedIDPool[j] = recommedIDPool[j], recommedIDPool[i]
	})
	return recommedIDPool[0:num], nil
}

type Random struct {
}

func (r Random) GetFeed(uerName string, num int) ([]string, error) {
	var recommedIDPool []string
	if len(productIDPool) <= recommedPoolSize {
		recommedIDPool = productIDPool[:]
	} else {
		begin := int(rand.Int31n(int32(len(productIDPool) - recommedPoolSize)))
		recommedIDPool = productIDPool[begin : begin+recommedPoolSize+1]
	}
	return recommedIDPool, nil
}

type No struct {
}

func (r No) GetFeed(uerName string, num int) ([]string, error) {
	return productIDPool[0:num], nil
}
