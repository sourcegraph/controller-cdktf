package googlecloudfunctions2function

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudfunctions2FunctionConfig struct {
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
	// A user-defined name of the function. Function names must be unique globally and match pattern 'projects/*\/locations/*\/functions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#name GoogleCloudfunctions2Function#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// build_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#build_config GoogleCloudfunctions2Function#build_config}
	BuildConfig *GoogleCloudfunctions2FunctionBuildConfig `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// User-provided description of a function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#description GoogleCloudfunctions2Function#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// event_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#event_trigger GoogleCloudfunctions2Function#event_trigger}
	EventTrigger *GoogleCloudfunctions2FunctionEventTrigger `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#id GoogleCloudfunctions2Function#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs associated with this Cloud Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#labels GoogleCloudfunctions2Function#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of this cloud function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#location GoogleCloudfunctions2Function#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#project GoogleCloudfunctions2Function#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#service_config GoogleCloudfunctions2Function#service_config}
	ServiceConfig *GoogleCloudfunctions2FunctionServiceConfig `field:"optional" json:"serviceConfig" yaml:"serviceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions2_function#timeouts GoogleCloudfunctions2Function#timeouts}
	Timeouts *GoogleCloudfunctions2FunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

