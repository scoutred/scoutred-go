package scoutred

import "time"

type Contractor struct {
	ID                 *int32          `json:"id"`
	LicenseNumber      *string         `json:"licenseNumber"`
	LicenseState       *string         `json:"licenseState"`
	FullName           *string         `json:"fullName"`
	CompanyName        *string         `json:"companyName"`
	CompanyDescription *string         `json:"companyDescription"`
	Phone              *string         `json:"phone"`
	Address1           *string         `json:"address1"`
	Address2           *string         `json:"address2"`
	City               *string         `json:"city"`
	State              *string         `json:"state"`
	Zip                *string         `json:"zip"`
	Email              *string         `json:"email"`
	Trade              ContractorTrade `json:"trade"`
	Created            *time.Time      `json:"created"`
	Updated            *time.Time      `json:"updated"`
}
