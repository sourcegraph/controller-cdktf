package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveCountMetricsGoodTotalOutputReference interface {
	cdktf.ComplexObject
	AmazonPrometheus() SloObjectiveCountMetricsGoodTotalAmazonPrometheusList
	AmazonPrometheusInput() interface{}
	Appdynamics() SloObjectiveCountMetricsGoodTotalAppdynamicsList
	AppdynamicsInput() interface{}
	AzureMonitor() SloObjectiveCountMetricsGoodTotalAzureMonitorList
	AzureMonitorInput() interface{}
	Bigquery() SloObjectiveCountMetricsGoodTotalBigqueryList
	BigqueryInput() interface{}
	Cloudwatch() SloObjectiveCountMetricsGoodTotalCloudwatchList
	CloudwatchInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Datadog() SloObjectiveCountMetricsGoodTotalDatadogList
	DatadogInput() interface{}
	Dynatrace() SloObjectiveCountMetricsGoodTotalDynatraceList
	DynatraceInput() interface{}
	Elasticsearch() SloObjectiveCountMetricsGoodTotalElasticsearchList
	ElasticsearchInput() interface{}
	// Experimental.
	Fqn() *string
	Gcm() SloObjectiveCountMetricsGoodTotalGcmList
	GcmInput() interface{}
	GrafanaLoki() SloObjectiveCountMetricsGoodTotalGrafanaLokiList
	GrafanaLokiInput() interface{}
	Graphite() SloObjectiveCountMetricsGoodTotalGraphiteList
	GraphiteInput() interface{}
	Honeycomb() SloObjectiveCountMetricsGoodTotalHoneycombList
	HoneycombInput() interface{}
	Influxdb() SloObjectiveCountMetricsGoodTotalInfluxdbList
	InfluxdbInput() interface{}
	Instana() SloObjectiveCountMetricsGoodTotalInstanaList
	InstanaInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lightstep() SloObjectiveCountMetricsGoodTotalLightstepList
	LightstepInput() interface{}
	LogicMonitor() SloObjectiveCountMetricsGoodTotalLogicMonitorList
	LogicMonitorInput() interface{}
	Newrelic() SloObjectiveCountMetricsGoodTotalNewrelicList
	NewrelicInput() interface{}
	Opentsdb() SloObjectiveCountMetricsGoodTotalOpentsdbList
	OpentsdbInput() interface{}
	Pingdom() SloObjectiveCountMetricsGoodTotalPingdomList
	PingdomInput() interface{}
	Prometheus() SloObjectiveCountMetricsGoodTotalPrometheusList
	PrometheusInput() interface{}
	Redshift() SloObjectiveCountMetricsGoodTotalRedshiftList
	RedshiftInput() interface{}
	Splunk() SloObjectiveCountMetricsGoodTotalSplunkList
	SplunkInput() interface{}
	SplunkObservability() SloObjectiveCountMetricsGoodTotalSplunkObservabilityList
	SplunkObservabilityInput() interface{}
	Sumologic() SloObjectiveCountMetricsGoodTotalSumologicList
	SumologicInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thousandeyes() SloObjectiveCountMetricsGoodTotalThousandeyesList
	ThousandeyesInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAmazonPrometheus(value interface{})
	PutAppdynamics(value interface{})
	PutAzureMonitor(value interface{})
	PutBigquery(value interface{})
	PutCloudwatch(value interface{})
	PutDatadog(value interface{})
	PutDynatrace(value interface{})
	PutElasticsearch(value interface{})
	PutGcm(value interface{})
	PutGrafanaLoki(value interface{})
	PutGraphite(value interface{})
	PutHoneycomb(value interface{})
	PutInfluxdb(value interface{})
	PutInstana(value interface{})
	PutLightstep(value interface{})
	PutLogicMonitor(value interface{})
	PutNewrelic(value interface{})
	PutOpentsdb(value interface{})
	PutPingdom(value interface{})
	PutPrometheus(value interface{})
	PutRedshift(value interface{})
	PutSplunk(value interface{})
	PutSplunkObservability(value interface{})
	PutSumologic(value interface{})
	PutThousandeyes(value interface{})
	ResetAmazonPrometheus()
	ResetAppdynamics()
	ResetAzureMonitor()
	ResetBigquery()
	ResetCloudwatch()
	ResetDatadog()
	ResetDynatrace()
	ResetElasticsearch()
	ResetGcm()
	ResetGrafanaLoki()
	ResetGraphite()
	ResetHoneycomb()
	ResetInfluxdb()
	ResetInstana()
	ResetLightstep()
	ResetLogicMonitor()
	ResetNewrelic()
	ResetOpentsdb()
	ResetPingdom()
	ResetPrometheus()
	ResetRedshift()
	ResetSplunk()
	ResetSplunkObservability()
	ResetSumologic()
	ResetThousandeyes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SloObjectiveCountMetricsGoodTotalOutputReference
type jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AmazonPrometheus() SloObjectiveCountMetricsGoodTotalAmazonPrometheusList {
	var returns SloObjectiveCountMetricsGoodTotalAmazonPrometheusList
	_jsii_.Get(
		j,
		"amazonPrometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AmazonPrometheusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonPrometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Appdynamics() SloObjectiveCountMetricsGoodTotalAppdynamicsList {
	var returns SloObjectiveCountMetricsGoodTotalAppdynamicsList
	_jsii_.Get(
		j,
		"appdynamics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AppdynamicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appdynamicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AzureMonitor() SloObjectiveCountMetricsGoodTotalAzureMonitorList {
	var returns SloObjectiveCountMetricsGoodTotalAzureMonitorList
	_jsii_.Get(
		j,
		"azureMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AzureMonitorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Bigquery() SloObjectiveCountMetricsGoodTotalBigqueryList {
	var returns SloObjectiveCountMetricsGoodTotalBigqueryList
	_jsii_.Get(
		j,
		"bigquery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) BigqueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bigqueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Cloudwatch() SloObjectiveCountMetricsGoodTotalCloudwatchList {
	var returns SloObjectiveCountMetricsGoodTotalCloudwatchList
	_jsii_.Get(
		j,
		"cloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) CloudwatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Datadog() SloObjectiveCountMetricsGoodTotalDatadogList {
	var returns SloObjectiveCountMetricsGoodTotalDatadogList
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) DatadogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Dynatrace() SloObjectiveCountMetricsGoodTotalDynatraceList {
	var returns SloObjectiveCountMetricsGoodTotalDynatraceList
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) DynatraceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Elasticsearch() SloObjectiveCountMetricsGoodTotalElasticsearchList {
	var returns SloObjectiveCountMetricsGoodTotalElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Gcm() SloObjectiveCountMetricsGoodTotalGcmList {
	var returns SloObjectiveCountMetricsGoodTotalGcmList
	_jsii_.Get(
		j,
		"gcm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GcmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GrafanaLoki() SloObjectiveCountMetricsGoodTotalGrafanaLokiList {
	var returns SloObjectiveCountMetricsGoodTotalGrafanaLokiList
	_jsii_.Get(
		j,
		"grafanaLoki",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GrafanaLokiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grafanaLokiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Graphite() SloObjectiveCountMetricsGoodTotalGraphiteList {
	var returns SloObjectiveCountMetricsGoodTotalGraphiteList
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GraphiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Honeycomb() SloObjectiveCountMetricsGoodTotalHoneycombList {
	var returns SloObjectiveCountMetricsGoodTotalHoneycombList
	_jsii_.Get(
		j,
		"honeycomb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) HoneycombInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"honeycombInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Influxdb() SloObjectiveCountMetricsGoodTotalInfluxdbList {
	var returns SloObjectiveCountMetricsGoodTotalInfluxdbList
	_jsii_.Get(
		j,
		"influxdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InfluxdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"influxdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Instana() SloObjectiveCountMetricsGoodTotalInstanaList {
	var returns SloObjectiveCountMetricsGoodTotalInstanaList
	_jsii_.Get(
		j,
		"instana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InstanaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Lightstep() SloObjectiveCountMetricsGoodTotalLightstepList {
	var returns SloObjectiveCountMetricsGoodTotalLightstepList
	_jsii_.Get(
		j,
		"lightstep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) LightstepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lightstepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) LogicMonitor() SloObjectiveCountMetricsGoodTotalLogicMonitorList {
	var returns SloObjectiveCountMetricsGoodTotalLogicMonitorList
	_jsii_.Get(
		j,
		"logicMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) LogicMonitorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logicMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Newrelic() SloObjectiveCountMetricsGoodTotalNewrelicList {
	var returns SloObjectiveCountMetricsGoodTotalNewrelicList
	_jsii_.Get(
		j,
		"newrelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) NewrelicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newrelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Opentsdb() SloObjectiveCountMetricsGoodTotalOpentsdbList {
	var returns SloObjectiveCountMetricsGoodTotalOpentsdbList
	_jsii_.Get(
		j,
		"opentsdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) OpentsdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opentsdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Pingdom() SloObjectiveCountMetricsGoodTotalPingdomList {
	var returns SloObjectiveCountMetricsGoodTotalPingdomList
	_jsii_.Get(
		j,
		"pingdom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PingdomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pingdomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Prometheus() SloObjectiveCountMetricsGoodTotalPrometheusList {
	var returns SloObjectiveCountMetricsGoodTotalPrometheusList
	_jsii_.Get(
		j,
		"prometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PrometheusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Redshift() SloObjectiveCountMetricsGoodTotalRedshiftList {
	var returns SloObjectiveCountMetricsGoodTotalRedshiftList
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) RedshiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Splunk() SloObjectiveCountMetricsGoodTotalSplunkList {
	var returns SloObjectiveCountMetricsGoodTotalSplunkList
	_jsii_.Get(
		j,
		"splunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SplunkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SplunkObservability() SloObjectiveCountMetricsGoodTotalSplunkObservabilityList {
	var returns SloObjectiveCountMetricsGoodTotalSplunkObservabilityList
	_jsii_.Get(
		j,
		"splunkObservability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SplunkObservabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkObservabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Sumologic() SloObjectiveCountMetricsGoodTotalSumologicList {
	var returns SloObjectiveCountMetricsGoodTotalSumologicList
	_jsii_.Get(
		j,
		"sumologic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SumologicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumologicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Thousandeyes() SloObjectiveCountMetricsGoodTotalThousandeyesList {
	var returns SloObjectiveCountMetricsGoodTotalThousandeyesList
	_jsii_.Get(
		j,
		"thousandeyes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ThousandeyesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thousandeyesInput",
		&returns,
	)
	return returns
}


func NewSloObjectiveCountMetricsGoodTotalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveCountMetricsGoodTotalOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveCountMetricsGoodTotalOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodTotalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsGoodTotalOutputReference_Override(s SloObjectiveCountMetricsGoodTotalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodTotalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutAmazonPrometheus(value interface{}) {
	if err := s.validatePutAmazonPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAmazonPrometheus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutAppdynamics(value interface{}) {
	if err := s.validatePutAppdynamicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppdynamics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutAzureMonitor(value interface{}) {
	if err := s.validatePutAzureMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureMonitor",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutBigquery(value interface{}) {
	if err := s.validatePutBigqueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBigquery",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutCloudwatch(value interface{}) {
	if err := s.validatePutCloudwatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatch",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutDatadog(value interface{}) {
	if err := s.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatadog",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutDynatrace(value interface{}) {
	if err := s.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutElasticsearch(value interface{}) {
	if err := s.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutGcm(value interface{}) {
	if err := s.validatePutGcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcm",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutGrafanaLoki(value interface{}) {
	if err := s.validatePutGrafanaLokiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrafanaLoki",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutGraphite(value interface{}) {
	if err := s.validatePutGraphiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGraphite",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutHoneycomb(value interface{}) {
	if err := s.validatePutHoneycombParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHoneycomb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutInfluxdb(value interface{}) {
	if err := s.validatePutInfluxdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfluxdb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutInstana(value interface{}) {
	if err := s.validatePutInstanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstana",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutLightstep(value interface{}) {
	if err := s.validatePutLightstepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLightstep",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutLogicMonitor(value interface{}) {
	if err := s.validatePutLogicMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogicMonitor",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutNewrelic(value interface{}) {
	if err := s.validatePutNewrelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNewrelic",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutOpentsdb(value interface{}) {
	if err := s.validatePutOpentsdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOpentsdb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutPingdom(value interface{}) {
	if err := s.validatePutPingdomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPingdom",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutPrometheus(value interface{}) {
	if err := s.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrometheus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutRedshift(value interface{}) {
	if err := s.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedshift",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutSplunk(value interface{}) {
	if err := s.validatePutSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutSplunkObservability(value interface{}) {
	if err := s.validatePutSplunkObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunkObservability",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutSumologic(value interface{}) {
	if err := s.validatePutSumologicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSumologic",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutThousandeyes(value interface{}) {
	if err := s.validatePutThousandeyesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThousandeyes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetAmazonPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetAmazonPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetAppdynamics() {
	_jsii_.InvokeVoid(
		s,
		"resetAppdynamics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetAzureMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetBigquery() {
	_jsii_.InvokeVoid(
		s,
		"resetBigquery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetCloudwatch() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		s,
		"resetDatadog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		s,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		s,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetGcm() {
	_jsii_.InvokeVoid(
		s,
		"resetGcm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetGrafanaLoki() {
	_jsii_.InvokeVoid(
		s,
		"resetGrafanaLoki",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		s,
		"resetGraphite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetHoneycomb() {
	_jsii_.InvokeVoid(
		s,
		"resetHoneycomb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetInfluxdb() {
	_jsii_.InvokeVoid(
		s,
		"resetInfluxdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetInstana() {
	_jsii_.InvokeVoid(
		s,
		"resetInstana",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetLightstep() {
	_jsii_.InvokeVoid(
		s,
		"resetLightstep",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetLogicMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetLogicMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetNewrelic() {
	_jsii_.InvokeVoid(
		s,
		"resetNewrelic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetOpentsdb() {
	_jsii_.InvokeVoid(
		s,
		"resetOpentsdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetPingdom() {
	_jsii_.InvokeVoid(
		s,
		"resetPingdom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		s,
		"resetRedshift",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetSplunk() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetSplunkObservability() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunkObservability",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetSumologic() {
	_jsii_.InvokeVoid(
		s,
		"resetSumologic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ResetThousandeyes() {
	_jsii_.InvokeVoid(
		s,
		"resetThousandeyes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

