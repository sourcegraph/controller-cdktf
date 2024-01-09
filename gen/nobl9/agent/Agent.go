package agent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/agent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent nobl9_agent}.
type Agent interface {
	cdktf.TerraformResource
	AgentType() *string
	SetAgentType(val *string)
	AgentTypeInput() *string
	AmazonPrometheusConfig() AgentAmazonPrometheusConfigOutputReference
	AmazonPrometheusConfigInput() *AgentAmazonPrometheusConfig
	AppdynamicsConfig() AgentAppdynamicsConfigOutputReference
	AppdynamicsConfigInput() *AgentAppdynamicsConfig
	AzureMonitorConfig() AgentAzureMonitorConfigOutputReference
	AzureMonitorConfigInput() *AgentAzureMonitorConfig
	BigqueryConfig() AgentBigqueryConfigOutputReference
	BigqueryConfigInput() *AgentBigqueryConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	ClientSecret() *string
	CloudwatchConfig() AgentCloudwatchConfigOutputReference
	CloudwatchConfigInput() *AgentCloudwatchConfig
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatadogConfig() AgentDatadogConfigOutputReference
	DatadogConfigInput() *AgentDatadogConfig
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DynatraceConfig() AgentDynatraceConfigOutputReference
	DynatraceConfigInput() *AgentDynatraceConfig
	ElasticsearchConfig() AgentElasticsearchConfigOutputReference
	ElasticsearchConfigInput() *AgentElasticsearchConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcmConfig() AgentGcmConfigOutputReference
	GcmConfigInput() *AgentGcmConfig
	GrafanaLokiConfig() AgentGrafanaLokiConfigOutputReference
	GrafanaLokiConfigInput() *AgentGrafanaLokiConfig
	GraphiteConfig() AgentGraphiteConfigOutputReference
	GraphiteConfigInput() *AgentGraphiteConfig
	HistoricalDataRetrieval() AgentHistoricalDataRetrievalOutputReference
	HistoricalDataRetrievalInput() *AgentHistoricalDataRetrieval
	Id() *string
	SetId(val *string)
	IdInput() *string
	InfluxdbConfig() AgentInfluxdbConfigOutputReference
	InfluxdbConfigInput() *AgentInfluxdbConfig
	InstanaConfig() AgentInstanaConfigOutputReference
	InstanaConfigInput() *AgentInstanaConfig
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LightstepConfig() AgentLightstepConfigOutputReference
	LightstepConfigInput() *AgentLightstepConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewrelicConfig() AgentNewrelicConfigOutputReference
	NewrelicConfigInput() *AgentNewrelicConfig
	// The tree node.
	Node() constructs.Node
	OpentsdbConfig() AgentOpentsdbConfigOutputReference
	OpentsdbConfigInput() *AgentOpentsdbConfig
	PingdomConfig() AgentPingdomConfigOutputReference
	PingdomConfigInput() *AgentPingdomConfig
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PrometheusConfig() AgentPrometheusConfigOutputReference
	PrometheusConfigInput() *AgentPrometheusConfig
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueryDelay() AgentQueryDelayOutputReference
	QueryDelayInput() *AgentQueryDelay
	// Experimental.
	RawOverrides() interface{}
	RedshiftConfig() AgentRedshiftConfigOutputReference
	RedshiftConfigInput() *AgentRedshiftConfig
	ReleaseChannel() *string
	SetReleaseChannel(val *string)
	ReleaseChannelInput() *string
	SourceOf() *[]*string
	SetSourceOf(val *[]*string)
	SourceOfInput() *[]*string
	SplunkConfig() AgentSplunkConfigOutputReference
	SplunkConfigInput() *AgentSplunkConfig
	SplunkObservabilityConfig() AgentSplunkObservabilityConfigOutputReference
	SplunkObservabilityConfigInput() *AgentSplunkObservabilityConfig
	Status() cdktf.StringMap
	SumologicConfig() AgentSumologicConfigOutputReference
	SumologicConfigInput() *AgentSumologicConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThousandeyesConfig() AgentThousandeyesConfigOutputReference
	ThousandeyesConfigInput() *AgentThousandeyesConfig
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAmazonPrometheusConfig(value *AgentAmazonPrometheusConfig)
	PutAppdynamicsConfig(value *AgentAppdynamicsConfig)
	PutAzureMonitorConfig(value *AgentAzureMonitorConfig)
	PutBigqueryConfig(value *AgentBigqueryConfig)
	PutCloudwatchConfig(value *AgentCloudwatchConfig)
	PutDatadogConfig(value *AgentDatadogConfig)
	PutDynatraceConfig(value *AgentDynatraceConfig)
	PutElasticsearchConfig(value *AgentElasticsearchConfig)
	PutGcmConfig(value *AgentGcmConfig)
	PutGrafanaLokiConfig(value *AgentGrafanaLokiConfig)
	PutGraphiteConfig(value *AgentGraphiteConfig)
	PutHistoricalDataRetrieval(value *AgentHistoricalDataRetrieval)
	PutInfluxdbConfig(value *AgentInfluxdbConfig)
	PutInstanaConfig(value *AgentInstanaConfig)
	PutLightstepConfig(value *AgentLightstepConfig)
	PutNewrelicConfig(value *AgentNewrelicConfig)
	PutOpentsdbConfig(value *AgentOpentsdbConfig)
	PutPingdomConfig(value *AgentPingdomConfig)
	PutPrometheusConfig(value *AgentPrometheusConfig)
	PutQueryDelay(value *AgentQueryDelay)
	PutRedshiftConfig(value *AgentRedshiftConfig)
	PutSplunkConfig(value *AgentSplunkConfig)
	PutSplunkObservabilityConfig(value *AgentSplunkObservabilityConfig)
	PutSumologicConfig(value *AgentSumologicConfig)
	PutThousandeyesConfig(value *AgentThousandeyesConfig)
	ResetAmazonPrometheusConfig()
	ResetAppdynamicsConfig()
	ResetAzureMonitorConfig()
	ResetBigqueryConfig()
	ResetCloudwatchConfig()
	ResetDatadogConfig()
	ResetDescription()
	ResetDisplayName()
	ResetDynatraceConfig()
	ResetElasticsearchConfig()
	ResetGcmConfig()
	ResetGrafanaLokiConfig()
	ResetGraphiteConfig()
	ResetHistoricalDataRetrieval()
	ResetId()
	ResetInfluxdbConfig()
	ResetInstanaConfig()
	ResetLightstepConfig()
	ResetNewrelicConfig()
	ResetOpentsdbConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPingdomConfig()
	ResetPrometheusConfig()
	ResetQueryDelay()
	ResetRedshiftConfig()
	ResetReleaseChannel()
	ResetSplunkConfig()
	ResetSplunkObservabilityConfig()
	ResetSumologicConfig()
	ResetThousandeyesConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Agent
type jsiiProxy_Agent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Agent) AgentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AgentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AmazonPrometheusConfig() AgentAmazonPrometheusConfigOutputReference {
	var returns AgentAmazonPrometheusConfigOutputReference
	_jsii_.Get(
		j,
		"amazonPrometheusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AmazonPrometheusConfigInput() *AgentAmazonPrometheusConfig {
	var returns *AgentAmazonPrometheusConfig
	_jsii_.Get(
		j,
		"amazonPrometheusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AppdynamicsConfig() AgentAppdynamicsConfigOutputReference {
	var returns AgentAppdynamicsConfigOutputReference
	_jsii_.Get(
		j,
		"appdynamicsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AppdynamicsConfigInput() *AgentAppdynamicsConfig {
	var returns *AgentAppdynamicsConfig
	_jsii_.Get(
		j,
		"appdynamicsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AzureMonitorConfig() AgentAzureMonitorConfigOutputReference {
	var returns AgentAzureMonitorConfigOutputReference
	_jsii_.Get(
		j,
		"azureMonitorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AzureMonitorConfigInput() *AgentAzureMonitorConfig {
	var returns *AgentAzureMonitorConfig
	_jsii_.Get(
		j,
		"azureMonitorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) BigqueryConfig() AgentBigqueryConfigOutputReference {
	var returns AgentBigqueryConfigOutputReference
	_jsii_.Get(
		j,
		"bigqueryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) BigqueryConfigInput() *AgentBigqueryConfig {
	var returns *AgentBigqueryConfig
	_jsii_.Get(
		j,
		"bigqueryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) CloudwatchConfig() AgentCloudwatchConfigOutputReference {
	var returns AgentCloudwatchConfigOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) CloudwatchConfigInput() *AgentCloudwatchConfig {
	var returns *AgentCloudwatchConfig
	_jsii_.Get(
		j,
		"cloudwatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DatadogConfig() AgentDatadogConfigOutputReference {
	var returns AgentDatadogConfigOutputReference
	_jsii_.Get(
		j,
		"datadogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DatadogConfigInput() *AgentDatadogConfig {
	var returns *AgentDatadogConfig
	_jsii_.Get(
		j,
		"datadogConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DynatraceConfig() AgentDynatraceConfigOutputReference {
	var returns AgentDynatraceConfigOutputReference
	_jsii_.Get(
		j,
		"dynatraceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) DynatraceConfigInput() *AgentDynatraceConfig {
	var returns *AgentDynatraceConfig
	_jsii_.Get(
		j,
		"dynatraceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ElasticsearchConfig() AgentElasticsearchConfigOutputReference {
	var returns AgentElasticsearchConfigOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ElasticsearchConfigInput() *AgentElasticsearchConfig {
	var returns *AgentElasticsearchConfig
	_jsii_.Get(
		j,
		"elasticsearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) GcmConfig() AgentGcmConfigOutputReference {
	var returns AgentGcmConfigOutputReference
	_jsii_.Get(
		j,
		"gcmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) GcmConfigInput() *AgentGcmConfig {
	var returns *AgentGcmConfig
	_jsii_.Get(
		j,
		"gcmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) GrafanaLokiConfig() AgentGrafanaLokiConfigOutputReference {
	var returns AgentGrafanaLokiConfigOutputReference
	_jsii_.Get(
		j,
		"grafanaLokiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) GrafanaLokiConfigInput() *AgentGrafanaLokiConfig {
	var returns *AgentGrafanaLokiConfig
	_jsii_.Get(
		j,
		"grafanaLokiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) GraphiteConfig() AgentGraphiteConfigOutputReference {
	var returns AgentGraphiteConfigOutputReference
	_jsii_.Get(
		j,
		"graphiteConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) GraphiteConfigInput() *AgentGraphiteConfig {
	var returns *AgentGraphiteConfig
	_jsii_.Get(
		j,
		"graphiteConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) HistoricalDataRetrieval() AgentHistoricalDataRetrievalOutputReference {
	var returns AgentHistoricalDataRetrievalOutputReference
	_jsii_.Get(
		j,
		"historicalDataRetrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) HistoricalDataRetrievalInput() *AgentHistoricalDataRetrieval {
	var returns *AgentHistoricalDataRetrieval
	_jsii_.Get(
		j,
		"historicalDataRetrievalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) InfluxdbConfig() AgentInfluxdbConfigOutputReference {
	var returns AgentInfluxdbConfigOutputReference
	_jsii_.Get(
		j,
		"influxdbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) InfluxdbConfigInput() *AgentInfluxdbConfig {
	var returns *AgentInfluxdbConfig
	_jsii_.Get(
		j,
		"influxdbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) InstanaConfig() AgentInstanaConfigOutputReference {
	var returns AgentInstanaConfigOutputReference
	_jsii_.Get(
		j,
		"instanaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) InstanaConfigInput() *AgentInstanaConfig {
	var returns *AgentInstanaConfig
	_jsii_.Get(
		j,
		"instanaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) LightstepConfig() AgentLightstepConfigOutputReference {
	var returns AgentLightstepConfigOutputReference
	_jsii_.Get(
		j,
		"lightstepConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) LightstepConfigInput() *AgentLightstepConfig {
	var returns *AgentLightstepConfig
	_jsii_.Get(
		j,
		"lightstepConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) NewrelicConfig() AgentNewrelicConfigOutputReference {
	var returns AgentNewrelicConfigOutputReference
	_jsii_.Get(
		j,
		"newrelicConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) NewrelicConfigInput() *AgentNewrelicConfig {
	var returns *AgentNewrelicConfig
	_jsii_.Get(
		j,
		"newrelicConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) OpentsdbConfig() AgentOpentsdbConfigOutputReference {
	var returns AgentOpentsdbConfigOutputReference
	_jsii_.Get(
		j,
		"opentsdbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) OpentsdbConfigInput() *AgentOpentsdbConfig {
	var returns *AgentOpentsdbConfig
	_jsii_.Get(
		j,
		"opentsdbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) PingdomConfig() AgentPingdomConfigOutputReference {
	var returns AgentPingdomConfigOutputReference
	_jsii_.Get(
		j,
		"pingdomConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) PingdomConfigInput() *AgentPingdomConfig {
	var returns *AgentPingdomConfig
	_jsii_.Get(
		j,
		"pingdomConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) PrometheusConfig() AgentPrometheusConfigOutputReference {
	var returns AgentPrometheusConfigOutputReference
	_jsii_.Get(
		j,
		"prometheusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) PrometheusConfigInput() *AgentPrometheusConfig {
	var returns *AgentPrometheusConfig
	_jsii_.Get(
		j,
		"prometheusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) QueryDelay() AgentQueryDelayOutputReference {
	var returns AgentQueryDelayOutputReference
	_jsii_.Get(
		j,
		"queryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) QueryDelayInput() *AgentQueryDelay {
	var returns *AgentQueryDelay
	_jsii_.Get(
		j,
		"queryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) RedshiftConfig() AgentRedshiftConfigOutputReference {
	var returns AgentRedshiftConfigOutputReference
	_jsii_.Get(
		j,
		"redshiftConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) RedshiftConfigInput() *AgentRedshiftConfig {
	var returns *AgentRedshiftConfig
	_jsii_.Get(
		j,
		"redshiftConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ReleaseChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SourceOf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SourceOfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SplunkConfig() AgentSplunkConfigOutputReference {
	var returns AgentSplunkConfigOutputReference
	_jsii_.Get(
		j,
		"splunkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SplunkConfigInput() *AgentSplunkConfig {
	var returns *AgentSplunkConfig
	_jsii_.Get(
		j,
		"splunkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SplunkObservabilityConfig() AgentSplunkObservabilityConfigOutputReference {
	var returns AgentSplunkObservabilityConfigOutputReference
	_jsii_.Get(
		j,
		"splunkObservabilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SplunkObservabilityConfigInput() *AgentSplunkObservabilityConfig {
	var returns *AgentSplunkObservabilityConfig
	_jsii_.Get(
		j,
		"splunkObservabilityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Status() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SumologicConfig() AgentSumologicConfigOutputReference {
	var returns AgentSumologicConfigOutputReference
	_jsii_.Get(
		j,
		"sumologicConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) SumologicConfigInput() *AgentSumologicConfig {
	var returns *AgentSumologicConfig
	_jsii_.Get(
		j,
		"sumologicConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ThousandeyesConfig() AgentThousandeyesConfigOutputReference {
	var returns AgentThousandeyesConfigOutputReference
	_jsii_.Get(
		j,
		"thousandeyesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) ThousandeyesConfigInput() *AgentThousandeyesConfig {
	var returns *AgentThousandeyesConfig
	_jsii_.Get(
		j,
		"thousandeyesConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent nobl9_agent} Resource.
func NewAgent(scope constructs.Construct, id *string, config *AgentConfig) Agent {
	_init_.Initialize()

	if err := validateNewAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Agent{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.agent.Agent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent nobl9_agent} Resource.
func NewAgent_Override(a Agent, scope constructs.Construct, id *string, config *AgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.agent.Agent",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Agent)SetAgentType(val *string) {
	if err := j.validateSetAgentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentType",
		val,
	)
}

func (j *jsiiProxy_Agent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Agent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Agent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Agent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Agent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Agent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Agent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Agent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Agent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Agent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_Agent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Agent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Agent)SetReleaseChannel(val *string) {
	if err := j.validateSetReleaseChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_Agent)SetSourceOf(val *[]*string) {
	if err := j.validateSetSourceOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceOf",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Agent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.agent.Agent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Agent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.agent.Agent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Agent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.agent.Agent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Agent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nobl9.agent.Agent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Agent) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Agent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Agent) PutAmazonPrometheusConfig(value *AgentAmazonPrometheusConfig) {
	if err := a.validatePutAmazonPrometheusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonPrometheusConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutAppdynamicsConfig(value *AgentAppdynamicsConfig) {
	if err := a.validatePutAppdynamicsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppdynamicsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutAzureMonitorConfig(value *AgentAzureMonitorConfig) {
	if err := a.validatePutAzureMonitorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAzureMonitorConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutBigqueryConfig(value *AgentBigqueryConfig) {
	if err := a.validatePutBigqueryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBigqueryConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutCloudwatchConfig(value *AgentCloudwatchConfig) {
	if err := a.validatePutCloudwatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutDatadogConfig(value *AgentDatadogConfig) {
	if err := a.validatePutDatadogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutDynatraceConfig(value *AgentDynatraceConfig) {
	if err := a.validatePutDynatraceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatraceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutElasticsearchConfig(value *AgentElasticsearchConfig) {
	if err := a.validatePutElasticsearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutGcmConfig(value *AgentGcmConfig) {
	if err := a.validatePutGcmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGcmConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutGrafanaLokiConfig(value *AgentGrafanaLokiConfig) {
	if err := a.validatePutGrafanaLokiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrafanaLokiConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutGraphiteConfig(value *AgentGraphiteConfig) {
	if err := a.validatePutGraphiteConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGraphiteConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutHistoricalDataRetrieval(value *AgentHistoricalDataRetrieval) {
	if err := a.validatePutHistoricalDataRetrievalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHistoricalDataRetrieval",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutInfluxdbConfig(value *AgentInfluxdbConfig) {
	if err := a.validatePutInfluxdbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInfluxdbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutInstanaConfig(value *AgentInstanaConfig) {
	if err := a.validatePutInstanaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutLightstepConfig(value *AgentLightstepConfig) {
	if err := a.validatePutLightstepConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLightstepConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutNewrelicConfig(value *AgentNewrelicConfig) {
	if err := a.validatePutNewrelicConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNewrelicConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutOpentsdbConfig(value *AgentOpentsdbConfig) {
	if err := a.validatePutOpentsdbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpentsdbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutPingdomConfig(value *AgentPingdomConfig) {
	if err := a.validatePutPingdomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPingdomConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutPrometheusConfig(value *AgentPrometheusConfig) {
	if err := a.validatePutPrometheusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrometheusConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutQueryDelay(value *AgentQueryDelay) {
	if err := a.validatePutQueryDelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryDelay",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutRedshiftConfig(value *AgentRedshiftConfig) {
	if err := a.validatePutRedshiftConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutSplunkConfig(value *AgentSplunkConfig) {
	if err := a.validatePutSplunkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSplunkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutSplunkObservabilityConfig(value *AgentSplunkObservabilityConfig) {
	if err := a.validatePutSplunkObservabilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSplunkObservabilityConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutSumologicConfig(value *AgentSumologicConfig) {
	if err := a.validatePutSumologicConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSumologicConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) PutThousandeyesConfig(value *AgentThousandeyesConfig) {
	if err := a.validatePutThousandeyesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThousandeyesConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Agent) ResetAmazonPrometheusConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonPrometheusConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetAppdynamicsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAppdynamicsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetAzureMonitorConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureMonitorConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetBigqueryConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetBigqueryConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetCloudwatchConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetDatadogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetDynatraceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatraceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetElasticsearchConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetGcmConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGcmConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetGrafanaLokiConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGrafanaLokiConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetGraphiteConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGraphiteConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetHistoricalDataRetrieval() {
	_jsii_.InvokeVoid(
		a,
		"resetHistoricalDataRetrieval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetInfluxdbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetInfluxdbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetInstanaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetLightstepConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLightstepConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetNewrelicConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNewrelicConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetOpentsdbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpentsdbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetPingdomConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPingdomConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetPrometheusConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPrometheusConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetQueryDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetRedshiftConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		a,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetSplunkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSplunkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetSplunkObservabilityConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSplunkObservabilityConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetSumologicConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSumologicConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) ResetThousandeyesConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetThousandeyesConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Agent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

