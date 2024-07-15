package googledataformrepositoryreleaseconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataformRepositoryReleaseConfigConfig struct {
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
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#git_commitish GoogleDataformRepositoryReleaseConfig#git_commitish}
	GitCommitish *string `field:"required" json:"gitCommitish" yaml:"gitCommitish"`
	// The release's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#name GoogleDataformRepositoryReleaseConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// code_compilation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#code_compilation_config GoogleDataformRepositoryReleaseConfig#code_compilation_config}
	CodeCompilationConfig *GoogleDataformRepositoryReleaseConfigCodeCompilationConfig `field:"optional" json:"codeCompilationConfig" yaml:"codeCompilationConfig"`
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#cron_schedule GoogleDataformRepositoryReleaseConfig#cron_schedule}
	CronSchedule *string `field:"optional" json:"cronSchedule" yaml:"cronSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#id GoogleDataformRepositoryReleaseConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#project GoogleDataformRepositoryReleaseConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#region GoogleDataformRepositoryReleaseConfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A reference to the Dataform repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#repository GoogleDataformRepositoryReleaseConfig#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#timeouts GoogleDataformRepositoryReleaseConfig#timeouts}
	Timeouts *GoogleDataformRepositoryReleaseConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository_release_config#time_zone GoogleDataformRepositoryReleaseConfig#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

