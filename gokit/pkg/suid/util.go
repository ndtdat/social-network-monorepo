package suid

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/concurrentset"
	"strconv"
	"sync"
	"time"
)

func NextID() (uint64, error) {
	return service.NextID()
}

func New() uint64 {
	id, err := service.NextID()
	if err != nil {
		panic(fmt.Sprintf("cannot generate sonyflake id due to %v", err))
	}

	return id
}

func NewStr() string {
	id, err := service.NextID()
	if err != nil {
		panic(fmt.Sprintf("cannot generate sonyflake id due to %v", err))
	}

	return strconv.FormatUint(id, 10)
}

//nolint:forbidigo
func TestUniqueness(nThread, numIDPerThread int) {
	set := concurrentset.New[uint64](512)
	var wg sync.WaitGroup
	for i := 0; i < nThread; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < numIDPerThread; j++ {
				id := New()
				if set.Contains(id) {
					//nolint:forbidigo,gosimple
					fmt.Println(fmt.Sprintf("Failed due to duplicated id %d", id))

					break
				}

				set.Add(id)
			}
		}()
	}

	wg.Wait()

	total := nThread * numIDPerThread

	//nolint:gosimple
	fmt.Println(fmt.Sprintf("Finished generating %d ids", total))
	//nolint:gosimple
	fmt.Println(fmt.Sprintf("Number of unique ids: %d", set.Count()))
}

// suid.TestRate(1, 10_000_000).
func TestRate(nThread, numIDPerThread int) {
	var wg sync.WaitGroup
	start := time.Now().UnixMicro()
	for i := 0; i < nThread; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < numIDPerThread; j++ {
				_ = New()
			}
		}()
	}

	wg.Wait()
	elapsed := time.Now().UnixMicro() - start
	total := nThread * numIDPerThread
	rate := float64(total) / float64(elapsed)

	//nolint:gosimple,forbidigo
	fmt.Println(fmt.Sprintf("Finished generating %d id in %d microseconds", total, elapsed))
	//nolint:gosimple,forbidigo
	fmt.Println(fmt.Sprintf("Rate: %f per microsecond", rate))
}
