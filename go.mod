module github.com/minio/minio

go 1.16

require (
	cloud.google.com/go/storage v1.8.0
	git.apache.org/thrift.git v0.13.0
	github.com/Azure/azure-pipeline-go v0.2.2
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest/adal v0.9.1 // indirect
	github.com/Shopify/sarama v1.27.2
	github.com/VividCortex/ewma v1.1.1
	github.com/alecthomas/participle v0.2.1
	github.com/bcicen/jstream v1.0.1
	github.com/beevik/ntp v0.3.0
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/cheggaaa/pb v1.0.29
	github.com/colinmarc/hdfs/v2 v2.2.0
	github.com/coredns/coredns v1.4.0
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dchest/siphash v1.2.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/djherbis/atime v1.0.0
	github.com/dswarbrick/smart v0.0.0-20190505152634-909a45200d6d
	github.com/dustin/go-humanize v1.0.0
	github.com/eclipse/paho.mqtt.golang v1.3.0
	github.com/fatih/color v1.10.0
	github.com/fatih/structs v1.1.0
	github.com/go-ldap/ldap/v3 v3.2.4
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/martian v2.1.1-0.20190517191504-25dcb96d9e51+incompatible // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2
	github.com/json-iterator/go v1.1.10
	github.com/klauspost/compress v1.11.12
	github.com/klauspost/cpuid/v2 v2.0.4
	github.com/klauspost/pgzip v1.2.5
	github.com/klauspost/readahead v1.3.1
	github.com/klauspost/reedsolomon v1.9.11
	github.com/lib/pq v1.9.0
	github.com/mattn/go-colorable v0.1.8
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/miekg/dns v1.1.35
	github.com/minio/argon2 v0.0.0-20210427164258-0025d10c2c04
	github.com/minio/cli v1.22.0
	github.com/minio/csvparser v0.0.0-20210427163918-ee4f0ffa388d
	github.com/minio/highwayhash v1.0.2
	github.com/minio/md5-simd v1.1.1 // indirect
	github.com/minio/minio-go/v7 v7.0.11-0.20210302210017-6ae69c73ce78
	github.com/minio/selfupdate v0.3.1
	github.com/minio/sha256-simd v1.0.0
	github.com/minio/simdjson-go v0.2.1
	github.com/minio/sio v0.2.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/montanaflynn/stats v0.5.0
	github.com/nats-io/nats-server/v2 v2.1.9
	github.com/nats-io/nats-streaming-server v0.21.1 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/nats-io/stan.go v0.8.3
	github.com/ncw/directio v1.0.5
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic/v7 v7.0.22
	github.com/philhofer/fwd v1.1.1
	github.com/pierrec/lz4 v2.5.2+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/procfs v0.6.0
	github.com/rjeczalik/notify v0.9.2
	github.com/rs/cors v1.7.0
	github.com/secure-io/sio-go v0.3.1
	github.com/shirou/gopsutil/v3 v3.21.1
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/tidwall/gjson v1.6.8
	github.com/tidwall/sjson v1.0.4
	github.com/tinylib/msgp v1.1.3
	github.com/valyala/tcplisten v0.0.0-20161114210144-ceec8f93295a
	github.com/willf/bitset v1.1.11 // indirect
	github.com/willf/bloom v2.0.3+incompatible
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
	go.etcd.io/etcd v0.0.0-20201125193152-8a03d2e9614b
	go.uber.org/zap v1.14.1
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
	golang.org/x/mod v0.4.0 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/sys v0.0.0-20210225134936-a50acf3fe073
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	golang.org/x/tools v0.1.1-0.20210201201750-4d4ee958a9b7 // indirect
	google.golang.org/api v0.25.0
	google.golang.org/genproto v0.0.0-20200527145253-8367513e4ece // indirect
	gopkg.in/yaml.v2 v2.3.0
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
