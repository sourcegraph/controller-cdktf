package authserverscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthServerScopeConfig struct {
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
	// Auth server ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#auth_server_id AuthServerScope#auth_server_id}
	AuthServerId *string `field:"required" json:"authServerId" yaml:"authServerId"`
	// Auth server scope name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#name AuthServerScope#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// EA Feature and thus it is simply ignored if the feature is off.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#consent AuthServerScope#consent}
	Consent *string `field:"optional" json:"consent" yaml:"consent"`
	// A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#default AuthServerScope#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#description AuthServerScope#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the end user displayed in a consent dialog box.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#display_name AuthServerScope#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#id AuthServerScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to publish metadata or not, matching API type despite the fact it could just be a boolean.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/auth_server_scope#metadata_publish AuthServerScope#metadata_publish}
	MetadataPublish *string `field:"optional" json:"metadataPublish" yaml:"metadataPublish"`
}

