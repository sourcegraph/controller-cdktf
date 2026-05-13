package googlesecuresourcemanagerhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecureSourceManagerHookConfig struct {
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
	// The ID for the Hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#hook_id GoogleSecureSourceManagerHook#hook_id}
	HookId *string `field:"required" json:"hookId" yaml:"hookId"`
	// The location for the Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#location GoogleSecureSourceManagerHook#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID for the Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#repository_id GoogleSecureSourceManagerHook#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// The target URI to which the payloads will be delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#target_uri GoogleSecureSourceManagerHook#target_uri}
	TargetUri *string `field:"required" json:"targetUri" yaml:"targetUri"`
	// Determines if the hook disabled or not. Set to true to stop sending traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#disabled GoogleSecureSourceManagerHook#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The events that trigger hook on. Possible values: ["PUSH", "PULL_REQUEST"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#events GoogleSecureSourceManagerHook#events}
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#id GoogleSecureSourceManagerHook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#project GoogleSecureSourceManagerHook#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// push_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#push_option GoogleSecureSourceManagerHook#push_option}
	PushOption *GoogleSecureSourceManagerHookPushOption `field:"optional" json:"pushOption" yaml:"pushOption"`
	// The sensitive query string to be appended to the target URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#sensitive_query_string GoogleSecureSourceManagerHook#sensitive_query_string}
	SensitiveQueryString *string `field:"optional" json:"sensitiveQueryString" yaml:"sensitiveQueryString"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_hook#timeouts GoogleSecureSourceManagerHook#timeouts}
	Timeouts *GoogleSecureSourceManagerHookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

