package googledialogflowcxagent


type GoogleDialogflowCxAgentPersonalizationSettings struct {
	// Default end user metadata, used when processing DetectIntent requests.
	//
	// Recommended to be filled as a template instead of hard-coded value, for example { "age": "$session.params.age" }.
	// The data will be merged with the [QueryParameters.end_user_metadata](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/QueryParameters#FIELDS.end_user_metadata)
	// in [DetectIntentRequest.query_params](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.sessions/detectIntent#body.request_body.FIELDS.query_params) during query processing.
	//
	// This field uses JSON data as a string. The value provided must be a valid JSON representation documented in [Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#default_end_user_metadata GoogleDialogflowCxAgent#default_end_user_metadata}
	DefaultEndUserMetadata *string `field:"optional" json:"defaultEndUserMetadata" yaml:"defaultEndUserMetadata"`
}

