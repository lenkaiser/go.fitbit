package main

import (
    "log"
    "fmt"
    "github.com/jgoulah/go.fitbit"
	/*"github.com/mrjones/oauth"*/
    "os"
)

func main() {

	config := &fitbit.Config{
		false, //Debug
		false, //Disable SSL
	}

	//Initialise FitbitAPI
	/*fapi, err := NewAPI("761d7f0836484d81999bfc1b3bc9c3a0", "b47420b3554642599267b080ea7e2759", config)*/
	fapi, err := fitbit.NewAPI("3029d5e2f3de4f42858e68a33ff63fe6", "e3e3b4043b1748eeb56b7644611fc950", config)
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

    /*profile,err := client.GetProfile()*/
    activities,err := client.GetRecentActivities()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Printf("%#v\n", activities)

/*
	config := &fitbit.Config{
		false, //Debug
		false, //Disable SSL
	}
	fapi, err := fitbit.NewAPI("3029d5e2f3de4f42858e68a33ff63fe6", "e3e3b4043b1748eeb56b7644611fc950", config)
	if err != nil {
		log.Fatal(err)
	}

	oauthConsumer := oauth.NewConsumer(fapi.applicationToken, fapi.applicationSecret, oauth.ServiceProvider{
		RequestTokenUrl:   fapi.requestURL,
		AuthorizeTokenUrl: fapi.authURL,
		AccessTokenUrl:    fapi.accessURL,
	})
	if oauthConsumer == nil {
		return nil, errors.New("failed to create new consumer for FitbitAPI")
	}


	// //Get request tokenURL
	requestToken, url, err := oauthConsumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		return nil, err
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Enter the verification code: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := oauthConsumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
	    return nil, err
	}
	c.accessToken = accessToken
	c.userID = c.accessToken.AdditionalData["encoded_user_id"]

    log.Printf("accessToken >> %s\n", accessToken.Token)
	log.Printf("accessSecret >> %s\n", accessToken.Secret)
	log.Printf("userID >> %s\n", c.userID)
*/
}
