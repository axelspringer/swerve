package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	phm "github.com/axelspringer/swerve/prometheus"

	"github.com/axelspringer/swerve/acm"
	"github.com/axelspringer/swerve/api"
	"github.com/axelspringer/swerve/cache"
	"github.com/axelspringer/swerve/config"
	"github.com/axelspringer/swerve/database"
	hp "github.com/axelspringer/swerve/http"
	"github.com/axelspringer/swerve/https"
	"github.com/axelspringer/swerve/log"
	"github.com/axelspringer/swerve/model"
	"github.com/pkg/errors"
)

// NewApplication creates a new instance
func NewApplication() *Application {
	return &Application{
		Config: config.NewConfiguration(),
	}
}

// Setup sets up the application
func (a *Application) Setup() error {
	err := a.Config.FromEnv()
	if err != nil {
		return errors.WithMessage(err, ErrConfigInvalid)
	}
	a.Config.FromParameter()
	err = a.Config.Validate()
	if err != nil {
		return errors.WithMessage(err, ErrConfigInvalid)
	}

	log.SetupLogger(a.Config.LogLevel, a.Config.LogFormatter)

	db, err := database.NewDatabase(a.Config.Database)
	if err != nil {
		return errors.WithMessage(err, ErrDatabaseServiceCreate)
	}
	if a.Config.Bootstrap {
		err = db.Prepare()
		if err != nil {
			log.Errorf("error while preparing the dynamodb %v", err)
			return errors.WithMessage(err, ErrTablePrepare)
		}
	}

	pro := phm.NewPHM()
	a.Cache = cache.NewCache(db)

	controlModel := model.NewModel(a.Cache)

	autocertManager := acm.NewACM(a.Cache.AllowHostPolicy,
		a.Cache,
		a.Config)

	a.HTTPServer = hp.NewHTTPServer(controlModel.GetRedirectByDomain,
		autocertManager.HTTPHandler,
		a.Config.HTTPListenerPort,
		pro.WrapHandler)

	a.HTTPSServer = https.NewHTTPSServer(controlModel.GetRedirectByDomain,
		autocertManager.GetCertificate,
		a.Config.HTTPSListenerPort,
		pro.WrapHandler)
	a.APIServer = api.NewAPIServer(controlModel, a.Config.API, pro.WrapHandler)

	return nil
}

// Run starts the application
func (a *Application) Run() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	if err := a.Cache.Observe(a.Config.CacheInterval); err != nil {
		log.Fatal(errors.WithMessage(err, ErrCacheObserver))
	}

	go func() {
		log.Fatal(a.APIServer.Listen())
	}()

	go func() {
		log.Fatal(a.HTTPServer.Listen())
	}()

	go func() {
		log.Fatal(a.HTTPSServer.Listen())
	}()

	time.Sleep(time.Second * 1)

	url := fmt.Sprintf("http://localhost:%d", a.Config.HTTPListenerPort)
	_, err := http.Get(url)
	if err != nil {
		log.Error(err)
	}

	<-sigchan
}
