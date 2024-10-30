package dataoktaemailcustomization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaEmailCustomizationConfig struct {
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
	// Brand ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/email_customization#brand_id DataOktaEmailCustomization#brand_id}
	BrandId *string `field:"required" json:"brandId" yaml:"brandId"`
	// The ID of the customization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/email_customization#customization_id DataOktaEmailCustomization#customization_id}
	CustomizationId *string `field:"required" json:"customizationId" yaml:"customizationId"`
	// Template Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/email_customization#template_name DataOktaEmailCustomization#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}

