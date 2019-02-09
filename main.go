package main

import (
	"os"
	"strconv"
)

func main() {
	list := newTodoList()

	cmd := os.Args[1]

	if cmd == "clear" {
		list.removeAll()
		list.print()
		return
	} else if cmd == "list" {
		list.print()
		return
	} else if cmd == "done" {
		list.markAllDone()
		list.print()
		return
	}

	arg := os.Args[2]

	switch cmd {
	case "add":
		list.add(todo{
			Text: arg,
		})
		break
	case "remove":
		i, _ := strconv.Atoi(arg)
		list.remove(i - 1)
		break
	case "toggle":
		i, _ := strconv.Atoi(arg)
		err := list.toggleTodo(i - 1)
		check(err)
		break
	}

	list.print()
}
