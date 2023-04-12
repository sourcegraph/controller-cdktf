package workspaceruntask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceRunTaskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The enforcement level of the task. Valid values are `advisory` and `mandatory`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/workspace_run_task#enforcement_level WorkspaceRunTask#enforcement_level}
	EnforcementLevel *string `field:"required" json:"enforcementLevel" yaml:"enforcementLevel"`
	// The id of the Run task to associate to the Workspace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/workspace_run_task#task_id WorkspaceRunTask#task_id}
	TaskId *string `field:"required" json:"taskId" yaml:"taskId"`
	// The id of the workspace to associate the Run task to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/workspace_run_task#workspace_id WorkspaceRunTask#workspace_id}
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/workspace_run_task#id WorkspaceRunTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The stage to run the task in. Valid values are `pre_plan`, `post_plan` and `pre_apply`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/workspace_run_task#stage WorkspaceRunTask#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

