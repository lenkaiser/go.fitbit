package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mrjones/oauth"
	"log"
)

const (
	AUTH_HOST = "www.fitbit.com"
	API_HOST  = "api.fitbit.com"
)

type Fitbit struct {
	config            *Config
	applicationToken  string
	applicationSecret string
	authURL           string
	requestURL        string
	accessURL         string
	apiURL            string
}

// Config holds options for the API that are not mandatory, but can optionally be changed
type Config struct {
	// Debug, when true, will enable debug messages written to Stdout
	Debug bool
	// DisableSSL, when true, reverts the used schema to HTTP (instead of HTTPS)
	DisableSSL bool
}

// DefaultConfig is used when config passed to NewAPI(..) is nil.
var DefaultConfig = &Config{
	Debug:      false,
	DisableSSL: false,
}

func NewAPI(token string, secret string, config *Config) (*Fitbit, error) {
	if config == nil {
		config = DefaultConfig
	}

	// conn type http or https
	connType := "https"
	if config.DisableSSL {
		connType = "http"
	}

	// create new API instance
	fapi := &Fitbit{
		config: config,

		applicationToken:  token,
		applicationSecret: secret,

		authURL:    fmt.Sprintf("%s://"+AUTH_HOST+"/oauth/authorize", connType),
		requestURL: fmt.Sprintf("%s://"+AUTH_HOST+"/oauth/request_token", connType),
		accessURL:  fmt.Sprintf("%s://"+AUTH_HOST+"/oauth/access_token", connType),
		apiURL:     fmt.Sprintf("%s://"+API_HOST+"/1/", connType),
	}
	return fapi, nil
}

// Client holds the unique information for a single end-user.
// It provides methods to interact with the Fitbit servers and send/retrieve information concerning the given end-user.
type Client struct {
	api         *Fitbit
	oc          *oauth.Consumer
	accessToken *oauth.AccessToken
	userID      string
}

// NewClient creates a new client
func (api *Fitbit) NewClient() (*Client, error) {
	c := &Client{
		api: api,
	}

	// initiate OAuth-consumer
	oauthConsumer := oauth.NewConsumer(api.applicationToken, api.applicationSecret, oauth.ServiceProvider{
		RequestTokenUrl:   api.requestURL,
		AuthorizeTokenUrl: api.authURL,
		AccessTokenUrl:    api.accessURL,
	})
	if oauthConsumer == nil {
		return nil, errors.New("failed to create new consumer for FitbitAPI")
	}

	//Set OAuth consumer to debug
	oauthConsumer.Debug(api.config.Debug)

	//Add to client
	//TODO: Replace this with DB values
	c.oc = oauthConsumer
	c.accessToken = &oauth.AccessToken{
		Token:  "cd7c9a14b886e46a3bb597a15ba32684",
		Secret: "bf170c1442a005952c4747eb0146b149",
		AdditionalData: map[string]string{
			"encoded_user_id": "2DXGXY",
		},
	}

	// //Get request tokenURL
	// requestToken, url, err := oauthConsumer.GetRequestTokenAndUrl("oob")
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("(1) Go to: " + url)
	// fmt.Println("(2) Enter the verification code: ")

	// verificationCode := ""
	// fmt.Scanln(&verificationCode)

	// accessToken, err := oauthConsumer.AuthorizeToken(requestToken, verificationCode)
	// if err != nil {
	// 	return nil, err
	// }
	// c.accessToken = accessToken
	// c.userID = c.accessToken.AdditionalData["encoded_user_id"]

	// log.Printf("accessToken >> %s\n", accessToken.Token)
	// log.Printf("accessSecret >> %s\n", accessToken.Secret)
	// log.Printf("userID >> %s\n", c.userID)

	return c, nil
}

//Set OAuth details
func (api *Fitbit) setOAuthDetails(token, secret string) error {
	if api == nil {
		return errors.New("FitbitAPI is nil")
	}

	api.applicationToken = token
	api.applicationSecret = secret

	return nil
}

//Get OAuth token
func (api *Fitbit) getOAuthToken() (string, error) {
	if api == nil {
		return "", errors.New("FitbitAPI is nil")
	}

	return api.applicationToken, nil
}

//Get OAuth secret
func (api *Fitbit) getOAuthSecret() (string, error) {
	if api == nil {
		return "", errors.New("FitbitAPI is nil")
	}

	return api.applicationSecret, nil
}

//Set userID
func (c *Client) setUserID(userID string) error {
	if c == nil {
		return errors.New("Client is nil")
	}

	//Set userID
	c.userID = userID

	return nil
}

//Test method

func main() {
	//Init config
	config := &Config{
		false, //Debug
		false, //Disable SSL
	}

	//Initialise FitbitAPI
	fapi, err := NewAPI("761d7f0836484d81999bfc1b3bc9c3a0", "b47420b3554642599267b080ea7e2759", config)
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

}
