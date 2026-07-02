package poller

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PollerConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Poller name. Must be unique within workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#name Poller#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#workspace Poller#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// aws_snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#aws_snapshot Poller#aws_snapshot}
	AwsSnapshot *PollerAwsSnapshot `field:"optional" json:"awsSnapshot" yaml:"awsSnapshot"`
	// chunk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#chunk Poller#chunk}
	Chunk *PollerChunk `field:"optional" json:"chunk" yaml:"chunk"`
	// cloudwatch_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#cloudwatch_metrics Poller#cloudwatch_metrics}
	CloudwatchMetrics *PollerCloudwatchMetrics `field:"optional" json:"cloudwatchMetrics" yaml:"cloudwatchMetrics"`
	// Datastream where poller will deliver data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#datastream Poller#datastream}
	Datastream *string `field:"optional" json:"datastream" yaml:"datastream"`
	// Whether to disable poller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#disabled Poller#disabled}
	Disabled any `field:"optional" json:"disabled" yaml:"disabled"`
	// gcp_monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#gcp_monitoring Poller#gcp_monitoring}
	GcpMonitoring *PollerGcpMonitoring `field:"optional" json:"gcpMonitoring" yaml:"gcpMonitoring"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#http Poller#http}
	Http *PollerHttp `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#id Poller#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Interval between poller runs. Only applicable to periodic poller kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#interval Poller#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// mongodbatlas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#mongodbatlas Poller#mongodbatlas}
	Mongodbatlas *PollerMongodbatlas `field:"optional" json:"mongodbatlas" yaml:"mongodbatlas"`
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#pubsub Poller#pubsub}
	Pubsub *PollerPubsub `field:"optional" json:"pubsub" yaml:"pubsub"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#retries Poller#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Skips validating any provided external API credentials against their external APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#skip_external_validation Poller#skip_external_validation}
	SkipExternalValidation any `field:"optional" json:"skipExternalValidation" yaml:"skipExternalValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#tags Poller#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}
