package models

import (
	"time"
)

//UpdatePatientProfileInput for request to update PatientProfile
type UpdatePatientProfileInput struct {
	FullName        string    `json:"fullname"`
	Gender          string    `json:"gender"`
	Contact         string    `json:"contact"`
	ReligionCulture string    `json:"religion_culture"`
	Telephone       string    `json:"telephone"`
	Languages       []string  `json:"languages"`
	Assessment      string    `json:"assessment"`
	History         []string  `json:"history"`
	Objective       []string  `json:"objective"`
	AttachedFiles   []string  `json:"attached_files"`
	ResponseTime    time.Time `json:"response_time"`
	TimeofCompalint time.Time `json:"time_of_compalint"`
}


//CreatePatientProfileInput for request to update PatientProfile
type CreatePatientProfileInput struct {
	FullName        string    `json:"fullname"`
	Gender          string    `json:"gender"`
	Contact         string    `json:"contact"`
	ReligionCulture string    `json:"religion_culture"`
	Telephone       string    `json:"telephone"`
	Languages       []string  `json:"languages"`
	Assessment      string    `json:"assessment"`
	History         []string  `json:"history"`
	Objective       []string  `json:"objective"`
	AttachedFiles   []string  `json:"attached_files"`
	ResponseTime    time.Time `json:"response_time"`
	TimeofCompalint time.Time `json:"time_of_compalint"`
}


//PatientProfile is a model for patient_profile_table
type PatientProfile struct {
	ID              string    `json:"id" gorm:"primary_key"`
	FullName        string    `json:"fullname"`
	Gender          string    `json:"gender"`
	Contact         string    `json:"contact"`
	ReligionCulture string    `json:"religion_culture"`
	Telephone       string    `json:"telephone"`
	Languages       []string  `json:"languages"`
	Assessment      string    `json:"assessment"`
	History         []string  `json:"history"`
	Objective       []string  `json:"objective"`
	AttachedFiles   []string  `json:"attached_files"`
	ResponseTime    time.Time `json:"response_time"`
	TimeofCompalint time.Time `json:"time_of_compalint"`
}
