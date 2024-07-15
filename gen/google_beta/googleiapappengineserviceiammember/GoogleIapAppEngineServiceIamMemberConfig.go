package googleiapappengineserviceiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIapAppEngineServiceIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#app_id GoogleIapAppEngineServiceIamMember#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#member GoogleIapAppEngineServiceIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#role GoogleIapAppEngineServiceIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#service GoogleIapAppEngineServiceIamMember#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#condition GoogleIapAppEngineServiceIamMember#condition}
	Condition *GoogleIapAppEngineServiceIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#id GoogleIapAppEngineServiceIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_app_engine_service_iam_member#project GoogleIapAppEngineServiceIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

