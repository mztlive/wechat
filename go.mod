module github.com/mztlive/wechat

go 1.19

replace (
	github.com/mztlive/foundation => ../foundation
	github.com/mztlive/logger => ../logger
	github.com/mztlive/utils => ../utils

)

require (
	github.com/golang/mock v1.6.0
	github.com/gookit/config/v2 v2.1.6
	github.com/mztlive/logger v0.0.0-00010101000000-000000000000
	github.com/silenceper/wechat/v2 v2.1.3
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	go.uber.org/zap v1.23.0
)

require (
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-redis/redis/v8 v8.11.6-0.20220405070650-99c79f7041fc // indirect
	github.com/gookit/goutil v0.5.12 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/tidwall/gjson v1.14.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.7 // indirect
)
