package googlestoragemanagedfolder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageManagedFolderConfig struct {
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
	// The name of the bucket that contains the managed folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#bucket GoogleStorageManagedFolder#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the managed folder expressed as a path. Must include trailing '/'. For example, 'example_dir/example_dir2/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#name GoogleStorageManagedFolder#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allows the deletion of a managed folder even if contains objects.
	//
	// If a non-empty managed folder is deleted, any objects
	// within the folder will remain in a simulated folder with the
	// same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#force_destroy GoogleStorageManagedFolder#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#id GoogleStorageManagedFolder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_managed_folder#timeouts GoogleStorageManagedFolder#timeouts}
	Timeouts *GoogleStorageManagedFolderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

