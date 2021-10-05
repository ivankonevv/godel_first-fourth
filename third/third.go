package third

import (
	"fmt"
)

type Book struct {
	Id          int32
	Title       string
	PubDate     string
	Price       float32
	IsAvailable bool
	Authors     []Author
}

type Author struct {
	Id        int32
	FirstName string
	LastName  string
	Country   string
}

type Booker interface {
	GetInfo() string
	ForLoop() string
}

func (b Book) GetInfo() string {
	return fmt.Sprintf(
		"Книга %s была опубликована %s. На момент выхода стоимость была около %.2f руб. Авторы: %v.",
		b.Title, b.PubDate, b.Price, b.Authors,
	)
}

func (a Author) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Фамилия: %s", a.FirstName, a.LastName)
}

func (b Book) ForLoop() string {
	return fmt.Sprint(b)
}

func (a Author) ForLoop() string {
	return fmt.Sprint(a)
}

func Third() {
	var booker Booker
	author := []Author{
		{
			Id:        1,
			FirstName: "Иван",
			LastName:  "Конев",
			Country:   "Беларусь",
		},
		{
			Id:        2,
			FirstName: "Дмитрий",
			LastName:  "Неизвестный",
			Country:   "Украина",
		},
		{
			Id:        3,
			FirstName: "Иван",
			LastName:  "Иванов",
			Country:   "Казахстан",
		},
	}
	book := []Book{
		{
			Id:          1,
			Title:       "Godel",
			PubDate:     "2021.02.02",
			Price:       1199.0,
			IsAvailable: true, Authors: []Author{author[0], author[2]},
		},
		{
			Id:          2,
			Title:       "Godel1",
			PubDate:     "2021.02.03",
			Price:       1099.0,
			IsAvailable: true, Authors: []Author{author[0]},
		},
		{
			Id:          3,
			Title:       "Godel2",
			PubDate:     "2021.02.04",
			Price:       1399.0,
			IsAvailable: true, Authors: []Author{author[0], author[1], author[2]},
		},
	}
	for i, b := range book {
		booker = &b
		fmt.Printf("%v | %v \n", i, booker.ForLoop())
		fmt.Println(booker.GetInfo())
	}
	for i, a := range author {
		booker = &a
		fmt.Printf("%v | %v \n", i, booker.ForLoop())
		fmt.Println(booker.GetInfo())
	}
}
