package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ······················task1_1
func function_int(num *int) {
	*num += 10
}

// ······················task1_2
func function_slice(nums *[]int) {
	for i, num := range *nums {
		(*nums)[i] = num * 2
	}
}

// ······················task2_1
func goroutinetask() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println(s, 2*i)
		}
	}("偶数：")
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println(s, 2*i-1)
		}
	}("奇数：")

	wg.Wait()
}

// ······················task2_2
type Task func()

func do_tasks(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for i := 0; i < len(tasks); i++ {
		go func(task Task) {
			defer wg.Done()
			now := time.Now()
			task()
			after := time.Now()
			fmt.Println(i+1, "方法的运行时间:", after.Sub(now))
		}(tasks[i])
	}
	wg.Wait()
}

// ······················task3_1
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (*Rectangle) Area() {
	fmt.Println("Rectangle-Area")
}

func (*Rectangle) Perimeter() {
	fmt.Println("Rectangle-Perimeter")
}

type Circle struct {
}

func (*Circle) Area() {
	fmt.Println("Circle-Area")
}

func (*Circle) Perimeter() {
	fmt.Println("Circle-Perimeter")
}

// ······················task3_2
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID string
}

func (employee *Employee) PrintInfo() {
	fmt.Println(*employee)
}

// ······················task4_1
func sendmessage_1() {
	var wg sync.WaitGroup
	wg.Add(2)
	channel := make(chan (int))
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			channel <- i
		}
		close(channel)
	}()

	go func() {
		defer wg.Done()
		for value := range channel {
			fmt.Println("read channel value:", value)
		}
	}()

	wg.Wait()
}

// ······················task4_2
func sendmessage_2() {
	var wg sync.WaitGroup
	wg.Add(2)
	channel := make(chan (int), 10)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			channel <- i
			fmt.Println("write channel value:", i)
		}
		close(channel)
	}()

	go func() {
		defer wg.Done()
		for value := range channel {
			fmt.Println("read channel value:", value)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}

var index int64 = 0

// ······················task5_1 lock
func addnum_1() {
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock()
				index++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("index:", index)
}

// ······················task5_2 atomic
func addnum_2() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&index, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("index:", index)
}
func main() {
	// digit := 8
	// function_int(&digit)
	// fmt.Println(digit)

	// nums := []int{1, 2, 3, 8, 5}
	// function_slice(&nums)
	// fmt.Println(nums)

	// goroutinetask()

	// tasks := []Task{
	// 	func() { fmt.Println("任务1"); time.Sleep(time.Second) },
	// 	func() { fmt.Println("任务2"); time.Sleep(time.Second * 2) },
	// 	func() { fmt.Println("任务3"); time.Sleep(time.Second * 3) },
	// }
	// do_tasks(tasks)

	// circle := Circle{}
	// rectangle := Rectangle{}
	// (&circle).Area()
	// circle.Area()
	// (&circle).Perimeter()
	// (&rectangle).Area()
	// (&rectangle).Perimeter()

	// person := Person{
	// 	Name: "QX", Age: 22,
	// }
	// employee := Employee{
	// 	Person:     person,
	// 	EmployeeID: "9527",
	// }
	// (&employee).PrintInfo()

	// sendmessage_1()
	// sendmessage_2()

	// addnum_1()
	// addnum_2()

}
