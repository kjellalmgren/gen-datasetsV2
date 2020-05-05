package offices

import (
	"gen-datasets/models"
	"strings"
)

// OFFICES Documentation
const OFFICES = "10;20;30"

// OFFICENAMES documentation
const OFFICENAMES = "UMEÅ;STOCKHOLM;KALMAR"

// CreateOffices documentation
func CreateOffices() models.OfficesType {

	offices := models.OfficesType{}
	office := models.OfficeType{}
	//
	items := strings.Split(OFFICES, ";")
	items1 := strings.Split(OFFICENAMES, ";")
	for i := range items {
		office.OfficeID = items[i]
		office.Name = items1[i]
		offices.AddItem(office)
	}
	return offices
}
