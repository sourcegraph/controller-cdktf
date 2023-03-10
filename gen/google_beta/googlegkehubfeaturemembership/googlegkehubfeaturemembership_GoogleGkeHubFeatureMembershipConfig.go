package googlegkehubfeaturemembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeHubFeatureMembershipConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// configmanagement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#configmanagement GoogleGkeHubFeatureMembership#configmanagement}
	Configmanagement *GoogleGkeHubFeatureMembershipConfigmanagement `field:"required" json:"configmanagement" yaml:"configmanagement"`
	// The name of the feature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#feature GoogleGkeHubFeatureMembership#feature}
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// The location of the feature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#location GoogleGkeHubFeatureMembership#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the membership.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#membership GoogleGkeHubFeatureMembership#membership}
	Membership *string `field:"required" json:"membership" yaml:"membership"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#id GoogleGkeHubFeatureMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project of the feature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#project GoogleGkeHubFeatureMembership#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#timeouts GoogleGkeHubFeatureMembership#timeouts}
	Timeouts *GoogleGkeHubFeatureMembershipTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

