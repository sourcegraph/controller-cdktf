package googlegeminigeminigcpenablementsetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGeminiGeminiGcpEnablementSettingConfig struct {
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
	// Id of the Gemini Gcp Enablement setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#gemini_gcp_enablement_setting_id GoogleGeminiGeminiGcpEnablementSetting#gemini_gcp_enablement_setting_id}
	GeminiGcpEnablementSettingId *string `field:"required" json:"geminiGcpEnablementSettingId" yaml:"geminiGcpEnablementSettingId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#location GoogleGeminiGeminiGcpEnablementSetting#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Whether customer data sharing should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#enable_customer_data_sharing GoogleGeminiGeminiGcpEnablementSetting#enable_customer_data_sharing}
	EnableCustomerDataSharing interface{} `field:"optional" json:"enableCustomerDataSharing" yaml:"enableCustomerDataSharing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#id GoogleGeminiGeminiGcpEnablementSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#labels GoogleGeminiGeminiGcpEnablementSetting#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#project GoogleGeminiGeminiGcpEnablementSetting#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_gemini_gcp_enablement_setting#timeouts GoogleGeminiGeminiGcpEnablementSetting#timeouts}
	Timeouts *GoogleGeminiGeminiGcpEnablementSettingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

