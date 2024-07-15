package googledataplextask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataplexTaskConfig struct {
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
	// execution_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#execution_spec GoogleDataplexTask#execution_spec}
	ExecutionSpec *GoogleDataplexTaskExecutionSpec `field:"required" json:"executionSpec" yaml:"executionSpec"`
	// trigger_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#trigger_spec GoogleDataplexTask#trigger_spec}
	TriggerSpec *GoogleDataplexTaskTriggerSpec `field:"required" json:"triggerSpec" yaml:"triggerSpec"`
	// User-provided description of the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#description GoogleDataplexTask#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#display_name GoogleDataplexTask#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#id GoogleDataplexTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the task.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#labels GoogleDataplexTask#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The lake in which the task will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#lake GoogleDataplexTask#lake}
	Lake *string `field:"optional" json:"lake" yaml:"lake"`
	// The location in which the task will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#location GoogleDataplexTask#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// notebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#notebook GoogleDataplexTask#notebook}
	Notebook *GoogleDataplexTaskNotebook `field:"optional" json:"notebook" yaml:"notebook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#project GoogleDataplexTask#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#spark GoogleDataplexTask#spark}
	Spark *GoogleDataplexTaskSpark `field:"optional" json:"spark" yaml:"spark"`
	// The task Id of the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#task_id GoogleDataplexTask#task_id}
	TaskId *string `field:"optional" json:"taskId" yaml:"taskId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#timeouts GoogleDataplexTask#timeouts}
	Timeouts *GoogleDataplexTaskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

