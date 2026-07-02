package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveCountMetricsBadOutputReference interface {
	cdktf.ComplexObject
	AmazonPrometheus() SloObjectiveCountMetricsBadAmazonPrometheusList
	AmazonPrometheusInput() any
	Appdynamics() SloObjectiveCountMetricsBadAppdynamicsList
	AppdynamicsInput() any
	AzureMonitor() SloObjectiveCountMetricsBadAzureMonitorList
	AzureMonitorInput() any
	Bigquery() SloObjectiveCountMetricsBadBigqueryList
	BigqueryInput() any
	Cloudwatch() SloObjectiveCountMetricsBadCloudwatchList
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
	Datadog() SloObjectiveCountMetricsBadDatadogList
	DatadogInput() any
	Dynatrace() SloObjectiveCountMetricsBadDynatraceList
	DynatraceInput() any
	Elasticsearch() SloObjectiveCountMetricsBadElasticsearchList
	ElasticsearchInput() any
	// Experimental.
	Fqn() *string
	Gcm() SloObjectiveCountMetricsBadGcmList
	GcmInput() any
	GrafanaLoki() SloObjectiveCountMetricsBadGrafanaLokiList
	GrafanaLokiInput() any
	Graphite() SloObjectiveCountMetricsBadGraphiteList
	GraphiteInput() any
	Honeycomb() SloObjectiveCountMetricsBadHoneycombList
	HoneycombInput() any
	Influxdb() SloObjectiveCountMetricsBadInfluxdbList
	InfluxdbInput() any
	Instana() SloObjectiveCountMetricsBadInstanaList
	InstanaInput() any
	InternalValue() any
	SetInternalValue(val any)
	Lightstep() SloObjectiveCountMetricsBadLightstepList
	LightstepInput() any
	LogicMonitor() SloObjectiveCountMetricsBadLogicMonitorList
	LogicMonitorInput() any
	Newrelic() SloObjectiveCountMetricsBadNewrelicList
	NewrelicInput() any
	Opentsdb() SloObjectiveCountMetricsBadOpentsdbList
	OpentsdbInput() any
	Pingdom() SloObjectiveCountMetricsBadPingdomList
	PingdomInput() any
	Prometheus() SloObjectiveCountMetricsBadPrometheusList
	PrometheusInput() any
	Redshift() SloObjectiveCountMetricsBadRedshiftList
	RedshiftInput() any
	Splunk() SloObjectiveCountMetricsBadSplunkList
	SplunkInput() any
	SplunkObservability() SloObjectiveCountMetricsBadSplunkObservabilityList
	SplunkObservabilityInput() any
	Sumologic() SloObjectiveCountMetricsBadSumologicList
	SumologicInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thousandeyes() SloObjectiveCountMetricsBadThousandeyesList
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

// The jsii proxy struct for SloObjectiveCountMetricsBadOutputReference
type jsiiProxy_SloObjectiveCountMetricsBadOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) AmazonPrometheus() SloObjectiveCountMetricsBadAmazonPrometheusList {
	var returns SloObjectiveCountMetricsBadAmazonPrometheusList
	_jsii_.Get(
		j,
		"amazonPrometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) AmazonPrometheusInput() any {
	var returns any
	_jsii_.Get(
		j,
		"amazonPrometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Appdynamics() SloObjectiveCountMetricsBadAppdynamicsList {
	var returns SloObjectiveCountMetricsBadAppdynamicsList
	_jsii_.Get(
		j,
		"appdynamics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) AppdynamicsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"appdynamicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) AzureMonitor() SloObjectiveCountMetricsBadAzureMonitorList {
	var returns SloObjectiveCountMetricsBadAzureMonitorList
	_jsii_.Get(
		j,
		"azureMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) AzureMonitorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"azureMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Bigquery() SloObjectiveCountMetricsBadBigqueryList {
	var returns SloObjectiveCountMetricsBadBigqueryList
	_jsii_.Get(
		j,
		"bigquery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) BigqueryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"bigqueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Cloudwatch() SloObjectiveCountMetricsBadCloudwatchList {
	var returns SloObjectiveCountMetricsBadCloudwatchList
	_jsii_.Get(
		j,
		"cloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) CloudwatchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"cloudwatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Datadog() SloObjectiveCountMetricsBadDatadogList {
	var returns SloObjectiveCountMetricsBadDatadogList
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) DatadogInput() any {
	var returns any
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Dynatrace() SloObjectiveCountMetricsBadDynatraceList {
	var returns SloObjectiveCountMetricsBadDynatraceList
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) DynatraceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Elasticsearch() SloObjectiveCountMetricsBadElasticsearchList {
	var returns SloObjectiveCountMetricsBadElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ElasticsearchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Gcm() SloObjectiveCountMetricsBadGcmList {
	var returns SloObjectiveCountMetricsBadGcmList
	_jsii_.Get(
		j,
		"gcm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GcmInput() any {
	var returns any
	_jsii_.Get(
		j,
		"gcmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GrafanaLoki() SloObjectiveCountMetricsBadGrafanaLokiList {
	var returns SloObjectiveCountMetricsBadGrafanaLokiList
	_jsii_.Get(
		j,
		"grafanaLoki",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GrafanaLokiInput() any {
	var returns any
	_jsii_.Get(
		j,
		"grafanaLokiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Graphite() SloObjectiveCountMetricsBadGraphiteList {
	var returns SloObjectiveCountMetricsBadGraphiteList
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GraphiteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Honeycomb() SloObjectiveCountMetricsBadHoneycombList {
	var returns SloObjectiveCountMetricsBadHoneycombList
	_jsii_.Get(
		j,
		"honeycomb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) HoneycombInput() any {
	var returns any
	_jsii_.Get(
		j,
		"honeycombInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Influxdb() SloObjectiveCountMetricsBadInfluxdbList {
	var returns SloObjectiveCountMetricsBadInfluxdbList
	_jsii_.Get(
		j,
		"influxdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) InfluxdbInput() any {
	var returns any
	_jsii_.Get(
		j,
		"influxdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Instana() SloObjectiveCountMetricsBadInstanaList {
	var returns SloObjectiveCountMetricsBadInstanaList
	_jsii_.Get(
		j,
		"instana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) InstanaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"instanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Lightstep() SloObjectiveCountMetricsBadLightstepList {
	var returns SloObjectiveCountMetricsBadLightstepList
	_jsii_.Get(
		j,
		"lightstep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) LightstepInput() any {
	var returns any
	_jsii_.Get(
		j,
		"lightstepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) LogicMonitor() SloObjectiveCountMetricsBadLogicMonitorList {
	var returns SloObjectiveCountMetricsBadLogicMonitorList
	_jsii_.Get(
		j,
		"logicMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) LogicMonitorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logicMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Newrelic() SloObjectiveCountMetricsBadNewrelicList {
	var returns SloObjectiveCountMetricsBadNewrelicList
	_jsii_.Get(
		j,
		"newrelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) NewrelicInput() any {
	var returns any
	_jsii_.Get(
		j,
		"newrelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Opentsdb() SloObjectiveCountMetricsBadOpentsdbList {
	var returns SloObjectiveCountMetricsBadOpentsdbList
	_jsii_.Get(
		j,
		"opentsdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) OpentsdbInput() any {
	var returns any
	_jsii_.Get(
		j,
		"opentsdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Pingdom() SloObjectiveCountMetricsBadPingdomList {
	var returns SloObjectiveCountMetricsBadPingdomList
	_jsii_.Get(
		j,
		"pingdom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PingdomInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pingdomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Prometheus() SloObjectiveCountMetricsBadPrometheusList {
	var returns SloObjectiveCountMetricsBadPrometheusList
	_jsii_.Get(
		j,
		"prometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PrometheusInput() any {
	var returns any
	_jsii_.Get(
		j,
		"prometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Redshift() SloObjectiveCountMetricsBadRedshiftList {
	var returns SloObjectiveCountMetricsBadRedshiftList
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) RedshiftInput() any {
	var returns any
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Splunk() SloObjectiveCountMetricsBadSplunkList {
	var returns SloObjectiveCountMetricsBadSplunkList
	_jsii_.Get(
		j,
		"splunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SplunkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"splunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SplunkObservability() SloObjectiveCountMetricsBadSplunkObservabilityList {
	var returns SloObjectiveCountMetricsBadSplunkObservabilityList
	_jsii_.Get(
		j,
		"splunkObservability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SplunkObservabilityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"splunkObservabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Sumologic() SloObjectiveCountMetricsBadSumologicList {
	var returns SloObjectiveCountMetricsBadSumologicList
	_jsii_.Get(
		j,
		"sumologic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SumologicInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sumologicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Thousandeyes() SloObjectiveCountMetricsBadThousandeyesList {
	var returns SloObjectiveCountMetricsBadThousandeyesList
	_jsii_.Get(
		j,
		"thousandeyes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ThousandeyesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"thousandeyesInput",
		&returns,
	)
	return returns
}

func NewSloObjectiveCountMetricsBadOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveCountMetricsBadOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveCountMetricsBadOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveCountMetricsBadOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsBadOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsBadOutputReference_Override(s SloObjectiveCountMetricsBadOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsBadOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutAmazonPrometheus(value any) {
	if err := s.validatePutAmazonPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAmazonPrometheus",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutAppdynamics(value any) {
	if err := s.validatePutAppdynamicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppdynamics",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutAzureMonitor(value any) {
	if err := s.validatePutAzureMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureMonitor",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutBigquery(value any) {
	if err := s.validatePutBigqueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBigquery",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutCloudwatch(value any) {
	if err := s.validatePutCloudwatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatch",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutDatadog(value any) {
	if err := s.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatadog",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutDynatrace(value any) {
	if err := s.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynatrace",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutElasticsearch(value any) {
	if err := s.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElasticsearch",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutGcm(value any) {
	if err := s.validatePutGcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcm",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutGrafanaLoki(value any) {
	if err := s.validatePutGrafanaLokiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrafanaLoki",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutGraphite(value any) {
	if err := s.validatePutGraphiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGraphite",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutHoneycomb(value any) {
	if err := s.validatePutHoneycombParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHoneycomb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutInfluxdb(value any) {
	if err := s.validatePutInfluxdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfluxdb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutInstana(value any) {
	if err := s.validatePutInstanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstana",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutLightstep(value any) {
	if err := s.validatePutLightstepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLightstep",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutLogicMonitor(value any) {
	if err := s.validatePutLogicMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogicMonitor",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutNewrelic(value any) {
	if err := s.validatePutNewrelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNewrelic",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutOpentsdb(value any) {
	if err := s.validatePutOpentsdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOpentsdb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutPingdom(value any) {
	if err := s.validatePutPingdomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPingdom",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutPrometheus(value any) {
	if err := s.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrometheus",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutRedshift(value any) {
	if err := s.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedshift",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutSplunk(value any) {
	if err := s.validatePutSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunk",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutSplunkObservability(value any) {
	if err := s.validatePutSplunkObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunkObservability",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutSumologic(value any) {
	if err := s.validatePutSumologicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSumologic",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) PutThousandeyes(value any) {
	if err := s.validatePutThousandeyesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThousandeyes",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetAmazonPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetAmazonPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetAppdynamics() {
	_jsii_.InvokeVoid(
		s,
		"resetAppdynamics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetAzureMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetBigquery() {
	_jsii_.InvokeVoid(
		s,
		"resetBigquery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetCloudwatch() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		s,
		"resetDatadog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		s,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		s,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetGcm() {
	_jsii_.InvokeVoid(
		s,
		"resetGcm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetGrafanaLoki() {
	_jsii_.InvokeVoid(
		s,
		"resetGrafanaLoki",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		s,
		"resetGraphite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetHoneycomb() {
	_jsii_.InvokeVoid(
		s,
		"resetHoneycomb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetInfluxdb() {
	_jsii_.InvokeVoid(
		s,
		"resetInfluxdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetInstana() {
	_jsii_.InvokeVoid(
		s,
		"resetInstana",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetLightstep() {
	_jsii_.InvokeVoid(
		s,
		"resetLightstep",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetLogicMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetLogicMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetNewrelic() {
	_jsii_.InvokeVoid(
		s,
		"resetNewrelic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetOpentsdb() {
	_jsii_.InvokeVoid(
		s,
		"resetOpentsdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetPingdom() {
	_jsii_.InvokeVoid(
		s,
		"resetPingdom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		s,
		"resetRedshift",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetSplunk() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetSplunkObservability() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunkObservability",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetSumologic() {
	_jsii_.InvokeVoid(
		s,
		"resetSumologic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ResetThousandeyes() {
	_jsii_.InvokeVoid(
		s,
		"resetThousandeyes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_SloObjectiveCountMetricsBadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
