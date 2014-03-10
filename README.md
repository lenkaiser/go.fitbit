# GoFitbit

Wrapper for the OAuth based REST API of Fitbit.com. For the full documentation please see <a href=""http://dev.fitbit.com">dev.fitbit.com</a>.

The development of this library is still in ALPHA so it still could be a bit buggy. So feel free to request merges and bug fix any problems in this repo.

# Usage

Before you can get started with the API you'll have to register your application at <a href="http://dev.fitbit.com">dev.fitbit.com</a> and obtain consumer key and secret for your application.

This library will handle the OAuth authorisation. To register your application you'll have to grant access to your application with your fitbit account. This can be done with the <a href="https://github.com/lenkaiser/gofitbit-client">gofitbit-client</a> repo.

```go
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
	fmt.Println("API initialised")

	//Add client
	client, err := fapi.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New client initialised")

	//Call methods from client
	client.getProfile()
	client.getRecentActivities()
}
```


# Changelog
- Version 0.1: 10 March 2014
 - Initial commit
 - Partial support for activities and profile
 - Protected setup so client only calls public methods
 - LICENSE and README.md files