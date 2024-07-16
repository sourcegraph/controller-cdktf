package googleintegrationconnectorsconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIntegrationConnectorsConnectionConfig struct {
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
	// connectorVersion of the Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#connector_version GoogleIntegrationConnectorsConnection#connector_version}
	ConnectorVersion *string `field:"required" json:"connectorVersion" yaml:"connectorVersion"`
	// Location in which Connection needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#location GoogleIntegrationConnectorsConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of Connection needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#name GoogleIntegrationConnectorsConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#auth_config GoogleIntegrationConnectorsConnection#auth_config}
	AuthConfig *GoogleIntegrationConnectorsConnectionAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// config_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#config_variable GoogleIntegrationConnectorsConnection#config_variable}
	ConfigVariable interface{} `field:"optional" json:"configVariable" yaml:"configVariable"`
	// An arbitrary description for the Conection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#description GoogleIntegrationConnectorsConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#destination_config GoogleIntegrationConnectorsConnection#destination_config}
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// eventing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#eventing_config GoogleIntegrationConnectorsConnection#eventing_config}
	EventingConfig *GoogleIntegrationConnectorsConnectionEventingConfig `field:"optional" json:"eventingConfig" yaml:"eventingConfig"`
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION", "ONLY_EVENTING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#eventing_enablement_type GoogleIntegrationConnectorsConnection#eventing_enablement_type}
	EventingEnablementType *string `field:"optional" json:"eventingEnablementType" yaml:"eventingEnablementType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#id GoogleIntegrationConnectorsConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#labels GoogleIntegrationConnectorsConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// lock_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#lock_config GoogleIntegrationConnectorsConnection#lock_config}
	LockConfig *GoogleIntegrationConnectorsConnectionLockConfig `field:"optional" json:"lockConfig" yaml:"lockConfig"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#log_config GoogleIntegrationConnectorsConnection#log_config}
	LogConfig *GoogleIntegrationConnectorsConnectionLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#node_config GoogleIntegrationConnectorsConnection#node_config}
	NodeConfig *GoogleIntegrationConnectorsConnectionNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#project GoogleIntegrationConnectorsConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Service account needed for runtime plane to access Google Cloud resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#service_account GoogleIntegrationConnectorsConnection#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// ssl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#ssl_config GoogleIntegrationConnectorsConnection#ssl_config}
	SslConfig *GoogleIntegrationConnectorsConnectionSslConfig `field:"optional" json:"sslConfig" yaml:"sslConfig"`
	// Suspended indicates if a user has suspended a connection or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#suspended GoogleIntegrationConnectorsConnection#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#timeouts GoogleIntegrationConnectorsConnection#timeouts}
	Timeouts *GoogleIntegrationConnectorsConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

