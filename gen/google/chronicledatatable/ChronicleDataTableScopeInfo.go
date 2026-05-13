package chronicledatatable


type ChronicleDataTableScopeInfo struct {
	// Contains the list of scope names of the data table.
	//
	// If the list is empty,
	// the data table is treated as unscoped. The scope names should be
	// full resource names and should be of the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope_name}"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table#data_access_scopes ChronicleDataTable#data_access_scopes}
	DataAccessScopes *[]*string `field:"required" json:"dataAccessScopes" yaml:"dataAccessScopes"`
}

