package googlesecuritypostureposture

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecurityposturePostureConfig struct {
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
	// Location of the resource, eg: global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#location GoogleSecurityposturePosture#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#parent GoogleSecurityposturePosture#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// policy_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#policy_sets GoogleSecurityposturePosture#policy_sets}
	PolicySets interface{} `field:"required" json:"policySets" yaml:"policySets"`
	// Id of the posture. It is an immutable field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#posture_id GoogleSecurityposturePosture#posture_id}
	PostureId *string `field:"required" json:"postureId" yaml:"postureId"`
	// State of the posture.
	//
	// Update to state field should not be triggered along with
	// with other field updates. Possible values: ["DEPRECATED", "DRAFT", "ACTIVE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#state GoogleSecurityposturePosture#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// Description of the posture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#description GoogleSecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#id GoogleSecurityposturePosture#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#timeouts GoogleSecurityposturePosture#timeouts}
	Timeouts *GoogleSecurityposturePostureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

