package dto

type PersonDto struct {
	ID       string `json:"-" db:"id"`
	FullName string `json:"fullName" db:"fullname"`
	Title    string `json:"title" db:"title"`
	Age      string `json:"age" db:"age"'`
}

func (p *PersonDto) SetFullName(fullName string) {
	p.FullName = fullName
}

func (p *PersonDto) GetFullName() string {
	return p.FullName
}

func (p *PersonDto) SetTitle(title string) {
	p.Title = title
}

func (p *PersonDto) GetTitle() string {
	return p.Title
}

func (p *PersonDto) SetAge(age string) {
	p.Age = age
}

func (p *PersonDto) GetAge() string {
	return p.Age
}
