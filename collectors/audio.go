package collector

import (
	"log"

	"github.com/nestorneo/ecodata/config"
)

type Audio struct {
	localConfig config.Server
}

func (collector *Audio) Run() {
	log.Println("*************** AUDIO PROCESSOR **************")
}
