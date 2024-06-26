// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

// Application gateway resource.
type ApplicationGateway_STATUS_ApplicationGateway_SubResourceEmbedded_ARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the application gateway, if configured.
	Identity *ManagedServiceIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the application gateway.
	Properties *ApplicationGatewayPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

// Properties of the application gateway.
type ApplicationGatewayPropertiesFormat_STATUS_ARM struct {
	// AuthenticationCertificates: Authentication certificates of the application gateway resource. For default limits, see
	// [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	AuthenticationCertificates []ApplicationGatewayAuthenticationCertificate_STATUS_ARM `json:"authenticationCertificates,omitempty"`

	// AutoscaleConfiguration: Autoscale Configuration.
	AutoscaleConfiguration *ApplicationGatewayAutoscaleConfiguration_STATUS_ARM `json:"autoscaleConfiguration,omitempty"`

	// BackendAddressPools: Backend address pool of the application gateway resource. For default limits, see [Application
	// Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendAddressPools []ApplicationGatewayBackendAddressPool_STATUS_ARM `json:"backendAddressPools,omitempty"`

	// BackendHttpSettingsCollection: Backend http settings of the application gateway resource. For default limits, see
	// [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettings_STATUS_ARM `json:"backendHttpSettingsCollection,omitempty"`

	// BackendSettingsCollection: Backend settings of the application gateway resource. For default limits, see [Application
	// Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendSettingsCollection []ApplicationGatewayBackendSettings_STATUS_ARM `json:"backendSettingsCollection,omitempty"`

	// CustomErrorConfigurations: Custom error configurations of the application gateway resource.
	CustomErrorConfigurations []ApplicationGatewayCustomError_STATUS_ARM `json:"customErrorConfigurations,omitempty"`

	// EnableFips: Whether FIPS is enabled on the application gateway resource.
	EnableFips *bool `json:"enableFips,omitempty"`

	// EnableHttp2: Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 *bool `json:"enableHttp2,omitempty"`

	// FirewallPolicy: Reference to the FirewallPolicy resource.
	FirewallPolicy *ApplicationGatewaySubResource_STATUS_ARM `json:"firewallPolicy,omitempty"`

	// ForceFirewallPolicyAssociation: If true, associates a firewall policy with an application gateway regardless whether the
	// policy differs from the WAF Config.
	ForceFirewallPolicyAssociation *bool `json:"forceFirewallPolicyAssociation,omitempty"`

	// FrontendIPConfigurations: Frontend IP addresses of the application gateway resource. For default limits, see
	// [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendIPConfigurations []ApplicationGatewayFrontendIPConfiguration_STATUS_ARM `json:"frontendIPConfigurations,omitempty"`

	// FrontendPorts: Frontend ports of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendPorts []ApplicationGatewayFrontendPort_STATUS_ARM `json:"frontendPorts,omitempty"`

	// GatewayIPConfigurations: Subnets of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	GatewayIPConfigurations []ApplicationGatewayIPConfiguration_STATUS_ApplicationGateway_SubResourceEmbedded_ARM `json:"gatewayIPConfigurations,omitempty"`

	// GlobalConfiguration: Global Configuration.
	GlobalConfiguration *ApplicationGatewayGlobalConfiguration_STATUS_ARM `json:"globalConfiguration,omitempty"`

	// HttpListeners: Http listeners of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	HttpListeners []ApplicationGatewayHttpListener_STATUS_ARM `json:"httpListeners,omitempty"`

	// Listeners: Listeners of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	Listeners []ApplicationGatewayListener_STATUS_ARM `json:"listeners,omitempty"`

	// LoadDistributionPolicies: Load distribution policies of the application gateway resource.
	LoadDistributionPolicies []ApplicationGatewayLoadDistributionPolicy_STATUS_ARM `json:"loadDistributionPolicies,omitempty"`

	// OperationalState: Operational state of the application gateway resource.
	OperationalState *ApplicationGatewayPropertiesFormat_OperationalState_STATUS `json:"operationalState,omitempty"`

	// PrivateEndpointConnections: Private Endpoint connections on application gateway.
	PrivateEndpointConnections []ApplicationGatewayPrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// PrivateLinkConfigurations: PrivateLink configurations on application gateway.
	PrivateLinkConfigurations []ApplicationGatewayPrivateLinkConfiguration_STATUS_ARM `json:"privateLinkConfigurations,omitempty"`

	// Probes: Probes of the application gateway resource.
	Probes []ApplicationGatewayProbe_STATUS_ARM `json:"probes,omitempty"`

	// ProvisioningState: The provisioning state of the application gateway resource.
	ProvisioningState *ApplicationGatewayProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RedirectConfigurations: Redirect configurations of the application gateway resource. For default limits, see
	// [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	RedirectConfigurations []ApplicationGatewayRedirectConfiguration_STATUS_ARM `json:"redirectConfigurations,omitempty"`

	// RequestRoutingRules: Request routing rules of the application gateway resource.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRule_STATUS_ARM `json:"requestRoutingRules,omitempty"`

	// ResourceGuid: The resource GUID property of the application gateway resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// RewriteRuleSets: Rewrite rules for the application gateway resource.
	RewriteRuleSets []ApplicationGatewayRewriteRuleSet_STATUS_ARM `json:"rewriteRuleSets,omitempty"`

	// RoutingRules: Routing rules of the application gateway resource.
	RoutingRules []ApplicationGatewayRoutingRule_STATUS_ARM `json:"routingRules,omitempty"`

	// Sku: SKU of the application gateway resource.
	Sku *ApplicationGatewaySku_STATUS_ARM `json:"sku,omitempty"`

	// SslCertificates: SSL certificates of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslCertificates []ApplicationGatewaySslCertificate_STATUS_ARM `json:"sslCertificates,omitempty"`

	// SslPolicy: SSL policy of the application gateway resource.
	SslPolicy *ApplicationGatewaySslPolicy_STATUS_ARM `json:"sslPolicy,omitempty"`

	// SslProfiles: SSL profiles of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslProfiles []ApplicationGatewaySslProfile_STATUS_ARM `json:"sslProfiles,omitempty"`

	// TrustedClientCertificates: Trusted client certificates of the application gateway resource. For default limits, see
	// [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedClientCertificates []ApplicationGatewayTrustedClientCertificate_STATUS_ARM `json:"trustedClientCertificates,omitempty"`

	// TrustedRootCertificates: Trusted Root certificates of the application gateway resource. For default limits, see
	// [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedRootCertificates []ApplicationGatewayTrustedRootCertificate_STATUS_ARM `json:"trustedRootCertificates,omitempty"`

	// UrlPathMaps: URL path map of the application gateway resource. For default limits, see [Application Gateway
	// limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	UrlPathMaps []ApplicationGatewayUrlPathMap_STATUS_ARM `json:"urlPathMaps,omitempty"`

	// WebApplicationFirewallConfiguration: Web application firewall configuration.
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfiguration_STATUS_ARM `json:"webApplicationFirewallConfiguration,omitempty"`
}

