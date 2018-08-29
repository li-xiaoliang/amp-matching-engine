package main

import "github.com/Proofsuite/amp-matching-engine/cmd"

func main() {
	cmd.Execute()

	// if err := app.LoadConfig("./config"); err != nil {
	// 	panic(fmt.Errorf("Invalid application configuration: %s", err))
	// }

	// if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
	// 	panic(fmt.Errorf("Failed to read the error message file: %s", err))
	// }

	// log.SetFlags(log.LstdFlags | log.Llongfile)
	// log.SetPrefix("\nLOG: ")
	// logger := logrus.New()

	// // connect to the database
	// if _, err := daos.InitSession(nil); err != nil {
	// 	panic(err)
	// }

	// http.Handle("/", NewRouter(logger))
	// http.HandleFunc("/socket", ws.ConnectionEndpoint)

	// // start the server
	// address := fmt.Sprintf(":%v", app.Config.ServerPort)
	// logger.Infof("server %v is started at %v\n", app.Version, address)
	// panic(http.ListenAndServe(address, nil))
}

// func NewRouter(logger *logrus.Logger) *routing.Router {
// 	router := routing.New()

// 	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
// 		c.Abort() // skip all other middlewares/handlers
// 		return c.Write("OK " + app.Version)
// 	})

// 	router.Use(
// 		app.Init(logger),
// 		content.TypeNegotiator(content.JSON),
// 		cors.Handler(cors.Options{
// 			AllowOrigins: "*",
// 			AllowHeaders: "*",
// 			AllowMethods: "*",
// 		}),
// 	)

// 	rg := router.Group("")

// 	rabbitmq.InitConnection(app.Config.Rabbitmq)
// 	ethereum.InitConnection(app.Config.Ethereum)
// 	redisClient := redis.InitConnection(app.Config.Redis)

// 	// instantiate engine
// 	engineResource, err := engine.InitEngine(redisClient)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// get daos for dependency injection
// 	orderDao := daos.NewOrderDao()
// 	tokenDao := daos.NewTokenDao()
// 	pairDao := daos.NewPairDao()
// 	tradeDao := daos.NewTradeDao()
// 	accountDao := daos.NewAccountDao()

// 	// get services for injection
// 	accountService := services.NewAccountService(accountDao, tokenDao)
// 	ohlcvService := services.NewOHLCVService(tradeDao)
// 	tokenService := services.NewTokenService(tokenDao)
// 	tradeService := services.NewTradeService(tradeDao)
// 	pairService := services.NewPairService(pairDao, tokenDao, engineResource, tradeService)
// 	orderService := services.NewOrderService(orderDao, pairDao, accountDao, tradeDao, engineResource)
// 	orderBookService := services.NewOrderBookService(pairDao, tokenDao, engineResource)
// 	cronService := crons.NewCronService(ohlcvService)
// 	// walletService := services.NewWalletService(walletDao, balanceDao)

// 	endpoints.ServeAccountResource(rg, accountService)
// 	endpoints.ServeTokenResource(rg, tokenService)
// 	endpoints.ServePairResource(rg, pairService)
// 	endpoints.ServeOrderBookResource(rg, orderBookService)
// 	endpoints.ServeOHLCVResource(rg, ohlcvService)
// 	endpoints.ServeTradeResource(rg, tradeService)
// 	endpoints.ServeOrderResource(rg, orderService, engineResource)

// 	//initialize rabbitmq subscriptions
// 	orderService.SubscribeQueue(engineResource.HandleOrders)
// 	engineResource.SubscribeResponseQueue(orderService.HandleEngineResponse)

// 	// fmt.Printf("\n%+v\n", app.Config.TickDuration)
// 	cronService.InitCrons()
// 	return router
// }

// rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
// rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
// 	SigningMethod: app.Config.JWTSigningMethod,
// 	TokenHandler:  apis.JWTHandler,
// }))
