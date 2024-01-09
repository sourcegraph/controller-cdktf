package agent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AgentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The type of the Agent. Check [Supported Agent types | Nobl9 Documentation](https://docs.nobl9.com/Sources/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#agent_type Agent#agent_type}
	AgentType *string `field:"required" json:"agentType" yaml:"agentType"`
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#name Agent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the Nobl9 project the resource sits in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#project Agent#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Source of Metrics and/or Services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#source_of Agent#source_of}
	SourceOf *[]*string `field:"required" json:"sourceOf" yaml:"sourceOf"`
	// amazon_prometheus_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#amazon_prometheus_config Agent#amazon_prometheus_config}
	AmazonPrometheusConfig *AgentAmazonPrometheusConfig `field:"optional" json:"amazonPrometheusConfig" yaml:"amazonPrometheusConfig"`
	// appdynamics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#appdynamics_config Agent#appdynamics_config}
	AppdynamicsConfig *AgentAppdynamicsConfig `field:"optional" json:"appdynamicsConfig" yaml:"appdynamicsConfig"`
	// azure_monitor_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#azure_monitor_config Agent#azure_monitor_config}
	AzureMonitorConfig *AgentAzureMonitorConfig `field:"optional" json:"azureMonitorConfig" yaml:"azureMonitorConfig"`
	// bigquery_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#bigquery_config Agent#bigquery_config}
	BigqueryConfig *AgentBigqueryConfig `field:"optional" json:"bigqueryConfig" yaml:"bigqueryConfig"`
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#cloudwatch_config Agent#cloudwatch_config}
	CloudwatchConfig *AgentCloudwatchConfig `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// datadog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#datadog_config Agent#datadog_config}
	DatadogConfig *AgentDatadogConfig `field:"optional" json:"datadogConfig" yaml:"datadogConfig"`
	// Optional description of the resource.
	//
	// Here, you can add details about who is responsible for the integration (team/owner) or the purpose of creating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#description Agent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#display_name Agent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// dynatrace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#dynatrace_config Agent#dynatrace_config}
	DynatraceConfig *AgentDynatraceConfig `field:"optional" json:"dynatraceConfig" yaml:"dynatraceConfig"`
	// elasticsearch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#elasticsearch_config Agent#elasticsearch_config}
	ElasticsearchConfig *AgentElasticsearchConfig `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// gcm_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#gcm_config Agent#gcm_config}
	GcmConfig *AgentGcmConfig `field:"optional" json:"gcmConfig" yaml:"gcmConfig"`
	// grafana_loki_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#grafana_loki_config Agent#grafana_loki_config}
	GrafanaLokiConfig *AgentGrafanaLokiConfig `field:"optional" json:"grafanaLokiConfig" yaml:"grafanaLokiConfig"`
	// graphite_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#graphite_config Agent#graphite_config}
	GraphiteConfig *AgentGraphiteConfig `field:"optional" json:"graphiteConfig" yaml:"graphiteConfig"`
	// historical_data_retrieval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#historical_data_retrieval Agent#historical_data_retrieval}
	HistoricalDataRetrieval *AgentHistoricalDataRetrieval `field:"optional" json:"historicalDataRetrieval" yaml:"historicalDataRetrieval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#id Agent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// influxdb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#influxdb_config Agent#influxdb_config}
	InfluxdbConfig *AgentInfluxdbConfig `field:"optional" json:"influxdbConfig" yaml:"influxdbConfig"`
	// instana_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#instana_config Agent#instana_config}
	InstanaConfig *AgentInstanaConfig `field:"optional" json:"instanaConfig" yaml:"instanaConfig"`
	// lightstep_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#lightstep_config Agent#lightstep_config}
	LightstepConfig *AgentLightstepConfig `field:"optional" json:"lightstepConfig" yaml:"lightstepConfig"`
	// newrelic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#newrelic_config Agent#newrelic_config}
	NewrelicConfig *AgentNewrelicConfig `field:"optional" json:"newrelicConfig" yaml:"newrelicConfig"`
	// opentsdb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#opentsdb_config Agent#opentsdb_config}
	OpentsdbConfig *AgentOpentsdbConfig `field:"optional" json:"opentsdbConfig" yaml:"opentsdbConfig"`
	// pingdom_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#pingdom_config Agent#pingdom_config}
	PingdomConfig *AgentPingdomConfig `field:"optional" json:"pingdomConfig" yaml:"pingdomConfig"`
	// prometheus_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#prometheus_config Agent#prometheus_config}
	PrometheusConfig *AgentPrometheusConfig `field:"optional" json:"prometheusConfig" yaml:"prometheusConfig"`
	// query_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#query_delay Agent#query_delay}
	QueryDelay *AgentQueryDelay `field:"optional" json:"queryDelay" yaml:"queryDelay"`
	// redshift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#redshift_config Agent#redshift_config}
	RedshiftConfig *AgentRedshiftConfig `field:"optional" json:"redshiftConfig" yaml:"redshiftConfig"`
	// Release channel of the created datasource [stable/beta].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#release_channel Agent#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// splunk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#splunk_config Agent#splunk_config}
	SplunkConfig *AgentSplunkConfig `field:"optional" json:"splunkConfig" yaml:"splunkConfig"`
	// splunk_observability_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#splunk_observability_config Agent#splunk_observability_config}
	SplunkObservabilityConfig *AgentSplunkObservabilityConfig `field:"optional" json:"splunkObservabilityConfig" yaml:"splunkObservabilityConfig"`
	// sumologic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#sumologic_config Agent#sumologic_config}
	SumologicConfig *AgentSumologicConfig `field:"optional" json:"sumologicConfig" yaml:"sumologicConfig"`
	// thousandeyes_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#thousandeyes_config Agent#thousandeyes_config}
	ThousandeyesConfig *AgentThousandeyesConfig `field:"optional" json:"thousandeyesConfig" yaml:"thousandeyesConfig"`
}

