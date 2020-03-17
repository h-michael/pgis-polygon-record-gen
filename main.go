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
	Lat float32
	Lon float32
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
	ll1, ll2, ll3, ll4 := buildFourPoints()
	fmt.Printf("\t(ST_PolygonFromText('POLYGON((%f %f, %f %f, %f %f, %f %f, %f %f))', 4326))", ll1.Lat, ll1.Lon, ll2.Lat, ll2.Lon, ll3.Lat, ll3.Lon, ll4.Lat, ll4.Lon, ll1.Lat, ll1.Lon)
}

func buildFourPoints() (LatLon, LatLon, LatLon, LatLon) {
	ll1 := buildBaseRandomLatLon()
	offset := randomOffset()
	ll2 := LatLon{
		Lat: ll1.Lat + offset,
		Lon: ll1.Lon,
	}

	ll3 := LatLon{
		Lat: ll1.Lat + offset,
		Lon: ll1.Lon + offset,
	}

	ll4 := LatLon{
		Lat: ll1.Lat,
		Lon: ll1.Lon + offset,
	}

	return ll1, ll2, ll3, ll4
}

func buildBaseRandomLatLon() LatLon {
	return LatLon{
		Lat: randomInvert(random(0, 89)),
		Lon: randomInvert(random(0, 179)),
	}
}

func randomOffset() float32 {
	return random(1, 999) / 1000
}

func random(min, max int) float32 {
	return float32(rand.Intn(max-min) + min)
}

func randomInvert(target float32) float32 {
	if (rand.Intn(2) % 2) == 0 {
		return target
	} else {
		return -1 * target
	}
}
