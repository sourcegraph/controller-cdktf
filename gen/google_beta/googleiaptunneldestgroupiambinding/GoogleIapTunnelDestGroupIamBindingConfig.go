package googleiaptunneldestgroupiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIapTunnelDestGroupIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#dest_group GoogleIapTunnelDestGroupIamBinding#dest_group}.
	DestGroup *string `field:"required" json:"destGroup" yaml:"destGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#members GoogleIapTunnelDestGroupIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#role GoogleIapTunnelDestGroupIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#condition GoogleIapTunnelDestGroupIamBinding#condition}
	Condition *GoogleIapTunnelDestGroupIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#id GoogleIapTunnelDestGroupIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#project GoogleIapTunnelDestGroupIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iap_tunnel_dest_group_iam_binding#region GoogleIapTunnelDestGroupIamBinding#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

