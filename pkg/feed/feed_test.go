package feed

import (
	"entry_task/config"
	"entry_task/database"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFeed(b *testing.B) {
	db, err := database.NewMyDB(&config.Mysql{})
	if err != nil {
		panic(fmt.Errorf("database.NewMyDB err %s", err.Error()))
	}
	Init(db)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		_, err := GetFeed("", 10)
		if err != nil {
			panic(err)
		}
	}
}
