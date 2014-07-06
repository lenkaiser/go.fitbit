package main

import (
    "log"
    "fmt"
    "os"
    ".." // local fitbit api
)

// this example exists to generate a token since fmt.Scanln 
// isn't working in the tests this should be run first
func main() {

	config := &fitbit.Config{
		false, //Debug
		false, //Disable SSL
	}

	//Initialise FitbitAPI
    client_key    := os.Getenv("FITBIT_CLIENT_KEY")
    client_secret := os.Getenv("FITBIT_CLIENT_SECRET")
    if ((client_key == "") || (client_secret == "")) {
        log.Fatal("Please export FITBIT_CLIENT_KEY and FITBIT_CLIENT_SECRET")
    }
    fapi, err := fitbit.NewAPI(client_key, client_secret, config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FitbitAPI initialised")

	//Add client
	_, err = fapi.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New client initialised")

}
