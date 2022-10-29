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
	ginsession "github.com/go-session/gin-session"
)

const (
	UUID        = "uuid"
	BUNDLEPARAM = "bundleId"
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

	id, _ := c.Get("uuid")

	c.JSON(http.StatusOK,
		&Bundle{
			Id:           id.(string),
			BundleStatus: "ACCEPTED_TO_PROCESS",
		})
}

// GetBundleStatus - Get user by user name
func GetBundleStatus(c *gin.Context) {

	// CURRENT WORKAROUND PENDING TO ADD DB just using context by now

	targetId := c.Param(BUNDLEPARAM)
	store := ginsession.FromContext(c)
	content, found := store.Get(targetId)
	fileLocation, foundFileName := store.Get(targetId + "_bundle_location")

	// if not found in cached then it should be in the db more expensive. request
	if !found || !foundFileName {
		c.JSON(
			http.StatusNotFound, &Bundle{
				Id:           targetId,
				BundleStatus: "NOT_FOUND",
			},
		)
		return
	}

	fileLocationStr := fmt.Sprintf("%s", fileLocation)
	if _, err := os.Stat(fileLocationStr); errors.Is(err, os.ErrNotExist) {
		store.Set(targetId, "PROCCESED")
		store.Save()
	}

	c.JSON(http.StatusOK,
		&Bundle{
			Id:           targetId,
			BundleStatus: fmt.Sprintf("%s", content),
		},
	)
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
	go func(raw []byte, c *gin.Context) {
		id, _ := c.Get(UUID)
		log.Println("still working in backend pal ....")
		f, err := os.CreateTemp(
			c.MustGet("SERVER_STAGING_TMP_FILES").(string),
			c.MustGet("SERVER_PREFIX_TMP_FILES").(string)+fmt.Sprintf("_%s_", id))
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		log.Println("new file at", f.Name())
		n, err := f.Write(raw)

		if err != nil {
			log.Println(err)
			return
		}
		log.Println("written bytes n", n)
		// identifies this request id at server level
		store := ginsession.FromContext(c)

		// ALL IN CACHE ARE BEING USED FOR LATER PROCESSING
		store.Set(fmt.Sprintf("%s", id), "LOCAL_STORED_NOT_PROCESSED")
		store.Set(fmt.Sprintf("%s_bundle_location", id), f.Name())
		if err = store.Save(); err != nil {
			log.Println(err)
		}

	}(data, c.Copy())

	return nil
}
