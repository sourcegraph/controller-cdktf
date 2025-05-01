package googleeventarcpipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleEventarcPipelineConfig struct {
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
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#destinations GoogleEventarcPipeline#destinations}
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#location GoogleEventarcPipeline#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user-provided ID to be assigned to the Pipeline. It should match the format '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#pipeline_id GoogleEventarcPipeline#pipeline_id}
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
	// User-defined annotations. See https://google.aip.dev/128#annotations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#annotations GoogleEventarcPipeline#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt the event data.
	//
	// If not set, an internal Google-owned key
	// will be used to encrypt messages. It must match the pattern
	// "projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#crypto_key_name GoogleEventarcPipeline#crypto_key_name}
	CryptoKeyName *string `field:"optional" json:"cryptoKeyName" yaml:"cryptoKeyName"`
	// Display name of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#display_name GoogleEventarcPipeline#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#id GoogleEventarcPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input_payload_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#input_payload_format GoogleEventarcPipeline#input_payload_format}
	InputPayloadFormat *GoogleEventarcPipelineInputPayloadFormat `field:"optional" json:"inputPayloadFormat" yaml:"inputPayloadFormat"`
	// User labels attached to the Pipeline that can be used to group resources.
	//
	// An object containing a list of "key": value pairs. Example: {
	// "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#labels GoogleEventarcPipeline#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#logging_config GoogleEventarcPipeline#logging_config}
	LoggingConfig *GoogleEventarcPipelineLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// mediations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#mediations GoogleEventarcPipeline#mediations}
	Mediations interface{} `field:"optional" json:"mediations" yaml:"mediations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#project GoogleEventarcPipeline#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#retry_policy GoogleEventarcPipeline#retry_policy}
	RetryPolicy *GoogleEventarcPipelineRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#timeouts GoogleEventarcPipeline#timeouts}
	Timeouts *GoogleEventarcPipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

