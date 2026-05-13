package googlediscoveryenginewidgetconfig


type GoogleDiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigs struct {
	// facet_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#facet_field GoogleDiscoveryEngineWidgetConfig#facet_field}
	FacetField interface{} `field:"optional" json:"facetField" yaml:"facetField"`
	// fields_ui_components_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#fields_ui_components_map GoogleDiscoveryEngineWidgetConfig#fields_ui_components_map}
	FieldsUiComponentsMap interface{} `field:"optional" json:"fieldsUiComponentsMap" yaml:"fieldsUiComponentsMap"`
	// The name of the data store.
	//
	// It should be data store resource name. Format:
	// 'projects/{project}/locations/{location}/collections/{collectionId}/dataStores/{dataStoreId}'.
	// For APIs under 'WidgetService', such as [WidgetService.LookUpWidgetConfig][],
	// the project number and location part is erased in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config#name GoogleDiscoveryEngineWidgetConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

