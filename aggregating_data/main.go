package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	start := time.Now()
// 	userName := fetchUser()
// 	likes := fetchUserLikes(userName)
// 	match := fetchUserMatch(userName)
// 	fmt.Println(likes)
// 	fmt.Println(match)
// 	fmt.Println(time.Since(start))
// } => { 11 ANNA 3.0019714s }

// Áp goroutine 1

// func main() {

// 	start := time.Now()
// 	userName := fetchUser()
// 	likes := make(chan int, 1)
// 	match := make(chan string, 1)

// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go func(n string) {
// 		defer wg.Done()
// 		likes <- fetchUserLikes(n)
// 	}(userName)

// 	go func(n string) {
// 		defer wg.Done()
// 		match <- fetchUserMatch(n)
// 	}(userName)

// 	wg.Wait()

// 	fmt.Println(<-likes)
// 	fmt.Println(<-match)
// 	fmt.Println(time.Since(start))
// } =>{11
// ANNA
// 2.0014738s}

func main() {
	start := time.Now()
	userName := fetchUser()

	likesChan := make(chan int)
	matchChan := make(chan string)

	// Dùng sync.WaitGroup để đợi goroutines hoàn thành
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine lấy likes
	go func() {
		defer wg.Done()
		likesChan <- fetchUserLikes(userName)
	}()

	// Goroutine lấy match
	go func() {
		defer wg.Done()
		matchChan <- fetchUserMatch(userName)
	}()

	// Goroutine đóng channel sau khi tất cả hoàn thành
	go func() {
		wg.Wait()
		close(likesChan)
		close(matchChan)
	}()

	// Nhận dữ liệu không cần WaitGroup.Wait()
	var likes int
	var match string
	for i := 0; i < 2; i++ {
		select {
		case likes = <-likesChan:
		case match = <-matchChan:
		}
	}

	fmt.Println(likes)
	fmt.Println(match)
	fmt.Println(time.Since(start))
}

func fetchUser() string {

	time.Sleep(time.Second * 1)

	return "BOB"

}

func fetchUserLikes(userName string) int {
	time.Sleep(time.Second * 1)

	return 11
}

func fetchUserMatch(username string) string {

	time.Sleep(time.Second * 1)
	return "ANNA"
}
