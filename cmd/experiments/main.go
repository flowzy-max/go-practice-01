package main

import "fmt"

func main() {
	var name string = "Артём"
	var group string = "ЭФБО-03-26"
	var course string = "Бэкенд-разработка на языке Go"
	var completed int = 1
	var ready bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", ready)

}
