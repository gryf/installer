//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

// DomainTopicsClientCreateOrUpdateResponse contains the response from method DomainTopicsClient.CreateOrUpdate.
type DomainTopicsClientCreateOrUpdateResponse struct {
	DomainTopic
}

// DomainTopicsClientDeleteResponse contains the response from method DomainTopicsClient.Delete.
type DomainTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainTopicsClientGetResponse contains the response from method DomainTopicsClient.Get.
type DomainTopicsClientGetResponse struct {
	DomainTopic
}

// DomainTopicsClientListByDomainResponse contains the response from method DomainTopicsClient.ListByDomain.
type DomainTopicsClientListByDomainResponse struct {
	DomainTopicsListResult
}

// DomainsClientCreateOrUpdateResponse contains the response from method DomainsClient.CreateOrUpdate.
type DomainsClientCreateOrUpdateResponse struct {
	Domain
}

// DomainsClientDeleteResponse contains the response from method DomainsClient.Delete.
type DomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainsClientGetResponse contains the response from method DomainsClient.Get.
type DomainsClientGetResponse struct {
	Domain
}

// DomainsClientListByResourceGroupResponse contains the response from method DomainsClient.ListByResourceGroup.
type DomainsClientListByResourceGroupResponse struct {
	DomainsListResult
}

// DomainsClientListBySubscriptionResponse contains the response from method DomainsClient.ListBySubscription.
type DomainsClientListBySubscriptionResponse struct {
	DomainsListResult
}

// DomainsClientListSharedAccessKeysResponse contains the response from method DomainsClient.ListSharedAccessKeys.
type DomainsClientListSharedAccessKeysResponse struct {
	DomainSharedAccessKeys
}

// DomainsClientRegenerateKeyResponse contains the response from method DomainsClient.RegenerateKey.
type DomainsClientRegenerateKeyResponse struct {
	DomainSharedAccessKeys
}

// DomainsClientUpdateResponse contains the response from method DomainsClient.Update.
type DomainsClientUpdateResponse struct {
	Domain
}

