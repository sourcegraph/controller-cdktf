package googleservicedirectoryservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleServiceDirectoryServiceConfig struct {
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
	// The resource name of the namespace this service will belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service#namespace GoogleServiceDirectoryService#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service#service_id GoogleServiceDirectoryService#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service#id GoogleServiceDirectoryService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metadata for the service.
	//
	// This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 2000 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service#metadata GoogleServiceDirectoryService#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_service#timeouts GoogleServiceDirectoryService#timeouts}
	Timeouts *GoogleServiceDirectoryServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

