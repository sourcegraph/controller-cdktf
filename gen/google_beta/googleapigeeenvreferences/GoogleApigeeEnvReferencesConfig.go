package googleapigeeenvreferences

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeEnvReferencesConfig struct {
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
	// The Apigee environment group associated with the Apigee environment, in the format 'organizations/{{org_name}}/environments/{{env_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#env_id GoogleApigeeEnvReferences#env_id}
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// Required. The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#name GoogleApigeeEnvReferences#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required.
	//
	// The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#refers GoogleApigeeEnvReferences#refers}
	Refers *string `field:"required" json:"refers" yaml:"refers"`
	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#resource_type GoogleApigeeEnvReferences#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Optional. A human-readable description of this reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#description GoogleApigeeEnvReferences#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#id GoogleApigeeEnvReferences#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_env_references#timeouts GoogleApigeeEnvReferences#timeouts}
	Timeouts *GoogleApigeeEnvReferencesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

