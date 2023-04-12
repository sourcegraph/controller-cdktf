package accessapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#domain AccessApplication#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#name AccessApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#account_id AccessApplication#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#allowed_idps AccessApplication#allowed_idps}.
	AllowedIdps *[]*string `field:"optional" json:"allowedIdps" yaml:"allowedIdps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#app_launcher_visible AccessApplication#app_launcher_visible}.
	AppLauncherVisible interface{} `field:"optional" json:"appLauncherVisible" yaml:"appLauncherVisible"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#auto_redirect_to_identity AccessApplication#auto_redirect_to_identity}.
	AutoRedirectToIdentity interface{} `field:"optional" json:"autoRedirectToIdentity" yaml:"autoRedirectToIdentity"`
	// cors_headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#cors_headers AccessApplication#cors_headers}
	CorsHeaders interface{} `field:"optional" json:"corsHeaders" yaml:"corsHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#custom_deny_message AccessApplication#custom_deny_message}.
	CustomDenyMessage *string `field:"optional" json:"customDenyMessage" yaml:"customDenyMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#custom_deny_url AccessApplication#custom_deny_url}.
	CustomDenyUrl *string `field:"optional" json:"customDenyUrl" yaml:"customDenyUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#enable_binding_cookie AccessApplication#enable_binding_cookie}.
	EnableBindingCookie interface{} `field:"optional" json:"enableBindingCookie" yaml:"enableBindingCookie"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#http_only_cookie_attribute AccessApplication#http_only_cookie_attribute}.
	HttpOnlyCookieAttribute interface{} `field:"optional" json:"httpOnlyCookieAttribute" yaml:"httpOnlyCookieAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#id AccessApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#logo_url AccessApplication#logo_url}.
	LogoUrl *string `field:"optional" json:"logoUrl" yaml:"logoUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#same_site_cookie_attribute AccessApplication#same_site_cookie_attribute}.
	SameSiteCookieAttribute *string `field:"optional" json:"sameSiteCookieAttribute" yaml:"sameSiteCookieAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#service_auth_401_redirect AccessApplication#service_auth_401_redirect}.
	ServiceAuth401Redirect interface{} `field:"optional" json:"serviceAuth401Redirect" yaml:"serviceAuth401Redirect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#session_duration AccessApplication#session_duration}.
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#skip_interstitial AccessApplication#skip_interstitial}.
	SkipInterstitial interface{} `field:"optional" json:"skipInterstitial" yaml:"skipInterstitial"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#type AccessApplication#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_application#zone_id AccessApplication#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

