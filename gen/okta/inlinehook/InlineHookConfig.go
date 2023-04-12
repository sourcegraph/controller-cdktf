package inlinehook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InlineHookConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#channel InlineHook#channel}.
	Channel *map[string]*string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#name InlineHook#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#type InlineHook#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#version InlineHook#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#auth InlineHook#auth}.
	Auth *map[string]*string `field:"optional" json:"auth" yaml:"auth"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#headers InlineHook#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#id InlineHook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/inline_hook#status InlineHook#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

