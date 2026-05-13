package googlefirebaseailogicconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAiLogicConfigConfig struct {
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
	// generative_language_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#generative_language_config GoogleFirebaseAiLogicConfig#generative_language_config}
	GenerativeLanguageConfig *GoogleFirebaseAiLogicConfigGenerativeLanguageConfig `field:"optional" json:"generativeLanguageConfig" yaml:"generativeLanguageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#id GoogleFirebaseAiLogicConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#location GoogleFirebaseAiLogicConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#project GoogleFirebaseAiLogicConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// telemetry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#telemetry_config GoogleFirebaseAiLogicConfig#telemetry_config}
	TelemetryConfig *GoogleFirebaseAiLogicConfigTelemetryConfig `field:"optional" json:"telemetryConfig" yaml:"telemetryConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#timeouts GoogleFirebaseAiLogicConfig#timeouts}
	Timeouts *GoogleFirebaseAiLogicConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#traffic_filter GoogleFirebaseAiLogicConfig#traffic_filter}
	TrafficFilter *GoogleFirebaseAiLogicConfigTrafficFilter `field:"optional" json:"trafficFilter" yaml:"trafficFilter"`
}

