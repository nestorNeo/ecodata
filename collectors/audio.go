package collectors

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/nestorneo/ecodata/config"
	"github.com/nestorneo/ecodata/contract"
)

type Audio struct {
	Cfg config.Server
}

func (collector *Audio) Run() {
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
				os.Remove(fullPath)
				continue
			}

			fmt.Println(record.Name, record.Dispositivo, record.Inicio, record.Fin, "  ", len(record.Content))
			os.Remove(fullPath)
			log.Println(fullPath, "Procesado...")
		}
	}
	log.Println("audio collector is done")
}
