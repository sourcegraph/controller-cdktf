package googledataformrepositoryworkflowconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataformRepositoryWorkflowConfigConfig struct {
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
	// The workflow's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#name GoogleDataformRepositoryWorkflowConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the release config whose releaseCompilationResult should be executed.
	//
	// Must be in the format projects/* /locations/* /repositories/* /releaseConfigs/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#release_config GoogleDataformRepositoryWorkflowConfig#release_config}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	ReleaseConfig *string `field:"required" json:"releaseConfig" yaml:"releaseConfig"`
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#cron_schedule GoogleDataformRepositoryWorkflowConfig#cron_schedule}
	CronSchedule *string `field:"optional" json:"cronSchedule" yaml:"cronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#id GoogleDataformRepositoryWorkflowConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// invocation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#invocation_config GoogleDataformRepositoryWorkflowConfig#invocation_config}
	InvocationConfig *GoogleDataformRepositoryWorkflowConfigInvocationConfig `field:"optional" json:"invocationConfig" yaml:"invocationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#project GoogleDataformRepositoryWorkflowConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#region GoogleDataformRepositoryWorkflowConfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A reference to the Dataform repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#repository GoogleDataformRepositoryWorkflowConfig#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#timeouts GoogleDataformRepositoryWorkflowConfig#timeouts}
	Timeouts *GoogleDataformRepositoryWorkflowConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_workflow_config#time_zone GoogleDataformRepositoryWorkflowConfig#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

