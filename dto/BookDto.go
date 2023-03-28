package dto

type BookDto struct {
	ID        string    `json:"-" db:"id"`
	Title     string    `json:"title" db:"title"`
	Author    string    `json:"author" db:"author"`
	PageCount string    `json:"pageCount" db:"pageCount"`
	Person    PersonDto `json:"person" db:"person"`
}

func (p *BookDto) SetTitle(title string) {
	p.Title = title
}

func (p *BookDto) GetTitle() string {
	return p.Title
}

func (p *BookDto) SetAuthor(title string) {
	p.Title = title
}

func (p *BookDto) GetAuthor() string {
	return p.Author
}

func (p *BookDto) SetPageCount(pageCount string) {
	p.PageCount = pageCount
}

func (p *BookDto) GetPageCount() string {
	return p.PageCount
}

func (p *BookDto) SetPerson(person PersonDto) {
	p.Person = person
}

func (p *BookDto) GetPerson() PersonDto {
	return p.Person
}
