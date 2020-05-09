package models

// OfficeType documentation
type OfficeType struct {
	OfficeID string `json:"officeID"`
	Name     string `json:"Name"`
	RegionID string `json:"RegionID"`
}

// OfficesType documentation
type OfficesType struct {
	Offices []OfficeType
}

//AddItem description
func (offices *OfficesType) AddItem(office OfficeType) {
	offices.Offices = append(offices.Offices, office)
}
