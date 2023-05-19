package googlecloudiotregistry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudiotRegistryConfig struct {
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
	// A unique name for the resource, required by device registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#name GoogleCloudiotRegistry#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#credentials GoogleCloudiotRegistry#credentials}
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// event_notification_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#event_notification_configs GoogleCloudiotRegistry#event_notification_configs}
	EventNotificationConfigs interface{} `field:"optional" json:"eventNotificationConfigs" yaml:"eventNotificationConfigs"`
	// Activate or deactivate HTTP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#http_config GoogleCloudiotRegistry#http_config}
	HttpConfig *map[string]*string `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#id GoogleCloudiotRegistry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The default logging verbosity for activity from devices in this registry.
	//
	// Specifies which events should be written to logs. For
	// example, if the LogLevel is ERROR, only events that terminate in
	// errors will be logged. LogLevel is inclusive; enabling INFO logging
	// will also enable ERROR logging. Default value: "NONE" Possible values: ["NONE", "ERROR", "INFO", "DEBUG"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#log_level GoogleCloudiotRegistry#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Activate or deactivate MQTT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#mqtt_config GoogleCloudiotRegistry#mqtt_config}
	MqttConfig *map[string]*string `field:"optional" json:"mqttConfig" yaml:"mqttConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#project GoogleCloudiotRegistry#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the created registry should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#region GoogleCloudiotRegistry#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A PubSub topic to publish device state updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#state_notification_config GoogleCloudiotRegistry#state_notification_config}
	StateNotificationConfig *map[string]*string `field:"optional" json:"stateNotificationConfig" yaml:"stateNotificationConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_registry#timeouts GoogleCloudiotRegistry#timeouts}
	Timeouts *GoogleCloudiotRegistryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

