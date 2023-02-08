package deviceposturerule


type DevicePostureRuleInput struct {
	// The workspace one device compliance status.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#compliance_status DevicePostureRule#compliance_status}
	ComplianceStatus *string `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// The workspace one connection id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#connection_id DevicePostureRule#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The domain that the client must join.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#domain DevicePostureRule#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// True if the firewall must be enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#enabled DevicePostureRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Checks if the file should exist.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#exists DevicePostureRule#exists}
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// The Teams List id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#id DevicePostureRule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The version comparison operator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#operator DevicePostureRule#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The path to the file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#path DevicePostureRule#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// True if all drives must be encrypted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#require_all DevicePostureRule#require_all}
	RequireAll interface{} `field:"optional" json:"requireAll" yaml:"requireAll"`
	// Checks if the application should be running.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#running DevicePostureRule#running}
	Running interface{} `field:"optional" json:"running" yaml:"running"`
	// The sha256 hash of the file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#sha256 DevicePostureRule#sha256}
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
	// The thumbprint of the file certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#thumbprint DevicePostureRule#thumbprint}
	Thumbprint *string `field:"optional" json:"thumbprint" yaml:"thumbprint"`
	// The operating system semantic version.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#version DevicePostureRule#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

