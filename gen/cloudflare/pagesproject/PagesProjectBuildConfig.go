package pagesproject


type PagesProjectBuildConfig struct {
	// Command used to build project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#build_command PagesProject#build_command}
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// Output directory of the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#destination_dir PagesProject#destination_dir}
	DestinationDir *string `field:"optional" json:"destinationDir" yaml:"destinationDir"`
	// Directory to run the command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#root_dir PagesProject#root_dir}
	RootDir *string `field:"optional" json:"rootDir" yaml:"rootDir"`
	// The classifying tag for analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#web_analytics_tag PagesProject#web_analytics_tag}
	WebAnalyticsTag *string `field:"optional" json:"webAnalyticsTag" yaml:"webAnalyticsTag"`
	// The auth token for analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#web_analytics_token PagesProject#web_analytics_token}
	WebAnalyticsToken *string `field:"optional" json:"webAnalyticsToken" yaml:"webAnalyticsToken"`
}

