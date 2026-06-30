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
	AmazonPrometheusInput() any
	Appdynamics() SloObjectiveCountMetricsGoodTotalAppdynamicsList
	AppdynamicsInput() any
	AzureMonitor() SloObjectiveCountMetricsGoodTotalAzureMonitorList
	AzureMonitorInput() any
	Bigquery() SloObjectiveCountMetricsGoodTotalBigqueryList
	BigqueryInput() any
	Cloudwatch() SloObjectiveCountMetricsGoodTotalCloudwatchList
	CloudwatchInput() any
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	DatadogInput() any
	Dynatrace() SloObjectiveCountMetricsGoodTotalDynatraceList
	DynatraceInput() any
	Elasticsearch() SloObjectiveCountMetricsGoodTotalElasticsearchList
	ElasticsearchInput() any
	// Experimental.
	Fqn() *string
	Gcm() SloObjectiveCountMetricsGoodTotalGcmList
	GcmInput() any
	GrafanaLoki() SloObjectiveCountMetricsGoodTotalGrafanaLokiList
	GrafanaLokiInput() any
	Graphite() SloObjectiveCountMetricsGoodTotalGraphiteList
	GraphiteInput() any
	Honeycomb() SloObjectiveCountMetricsGoodTotalHoneycombList
	HoneycombInput() any
	Influxdb() SloObjectiveCountMetricsGoodTotalInfluxdbList
	InfluxdbInput() any
	Instana() SloObjectiveCountMetricsGoodTotalInstanaList
	InstanaInput() any
	InternalValue() any
	SetInternalValue(val any)
	Lightstep() SloObjectiveCountMetricsGoodTotalLightstepList
	LightstepInput() any
	LogicMonitor() SloObjectiveCountMetricsGoodTotalLogicMonitorList
	LogicMonitorInput() any
	Newrelic() SloObjectiveCountMetricsGoodTotalNewrelicList
	NewrelicInput() any
	Opentsdb() SloObjectiveCountMetricsGoodTotalOpentsdbList
	OpentsdbInput() any
	Pingdom() SloObjectiveCountMetricsGoodTotalPingdomList
	PingdomInput() any
	Prometheus() SloObjectiveCountMetricsGoodTotalPrometheusList
	PrometheusInput() any
	Redshift() SloObjectiveCountMetricsGoodTotalRedshiftList
	RedshiftInput() any
	Splunk() SloObjectiveCountMetricsGoodTotalSplunkList
	SplunkInput() any
	SplunkObservability() SloObjectiveCountMetricsGoodTotalSplunkObservabilityList
	SplunkObservabilityInput() any
	Sumologic() SloObjectiveCountMetricsGoodTotalSumologicList
	SumologicInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thousandeyes() SloObjectiveCountMetricsGoodTotalThousandeyesList
	ThousandeyesInput() any
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutAmazonPrometheus(value any)
	PutAppdynamics(value any)
	PutAzureMonitor(value any)
	PutBigquery(value any)
	PutCloudwatch(value any)
	PutDatadog(value any)
	PutDynatrace(value any)
	PutElasticsearch(value any)
	PutGcm(value any)
	PutGrafanaLoki(value any)
	PutGraphite(value any)
	PutHoneycomb(value any)
	PutInfluxdb(value any)
	PutInstana(value any)
	PutLightstep(value any)
	PutLogicMonitor(value any)
	PutNewrelic(value any)
	PutOpentsdb(value any)
	PutPingdom(value any)
	PutPrometheus(value any)
	PutRedshift(value any)
	PutSplunk(value any)
	PutSplunkObservability(value any)
	PutSumologic(value any)
	PutThousandeyes(value any)
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
	Resolve(_context cdktf.IResolveContext) any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AmazonPrometheusInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AppdynamicsInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) AzureMonitorInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) BigqueryInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) CloudwatchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"cloudwatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ComplexObjectIndex() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) DatadogInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) DynatraceInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ElasticsearchInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GcmInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GrafanaLokiInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GraphiteInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) HoneycombInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InfluxdbInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InstanaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"instanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) InternalValue() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) LightstepInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) LogicMonitorInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) NewrelicInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) OpentsdbInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PingdomInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PrometheusInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) RedshiftInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SplunkInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SplunkObservabilityInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SumologicInput() any {
	var returns any
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

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) ThousandeyesInput() any {
	var returns any
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
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsGoodTotalOutputReference_Override(s SloObjectiveCountMetricsGoodTotalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodTotalOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutAmazonPrometheus(value any) {
	if err := s.validatePutAmazonPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAmazonPrometheus",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutAppdynamics(value any) {
	if err := s.validatePutAppdynamicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppdynamics",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutAzureMonitor(value any) {
	if err := s.validatePutAzureMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureMonitor",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutBigquery(value any) {
	if err := s.validatePutBigqueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBigquery",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutCloudwatch(value any) {
	if err := s.validatePutCloudwatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatch",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutDatadog(value any) {
	if err := s.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatadog",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutDynatrace(value any) {
	if err := s.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynatrace",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutElasticsearch(value any) {
	if err := s.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElasticsearch",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutGcm(value any) {
	if err := s.validatePutGcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcm",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutGrafanaLoki(value any) {
	if err := s.validatePutGrafanaLokiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrafanaLoki",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutGraphite(value any) {
	if err := s.validatePutGraphiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGraphite",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutHoneycomb(value any) {
	if err := s.validatePutHoneycombParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHoneycomb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutInfluxdb(value any) {
	if err := s.validatePutInfluxdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfluxdb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutInstana(value any) {
	if err := s.validatePutInstanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstana",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutLightstep(value any) {
	if err := s.validatePutLightstepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLightstep",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutLogicMonitor(value any) {
	if err := s.validatePutLogicMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogicMonitor",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutNewrelic(value any) {
	if err := s.validatePutNewrelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNewrelic",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutOpentsdb(value any) {
	if err := s.validatePutOpentsdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOpentsdb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutPingdom(value any) {
	if err := s.validatePutPingdomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPingdom",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutPrometheus(value any) {
	if err := s.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrometheus",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutRedshift(value any) {
	if err := s.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedshift",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutSplunk(value any) {
	if err := s.validatePutSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunk",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutSplunkObservability(value any) {
	if err := s.validatePutSplunkObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunkObservability",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutSumologic(value any) {
	if err := s.validatePutSumologicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSumologic",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) PutThousandeyes(value any) {
	if err := s.validatePutThousandeyesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThousandeyes",
		[]any{value},
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodTotalOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		s,
		"resolve",
		[]any{_context},
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
