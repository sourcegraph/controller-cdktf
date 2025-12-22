package poller


type PollerAwsSnapshot struct {
	// AWS role to assume when scraping AWS API. External ID will be set to datastream ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#assume_role_arn Poller#assume_role_arn}
	AssumeRoleArn *string `field:"required" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Set of AWS API actions poller is allowed to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#include_actions Poller#include_actions}
	IncludeActions *[]*string `field:"required" json:"includeActions" yaml:"includeActions"`
	// AWS Region to scrape from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#region Poller#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

