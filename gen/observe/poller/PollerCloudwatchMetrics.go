package poller

type PollerCloudwatchMetrics struct {
	// AWS role to assume when scraping AWS CloudWatch Metrics. External ID will be set to datastream ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#assume_role_arn Poller#assume_role_arn}
	AssumeRoleArn *string `field:"required" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#query Poller#query}
	Query any `field:"required" json:"query" yaml:"query"`
	// AWS Region to scrape from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#region Poller#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Collection delay. Must account for metrics availability via CloudWatch API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#delay Poller#delay}
	Delay *string `field:"optional" json:"delay" yaml:"delay"`
	// Metric resolution. Must be a multiple of 60s. When omitted, poller interval will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#period Poller#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
}
