package collectors

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/nestorneo/ecodata/config"
	"github.com/nestorneo/ecodata/contract"
	"go.mongodb.org/mongo-driver/mongo"
)

type Contaminacion struct {
	Cfg    config.Server
	client *mongo.Client
	audio  *mongo.Collection
	aire   *mongo.Collection
	co2    *mongo.Collection
	basura *mongo.Collection
}

// default collections
const (
	AUDIO  = "audio"
	CO2    = "co2"
	Aire   = "aire"
	Basura = "basura"
)

func NewCollector(cfg config.Server, client *mongo.Client) (*Contaminacion, error) {
	ct := &Contaminacion{
		cfg,
		client,
		nil,
		nil,
		nil,
		nil,
	}

	return ct, ct.Init()
}

func (collector *Contaminacion) Init() error {
	if collector.Cfg.DBAccess.Enable {
		collector.audio = collector.client.Database(collector.Cfg.DBAccess.DatabaseName).Collection(AUDIO)
		collector.co2 = collector.client.Database(collector.Cfg.DBAccess.DatabaseName).Collection(CO2)
		collector.aire = collector.client.Database(collector.Cfg.DBAccess.DatabaseName).Collection(Aire)
		collector.basura = collector.client.Database(collector.Cfg.DBAccess.DatabaseName).Collection(Basura)
	}
	return nil
}

func (collector *Contaminacion) Run() {

	log.Println(uuid.NewString(), "*************** AUDIO PROCESSOR **************")

	stangingDir := os.TempDir()

	if collector.Cfg.StagingArea != "" {
		stangingDir = collector.Cfg.StagingArea
	}

	files, err := ioutil.ReadDir(stangingDir)
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), collector.Cfg.PrefixForTempFile) && !file.IsDir() {
			log.Println("** *** * files to process are pal", file.Name())
			fullPath := filepath.Join(stangingDir, file.Name())
			fmt.Println(fullPath)
			record, err := contract.InspectBundle(fullPath)
			// TODO delete files with errors for security
			if err != nil {
				log.Println("error processing audio record, know issue")
				log.Println(err)

				parts := strings.Split(file.Name(), "_")
				if len(parts) == 3 {
					collector.basura.InsertOne(
						context.TODO(), map[string]any{
							"request_id": parts[1],
							"filename":   fullPath,
							"error":      err.Error(),
						},
					)
				}

				os.Remove(fullPath)
				continue
			}

			if collector.Cfg.DBAccess.Enable && collector.client != nil {
				switch record.Tipo {
				case "AUDIO":
					_, err = collector.audio.InsertOne(context.TODO(), record)

					break
				case "CO2":
					_, err = collector.co2.InsertOne(context.TODO(), record)
					break
				default:
					log.Println("SABEEEE ")
				}
				if err != nil {
					log.Println("db error", err)
				}
			}

			fmt.Println(record.Name, record.Dispositivo, record.Inicio, record.Fin, "  ", len(record.Content))
			os.Remove(fullPath)
			log.Println(fullPath, "Procesado...")
		}
	}
	log.Println("audio collector is done")
}
