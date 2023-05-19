package accessapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessApplicationConfig struct {
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
	// Friendly name of the Access Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#account_id AccessApplication#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The identity providers selected for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#allowed_idps AccessApplication#allowed_idps}
	AllowedIdps *[]*string `field:"optional" json:"allowedIdps" yaml:"allowedIdps"`
	// Option to show/hide applications in App Launcher. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#app_launcher_visible AccessApplication#app_launcher_visible}
	AppLauncherVisible interface{} `field:"optional" json:"appLauncherVisible" yaml:"appLauncherVisible"`
	// Option to skip identity provider selection if only one is configured in `allowed_idps`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#auto_redirect_to_identity AccessApplication#auto_redirect_to_identity}
	AutoRedirectToIdentity interface{} `field:"optional" json:"autoRedirectToIdentity" yaml:"autoRedirectToIdentity"`
	// cors_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#cors_headers AccessApplication#cors_headers}
	CorsHeaders interface{} `field:"optional" json:"corsHeaders" yaml:"corsHeaders"`
	// Option that returns a custom error message when a user is denied access to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#custom_deny_message AccessApplication#custom_deny_message}
	CustomDenyMessage *string `field:"optional" json:"customDenyMessage" yaml:"customDenyMessage"`
	// Option that redirects to a custom URL when a user is denied access to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#custom_deny_url AccessApplication#custom_deny_url}
	CustomDenyUrl *string `field:"optional" json:"customDenyUrl" yaml:"customDenyUrl"`
	// The complete URL of the asset you wish to put Cloudflare Access in front of.
	//
	// Can include subdomains or paths. Or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#domain AccessApplication#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional "binding" cookie on requests.
	//
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#enable_binding_cookie AccessApplication#enable_binding_cookie}
	EnableBindingCookie interface{} `field:"optional" json:"enableBindingCookie" yaml:"enableBindingCookie"`
	// Option to add the `HttpOnly` cookie flag to access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#http_only_cookie_attribute AccessApplication#http_only_cookie_attribute}
	HttpOnlyCookieAttribute interface{} `field:"optional" json:"httpOnlyCookieAttribute" yaml:"httpOnlyCookieAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#id AccessApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Image URL for the logo shown in the app launcher dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#logo_url AccessApplication#logo_url}
	LogoUrl *string `field:"optional" json:"logoUrl" yaml:"logoUrl"`
	// saas_app block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#saas_app AccessApplication#saas_app}
	SaasApp *AccessApplicationSaasApp `field:"optional" json:"saasApp" yaml:"saasApp"`
	// Defines the same-site cookie setting for access tokens. Available values: `none`, `lax`, `strict`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#same_site_cookie_attribute AccessApplication#same_site_cookie_attribute}
	SameSiteCookieAttribute *string `field:"optional" json:"sameSiteCookieAttribute" yaml:"sameSiteCookieAttribute"`
	// Option to return a 401 status code in service authentication rules on failed requests. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#service_auth_401_redirect AccessApplication#service_auth_401_redirect}
	ServiceAuth401Redirect interface{} `field:"optional" json:"serviceAuth401Redirect" yaml:"serviceAuth401Redirect"`
	// How often a user will be forced to re-authorise.
	//
	// Must be in the format `48h` or `2h45m`. Defaults to `24h`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#session_duration AccessApplication#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Option to skip the authorization interstitial when using the CLI. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#skip_interstitial AccessApplication#skip_interstitial}
	SkipInterstitial interface{} `field:"optional" json:"skipInterstitial" yaml:"skipInterstitial"`
	// The application type. Available values: `app_launcher`, `bookmark`, `biso`, `dash_sso`, `saas`, `self_hosted`, `ssh`, `vnc`, `warp`. Defaults to `self_hosted`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#type AccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application#zone_id AccessApplication#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

