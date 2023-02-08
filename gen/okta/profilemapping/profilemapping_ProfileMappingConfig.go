package profilemapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProfileMappingConfig struct {
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
	// The source id of the mapping to manage.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#source_id ProfileMapping#source_id}
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
	// The target id of the mapping to manage.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#target_id ProfileMapping#target_id}
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Whether apply the changes to all users with this profile after updating or creating the these mappings.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#always_apply ProfileMapping#always_apply}
	AlwaysApply interface{} `field:"optional" json:"alwaysApply" yaml:"alwaysApply"`
	// When turned on this flag will trigger the provider to delete mapping properties that are not defined in config.
	//
	// By default, we do not delete missing properties.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#delete_when_absent ProfileMapping#delete_when_absent}
	DeleteWhenAbsent interface{} `field:"optional" json:"deleteWhenAbsent" yaml:"deleteWhenAbsent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#id ProfileMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mappings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#mappings ProfileMapping#mappings}
	Mappings interface{} `field:"optional" json:"mappings" yaml:"mappings"`
}

