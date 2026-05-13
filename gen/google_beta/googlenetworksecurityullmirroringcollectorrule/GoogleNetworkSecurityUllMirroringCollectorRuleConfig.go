package googlenetworksecurityullmirroringcollectorrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityUllMirroringCollectorRuleConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#location GoogleNetworkSecurityUllMirroringCollectorRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#match GoogleNetworkSecurityUllMirroringCollectorRule#match}
	Match *GoogleNetworkSecurityUllMirroringCollectorRuleMatch `field:"required" json:"match" yaml:"match"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#ull_mirroring_collector GoogleNetworkSecurityUllMirroringCollectorRule#ull_mirroring_collector}
	UllMirroringCollector *string `field:"required" json:"ullMirroringCollector" yaml:"ullMirroringCollector"`
	// ID for the new UllMirroringCollectorRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#ull_mirroring_collector_rule_id GoogleNetworkSecurityUllMirroringCollectorRule#ull_mirroring_collector_rule_id}
	UllMirroringCollectorRuleId *string `field:"required" json:"ullMirroringCollectorRuleId" yaml:"ullMirroringCollectorRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#id GoogleNetworkSecurityUllMirroringCollectorRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#labels GoogleNetworkSecurityUllMirroringCollectorRule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#project GoogleNetworkSecurityUllMirroringCollectorRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_ull_mirroring_collector_rule#timeouts GoogleNetworkSecurityUllMirroringCollectorRule#timeouts}
	Timeouts *GoogleNetworkSecurityUllMirroringCollectorRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

