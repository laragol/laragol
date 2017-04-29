package Models

type Country struct {
	ID   uint64 `json:"id"   gorm:"primary_key"`
	Name string `json:"name" gorm:"size:100"`
}
