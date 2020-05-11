package offices

import (
	"gen-datasets/models"
	"strings"
)

// OFFICENAMES documentation
const OFFICENAMES = "Malmö;Göteborg;stockholm;Uppsala;Karlstad;örebro;Luleå;Kiruna"

// REGION Documentation
const REGION = "SYD;SYD;ÖST;ÖST;VÄST;VÄST;NORR;NORR"

// CreateOffices documentation
func CreateOffices() models.OfficesType {

	officeIDs := [8]float64{11, 12, 21, 22, 31, 32, 41, 42}
	regionIDs := [8]float64{10.0, 10.0, 20.0, 20.0, 30.0, 30.0, 40.0, 40.0}
	//
	offices := models.OfficesType{}
	office := models.OfficeType{}
	//
	//for x := range officeIDs {
	//	fmt.Printf("%.1f - %.1f\r\n", officeIDs[x], regionIDs[x])
	//}
	items1 := strings.Split(OFFICENAMES, ";")
	items3 := strings.Split(REGION, ";")
	for i := range items1 {
		//office.OfficeID = items[i]
		office.OfficeID = officeIDs[i]
		office.Name = items1[i]
		office.RegionID = regionIDs[i]
		office.Region = items3[i]
		offices.AddItem(office)
	}
	return offices
}
