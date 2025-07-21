package log

import (
	"net/url"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	//kafka "github.com/segmentio/kafka-go"
	"github.com/Rivalz-ai/framework-be/framework/log/pubsub/zapx"
	"github.com/joho/godotenv"

	//"github.com/Rivalz-ai/framework-be/framework/utils"
	"context"
	"encoding/json"
	"fmt"
	dlog "log"

	"go.elastic.co/apm/v2"
)

type Logger struct {
	id               string
	host             string
	service          string
	logEngineer      *zap.Logger
	logCacheEngineer *zapx.CachedLogger
	mode             int
	dest             int
	env              string
}

var log Logger
var wg sync.WaitGroup
var logLevel int8

//var logger *zap.Logger

/*
Log mode:
	- 0: production
	- 1: staging
	- 2: local
	- 3: dev
Log destination:
	-0: stdout
	-1: file
	-2: pub/sub
args: pub/sub kafka
	args[0]:brokers list, ex: 10.148.0.177:9092,10.148.15.194:9092,10.148.15.198:9092
	args[1]:topic
	args[2]:username
	args[3]:password

*/
/*
args[0]=loglevel string
*/
func Initial(service_name string, args ...interface{}) {
	//
	fmt.Println("Logger Initial")
	//get ENV
	err_env := godotenv.Load(os.ExpandEnv("/config/.env"))
	if err_env != nil {
		err := godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			panic(err)
		}
	}
	env := os.Getenv("ENV") //2: local,1: development,0:product
	logLevelCfg := os.Getenv("LOG_LEVEL")
	mode := 2
	dest := 0
	if env == "" || env == "local" { //local
		mode = 2
		dest = 0
	} else {
		if env == "prd" {
			mode = 0
			dest = 0
		} else if env == "dev" {
			mode = 3
			dest = 0
		} else {
			mode = 1
			dest = 0
		}
	}
	//
	log.service = service_name
	hostname, _ := os.Hostname()
	log.host = hostname
	var err error
	log.mode = mode
	log.dest = dest
	log.env = env
	var cfg zap.Config
	cfg = zap.NewProductionConfig()
	// cfg.OutputPaths = []string{"judger.log"}
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	if mode == 2 {
		cfg.Encoding = "console"
	} else {
		cfg.Encoding = "json"
	}
	log.logEngineer, err = cfg.Build()
	defer log.logEngineer.Sync()
	if err != nil {
		panic(err)
	}
	//

	if len(args) > 0 {
		log_level := logLevelCfg
		if log_level == "" { //default will log error,warn,info
			logLevel = 3
		} else if log_level == "error" {
			logLevel = 1
		} else if log_level == "warn" { //if log level=warn will log: warn & error
			logLevel = 2
		} else { //info =>if log level=info  will log: warn & error & info
			logLevel = 3
		}
	} else { //default will log error,warn,info
		logLevel = 3
	}
	fmt.Println("Logger Initial success")
}
func SetDestKafka(config_map map[string]string) {
	if config_map["BROKERS"] == "" {
		panic("Logger SetDestKafka: LOGGER KAFKA BROKER not found")
	}
	if config_map["TOPIC"] == "" {
		panic("Logger SetDestKafka: LOGGER KAFKA TOPIC not found")
	}
	brokers := config_map["BROKERS"]
	topic := config_map["TOPIC"]
	var err error
	log.dest = 2
	//stderr := zapx.SinkURL{url.URL{Opaque: "stderr"}}
	sinkUrl := zapx.SinkURL{url.URL{Scheme: "kafka", Host: brokers, RawQuery: fmt.Sprintf("topic=%s&username=%s&password=%s", topic, config_map["USERNAME"], config_map["PASSWORD"])}}
	log.logCacheEngineer, err = zapx.NewCachedLoggerConfig().AddSinks(sinkUrl).Build(log.mode)
	defer log.logCacheEngineer.Flush(nil)
	if err != nil {
		dlog.Fatalf(err.Error())
	}
	defer log.logCacheEngineer.Flush(&wg)
	fmt.Println("Logger connected to Kafka success")
}
func Info(msg string, args ...interface{}) {
	if logLevel < 3 { //error,warn
		return
	}
	key := ""
	data := ""
	if len(args) > 0 {
		key = ItoString(args[0])
	}
	if len(args) > 1 {
		data = StructToJson(args[1])
	}
	var arr []zapcore.Field
	arr = append(arr, zap.String("message", msg))
	arr = append(arr, zap.String("host", log.host))
	arr = append(arr, zap.String("service", log.service))
	arr = append(arr, zap.String("key", key))
	arr = append(arr, zap.String("env", log.env))
	arr = append(arr, zap.String("data", data))
	if len(args) > 2 { //context contains apm txn id and trace id
		ctx, ok := args[2].(context.Context)
		if ok {
			txn := apm.TransactionFromContext(ctx)
			if txn != nil {
				traceContext := txn.TraceContext()
				arr = append(arr, zap.String("transaction.id", traceContext.Span.String()))
				arr = append(arr, zap.String("trace.id", traceContext.Trace.String()))
			}
		}
		/*for i := 3; i <= len(args)/2; i++ {
			arr=append(arr,zap.String(ItoString(args[i]), ItoString(args[i+1])))
		}*/
	}
	if log.dest == 2 {
		log.logCacheEngineer.Info(
			"",
			arr...,
		)
		defer log.logCacheEngineer.Flush(&wg)
		//defer log.logCacheEngineer.Flush(nil)
	} else {
		dlog.Println(msg, key, data)
		/*log.logEngineer.Info(
			"",
			zap.String("id",uuid.New().String()),
			zap.String("msg",msg),
			zap.String("host", log.host),
			zap.String("service", log.service),
			zap.String("key_msg", key_msg),
		)
		defer log.logEngineer.Sync()*/
	}

}

