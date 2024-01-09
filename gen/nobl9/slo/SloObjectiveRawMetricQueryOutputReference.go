package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveRawMetricQueryOutputReference interface {
	cdktf.ComplexObject
	AmazonPrometheus() SloObjectiveRawMetricQueryAmazonPrometheusList
	AmazonPrometheusInput() interface{}
	Appdynamics() SloObjectiveRawMetricQueryAppdynamicsList
	AppdynamicsInput() interface{}
	AzureMonitor() SloObjectiveRawMetricQueryAzureMonitorList
	AzureMonitorInput() interface{}
	Bigquery() SloObjectiveRawMetricQueryBigqueryList
	BigqueryInput() interface{}
	Cloudwatch() SloObjectiveRawMetricQueryCloudwatchList
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
	Datadog() SloObjectiveRawMetricQueryDatadogList
	DatadogInput() interface{}
	Dynatrace() SloObjectiveRawMetricQueryDynatraceList
	DynatraceInput() interface{}
	Elasticsearch() SloObjectiveRawMetricQueryElasticsearchList
	ElasticsearchInput() interface{}
	// Experimental.
	Fqn() *string
	Gcm() SloObjectiveRawMetricQueryGcmList
	GcmInput() interface{}
	GrafanaLoki() SloObjectiveRawMetricQueryGrafanaLokiList
	GrafanaLokiInput() interface{}
	Graphite() SloObjectiveRawMetricQueryGraphiteList
	GraphiteInput() interface{}
	Influxdb() SloObjectiveRawMetricQueryInfluxdbList
	InfluxdbInput() interface{}
	Instana() SloObjectiveRawMetricQueryInstanaList
	InstanaInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lightstep() SloObjectiveRawMetricQueryLightstepList
	LightstepInput() interface{}
	Newrelic() SloObjectiveRawMetricQueryNewrelicList
	NewrelicInput() interface{}
	Opentsdb() SloObjectiveRawMetricQueryOpentsdbList
	OpentsdbInput() interface{}
	Pingdom() SloObjectiveRawMetricQueryPingdomList
	PingdomInput() interface{}
	Prometheus() SloObjectiveRawMetricQueryPrometheusList
	PrometheusInput() interface{}
	Redshift() SloObjectiveRawMetricQueryRedshiftList
	RedshiftInput() interface{}
	Splunk() SloObjectiveRawMetricQuerySplunkList
	SplunkInput() interface{}
	SplunkObservability() SloObjectiveRawMetricQuerySplunkObservabilityList
	SplunkObservabilityInput() interface{}
	Sumologic() SloObjectiveRawMetricQuerySumologicList
	SumologicInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thousandeyes() SloObjectiveRawMetricQueryThousandeyesList
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
	PutInfluxdb(value interface{})
	PutInstana(value interface{})
	PutLightstep(value interface{})
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
	ResetInfluxdb()
	ResetInstana()
	ResetLightstep()
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

// The jsii proxy struct for SloObjectiveRawMetricQueryOutputReference
type jsiiProxy_SloObjectiveRawMetricQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) AmazonPrometheus() SloObjectiveRawMetricQueryAmazonPrometheusList {
	var returns SloObjectiveRawMetricQueryAmazonPrometheusList
	_jsii_.Get(
		j,
		"amazonPrometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) AmazonPrometheusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonPrometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Appdynamics() SloObjectiveRawMetricQueryAppdynamicsList {
	var returns SloObjectiveRawMetricQueryAppdynamicsList
	_jsii_.Get(
		j,
		"appdynamics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) AppdynamicsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appdynamicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) AzureMonitor() SloObjectiveRawMetricQueryAzureMonitorList {
	var returns SloObjectiveRawMetricQueryAzureMonitorList
	_jsii_.Get(
		j,
		"azureMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) AzureMonitorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Bigquery() SloObjectiveRawMetricQueryBigqueryList {
	var returns SloObjectiveRawMetricQueryBigqueryList
	_jsii_.Get(
		j,
		"bigquery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) BigqueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bigqueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Cloudwatch() SloObjectiveRawMetricQueryCloudwatchList {
	var returns SloObjectiveRawMetricQueryCloudwatchList
	_jsii_.Get(
		j,
		"cloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) CloudwatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Datadog() SloObjectiveRawMetricQueryDatadogList {
	var returns SloObjectiveRawMetricQueryDatadogList
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) DatadogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Dynatrace() SloObjectiveRawMetricQueryDynatraceList {
	var returns SloObjectiveRawMetricQueryDynatraceList
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) DynatraceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Elasticsearch() SloObjectiveRawMetricQueryElasticsearchList {
	var returns SloObjectiveRawMetricQueryElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Gcm() SloObjectiveRawMetricQueryGcmList {
	var returns SloObjectiveRawMetricQueryGcmList
	_jsii_.Get(
		j,
		"gcm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GcmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GrafanaLoki() SloObjectiveRawMetricQueryGrafanaLokiList {
	var returns SloObjectiveRawMetricQueryGrafanaLokiList
	_jsii_.Get(
		j,
		"grafanaLoki",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GrafanaLokiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grafanaLokiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Graphite() SloObjectiveRawMetricQueryGraphiteList {
	var returns SloObjectiveRawMetricQueryGraphiteList
	_jsii_.Get(
		j,
		"graphite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GraphiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graphiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Influxdb() SloObjectiveRawMetricQueryInfluxdbList {
	var returns SloObjectiveRawMetricQueryInfluxdbList
	_jsii_.Get(
		j,
		"influxdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) InfluxdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"influxdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Instana() SloObjectiveRawMetricQueryInstanaList {
	var returns SloObjectiveRawMetricQueryInstanaList
	_jsii_.Get(
		j,
		"instana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) InstanaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Lightstep() SloObjectiveRawMetricQueryLightstepList {
	var returns SloObjectiveRawMetricQueryLightstepList
	_jsii_.Get(
		j,
		"lightstep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) LightstepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lightstepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Newrelic() SloObjectiveRawMetricQueryNewrelicList {
	var returns SloObjectiveRawMetricQueryNewrelicList
	_jsii_.Get(
		j,
		"newrelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) NewrelicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newrelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Opentsdb() SloObjectiveRawMetricQueryOpentsdbList {
	var returns SloObjectiveRawMetricQueryOpentsdbList
	_jsii_.Get(
		j,
		"opentsdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) OpentsdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opentsdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Pingdom() SloObjectiveRawMetricQueryPingdomList {
	var returns SloObjectiveRawMetricQueryPingdomList
	_jsii_.Get(
		j,
		"pingdom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PingdomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pingdomInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Prometheus() SloObjectiveRawMetricQueryPrometheusList {
	var returns SloObjectiveRawMetricQueryPrometheusList
	_jsii_.Get(
		j,
		"prometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PrometheusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Redshift() SloObjectiveRawMetricQueryRedshiftList {
	var returns SloObjectiveRawMetricQueryRedshiftList
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) RedshiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Splunk() SloObjectiveRawMetricQuerySplunkList {
	var returns SloObjectiveRawMetricQuerySplunkList
	_jsii_.Get(
		j,
		"splunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) SplunkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) SplunkObservability() SloObjectiveRawMetricQuerySplunkObservabilityList {
	var returns SloObjectiveRawMetricQuerySplunkObservabilityList
	_jsii_.Get(
		j,
		"splunkObservability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) SplunkObservabilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkObservabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Sumologic() SloObjectiveRawMetricQuerySumologicList {
	var returns SloObjectiveRawMetricQuerySumologicList
	_jsii_.Get(
		j,
		"sumologic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) SumologicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumologicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Thousandeyes() SloObjectiveRawMetricQueryThousandeyesList {
	var returns SloObjectiveRawMetricQueryThousandeyesList
	_jsii_.Get(
		j,
		"thousandeyes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ThousandeyesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thousandeyesInput",
		&returns,
	)
	return returns
}


func NewSloObjectiveRawMetricQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveRawMetricQueryOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveRawMetricQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveRawMetricQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveRawMetricQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveRawMetricQueryOutputReference_Override(s SloObjectiveRawMetricQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveRawMetricQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveRawMetricQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutAmazonPrometheus(value interface{}) {
	if err := s.validatePutAmazonPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAmazonPrometheus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutAppdynamics(value interface{}) {
	if err := s.validatePutAppdynamicsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAppdynamics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutAzureMonitor(value interface{}) {
	if err := s.validatePutAzureMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureMonitor",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutBigquery(value interface{}) {
	if err := s.validatePutBigqueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBigquery",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutCloudwatch(value interface{}) {
	if err := s.validatePutCloudwatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatch",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutDatadog(value interface{}) {
	if err := s.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDatadog",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutDynatrace(value interface{}) {
	if err := s.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutElasticsearch(value interface{}) {
	if err := s.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutGcm(value interface{}) {
	if err := s.validatePutGcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcm",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutGrafanaLoki(value interface{}) {
	if err := s.validatePutGrafanaLokiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGrafanaLoki",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutGraphite(value interface{}) {
	if err := s.validatePutGraphiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGraphite",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutInfluxdb(value interface{}) {
	if err := s.validatePutInfluxdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfluxdb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutInstana(value interface{}) {
	if err := s.validatePutInstanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstana",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutLightstep(value interface{}) {
	if err := s.validatePutLightstepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLightstep",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutNewrelic(value interface{}) {
	if err := s.validatePutNewrelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNewrelic",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutOpentsdb(value interface{}) {
	if err := s.validatePutOpentsdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOpentsdb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutPingdom(value interface{}) {
	if err := s.validatePutPingdomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPingdom",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutPrometheus(value interface{}) {
	if err := s.validatePutPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrometheus",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutRedshift(value interface{}) {
	if err := s.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedshift",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutSplunk(value interface{}) {
	if err := s.validatePutSplunkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutSplunkObservability(value interface{}) {
	if err := s.validatePutSplunkObservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSplunkObservability",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutSumologic(value interface{}) {
	if err := s.validatePutSumologicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSumologic",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) PutThousandeyes(value interface{}) {
	if err := s.validatePutThousandeyesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThousandeyes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetAmazonPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetAmazonPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetAppdynamics() {
	_jsii_.InvokeVoid(
		s,
		"resetAppdynamics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetAzureMonitor() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureMonitor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetBigquery() {
	_jsii_.InvokeVoid(
		s,
		"resetBigquery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetCloudwatch() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		s,
		"resetDatadog",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		s,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		s,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetGcm() {
	_jsii_.InvokeVoid(
		s,
		"resetGcm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetGrafanaLoki() {
	_jsii_.InvokeVoid(
		s,
		"resetGrafanaLoki",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetGraphite() {
	_jsii_.InvokeVoid(
		s,
		"resetGraphite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetInfluxdb() {
	_jsii_.InvokeVoid(
		s,
		"resetInfluxdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetInstana() {
	_jsii_.InvokeVoid(
		s,
		"resetInstana",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetLightstep() {
	_jsii_.InvokeVoid(
		s,
		"resetLightstep",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetNewrelic() {
	_jsii_.InvokeVoid(
		s,
		"resetNewrelic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetOpentsdb() {
	_jsii_.InvokeVoid(
		s,
		"resetOpentsdb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetPingdom() {
	_jsii_.InvokeVoid(
		s,
		"resetPingdom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetPrometheus() {
	_jsii_.InvokeVoid(
		s,
		"resetPrometheus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		s,
		"resetRedshift",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetSplunk() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetSplunkObservability() {
	_jsii_.InvokeVoid(
		s,
		"resetSplunkObservability",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetSumologic() {
	_jsii_.InvokeVoid(
		s,
		"resetSumologic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ResetThousandeyes() {
	_jsii_.InvokeVoid(
		s,
		"resetThousandeyes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SloObjectiveRawMetricQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

