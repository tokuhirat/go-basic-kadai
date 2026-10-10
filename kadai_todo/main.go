package main

import "fmt"

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func (item *ToDo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", item.ID)
	item.Completed = true
}
