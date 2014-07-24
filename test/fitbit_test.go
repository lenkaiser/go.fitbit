package testing 

import (
    "testing"
    "log"
    "fmt"
    "os"
    ".." // local fitbit api
)
//Test method

func TestMain(*testing.T) {
	//Init config
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
	client, err := fapi.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New client initialised")

    profile,err := client.GetProfile()
    /*activities,err := client.GetRecentActivities()*/

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    /*fmt.Printf("%#v\n", activities)*/
    fmt.Printf("name: %#v\n", profile.User.FullName)

	// log.Printf("LOG BODY MEASUREMENTS")
	// _, err = client.LogBodyMeasurements(time.Time, bicep, calf, chest, fat, forearm, hips, neck, thigh, waist, weight)
	// if err != nil {
	// 	log.Printf("measurement error: %s", err)
	// }

	// log.Printf("LOG BODY WEIGHT")
	// weightData, err := client.LogBodyWeight(time.Now(), 64)
	// if err != nil {
	// 	log.Printf("weight error: %s", err)
	// } else {
	// 	log.Printf("DELETE BODY WEIGHT")

	// 	err = client.DeleteBodyWeight(weightData.WeightLog.LogID)
	// 	if err != nil {
	// 		log.Printf("delete weight: %s", err)
	// 	}
	// }

	// log.Printf("LOG BODY FAT")
	// fatData, err := client.LogBodyFat(time.Now(), 14)
	// if err != nil {
	// 	log.Printf("fat error: %s", err)
	// } else {
	// 	log.Printf("DELETE BODY WEIGHT")

	// 	err = client.DeleteBodyFat(fatData.FatLog.LogID)
	// 	if err != nil {
	// 		log.Printf("delete fat: %s", err)
	// 	}
	// }
}
