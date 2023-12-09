package botmanagement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotManagementConfig struct {
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
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#zone_id BotManagement#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#auto_update_model BotManagement#auto_update_model}
	AutoUpdateModel interface{} `field:"optional" json:"autoUpdateModel" yaml:"autoUpdateModel"`
	// Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#enable_js BotManagement#enable_js}
	EnableJs interface{} `field:"optional" json:"enableJs" yaml:"enableJs"`
	// Whether to enable Bot Fight Mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#fight_mode BotManagement#fight_mode}
	FightMode interface{} `field:"optional" json:"fightMode" yaml:"fightMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#id BotManagement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to optimize Super Bot Fight Mode protections for Wordpress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#optimize_wordpress BotManagement#optimize_wordpress}
	OptimizeWordpress interface{} `field:"optional" json:"optimizeWordpress" yaml:"optimizeWordpress"`
	// Super Bot Fight Mode (SBFM) action to take on definitely automated requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#sbfm_definitely_automated BotManagement#sbfm_definitely_automated}
	SbfmDefinitelyAutomated *string `field:"optional" json:"sbfmDefinitelyAutomated" yaml:"sbfmDefinitelyAutomated"`
	// Super Bot Fight Mode (SBFM) action to take on likely automated requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#sbfm_likely_automated BotManagement#sbfm_likely_automated}
	SbfmLikelyAutomated *string `field:"optional" json:"sbfmLikelyAutomated" yaml:"sbfmLikelyAutomated"`
	// Super Bot Fight Mode (SBFM) to enable static resource protection.
	//
	// Enable if static resources on your application need bot protection. Note: Static resource protection can also result in legitimate traffic being blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#sbfm_static_resource_protection BotManagement#sbfm_static_resource_protection}
	SbfmStaticResourceProtection interface{} `field:"optional" json:"sbfmStaticResourceProtection" yaml:"sbfmStaticResourceProtection"`
	// Super Bot Fight Mode (SBFM) action to take on verified bots requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#sbfm_verified_bots BotManagement#sbfm_verified_bots}
	SbfmVerifiedBots *string `field:"optional" json:"sbfmVerifiedBots" yaml:"sbfmVerifiedBots"`
	// Whether to disable tracking the highest bot score for a session in the Bot Management cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/bot_management#suppress_session_score BotManagement#suppress_session_score}
	SuppressSessionScore interface{} `field:"optional" json:"suppressSessionScore" yaml:"suppressSessionScore"`
}

