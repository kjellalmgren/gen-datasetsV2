package main

import (
	"bufio"
	"fmt"
	"gen-datasets/version"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Banner text
const (
	// TETRACON banner
	TETRACON = `
_________    __
|__    __|   | |
   |  |  ___ | |_   ____  ___   ___ ___  _ __ 
   |  | / _ \|  _| /  __|/ _ \ / __/ _ \| '_ \
   |  | \ __/| |_  | |  | (_| | (_| (_) | | | | 
   |__| \___| \__| |_|   \__,_|\___\___/|_| |_| 
Server-version: %s Model-version: %s Model-date: %s
`
)

// NUMBERS Documentation
const NUMBERS = 1000

// MIN - Hundra tusen
const MIN = 1000

// MAX - 10000
const MAX = 10000 // 50 miljoner
// NumberofRegion Documentation
const NumberofRegion = 4

// NumberofOffice Documentation
const NumberofOffice = 2

// REGIONS Documentation
const REGIONS = "SYD;NORR;VÄST;ÖST"

// OFFICES Documentation
const OFFICES = "10;20"

// OfficeType documentation
type OfficeType struct {
	OfficeID string `json:"officeID"`
	Name     string `json:"Name"`
}

// OfficesType documentation
type OfficesType struct {
	Offices []OfficeType
}

// RegionType Documentation
type RegionType struct {
	Name string `json:"name"`
}

// RegionsType Documntation
type RegionsType struct {
	Regions []RegionType
}

var (
	srv  bool
	vrsn bool
	date bool
)

var (
	min int64
	max int64
)

//AddItem description
func (regions *RegionsType) AddItem(region RegionType) {
	regions.Regions = append(regions.Regions, region)
}

//AddItem description
func (offices *OfficesType) AddItem(office OfficeType) {
	offices.Offices = append(offices.Offices, office)
}

// CreateRegions documentation
func CreateRegions() RegionsType {

	regions := RegionsType{}
	region := RegionType{}
	//
	items := strings.Split(REGIONS, ";")
	for i := range items {
		region.Name = items[i]
		regions.AddItem(region)
	}
	return regions
}

// CreateOffices documentation
func CreateOffices() OfficesType {

	offices := OfficesType{}
	office := OfficeType{}
	//
	items := strings.Split(OFFICES, ";")
	for i := range items {
		office.OfficeID = items[i]
		office.Name = items[i+1]
	}
	return offices
}

// init documwentation
func init() {
	//regions := {"SYD", "NORR", "ÖST", "VÄST"}

	// instanciate a new logger
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
	color.Set(color.FgHiGreen)
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion(), version.ModelDate()))
	color.Unset()
}

// our main function
func main() {

	regions := CreateRegions()
	fmt.Printf("Regions Length: %d", len(regions.Regions))
	offices := CreateOffices()
	fmt.Printf("Offices Length: %d", len(offices.Offices))
	//
	color.Set(color.FgHiYellow)
	fmt.Printf("Started on server: ")
	color.Set(color.FgHiRed)
	fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
	fmt.Printf("\r\n")
	color.Set(color.FgHiGreen)
	fmt.Printf("gen-datasets started...\r\n")
	color.Unset()
	startTime1 := time.Now()
	//
	//min = 10000 // minimum value
	//max = 90000 // max value to generate
	fmt.Printf("Range to be used: (%d - %d) number of records to produce %d\r\n",
		MIN, MAX, NUMBERS)
	// #############################
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Start init integer array\r\n")
	a := [NUMBERS]float64{}
	fmt.Printf("Start random order array\r\n")

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// #############################
	header := []byte("Region,Office,Reveue,Segment\r\n")
	//
	err := ioutil.WriteFile("csv/segment_training_v4.csv", header, 0644)
	check(err)
	f, err := os.Create("csv/segment_training_v4.csv")
	check(err)
	defer f.Close()
	//
	//
	w := bufio.NewWriter(f)
	b1, err := w.WriteString(fmt.Sprintf("%s", header))
	btot := 0
	//
	for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := randomNumber(MIN, MAX)
		b2, err := w.WriteString(fmt.Sprintf("10.0,100.0,%.1f,%d\r\n",
			float64(j), getSegment(j)))
		check(err)
		btot = btot + b2
	}
	w.Flush()
	fmt.Printf("Wrote %d bytes\r\n", btot+b1)
	color.Set(color.FgHiGreen)
	//fmt.Printf("- File: %s, # of lines: %d, processing time: %s \r\n",
	//	fileName, lineCount, time.Since(startTime1))
	fmt.Printf("gen-datasets finnished in %s...\r\n", time.Since(startTime1))
}

// randomNumber
func randomNumber(min float64, max float64) float64 {

	var rn float64
	rn = 0
	for rn == 0 || (rn < min && rn > max) {
		rand.Seed(time.Now().UnixNano())
		rn = min + rand.Float64()*(max-min) // rand.int63(max)
	}
	return rn
}

//
//	getHostName documentation
func getHostname() string {

	hostname, err1 := os.Hostname()
	if err1 != nil {
		//log.Printf("Hostname: %s", hostname)
		fmt.Println("Error when try to resolve Hostname: ", hostname)
	}
	return hostname
}

// getSegment for value
func getSegment(i float64) int {
	//revenue = rand.Intn(max-min) + min
	//value := max - min // max - min = 1 000 000 - 100 000 = 900 000
	value := float64(i)
	_max := float64(MAX)
	percent := value / _max
	segment := 0
	if (percent > float64(0.0)) && (percent <= float64(0.25)) {
		segment = 0
	}
	if (percent > float64(0.25)) && (percent <= float64(0.50)) {
		segment = 1
	}
	if (percent > float64(0.50)) && (percent <= float64(0.75)) {
		segment = 2
	}
	if (percent > float64(0.75)) && (percent <= float64(1.00)) {
		segment = 3
	}
	//fmt.Println(fmt.Sprintf("value: %d - %d - %f (%d)", max, i, percent, segment))
	return segment
}

// check for error
func check(e error) {
	if e != nil {
		panic(e)
	}
}
