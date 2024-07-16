package googlebinaryauthorizationattestoriambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBinaryAuthorizationAttestorIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_binding#attestor GoogleBinaryAuthorizationAttestorIamBinding#attestor}.
	Attestor *string `field:"required" json:"attestor" yaml:"attestor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_binding#members GoogleBinaryAuthorizationAttestorIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_binding#role GoogleBinaryAuthorizationAttestorIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_binding#condition GoogleBinaryAuthorizationAttestorIamBinding#condition}
	Condition *GoogleBinaryAuthorizationAttestorIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_binding#id GoogleBinaryAuthorizationAttestorIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_attestor_iam_binding#project GoogleBinaryAuthorizationAttestorIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

