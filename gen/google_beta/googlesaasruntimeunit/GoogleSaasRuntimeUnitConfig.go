package googlesaasruntimeunit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSaasRuntimeUnitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#location GoogleSaasRuntimeUnit#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID value for the new unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#unit_id GoogleSaasRuntimeUnit#unit_id}
	UnitId *string `field:"required" json:"unitId" yaml:"unitId"`
	// Annotations is an unstructured key-value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// More info: https://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#annotations GoogleSaasRuntimeUnit#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#id GoogleSaasRuntimeUnit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels on the resource, which can be used for categorization. similar to Kubernetes resource labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#labels GoogleSaasRuntimeUnit#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#maintenance GoogleSaasRuntimeUnit#maintenance}
	Maintenance *GoogleSaasRuntimeUnitMaintenance `field:"optional" json:"maintenance" yaml:"maintenance"`
	// Indicates whether the Unit life cycle is controlled by the user or by the system. Immutable once created. Possible values: MANAGEMENT_MODE_USER MANAGEMENT_MODE_SYSTEM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#management_mode GoogleSaasRuntimeUnit#management_mode}
	ManagementMode *string `field:"optional" json:"managementMode" yaml:"managementMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#project GoogleSaasRuntimeUnit#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Reference to the Saas Tenant resource this unit belongs to.
	//
	// This for
	// example informs the maintenance policies to use for scheduling future
	// updates on a unit. (optional and immutable once created)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#tenant GoogleSaasRuntimeUnit#tenant}
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#timeouts GoogleSaasRuntimeUnit#timeouts}
	Timeouts *GoogleSaasRuntimeUnitTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Reference to the UnitKind this Unit belongs to. Immutable once set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_unit#unit_kind GoogleSaasRuntimeUnit#unit_kind}
	UnitKind *string `field:"optional" json:"unitKind" yaml:"unitKind"`
}

