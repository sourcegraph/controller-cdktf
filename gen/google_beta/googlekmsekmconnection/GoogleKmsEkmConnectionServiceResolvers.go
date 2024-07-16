package googlekmsekmconnection


type GoogleKmsEkmConnectionServiceResolvers struct {
	// Required. The hostname of the EKM replica used at TLS and HTTP layers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_ekm_connection#hostname GoogleKmsEkmConnection#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// server_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_ekm_connection#server_certificates GoogleKmsEkmConnection#server_certificates}
	ServerCertificates interface{} `field:"required" json:"serverCertificates" yaml:"serverCertificates"`
	// Required.
	//
	// The resource name of the Service Directory service pointing to an EKM replica, in the format projects/* /locations/* /namespaces/* /services/*
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_ekm_connection#service_directory_service GoogleKmsEkmConnection#service_directory_service}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	ServiceDirectoryService *string `field:"required" json:"serviceDirectoryService" yaml:"serviceDirectoryService"`
	// Optional.
	//
	// The filter applied to the endpoints of the resolved service. If no filter is specified, all endpoints will be considered. An endpoint will be chosen arbitrarily from the filtered list for each request. For endpoint filter syntax and examples, see https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_ekm_connection#endpoint_filter GoogleKmsEkmConnection#endpoint_filter}
	EndpointFilter *string `field:"optional" json:"endpointFilter" yaml:"endpointFilter"`
}

