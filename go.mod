module GMS

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.15
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.16.0
	github.com/micro/go-micro/v2 v2.9.1-0.20200723075038-fbdf1f2c1c4c
	github.com/micro/go-plugins v0.22.0
	go.uber.org/zap v1.15.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d
