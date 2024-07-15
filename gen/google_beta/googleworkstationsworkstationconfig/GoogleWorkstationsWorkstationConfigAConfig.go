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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#location GoogleWorkstationsWorkstationConfigA#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the parent workstation cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#workstation_cluster_id GoogleWorkstationsWorkstationConfigA#workstation_cluster_id}
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// The ID to be assigned to the workstation cluster config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#workstation_config_id GoogleWorkstationsWorkstationConfigA#workstation_config_id}
	WorkstationConfigId *string `field:"required" json:"workstationConfigId" yaml:"workstationConfigId"`
	// Client-specified annotations. This is distinct from labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#annotations GoogleWorkstationsWorkstationConfigA#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#container GoogleWorkstationsWorkstationConfigA#container}
	Container *GoogleWorkstationsWorkstationConfigContainer `field:"optional" json:"container" yaml:"container"`
	// Disables support for plain TCP connections in the workstation.
	//
	// By default the service supports TCP connections via a websocket relay. Setting this option to true disables that relay, which prevents the usage of services that require plain tcp connections, such as ssh. When enabled, all communication must occur over https or wss.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#disable_tcp_connections GoogleWorkstationsWorkstationConfigA#disable_tcp_connections}
	DisableTcpConnections interface{} `field:"optional" json:"disableTcpConnections" yaml:"disableTcpConnections"`
	// Human-readable name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#display_name GoogleWorkstationsWorkstationConfigA#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Whether to enable Linux 'auditd' logging on the workstation.
	//
	// When enabled, a service account must also be specified that has 'logging.buckets.write' permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#enable_audit_agent GoogleWorkstationsWorkstationConfigA#enable_audit_agent}
	EnableAuditAgent interface{} `field:"optional" json:"enableAuditAgent" yaml:"enableAuditAgent"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#encryption_key GoogleWorkstationsWorkstationConfigA#encryption_key}
	EncryptionKey *GoogleWorkstationsWorkstationConfigEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// ephemeral_directories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#ephemeral_directories GoogleWorkstationsWorkstationConfigA#ephemeral_directories}
	EphemeralDirectories interface{} `field:"optional" json:"ephemeralDirectories" yaml:"ephemeralDirectories"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#host GoogleWorkstationsWorkstationConfigA#host}
	Host *GoogleWorkstationsWorkstationConfigHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#id GoogleWorkstationsWorkstationConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// How long to wait before automatically stopping an instance that hasn't recently received any user traffic.
	//
	// A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#idle_timeout GoogleWorkstationsWorkstationConfigA#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#labels GoogleWorkstationsWorkstationConfigA#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// persistent_directories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#persistent_directories GoogleWorkstationsWorkstationConfigA#persistent_directories}
	PersistentDirectories interface{} `field:"optional" json:"persistentDirectories" yaml:"persistentDirectories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#project GoogleWorkstationsWorkstationConfigA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// readiness_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#readiness_checks GoogleWorkstationsWorkstationConfigA#readiness_checks}
	ReadinessChecks interface{} `field:"optional" json:"readinessChecks" yaml:"readinessChecks"`
	// Specifies the zones used to replicate the VM and disk resources within the region.
	//
	// If set, exactly two zones within the workstation cluster's region must be specified—for example, '['us-central1-a', 'us-central1-f']'.
	// If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#replica_zones GoogleWorkstationsWorkstationConfigA#replica_zones}
	ReplicaZones *[]*string `field:"optional" json:"replicaZones" yaml:"replicaZones"`
	// How long to wait before automatically stopping a workstation after it was started.
	//
	// A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if 'encryption_key' is set. Defaults to 12 hours.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#running_timeout GoogleWorkstationsWorkstationConfigA#running_timeout}
	RunningTimeout *string `field:"optional" json:"runningTimeout" yaml:"runningTimeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#timeouts GoogleWorkstationsWorkstationConfigA#timeouts}
	Timeouts *GoogleWorkstationsWorkstationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

