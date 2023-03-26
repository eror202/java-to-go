package entity

type Person struct {
	ID       string `json:"-" db:"id"`
	FullName string `json:"fullName" db:"fullname"`
	Title    string `json:"title" db:"title"`
	Age      string `json:"age" db:"age"'`
}

func (p *Person) SetFullName(fullName string) {
	p.FullName = fullName
}

func (p *Person) GetFullName() string {
	return p.FullName
}

func (p *Person) SetTitle(title string) {
	p.Title = title
}

func (p *Person) GetTitle() string {
	return p.Title
}

func (p *Person) SetAge(age string) {
	p.Age = age
}

func (p *Person) GetAge() string {
	return p.Age
}
