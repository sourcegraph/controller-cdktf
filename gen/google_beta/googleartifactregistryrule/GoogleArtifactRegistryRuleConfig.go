package googleartifactregistryrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleArtifactRegistryRuleConfig struct {
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
	// The last part of the repository name, for example: "repo1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#repository_id GoogleArtifactRegistryRule#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// The rule id to use for this repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#rule_id GoogleArtifactRegistryRule#rule_id}
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// The action this rule takes. Possible values: ["ACTION_UNSPECIFIED", "ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#action GoogleArtifactRegistryRule#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#condition GoogleArtifactRegistryRule#condition}
	Condition *GoogleArtifactRegistryRuleCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#id GoogleArtifactRegistryRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the repository's location.
	//
	// In addition to specific regions,
	// special values for multi-region locations are 'asia', 'europe', and 'us'.
	// See [here](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations),
	// or use the
	// [google_artifact_registry_locations](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/artifact_registry_locations)
	// data source for possible values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#location GoogleArtifactRegistryRule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The operation the rule applies to. Possible values: ["OPERATION_UNSPECIFIED", "DOWNLOAD"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#operation GoogleArtifactRegistryRule#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// The package ID the rule applies to. If empty, this rule applies to all packages inside the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#package_id GoogleArtifactRegistryRule#package_id}
	PackageId *string `field:"optional" json:"packageId" yaml:"packageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#project GoogleArtifactRegistryRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_artifact_registry_rule#timeouts GoogleArtifactRegistryRule#timeouts}
	Timeouts *GoogleArtifactRegistryRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

