# Практическая работа № 1. Знакомство с Go

Курс «Бэкенд-разработка на языке Go».

## Окружение

- Операционная система: macos.
- Версия Go: go version go1.27.2 darwin/arm64
- Редактор: VS Code

## Выполнено

первое и второе задание

## Запуск первой программы

Откройте терминал в папке проекта, где находятся `main.go` и `go.mod`.

Привет! Это моя первая программа на Go. (main.go)


## Проверка сборки и оформления

cmd/experiments/main.go

## Дополнительные программы

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


## Проверенные данные

Участник: Артём
Группа: ЭФБО-03-26
Курс: Бэкенд-разработка на языке Go
Завершено работ: 1
Интерес к бэкенду: true

## Что осталось сделать

доделать практическую работу 1