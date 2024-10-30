package adminrolecustom

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdminRoleCustomConfig struct {
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
	// A human-readable description of the new Role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_custom#description AdminRoleCustom#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name given to the new Role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_custom#label AdminRoleCustom#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_custom#id AdminRoleCustom#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The permissions that the new Role grants.
	//
	// At least one
	// 				permission must be specified when creating custom role. Valid values: "okta.authzServers.manage",
	// 			  "okta.authzServers.read",
	// 			  "okta.apps.assignment.manage",
	// 			  "okta.apps.manage",
	// 			  "okta.apps.read",
	// 			  "okta.customizations.manage",
	// 			  "okta.customizations.read",
	// 			  "okta.groups.appAssignment.manage",
	// 			  "okta.groups.create",
	// 			  "okta.groups.manage",
	// 			  "okta.groups.members.manage",
	// 			  "okta.groups.read",
	// 			  "okta.profilesources.import.run",
	// 			  "okta.users.appAssignment.manage",
	// 			  "okta.users.create",
	// 			  "okta.users.credentials.expirePassword",
	// 			  "okta.users.credentials.manage",
	// 			  "okta.users.credentials.resetFactors",
	// 			  "okta.users.credentials.resetPassword",
	// 			  "okta.users.groupMembership.manage",
	// 			  "okta.users.lifecycle.activate",
	// 			  "okta.users.lifecycle.clearSessions",
	// 			  "okta.users.lifecycle.deactivate",
	// 			  "okta.users.lifecycle.delete",
	// 			  "okta.users.lifecycle.manage",
	// 			  "okta.users.lifecycle.suspend",
	// 			  "okta.users.lifecycle.unlock",
	// 			  "okta.users.lifecycle.unsuspend",
	// 			  "okta.users.manage",
	// 			  "okta.users.read",
	// 			  "okta.users.userprofile.manage",
	// 			  "okta.workflows.invoke".,
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_custom#permissions AdminRoleCustom#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

