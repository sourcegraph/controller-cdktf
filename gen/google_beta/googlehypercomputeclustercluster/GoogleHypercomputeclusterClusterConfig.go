package googlehypercomputeclustercluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleHypercomputeclusterClusterConfig struct {
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
	// ID of the cluster to create.
	//
	// Must start with a lowercase letter,
	// use only lowercase letters and numbers, and be at most 10 characters long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#cluster_id GoogleHypercomputeclusterCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#location GoogleHypercomputeclusterCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// compute_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#compute_resources GoogleHypercomputeclusterCluster#compute_resources}
	ComputeResources interface{} `field:"optional" json:"computeResources" yaml:"computeResources"`
	// User-provided description of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#description GoogleHypercomputeclusterCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#id GoogleHypercomputeclusterCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// [Labels](https://cloud.google.com/compute/docs/labeling-resources) applied to the cluster. Labels can be used to organize clusters and to filter them in queries.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#labels GoogleHypercomputeclusterCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#network_resources GoogleHypercomputeclusterCluster#network_resources}
	NetworkResources interface{} `field:"optional" json:"networkResources" yaml:"networkResources"`
	// orchestrator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#orchestrator GoogleHypercomputeclusterCluster#orchestrator}
	Orchestrator *GoogleHypercomputeclusterClusterOrchestrator `field:"optional" json:"orchestrator" yaml:"orchestrator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#project GoogleHypercomputeclusterCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// storage_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#storage_resources GoogleHypercomputeclusterCluster#storage_resources}
	StorageResources interface{} `field:"optional" json:"storageResources" yaml:"storageResources"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#timeouts GoogleHypercomputeclusterCluster#timeouts}
	Timeouts *GoogleHypercomputeclusterClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

