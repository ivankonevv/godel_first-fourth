package first

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

func (b *Book) GetInfo() string {
	return fmt.Sprintf(
		"Книга %s была опубликована %s. На момент выхода стоимость была около %.2f руб. Авторами произведения являются %v",
		b.Title, b.PubDate, b.Price, b.Authors,
	)
}

func First() {
	var author Author
	author.Id = 1
	author.FirstName = "Иван"
	author.LastName = "Конев"
	author.Country = "Беларусь"
	authors := append([]Author{}, author)
	book := Book{Id: 1, Title: "Godel", PubDate: "2021.02.02", Price: 1199.0, IsAvailable: true, Authors: authors}
	fmt.Println(book.GetInfo())
}
