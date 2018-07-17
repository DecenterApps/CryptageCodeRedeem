package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"
	"github.com/DecenterApps/CryptageCodeRedeem/endpoint"
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/dao"
	"github.com/DecenterApps/CryptageCodeRedeem/service"
	cryptageError "github.com/DecenterApps/CryptageCodeRedeem/error"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := cryptageError.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	db, err := dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	http.Handle("/", buildRouter(logger, db))

	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort()
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator("text/html"),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("")
	endpoint.ServeCouponResource(rg, service.NewCouponService(dao.NewCouponDAO()), service.NewCardService(dao.NewCardDAO(getCryptageCards())))

	return router
}

func getCryptageCards() app.Cryptage {
	var cryptage app.Cryptage
	cardsYaml, _ := ioutil.ReadFile(app.Config.Cards)
	yaml.Unmarshal(cardsYaml, &cryptage)

	for i := 0; i < len(cryptage.Cards); i++ {
		for j := 1; j <= len(cryptage.Cards[uint(i)]); j++ {
			card := cryptage.Cards[uint(i)][uint(j)]
			cryptage.Cards[uint(i)][card.Level] = card
		}

		delete(cryptage.Cards[uint(i)], uint(len(cryptage.Cards[uint(i)])-1))
	}

	return cryptage
}