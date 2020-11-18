package main

import (
	"bufio"
	"flag"
	"fmt"
	"gen-datasets/offices"
	"gen-datasets/randoms"
	"gen-datasets/segments"
	"gen-datasets/version"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
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

var (
	srv  bool
	vrsn bool
	date bool
)

var (
	min int64
	max int64
)

// NUMBERSV4 Documentation
const NUMBERSV4 = 300000

// NUMBERSV5 documentation
const NUMBERSV5 = 1000000

// MIN - Hundra tusen
const MIN = 100000
const MAX = 10000000 // not used in V5
//
const V4MIN = 0
const V4MAX = 2000000 // NUMBERSV4 is 15% V4MAX

// NumberofRegion Documentation
const NumberofRegion = 4

// NumberofOffice Documentation
const NumberofOffice = 2

// generate v4 of the data-sets
var (
	v4 bool
	v5 bool
)

// init documwentation
func init() {

	// instanciate a new logger
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
	color.Set(color.FgHiGreen)
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion(), version.ModelDate()))
	color.Unset()
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&v4, "v4", false, "Generate old version of the datasets")
	flag.BoolVar(&v5, "v5", false, "Generate version 5 of the datasets")
	//
	flag.Usage = func() {
		//fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion(), version.ModelDate()))
		flag.PrintDefaults()
	}

	flag.Parse()
	//
	if vrsn {
		fmt.Printf("flag version %s\n", version.ServerVersion())
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		usageAndExit("Need parameters...", 0)
	}

	// parse the arg
	arg := flag.Args()[0]

	if arg == "v4" {
		v4 = true
	}
	if arg == "v5" {
		v5 = true
	}
	if arg == "help" {
		usageAndExit("description", 0)
	}

	if arg == "version" {
		fmt.Printf("GEN-DatasetV2 version history, model-date %s, model-version: %s, server-version: %s\n", version.ModelDate(), version.ModelVersion(), version.ServerVersion())
		os.Exit(0)
	}
}

// our main function, here we go...
func main() {

	//regions := regions.CreateRegions()
	//for i := range regions.Regions {
	//	fmt.Printf("Region: %s\r\n", regions.Regions[i].Name)
	//}
	//
	if v4 == true {
		color.Set(color.FgHiYellow)
		fmt.Printf("Started on server: ")
		color.Set(color.FgHiRed)
		fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
		fmt.Printf("\r\n")
		color.Set(color.FgHiGreen)
		fmt.Printf("gen-datasets v4 started...\r\n")
		color.Unset()
		generateV4Datasets(V4MIN, V4MAX)
	}
	//
	if v5 == true {
		color.Set(color.FgHiYellow)
		fmt.Printf("Started on server: ")
		color.Set(color.FgHiRed)
		fmt.Fprint(os.Stdout, fmt.Sprintf(getHostname()))
		fmt.Printf("\r\n")
		color.Set(color.FgHiGreen)
		fmt.Printf("gen-datasets v5 started...\r\n")
		color.Unset()
		generateV5Datasets()
	}
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

// check for error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// usageAndExit documentation
func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

// generateV4Datasets documentation
func generateV4Datasets(MIN float64, MAX float64) {

	startTime1 := time.Now()
	fmt.Printf("Range to be used: (%f - %f) number of records to produce: %d\r\n",
		MIN, MAX, NUMBERSV4)
	// #############################
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Start init integer array\r\n")
	a := [NUMBERSV4]float64{}
	fmt.Printf("Start random order array\r\n")

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// #############################
	header := []byte("region,office,revenue,segment\r\n")
	//
	err := ioutil.WriteFile("csv/segment_training_v4.csv", header, 0644)
	check(err)
	errv := ioutil.WriteFile("csv/segment_evaluation_v4.csv", header, 0644)
	check(errv)
	f, err := os.Create("csv/segment_training_v4.csv")
	check(err)
	fv, errv := os.Create("csv/segment_evaluation_v4.csv")
	check(err)
	defer f.Close()
	defer fv.Close()
	//
	//
	w := bufio.NewWriter(f)
	b1, err := w.WriteString(fmt.Sprintf("%s", header))
	wv := bufio.NewWriter(fv)
	bv, err := wv.WriteString(fmt.Sprintf("%s", header))
	btot := 0
	bvtot := 0
	//
	proc := 0.30
	v := int64(NUMBERSV4 * proc) // Get % of MAX for evaluation set
	fmt.Printf("Evaluation set: %d is %d %% of training set\r\n", v, (int64(proc * 100)))
	for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := randoms.RandomNumberv4(MIN, MAX)
		b1, err := w.WriteString(fmt.Sprintf("10,100,%.1f,%d\r\n",
			float64(j), segments.GetSegmentv4(j, MAX)))
		check(err)
		if v > 0 {
			bv, err := wv.WriteString(fmt.Sprintf("10,100,%.1f,%d\r\n",
				float64(j), segments.GetSegmentv4(j, MAX)))
			check(err)
			v--
			bvtot = bvtot + bv
		}
		btot = btot + b1
	}
	w.Flush()  // Flush training set
	wv.Flush() // Flush evaluation set
	fmt.Printf("Wrote %d bytes for training set\r\n", btot+b1)
	fmt.Printf("Wrote %d bytes for evaluation set ", bvtot+bv)
	color.Set(color.FgHiGreen)
	//fmt.Printf("- File: %s, # of lines: %d, processing time: %s \r\n",
	//	fileName, lineCount, time.Since(startTime1))
	fmt.Printf("gen-datasets v4 finnished in %s...\r\n", time.Since(startTime1))
}

