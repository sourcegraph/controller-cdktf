package slo

type SloObjectiveCountMetricsTotal struct {
	// amazon_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#amazon_prometheus Slo#amazon_prometheus}
	AmazonPrometheus any `field:"optional" json:"amazonPrometheus" yaml:"amazonPrometheus"`
	// appdynamics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#appdynamics Slo#appdynamics}
	Appdynamics any `field:"optional" json:"appdynamics" yaml:"appdynamics"`
	// azure_monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#azure_monitor Slo#azure_monitor}
	AzureMonitor any `field:"optional" json:"azureMonitor" yaml:"azureMonitor"`
	// bigquery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#bigquery Slo#bigquery}
	Bigquery any `field:"optional" json:"bigquery" yaml:"bigquery"`
	// cloudwatch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#cloudwatch Slo#cloudwatch}
	Cloudwatch any `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#datadog Slo#datadog}
	Datadog any `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#dynatrace Slo#dynatrace}
	Dynatrace any `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#elasticsearch Slo#elasticsearch}
	Elasticsearch any `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// gcm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#gcm Slo#gcm}
	Gcm any `field:"optional" json:"gcm" yaml:"gcm"`
	// grafana_loki block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#grafana_loki Slo#grafana_loki}
	GrafanaLoki any `field:"optional" json:"grafanaLoki" yaml:"grafanaLoki"`
	// graphite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#graphite Slo#graphite}
	Graphite any `field:"optional" json:"graphite" yaml:"graphite"`
	// honeycomb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#honeycomb Slo#honeycomb}
	Honeycomb any `field:"optional" json:"honeycomb" yaml:"honeycomb"`
	// influxdb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#influxdb Slo#influxdb}
	Influxdb any `field:"optional" json:"influxdb" yaml:"influxdb"`
	// instana block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#instana Slo#instana}
	Instana any `field:"optional" json:"instana" yaml:"instana"`
	// lightstep block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#lightstep Slo#lightstep}
	Lightstep any `field:"optional" json:"lightstep" yaml:"lightstep"`
	// logic_monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#logic_monitor Slo#logic_monitor}
	LogicMonitor any `field:"optional" json:"logicMonitor" yaml:"logicMonitor"`
	// newrelic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#newrelic Slo#newrelic}
	Newrelic any `field:"optional" json:"newrelic" yaml:"newrelic"`
	// opentsdb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#opentsdb Slo#opentsdb}
	Opentsdb any `field:"optional" json:"opentsdb" yaml:"opentsdb"`
	// pingdom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#pingdom Slo#pingdom}
	Pingdom any `field:"optional" json:"pingdom" yaml:"pingdom"`
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#prometheus Slo#prometheus}
	Prometheus any `field:"optional" json:"prometheus" yaml:"prometheus"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#redshift Slo#redshift}
	Redshift any `field:"optional" json:"redshift" yaml:"redshift"`
	// splunk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#splunk Slo#splunk}
	Splunk any `field:"optional" json:"splunk" yaml:"splunk"`
	// splunk_observability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#splunk_observability Slo#splunk_observability}
	SplunkObservability any `field:"optional" json:"splunkObservability" yaml:"splunkObservability"`
	// sumologic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#sumologic Slo#sumologic}
	Sumologic any `field:"optional" json:"sumologic" yaml:"sumologic"`
	// thousandeyes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#thousandeyes Slo#thousandeyes}
	Thousandeyes any `field:"optional" json:"thousandeyes" yaml:"thousandeyes"`
}
