package entity

type Application struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Text string `json:"text"`
}

func (a *Application) ToText() string {
	return "Новая заявка!\nИмя: " + a.Name + "\nТелефон: " + a.Phone + "\nТекст: " + a.Text
}
