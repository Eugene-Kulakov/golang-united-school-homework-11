package batch

import (
	//"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	//var wg sync.WaitGroup
	sem := make(chan struct{}, pool)
	res_chan := make(chan user, n)
	for i := int64(0); i != n; i++ {
		//wg.Add(1)
		sem <- struct{}{}
		go func(j int64) {
			//defer wg.Done()
			user := getOne(j)
			res_chan <- user
			<-sem
		}(i)
	}
	for i := int64(0); i != n; i++ {
		u := <-res_chan
		res = append(res, u)
	}
	return
}
