package models

// RegionType Documentation
type RegionType struct {
	Name string `json:"name"`
}

// RegionsType Documntation
type RegionsType struct {
	Regions []RegionType
}

//AddItem description
func (regions *RegionsType) AddItem(region RegionType) {
	regions.Regions = append(regions.Regions, region)
}

// OfficeType documentation
type OfficeType struct {
	OfficeID string `json:"officeID"`
	Name     string `json:"Name"`
}

// OfficesType documentation
type OfficesType struct {
	Offices []OfficeType
}

//AddItem description
func (offices *OfficesType) AddItem(office OfficeType) {
	offices.Offices = append(offices.Offices, office)
}
