package alertsource


type AlertSourceJiraOptions struct {
	// Which projects in Jira should this alert source watch for new issues?
	//
	// IDs can either be IDs of the projects in Jira, or ID of catalog entries in the 'Jira Project' catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_source#project_ids AlertSource#project_ids}
	ProjectIds *[]*string `field:"optional" json:"projectIds" yaml:"projectIds"`
}

