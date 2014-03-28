package queue

import "fmt"

type Thunder struct {
	Queue
	contentType string
}

func (th *Thunder) newFunc() (*int, func() error) {
	var a = 3
	return &a, func() error {
		fmt.Println("hello")
		a = a + 3
		return nil
	}
}

func (th *Thunder) Fetch() {
	a, f := th.newFunc()
	r := th.Add(f)
	fmt.Println(<-r, *a)
}
