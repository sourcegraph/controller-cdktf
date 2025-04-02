package directcloudwatch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectCloudwatchConfig struct {
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
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#name DirectCloudwatch#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the Nobl9 project the resource sits in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#project DirectCloudwatch#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Optional description of the resource.
	//
	// Here, you can add details about who is responsible for the integration (team/owner) or the purpose of creating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#description DirectCloudwatch#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#display_name DirectCloudwatch#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// historical_data_retrieval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#historical_data_retrieval DirectCloudwatch#historical_data_retrieval}
	HistoricalDataRetrieval *DirectCloudwatchHistoricalDataRetrieval `field:"optional" json:"historicalDataRetrieval" yaml:"historicalDataRetrieval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#id DirectCloudwatch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// [Logs documentation](https://docs.nobl9.com/features/slo-troubleshooting/event-logs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#log_collection_enabled DirectCloudwatch#log_collection_enabled}
	LogCollectionEnabled interface{} `field:"optional" json:"logCollectionEnabled" yaml:"logCollectionEnabled"`
	// query_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#query_delay DirectCloudwatch#query_delay}
	QueryDelay *DirectCloudwatchQueryDelay `field:"optional" json:"queryDelay" yaml:"queryDelay"`
	// Release channel of the created data source [stable/beta].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#release_channel DirectCloudwatch#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// [required] | ARN of the AWS IAM Role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#role_arn DirectCloudwatch#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// This value indicated whether the field was a source of metrics and/or services.
	//
	// 'source_of' is deprecated and not used anywhere; however, it's kept for backward compatibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_cloudwatch#source_of DirectCloudwatch#source_of}
	SourceOf *[]*string `field:"optional" json:"sourceOf" yaml:"sourceOf"`
}

