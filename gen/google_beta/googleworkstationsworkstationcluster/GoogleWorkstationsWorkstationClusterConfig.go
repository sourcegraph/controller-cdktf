package googleworkstationsworkstationcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkstationsWorkstationClusterConfig struct {
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
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#network GoogleWorkstationsWorkstationCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	//
	// Must be part of the subnetwork specified for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#subnetwork GoogleWorkstationsWorkstationCluster#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
	// ID to use for the workstation cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#workstation_cluster_id GoogleWorkstationsWorkstationCluster#workstation_cluster_id}
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// Client-specified annotations. This is distinct from labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#annotations GoogleWorkstationsWorkstationCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Human-readable name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#display_name GoogleWorkstationsWorkstationCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// domain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#domain_config GoogleWorkstationsWorkstationCluster#domain_config}
	DomainConfig *GoogleWorkstationsWorkstationClusterDomainConfig `field:"optional" json:"domainConfig" yaml:"domainConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#id GoogleWorkstationsWorkstationCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#labels GoogleWorkstationsWorkstationCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the workstation cluster should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#location GoogleWorkstationsWorkstationCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// private_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#private_cluster_config GoogleWorkstationsWorkstationCluster#private_cluster_config}
	PrivateClusterConfig *GoogleWorkstationsWorkstationClusterPrivateClusterConfig `field:"optional" json:"privateClusterConfig" yaml:"privateClusterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#project GoogleWorkstationsWorkstationCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_cluster#timeouts GoogleWorkstationsWorkstationCluster#timeouts}
	Timeouts *GoogleWorkstationsWorkstationClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

