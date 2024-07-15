package googleworkstationsworkstation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkstationsWorkstationConfig struct {
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
	// The location where the workstation parent resources reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#location GoogleWorkstationsWorkstation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the parent workstation cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#workstation_cluster_id GoogleWorkstationsWorkstation#workstation_cluster_id}
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// The ID of the parent workstation cluster config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#workstation_config_id GoogleWorkstationsWorkstation#workstation_config_id}
	WorkstationConfigId *string `field:"required" json:"workstationConfigId" yaml:"workstationConfigId"`
	// ID to use for the workstation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#workstation_id GoogleWorkstationsWorkstation#workstation_id}
	WorkstationId *string `field:"required" json:"workstationId" yaml:"workstationId"`
	// Client-specified annotations. This is distinct from labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#annotations GoogleWorkstationsWorkstation#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Human-readable name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#display_name GoogleWorkstationsWorkstation#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// 'Client-specified environment variables passed to the workstation container's entrypoint.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#env GoogleWorkstationsWorkstation#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#id GoogleWorkstationsWorkstation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#labels GoogleWorkstationsWorkstation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#project GoogleWorkstationsWorkstation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation#timeouts GoogleWorkstationsWorkstation#timeouts}
	Timeouts *GoogleWorkstationsWorkstationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

