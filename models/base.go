package models

type Base struct {
	ID        int `json:"id" db:"id"`
	CreatedAt int `json:"created_at" db:"created_at"`
	UpdatedAt int `json:"updated_at" db:"updated_at"`
	DeletedAt int `json:"deleted_at" db:"deleted_at"`
}

type Meta struct {
	TotalData   int `json:"total"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	CurrentPage int `json:"current_page"`
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func (m Meta) GetLimitAndOffset() (int, int) {
	return m.PerPage, (m.CurrentPage - 1) * m.PerPage
}

func (m Meta) SetTotalData(totalData int) Meta {

	m.TotalData = totalData
	m.TotalPage = totalData / m.PerPage
	if totalData%m.PerPage != 0 {
		m.TotalPage++
	}
	return m
}
