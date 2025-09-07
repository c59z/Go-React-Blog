package config

// GitHub login configuration, see https://docs.github.com/en/developers/apps/building-oauth-apps/authorizing-oauth-apps
type GitHub struct {
	Enable       bool   `json:"enable" yaml:"enable"`               // Whether to enable GitHub login
	ClientID     string `json:"client_id" yaml:"client_id"`         // Client ID of the GitHub OAuth app
	ClientSecret string `json:"client_secret" yaml:"client_secret"` // Client Secret of the GitHub OAuth app
	RedirectURI  string `json:"redirect_uri" yaml:"redirect_uri"`   // Callback URI
}

// GitHubLoginURL generates the GitHub OAuth login URL
func (gh GitHub) GitHubLoginURL(state string) string {
	return "https://github.com/login/oauth/authorize?" +
		"client_id=" + gh.ClientID + "&" +
		"redirect_uri=" + gh.RedirectURI + "&" +
		"scope=user:email" + "&" +
		"state=" + state // Prevent CSRF attacks
}
