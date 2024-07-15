package googlevmwareenginecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVmwareengineClusterConfig struct {
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
	// The ID of the Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#name GoogleVmwareengineCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The resource name of the private cloud to create a new cluster in.
	//
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#parent GoogleVmwareengineCluster#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#id GoogleVmwareengineCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// node_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#node_type_configs GoogleVmwareengineCluster#node_type_configs}
	NodeTypeConfigs interface{} `field:"optional" json:"nodeTypeConfigs" yaml:"nodeTypeConfigs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_cluster#timeouts GoogleVmwareengineCluster#timeouts}
	Timeouts *GoogleVmwareengineClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