// generateV5Datasets documenation
func generateV5Datasets() {

	startTime1 := time.Now()
	// Display regions and offices
	//fmt.Printf("Range to be used: (%d - %d) number of records to produce %d\r\n",MIN, MAX, NUMBERSV5)
	//
	fmt.Printf("Start init integer array\r\n")
	a := [NUMBERSV5]float64{}
	fmt.Printf("Start random order array\r\n")
	//
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// #############################
	header := []byte("region,office,revenue,segment\r\n")
	err := ioutil.WriteFile("csv/segment_training_v5.csv", header, 0644)
	check(err)
	f, err := os.Create("csv/segment_training_v5.csv")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	b1, err := w.WriteString(fmt.Sprintf("%s", header))
	btot := 0
	// #############################
	//
	offices := offices.CreateOffices()
	for j := range offices.Offices {
		fmt.Printf("RegionID: %.1f ", offices.Offices[j].RegionID)
		fmt.Printf("Region: %s ", offices.Offices[j].Region)
		fmt.Printf("OfficeId: %.1f - ", offices.Offices[j].OfficeID)
		fmt.Printf("OfficeName: %s\n", offices.Offices[j].Name)
		switch offices.Offices[j].RegionID {
		case 10.0:
			{
				switch offices.Offices[j].OfficeID {
				case 11.0:
					{
						max := 12000000.0
						distributions := []float64{0.25, 0.50, 0.75, 1.0}
						for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
							r := randoms.RandomNumberv5(MIN, max)
							b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
								offices.Offices[j].RegionID,
								offices.Offices[j].OfficeID,
								//float64(r), segments.GetSegmentv4(r, MAX)))
								float64(math.RoundToEven(r)),
								segments.GetSegmentv5(r, max, distributions)))
							check(err)
							btot = btot + b2
						}
						fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
					}

				case 12.0:
					{
						/*
							max := 12000000.0
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				}
			}

		case 20.0:
			{
				switch offices.Offices[j].OfficeID {
				case 21.0:
					{
						/*
							max := 12000000.0
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				case 22.0:
					{
						/*
							max := 12000000.0
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				}
			}

		case 30.0:
			{
				switch offices.Offices[j].OfficeID {
				case 31.0:
					{
						/*
							max := 15000000.0
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				case 32.0:
					{
						/*
							max := 15000000.0
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				}
			}

		case 40:
			{
				switch offices.Offices[j].OfficeID {
				case 41:
					{
						/*
							max := 8000000.0
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				case 42:
					{
						/*
							distributions := []float64{0.25, 0.50, 0.75, 1.0}
							max := 5000000.0
							for i := int64(len(a)) - 1; i > 0; i-- { // Fisher–Yates shuffle
								r := randoms.RandomNumberv5(MIN, max)
								b2, err := w.WriteString(fmt.Sprintf("%.1f,%.1f,%.1f,%d\r\n",
									offices.Offices[j].RegionID,
									offices.Offices[j].OfficeID,
									//float64(r), segments.GetSegmentv4(r, MAX)))
									float64(math.RoundToEven(r)),
									segments.GetSegmentv5(r, max, distributions)))
								check(err)
								btot = btot + b2
							}
							fmt.Printf("Range to be used: (%dSEK - %.0fSEK) number of records to produce %d\r\n", MIN, max, NUMBERSV5)
						*/
					}
				}
			}
		}

	}
	w.Flush()
	fmt.Printf("Wrote %d bytes\r\n", btot+b1)
	color.Set(color.FgHiGreen)
	fmt.Printf("gen-datasets v5 finnished in %s...\r\n", time.Since(startTime1))
}
