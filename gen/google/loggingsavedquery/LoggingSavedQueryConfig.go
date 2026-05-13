package loggingsavedquery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingSavedQueryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#display_name LoggingSavedQuery#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location of the resource see [supported regions](https://docs.cloud.google.com/logging/docs/region-support#bucket-regions).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#location LoggingSavedQuery#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the saved query. For example: 'my-saved-query'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#name LoggingSavedQuery#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#parent LoggingSavedQuery#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The visibility of the saved query. Possible values: ["SHARED", "PRIVATE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#visibility LoggingSavedQuery#visibility}
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// A description of the saved query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#description LoggingSavedQuery#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#id LoggingSavedQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#logging_query LoggingSavedQuery#logging_query}
	LoggingQuery *LoggingSavedQueryLoggingQuery `field:"optional" json:"loggingQuery" yaml:"loggingQuery"`
	// ops_analytics_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#ops_analytics_query LoggingSavedQuery#ops_analytics_query}
	OpsAnalyticsQuery *LoggingSavedQueryOpsAnalyticsQuery `field:"optional" json:"opsAnalyticsQuery" yaml:"opsAnalyticsQuery"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#timeouts LoggingSavedQuery#timeouts}
	Timeouts *LoggingSavedQueryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

