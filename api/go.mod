module tapera.mitraintegrasi/api

go 1.14

replace tapera/util => ../internal/util

replace tapera/conf => ../internal/conf

replace tapera.mitraintegrasi/service => ../internal/service

replace tapera.mitraintegrasi/grpc/client => ../internal/grpc/client

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/validator/v10 v10.3.0
	github.com/google/go-tpm v0.2.0 // indirect
	github.com/googleapis/gax-go v1.0.3 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/schema v1.1.0
	github.com/hashicorp/vault/api v1.0.4 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/processout/grpc-go-pool v1.2.1
	github.com/rs/zerolog v1.19.0
	github.com/salrashid123/oauth2 v0.0.0-20200503195646-e37a24dfdeb3
	github.com/satori/go.uuid v1.2.0
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.3
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/tools v0.0.0-20200626171337-aa94e735be7f // indirect
	google.golang.org/api v0.28.0 // indirect
	google.golang.org/grpc v1.29.1
	tapera.mitraintegrasi/grpc/client v1.0.0
	tapera.mitraintegrasi/service v1.0.0
	tapera/conf v1.0.0
	tapera/util v1.0.0
)
