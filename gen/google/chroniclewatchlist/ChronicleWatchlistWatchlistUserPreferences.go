package chroniclewatchlist

type ChronicleWatchlistWatchlistUserPreferences struct {
	// Optional. Whether the watchlist is pinned on the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/chronicle_watchlist#pinned ChronicleWatchlist#pinned}
	Pinned any `field:"optional" json:"pinned" yaml:"pinned"`
}
