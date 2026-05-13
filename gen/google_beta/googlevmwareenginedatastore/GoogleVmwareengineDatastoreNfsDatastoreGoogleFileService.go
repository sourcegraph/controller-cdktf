package googlevmwareenginedatastore


type GoogleVmwareengineDatastoreNfsDatastoreGoogleFileService struct {
	// Google filestore instance resource name e.g. projects/my-project/locations/me-west1-b/instances/my-instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#filestore_instance GoogleVmwareengineDatastore#filestore_instance}
	FilestoreInstance *string `field:"optional" json:"filestoreInstance" yaml:"filestoreInstance"`
	// Google netapp volume resource name e.g. projects/my-project/locations/me-west1-b/volumes/my-volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_datastore#netapp_volume GoogleVmwareengineDatastore#netapp_volume}
	NetappVolume *string `field:"optional" json:"netappVolume" yaml:"netappVolume"`
}

