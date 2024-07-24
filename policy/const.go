package policy

const (
	singular = "policy"
	plural   = "policies"
)

// Valid values for Policy.Rule.RuleType.
const (
	RuleTypeConfig        = "Config"
	RuleTypeAuditEvent    = "AuditEvent"
	RuleTypeNetwork       = "Network"
	RuleTypeIAM           = "IAM"
	RuleTypeAnomaly       = "Anomaly"
	RuleTypeData          = "DLP"
	RuleTypeNetworkConfig = "NetworkConfig"
)

// Valid values for Policy.PolicyType.
const (
	PolicyTypeConfig     = "config"
	PolicyTypeAuditEvent = "audit_event"
	PolicyTypeNetwork    = "network"
	PolicyTypeIAM        = "iam"
	PolicyTypeAnomaly    = "anomaly"
	PolicyTypeData       = "data"
)

// Valid values for Policy.Rule.Severity.
const (
	SeverityLow           = "low"
	SeverityMedium        = "medium"
	SeverityHigh          = "high"
	SeverityCritical      = "critical"
	SeverityInformational = "informational"
)

var Suffix = []string{"policy"}
var BridgecrewSuffix = []string{"bridgecrew/api/v1/policies"}
