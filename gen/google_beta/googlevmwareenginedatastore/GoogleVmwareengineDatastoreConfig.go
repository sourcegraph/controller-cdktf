package googlevmwareenginedatastore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVmwareengineDatastoreConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#location GoogleVmwareengineDatastore#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user-provided identifier of the datastore to be created.
	//
	// This identifier must be unique among each 'Datastore' within the parent
	// and becomes the final token in the name URI.
	// The identifier must meet the following requirements:
	//
	// * Only contains 1-63 alphanumeric characters and hyphens
	// * Begins with an alphabetical character
	// * Ends with a non-hyphen character
	// * Not formatted as a UUID
	// * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034)
	// (section 3.5)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#name GoogleVmwareengineDatastore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// nfs_datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#nfs_datastore GoogleVmwareengineDatastore#nfs_datastore}
	NfsDatastore *GoogleVmwareengineDatastoreNfsDatastore `field:"required" json:"nfsDatastore" yaml:"nfsDatastore"`
	// User-provided description for this datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#description GoogleVmwareengineDatastore#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#id GoogleVmwareengineDatastore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#project GoogleVmwareengineDatastore#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#timeouts GoogleVmwareengineDatastore#timeouts}
	Timeouts *GoogleVmwareengineDatastoreTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

