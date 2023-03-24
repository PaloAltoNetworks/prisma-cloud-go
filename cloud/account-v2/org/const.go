package org

const (
	singular = "org cloud account"
	plural   = "org cloud accounts"
)

var DeleteSuffix = []string{"cloud"}

var Suffix = []string{"cas", "v1"}

var ListSuffix = []string{"v1", "cloudAccounts"}

const (
	TypeAwsOrg   = "aws"
	TypeAzureOrg = "azure"
)
