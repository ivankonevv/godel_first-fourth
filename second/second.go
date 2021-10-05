package second

import (
	"first_test/first"
	"fmt"
)

func Second() {
	var a = [3]first.Author{
		{
			Id:        1,
			FirstName: "Иван",
			LastName:  "Конев",
			Country:   "Беларусь",
		},
		{
			Id:        2,
			FirstName: "Евгений",
			LastName:  "Онегин",
			Country:   "Россия",
		},
		{
			Id:        3,
			FirstName: "Алексей",
			LastName:  "Мангутов",
			Country:   "Беларусь",
		},
	}
	for i, v := range a {
		fmt.Printf("%v | %v \n", i, v)
	}
	var b = [3]first.Book{
		{
			Id:          1,
			Title:       "Первая книга",
			PubDate:     "2000.12.12",
			IsAvailable: true,
			Authors:     []first.Author{a[0]},
		},
		{
			Id:          2,
			Title:       "Вторая книга",
			PubDate:     "2004.12.12",
			IsAvailable: false,
			Authors:     []first.Author{a[1], a[2]},
		},
		{
			Id:          3,
			Title:       "Третья книга",
			PubDate:     "2001.12.12",
			IsAvailable: true,
			Authors:     []first.Author{a[2]},
		},
	}
	for i, v := range b {
		fmt.Printf("%v | %v \n", i, v)
	}

}
