package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

// type coordinates struct {
// 	Longitude float64 `json:"longitude"`
// 	Latitude  float64 `json:"latitude"`
// 	Altitude  int64   `json:"altitude"`
// 	Timestamp int64   `json:"timestamp"`
// }

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

}

func getDuration(c *gin.Context) {

}

func getGeoData() (*geoFeatures, error) {
	file, err := os.Open("./path.geojso")
	if err != nil {
		return nil, err
	}

	var featuresCollection geoFeatures

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &featuresCollection)

	return &featuresCollection, nil
}
