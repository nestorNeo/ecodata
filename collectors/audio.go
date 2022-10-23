package collectors

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/nestorneo/ecodata/config"
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
		}
	}
}
