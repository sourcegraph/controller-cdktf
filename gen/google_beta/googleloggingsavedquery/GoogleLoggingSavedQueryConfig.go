package googleloggingsavedquery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleLoggingSavedQueryConfig struct {
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
	// The user-visible display name of the saved query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#display_name GoogleLoggingSavedQuery#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location of the resource see [supported regions](https://docs.cloud.google.com/logging/docs/region-support#bucket-regions).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#location GoogleLoggingSavedQuery#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the saved query. For example: 'my-saved-query'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#name GoogleLoggingSavedQuery#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#parent GoogleLoggingSavedQuery#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The visibility of the saved query. Possible values: ["SHARED", "PRIVATE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#visibility GoogleLoggingSavedQuery#visibility}
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// A description of the saved query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#description GoogleLoggingSavedQuery#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#id GoogleLoggingSavedQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#logging_query GoogleLoggingSavedQuery#logging_query}
	LoggingQuery *GoogleLoggingSavedQueryLoggingQuery `field:"optional" json:"loggingQuery" yaml:"loggingQuery"`
	// ops_analytics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#ops_analytics_query GoogleLoggingSavedQuery#ops_analytics_query}
	OpsAnalyticsQuery *GoogleLoggingSavedQueryOpsAnalyticsQuery `field:"optional" json:"opsAnalyticsQuery" yaml:"opsAnalyticsQuery"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_logging_saved_query#timeouts GoogleLoggingSavedQuery#timeouts}
	Timeouts *GoogleLoggingSavedQueryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

