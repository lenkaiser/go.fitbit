package fitbit

import (
	"errors"
	"fmt"
	"github.com/mrjones/oauth"
    "log"
    "io/ioutil"
    "os"
    "strings"
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
    c.oc = oauthConsumer

    var oauth_creds_file string = fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".fitbit-oauth")
    if _, err := os.Stat(oauth_creds_file); err == nil {
        contents,_ := ioutil.ReadFile(oauth_creds_file)
        keys := strings.Split(string(contents), "\n");
        c.accessToken = &oauth.AccessToken{
            Token:  keys[0],
            Secret: keys[1],
        }
    } else {

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

        log.Printf("accessToken >> %s\n", accessToken.Token)
        log.Printf("accessSecret >> %s\n", accessToken.Secret)

        contents := fmt.Sprintf("%s\n%s", accessToken.Token, accessToken.Secret);
        ioutil.WriteFile(oauth_creds_file, []byte(contents), 0x777)

        c.accessToken = &oauth.AccessToken{
            Token:  accessToken.Token,
            Secret: accessToken.Secret,
        }
    }

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

