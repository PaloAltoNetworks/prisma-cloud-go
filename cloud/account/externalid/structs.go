package externalid

type ExternalIdReq struct {
	AccountId    string   `json:"accountId"`
	AccountType  string   `json:"accountType"`
	AwsPartition string   `json:"awsPartition"`
	Features     []string `json:"features"`
}

type ExternalId struct {
	AccountId                         string   `json:"accountId"`
	AccountType                       string   `json:"accountType"`
	AwsPartition                      string   `json:"awsPartition"`
	Features                          []string `json:"features"`
	ExternalId                        string   `json:"externalId"`
	CreateStackLinkWithS3PresignedUrl string   `json:"createStackLinkWithS3PresignedUrl"`
	EventBridgeRuleNamePrefix         string   `json:"eventBridgeRuleNamePrefix"`
}
