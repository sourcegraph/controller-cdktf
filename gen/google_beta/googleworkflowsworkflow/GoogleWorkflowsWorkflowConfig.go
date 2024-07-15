package googleworkflowsworkflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkflowsWorkflowConfig struct {
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
	// Describes the level of platform logging to apply to calls and call responses during executions of this workflow.
	//
	// If both the workflow and the execution specify a logging level,
	// the execution level takes precedence. Possible values: ["CALL_LOG_LEVEL_UNSPECIFIED", "LOG_ALL_CALLS", "LOG_ERRORS_ONLY", "LOG_NONE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#call_log_level GoogleWorkflowsWorkflow#call_log_level}
	CallLogLevel *string `field:"optional" json:"callLogLevel" yaml:"callLogLevel"`
	// The KMS key used to encrypt workflow and execution data.
	//
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#crypto_key_name GoogleWorkflowsWorkflow#crypto_key_name}
	CryptoKeyName *string `field:"optional" json:"cryptoKeyName" yaml:"cryptoKeyName"`
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#description GoogleWorkflowsWorkflow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#id GoogleWorkflowsWorkflow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to this Workflow.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#labels GoogleWorkflowsWorkflow#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the Workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#name GoogleWorkflowsWorkflow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#name_prefix GoogleWorkflowsWorkflow#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#project GoogleWorkflowsWorkflow#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#region GoogleWorkflowsWorkflow#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Name of the service account associated with the latest workflow version.
	//
	// This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the unique_id of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#service_account GoogleWorkflowsWorkflow#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Workflow code to be executed. The size limit is 128KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#source_contents GoogleWorkflowsWorkflow#source_contents}
	SourceContents *string `field:"optional" json:"sourceContents" yaml:"sourceContents"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#timeouts GoogleWorkflowsWorkflow#timeouts}
	Timeouts *GoogleWorkflowsWorkflowTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// User-defined environment variables associated with this workflow revision.
	//
	// This map has a maximum length of 20. Each string can take up to 4KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or “WORKFLOWS".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workflows_workflow#user_env_vars GoogleWorkflowsWorkflow#user_env_vars}
	UserEnvVars *map[string]*string `field:"optional" json:"userEnvVars" yaml:"userEnvVars"`
}

