package testing 

import (
    "testing"
    "log"
    "fmt"
    "os"
    "github.com/jgoulah/go.fitbit"
)
//Test method

/*func main() {*/
func TestMain(*testing.T) {
	//Init config
	config := &fitbit.Config{
		false, //Debug
		false, //Disable SSL
	}

	//Initialise FitbitAPI
    fapi, err := fitbit.NewAPI("761d7f0836484d81999bfc1b3bc9c3a0", "b47420b3554642599267b080ea7e2759", config)
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

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

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
