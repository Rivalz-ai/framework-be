package apmelk

import (
	"fmt"

	"github.com/Rivalz-ai/framework-be/framework/config/vault"
	"github.com/Rivalz-ai/framework-be/framework/utils"

	//"github.com/Rivalz-ai/framework-be/framework/log"
	"context"
	"os"
	"strings"

	"go.elastic.co/apm/v2"
)

var ServiceName string
var service_enable bool

type APMAgent struct {
	serviceName string
	secretToken string
	config      *vault.Vault
	isDebug     bool
	enable      bool
	cfg         map[string]string
}
type APMTransaction struct {
	t      *apm.Transaction
	tracer *apm.Tracer
}
type APMSpan struct {
	s *apm.Span
}

func (agent *APMAgent) SetDebug(debug bool) *APMAgent {
	agent.isDebug = debug
	return agent
}

func (agent *APMAgent) GetServiceName() string {
	return agent.serviceName
}

func (agent *APMAgent) SetServiceName(name string) *APMAgent {
	agent.serviceName = name
	return agent
}
func (agent *APMAgent) GetEnableAPM() bool {
	return agent.enable
}
func GetGlobalEnableAPM() bool {
	return service_enable
}
func (agent *APMAgent) GetConfig() map[string]string {
	return agent.cfg
}
func (agent *APMAgent) Initial(vault *vault.Vault) {
	if vault == nil {
		return
	}
	agent.config = vault
	agent.serviceName = agent.config.GetServiceName()
	ServiceName = strings.ReplaceAll(agent.serviceName, ".", "_")
	ServiceName = strings.ReplaceAll(ServiceName, "/", "_")
	//APM Global
	globalConfigPath := "apm"
	//APM service
	srvConfigPath := fmt.Sprintf("%s/%s", strings.ReplaceAll(agent.serviceName, ".", "/"), "apm")
	//global config
	global_cfg := GetConfig(agent.config, globalConfigPath)
	if global_cfg != nil {
		if global_cfg["ENABLE"] == "true" {
			agent.enable = true
		}
	}
	if !agent.enable {
		return
	}
	check, err := agent.config.CheckItemExist(srvConfigPath)
	if err != nil {
		return
	}
	if !check {
		return
	}
	//service config
	srv_cfg := GetConfig(agent.config, srvConfigPath)
	//merge config
	for k, v := range srv_cfg {
		if v == "" {
			continue
		}
		global_cfg[k] = v
	}
	agent.cfg = global_cfg
	//cfg["ENABLE"] => global enable APM server => effect for all service
	//cfg["ENABLE_GPRC"] => only service enable
	if agent.cfg["ENABLE_APM"] == "true" && agent.cfg["ENABLE"] == "true" {
		agent.enable = true
		service_enable = true
		agent.SetAPMenv(agent.cfg)
		if agent.cfg["ENABLE_AGENT_DEBUG"] == "1" || agent.cfg["ENABLE_AGENT_DEBUG"] == "true" {
			agent.isDebug = true
			os.Setenv("ELASTIC_APM_LOG_FILE", "stderr")
			os.Setenv("ELASTIC_APM_LOG_LEVEL", "debug")
		}
	}
}
func (agent *APMAgent) SetAPMenv(cfg map[string]string) {
	//m := utils.DictionaryString()
	fmt.Println("Set APM ENV: " + cfg["ELASTIC_APM_SERVER_URL"])
	os.Setenv("ELASTIC_APM_SECRET_TOKEN", cfg["ELASTIC_APM_SECRET_TOKEN"])
	os.Setenv("ELASTIC_APM_ENVIRONMENT", os.Getenv("ENV"))
	os.Setenv("ELASTIC_APM_SERVER_URL", cfg["ELASTIC_APM_SERVER_URL"])
	os.Setenv("ELASTIC_APM_SERVICE_NAME", cfg["ELASTIC_APM_SERVICE_NAME"])
	//return m
}

