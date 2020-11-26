package models

// OfficeType documentation
type OfficeType struct {
	OfficeID int64  `json:"officeID"`
	Name     string `json:"Name"`
	RegionID int64  `json:"RegionID"`
	Region   string `json:"Region"`
}

// OfficesType documentation
type OfficesType struct {
	Offices []OfficeType
}

//AddItem description
func (offices *OfficesType) AddItem(office OfficeType) {
	offices.Offices = append(offices.Offices, office)
}
