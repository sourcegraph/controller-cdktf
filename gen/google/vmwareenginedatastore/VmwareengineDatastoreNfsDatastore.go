package vmwareenginedatastore


type VmwareengineDatastoreNfsDatastore struct {
	// google_file_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_datastore#google_file_service VmwareengineDatastore#google_file_service}
	GoogleFileService *VmwareengineDatastoreNfsDatastoreGoogleFileService `field:"optional" json:"googleFileService" yaml:"googleFileService"`
	// third_party_file_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_datastore#third_party_file_service VmwareengineDatastore#third_party_file_service}
	ThirdPartyFileService *VmwareengineDatastoreNfsDatastoreThirdPartyFileService `field:"optional" json:"thirdPartyFileService" yaml:"thirdPartyFileService"`
}

