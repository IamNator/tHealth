package models

import (
	"time"

	"github.com/lib/pq"
)

//LocalFacility  .. (ician) ...f
type LocalFacility struct {
	ID                  int            `json:"id" db:"id" gorm:"primary_key;autoIncrement;"`
	FacilityName        string         `json:"facility_name" db:"facility_name" gorm:"type:text"`
	FacilityType        string         `json:"facility_type" db:"facility_type" gorm:"type:text"`
	FacilityLocation    string         `json:"facility_location" db:"facility_location" gorm:"type:text"`
	OIC                 string         `json:"o_i_c" db:"o_i_c"`
	Staffing            pq.StringArray `json:"staffing" db:"staffing" gorm:"type:text[]"`
	Supplies            pq.StringArray `json:"supplies" db:"supplies" gorm:"type:text[]"`
	DiagnosticEquipment pq.StringArray `json:"diagnostic_quiptment" db:"diagnostic_equipment" gorm:"type:text[]"`
	TelehealthEquipment pq.StringArray `json:"telehealth_equipment" db:"telehealth_equipment" gorm:"type:text[]"`
	PharmacyNear        time.Duration  `json:"pharmacy_near" db:"pharmacy_near" gorm:"type:interval"`
	MajorMedicCentre    time.Duration  `json:"major_medic_centre" db:"major_medic_centre" gorm:"type:interval"`
	HospitalNear        time.Duration  `json:"hospital_near" db:"hospital_near" gorm:"type:interval"`
	PhysicianNear       time.Duration  `json:"physician_near" db:"physician_near"  gorm:"type:interval"`
	Environment         string         `json:"environment" db:"environment"`
	EconomicStatus      string         `json:"economic_status" db:"economic_status"`
}
