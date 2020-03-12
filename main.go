package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var table = flag.String("t", "", "required, target table name")
var column = flag.String("c", "", "required, target column name")
var size = flag.Int("s", 1000, "optional, default 1000, record size")

func main() {
	flag.Parse()

	if *table == "" {
		fmt.Println("'-t' must be specified")
		return
	}

	if *column == "" {
		fmt.Println("'-c' must be specified")
		return
	}

	rand.Seed(time.Now().UnixNano())

	buildInsertStatement()
}

type LatLon struct {
	Lat int
	Lon int
}

func buildInsertStatement() {
	fmt.Printf("INSERT INTO \"%s\" (\"%s\") VALUES\n", *table, *column)
	for i := 0; i < *size; i++ {
		buildRecordValue()
		if i == (*size - 1) {
			fmt.Println(";")
		} else {
			fmt.Println(",")
		}
	}
}

func buildRecordValue() {
	ll1 := buildRandomLatLon()
	ll2 := buildRandomLatLon()
	ll3 := buildRandomLatLon()
	ll4 := buildRandomLatLon()
	fmt.Printf("\t(ST_PolygonFromText('POLYGON((%d %d, %d %d, %d %d, %d %d, %d %d))', 4326))", ll1.Lat, ll1.Lon, ll2.Lat, ll2.Lon, ll3.Lat, ll3.Lon, ll4.Lat, ll4.Lon, ll1.Lat, ll1.Lon)
}

func buildRandomLatLon() LatLon {
	return LatLon{
		Lat: randomInvert(random(0, 90)),
		Lon: randomInvert(random(0, 180)),
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomInvert(target int) int {
	if (rand.Intn(2) % 2) == 0 {
		return target
	} else {
		return -1 * target
	}
}
