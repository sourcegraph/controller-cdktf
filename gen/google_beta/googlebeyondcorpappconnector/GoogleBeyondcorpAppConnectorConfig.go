package googlebeyondcorpappconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBeyondcorpAppConnectorConfig struct {
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
	// ID of the AppConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#name GoogleBeyondcorpAppConnector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// principal_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#principal_info GoogleBeyondcorpAppConnector#principal_info}
	PrincipalInfo *GoogleBeyondcorpAppConnectorPrincipalInfo `field:"required" json:"principalInfo" yaml:"principalInfo"`
	// An arbitrary user-provided name for the AppConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#display_name GoogleBeyondcorpAppConnector#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#id GoogleBeyondcorpAppConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#labels GoogleBeyondcorpAppConnector#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#project GoogleBeyondcorpAppConnector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the AppConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#region GoogleBeyondcorpAppConnector#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#timeouts GoogleBeyondcorpAppConnector#timeouts}
	Timeouts *GoogleBeyondcorpAppConnectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

