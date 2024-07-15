package googledatafusioninstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataFusionInstanceConfig struct {
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
	// The ID of the instance or a fully qualified identifier for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#name GoogleDataFusionInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Represents the type of Data Fusion instance.
	//
	// Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	// using point and click UI. However, there are certain limitations, such as fewer number
	// of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	// available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	// with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	// pipelines at low cost. Possible values: ["BASIC", "ENTERPRISE", "DEVELOPER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#type GoogleDataFusionInstance#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#accelerators GoogleDataFusionInstance#accelerators}
	Accelerators interface{} `field:"optional" json:"accelerators" yaml:"accelerators"`
	// crypto_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#crypto_key_config GoogleDataFusionInstance#crypto_key_config}
	CryptoKeyConfig *GoogleDataFusionInstanceCryptoKeyConfig `field:"optional" json:"cryptoKeyConfig" yaml:"cryptoKeyConfig"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#dataproc_service_account GoogleDataFusionInstance#dataproc_service_account}
	DataprocServiceAccount *string `field:"optional" json:"dataprocServiceAccount" yaml:"dataprocServiceAccount"`
	// An optional description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#description GoogleDataFusionInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#display_name GoogleDataFusionInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Option to enable granular role-based access control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#enable_rbac GoogleDataFusionInstance#enable_rbac}
	EnableRbac interface{} `field:"optional" json:"enableRbac" yaml:"enableRbac"`
	// Option to enable Stackdriver Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#enable_stackdriver_logging GoogleDataFusionInstance#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#enable_stackdriver_monitoring GoogleDataFusionInstance#enable_stackdriver_monitoring}
	EnableStackdriverMonitoring interface{} `field:"optional" json:"enableStackdriverMonitoring" yaml:"enableStackdriverMonitoring"`
	// event_publish_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#event_publish_config GoogleDataFusionInstance#event_publish_config}
	EventPublishConfig *GoogleDataFusionInstanceEventPublishConfig `field:"optional" json:"eventPublishConfig" yaml:"eventPublishConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#id GoogleDataFusionInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine VMs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#labels GoogleDataFusionInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#network_config GoogleDataFusionInstance#network_config}
	NetworkConfig *GoogleDataFusionInstanceNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#options GoogleDataFusionInstance#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Specifies whether the Data Fusion instance should be private.
	//
	// If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#private_instance GoogleDataFusionInstance#private_instance}
	PrivateInstance interface{} `field:"optional" json:"privateInstance" yaml:"privateInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#project GoogleDataFusionInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the Data Fusion instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#region GoogleDataFusionInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#timeouts GoogleDataFusionInstance#timeouts}
	Timeouts *GoogleDataFusionInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Current version of the Data Fusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#version GoogleDataFusionInstance#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#zone GoogleDataFusionInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

