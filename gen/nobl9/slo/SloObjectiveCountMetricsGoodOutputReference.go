package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveCountMetricsGoodOutputReference interface {
	cdktf.ComplexObject
	AmazonPrometheus() SloObjectiveCountMetricsGoodAmazonPrometheusList
	AmazonPrometheusInput() any
	Appdynamics() SloObjectiveCountMetricsGoodAppdynamicsList
	AppdynamicsInput() any
	AzureMonitor() SloObjectiveCountMetricsGoodAzureMonitorList
	AzureMonitorInput() any
	Bigquery() SloObjectiveCountMetricsGoodBigqueryList
	BigqueryInput() any
	Cloudwatch() SloObjectiveCountMetricsGoodCloudwatchList
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
	Datadog() SloObjectiveCountMetricsGoodDatadogList
	DatadogInput() any
	Dynatrace() SloObjectiveCountMetricsGoodDynatraceList
	DynatraceInput() any
	Elasticsearch() SloObjectiveCountMetricsGoodElasticsearchList
	ElasticsearchInput() any
	// Experimental.
	Fqn() *string
	Gcm() SloObjectiveCountMetricsGoodGcmList
	GcmInput() any
	GrafanaLoki() SloObjectiveCountMetricsGoodGrafanaLokiList
	GrafanaLokiInput() any
	Graphite() SloObjectiveCountMetricsGoodGraphiteList
	GraphiteInput() any
	Honeycomb() SloObjectiveCountMetricsGoodHoneycombList
	HoneycombInput() any
	Influxdb() SloObjectiveCountMetricsGoodInfluxdbList
	InfluxdbInput() any
	Instana() SloObjectiveCountMetricsGoodInstanaList
	InstanaInput() any
	InternalValue() any
	SetInternalValue(val any)
	Lightstep() SloObjectiveCountMetricsGoodLightstepList
	LightstepInput() any
	LogicMonitor() SloObjectiveCountMetricsGoodLogicMonitorList
	LogicMonitorInput() any
	Newrelic() SloObjectiveCountMetricsGoodNewrelicList
	NewrelicInput() any
	Opentsdb() SloObjectiveCountMetricsGoodOpentsdbList
	OpentsdbInput() any
	Pingdom() SloObjectiveCountMetricsGoodPingdomList
	PingdomInput() any
	Prometheus() SloObjectiveCountMetricsGoodPrometheusList
	PrometheusInput() any
	Redshift() SloObjectiveCountMetricsGoodRedshiftList
	RedshiftInput() any
	Splunk() SloObjectiveCountMetricsGoodSplunkList
	SplunkInput() any
	SplunkObservability() SloObjectiveCountMetricsGoodSplunkObservabilityList
	SplunkObservabilityInput() any
	Sumologic() SloObjectiveCountMetricsGoodSumologicList
	SumologicInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thousandeyes() SloObjectiveCountMetricsGoodThousandeyesList
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

// The jsii proxy struct for SloObjectiveCountMetricsGoodOutputReference
type jsiiProxy_SloObjectiveCountMetricsGoodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) AmazonPrometheus() SloObjectiveCountMetricsGoodAmazonPrometheusList {
	var returns SloObjectiveCountMetricsGoodAmazonPrometheusList
	_jsii_.Get(
		j,
		"amazonPrometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) AmazonPrometheusInput() any {
	var returns any
	_jsii_.Get(
		j,
		"amazonPrometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Appdynamics() SloObjectiveCountMetricsGoodAppdynamicsList {
	var returns SloObjectiveCountMetricsGoodAppdynamicsList
	_jsii_.Get(
		j,
		"appdynamics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) AppdynamicsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"appdynamicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) AzureMonitor() SloObjectiveCountMetricsGoodAzureMonitorList {
	var returns SloObjectiveCountMetricsGoodAzureMonitorList
	_jsii_.Get(
		j,
		"azureMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) AzureMonitorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"azureMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Bigquery() SloObjectiveCountMetricsGoodBigqueryList {
	var returns SloObjectiveCountMetricsGoodBigqueryList
	_jsii_.Get(
		j,
		"bigquery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) BigqueryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"bigqueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Cloudwatch() SloObjectiveCountMetricsGoodCloudwatchList {
	var returns SloObjectiveCountMetricsGoodCloudwatchList
	_jsii_.Get(
		j,
		"cloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) CloudwatchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"cloudwatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Datadog() SloObjectiveCountMetricsGoodDatadogList {
	var returns SloObjectiveCountMetricsGoodDatadogList
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) DatadogInput() any {
	var returns any
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Dynatrace() SloObjectiveCountMetricsGoodDynatraceList {
	var returns SloObjectiveCountMetricsGoodDynatraceList
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) DynatraceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Elasticsearch() SloObjectiveCountMetricsGoodElasticsearchList {
	var returns SloObjectiveCountMetricsGoodElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ElasticsearchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Gcm() SloObjectiveCountMetricsGoodGcmList {
	var returns SloObjectiveCountMetricsGoodGcmList
	_jsii_.Get(
		j,
		"gcm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GcmInput() any {
	var returns any
	_jsii_.Get(
		j,
		"gcmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GrafanaLoki() SloObjectiveCountMetricsGoodGrafanaLokiList {
	var returns SloObjectiveCountMetricsGoodGrafanaLokiList
	_jsii_.Get(
		j,
		"grafanaLoki",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GrafanaLokiInput() any {
	var returns any
	_jsii_.Get(
		j,
		"grafanaLokiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Graphite() SloObjectiveCountMetricsGoodGraphiteList {
	var returns SloObjectiveCountMetricsGoodGraphiteList
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GraphiteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Honeycomb() SloObjectiveCountMetricsGoodHoneycombList {
	var returns SloObjectiveCountMetricsGoodHoneycombList
	_jsii_.Get(
		j,
		"honeycomb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) HoneycombInput() any {
	var returns any
	_jsii_.Get(
		j,
		"honeycombInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Influxdb() SloObjectiveCountMetricsGoodInfluxdbList {
	var returns SloObjectiveCountMetricsGoodInfluxdbList
	_jsii_.Get(
		j,
		"influxdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) InfluxdbInput() any {
	var returns any
	_jsii_.Get(
		j,
		"influxdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Instana() SloObjectiveCountMetricsGoodInstanaList {
	var returns SloObjectiveCountMetricsGoodInstanaList
	_jsii_.Get(
		j,
		"instana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) InstanaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"instanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Lightstep() SloObjectiveCountMetricsGoodLightstepList {
	var returns SloObjectiveCountMetricsGoodLightstepList
	_jsii_.Get(
		j,
		"lightstep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) LightstepInput() any {
	var returns any
	_jsii_.Get(
		j,
		"lightstepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) LogicMonitor() SloObjectiveCountMetricsGoodLogicMonitorList {
	var returns SloObjectiveCountMetricsGoodLogicMonitorList
	_jsii_.Get(
		j,
		"logicMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) LogicMonitorInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logicMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Newrelic() SloObjectiveCountMetricsGoodNewrelicList {
	var returns SloObjectiveCountMetricsGoodNewrelicList
	_jsii_.Get(
		j,
		"newrelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) NewrelicInput() any {
	var returns any
	_jsii_.Get(
		j,
		"newrelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Opentsdb() SloObjectiveCountMetricsGoodOpentsdbList {
	var returns SloObjectiveCountMetricsGoodOpentsdbList
	_jsii_.Get(
		j,
		"opentsdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) OpentsdbInput() any {
	var returns any
	_jsii_.Get(
		j,
		"opentsdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Pingdom() SloObjectiveCountMetricsGoodPingdomList {
	var returns SloObjectiveCountMetricsGoodPingdomList
	_jsii_.Get(
		j,
		"pingdom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PingdomInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pingdomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Prometheus() SloObjectiveCountMetricsGoodPrometheusList {
	var returns SloObjectiveCountMetricsGoodPrometheusList
	_jsii_.Get(
		j,
		"prometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PrometheusInput() any {
	var returns any
	_jsii_.Get(
		j,
		"prometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Redshift() SloObjectiveCountMetricsGoodRedshiftList {
	var returns SloObjectiveCountMetricsGoodRedshiftList
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) RedshiftInput() any {
	var returns any
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Splunk() SloObjectiveCountMetricsGoodSplunkList {
	var returns SloObjectiveCountMetricsGoodSplunkList
	_jsii_.Get(
		j,
		"splunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SplunkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"splunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SplunkObservability() SloObjectiveCountMetricsGoodSplunkObservabilityList {
	var returns SloObjectiveCountMetricsGoodSplunkObservabilityList
	_jsii_.Get(
		j,
		"splunkObservability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SplunkObservabilityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"splunkObservabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Sumologic() SloObjectiveCountMetricsGoodSumologicList {
	var returns SloObjectiveCountMetricsGoodSumologicList
	_jsii_.Get(
		j,
		"sumologic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SumologicInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sumologicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Thousandeyes() SloObjectiveCountMetricsGoodThousandeyesList {
	var returns SloObjectiveCountMetricsGoodThousandeyesList
	_jsii_.Get(
		j,
		"thousandeyes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ThousandeyesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"thousandeyesInput",
		&returns,
	)
	return returns
}

func NewSloObjectiveCountMetricsGoodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveCountMetricsGoodOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveCountMetricsGoodOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveCountMetricsGoodOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsGoodOutputReference_Override(s SloObjectiveCountMetricsGoodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutAmazonPrometheus(value any) {
	if err := s.validatePutAmazonPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAmazonPrometheus",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutAppdynamics(value any) {
	if err := s.validatePutAppdynamicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppdynamics",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutAzureMonitor(value any) {
	if err := s.validatePutAzureMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureMonitor",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutBigquery(value any) {
	if err := s.validatePutBigqueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBigquery",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutCloudwatch(value any) {
	if err := s.validatePutCloudwatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatch",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutDatadog(value any) {
	if err := s.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatadog",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutDynatrace(value any) {
	if err := s.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynatrace",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutElasticsearch(value any) {
	if err := s.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElasticsearch",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutGcm(value any) {
	if err := s.validatePutGcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcm",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutGrafanaLoki(value any) {
	if err := s.validatePutGrafanaLokiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrafanaLoki",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutGraphite(value any) {
	if err := s.validatePutGraphiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGraphite",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutHoneycomb(value any) {
	if err := s.validatePutHoneycombParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHoneycomb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutInfluxdb(value any) {
	if err := s.validatePutInfluxdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfluxdb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutInstana(value any) {
	if err := s.validatePutInstanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstana",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutLightstep(value any) {
	if err := s.validatePutLightstepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLightstep",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutLogicMonitor(value any) {
	if err := s.validatePutLogicMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogicMonitor",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutNewrelic(value any) {
	if err := s.validatePutNewrelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNewrelic",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutOpentsdb(value any) {
	if err := s.validatePutOpentsdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOpentsdb",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutPingdom(value any) {
	if err := s.validatePutPingdomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPingdom",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutPrometheus(value any) {
	if err := s.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrometheus",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutRedshift(value any) {
	if err := s.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedshift",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutSplunk(value any) {
	if err := s.validatePutSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunk",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutSplunkObservability(value any) {
	if err := s.validatePutSplunkObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunkObservability",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutSumologic(value any) {
	if err := s.validatePutSumologicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSumologic",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) PutThousandeyes(value any) {
	if err := s.validatePutThousandeyesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThousandeyes",
		[]any{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetAmazonPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetAmazonPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetAppdynamics() {
	_jsii_.InvokeVoid(
		s,
		"resetAppdynamics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetAzureMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetBigquery() {
	_jsii_.InvokeVoid(
		s,
		"resetBigquery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetCloudwatch() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		s,
		"resetDatadog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		s,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		s,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetGcm() {
	_jsii_.InvokeVoid(
		s,
		"resetGcm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetGrafanaLoki() {
	_jsii_.InvokeVoid(
		s,
		"resetGrafanaLoki",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		s,
		"resetGraphite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetHoneycomb() {
	_jsii_.InvokeVoid(
		s,
		"resetHoneycomb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetInfluxdb() {
	_jsii_.InvokeVoid(
		s,
		"resetInfluxdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetInstana() {
	_jsii_.InvokeVoid(
		s,
		"resetInstana",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetLightstep() {
	_jsii_.InvokeVoid(
		s,
		"resetLightstep",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetLogicMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetLogicMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetNewrelic() {
	_jsii_.InvokeVoid(
		s,
		"resetNewrelic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetOpentsdb() {
	_jsii_.InvokeVoid(
		s,
		"resetOpentsdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetPingdom() {
	_jsii_.InvokeVoid(
		s,
		"resetPingdom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		s,
		"resetRedshift",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetSplunk() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetSplunkObservability() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunkObservability",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetSumologic() {
	_jsii_.InvokeVoid(
		s,
		"resetSumologic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ResetThousandeyes() {
	_jsii_.InvokeVoid(
		s,
		"resetThousandeyes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
