package googlelogginglogscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleLoggingLogScopeConfig struct {
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
	// The resource name of the log scope. For example: \'projects/my-project/locations/global/logScopes/my-log-scope\'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#name GoogleLoggingLogScope#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Names of one or more parent resources : *  \'projects/[PROJECT_ID]\' May alternatively be one or more views : * \'projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/views/[VIEW_ID]\' A log scope can include a maximum of 50 projects and a maximum of 100 resources in total.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#resource_names GoogleLoggingLogScope#resource_names}
	ResourceNames *[]*string `field:"required" json:"resourceNames" yaml:"resourceNames"`
	// Describes this log scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#description GoogleLoggingLogScope#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#id GoogleLoggingLogScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the resource. The only supported location is global so far.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#location GoogleLoggingLogScope#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The parent of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#parent GoogleLoggingLogScope#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_logging_log_scope#timeouts GoogleLoggingLogScope#timeouts}
	Timeouts *GoogleLoggingLogScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

