package cmap

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sync"
	"time"
)

//nolint:forbidigo
func BenchmarkReadWrite(nShard int, nElement int, nThread int) {
	m := New[int, int](nShard)

	var (
		wg    sync.WaitGroup
		start = time.Now().UnixMilli()
	)

	nElementPerThread := nElement / nThread

	for i := 0; i <= nThread; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			for j := index * nElementPerThread; j < (index+1)*nElementPerThread; j++ {
				m.Set(j, j)
			}
		}(i)
	}

	wg.Wait()
	p := message.NewPrinter(language.English)
	//nolint:forbidigo
	fmt.Println(
		p.Sprintf(
			"It takes %d (ms) with nShard=%d, nElement=%d, nThread=%d for WRITE",
			time.Now().UnixMilli()-start, nShard, nElement, nThread,
		),
	)

	start = time.Now().UnixMilli()
	for i := 0; i <= nThread; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			for j := index * nElementPerThread; j < (index+1)*nElementPerThread; j++ {
				v, _ := m.Get(j)
				if v != j {
					fmt.Println("Wrong")
				}
			}
		}(i)
	}

	wg.Wait()
	p = message.NewPrinter(language.English)
	//nolint:forbidigo
	fmt.Println(
		p.Sprintf(
			"It takes %d (ms) with nShard=%d, nElement=%d, nThread=%d for READ",
			time.Now().UnixMilli()-start, nShard, nElement, nThread,
		),
	)
}

//nolint:forbidigo
func BenchmarkSyncReadWrite(nElement int, nThread int) {
	m := sync.Map{}

	var (
		wg    sync.WaitGroup
		start = time.Now().UnixMilli()
	)

	nElementPerThread := nElement / nThread

	for i := 0; i <= nThread; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			for j := index * nElementPerThread; j < (index+1)*nElementPerThread; j++ {
				m.Store(j, j)
			}
		}(i)
	}

	wg.Wait()
	p := message.NewPrinter(language.English)
	//nolint:forbidigo
	fmt.Println(
		p.Sprintf(
			"It takes %d (ms) with nElement=%d, nThread=%d for WRITE",
			time.Now().UnixMilli()-start, nElement, nThread,
		),
	)

	start = time.Now().UnixMilli()
	for i := 0; i <= nThread; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			for j := index * nElementPerThread; j < (index+1)*nElementPerThread; j++ {
				v, _ := m.Load(j)
				if v != j {
					fmt.Println("Wrong")
				}
			}
		}(i)
	}

	wg.Wait()
	p = message.NewPrinter(language.English)
	//nolint:forbidigo
	fmt.Println(
		p.Sprintf(
			"It takes %d (ms) with nElement=%d, nThread=%d for READ",
			time.Now().UnixMilli()-start, nElement, nThread,
		),
	)
}
