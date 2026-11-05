package main

import "fmt"

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В Работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

func countStatus(statuses []string, target string) int {
	k := 0
	for _, b := range statuses {
		if b == target {
			k++
		}
	}
	return k
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))
	fmt.Println(len(statuses))
	for i, v := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(v))
	}
	fmt.Println(countStatus(statuses, "done"))
	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	fmt.Println("new=", counts["new"], "in_progress=", counts["in_progress"], "done=", counts["done"])

}
