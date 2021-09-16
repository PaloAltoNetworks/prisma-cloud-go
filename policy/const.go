package policy

const (
	singular = "policy"
	plural   = "policies"
)

// Valid values for Policy.Rule.RuleType.
const (
	RuleTypeConfig     = "Config"
	RuleTypeAuditEvent = "AuditEvent"
	RuleTypeNetwork    = "Network"
	RuleTypeAnomaly    = "Anomaly"
)

// Valid values for Policy.PolicyType.
const (
	PolicyTypeConfig     = "config"
	PolicyTypeAuditEvent = "audit_event"
	PolicyTypeNetwork    = "network"
	PolicyTypeAnomaly    = "anomaly"
)

// Valid values for Policy.Rule.Severity.
const (
	SeverityLow    = "low"
	SeverityMedium = "medium"
	SeverityHigh   = "high"
)

var Suffix = []string{"policy"}
