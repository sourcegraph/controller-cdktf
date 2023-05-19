package googleworkstationsworkstationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkstationsWorkstationConfigAConfig struct {
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
	// The location where the workstation cluster config should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#location GoogleWorkstationsWorkstationConfigA#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the parent workstation cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#workstation_cluster_id GoogleWorkstationsWorkstationConfigA#workstation_cluster_id}
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// The ID to be assigned to the workstation cluster config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#workstation_config_id GoogleWorkstationsWorkstationConfigA#workstation_config_id}
	WorkstationConfigId *string `field:"required" json:"workstationConfigId" yaml:"workstationConfigId"`
	// Client-specified annotations. This is distinct from labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#annotations GoogleWorkstationsWorkstationConfigA#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#container GoogleWorkstationsWorkstationConfigA#container}
	Container *GoogleWorkstationsWorkstationConfigContainer `field:"optional" json:"container" yaml:"container"`
	// Human-readable name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#display_name GoogleWorkstationsWorkstationConfigA#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#encryption_key GoogleWorkstationsWorkstationConfigA#encryption_key}
	EncryptionKey *GoogleWorkstationsWorkstationConfigEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#host GoogleWorkstationsWorkstationConfigA#host}
	Host *GoogleWorkstationsWorkstationConfigHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#id GoogleWorkstationsWorkstationConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#labels GoogleWorkstationsWorkstationConfigA#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// persistent_directories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#persistent_directories GoogleWorkstationsWorkstationConfigA#persistent_directories}
	PersistentDirectories interface{} `field:"optional" json:"persistentDirectories" yaml:"persistentDirectories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#project GoogleWorkstationsWorkstationConfigA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#timeouts GoogleWorkstationsWorkstationConfigA#timeouts}
	Timeouts *GoogleWorkstationsWorkstationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

