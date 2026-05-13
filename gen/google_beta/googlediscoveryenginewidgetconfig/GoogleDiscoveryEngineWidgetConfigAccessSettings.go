package googlediscoveryenginewidgetconfig


type GoogleDiscoveryEngineWidgetConfigAccessSettings struct {
	// List of domains that are allowed to integrate the search widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#allowlisted_domains GoogleDiscoveryEngineWidgetConfig#allowlisted_domains}
	AllowlistedDomains *[]*string `field:"optional" json:"allowlistedDomains" yaml:"allowlistedDomains"`
	// Whether public unauthenticated access is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#allow_public_access GoogleDiscoveryEngineWidgetConfig#allow_public_access}
	AllowPublicAccess interface{} `field:"optional" json:"allowPublicAccess" yaml:"allowPublicAccess"`
	// Whether web app access is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#enable_web_app GoogleDiscoveryEngineWidgetConfig#enable_web_app}
	EnableWebApp interface{} `field:"optional" json:"enableWebApp" yaml:"enableWebApp"`
	// Language code for user interface. Use language tags defined by [BCP47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). If unset, the default language code is "en-US".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#language_code GoogleDiscoveryEngineWidgetConfig#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The workforce identity pool provider used to access the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#workforce_identity_pool_provider GoogleDiscoveryEngineWidgetConfig#workforce_identity_pool_provider}
	WorkforceIdentityPoolProvider *string `field:"optional" json:"workforceIdentityPoolProvider" yaml:"workforceIdentityPoolProvider"`
}

