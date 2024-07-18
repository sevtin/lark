module lark

go 1.22.1

require (
	github.com/IBM/sarama v1.41.2
	github.com/apache/rocketmq-client-go/v2 v2.1.2
	github.com/aws/aws-sdk-go v1.50.32
	github.com/bwmarrin/snowflake v0.3.0
	github.com/chromedp/cdproto v0.0.0-20240102194822-c006b26f21c7
	github.com/chromedp/chromedp v0.9.3
	github.com/disintegration/imaging v1.6.2
	github.com/elastic/go-elasticsearch/v7 v7.17.10
	github.com/emersion/go-message v0.17.0
	github.com/emersion/go-msgauth v0.6.8
	github.com/emersion/go-smtp v0.20.1
	github.com/emirpasic/gods v1.18.1
	github.com/gin-contrib/pprof v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/validator/v10 v10.15.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redsync/redsync/v4 v4.9.4
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/protobuf v1.5.3
	github.com/google/go-github/v57 v57.0.0
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/jinzhu/copier v0.4.0
	github.com/jinzhu/inflection v1.0.0
	github.com/json-iterator/go v1.1.12
	github.com/mileusna/spf v0.9.5
	github.com/minio/minio-go/v7 v7.0.63
	github.com/mmcloughlin/geohash v0.10.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nicksnyder/go-i18n/v2 v2.3.0
	github.com/opentracing-contrib/go-gin v0.0.0-20230313182354-1499e56cc906
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/panjf2000/ants/v2 v2.8.1
	github.com/pion/opus v0.0.0-20230805182210-70baa277ae2f
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.19.0
	github.com/quic-go/quic-go v0.41.0
	github.com/redis/go-redis/v9 v9.1.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/shimingyah/pool v1.0.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/shopspring/decimal v1.3.1
	github.com/smartwalle/alipay/v3 v3.2.16
	github.com/smartwalle/paypal v1.0.2
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/streadway/amqp v1.1.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/urfave/cli v1.22.14
	github.com/xtaci/kcp-go/v5 v5.6.3
	github.com/xtaci/kcptun v0.0.0-20230811044945-55e670033ba3
	github.com/xtaci/smux v1.5.24
	github.com/xtaci/tcpraw v1.2.25
	go.etcd.io/etcd/api/v3 v3.5.12
	go.etcd.io/etcd/client/v3 v3.5.12
	go.mongodb.org/mongo-driver v1.12.1
	go.uber.org/dig v1.17.0
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.21.0
	golang.org/x/image v0.12.0
	golang.org/x/net v0.22.0
	golang.org/x/oauth2 v0.18.0
	golang.org/x/text v0.14.0
	golang.org/x/time v0.5.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.4.6
	gorm.io/gorm v1.25.4
	gorm.io/plugin/dbresolver v1.4.7
	vitess.io/vitess v0.19.0
)

require (
	cloud.google.com/go/compute v1.25.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.10.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/coreos/go-iptables v0.7.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.4.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21 // indirect
	github.com/emersion/go-textwrapper v0.0.0-20200911093747-65d896831594 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-redis/redis/v7 v7.4.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.2 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20230817174616-7a8ec2ada47b // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.2 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/klauspost/reedsolomon v1.11.8 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.55 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/onsi/ginkgo/v2 v2.9.5 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.49.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smartwalle/ncrypto v1.0.3 // indirect
	github.com/smartwalle/ngx v1.0.7 // indirect
	github.com/smartwalle/nsign v1.0.8 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/templexxx/cpu v0.1.0 // indirect
	github.com/templexxx/xorsimd v0.4.2 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.12 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/mock v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/mod v0.16.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240304212257-790db918fca8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240304212257-790db918fca8 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	stathat.com/c/consistent v1.0.0 // indirect
)
