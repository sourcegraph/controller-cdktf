package linkvalue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinkValueConfig struct {
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
	// Name of the 'primary' relationship being assigned.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_value#primary_name LinkValue#primary_name}
	PrimaryName *string `field:"required" json:"primaryName" yaml:"primaryName"`
	// User ID to be assigned to 'primary' for the 'associated' user in the specified relationship.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_value#primary_user_id LinkValue#primary_user_id}
	PrimaryUserId *string `field:"required" json:"primaryUserId" yaml:"primaryUserId"`
	// Set of User IDs or login values of the users to be assigned the 'associated' relationship.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_value#associated_user_ids LinkValue#associated_user_ids}
	AssociatedUserIds *[]*string `field:"optional" json:"associatedUserIds" yaml:"associatedUserIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_value#id LinkValue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

