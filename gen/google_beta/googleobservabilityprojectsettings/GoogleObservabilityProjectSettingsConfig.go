package googleobservabilityprojectsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleObservabilityProjectSettingsConfig struct {
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
	// The location of the settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_observability_project_settings#location GoogleObservabilityProjectSettings#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The default storage location for new resources, e.g. buckets. Only valid for global location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_observability_project_settings#default_storage_location GoogleObservabilityProjectSettings#default_storage_location}
	DefaultStorageLocation *string `field:"optional" json:"defaultStorageLocation" yaml:"defaultStorageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_observability_project_settings#id GoogleObservabilityProjectSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The default Cloud KMS key to use for new resources. Only valid for regional locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_observability_project_settings#kms_key_name GoogleObservabilityProjectSettings#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_observability_project_settings#project GoogleObservabilityProjectSettings#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_observability_project_settings#timeouts GoogleObservabilityProjectSettings#timeouts}
	Timeouts *GoogleObservabilityProjectSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

