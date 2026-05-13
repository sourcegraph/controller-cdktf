package googlesaasruntimerelease

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSaasRuntimeReleaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#location GoogleSaasRuntimeRelease#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID value for the new release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#release_id GoogleSaasRuntimeRelease#release_id}
	ReleaseId *string `field:"required" json:"releaseId" yaml:"releaseId"`
	// Reference to the UnitKind this Release corresponds to (required and immutable once created).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#unit_kind GoogleSaasRuntimeRelease#unit_kind}
	UnitKind *string `field:"required" json:"unitKind" yaml:"unitKind"`
	// Annotations is an unstructured key-value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// More info: https://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#annotations GoogleSaasRuntimeRelease#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// blueprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#blueprint GoogleSaasRuntimeRelease#blueprint}
	Blueprint *GoogleSaasRuntimeReleaseBlueprint `field:"optional" json:"blueprint" yaml:"blueprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#id GoogleSaasRuntimeRelease#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input_variable_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#input_variable_defaults GoogleSaasRuntimeRelease#input_variable_defaults}
	InputVariableDefaults interface{} `field:"optional" json:"inputVariableDefaults" yaml:"inputVariableDefaults"`
	// The labels on the resource, which can be used for categorization. similar to Kubernetes resource labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#labels GoogleSaasRuntimeRelease#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#project GoogleSaasRuntimeRelease#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// release_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#release_requirements GoogleSaasRuntimeRelease#release_requirements}
	ReleaseRequirements *GoogleSaasRuntimeReleaseReleaseRequirements `field:"optional" json:"releaseRequirements" yaml:"releaseRequirements"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_release#timeouts GoogleSaasRuntimeRelease#timeouts}
	Timeouts *GoogleSaasRuntimeReleaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

