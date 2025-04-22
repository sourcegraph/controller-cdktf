package googlegkeonpremvmwareadmincluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeonpremVmwareAdminClusterConfig struct {
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
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#location GoogleGkeonpremVmwareAdminCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The VMware admin cluster resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#name GoogleGkeonpremVmwareAdminCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#network_config GoogleGkeonpremVmwareAdminCluster#network_config}
	NetworkConfig *GoogleGkeonpremVmwareAdminClusterNetworkConfig `field:"required" json:"networkConfig" yaml:"networkConfig"`
	// addon_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#addon_node GoogleGkeonpremVmwareAdminCluster#addon_node}
	AddonNode *GoogleGkeonpremVmwareAdminClusterAddonNode `field:"optional" json:"addonNode" yaml:"addonNode"`
	// Annotations on the VMware Admin Cluster.
	//
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#annotations GoogleGkeonpremVmwareAdminCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// anti_affinity_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#anti_affinity_groups GoogleGkeonpremVmwareAdminCluster#anti_affinity_groups}
	AntiAffinityGroups *GoogleGkeonpremVmwareAdminClusterAntiAffinityGroups `field:"optional" json:"antiAffinityGroups" yaml:"antiAffinityGroups"`
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#authorization GoogleGkeonpremVmwareAdminCluster#authorization}
	Authorization *GoogleGkeonpremVmwareAdminClusterAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// auto_repair_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#auto_repair_config GoogleGkeonpremVmwareAdminCluster#auto_repair_config}
	AutoRepairConfig *GoogleGkeonpremVmwareAdminClusterAutoRepairConfig `field:"optional" json:"autoRepairConfig" yaml:"autoRepairConfig"`
	// The bootstrap cluster this VMware admin cluster belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#bootstrap_cluster_membership GoogleGkeonpremVmwareAdminCluster#bootstrap_cluster_membership}
	BootstrapClusterMembership *string `field:"optional" json:"bootstrapClusterMembership" yaml:"bootstrapClusterMembership"`
	// control_plane_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#control_plane_node GoogleGkeonpremVmwareAdminCluster#control_plane_node}
	ControlPlaneNode *GoogleGkeonpremVmwareAdminClusterControlPlaneNode `field:"optional" json:"controlPlaneNode" yaml:"controlPlaneNode"`
	// A human readable description of this VMware admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#description GoogleGkeonpremVmwareAdminCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#id GoogleGkeonpremVmwareAdminCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The OS image type for the VMware admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#image_type GoogleGkeonpremVmwareAdminCluster#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#load_balancer GoogleGkeonpremVmwareAdminCluster#load_balancer}
	LoadBalancer *GoogleGkeonpremVmwareAdminClusterLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// The Anthos clusters on the VMware version for the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#on_prem_version GoogleGkeonpremVmwareAdminCluster#on_prem_version}
	OnPremVersion *string `field:"optional" json:"onPremVersion" yaml:"onPremVersion"`
	// platform_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#platform_config GoogleGkeonpremVmwareAdminCluster#platform_config}
	PlatformConfig *GoogleGkeonpremVmwareAdminClusterPlatformConfig `field:"optional" json:"platformConfig" yaml:"platformConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#project GoogleGkeonpremVmwareAdminCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#timeouts GoogleGkeonpremVmwareAdminCluster#timeouts}
	Timeouts *GoogleGkeonpremVmwareAdminClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vcenter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#vcenter GoogleGkeonpremVmwareAdminCluster#vcenter}
	Vcenter *GoogleGkeonpremVmwareAdminClusterVcenter `field:"optional" json:"vcenter" yaml:"vcenter"`
}

