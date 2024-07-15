package googleappenginedomainmapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAppEngineDomainMappingConfig struct {
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
	// Relative name of the domain serving the application. Example: example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_domain_mapping#domain_name GoogleAppEngineDomainMapping#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_domain_mapping#id GoogleAppEngineDomainMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the domain creation should override any existing mappings for this domain.
	//
	// By default, overrides are rejected. Default value: "STRICT" Possible values: ["STRICT", "OVERRIDE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_domain_mapping#override_strategy GoogleAppEngineDomainMapping#override_strategy}
	OverrideStrategy *string `field:"optional" json:"overrideStrategy" yaml:"overrideStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_domain_mapping#project GoogleAppEngineDomainMapping#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// ssl_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_domain_mapping#ssl_settings GoogleAppEngineDomainMapping#ssl_settings}
	SslSettings *GoogleAppEngineDomainMappingSslSettings `field:"optional" json:"sslSettings" yaml:"sslSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_domain_mapping#timeouts GoogleAppEngineDomainMapping#timeouts}
	Timeouts *GoogleAppEngineDomainMappingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

