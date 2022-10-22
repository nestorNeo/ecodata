/*
 * ESTE ES EL BUENO !!!! No sobreescribir
 * ecodata - OpenAPI 3.0
 *
 * API to store contamination events and data from different sources
 *
 * API version: 1.0.0
 * Contact: nestor.salvador@gdl.cinvestav.mx
 */

package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CreateBundle - Upload bundle from device
func CreateBundle(c *gin.Context) {
	log.Println("got a new packet")
	fmt.Printf("Raw Request Body : %s", c.Request.Body)
	if err := processBundle(c); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"mensaje": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetBundleStatus - Get user by user name
func GetBundleStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func processBundle(c *gin.Context) error {
	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return err
	}

	if len(data) == 0 {
		return errors.New("empty file please send a valid file")
	}

	// spawn a processing go-routine
	go func(raw []byte) {
		log.Println("still working in backend pal ....")
		f, err := os.CreateTemp("", "audioproc")
		defer f.Close()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("new file at", f.Name())
		n, err := f.Write(raw)

		if err != nil {
			log.Println(err)
			return
		}
		log.Println("written bytes n", n)
	}(data)

	return nil
}
