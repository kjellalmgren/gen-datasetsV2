package models

// OfficeType documentation
type OfficeType struct {
	OfficeID float64 `json:"officeID"`
	Name     string  `json:"Name"`
	RegionID float64 `json:"RegionID"`
	Region   string  `json:"Region"`
}

// OfficesType documentation
type OfficesType struct {
	Offices []OfficeType
}

//AddItem description
func (offices *OfficesType) AddItem(office OfficeType) {
	offices.Offices = append(offices.Offices, office)
}
