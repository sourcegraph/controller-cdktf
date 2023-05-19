package deviceposturerule


type DevicePostureRuleInput struct {
	// Specific volume(s) to check for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#check_disks DevicePostureRule#check_disks}
	CheckDisks *[]*string `field:"optional" json:"checkDisks" yaml:"checkDisks"`
	// The workspace one device compliance status. Available values: `compliant`, `noncompliant`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#compliance_status DevicePostureRule#compliance_status}
	ComplianceStatus *string `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// The workspace one connection id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#connection_id DevicePostureRule#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The domain that the client must join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#domain DevicePostureRule#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// True if the firewall must be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#enabled DevicePostureRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Checks if the file should exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#exists DevicePostureRule#exists}
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// The Teams List id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#id DevicePostureRule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The version comparison operator. Available values: `>`, `>=`, `<`, `<=`, `==`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#operator DevicePostureRule#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// OS signal score from Crowdstrike. Value must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#os DevicePostureRule#os}
	Os *string `field:"optional" json:"os" yaml:"os"`
	// The operating system excluding version information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#os_distro_name DevicePostureRule#os_distro_name}
	OsDistroName *string `field:"optional" json:"osDistroName" yaml:"osDistroName"`
	// The operating system version excluding OS name information or release name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#os_distro_revision DevicePostureRule#os_distro_revision}
	OsDistroRevision *string `field:"optional" json:"osDistroRevision" yaml:"osDistroRevision"`
	// Overall ZTA score from Crowdstrike. Value must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#overall DevicePostureRule#overall}
	Overall *string `field:"optional" json:"overall" yaml:"overall"`
	// The path to the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#path DevicePostureRule#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// True if all drives must be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#require_all DevicePostureRule#require_all}
	RequireAll interface{} `field:"optional" json:"requireAll" yaml:"requireAll"`
	// Checks if the application should be running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#running DevicePostureRule#running}
	Running interface{} `field:"optional" json:"running" yaml:"running"`
	// Sensor signal score from Crowdstrike. Value must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#sensor_config DevicePostureRule#sensor_config}
	SensorConfig *string `field:"optional" json:"sensorConfig" yaml:"sensorConfig"`
	// The sha256 hash of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#sha256 DevicePostureRule#sha256}
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
	// The thumbprint of the file certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#thumbprint DevicePostureRule#thumbprint}
	Thumbprint *string `field:"optional" json:"thumbprint" yaml:"thumbprint"`
	// The operating system semantic version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#version DevicePostureRule#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The version comparison operator for crowdstrike. Available values: `>`, `>=`, `<`, `<=`, `==`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_posture_rule#version_operator DevicePostureRule#version_operator}
	VersionOperator *string `field:"optional" json:"versionOperator" yaml:"versionOperator"`
}

