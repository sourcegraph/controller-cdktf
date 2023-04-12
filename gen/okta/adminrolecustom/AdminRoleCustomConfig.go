package adminrolecustom

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdminRoleCustomConfig struct {
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
	// A human-readable description of the new Role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom#description AdminRoleCustom#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name given to the new Role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom#label AdminRoleCustom#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom#id AdminRoleCustom#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The permissions that the new Role grants.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/admin_role_custom#permissions AdminRoleCustom#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

