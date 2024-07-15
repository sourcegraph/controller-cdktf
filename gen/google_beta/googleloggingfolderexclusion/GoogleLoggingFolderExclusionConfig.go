package googleloggingfolderexclusion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleLoggingFolderExclusionConfig struct {
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
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_folder_exclusion#filter GoogleLoggingFolderExclusion#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_folder_exclusion#folder GoogleLoggingFolderExclusion#folder}.
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The name of the logging exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_folder_exclusion#name GoogleLoggingFolderExclusion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_folder_exclusion#description GoogleLoggingFolderExclusion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_folder_exclusion#disabled GoogleLoggingFolderExclusion#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_folder_exclusion#id GoogleLoggingFolderExclusion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

