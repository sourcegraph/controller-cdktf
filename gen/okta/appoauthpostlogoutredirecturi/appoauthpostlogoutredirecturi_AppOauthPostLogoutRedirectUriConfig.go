package appoauthpostlogoutredirecturi

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppOauthPostLogoutRedirectUriConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth_post_logout_redirect_uri#app_id AppOauthPostLogoutRedirectUri#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Post Logout Redirect URI to append to Okta OIDC application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth_post_logout_redirect_uri#uri AppOauthPostLogoutRedirectUri#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth_post_logout_redirect_uri#id AppOauthPostLogoutRedirectUri#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

