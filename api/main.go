package main

import (
	"fmt"
	"time"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	pool "github.com/processout/grpc-go-pool"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"tapera.mitraintegrasi/api/constant"
	"tapera.mitraintegrasi/api/controller/bri"
	_ "tapera.mitraintegrasi/api/docs"
	"tapera.mitraintegrasi/api/middleware"
	sbri "tapera.mitraintegrasi/service/bri"

	"tapera/conf"
	"tapera/util"
	"tapera/util/env"

	csbc "tapera.mitraintegrasi/grpc/client/cancelsubscribebri/v1"
	ppbc "tapera.mitraintegrasi/grpc/client/pendaftaranpesertabri/v1"
	sbc "tapera.mitraintegrasi/grpc/client/subscriptionbri/v1"
)

// @title Tapera API
// @version v1.0.0
// @description This is Tapera API listing descriptions.
// @termsOfService http://Tapera.org/terms/

// @contact.name Tapera API Support
// @contact.url http://www.Tapera.org/support
// @contact.email support@Tapera.org

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// load env variable if the is .env file exists
	if util.FileExists(".env") {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	// create appconfig
	appConf := conf.EnvConfig()
	conf.AppConf = appConf

	// create factory based on env
	envf := appConf.NewEnvFactory()

	// create logger
	logger := envf.Logger()

	// create db connection
	// db := envf.CockcroachDb()
	// defer db.Close()

	// create mux router
	r := mux.NewRouter()

	// add middlewares
	r.Use(handlers.CompressHandler,
		middleware.AddRequestID,
		middleware.Logger(logger),
		middleware.ReqLogger(),
	)

	// register swagger route
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// create grpc connection pool
	ppbGrpcAddr := env.Str(constant.EnvGrpcServerPendaftaranPeserta, true, nil)
	ppbGrpcClientCnPool := createGrpcClientCnPool(ppbGrpcAddr)
	defer ppbGrpcClientCnPool.Close()

	// create grpc client manager
	ppbClientMgr := ppbc.NewGrpcClientManager(ppbGrpcClientCnPool)

	// create grpc connection pool
	csbGrpcAddr := env.Str(constant.EnvGrpcServerCancelSubscribeBri, true, nil)
	csbGpcClientCnPool := createGrpcClientCnPool(csbGrpcAddr)
	defer csbGpcClientCnPool.Close()

	// create grpc client manager
	csbClientMgr := csbc.NewGrpcClientManager(csbGpcClientCnPool)

	// create grpc connection pool
	sbGrpcAddr := env.Str(constant.EnvGrpcServerSubscriptionBri, true, nil)
	sbGpcClientCnPool := createGrpcClientCnPool(sbGrpcAddr)
	defer sbGpcClientCnPool.Close()

	// create grpc client manager
	sbClientMgr := sbc.NewGrpcClientManager(sbGpcClientCnPool)

	//create bri controller
	bri.NewController(sbri.NewService(ppbClientMgr, csbClientMgr, sbClientMgr)).Route(r)

	// run http server
	logger.Info().Msgf("server is listening to port %d", appConf.Port())
	if err := http.ListenAndServe(fmt.Sprintf(":%d", appConf.Port()), handlers.RecoveryHandler()(r)); err != nil {
		panic(err)
	}
	logger.Info().Msg("exit")
}

func createGrpcClientCnPool(addr string) *pool.Pool {
	// create grpc connection pool see https://github.com/processout/grpc-go-pool
	maxCn := env.Int(constant.EnvGrpcClientMaxOpenConn, false, 10)
	maxIddleCn := env.Int(constant.EnvGrpcClientMaxIdleConn, false, 0)
	idleTimeOut := env.Int(constant.EnvGrpcClientIdleTimeoutSec, false, 10)
	cert := env.Str(constant.EnvGrpcClientCert, false, nil)

	if len(cert) != 0 {
		creds, err := credentials.NewClientTLSFromFile(cert, "")
		if err != nil {
			panic(err)
		}

		gpool, err := pool.New(func() (*grpc.ClientConn, error) {
			return grpc.Dial(addr, grpc.WithTransportCredentials(creds))
		}, maxIddleCn, maxCn, time.Duration(idleTimeOut)*time.Second)

		if err != nil {
			panic(err)
		}
		return gpool
	}

	gpool, err := pool.New(func() (*grpc.ClientConn, error) {
		return grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	}, maxIddleCn, maxCn, time.Duration(idleTimeOut)*time.Second)

	if err != nil {
		panic(err)
	}
	return gpool
}