/*
msg: log message
args[3]:

	0: key
	1: data structure, interface
	2: context.Context => apm trace context
*/
func Warn(msg string, args ...interface{}) {
	if logLevel < 2 { //error
		return
	}
	key := ""
	data := ""
	if len(args) > 0 {
		key = ItoString(args[0])
	}
	if len(args) > 1 {
		data = StructToJson(args[1])
	}
	var arr []zapcore.Field
	arr = append(arr, zap.String("message", msg))
	arr = append(arr, zap.String("host", log.host))
	arr = append(arr, zap.String("service", log.service))
	arr = append(arr, zap.String("key", key))
	arr = append(arr, zap.String("env", log.env))
	arr = append(arr, zap.String("data", data))
	if len(args) > 2 { //context contains apm txn id and trace id
		ctx, ok := args[2].(context.Context)
		if ok {
			txn := apm.TransactionFromContext(ctx)
			if txn != nil {
				traceContext := txn.TraceContext()
				arr = append(arr, zap.String("transaction.id", traceContext.Span.String()))
				arr = append(arr, zap.String("trace.id", traceContext.Trace.String()))
			}
		}
		/*for i := 3; i <= len(args)/2; i++ {
			arr=append(arr,zap.String(ItoString(args[i]), ItoString(args[i+1])))
		}*/
	}
	if log.dest == 2 {
		log.logCacheEngineer.Warn(
			"",
			arr...,
		)
		defer log.logCacheEngineer.Flush(&wg)
	} else {
		dlog.Println(msg, key, data)
		/*log.logEngineer.Warn(
			"",
			zap.String("id",uuid.New().String()),
			zap.String("msg",msg),
			zap.String("host", log.host),
			zap.String("service", log.service),
			zap.String("key_msg", key_msg),
		)
		defer log.logEngineer.Sync()*/
	}
}

/*
msg: log message
args[3]:

	0: key
	1: data structure, interface
	2: context.Context => apm trace context
*/
func Error(msg string, args ...interface{}) {
	key := ""
	data := ""
	if len(args) > 0 {
		key = ItoString(args[0])
	}
	if len(args) > 1 {
		data = StructToJson(args[1])
	}
	var arr []zapcore.Field
	//
	arr = append(arr, zap.String("message", msg)) //=>for compatible with APM
	//arr=append(arr,zap.String("msg",msg))
	arr = append(arr, zap.String("host", log.host))
	arr = append(arr, zap.String("service", log.service))
	arr = append(arr, zap.String("key", key))
	arr = append(arr, zap.String("env", log.env))
	arr = append(arr, zap.String("data", data))
	if len(args) > 2 { //context contains apm txn id and trace id
		ctx, ok := args[2].(context.Context)
		if ok {
			txn := apm.TransactionFromContext(ctx)
			if txn != nil {
				traceContext := txn.TraceContext()
				arr = append(arr, zap.String("transaction.id", traceContext.Span.String()))
				arr = append(arr, zap.String("trace.id", traceContext.Trace.String()))
			}
		}
		/*for i := 3; i <= len(args)/2; i++ {
			arr=append(arr,zap.String(ItoString(args[i]), ItoString(args[i+1])))
		}*/
	}
	if log.dest == 2 {
		log.logCacheEngineer.Error(
			"",
			arr...,
		)
		defer log.logCacheEngineer.Flush(&wg)
	} else {
		dlog.Println("Error", msg, key, data)
		/*
			log.logEngineer.WithOptions(zap.AddCallerSkip(2)).Error(
				"",
				arr...,
			)
			defer log.logEngineer.Sync()
		*/
	}
}

func ErrorF(msg string, args ...interface{}) {
	Error(msg, args...)
	os.Exit(0)
}

/*
- 0: production
- 1: deveopment
- 2: local
*/
func LogMode() int {
	return log.mode
}

func ItoString(value interface{}) string {
	if value == nil {
		return ""
	}
	str := fmt.Sprintf("%v", value)
	return str
}
func StructToJson(v interface{}) string {
	if v == nil {
		return ""
	}
	out, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return (string(out))
}
