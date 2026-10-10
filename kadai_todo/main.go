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

func printTodos(items []ToDo) {
	for _, item := range items {
		var status string
		if item.Completed {
			status = "[完　了]"
		} else {
			status = "[未完了]"
		}
		fmt.Printf("%s（ID：%d）%s\n", status, item.ID, item.Title)
	}
}
