package entity

type Book struct {
	ID        string `json:"-" db:"id"`
	Title     string `json:"title" db:"title"`
	Author    string `json:"author" db:"author"`
	PageCount string `json:"pageCount" db:"pageCount"`
	Person    Person `json:"person" db:"person"`
}

func (p *Book) SetTitle(title string) {
	p.Title = title
}

func (p *Book) GetTitle() string {
	return p.Title
}

func (p *Book) SetAuthor(title string) {
	p.Title = title
}

func (p *Book) GetAuthor() string {
	return p.Author
}

func (p *Book) SetPageCount(pageCount string) {
	p.PageCount = pageCount
}

func (p *Book) GetPageCount() string {
	return p.PageCount
}

func (p *Book) SetPerson(person Person) {
	p.Person = person
}

func (p *Book) GetPerson() Person {
	return p.Person
}
