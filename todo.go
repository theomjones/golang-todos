package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func (t *todo) toggle() {
	t.Done = !t.Done
}

func (t *todo) setText(text string) {
	t.Text = text
}

// LIST

type todoList struct {
	Todos []todo `json:"todos"`
}

func newTodoList() todoList {
	filename := "todos.json"
	existing, err := ioutil.ReadFile(filename)

	list := todoList{}

	if len(existing) < 1 || err != nil {
		ioutil.WriteFile(filename, list.toJSON(), 0666)
	} else {
		err = json.Unmarshal(existing, &list)
		check(err)
	}

	return list
}

func (tl todoList) save() {
	filename := "todos.json"

	ioutil.WriteFile(filename, tl.toJSON(), 0666)
}

func (tl *todoList) add(t todo) {
	tl.Todos = append(tl.Todos, t)
	tl.save()
}

func (tl *todoList) remove(i int) {
	tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
	tl.save()
}

func (tl *todoList) markAllDone() {
	newList := []todo{}

	for _, t := range tl.Todos {
		t.Done = true
		newList = append(newList, t)
	}
	tl.Todos = newList
	tl.save()
}

func (tl *todoList) removeAll() {
	tl.Todos = []todo{}
	tl.save()
}

func (tl todoList) print() {
	for i, t := range tl.Todos {
		fmt.Printf("%v) %s - [%v]\n", i+1, t.Text, t.Done)
	}
}

func (tl todoList) toggleTodo(i int) error {
	if i < 0 || (i > len(tl.Todos)) {
		return errors.New("No todo at this index")
	}
	tl.Todos[i].toggle()
	tl.save()
	return nil
}

func (tl todoList) toJSON() []byte {
	s, err := json.Marshal(tl)
	check(err)
	return s
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
