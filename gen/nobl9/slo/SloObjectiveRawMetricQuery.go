package slo


type SloObjectiveRawMetricQuery struct {
	// amazon_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#amazon_prometheus Slo#amazon_prometheus}
	AmazonPrometheus interface{} `field:"optional" json:"amazonPrometheus" yaml:"amazonPrometheus"`
	// appdynamics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#appdynamics Slo#appdynamics}
	Appdynamics interface{} `field:"optional" json:"appdynamics" yaml:"appdynamics"`
	// azure_monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#azure_monitor Slo#azure_monitor}
	AzureMonitor interface{} `field:"optional" json:"azureMonitor" yaml:"azureMonitor"`
	// bigquery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#bigquery Slo#bigquery}
	Bigquery interface{} `field:"optional" json:"bigquery" yaml:"bigquery"`
	// cloudwatch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#cloudwatch Slo#cloudwatch}
	Cloudwatch interface{} `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#datadog Slo#datadog}
	Datadog interface{} `field:"optional" json:"datadog" yaml:"datadog"`
	// dynatrace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#dynatrace Slo#dynatrace}
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#elasticsearch Slo#elasticsearch}
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// gcm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#gcm Slo#gcm}
	Gcm interface{} `field:"optional" json:"gcm" yaml:"gcm"`
	// grafana_loki block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#grafana_loki Slo#grafana_loki}
	GrafanaLoki interface{} `field:"optional" json:"grafanaLoki" yaml:"grafanaLoki"`
	// graphite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#graphite Slo#graphite}
	Graphite interface{} `field:"optional" json:"graphite" yaml:"graphite"`
	// influxdb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#influxdb Slo#influxdb}
	Influxdb interface{} `field:"optional" json:"influxdb" yaml:"influxdb"`
	// instana block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#instana Slo#instana}
	Instana interface{} `field:"optional" json:"instana" yaml:"instana"`
	// lightstep block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#lightstep Slo#lightstep}
	Lightstep interface{} `field:"optional" json:"lightstep" yaml:"lightstep"`
	// newrelic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#newrelic Slo#newrelic}
	Newrelic interface{} `field:"optional" json:"newrelic" yaml:"newrelic"`
	// opentsdb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#opentsdb Slo#opentsdb}
	Opentsdb interface{} `field:"optional" json:"opentsdb" yaml:"opentsdb"`
	// pingdom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#pingdom Slo#pingdom}
	Pingdom interface{} `field:"optional" json:"pingdom" yaml:"pingdom"`
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#prometheus Slo#prometheus}
	Prometheus interface{} `field:"optional" json:"prometheus" yaml:"prometheus"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#redshift Slo#redshift}
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// splunk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#splunk Slo#splunk}
	Splunk interface{} `field:"optional" json:"splunk" yaml:"splunk"`
	// splunk_observability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#splunk_observability Slo#splunk_observability}
	SplunkObservability interface{} `field:"optional" json:"splunkObservability" yaml:"splunkObservability"`
	// sumologic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#sumologic Slo#sumologic}
	Sumologic interface{} `field:"optional" json:"sumologic" yaml:"sumologic"`
	// thousandeyes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#thousandeyes Slo#thousandeyes}
	Thousandeyes interface{} `field:"optional" json:"thousandeyes" yaml:"thousandeyes"`
}

