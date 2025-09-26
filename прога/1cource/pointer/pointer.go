package main

import "fmt"

func main() {
	NameList := []string{}
	var CollectionSum int
	var name string

	var sum int
	for {
		name = ""
		sum = 0
		fmt.Print("Введите имя (введите 'все' чтобы закончить): ")
		fmt.Scan(&name)
		if name == "все" {
			break
		}
		fmt.Print("Введите сумму пожертвования: ")
		fmt.Scan(&sum)
		fmt.Println(name, sum)
		if name == "" && sum == 0 {
			fmt.Print("Неверно введен текст!")

		}
		if name != "" && sum == 0 {
			fmt.Print("Введите сумму пожертвования еще раз: ")
			fmt.Scan(&sum)
		} else if name == "" && sum != 0 {

			fmt.Print("Введите имя еще раз: ")
			fmt.Scan(&name)
		}

		NameList = append(NameList, name)
		CollectionSum += sum

	}
	fmt.Print("Список пожертвовавших: ", NameList)

	fmt.Println("Сумма сборов: ", CollectionSum)
}