// Identity for the resource.
type ManagedServiceIdentity_STATUS_ARM struct {
	// PrincipalId: The principal id of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant id of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual
	// machine.
	Type *ManagedServiceIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user identities associated with resource. The user identity dictionary key
	// references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]ManagedServiceIdentity_UserAssignedIdentities_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Authentication certificates of an application gateway.
type ApplicationGatewayAuthenticationCertificate_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Application Gateway autoscale configuration.
type ApplicationGatewayAutoscaleConfiguration_STATUS_ARM struct {
	// MaxCapacity: Upper bound on number of Application Gateway capacity.
	MaxCapacity *int `json:"maxCapacity,omitempty"`

	// MinCapacity: Lower bound on number of Application Gateway capacity.
	MinCapacity *int `json:"minCapacity,omitempty"`
}

// Backend Address Pool of an application gateway.
type ApplicationGatewayBackendAddressPool_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Backend address pool settings of an application gateway.
type ApplicationGatewayBackendHttpSettings_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Backend address pool settings of an application gateway.
type ApplicationGatewayBackendSettings_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Customer error of an application gateway.
type ApplicationGatewayCustomError_STATUS_ARM struct {
	// CustomErrorPageUrl: Error page URL of the application gateway customer error.
	CustomErrorPageUrl *string `json:"customErrorPageUrl,omitempty"`

	// StatusCode: Status code of the application gateway customer error.
	StatusCode *ApplicationGatewayCustomError_StatusCode_STATUS `json:"statusCode,omitempty"`
}

// Frontend IP configuration of an application gateway.
type ApplicationGatewayFrontendIPConfiguration_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Frontend port of an application gateway.
type ApplicationGatewayFrontendPort_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Application Gateway global configuration.
type ApplicationGatewayGlobalConfiguration_STATUS_ARM struct {
	// EnableRequestBuffering: Enable request buffering.
	EnableRequestBuffering *bool `json:"enableRequestBuffering,omitempty"`

	// EnableResponseBuffering: Enable response buffering.
	EnableResponseBuffering *bool `json:"enableResponseBuffering,omitempty"`
}

// Http listener of an application gateway.
type ApplicationGatewayHttpListener_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// IP configuration of an application gateway. Currently 1 public and 1 private IP configuration is allowed.
type ApplicationGatewayIPConfiguration_STATUS_ApplicationGateway_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Listener of an application gateway.
type ApplicationGatewayListener_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Load Distribution Policy of an application gateway.
type ApplicationGatewayLoadDistributionPolicy_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Private Endpoint connection on an application gateway.
type ApplicationGatewayPrivateEndpointConnection_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Private Link Configuration on an application gateway.
type ApplicationGatewayPrivateLinkConfiguration_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Probe of the application gateway.
type ApplicationGatewayProbe_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Redirect configuration of an application gateway.
type ApplicationGatewayRedirectConfiguration_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Request routing rule of an application gateway.
type ApplicationGatewayRequestRoutingRule_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Rewrite rule set of an application gateway.
type ApplicationGatewayRewriteRuleSet_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Routing rule of an application gateway.
type ApplicationGatewayRoutingRule_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// SKU of an application gateway.
type ApplicationGatewaySku_STATUS_ARM struct {
	// Capacity: Capacity (instance count) of an application gateway.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of an application gateway SKU.
	Name *ApplicationGatewaySku_Name_STATUS `json:"name,omitempty"`

	// Tier: Tier of an application gateway.
	Tier *ApplicationGatewaySku_Tier_STATUS `json:"tier,omitempty"`
}

// SSL certificates of an application gateway.
type ApplicationGatewaySslCertificate_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Application Gateway Ssl policy.
type ApplicationGatewaySslPolicy_STATUS_ARM struct {
	// CipherSuites: Ssl cipher suites to be enabled in the specified order to application gateway.
	CipherSuites []CipherSuitesEnum_STATUS `json:"cipherSuites,omitempty"`

	// DisabledSslProtocols: Ssl protocols to be disabled on application gateway.
	DisabledSslProtocols []ProtocolsEnum_STATUS `json:"disabledSslProtocols,omitempty"`

	// MinProtocolVersion: Minimum version of Ssl protocol to be supported on application gateway.
	MinProtocolVersion *ProtocolsEnum_STATUS `json:"minProtocolVersion,omitempty"`

	// PolicyName: Name of Ssl predefined policy.
	PolicyName *PolicyNameEnum_STATUS `json:"policyName,omitempty"`

	// PolicyType: Type of Ssl Policy.
	PolicyType *ApplicationGatewaySslPolicy_PolicyType_STATUS `json:"policyType,omitempty"`
}

// SSL profile of an application gateway.
type ApplicationGatewaySslProfile_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Reference to another subresource.
type ApplicationGatewaySubResource_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Trusted client certificates of an application gateway.
type ApplicationGatewayTrustedClientCertificate_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Trusted Root certificates of an application gateway.
type ApplicationGatewayTrustedRootCertificate_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// UrlPathMaps give a url path to the backend mapping information for PathBasedRouting.
type ApplicationGatewayUrlPathMap_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Application gateway web application firewall configuration.
type ApplicationGatewayWebApplicationFirewallConfiguration_STATUS_ARM struct {
	// DisabledRuleGroups: The disabled rule groups.
	DisabledRuleGroups []ApplicationGatewayFirewallDisabledRuleGroup_STATUS_ARM `json:"disabledRuleGroups,omitempty"`

	// Enabled: Whether the web application firewall is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`

	// Exclusions: The exclusion list.
	Exclusions []ApplicationGatewayFirewallExclusion_STATUS_ARM `json:"exclusions,omitempty"`

	// FileUploadLimitInMb: Maximum file upload size in Mb for WAF.
	FileUploadLimitInMb *int `json:"fileUploadLimitInMb,omitempty"`

	// FirewallMode: Web application firewall mode.
	FirewallMode *ApplicationGatewayWebApplicationFirewallConfiguration_FirewallMode_STATUS `json:"firewallMode,omitempty"`

	// MaxRequestBodySize: Maximum request body size for WAF.
	MaxRequestBodySize *int `json:"maxRequestBodySize,omitempty"`

	// MaxRequestBodySizeInKb: Maximum request body size in Kb for WAF.
	MaxRequestBodySizeInKb *int `json:"maxRequestBodySizeInKb,omitempty"`

	// RequestBodyCheck: Whether allow WAF to check request Body.
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty"`

	// RuleSetType: The type of the web application firewall rule set. Possible values are: 'OWASP'.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: The version of the rule set type.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

type ManagedServiceIdentity_Type_STATUS string

const (
	ManagedServiceIdentity_Type_STATUS_None                       = ManagedServiceIdentity_Type_STATUS("None")
	ManagedServiceIdentity_Type_STATUS_SystemAssigned             = ManagedServiceIdentity_Type_STATUS("SystemAssigned")
	ManagedServiceIdentity_Type_STATUS_SystemAssignedUserAssigned = ManagedServiceIdentity_Type_STATUS("SystemAssigned, UserAssigned")
	ManagedServiceIdentity_Type_STATUS_UserAssigned               = ManagedServiceIdentity_Type_STATUS("UserAssigned")
)

type ManagedServiceIdentity_UserAssignedIdentities_STATUS_ARM struct {
	// ClientId: The client id of user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal id of user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`
}

// Allows to disable rules within a rule group or an entire rule group.
type ApplicationGatewayFirewallDisabledRuleGroup_STATUS_ARM struct {
	// RuleGroupName: The name of the rule group that will be disabled.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: The list of rules that will be disabled. If null, all rules of the rule group will be disabled.
	Rules []int `json:"rules,omitempty"`
}

// Allow to exclude some variable satisfy the condition for the WAF check.
type ApplicationGatewayFirewallExclusion_STATUS_ARM struct {
	// MatchVariable: The variable to be excluded.
	MatchVariable *string `json:"matchVariable,omitempty"`

	// Selector: When matchVariable is a collection, operator used to specify which elements in the collection this exclusion
	// applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: When matchVariable is a collection, operate on the selector to specify which elements in the
	// collection this exclusion applies to.
	SelectorMatchOperator *string `json:"selectorMatchOperator,omitempty"`
}