/*
this method no need call flush, because it will always create tracer, but performance is not so good
when we call End(), End() method will see tracer!=nil then if will call Flush manually

Parameters:

	 name: txn name, for find this txn in APM kibana
	 args[4]:
		0: group name
		1: custome data
		2: ctx: parent context, to link with parent apm txn
		3: uid => key for search this txn in kibana, key search: label.uid=
*/
func NewTransactionWithTracer(ctx context.Context, name string, args ...interface{}) (*APMTransaction, context.Context, error) {
	if !service_enable {
		return &APMTransaction{}, ctx, nil
	}
	group := ""
	var custome_data interface{}
	if len(args) > 0 {
		group = utils.ItoString(args[0])
	}
	if len(args) > 1 {
		custome_data = args[1]
	}
	opts := apm.TransactionOptions{}
	if len(args) > 2 { //opt was passed from external
		traceContext := args[2].(apm.TraceContext)
		opts = apm.TransactionOptions{
			TraceContext: traceContext,
		}
	} else { //opt in context
		tx := apm.TransactionFromContext(ctx)
		if tx != nil {
			traceContext := tx.TraceContext()
			opts = apm.TransactionOptions{
				TraceContext: traceContext,
			}
		}
	}

	tracer, err := apm.NewTracer(ServiceName, "")
	if err != nil {
		return nil, nil, err
	}

	t := tracer.StartTransactionOptions(name, group, opts)
	c := apm.ContextWithTransaction(ctx, t)
	t.TransactionData.Context.SetCustom("data", custome_data)
	if len(args) > 3 { //for search in kibana UI
		uid := utils.ItoString(args[3])
		t.TransactionData.Context.SetLabel("uid", uid)
	}
	return &APMTransaction{
		tracer: tracer,
		t:      t,
	}, c, nil
}

/*
	this method use default global Tracer object, so when Dev call End() method, End() method  see tracer=nil, it will not call Flush method
*/
/*
func NewTransaction(ctx context.Context,name string,args...interface{}) (*APMTransaction,context.Context){
	if !enable{
		return  &APMTransaction{},ctx
	}
	group:=""
	var custome_data interface{}
	if len(args)>0{
		group=utils.ItoString(args[0])
	}
	if len(args)>1{
		custome_data=args[1]
	}
	t := apm.DefaultTracer().StartTransaction(name, group)
	c := apm.ContextWithTransaction(ctx, t)
	t.TransactionData.Context.SetCustom("data",custome_data)
	return &APMTransaction{
			t:t,
	},c
}
*/
/*
 name: txn name, for find this txn in APM kibana
 args[4]:
	0: group name
	1: custome data
	2: uid => key for search this txn in kibana, key search: label.uid=
	3: ctx: parent context, to link with parent apm txn
*/
func NewTransaction(ctx context.Context, name string, args ...interface{}) (*APMTransaction, context.Context) {
	if !service_enable {
		return &APMTransaction{}, ctx
	}
	group := ""
	var custome_data interface{}
	if len(args) > 0 { //group name
		group = utils.ItoString(args[0])
	}
	if len(args) > 1 { //custome data
		custome_data = args[1]
	}

	opts := apm.TransactionOptions{}
	if len(args) > 3 { //opt was passed from external
		traceContext := args[3].(apm.TraceContext)
		opts = apm.TransactionOptions{
			TraceContext: traceContext,
		}
	} else { //opt in context
		tx := apm.TransactionFromContext(ctx)
		if tx != nil {
			traceContext := tx.TraceContext()
			opts = apm.TransactionOptions{
				TraceContext: traceContext,
			}
		}
	}

	//
	t := apm.DefaultTracer().StartTransactionOptions(name, group, opts)
	c := apm.ContextWithTransaction(ctx, t)
	//more detail about data on this txn
	t.TransactionData.Context.SetCustom("data", custome_data)
	if len(args) > 2 { //for search in kibana UI
		uid := utils.ItoString(args[2])
		t.TransactionData.Context.SetLabel("uid", uid)
	}
	//t.TransactionData.Context.SetLa("data",custome_data)
	return &APMTransaction{
		t: t,
	}, c
}

// should call this method before service shutdown, if not some apm transactions will be loss, because apm transaction queue not full or next period for flush data to apm server
func Flush() {
	if !service_enable {
		return
	}
	apm.DefaultTracer().Flush(nil)
}
func (t APMTransaction) End() {
	if !service_enable {
		return
	}
	if t.t != nil {
		t.t.End()
		if t.tracer != nil {
			t.tracer.Flush(nil)
			t.tracer.Close()
		}
	}
}
func StartSpan(c context.Context, name string, args ...string) *APMSpan {
	if !service_enable {
		return &APMSpan{}
	}
	_type := ""
	if len(args) > 0 {
		_type = args[0]
	}
	span, _ := apm.StartSpan(c, name, _type)
	return &APMSpan{
		s: span,
	}
}
func (s *APMSpan) End() {
	if !service_enable {
		return
	}
	s.s.End()
}

func (t APMTransaction) GetOriginTxn() *apm.Transaction {
	return t.t
}
func GetTraceIDFromBytes(arr []byte) apm.TraceID {
	var trace_id apm.TraceID
	if len(arr) != 16 {
		return trace_id
	}
	return apm.TraceID{arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11], arr[12], arr[13], arr[14], arr[15]}
}
func GetSpanIDFromBytes(arr []byte) apm.SpanID {
	var trace_id apm.SpanID
	if len(arr) != 8 {
		return trace_id
	}
	return apm.SpanID{arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7]}
}
func GetTraceOptionsUInt8(v uint8) apm.TraceOptions {
	return apm.TraceOptions(v)
}
