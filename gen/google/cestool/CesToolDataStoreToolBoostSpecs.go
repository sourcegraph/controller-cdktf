package cestool


type CesToolDataStoreToolBoostSpecs struct {
	// The Data Store where the boosting configuration is applied. Full resource name of DataStore, such as projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#data_stores CesTool#data_stores}
	DataStores *[]*string `field:"required" json:"dataStores" yaml:"dataStores"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#spec CesTool#spec}
	Spec interface{} `field:"required" json:"spec" yaml:"spec"`
}

