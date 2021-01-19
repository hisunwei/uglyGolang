package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	start := time.Now()
	max := 10000 * 10000
	fmt.Printf("try with max:%d\n", max)
	knownPrimes := make([]int, 0)
	knownPrimes = append(knownPrimes, 2, 3, 5, 7, 11)


	next := knownPrimes[len(knownPrimes) -1 ] + 1
	//
	//knownPrimes1 := SingleThreadPrimes(next, max - next, knownPrimes)
	//fmt.Printf(">>>> 1thread result:%d cost:%v first10:%v last10:%v \n", len(knownPrimes1), time.Since(start), knownPrimes1[:10], knownPrimes1[len(knownPrimes1)-10:])


	start = time.Now()
	for true {
		offset := next * next - next
		isEnd := false
		if next + offset > max {
			offset = max - next
			isEnd = true
		}
		//fmt.Printf("start round p:%d offset:%d size:%d", next, offset, len(knownPrimes))
		knownPrimes = MultiExecute(next, offset, knownPrimes, 8)
		//fmt.Printf("end round p:%d offset:%d size:%d\n", next, offset, len(knownPrimes))
		if isEnd {
			break
		}
		next = next + offset
	}
	fmt.Printf(">>>> result:%d cost:%v first10:%v last10:%v \n", len(knownPrimes), time.Since(start), knownPrimes[:10], knownPrimes[len(knownPrimes)-10:])
}

func MultiExecute(p int, offset int, knowns []int, n int) []int {
	wg := sync.WaitGroup{}
	wg.Add(n)

	subOffset := offset / (n - 1)
	results := make([][]int, 0)

	for i:=0; i<n; i ++ {
		result := make([]int, 0)
		results = append(results, result)
	}
	for i:=0; i<n; i++ {
		go func(k int) {
			p1 := p+ k*subOffset
			subOffset1 := subOffset
			if p1 + subOffset > p + offset {
				subOffset1 = p + offset - p1
			}
			r := Primes(p1, subOffset1, knowns)
			results[k] = append(results[k], r...)
			wg.Done()
		}(i)
	}

	wg.Wait()
	for _, result := range results {
		knowns = append(knowns, result...)
	}
	return knowns
}

func Primes(p int, offset int, knowns []int) []int {

	knowns2 := make([]int, 0)

	for i:=0; i< offset; i++ {
		isPrime := true
		pi := p + i
	    for _, k := range knowns {
	    	if k*k <= pi {
	    		if pi % k == 0 {
	    			isPrime = false
	    			break
				}
	    	} else {
				break
	    	}
	    }

	    if isPrime {
			knowns2 = append(knowns2, pi)
		}
	}

	return knowns2
}



func SingleThreadPrimes(p int, offset int, knowns []int) []int {

	for i:=0; i< offset; i++ {
		isPrime := true
		pi := p + i
		for _, k := range knowns {
			if k*k <= pi {
				if pi % k == 0 {
					isPrime = false
					break
				}
			} else {
				break
			}
		}

		if isPrime {
			knowns = append(knowns, pi)
		}
	}

	return knowns
}


