package templatesms

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TemplateSmsConfig struct {
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
	// SMS default template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_sms#template TemplateSms#template}
	Template *string `field:"required" json:"template" yaml:"template"`
	// SMS template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_sms#type TemplateSms#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_sms#id TemplateSms#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// translations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/template_sms#translations TemplateSms#translations}
	Translations interface{} `field:"optional" json:"translations" yaml:"translations"`
}

