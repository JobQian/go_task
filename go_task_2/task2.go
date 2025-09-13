package main

import (
	"fmt"
	"sync"
	"time"
)

// ······················task1
func function_int(num *int) {
	*num += 10
}

func function_slice(nums *[]int) {
	for i, num := range *nums {
		(*nums)[i] = num * 2
	}
}

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

	person := Person{
		Name: "QX", Age: 22,
	}
	employee := Employee{
		Person:     person,
		EmployeeID: "9527",
	}
	(&employee).PrintInfo()

}
