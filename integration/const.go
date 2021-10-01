package integration

const (
	singular       = "integration"
	plural         = "integrations"
	JiraAuthUrl    = "AuthUrl"
	JiraSecretkey  = "SecretKey"
	JiraOauthtoken = "OauthToken"
)

var Suffix = []string{"integration"}
var JiraUrlSuffix = []string{"oauth", "authenticate", "jira"}
var JiraTokenSuffix = []string{"oauth", "token", "jira"}
var JiraSecretKeySuffix = []string{"plugins", "servlet", "oauth", "authorize"}
