package provider


type IncidentProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs#alias IncidentProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// API key for incident.io (https://app.incident.io/settings/api-keys). Sourced from the `INCIDENT_API_KEY` environment variable, if set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs#api_key IncidentProvider#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// URL of the incident.io API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs#endpoint IncidentProvider#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

