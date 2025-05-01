package googlechroniclewatchlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleChronicleWatchlistConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters (https://google.aip.dev/148).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#display_name GoogleChronicleWatchlist#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// entity_population_mechanism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#entity_population_mechanism GoogleChronicleWatchlist#entity_population_mechanism}
	EntityPopulationMechanism *GoogleChronicleWatchlistEntityPopulationMechanism `field:"required" json:"entityPopulationMechanism" yaml:"entityPopulationMechanism"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#instance GoogleChronicleWatchlist#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#location GoogleChronicleWatchlist#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional. Description of the watchlist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#description GoogleChronicleWatchlist#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#id GoogleChronicleWatchlist#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#multiplying_factor GoogleChronicleWatchlist#multiplying_factor}
	MultiplyingFactor *float64 `field:"optional" json:"multiplyingFactor" yaml:"multiplyingFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#project GoogleChronicleWatchlist#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#timeouts GoogleChronicleWatchlist#timeouts}
	Timeouts *GoogleChronicleWatchlistTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// The ID to use for the watchlist,
	// which will become the final component of the watchlist's resource name.
	// This value should be 4-63 characters, and valid characters
	// are /a-z-/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#watchlist_id GoogleChronicleWatchlist#watchlist_id}
	WatchlistId *string `field:"optional" json:"watchlistId" yaml:"watchlistId"`
	// watchlist_user_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_watchlist#watchlist_user_preferences GoogleChronicleWatchlist#watchlist_user_preferences}
	WatchlistUserPreferences *GoogleChronicleWatchlistWatchlistUserPreferences `field:"optional" json:"watchlistUserPreferences" yaml:"watchlistUserPreferences"`
}