// EventSubscriptionsClientCreateOrUpdateResponse contains the response from method EventSubscriptionsClient.CreateOrUpdate.
type EventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// EventSubscriptionsClientDeleteResponse contains the response from method EventSubscriptionsClient.Delete.
type EventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// EventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method EventSubscriptionsClient.GetDeliveryAttributes.
type EventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// EventSubscriptionsClientGetFullURLResponse contains the response from method EventSubscriptionsClient.GetFullURL.
type EventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// EventSubscriptionsClientGetResponse contains the response from method EventSubscriptionsClient.Get.
type EventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// EventSubscriptionsClientListByDomainTopicResponse contains the response from method EventSubscriptionsClient.ListByDomainTopic.
type EventSubscriptionsClientListByDomainTopicResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListByResourceResponse contains the response from method EventSubscriptionsClient.ListByResource.
type EventSubscriptionsClientListByResourceResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListGlobalByResourceGroupForTopicType.
type EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalByResourceGroupResponse contains the response from method EventSubscriptionsClient.ListGlobalByResourceGroup.
type EventSubscriptionsClientListGlobalByResourceGroupResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalBySubscriptionForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListGlobalBySubscriptionForTopicType.
type EventSubscriptionsClientListGlobalBySubscriptionForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalBySubscriptionResponse contains the response from method EventSubscriptionsClient.ListGlobalBySubscription.
type EventSubscriptionsClientListGlobalBySubscriptionResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalByResourceGroupForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListRegionalByResourceGroupForTopicType.
type EventSubscriptionsClientListRegionalByResourceGroupForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalByResourceGroupResponse contains the response from method EventSubscriptionsClient.ListRegionalByResourceGroup.
type EventSubscriptionsClientListRegionalByResourceGroupResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalBySubscriptionForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListRegionalBySubscriptionForTopicType.
type EventSubscriptionsClientListRegionalBySubscriptionForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalBySubscriptionResponse contains the response from method EventSubscriptionsClient.ListRegionalBySubscription.
type EventSubscriptionsClientListRegionalBySubscriptionResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientUpdateResponse contains the response from method EventSubscriptionsClient.Update.
type EventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// ExtensionTopicsClientGetResponse contains the response from method ExtensionTopicsClient.Get.
type ExtensionTopicsClientGetResponse struct {
	ExtensionTopic
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByResourceResponse contains the response from method PrivateEndpointConnectionsClient.ListByResource.
type PrivateEndpointConnectionsClientListByResourceResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByResourceResponse contains the response from method PrivateLinkResourcesClient.ListByResource.
type PrivateLinkResourcesClientListByResourceResponse struct {
	PrivateLinkResourcesListResult
}

// SystemTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method SystemTopicEventSubscriptionsClient.CreateOrUpdate.
type SystemTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// SystemTopicEventSubscriptionsClientDeleteResponse contains the response from method SystemTopicEventSubscriptionsClient.Delete.
type SystemTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method SystemTopicEventSubscriptionsClient.GetDeliveryAttributes.
type SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// SystemTopicEventSubscriptionsClientGetFullURLResponse contains the response from method SystemTopicEventSubscriptionsClient.GetFullURL.
type SystemTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// SystemTopicEventSubscriptionsClientGetResponse contains the response from method SystemTopicEventSubscriptionsClient.Get.
type SystemTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// SystemTopicEventSubscriptionsClientListBySystemTopicResponse contains the response from method SystemTopicEventSubscriptionsClient.ListBySystemTopic.
type SystemTopicEventSubscriptionsClientListBySystemTopicResponse struct {
	EventSubscriptionsListResult
}

// SystemTopicEventSubscriptionsClientUpdateResponse contains the response from method SystemTopicEventSubscriptionsClient.Update.
type SystemTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// SystemTopicsClientCreateOrUpdateResponse contains the response from method SystemTopicsClient.CreateOrUpdate.
type SystemTopicsClientCreateOrUpdateResponse struct {
	SystemTopic
}

// SystemTopicsClientDeleteResponse contains the response from method SystemTopicsClient.Delete.
type SystemTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// SystemTopicsClientGetResponse contains the response from method SystemTopicsClient.Get.
type SystemTopicsClientGetResponse struct {
	SystemTopic
}

// SystemTopicsClientListByResourceGroupResponse contains the response from method SystemTopicsClient.ListByResourceGroup.
type SystemTopicsClientListByResourceGroupResponse struct {
	SystemTopicsListResult
}

// SystemTopicsClientListBySubscriptionResponse contains the response from method SystemTopicsClient.ListBySubscription.
type SystemTopicsClientListBySubscriptionResponse struct {
	SystemTopicsListResult
}

// SystemTopicsClientUpdateResponse contains the response from method SystemTopicsClient.Update.
type SystemTopicsClientUpdateResponse struct {
	SystemTopic
}

// TopicTypesClientGetResponse contains the response from method TopicTypesClient.Get.
type TopicTypesClientGetResponse struct {
	TopicTypeInfo
}

// TopicTypesClientListEventTypesResponse contains the response from method TopicTypesClient.ListEventTypes.
type TopicTypesClientListEventTypesResponse struct {
	EventTypesListResult
}

// TopicTypesClientListResponse contains the response from method TopicTypesClient.List.
type TopicTypesClientListResponse struct {
	TopicTypesListResult
}

// TopicsClientCreateOrUpdateResponse contains the response from method TopicsClient.CreateOrUpdate.
type TopicsClientCreateOrUpdateResponse struct {
	Topic
}

// TopicsClientDeleteResponse contains the response from method TopicsClient.Delete.
type TopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// TopicsClientGetResponse contains the response from method TopicsClient.Get.
type TopicsClientGetResponse struct {
	Topic
}

// TopicsClientListByResourceGroupResponse contains the response from method TopicsClient.ListByResourceGroup.
type TopicsClientListByResourceGroupResponse struct {
	TopicsListResult
}

// TopicsClientListBySubscriptionResponse contains the response from method TopicsClient.ListBySubscription.
type TopicsClientListBySubscriptionResponse struct {
	TopicsListResult
}

// TopicsClientListEventTypesResponse contains the response from method TopicsClient.ListEventTypes.
type TopicsClientListEventTypesResponse struct {
	EventTypesListResult
}

// TopicsClientListSharedAccessKeysResponse contains the response from method TopicsClient.ListSharedAccessKeys.
type TopicsClientListSharedAccessKeysResponse struct {
	TopicSharedAccessKeys
}

// TopicsClientRegenerateKeyResponse contains the response from method TopicsClient.RegenerateKey.
type TopicsClientRegenerateKeyResponse struct {
	TopicSharedAccessKeys
}

// TopicsClientUpdateResponse contains the response from method TopicsClient.Update.
type TopicsClientUpdateResponse struct {
	Topic
}
