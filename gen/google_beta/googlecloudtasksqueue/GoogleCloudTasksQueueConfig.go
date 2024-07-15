package googlecloudtasksqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudTasksQueueConfig struct {
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
	// The location of the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#location GoogleCloudTasksQueue#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// app_engine_routing_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#app_engine_routing_override GoogleCloudTasksQueue#app_engine_routing_override}
	AppEngineRoutingOverride *GoogleCloudTasksQueueAppEngineRoutingOverride `field:"optional" json:"appEngineRoutingOverride" yaml:"appEngineRoutingOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#id GoogleCloudTasksQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The queue name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#name GoogleCloudTasksQueue#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#project GoogleCloudTasksQueue#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rate_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#rate_limits GoogleCloudTasksQueue#rate_limits}
	RateLimits *GoogleCloudTasksQueueRateLimits `field:"optional" json:"rateLimits" yaml:"rateLimits"`
	// retry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#retry_config GoogleCloudTasksQueue#retry_config}
	RetryConfig *GoogleCloudTasksQueueRetryConfig `field:"optional" json:"retryConfig" yaml:"retryConfig"`
	// stackdriver_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#stackdriver_logging_config GoogleCloudTasksQueue#stackdriver_logging_config}
	StackdriverLoggingConfig *GoogleCloudTasksQueueStackdriverLoggingConfig `field:"optional" json:"stackdriverLoggingConfig" yaml:"stackdriverLoggingConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#timeouts GoogleCloudTasksQueue#timeouts}
	Timeouts *GoogleCloudTasksQueueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

