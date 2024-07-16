package googlegkeonprembaremetalcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeonpremBareMetalClusterConfig struct {
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
	// The Admin Cluster this Bare Metal User Cluster belongs to.
	//
	// This is the full resource name of the Admin Cluster's hub membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#admin_cluster_membership GoogleGkeonpremBareMetalCluster#admin_cluster_membership}
	AdminClusterMembership *string `field:"required" json:"adminClusterMembership" yaml:"adminClusterMembership"`
	// A human readable description of this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#bare_metal_version GoogleGkeonpremBareMetalCluster#bare_metal_version}
	BareMetalVersion *string `field:"required" json:"bareMetalVersion" yaml:"bareMetalVersion"`
	// control_plane block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#control_plane GoogleGkeonpremBareMetalCluster#control_plane}
	ControlPlane *GoogleGkeonpremBareMetalClusterControlPlane `field:"required" json:"controlPlane" yaml:"controlPlane"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#load_balancer GoogleGkeonpremBareMetalCluster#load_balancer}
	LoadBalancer *GoogleGkeonpremBareMetalClusterLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#location GoogleGkeonpremBareMetalCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The bare metal cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#name GoogleGkeonpremBareMetalCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#network_config GoogleGkeonpremBareMetalCluster#network_config}
	NetworkConfig *GoogleGkeonpremBareMetalClusterNetworkConfig `field:"required" json:"networkConfig" yaml:"networkConfig"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#storage GoogleGkeonpremBareMetalCluster#storage}
	Storage *GoogleGkeonpremBareMetalClusterStorage `field:"required" json:"storage" yaml:"storage"`
	// Annotations on the Bare Metal User Cluster.
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#annotations GoogleGkeonpremBareMetalCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#binary_authorization GoogleGkeonpremBareMetalCluster#binary_authorization}
	BinaryAuthorization *GoogleGkeonpremBareMetalClusterBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// cluster_operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#cluster_operations GoogleGkeonpremBareMetalCluster#cluster_operations}
	ClusterOperations *GoogleGkeonpremBareMetalClusterClusterOperations `field:"optional" json:"clusterOperations" yaml:"clusterOperations"`
	// A human readable description of this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#description GoogleGkeonpremBareMetalCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#id GoogleGkeonpremBareMetalCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#maintenance_config GoogleGkeonpremBareMetalCluster#maintenance_config}
	MaintenanceConfig *GoogleGkeonpremBareMetalClusterMaintenanceConfig `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// node_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#node_access_config GoogleGkeonpremBareMetalCluster#node_access_config}
	NodeAccessConfig *GoogleGkeonpremBareMetalClusterNodeAccessConfig `field:"optional" json:"nodeAccessConfig" yaml:"nodeAccessConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#node_config GoogleGkeonpremBareMetalCluster#node_config}
	NodeConfig *GoogleGkeonpremBareMetalClusterNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// os_environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#os_environment_config GoogleGkeonpremBareMetalCluster#os_environment_config}
	OsEnvironmentConfig *GoogleGkeonpremBareMetalClusterOsEnvironmentConfig `field:"optional" json:"osEnvironmentConfig" yaml:"osEnvironmentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#project GoogleGkeonpremBareMetalCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#proxy GoogleGkeonpremBareMetalCluster#proxy}
	Proxy *GoogleGkeonpremBareMetalClusterProxy `field:"optional" json:"proxy" yaml:"proxy"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#security_config GoogleGkeonpremBareMetalCluster#security_config}
	SecurityConfig *GoogleGkeonpremBareMetalClusterSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#timeouts GoogleGkeonpremBareMetalCluster#timeouts}
	Timeouts *GoogleGkeonpremBareMetalClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#upgrade_policy GoogleGkeonpremBareMetalCluster#upgrade_policy}
	UpgradePolicy *GoogleGkeonpremBareMetalClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
}

