package theme

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ThemeConfig struct {
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
	// Brand ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#brand_id Theme#brand_id}
	BrandId *string `field:"required" json:"brandId" yaml:"brandId"`
	// Path to local file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#background_image Theme#background_image}
	BackgroundImage *string `field:"optional" json:"backgroundImage" yaml:"backgroundImage"`
	// Variant for email templates (`OKTA_DEFAULT`, `FULL_THEME`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#email_template_touch_point_variant Theme#email_template_touch_point_variant}
	EmailTemplateTouchPointVariant *string `field:"optional" json:"emailTemplateTouchPointVariant" yaml:"emailTemplateTouchPointVariant"`
	// Variant for the Okta End-User Dashboard (`OKTA_DEFAULT`, `WHITE_LOGO_BACKGROUND`, `FULL_THEME`, `LOGO_ON_FULL_WHITE_BACKGROUND`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#end_user_dashboard_touch_point_variant Theme#end_user_dashboard_touch_point_variant}
	EndUserDashboardTouchPointVariant *string `field:"optional" json:"endUserDashboardTouchPointVariant" yaml:"endUserDashboardTouchPointVariant"`
	// Variant for the error page (`OKTA_DEFAULT`, `BACKGROUND_SECONDARY_COLOR`, `BACKGROUND_IMAGE`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#error_page_touch_point_variant Theme#error_page_touch_point_variant}
	ErrorPageTouchPointVariant *string `field:"optional" json:"errorPageTouchPointVariant" yaml:"errorPageTouchPointVariant"`
	// Path to local file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#favicon Theme#favicon}
	Favicon *string `field:"optional" json:"favicon" yaml:"favicon"`
	// Path to local file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#logo Theme#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Primary color contrast hex code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#primary_color_contrast_hex Theme#primary_color_contrast_hex}
	PrimaryColorContrastHex *string `field:"optional" json:"primaryColorContrastHex" yaml:"primaryColorContrastHex"`
	// Primary color hex code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#primary_color_hex Theme#primary_color_hex}
	PrimaryColorHex *string `field:"optional" json:"primaryColorHex" yaml:"primaryColorHex"`
	// Secondary color contrast hex code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#secondary_color_contrast_hex Theme#secondary_color_contrast_hex}
	SecondaryColorContrastHex *string `field:"optional" json:"secondaryColorContrastHex" yaml:"secondaryColorContrastHex"`
	// Secondary color hex code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#secondary_color_hex Theme#secondary_color_hex}
	SecondaryColorHex *string `field:"optional" json:"secondaryColorHex" yaml:"secondaryColorHex"`
	// Variant for the Okta Sign-In Page (`OKTA_DEFAULT`, `BACKGROUND_SECONDARY_COLOR`, `BACKGROUND_IMAGE`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#sign_in_page_touch_point_variant Theme#sign_in_page_touch_point_variant}
	SignInPageTouchPointVariant *string `field:"optional" json:"signInPageTouchPointVariant" yaml:"signInPageTouchPointVariant"`
	// Theme ID - Note: Okta API for theme only reads and updates therefore the okta_theme resource needs to act as a quasi data source.
	//
	// Do this by setting theme_id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/theme#theme_id Theme#theme_id}
	ThemeId *string `field:"optional" json:"themeId" yaml:"themeId"`
}

