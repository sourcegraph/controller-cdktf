package googleapigeeenvironment


type GoogleApigeeEnvironmentClientIpResolutionConfigHeaderIndexAlgorithm struct {
	// The index of the ip in the header.
	//
	// Positive indices 0, 1, 2, 3 chooses indices from the left (first ips). Negative indices -1, -2, -3 chooses indices from the right (last ips).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_environment#ip_header_index GoogleApigeeEnvironment#ip_header_index}
	IpHeaderIndex *float64 `field:"required" json:"ipHeaderIndex" yaml:"ipHeaderIndex"`
	// The name of the header to extract the client ip from. We are currently only supporting the X-Forwarded-For header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_environment#ip_header_name GoogleApigeeEnvironment#ip_header_name}
	IpHeaderName *string `field:"required" json:"ipHeaderName" yaml:"ipHeaderName"`
}

