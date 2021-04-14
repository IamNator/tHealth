package models

import (
	"time"

	"github.com/lib/pq"
)

//ConsultPhys  .. (ician) ...f
type ConsultPhys struct {
	ID              int            `json:"id" db:"id" gorm:"primary_key;autoIncrement;"`
	FullName        string         `json:"fullname,omitempty" db:"fullname"`
	Email           string         `json:"email,omitempty" db:"email"`
	Telephone       int64          `json:"telephone,omitempty" db:"telephone" gorm:"type:bigint"`
	Languages       pq.StringArray `json:"languages,omitempty" db:"languages" gorm:"type:text[]"`
	Specialty       pq.StringArray `json:"specialty,omitempty" db:"specialty" gorm:"type:text[]"`
	TimeIn          time.Time      `json:"time_in,omitempty" db:"time_in"`
	Training        pq.StringArray `json:"training,omitempty" db:"training" gorm:"type:text[]"`
	Certification   pq.StringArray `json:"certification,omitempty" db:"certification" gorm:"type:text[]"`
	SpecialSkills   pq.StringArray `json:"special_skills,omitempty" db:"specialSkills" gorm:"type:text[]"`
	SpecialInterest pq.StringArray `json:"special_interest,omitempty" db:"special_interest" gorm:"type:text[]"`
	Experience      pq.StringArray `json:"experience,omitempty" db:"experience" gorm:"type:text[]"`
	TimeOut         time.Time      `json:"time_out,omitempty" gorm:"timestamp_with_time_zone;NOT_NULL"`
	Location        string         `json:"location_" db:"location_"`
}

//CreateConsultPhys  ...
type CreateConsultPhys struct {
	FullName        string         `json:"fullname" db:"fullname" validate:"required"`
	Email           string         `json:"email" db:"email" validate:"required"`
	Telephone       int64          `json:"telephone" db:"telephone" gorm:"type:bigint" validate:"required"`
	Languages       pq.StringArray `json:"languages" db:"languages" gorm:"type:text[]" validate:"required"`
	Specialty       pq.StringArray `json:"specialty" db:"specialty" gorm:"type:text[]" validate:"required"`
	TimeIn          time.Time      `json:"time_in" db:"time_in"`
	Training        pq.StringArray `json:"training" db:"training" gorm:"type:text[]" validate:"required"`
	Certification   pq.StringArray `json:"certification" db:"certification" gorm:"type:text[]"`
	SpecialSkills   pq.StringArray `json:"special_skills" db:"specialSkills" gorm:"type:text[]"`
	SpecialInterest pq.StringArray `json:"special_interest" db:"special_interest" gorm:"type:text[]"`
	Experience      pq.StringArray `json:"experience" db:"experience" gorm:"type:text[]" validate:"required"`
	TimeOut         time.Time      `json:"time_out" gorm:"timestamp_with_time_zone;NOT_NULL"`
	Location        string         `json:"location_" db:"location_" validate:"required"`
}
