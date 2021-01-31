package commands

import (
	"github.com/d2r2/go-dht"
	"fmt"
	"log"
	models "models"
)


func Get_TEMP_HUMITY() models.Dht_Info {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",temperature, humidity, retried)
	dht_info := models.Dht_Info{}
	dht_info.Temp = temperature
	dht_info.Humity = humidity

	return dht_info

}
