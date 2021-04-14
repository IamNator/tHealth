package models

import "github.com/lib/pq"

//LocalProvider ...
type LocalProvider struct {
	ID            int            `json:"id" db:"id" gorm:"primary_key;autoIncrement;"`
	FullName      string         `json:"fullname" db:"fullname"`
	Gender        string         `json:"gender" db:"gender"`
	ProviderType  string         `json:"provider_type" db:"provider_type"`
	Training      pq.StringArray `json:"training" db:"training" gorm:"type:text[]"`
	Experience    pq.StringArray `json:"experience" db:"experience" gorm:"type:text[]"`
	ComputerSkill pq.StringArray `json:"computer_skill" db:"computer_skill" gorm:"type:text[]"`
	LocalFacility pq.StringArray `json:"local_facility" db:"local_facility" gorm:"type:text[]"`
}
