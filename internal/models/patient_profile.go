package models

import (
	"time"

	"github.com/lib/pq"
)

//UpdatePatientProfileInput for request to update PatientProfile
type UpdatePatientProfileInput struct {
	FullName        string         `json:"fullname,omitempty"`
	Gender          string         `json:"gender,omitempty"`
	Contact         string         `json:"contact,omitempty"`
	ReligionCulture string         `json:"religion_culture,omitempty"`
	Telephone       string         `json:"telephone,omitempty"`
	Languages       pq.StringArray `json:"languages,omitempty"`
	Assessment      string         `json:"assessment,omitempty"`
	History         pq.StringArray `json:"history,omitempty"`
	Objective       pq.StringArray `json:"objective,omitempty"`
	AttachedFiles   pq.StringArray `json:"attached_files,omitempty"`
	ResponseTime    time.Time      `json:"response_time,omitempty"`
	TimeofCompalint time.Time      `json:"time_of_compalint,omitempty"`
}

//CreatePatientProfileInput for request to update PatientProfile
type CreatePatientProfileInput struct {
	FullName        string         `json:"fullname,omitempty"`
	Gender          string         `json:"gender,omitempty"`
	Contact         string         `json:"contact,omitempty"`
	ReligionCulture string         `json:"religion_culture,omitempty"`
	Telephone       string         `json:"telephone,omitempty"`
	Languages       pq.StringArray `json:"languages,omitempty"`
	Assessment      string         `json:"assessment,omitempty"`
	History         pq.StringArray `json:"history,omitempty"`
	Objective       pq.StringArray `json:"objective,omitempty"`
	AttachedFiles   pq.StringArray `json:"attached_files,omitempty"`
	ResponseTime    time.Time      `json:"response_time,omitempty"`
	TimeofCompalint time.Time      `json:"time_of_compalint,omitempty"`
}

//PatientProfiles is a model for patient_profile_table
type PatientProfiles struct {
	ID              uint64         `db:"id" json:"id" gorm:"primary_key;autoIncrement;"`
	FullName        string         `db:"fullname" json:"fullname,omitempty"`
	Gender          string         `db:"gender" json:"gender,omitempty"`
	Contact         string         `db:"contact" json:"contact,omitempty"`
	ReligionCulture string         `db:"religion_culture" json:"religion_culture,omitempty"`
	Telephone       string         `db:"rtelephone" json:"telephone,omitempty"`
	Languages       pq.StringArray `db:"languages" json:"languages,omitempty" gorm:"type:text[]"`
	Assessment      string         `db:"assessment" json:"assessment,omitempty"`
	History         pq.StringArray `db:"history" json:"history,omitempty" gorm:"type:text[]"`
	Objective       pq.StringArray `db:"objective" json:"objective,omitempty" gorm:"type:text[]"`
	AttachedFiles   pq.StringArray `db:"attached_files" json:"attached_files,omitempty" gorm:"type:text[]"`
	ResponseTime    time.Time      `db:"response_time" json:"response_time,omitempty"`
	TimeofCompalint time.Time      `db:"time_of_complaint" json:"time_of_compalint,omitempty"`
}
