package models

import (
	"time"

	"github.com/lib/pq"
)

//UpdatePatientProfileInput for request to update PatientProfile
type UpdatePatientProfileInput struct {
	FullName        string         `json:"fullname"`
	Gender          string         `json:"gender"`
	Contact         string         `json:"contact"`
	ReligionCulture string         `json:"religion_culture"`
	Telephone       string         `json:"telephone"`
	Languages       pq.StringArray `json:"languages"`
	Assessment      string         `json:"assessment"`
	History         pq.StringArray `json:"history"`
	Objective       pq.StringArray `json:"objective"`
	AttachedFiles   pq.StringArray `json:"attached_files"`
	ResponseTime    time.Time      `json:"response_time"`
	TimeofCompalint time.Time      `json:"time_of_compalint"`
}

//CreatePatientProfileInput for request to update PatientProfile
type CreatePatientProfileInput struct {
	FullName        string         `json:"fullname"`
	Gender          string         `json:"gender"`
	Contact         string         `json:"contact"`
	ReligionCulture string         `json:"religion_culture"`
	Telephone       string         `json:"telephone"`
	Languages       pq.StringArray `json:"languages" gorm:"type:text[]"`
	Assessment      string         `json:"assessment"`
	History         pq.StringArray `json:"history" gorm:"type:text[]"`
	Objective       pq.StringArray `json:"objective" gorm:"type:text[]"`
	AttachedFiles   pq.StringArray `json:"attached_files" gorm:"type:text[]"`
	ResponseTime    time.Time      `json:"response_time"`
	TimeofCompalint time.Time      `json:"time_of_compalint"`
}

//PatientProfiles is a model for patient_profile_table
type PatientProfiles struct {
	ID              uint64         `db:"id" json:"id" gorm:"primary_key;autoIncrement;"`
	FullName        string         `db:"fullname" json:"fullname"`
	Gender          string         `db:"gender" json:"gender"`
	Contact         string         `db:"contact" json:"contact"`
	ReligionCulture string         `db:"religion_culture" json:"religion_culture"`
	Telephone       string         `db:"rtelephone" json:"telephone"`
	Languages       pq.StringArray `db:"languages" json:"languages" gorm:"type:text[]"`
	Assessment      string         `db:"assessment" json:"assessment"`
	History         pq.StringArray `db:"history" json:"history" gorm:"type:text[]"`
	Objective       pq.StringArray `db:"objective" json:"objective" gorm:"type:text[]"`
	AttachedFiles   pq.StringArray `db:"attached_files" json:"attached_files" gorm:"type:text[]"`
	ResponseTime    time.Time      `db:"response_time" json:"response_time"`
	TimeofCompalint time.Time      `db:"time_of_complaint" json:"time_of_compalint"`
}
