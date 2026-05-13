package googlecesexample

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCesExampleConfig struct {
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
	// Resource ID segment making up resource 'name', defining the app the example belongs to.
	//
	// It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#app GoogleCesExample#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// Display name of the example.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#display_name GoogleCesExample#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID to use for the example, which will become the final component of the example's resource name.
	//
	// In Terraform, this field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#example_id GoogleCesExample#example_id}
	ExampleId *string `field:"required" json:"exampleId" yaml:"exampleId"`
	// Resource ID segment making up resource 'name', defining what region the parent app is in.
	//
	// It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#location GoogleCesExample#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Human-readable description of the example.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#description GoogleCesExample#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The agent that initially handles the conversation.
	//
	// If not specified, the
	// example represents a conversation that is handled by the root agent.
	// Format: 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#entry_agent GoogleCesExample#entry_agent}
	EntryAgent *string `field:"optional" json:"entryAgent" yaml:"entryAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#id GoogleCesExample#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// messages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#messages GoogleCesExample#messages}
	Messages interface{} `field:"optional" json:"messages" yaml:"messages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#project GoogleCesExample#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_example#timeouts GoogleCesExample#timeouts}
	Timeouts *GoogleCesExampleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

