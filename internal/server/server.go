package server

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	earthRadius = 6371
)

type geoFeatures struct {
	Type     string       `json:"type"`
	Features []geoFeature `json:"features"`
}

type geoFeature struct {
	Type       string             `json:"type"`
	Properties geoFeatureProperty `json:"properties"`
	Geometry   geoFeatureGeometry `json:"geometry"`
}

type geoFeatureProperty struct{}

type geoFeatureGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type coordinates struct {
	Longitude float64
	Latitude  float64
	Altitude  float64
	Timestamp float64
}

func Run() {
	r := gin.Default()

	r.GET("/getPath", getPath)
	r.GET("/getDuration", getDuration)
	r.GET("/getDistance", getDistance)

	r.Run(":8080")
}

func getPath(c *gin.Context) {
	var featuresCollection *geoFeatures

	featuresCollection, err := getGeoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, featuresCollection)
}

func getDistance(c *gin.Context) {
	var totalDist float64

	data, err := getGeoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": http.StatusInternalServerError})
		return
	}

	var pointA coordinates
	var pointB coordinates

	for _, feature := range data.Features {
		for i := 0; i < len(feature.Geometry.Coordinates); i++ {
			if i == 0 {
				pointA = coordinates{
					Longitude: feature.Geometry.Coordinates[i][0],
					Latitude:  feature.Geometry.Coordinates[i][1],
				}
				continue
			}
			pointB = coordinates{
				Longitude: feature.Geometry.Coordinates[i][0],
				Latitude:  feature.Geometry.Coordinates[i][1],
			}

			totalDist += getDistanceBetweenPoints(pointA, pointB)

			pointA = pointB
		}
	}

	c.JSON(http.StatusOK, gin.H{"totalDistance": totalDist})
}

func getDuration(c *gin.Context) {

}

func getGeoData() (*geoFeatures, error) {
	file, err := os.Open("./path.geojson")
	if err != nil {
		return nil, err
	}

	var featuresCollection geoFeatures

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &featuresCollection)

	return &featuresCollection, nil
}

func getDistanceBetweenPoints(pointA coordinates, pointB coordinates) float64 {
	lonA := degToRad(pointA.Longitude)
	latA := degToRad(pointA.Latitude)
	lonB := degToRad(pointB.Longitude)
	latB := degToRad(pointB.Latitude)

	deltaLat := latB - latA
	deltaLon := lonB - lonA

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(latA)*math.Cos(latB)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRadius
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}
