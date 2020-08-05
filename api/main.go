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
	"tapera.mitraintegrasi/api/controller/bkn"
	"tapera.mitraintegrasi/api/controller/bri"
	_ "tapera.mitraintegrasi/api/docs"
	"tapera.mitraintegrasi/api/middleware"
	sbkn "tapera.mitraintegrasi/service/bkn"
	sbri "tapera.mitraintegrasi/service/bri"

	"tapera/conf"
	"tapera/util"
	"tapera/util/env"

	crbc "tapera.mitraintegrasi/grpc/client/cancelredemptionbri/v1"
	csbc "tapera.mitraintegrasi/grpc/client/cancelsubscribebri/v1"
	ppbc "tapera.mitraintegrasi/grpc/client/pendaftaranpesertabri/v1"
	pbc "tapera.mitraintegrasi/grpc/client/pesertabkn/v1"
	rbc "tapera.mitraintegrasi/grpc/client/redemptionbri/v1"
	rgpbc "tapera.mitraintegrasi/grpc/client/riwayatgolonganpesertabkn/v1"
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

	// create grpc connection pool
	rbGrpcAddr := env.Str(constant.EnvGrpcServerRedemptionBri, true, nil)
	rbGpcClientCnPool := createGrpcClientCnPool(rbGrpcAddr)
	defer rbGpcClientCnPool.Close()

	// create grpc client manager
	rbClientMgr := rbc.NewGrpcClientManager(rbGpcClientCnPool)

	// create grpc connection pool
	crbGrpcAddr := env.Str(constant.EnvGrpcServerCancelRedemptionBri, true, nil)
	crbGpcClientCnPool := createGrpcClientCnPool(crbGrpcAddr)
	defer crbGpcClientCnPool.Close()

	// create grpc client manager
	crbClientMgr := crbc.NewGrpcClientManager(crbGpcClientCnPool)

	//create bri controller
	bri.NewController(sbri.NewService(ppbClientMgr, csbClientMgr, sbClientMgr, rbClientMgr, crbClientMgr)).Route(r)

	// create grpc connection pool
	pbGrpcAddr := env.Str(constant.EnvGrpcServerPesertaBkn, true, nil)
	pbGpcClientCnPool := createGrpcClientCnPool(pbGrpcAddr)
	defer pbGpcClientCnPool.Close()

	// create grpc client manager
	pbClientMgr := pbc.NewGrpcClientManager(pbGpcClientCnPool)

	// create grpc connection pool
	rgpbGrpcAddr := env.Str(constant.EnvGrpcServerRiwayatGolonganPesertaBkn, true, nil)
	rgpbGpcClientCnPool := createGrpcClientCnPool(rgpbGrpcAddr)
	defer rgpbGpcClientCnPool.Close()

	// create grpc client manager
	rgpbClientMgr := rgpbc.NewGrpcClientManager(rgpbGpcClientCnPool)

	//create bkn controller
	bkn.NewController(sbkn.NewService(pbClientMgr, rgpbClientMgr)).Route(r)

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
