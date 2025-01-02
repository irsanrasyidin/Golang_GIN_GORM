package model

type LogicModel struct {
	ID         string `json:"ID" gorm:"column:id"`
	First_name string `json:"First_name" gorm:"column:first_name"`
	Last_name  string `json:"Last_name" gorm:"column:last_name"`
	Email      string `json:"Email" gorm:"column:email"`
	Gender     string `json:"Gender" gorm:"column:gender"`
	Avatar     string `json:"Avatar" gorm:"column:avatar"`
}

type Pagination struct {
	Page       int    `json:"Page"`
	PageType   string `json:"Page_Type"`
	PageSize   int    `json:"Page_Size"`
	TotalItems int    `json:"Total_Items"`
}

type ExecutionModel struct {
	Nama     string  `json:"Nama" gorm:"column:nama"`
	Masuk    string  `json:"Masuk" gorm:"column:masuk"`
	Keluar   string  `json:"Keluar" gorm:"column:keluar"`
	Duration float64 `json:"Duration" gorm:"column:duration"`
	Coba     int     `json:"Coba" gorm:"column:coba"`
	Status   string  `json:"Status" gorm:"column:status"`
}

type ExecutionResultModel struct {
	Nama         string    `json:"Nama"`
	Average      float64   `json:"Average"`
	Top          float64   `json:"Top"`
	Data         []float64 `json:"Data"`
	S_Deviasi    float64   `json:"S_Deviasi"`
	Outliner     int       `json:"Outliner"`
	OutlinerData []float64 `json:"OutlinerData"`
}
