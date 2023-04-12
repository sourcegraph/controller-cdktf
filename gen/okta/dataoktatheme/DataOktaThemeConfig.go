package dataoktatheme

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaThemeConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/theme#brand_id DataOktaTheme#brand_id}
	BrandId *string `field:"required" json:"brandId" yaml:"brandId"`
	// Theme ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/d/theme#theme_id DataOktaTheme#theme_id}
	ThemeId *string `field:"required" json:"themeId" yaml:"themeId"`
}

