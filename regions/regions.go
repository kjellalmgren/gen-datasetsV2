package regions

import (
	"gen-datasets/models"
	"strings"
)

// REGIONS Documentation
const REGIONS = "SYD;NORR;VÄST;ÖST"

// CreateRegions documentation
func CreateRegions() models.RegionsType {

	regions := models.RegionsType{}
	region := models.RegionType{}
	//
	items := strings.Split(REGIONS, ";")
	for i := range items {
		region.Name = items[i]
		regions.AddItem(region)
	}
	return regions
}
