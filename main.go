package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Location represents data about a location.
type Location struct {
	UID       int64  `json:"uid"`
	UserUID   int64  `json:"useruid"`
	CreatedOn string `json:"createdon"`
	Geom      string `json:"geom"`
}

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "jbriant", "dev")
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	return db
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// getLocation responds with the list of all locations as JSON.
func GetLocations(c *gin.Context) {
	db := setupDB()

	printMessage("Getting locations...")

	// Get all locations from location table
	rows, err := db.Query("SELECT uid AS uid, user_uid::text as useruid, created_on as createdon, st_astext(geom) as geom FROM location")

	// check errors
	if err != nil {
		panic(err)
	}

	// var response []JsonResponse
	var locations []Location

	// Foreach location
	for rows.Next() {
		var uid int64
		var useruid int64
		var createdon string
		var geom string

		err = rows.Scan(&uid, &useruid, &createdon, &geom)

		// check errors
		if err != nil {
			panic(err)
		}

		locations = append(locations, Location{UID: uid, UserUID: useruid, CreatedOn: createdon, Geom: geom})
	}

	c.IndentedJSON(http.StatusOK, locations)
}

// getLocation responds with the list of all locations as JSON.
func GetLocationById(c *gin.Context) {
	id := c.Param("locationId")

	db := setupDB()

	printMessage("Getting location...")

	var uid int64
	var useruid int64
	var createdon string
	var geom string

	// Get location matching locationId
	getlocation := `SELECT uid AS uid, 
					user_uid as useruid, 
					created_on as createdon, 
					st_astext(geom) as geom 
					FROM location
					WHERE uid=$1`

	err := db.QueryRow(getlocation, id).Scan(&uid, &useruid, &createdon, &geom)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	location := Location{UID: uid, UserUID: useruid, CreatedOn: createdon, Geom: geom}

	c.IndentedJSON(http.StatusOK, location)
}

func AddLocation(c *gin.Context) {
	var newLocation Location

	db := setupDB()

	printMessage("Inserting location...")

	if err := c.BindJSON(&newLocation); err != nil {
		return
	}

	// Get location matching locationId
	insertLocation := `INSERT INTO location(user_uid, created_on, geom)
					   VALUES ($1::integer, now(), st_geomfromtext($2))`

	_, err := db.Exec(insertLocation, newLocation.UserUID, newLocation.Geom)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, newLocation)
}

func UpdateLocation(c *gin.Context) {
	var updatedLocation Location
	id := c.Param("locationId")

	db := setupDB()

	printMessage("Updating location...")

	if err := c.BindJSON(&updatedLocation); err != nil {
		return
	}

	// Update location matching locationId
	insertLocation := `UPDATE location
					   SET geom=$2
					   WHERE uid=$1`

	_, err := db.Exec(insertLocation, id, updatedLocation.Geom)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, updatedLocation)

}

func DeleteLocation(c *gin.Context) {
	id := c.Param("locationId")

	db := setupDB()

	printMessage("Deleting location...")

	// Delete location matching locationId
	deleteLocation := `DELETE FROM location
					   WHERE uid=$1`

	res, err := db.Exec(deleteLocation, id)
	if err != nil {
		panic(err)
	}

	// verifying if row was deleted
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, count)
}

func main() {
	router := gin.Default()
	router.GET("/locations", GetLocations)
	router.POST("/locations", AddLocation)
	router.GET("/locations/:locationId", GetLocationById)
	router.PUT("/locations/:locationId", UpdateLocation)
	router.DELETE("/locations/:locationId", DeleteLocation)
	router.Run("localhost:8080")
}
