package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	id int
}

func (u *user) notify() {
	fmt.Println(u.name, u.email)
}

func main() {
	u := user{"Tinker", "111@qq.com"}
	a := admin{u, 1}
	a.notify() // 嵌入类型
	sendNotification(&u)
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
