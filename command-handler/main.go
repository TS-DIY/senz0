package main

import (
	"log"
	"net/http"
	"os"

	// if using environment to configure app
	// "github.com/kelseyhightower/envconfig"

	"gitlab.com/norzion/temp0/command-handler/app"
	"gitlab.com/norzion/temp0/command-handler/datastore"
	"gitlab.com/norzion/temp0/command-handler/handlers"
)

// Config struct containing all default service settings
type Config struct {
	Name       string
	ListenPort string
}

func main() {
	log.Println("starting service")

	config := Config{
		Name:       "ExampleService",
		ListenPort: "8000",
	}

	/***********************************
	 * if using environment to configure service - prepended string in process governs how to configure the env vars
	 ***********************************/
	// var config Config
	// if err := envconfig.Process("NORZION_SERVICE", &config); err != nil {
	// 	log.Fatal(err.Error())
	// }

	/***********************************
	 * loading config from a json file
	 ***********************************/
	// configFile, err := os.Open("configfile")
	// defer configFile.Close()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// json.NewDecoder(configFile).Decode(&config)

	/***********************************
	 * Setup datastore(s)
	 ***********************************/
	// if using a psql based datastore
	// ds, err := datastore.NewPSQLStore("postgres://norzion_testuser:flafgiraf@localhost/norzion_skel_service", 5)
	// if err != nil {
	// 	log.Fatalln("Failed to connect to datastore")
	// }

	ds := datastore.NewInMemoryStore()

	/***********************************
	 * Setup the app (essentially just a wrapper to avoid multiple function inputs in dependencies)
	 ***********************************/
	app := &app.App{
		Log:       log.New(os.Stdout, config.Name, log.Ldate|log.Ltime|log.Lshortfile),
		Datastore: ds,
	}

	/***********************************
	 * Setup server handler(s)
	 ***********************************/
	// if using http
	h, err := handlers.NewHTTPHandlerWithPanicRecovery(app)

	if err != nil {
		log.Fatal(err)
	}

	/***********************************
	 * Serve
	 ***********************************/
	log.Println(http.ListenAndServe(":"+config.ListenPort, h))
}
