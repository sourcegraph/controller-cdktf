package turnstilewidget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TurnstileWidgetConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#account_id TurnstileWidget#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Domains where the widget is deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#domains TurnstileWidget#domains}
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// Widget Mode. Available values: `non-interactive`, `invisible`, `managed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#mode TurnstileWidget#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Human readable widget name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#name TurnstileWidget#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If bot_fight_mode is set to true, Cloudflare issues computationally expensive challenges in response to malicious bots (Enterprise only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#bot_fight_mode TurnstileWidget#bot_fight_mode}
	BotFightMode interface{} `field:"optional" json:"botFightMode" yaml:"botFightMode"`
	// The identifier of this resource. This is the site key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#id TurnstileWidget#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Do not show any Cloudflare branding on the widget (Enterprise only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#offlabel TurnstileWidget#offlabel}
	Offlabel interface{} `field:"optional" json:"offlabel" yaml:"offlabel"`
	// Region where this widget can be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/turnstile_widget#region TurnstileWidget#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

