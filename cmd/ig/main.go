package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/autogrow/go-jelly/ig"
	"github.com/penguinpowernz/go-ian/util/tell"
)

type creds struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func main() {
	credsFile := os.Getenv("HOME") + "/.intelligrow/creds"
	creds, err := readCreds(credsFile)
	if err != nil {
		initCreds(credsFile)
		tell.Fatalf("you need to enter your credentials in the file %s to continue", credsFile)
	}

	cl, err := ig.NewClient(creds.User, creds.Pass)
	if err != nil {
		tell.Fatalf("failed to create IG client: %s", err)
	}

	var listDevices, listGrowrooms bool
	var id, gr string
	var printReadings, fmtJSON bool
	flag.BoolVar(&listDevices, "l", false, "list known devices")
	flag.BoolVar(&listGrowrooms, "g", false, "list growrooms")
	flag.StringVar(&id, "id", "", "serial number to work with")
	flag.StringVar(&gr, "gr", "", "growroom name to work with")
	flag.BoolVar(&printReadings, "r", false, "print readings")
	flag.BoolVar(&fmtJSON, "json", false, "format as JSON")
	flag.Parse()

	switch {
	case listGrowrooms:
		fmt.Println("Growrooms:")
		for _, name := range cl.ListGrowrooms() {
			fmt.Printf("- %s", name)
		}

	case listDevices:
		fmt.Printf("%-12s %-18s %-10s %s", "Type", "ID", "Name", "Growroom")
		for _, d := range cl.Devices() {
			fmt.Printf("%-12s %-18s %-10s %s", d.Type, d.ID, d.DeviceName, d.Growroom)
		}

	case id != "" && printReadings:
		if doser, errd := cl.IntelliDoseBySerial(id); errd == nil {
			fmt.Printf("%-20s: %0.2f mS/cm²", "Nutrient", doser.Metrics.Ec)
			fmt.Printf("%-20s: %0.2f pH", "Acidity", doser.Metrics.PH)
			fmt.Printf("%-20s: %0.2f °C", "Water", doser.Metrics.NutTemp)
			return
		}

		if clim, errc := cl.IntelliClimateBySerial(id); errc == nil {
			fmt.Printf("%-20s: %0.2f °C", "Air Temp", clim.Metrics.AirTemp)
			fmt.Printf("%-20s: %0.2f %%H", "RH", clim.Metrics.Rh)
			fmt.Printf("%-20s: %0.2f kPa", "VPD", clim.Metrics.Vpd)
			fmt.Printf("%-20s: %0.2f ppm", "CO2", clim.Metrics.Co2)
			return
		}

		log.Fatalf("no device found with serial %s", id)

	case gr != "" && printReadings:
		room, ok := cl.Growroom(gr)
		if !ok {
			log.Fatalf("Growroom %s not found", gr)
		}

		if len(room.Devices.IntelliClimates) > 0 {
			fmt.Printf("%-20s: %0.2f °C", "Air Temp", room.Climate.AirTemp)
			fmt.Printf("%-20s: %0.2f %%H", "RH", room.Climate.RH)
			fmt.Printf("%-20s: %0.2f kPa", "VPD", room.Climate.VPD)
			fmt.Printf("%-20s: %0.2f ppm", "CO2", room.Climate.CO2)
		}

		if len(room.Devices.IntelliDoses) > 0 {
			fmt.Printf("%-20s: %0.2f mS/cm²", "Nutrient", room.Rootzone.EC)
			fmt.Printf("%-20s: %0.2f pH", "Acidity", room.Rootzone.PH)
			fmt.Printf("%-20s: %0.2f °C", "Water", room.Rootzone.Temp)
		}

	}
}

func initCreds(credsFile string) {
	data, err := json.Marshal(creds{})
	if err != nil {
		log.Fatalf("failed to encode empty creds file to put at %s", credsFile)
	}

	if err := os.Mkdir(filepath.Dir(credsFile), 0755); err != nil {
		log.Fatalf("failed to create creds file at %s: %s", credsFile, err)
	}

	if err := ioutil.WriteFile(credsFile, data, 0644); err != nil {
		log.Fatalf("failed to create creds file at %s: %s", credsFile, err)
	}
}

func readCreds(credsFile string) (creds, error) {
	creds := creds{}
	data, err := ioutil.ReadFile(credsFile)
	if err != nil {
		return creds, err
	}

	err = json.Unmarshal(data, &creds)
	if err != nil {
		log.Fatalf("failed to read creds file at %s: %s", credsFile, err)
	}

	return creds, nil
}