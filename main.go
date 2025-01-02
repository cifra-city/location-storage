package location_storage

import (
	"os"

	"github.com/cifra-city/location-storage/internal /cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}