package googleapigeesharedflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeSharedflowConfig struct {
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
	// Path to the config zip bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_sharedflow#config_bundle GoogleApigeeSharedflow#config_bundle}
	ConfigBundle *string `field:"required" json:"configBundle" yaml:"configBundle"`
	// The ID of the shared flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_sharedflow#name GoogleApigeeSharedflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Apigee Organization name associated with the Apigee instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_sharedflow#org_id GoogleApigeeSharedflow#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// A hash of local config bundle in string, user needs to use a Terraform Hash function of their choice.
	//
	// A change in hash will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_sharedflow#detect_md5hash GoogleApigeeSharedflow#detect_md5hash}
	DetectMd5Hash *string `field:"optional" json:"detectMd5Hash" yaml:"detectMd5Hash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_sharedflow#id GoogleApigeeSharedflow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_sharedflow#timeouts GoogleApigeeSharedflow#timeouts}
	Timeouts *GoogleApigeeSharedflowTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

