package googlednsmanagedzone


type GoogleDnsManagedZoneServiceDirectoryConfigNamespace struct {
	// The fully qualified or partial URL of the service directory namespace that should be associated with the zone.
	//
	// This should be formatted like
	// 'https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace_id}'
	// or simply 'projects/{project}/locations/{location}/namespaces/{namespace_id}'
	// Ignored for 'public' visibility zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#namespace_url GoogleDnsManagedZone#namespace_url}
	NamespaceUrl *string `field:"required" json:"namespaceUrl" yaml:"namespaceUrl"`
}

