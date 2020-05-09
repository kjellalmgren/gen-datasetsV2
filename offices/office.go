package offices

import (
	"gen-datasets/models"
	"strings"
)

// OFFICENAMES documentation
const OFFICES = "10;20;30;40;50;60;70;80"
const OFFICENAMES = "Malmö;Göteborg;stockholm;Uppsala;Karlstad;örebro;Luleå;Kiruna"
const REGIONID = "10;10;20;20;30;30;40;40"
const REGION = "SYD;SYD;ÖST;ÖST;VÅST;VÄST;NORR;NORR"

// CreateOffices documentation
func CreateOffices() models.OfficesType {

	offices := models.OfficesType{}
	office := models.OfficeType{}
	//
	items := strings.Split(OFFICES, ";")
	items1 := strings.Split(OFFICENAMES, ";")
	items2 := strings.Split(REGIONID, ";")
	items3 := strings.Split(REGION, ";")
	for i := range items {
		office.OfficeID = items[i]
		office.Name = items1[i]
		office.RegionID = items2[i]
		office.Region = items3[i]
		offices.AddItem(office)
	}
	return offices
}
