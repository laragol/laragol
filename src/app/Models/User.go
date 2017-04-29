package Models

type User struct {
	ID    uint64 `json:"id"      gorm:"primary_key"`
	Name  string `json:"name"    gorm:"size:100"`
	Email string `json:"email"   gorm:"size:100"`
	//Country Country `json:"country" gorm:"ForeignKey:country_id"`
	//Country   Country `json:"country" gorm:"save_associations:false"`
	Country   Country `json:"country"`
	CountryID uint64  `json:"-"`
}
