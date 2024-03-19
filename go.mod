module cloud_platform

go 1.21

replace (
	github.com/go-oauth2/oauth2/v4 => ./thirdlib/oauth2/v4@v4.5.1
	github.com/go-redis/redis/v8 => ./thirdlib/redis
	github.com/h2non/filetype => ./thirdlib/filetype
	github.com/jellydator/ttlcache/v2 v2.11.1 => github.com/ReneKroon/ttlcache/v2 v2.11.1
	gorm.io/gen => ./thirdlib/gen
)

require (
	github.com/anthonynsimon/bild v0.13.0
	github.com/asim/go-micro/plugins/client/grpc/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/asim/go-micro/plugins/server/grpc/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/bilibili/gengine v1.5.7
	github.com/chenyahui/gin-cache v1.5.0
	github.com/cyanBone/dingtalk_robot v0.1.0
	github.com/forgoer/openssl v1.2.1
	github.com/gin-contrib/gzip v0.0.3
	github.com/gin-gonic/gin v1.8.1
	github.com/go-oauth2/gin-server v1.1.0
	github.com/go-oauth2/oauth2/v4 v4.5.1
	github.com/go-pay/gopay v1.5.79
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-session/session v3.1.2+incompatible
	github.com/goccy/go-json v0.9.11
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/h2non/filetype v1.1.3
	github.com/mholt/archiver/v3 v3.5.1
	github.com/mozillazg/go-pinyin v0.19.0
	github.com/mssola/user_agent v0.6.0
	github.com/oschwald/geoip2-golang v1.7.0
	github.com/qiniu/x v1.10.5
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/tealeg/xlsx v1.0.5
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.389
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.389
	github.com/tidwall/gjson v1.14.3
	github.com/volcengine/volc-sdk-golang v1.0.41
	github.com/xhit/go-simple-mail/v2 v2.13.0
	github.com/xuri/excelize/v2 v2.6.0
	github.com/xxl-job/xxl-job-executor-go v1.1.0
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90
	golang.org/x/net v0.1.0
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	google.golang.org/api v0.81.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/clickhouse v0.4.2
	howett.net/plist v1.0.0
)

require (
	cloud.google.com/go/compute v1.6.1 // indirect
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/ClickHouse/clickhouse-go/v2 v2.2.0 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/antlr/antlr4 v0.0.0-20210105192202-5c2b686f95e1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/clbanning/mxj v1.8.5-0.20200714211355-ff02cfb8ea28 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/denisenkom/go-mssqldb v0.12.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dsnet/compress v0.0.2-0.20210315054119-f66993602bf5 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-basic/ipv4 v1.0.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-test/deep v1.0.8 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.0.0-20170517235910-f1bb20e5a188 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/google/martian v2.1.0+incompatible // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grokify/html-strip-tags-go v0.0.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.11.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	github.com/jackc/pgx/v4 v4.15.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/minio/highwayhash v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nats-io/jwt/v2 v2.0.2 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nwaples/rardecode v1.1.3 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/oschwald/maxminddb-golang v1.9.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paulmach/orb v0.7.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.7.2 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/btree v0.0.0-20191029221954-400434d76274 // indirect
	github.com/tidwall/buntdb v1.1.2 // indirect
	github.com/tidwall/grect v0.0.0-20161006141115-ba9a043346eb // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtree v0.0.0-20180113144539-6cd427091e0e // indirect
	github.com/tidwall/tinyqueue v0.0.0-20180302190814-1e39f5511563 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/xuri/efp v0.0.0-20220407160117-ad0f7a785be8 // indirect
	github.com/xuri/nfp v0.0.0-20220409054826-5e722a1d9e22 // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/v3 v3.5.4 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.46.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/datatypes v1.0.6 // indirect
	gorm.io/driver/postgres v1.3.4 // indirect
	gorm.io/driver/sqlserver v1.3.2 // indirect
	gorm.io/hints v1.1.0 // indirect
	gorm.io/plugin/dbresolver v1.1.0 // indirect
)

require (
	github.com/ReneKroon/ttlcache/v2 v2.11.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.976
	github.com/aliyun/aliyun-oss-go-sdk v2.2.2+incompatible
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/asim/go-micro/plugins/broker/nats/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.7.0
	github.com/asim/go-micro/plugins/config/source/etcd/v4 v4.7.0
	github.com/asim/go-micro/plugins/registry/etcd/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/asim/go-micro/plugins/transport/grpc/v4 v4.0.0-20220415015530-034ba9a0dead // indirect
	github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4 v4.0.0-20220415015530-034ba9a0dead
	github.com/aws/aws-sdk-go v1.44.153
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/casbin/casbin/v2 v2.44.1
	github.com/casbin/gorm-adapter/v3 v3.5.1
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/fsnotify/fsnotify v1.5.4
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogf/gf v1.16.7
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8
	github.com/google/uuid v1.3.0
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.21.12+incompatible
	github.com/json-iterator/go v1.1.12
	github.com/mitchellh/mapstructure v1.5.0
	github.com/nats-io/nats.go v1.14.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.1
	github.com/prometheus/client_golang v1.13.0
	github.com/qiniu/go-sdk/v7 v7.12.0
	github.com/shamsher31/goimgext v1.0.0
	github.com/sony/sonyflake v1.0.0
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.0
	github.com/subosito/gotenv v1.4.1
	github.com/swaggo/gin-swagger v1.4.2
	github.com/swaggo/swag v1.8.1
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go-micro.dev/v4 v4.6.0
	go.uber.org/zap v1.17.0
	golang.org/x/text v0.4.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/protobuf v1.28.1
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gen v0.2.44
	gorm.io/gorm v1.23.8
)
