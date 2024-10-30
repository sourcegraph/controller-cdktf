package policyruleidpdiscovery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyRuleIdpDiscoveryConfig struct {
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
	// Policy Rule Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#name PolicyRuleIdpDiscovery#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// app_exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#app_exclude PolicyRuleIdpDiscovery#app_exclude}
	AppExclude interface{} `field:"optional" json:"appExclude" yaml:"appExclude"`
	// app_include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#app_include PolicyRuleIdpDiscovery#app_include}
	AppInclude interface{} `field:"optional" json:"appInclude" yaml:"appInclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#id PolicyRuleIdpDiscovery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The identifier for the Idp the rule should route to if all conditions are met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#idp_id PolicyRuleIdpDiscovery#idp_id}
	IdpId *string `field:"optional" json:"idpId" yaml:"idpId"`
	// Type of Idp. One of: `SAML2`, `IWA`, `AgentlessDSSO`, `X509`, `FACEBOOK`, `GOOGLE`, `LINKEDIN`, `MICROSOFT`, `OIDC`. Default: `OKTA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#idp_type PolicyRuleIdpDiscovery#idp_type}
	IdpType *string `field:"optional" json:"idpType" yaml:"idpType"`
	// Network selection mode: `ANYWHERE`, `ZONE`, `ON_NETWORK`, or `OFF_NETWORK`. Default: `ANYWHERE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#network_connection PolicyRuleIdpDiscovery#network_connection}
	NetworkConnection *string `field:"optional" json:"networkConnection" yaml:"networkConnection"`
	// Required if `network_connection` = `ZONE`. Indicates the network zones to exclude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#network_excludes PolicyRuleIdpDiscovery#network_excludes}
	NetworkExcludes *[]*string `field:"optional" json:"networkExcludes" yaml:"networkExcludes"`
	// Required if `network_connection` = `ZONE`. Indicates the network zones to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#network_includes PolicyRuleIdpDiscovery#network_includes}
	NetworkIncludes *[]*string `field:"optional" json:"networkIncludes" yaml:"networkIncludes"`
	// platform_include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#platform_include PolicyRuleIdpDiscovery#platform_include}
	PlatformInclude interface{} `field:"optional" json:"platformInclude" yaml:"platformInclude"`
	// Policy ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#policy_id PolicyRuleIdpDiscovery#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Rule priority.
	//
	// This attribute can be set to a valid priority. To avoid an endless diff situation an error is thrown if an invalid property is provided. The Okta API defaults to the last (lowest) if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#priority PolicyRuleIdpDiscovery#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Policy Rule Status: `ACTIVE` or `INACTIVE`. Default: `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#status PolicyRuleIdpDiscovery#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Profile attribute matching can only have a single value that describes the type indicated in `user_identifier_type`.
	//
	// This is the attribute or identifier that the `user_identifier_patterns` are checked against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#user_identifier_attribute PolicyRuleIdpDiscovery#user_identifier_attribute}
	UserIdentifierAttribute *string `field:"optional" json:"userIdentifierAttribute" yaml:"userIdentifierAttribute"`
	// user_identifier_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#user_identifier_patterns PolicyRuleIdpDiscovery#user_identifier_patterns}
	UserIdentifierPatterns interface{} `field:"optional" json:"userIdentifierPatterns" yaml:"userIdentifierPatterns"`
	// One of: `IDENTIFIER`, `ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_idp_discovery#user_identifier_type PolicyRuleIdpDiscovery#user_identifier_type}
	UserIdentifierType *string `field:"optional" json:"userIdentifierType" yaml:"userIdentifierType"`
}

