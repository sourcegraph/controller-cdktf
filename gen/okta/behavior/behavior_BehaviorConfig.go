package behavior

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BehaviorConfig struct {
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
	// Name of the behavior.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#name Behavior#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Behavior type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#type Behavior#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#id Behavior#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines the method and level of detail used to evaluate the behavior.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#location_granularity_type Behavior#location_granularity_type}
	LocationGranularityType *string `field:"optional" json:"locationGranularityType" yaml:"locationGranularityType"`
	// The number of recent authentications used to evaluate the behavior.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#number_of_authentications Behavior#number_of_authentications}
	NumberOfAuthentications *float64 `field:"optional" json:"numberOfAuthentications" yaml:"numberOfAuthentications"`
	// Radius from location (in kilometers).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#radius_from_location Behavior#radius_from_location}
	RadiusFromLocation *float64 `field:"optional" json:"radiusFromLocation" yaml:"radiusFromLocation"`
	// Behavior status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#status Behavior#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Velocity (in kilometers per hour).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/behavior#velocity Behavior#velocity}
	Velocity *float64 `field:"optional" json:"velocity" yaml:"velocity"`
}

