package googlecontainerattachedcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleContainerAttachedClusterConfig struct {
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
	// The Kubernetes distribution of the underlying attached cluster. Supported values: "eks", "aks".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#distribution GoogleContainerAttachedCluster#distribution}
	Distribution *string `field:"required" json:"distribution" yaml:"distribution"`
	// fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#fleet GoogleContainerAttachedCluster#fleet}
	Fleet *GoogleContainerAttachedClusterFleet `field:"required" json:"fleet" yaml:"fleet"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#location GoogleContainerAttachedCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#name GoogleContainerAttachedCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// oidc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#oidc_config GoogleContainerAttachedCluster#oidc_config}
	OidcConfig *GoogleContainerAttachedClusterOidcConfig `field:"required" json:"oidcConfig" yaml:"oidcConfig"`
	// The platform version for the cluster (e.g. '1.23.0-gke.1').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#platform_version GoogleContainerAttachedCluster#platform_version}
	PlatformVersion *string `field:"required" json:"platformVersion" yaml:"platformVersion"`
	// Optional.
	//
	// Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#annotations GoogleContainerAttachedCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#authorization GoogleContainerAttachedCluster#authorization}
	Authorization *GoogleContainerAttachedClusterAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#binary_authorization GoogleContainerAttachedCluster#binary_authorization}
	BinaryAuthorization *GoogleContainerAttachedClusterBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// Policy to determine what flags to send on delete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#deletion_policy GoogleContainerAttachedCluster#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#description GoogleContainerAttachedCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#id GoogleContainerAttachedCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#logging_config GoogleContainerAttachedCluster#logging_config}
	LoggingConfig *GoogleContainerAttachedClusterLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#monitoring_config GoogleContainerAttachedCluster#monitoring_config}
	MonitoringConfig *GoogleContainerAttachedClusterMonitoringConfig `field:"optional" json:"monitoringConfig" yaml:"monitoringConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#project GoogleContainerAttachedCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#proxy_config GoogleContainerAttachedCluster#proxy_config}
	ProxyConfig *GoogleContainerAttachedClusterProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#timeouts GoogleContainerAttachedCluster#timeouts}
	Timeouts *GoogleContainerAttachedClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

