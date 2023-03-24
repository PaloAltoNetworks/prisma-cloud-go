package org

type AccountAndCredentials struct {
	Account OrgAccount `json:"cloudAccount"`
}

type NameTypeId struct {
	Name      string `json:"name"`
	CloudType string `json:"cloudType"`
	AccountId string `json:"id"`
}

type CloudAccountResp struct {
	AccountId               string `json:"accountId"`
	Name                    string `json:"name"`
	CloudType               string `json:"cloudType"`
	Enabled                 bool   `json:"enabled"`
	ParentId                string `json:"parentId"`
	AccountType             string `json:"accountType"`
	Deleted                 bool   `json:"deleted"`
	ProtectionMode          string `json:"protectionMode"`
	CustomerName            string `json:"customerName"`
	CreatedEpochMillis      int    `json:"createdEpochMillis"`
	LastModifiedEpochMillis int    `json:"lastModifiedEpochMillis"`
	LastModifiedBy          string `json:"lastModifiedBy"`
}

type AccountResponse struct {
	CloudAccountResp        CloudAccountResp     `json:"cloudAccount"`
	DeploymentType          string               `json:"deploymentType"`
	TenantId                string               `json:"tenantId"`
	ServicePrincipalId      string               `json:"servicePrincipalId"`
	ClientId                string               `json:"clientId"`
	Key                     string               `json:"key"`
	MonitorFlowLogs         bool                 `json:"monitorFlowLogs"`
	EnvironmentType         string               `json:"environmentType"`
	Features                []Features           `json:"features"`
	RoleArn                 string               `json:"roleArn"`
	ExternalId              string               `json:"externalId"`
	HasMemberRole           bool                 `json:"hasMemberRole"`
	TemplateUrl             string               `json:"templateUrl"`
	GroupIds                []string             `json:"groupIds"`
	HierarchySelection      []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId   string               `json:"defaultAccountGroupId"`
	DefaultAccountGroupName string               `json:"defaultAccountGroupName"`
	MemberRoleName          string               `json:"memberRoleName"`
	MemberExternalId        string               `json:"memberExternalId"`
	MemberTemplateUrl       string               `json:"memberTemplateUrl"`
}

type AwsOrg struct {
	AccountId             string               `json:"accountId"`
	Enabled               bool                 `json:"enabled"`
	DefaultAccountGroupId string               `json:"defaultAccountGroupId"`
	Name                  string               `json:"name"`
	RoleArn               string               `json:"roleArn"`
	AccountType           string               `json:"accountType"`
	GroupIds              []string             `json:"groupIds"`
	HierarchySelection    []HierarchySelection `json:"hierarchySelection"`
	Features              []Features           `json:"features"`
}

type AwsOrgV2 struct {
	CloudAccountResp          CloudAccountResp     `json:"cloudAccount"`
	Name                      string               `json:"name"`
	RoleArn                   string               `json:"roleArn"`
	ExternalId                string               `json:"externalId"`
	HasMemberRole             bool                 `json:"hasMemberRole"`
	TemplateUrl               string               `json:"templateUrl"`
	DeploymentType            string               `json:"deploymentType"`
	GroupIds                  []string             `json:"groupIds"`
	EventbridgeRuleNamePrefix string               `json:"eventbridgeRuleNamePrefix"`
	DefaultAccountGroupId     string               `json:"defaultAccountGroupId"`
	Features                  []Features           `json:"features"`
	HierarchySelection        []HierarchySelection `json:"hierarchySelection"`
}

type OrgAccount struct {
	Name                  string   `json:"name"`
	CloudType             string   `json:"cloudType"`
	AccountType           string   `json:"accountType"`
	Enabled               bool     `json:"enabled"`
	LastModifiedTs        int      `json:"lastModifiedTs"`
	LastModifiedBy        string   `json:"lastModifiedBy"`
	StorageScanEnabled    bool     `json:"storageScanEnabled"`
	ProtectionMode        string   `json:"protectionMode"`
	IngestionMode         int      `json:"ingestionMode"`
	GroupIds              []string `json:"groupIds"`
	Groups                []Group  `json:"groups"`
	Status                string   `json:"status"`
	NumberOfChildAccounts int      `json:"numberOfChildAccounts"`
	AccountId             string   `json:"accountId,omitempty"`
	AddedOn               int      `json:"addedOn"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type HierarchySelection struct {
	ResourceId    string `json:"resourceId"`
	DisplayName   string `json:"displayName"`
	NodeType      string `json:"nodeType"`
	SelectionType string `json:"selectionType"`
}

type Features struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type Features1 struct {
	Name  string `json:"featureName"`
	State string `json:"featureState"`
}

//AZUREV2ORG

type AzureOrgV2 struct {
	Name                  string               `json:"name"`
	GroupIds              []string             `json:"groupIds"`
	CloudAccountResp      CloudAccountResp     `json:"cloudAccount"`
	ClientId              string               `json:"clientId"`
	EnvironmentType       string               `json:"environmentType"`
	Key                   string               `json:"key"`
	MonitorFlowLogs       bool                 `json:"monitorFlowLogs"`
	ServicePrincipalId    string               `json:"servicePrincipalId"`
	TemplateUrl           string               `json:"templateUrl"`
	HierarchySelection    []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId string               `json:"defaultAccountGroupId"`
	RootSyncEnabled       bool                 `json:"rootSyncEnabled"`
	TenantId              string               `json:"tenantId"`
	Features              []Features           `json:"features"`
}

type AzureOrg struct {
	CloudAccountResp      CloudAccountResp     `json:"cloudAccount"`
	ClientId              string               `json:"clientId"`
	RootSyncEnabled       bool                 `json:"rootSyncEnabled"`
	GroupIds              []string             `json:"groupIds"`
	EnvironmentType       string               `json:"environmentType"`
	Key                   string               `json:"key"`
	MonitorFlowLogs       bool                 `json:"monitorFlowLogs"`
	ServicePrincipalId    string               `json:"servicePrincipalId"`
	TemplateUrl           string               `json:"templateUrl"`
	TenantId              string               `json:"tenantId"`
	Features              []Features           `json:"features"`
	HierarchySelection    []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId string               `json:"defaultAccountGroupId"`
}
