package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceHiveMetastoreConfig struct {
	// The Hive metastore schema version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#version GoogleDataprocMetastoreService#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// auxiliary_versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#auxiliary_versions GoogleDataprocMetastoreService#auxiliary_versions}
	AuxiliaryVersions interface{} `field:"optional" json:"auxiliaryVersions" yaml:"auxiliaryVersions"`
	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml). The mappings override system defaults (some keys cannot be overridden).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#config_overrides GoogleDataprocMetastoreService#config_overrides}
	ConfigOverrides *map[string]*string `field:"optional" json:"configOverrides" yaml:"configOverrides"`
	// The protocol to use for the metastore service endpoint.
	//
	// If unspecified, defaults to 'THRIFT'. Default value: "THRIFT" Possible values: ["THRIFT", "GRPC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#endpoint_protocol GoogleDataprocMetastoreService#endpoint_protocol}
	EndpointProtocol *string `field:"optional" json:"endpointProtocol" yaml:"endpointProtocol"`
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#kerberos_config GoogleDataprocMetastoreService#kerberos_config}
	KerberosConfig *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

