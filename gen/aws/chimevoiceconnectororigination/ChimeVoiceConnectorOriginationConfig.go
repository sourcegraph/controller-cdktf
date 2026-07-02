package chimevoiceconnectororigination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimeVoiceConnectorOriginationConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/chime_voice_connector_origination#route ChimeVoiceConnectorOrigination#route}
	Route any `field:"required" json:"route" yaml:"route"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/chime_voice_connector_origination#voice_connector_id ChimeVoiceConnectorOrigination#voice_connector_id}.
	VoiceConnectorId *string `field:"required" json:"voiceConnectorId" yaml:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/chime_voice_connector_origination#disabled ChimeVoiceConnectorOrigination#disabled}.
	Disabled any `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/chime_voice_connector_origination#id ChimeVoiceConnectorOrigination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
