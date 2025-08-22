package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PrereleaseVersion represents the PrereleaseVersion schema from the OpenAPI specification
type PrereleaseVersion struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// AppStoreReviewDetailUpdateRequest represents the AppStoreReviewDetailUpdateRequest schema from the OpenAPI specification
type AppStoreReviewDetailUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppsResponse represents the AppsResponse schema from the OpenAPI specification
type AppsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []App `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// AppBetaTestersLinkagesRequest represents the AppBetaTestersLinkagesRequest schema from the OpenAPI specification
type AppBetaTestersLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// CertificatesResponse represents the CertificatesResponse schema from the OpenAPI specification
type CertificatesResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []Certificate `json:"data"`
	Links PagedDocumentLinks `json:"links"`
}

// App represents the App schema from the OpenAPI specification
type App struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// BundleIdCapabilityCreateRequest represents the BundleIdCapabilityCreateRequest schema from the OpenAPI specification
type BundleIdCapabilityCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPrice represents the AppPrice schema from the OpenAPI specification
type AppPrice struct {
	TypeField string `json:"type"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// BetaTesterAppsLinkagesRequest represents the BetaTesterAppsLinkagesRequest schema from the OpenAPI specification
type BetaTesterAppsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// UserUpdateRequest represents the UserUpdateRequest schema from the OpenAPI specification
type UserUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// PagingInformation represents the PagingInformation schema from the OpenAPI specification
type PagingInformation struct {
	Paging map[string]interface{} `json:"paging"`
}

// RoutingAppCoverageCreateRequest represents the RoutingAppCoverageCreateRequest schema from the OpenAPI specification
type RoutingAppCoverageCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaAppLocalizationResponse represents the BetaAppLocalizationResponse schema from the OpenAPI specification
type BetaAppLocalizationResponse struct {
	Links DocumentLinks `json:"links"`
	Data BetaAppLocalization `json:"data"`
	Included []App `json:"included,omitempty"`
}

// ProfileCreateRequest represents the ProfileCreateRequest schema from the OpenAPI specification
type ProfileCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BundleIdsResponse represents the BundleIdsResponse schema from the OpenAPI specification
type BundleIdsResponse struct {
	Data []BundleId `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaBuildLocalizationsResponse represents the BetaBuildLocalizationsResponse schema from the OpenAPI specification
type BetaBuildLocalizationsResponse struct {
	Data []BetaBuildLocalization `json:"data"`
	Included []Build `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaBuildLocalizationResponse represents the BetaBuildLocalizationResponse schema from the OpenAPI specification
type BetaBuildLocalizationResponse struct {
	Included []Build `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
	Data BetaBuildLocalization `json:"data"`
}

// BetaTesterBetaGroupsLinkagesResponse represents the BetaTesterBetaGroupsLinkagesResponse schema from the OpenAPI specification
type BetaTesterBetaGroupsLinkagesResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []map[string]interface{} `json:"data"`
}

// BetaTesterInvitationCreateRequest represents the BetaTesterInvitationCreateRequest schema from the OpenAPI specification
type BetaTesterInvitationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BuildResponse represents the BuildResponse schema from the OpenAPI specification
type BuildResponse struct {
	Data Build `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// BuildIcon represents the BuildIcon schema from the OpenAPI specification
type BuildIcon struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
}

// AppEncryptionDeclarationBuildsLinkagesRequest represents the AppEncryptionDeclarationBuildsLinkagesRequest schema from the OpenAPI specification
type AppEncryptionDeclarationBuildsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// Profile represents the Profile schema from the OpenAPI specification
type Profile struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// AppStoreReviewAttachmentsResponse represents the AppStoreReviewAttachmentsResponse schema from the OpenAPI specification
type AppStoreReviewAttachmentsResponse struct {
	Data []AppStoreReviewAttachment `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppPriceTierResponse represents the AppPriceTierResponse schema from the OpenAPI specification
type AppPriceTierResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppPriceTier `json:"data"`
	Included []AppPricePoint `json:"included,omitempty"`
}

// BetaAppReviewSubmissionsResponse represents the BetaAppReviewSubmissionsResponse schema from the OpenAPI specification
type BetaAppReviewSubmissionsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []BetaAppReviewSubmission `json:"data"`
	Included []Build `json:"included,omitempty"`
}

// AppInfo represents the AppInfo schema from the OpenAPI specification
type AppInfo struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// BetaTesterBuildsLinkagesResponse represents the BetaTesterBuildsLinkagesResponse schema from the OpenAPI specification
type BetaTesterBuildsLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppStoreVersionSubmission represents the AppStoreVersionSubmission schema from the OpenAPI specification
type AppStoreVersionSubmission struct {
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Id string `json:"id"`
}

// AppPreviewSetAppPreviewsLinkagesRequest represents the AppPreviewSetAppPreviewsLinkagesRequest schema from the OpenAPI specification
type AppPreviewSetAppPreviewsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// UserInvitationsResponse represents the UserInvitationsResponse schema from the OpenAPI specification
type UserInvitationsResponse struct {
	Data []UserInvitation `json:"data"`
	Included []App `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaAppReviewDetailResponse represents the BetaAppReviewDetailResponse schema from the OpenAPI specification
type BetaAppReviewDetailResponse struct {
	Data BetaAppReviewDetail `json:"data"`
	Included []App `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// DeviceCreateRequest represents the DeviceCreateRequest schema from the OpenAPI specification
type DeviceCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppInfoResponse represents the AppInfoResponse schema from the OpenAPI specification
type AppInfoResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppInfo `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// AppInfoLocalizationCreateRequest represents the AppInfoLocalizationCreateRequest schema from the OpenAPI specification
type AppInfoLocalizationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaLicenseAgreementsResponse represents the BetaLicenseAgreementsResponse schema from the OpenAPI specification
type BetaLicenseAgreementsResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []BetaLicenseAgreement `json:"data"`
	Included []App `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
}

// BetaAppLocalization represents the BetaAppLocalization schema from the OpenAPI specification
type BetaAppLocalization struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// AppPreviewUpdateRequest represents the AppPreviewUpdateRequest schema from the OpenAPI specification
type AppPreviewUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BundleIdUpdateRequest represents the BundleIdUpdateRequest schema from the OpenAPI specification
type BundleIdUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaGroupResponse represents the BetaGroupResponse schema from the OpenAPI specification
type BetaGroupResponse struct {
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
	Data BetaGroup `json:"data"`
}

// IdfaDeclaration represents the IdfaDeclaration schema from the OpenAPI specification
type IdfaDeclaration struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// AppInfoLocalization represents the AppInfoLocalization schema from the OpenAPI specification
type AppInfoLocalization struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// CertificateCreateRequest represents the CertificateCreateRequest schema from the OpenAPI specification
type CertificateCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPreviewsResponse represents the AppPreviewsResponse schema from the OpenAPI specification
type AppPreviewsResponse struct {
	Data []AppPreview `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaAppLocalizationCreateRequest represents the BetaAppLocalizationCreateRequest schema from the OpenAPI specification
type BetaAppLocalizationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaLicenseAgreement represents the BetaLicenseAgreement schema from the OpenAPI specification
type BetaLicenseAgreement struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// AppStoreReviewDetailCreateRequest represents the AppStoreReviewDetailCreateRequest schema from the OpenAPI specification
type AppStoreReviewDetailCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppScreenshotSetsResponse represents the AppScreenshotSetsResponse schema from the OpenAPI specification
type AppScreenshotSetsResponse struct {
	Included []AppScreenshot `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppScreenshotSet `json:"data"`
}

// AppPreOrderCreateRequest represents the AppPreOrderCreateRequest schema from the OpenAPI specification
type AppPreOrderCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// PreReleaseVersionsResponse represents the PreReleaseVersionsResponse schema from the OpenAPI specification
type PreReleaseVersionsResponse struct {
	Data []PrereleaseVersion `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaAppReviewSubmissionResponse represents the BetaAppReviewSubmissionResponse schema from the OpenAPI specification
type BetaAppReviewSubmissionResponse struct {
	Links DocumentLinks `json:"links"`
	Data BetaAppReviewSubmission `json:"data"`
	Included []Build `json:"included,omitempty"`
}

// BuildBetaDetailsResponse represents the BuildBetaDetailsResponse schema from the OpenAPI specification
type BuildBetaDetailsResponse struct {
	Data []BuildBetaDetail `json:"data"`
	Included []Build `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaTesterInvitation represents the BetaTesterInvitation schema from the OpenAPI specification
type BetaTesterInvitation struct {
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Id string `json:"id"`
}

// AppStoreReviewAttachmentUpdateRequest represents the AppStoreReviewAttachmentUpdateRequest schema from the OpenAPI specification
type AppStoreReviewAttachmentUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// DocumentLinks represents the DocumentLinks schema from the OpenAPI specification
type DocumentLinks struct {
	Self string `json:"self"`
}

// TerritoryResponse represents the TerritoryResponse schema from the OpenAPI specification
type TerritoryResponse struct {
	Data Territory `json:"data"`
	Links DocumentLinks `json:"links"`
}

// Territory represents the Territory schema from the OpenAPI specification
type Territory struct {
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
}

// BuildIconsResponse represents the BuildIconsResponse schema from the OpenAPI specification
type BuildIconsResponse struct {
	Data []BuildIcon `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BuildIndividualTestersLinkagesResponse represents the BuildIndividualTestersLinkagesResponse schema from the OpenAPI specification
type BuildIndividualTestersLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaBuildLocalization represents the BetaBuildLocalization schema from the OpenAPI specification
type BetaBuildLocalization struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// IdfaDeclarationCreateRequest represents the IdfaDeclarationCreateRequest schema from the OpenAPI specification
type IdfaDeclarationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPreviewSetsResponse represents the AppPreviewSetsResponse schema from the OpenAPI specification
type AppPreviewSetsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppPreviewSet `json:"data"`
	Included []AppPreview `json:"included,omitempty"`
}

// AppCategoriesResponse represents the AppCategoriesResponse schema from the OpenAPI specification
type AppCategoriesResponse struct {
	Data []AppCategory `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaBuildLocalizationCreateRequest represents the BetaBuildLocalizationCreateRequest schema from the OpenAPI specification
type BetaBuildLocalizationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPreOrderUpdateRequest represents the AppPreOrderUpdateRequest schema from the OpenAPI specification
type AppPreOrderUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaTesterInvitationResponse represents the BetaTesterInvitationResponse schema from the OpenAPI specification
type BetaTesterInvitationResponse struct {
	Data BetaTesterInvitation `json:"data"`
	Links DocumentLinks `json:"links"`
}

// EndUserLicenseAgreementUpdateRequest represents the EndUserLicenseAgreementUpdateRequest schema from the OpenAPI specification
type EndUserLicenseAgreementUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppStoreVersionBuildLinkageRequest represents the AppStoreVersionBuildLinkageRequest schema from the OpenAPI specification
type AppStoreVersionBuildLinkageRequest struct {
	Data map[string]interface{} `json:"data"`
}

// CapabilityOption represents the CapabilityOption schema from the OpenAPI specification
type CapabilityOption struct {
	Enabledbydefault bool `json:"enabledByDefault,omitempty"`
	Key string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
	Supportswildcard bool `json:"supportsWildcard,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// AppPreviewSetCreateRequest represents the AppPreviewSetCreateRequest schema from the OpenAPI specification
type AppPreviewSetCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// EndUserLicenseAgreementCreateRequest represents the EndUserLicenseAgreementCreateRequest schema from the OpenAPI specification
type EndUserLicenseAgreementCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// ImageAsset represents the ImageAsset schema from the OpenAPI specification
type ImageAsset struct {
	Height int `json:"height,omitempty"`
	Templateurl string `json:"templateUrl,omitempty"`
	Width int `json:"width,omitempty"`
}

// AppScreenshotsResponse represents the AppScreenshotsResponse schema from the OpenAPI specification
type AppScreenshotsResponse struct {
	Data []AppScreenshot `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppScreenshotSetAppScreenshotsLinkagesRequest represents the AppScreenshotSetAppScreenshotsLinkagesRequest schema from the OpenAPI specification
type AppScreenshotSetAppScreenshotsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// BetaTester represents the BetaTester schema from the OpenAPI specification
type BetaTester struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// AppInfoLocalizationUpdateRequest represents the AppInfoLocalizationUpdateRequest schema from the OpenAPI specification
type AppInfoLocalizationUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// UserInvitationCreateRequest represents the UserInvitationCreateRequest schema from the OpenAPI specification
type UserInvitationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPriceResponse represents the AppPriceResponse schema from the OpenAPI specification
type AppPriceResponse struct {
	Data AppPrice `json:"data"`
	Links DocumentLinks `json:"links"`
}

// InAppPurchaseResponse represents the InAppPurchaseResponse schema from the OpenAPI specification
type InAppPurchaseResponse struct {
	Data InAppPurchase `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppPreviewCreateRequest represents the AppPreviewCreateRequest schema from the OpenAPI specification
type AppPreviewCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BundleIdCapability represents the BundleIdCapability schema from the OpenAPI specification
type BundleIdCapability struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
}

// BetaTesterBuildsLinkagesRequest represents the BetaTesterBuildsLinkagesRequest schema from the OpenAPI specification
type BetaTesterBuildsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// BundleIdCapabilityResponse represents the BundleIdCapabilityResponse schema from the OpenAPI specification
type BundleIdCapabilityResponse struct {
	Data BundleIdCapability `json:"data"`
	Links DocumentLinks `json:"links"`
}

// BuildUpdateRequest represents the BuildUpdateRequest schema from the OpenAPI specification
type BuildUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// CertificateResponse represents the CertificateResponse schema from the OpenAPI specification
type CertificateResponse struct {
	Data Certificate `json:"data"`
	Links DocumentLinks `json:"links"`
}

// ProfilesResponse represents the ProfilesResponse schema from the OpenAPI specification
type ProfilesResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []Profile `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// UserVisibleAppsLinkagesRequest represents the UserVisibleAppsLinkagesRequest schema from the OpenAPI specification
type UserVisibleAppsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// BetaGroup represents the BetaGroup schema from the OpenAPI specification
type BetaGroup struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// Certificate represents the Certificate schema from the OpenAPI specification
type Certificate struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
}

// AppPreOrderResponse represents the AppPreOrderResponse schema from the OpenAPI specification
type AppPreOrderResponse struct {
	Data AppPreOrder `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppCategoryResponse represents the AppCategoryResponse schema from the OpenAPI specification
type AppCategoryResponse struct {
	Data AppCategory `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// BetaGroupUpdateRequest represents the BetaGroupUpdateRequest schema from the OpenAPI specification
type BetaGroupUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BuildBetaNotificationResponse represents the BuildBetaNotificationResponse schema from the OpenAPI specification
type BuildBetaNotificationResponse struct {
	Data BuildBetaNotification `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppScreenshotSet represents the AppScreenshotSet schema from the OpenAPI specification
type AppScreenshotSet struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// BetaLicenseAgreementUpdateRequest represents the BetaLicenseAgreementUpdateRequest schema from the OpenAPI specification
type BetaLicenseAgreementUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppStoreVersionLocalizationsResponse represents the AppStoreVersionLocalizationsResponse schema from the OpenAPI specification
type AppStoreVersionLocalizationsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppStoreVersionLocalization `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// BetaAppReviewDetailUpdateRequest represents the BetaAppReviewDetailUpdateRequest schema from the OpenAPI specification
type BetaAppReviewDetailUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// ProfileResponse represents the ProfileResponse schema from the OpenAPI specification
type ProfileResponse struct {
	Links DocumentLinks `json:"links"`
	Data Profile `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// BundleIdCapabilityUpdateRequest represents the BundleIdCapabilityUpdateRequest schema from the OpenAPI specification
type BundleIdCapabilityUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPricePointResponse represents the AppPricePointResponse schema from the OpenAPI specification
type AppPricePointResponse struct {
	Data AppPricePoint `json:"data"`
	Included []Territory `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// BuildBetaGroupsLinkagesRequest represents the BuildBetaGroupsLinkagesRequest schema from the OpenAPI specification
type BuildBetaGroupsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// BetaTestersResponse represents the BetaTestersResponse schema from the OpenAPI specification
type BetaTestersResponse struct {
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []BetaTester `json:"data"`
}

// AgeRatingDeclarationUpdateRequest represents the AgeRatingDeclarationUpdateRequest schema from the OpenAPI specification
type AgeRatingDeclarationUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// Device represents the Device schema from the OpenAPI specification
type Device struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// BundleIdCreateRequest represents the BundleIdCreateRequest schema from the OpenAPI specification
type BundleIdCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPreviewSetResponse represents the AppPreviewSetResponse schema from the OpenAPI specification
type AppPreviewSetResponse struct {
	Data AppPreviewSet `json:"data"`
	Included []AppPreview `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// AppUpdateRequest represents the AppUpdateRequest schema from the OpenAPI specification
type AppUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppStoreVersionBuildLinkageResponse represents the AppStoreVersionBuildLinkageResponse schema from the OpenAPI specification
type AppStoreVersionBuildLinkageResponse struct {
	Data map[string]interface{} `json:"data"`
	Links DocumentLinks `json:"links"`
}

// BetaTesterBetaGroupsLinkagesRequest represents the BetaTesterBetaGroupsLinkagesRequest schema from the OpenAPI specification
type BetaTesterBetaGroupsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesRequest represents the GameCenterEnabledVersionCompatibleVersionsLinkagesRequest schema from the OpenAPI specification
type GameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// BetaAppReviewDetailsResponse represents the BetaAppReviewDetailsResponse schema from the OpenAPI specification
type BetaAppReviewDetailsResponse struct {
	Data []BetaAppReviewDetail `json:"data"`
	Included []App `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaTesterAppsLinkagesResponse represents the BetaTesterAppsLinkagesResponse schema from the OpenAPI specification
type BetaTesterAppsLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppPreviewResponse represents the AppPreviewResponse schema from the OpenAPI specification
type AppPreviewResponse struct {
	Data AppPreview `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppEncryptionDeclarationsResponse represents the AppEncryptionDeclarationsResponse schema from the OpenAPI specification
type AppEncryptionDeclarationsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppEncryptionDeclaration `json:"data"`
	Included []App `json:"included,omitempty"`
}

// UserResponse represents the UserResponse schema from the OpenAPI specification
type UserResponse struct {
	Data User `json:"data"`
	Included []App `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// BuildBetaDetailResponse represents the BuildBetaDetailResponse schema from the OpenAPI specification
type BuildBetaDetailResponse struct {
	Data BuildBetaDetail `json:"data"`
	Included []Build `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// AppScreenshotUpdateRequest represents the AppScreenshotUpdateRequest schema from the OpenAPI specification
type AppScreenshotUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppStoreVersionPhasedRelease represents the AppStoreVersionPhasedRelease schema from the OpenAPI specification
type AppStoreVersionPhasedRelease struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// AppScreenshotSetCreateRequest represents the AppScreenshotSetCreateRequest schema from the OpenAPI specification
type AppScreenshotSetCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppStoreVersionLocalizationResponse represents the AppStoreVersionLocalizationResponse schema from the OpenAPI specification
type AppStoreVersionLocalizationResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppStoreVersionLocalization `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// UserInvitationResponse represents the UserInvitationResponse schema from the OpenAPI specification
type UserInvitationResponse struct {
	Data UserInvitation `json:"data"`
	Included []App `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// PrereleaseVersionResponse represents the PrereleaseVersionResponse schema from the OpenAPI specification
type PrereleaseVersionResponse struct {
	Data PrereleaseVersion `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// AppMediaAssetState represents the AppMediaAssetState schema from the OpenAPI specification
type AppMediaAssetState struct {
	Errors []AppMediaStateError `json:"errors,omitempty"`
	State string `json:"state,omitempty"`
	Warnings []AppMediaStateError `json:"warnings,omitempty"`
}

// AppStoreVersionSubmissionCreateRequest represents the AppStoreVersionSubmissionCreateRequest schema from the OpenAPI specification
type AppStoreVersionSubmissionCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// PerfPowerMetricsResponse represents the PerfPowerMetricsResponse schema from the OpenAPI specification
type PerfPowerMetricsResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []PerfPowerMetric `json:"data"`
	Links PagedDocumentLinks `json:"links"`
}

// UserVisibleAppsLinkagesResponse represents the UserVisibleAppsLinkagesResponse schema from the OpenAPI specification
type UserVisibleAppsLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// PagedDocumentLinks represents the PagedDocumentLinks schema from the OpenAPI specification
type PagedDocumentLinks struct {
	First string `json:"first,omitempty"`
	Next string `json:"next,omitempty"`
	Self string `json:"self"`
}

// GameCenterEnabledVersionsResponse represents the GameCenterEnabledVersionsResponse schema from the OpenAPI specification
type GameCenterEnabledVersionsResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []GameCenterEnabledVersion `json:"data"`
	Included []GameCenterEnabledVersion `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
}

// DiagnosticLog represents the DiagnosticLog schema from the OpenAPI specification
type DiagnosticLog struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
}

// AppPriceTier represents the AppPriceTier schema from the OpenAPI specification
type AppPriceTier struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// BetaGroupCreateRequest represents the BetaGroupCreateRequest schema from the OpenAPI specification
type BetaGroupCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// UploadOperation represents the UploadOperation schema from the OpenAPI specification
type UploadOperation struct {
	Offset int `json:"offset,omitempty"`
	Requestheaders []UploadOperationHeader `json:"requestHeaders,omitempty"`
	Url string `json:"url,omitempty"`
	Length int `json:"length,omitempty"`
	Method string `json:"method,omitempty"`
}

// AppStoreReviewDetail represents the AppStoreReviewDetail schema from the OpenAPI specification
type AppStoreReviewDetail struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// BetaGroupBetaTestersLinkagesResponse represents the BetaGroupBetaTestersLinkagesResponse schema from the OpenAPI specification
type BetaGroupBetaTestersLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppStoreReviewDetailResponse represents the AppStoreReviewDetailResponse schema from the OpenAPI specification
type AppStoreReviewDetailResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppStoreReviewDetail `json:"data"`
	Included []AppStoreReviewAttachment `json:"included,omitempty"`
}

// RoutingAppCoverage represents the RoutingAppCoverage schema from the OpenAPI specification
type RoutingAppCoverage struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// InAppPurchase represents the InAppPurchase schema from the OpenAPI specification
type InAppPurchase struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// AppPreOrder represents the AppPreOrder schema from the OpenAPI specification
type AppPreOrder struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// AppStoreVersionLocalizationUpdateRequest represents the AppStoreVersionLocalizationUpdateRequest schema from the OpenAPI specification
type AppStoreVersionLocalizationUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// IdfaDeclarationResponse represents the IdfaDeclarationResponse schema from the OpenAPI specification
type IdfaDeclarationResponse struct {
	Links DocumentLinks `json:"links"`
	Data IdfaDeclaration `json:"data"`
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// AgeRatingDeclarationResponse represents the AgeRatingDeclarationResponse schema from the OpenAPI specification
type AgeRatingDeclarationResponse struct {
	Links DocumentLinks `json:"links"`
	Data AgeRatingDeclaration `json:"data"`
}

// AppStoreVersionPhasedReleaseCreateRequest represents the AppStoreVersionPhasedReleaseCreateRequest schema from the OpenAPI specification
type AppStoreVersionPhasedReleaseCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BuildBetaNotificationCreateRequest represents the BuildBetaNotificationCreateRequest schema from the OpenAPI specification
type BuildBetaNotificationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BuildsResponse represents the BuildsResponse schema from the OpenAPI specification
type BuildsResponse struct {
	Data []Build `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppInfoLocalizationsResponse represents the AppInfoLocalizationsResponse schema from the OpenAPI specification
type AppInfoLocalizationsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppInfoLocalization `json:"data"`
}

// AppStoreReviewAttachment represents the AppStoreReviewAttachment schema from the OpenAPI specification
type AppStoreReviewAttachment struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// PerfPowerMetric represents the PerfPowerMetric schema from the OpenAPI specification
type PerfPowerMetric struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// BetaAppLocalizationsResponse represents the BetaAppLocalizationsResponse schema from the OpenAPI specification
type BetaAppLocalizationsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []BetaAppLocalization `json:"data"`
	Included []App `json:"included,omitempty"`
}

// AppStoreVersionLocalizationCreateRequest represents the AppStoreVersionLocalizationCreateRequest schema from the OpenAPI specification
type AppStoreVersionLocalizationCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// EndUserLicenseAgreement represents the EndUserLicenseAgreement schema from the OpenAPI specification
type EndUserLicenseAgreement struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// AppPreviewSet represents the AppPreviewSet schema from the OpenAPI specification
type AppPreviewSet struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// BuildBetaNotification represents the BuildBetaNotification schema from the OpenAPI specification
type BuildBetaNotification struct {
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Id string `json:"id"`
}

// AppScreenshotResponse represents the AppScreenshotResponse schema from the OpenAPI specification
type AppScreenshotResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppScreenshot `json:"data"`
}

// AppPriceTiersResponse represents the AppPriceTiersResponse schema from the OpenAPI specification
type AppPriceTiersResponse struct {
	Data []AppPriceTier `json:"data"`
	Included []AppPricePoint `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppPricesResponse represents the AppPricesResponse schema from the OpenAPI specification
type AppPricesResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppPrice `json:"data"`
}

// BundleIdResponse represents the BundleIdResponse schema from the OpenAPI specification
type BundleIdResponse struct {
	Data BundleId `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// AppStoreVersionPhasedReleaseUpdateRequest represents the AppStoreVersionPhasedReleaseUpdateRequest schema from the OpenAPI specification
type AppStoreVersionPhasedReleaseUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BundleId represents the BundleId schema from the OpenAPI specification
type BundleId struct {
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
}

// BuildAppEncryptionDeclarationLinkageRequest represents the BuildAppEncryptionDeclarationLinkageRequest schema from the OpenAPI specification
type BuildAppEncryptionDeclarationLinkageRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BundleIdCapabilitiesResponse represents the BundleIdCapabilitiesResponse schema from the OpenAPI specification
type BundleIdCapabilitiesResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []BundleIdCapability `json:"data"`
}

// EndUserLicenseAgreementResponse represents the EndUserLicenseAgreementResponse schema from the OpenAPI specification
type EndUserLicenseAgreementResponse struct {
	Included []Territory `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
	Data EndUserLicenseAgreement `json:"data"`
}

// BetaGroupBetaTestersLinkagesRequest represents the BetaGroupBetaTestersLinkagesRequest schema from the OpenAPI specification
type BetaGroupBetaTestersLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// AppPreviewSetAppPreviewsLinkagesResponse represents the AppPreviewSetAppPreviewsLinkagesResponse schema from the OpenAPI specification
type AppPreviewSetAppPreviewsLinkagesResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []map[string]interface{} `json:"data"`
}

// AppScreenshotSetResponse represents the AppScreenshotSetResponse schema from the OpenAPI specification
type AppScreenshotSetResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppScreenshotSet `json:"data"`
	Included []AppScreenshot `json:"included,omitempty"`
}

// AppStoreVersionLocalization represents the AppStoreVersionLocalization schema from the OpenAPI specification
type AppStoreVersionLocalization struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// BetaLicenseAgreementResponse represents the BetaLicenseAgreementResponse schema from the OpenAPI specification
type BetaLicenseAgreementResponse struct {
	Included []App `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
	Data BetaLicenseAgreement `json:"data"`
}

// BetaGroupBuildsLinkagesResponse represents the BetaGroupBuildsLinkagesResponse schema from the OpenAPI specification
type BetaGroupBuildsLinkagesResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []map[string]interface{} `json:"data"`
}

// AppInfosResponse represents the AppInfosResponse schema from the OpenAPI specification
type AppInfosResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppInfo `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
}

// BetaGroupBuildsLinkagesRequest represents the BetaGroupBuildsLinkagesRequest schema from the OpenAPI specification
type BetaGroupBuildsLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// AppStoreVersionSubmissionResponse represents the AppStoreVersionSubmissionResponse schema from the OpenAPI specification
type AppStoreVersionSubmissionResponse struct {
	Data AppStoreVersionSubmission `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppPricePointsResponse represents the AppPricePointsResponse schema from the OpenAPI specification
type AppPricePointsResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppPricePoint `json:"data"`
	Included []Territory `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
}

// IdfaDeclarationUpdateRequest represents the IdfaDeclarationUpdateRequest schema from the OpenAPI specification
type IdfaDeclarationUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppPricePoint represents the AppPricePoint schema from the OpenAPI specification
type AppPricePoint struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// AppPreview represents the AppPreview schema from the OpenAPI specification
type AppPreview struct {
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
}

// RoutingAppCoverageResponse represents the RoutingAppCoverageResponse schema from the OpenAPI specification
type RoutingAppCoverageResponse struct {
	Data RoutingAppCoverage `json:"data"`
	Links DocumentLinks `json:"links"`
}

// BuildBetaDetailUpdateRequest represents the BuildBetaDetailUpdateRequest schema from the OpenAPI specification
type BuildBetaDetailUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaAppReviewDetail represents the BetaAppReviewDetail schema from the OpenAPI specification
type BetaAppReviewDetail struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// AppMediaStateError represents the AppMediaStateError schema from the OpenAPI specification
type AppMediaStateError struct {
	Description string `json:"description,omitempty"`
	Code string `json:"code,omitempty"`
}

// AppStoreVersionPhasedReleaseResponse represents the AppStoreVersionPhasedReleaseResponse schema from the OpenAPI specification
type AppStoreVersionPhasedReleaseResponse struct {
	Data AppStoreVersionPhasedRelease `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppScreenshotCreateRequest represents the AppScreenshotCreateRequest schema from the OpenAPI specification
type AppScreenshotCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// CapabilitySetting represents the CapabilitySetting schema from the OpenAPI specification
type CapabilitySetting struct {
	Allowedinstances string `json:"allowedInstances,omitempty"`
	Description string `json:"description,omitempty"`
	Enabledbydefault bool `json:"enabledByDefault,omitempty"`
	Key string `json:"key,omitempty"`
	Mininstances int `json:"minInstances,omitempty"`
	Name string `json:"name,omitempty"`
	Options []CapabilityOption `json:"options,omitempty"`
	Visible bool `json:"visible,omitempty"`
}

// AppStoreReviewAttachmentResponse represents the AppStoreReviewAttachmentResponse schema from the OpenAPI specification
type AppStoreReviewAttachmentResponse struct {
	Links DocumentLinks `json:"links"`
	Data AppStoreReviewAttachment `json:"data"`
}

// AppEncryptionDeclaration represents the AppEncryptionDeclaration schema from the OpenAPI specification
type AppEncryptionDeclaration struct {
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// AppScreenshot represents the AppScreenshot schema from the OpenAPI specification
type AppScreenshot struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// UserInvitation represents the UserInvitation schema from the OpenAPI specification
type UserInvitation struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// BuildBetaDetail represents the BuildBetaDetail schema from the OpenAPI specification
type BuildBetaDetail struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// AppEncryptionDeclarationResponse represents the AppEncryptionDeclarationResponse schema from the OpenAPI specification
type AppEncryptionDeclarationResponse struct {
	Data AppEncryptionDeclaration `json:"data"`
	Included []App `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// AppInfoUpdateRequest represents the AppInfoUpdateRequest schema from the OpenAPI specification
type AppInfoUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppResponse represents the AppResponse schema from the OpenAPI specification
type AppResponse struct {
	Links DocumentLinks `json:"links"`
	Data App `json:"data"`
	Included []interface{} `json:"included,omitempty"`
}

// AppStoreVersionResponse represents the AppStoreVersionResponse schema from the OpenAPI specification
type AppStoreVersionResponse struct {
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
	Data AppStoreVersion `json:"data"`
}

// ResourceLinks represents the ResourceLinks schema from the OpenAPI specification
type ResourceLinks struct {
	Self string `json:"self"`
}

// TerritoriesResponse represents the TerritoriesResponse schema from the OpenAPI specification
type TerritoriesResponse struct {
	Data []Territory `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// GameCenterEnabledVersionCompatibleVersionsLinkagesResponse represents the GameCenterEnabledVersionCompatibleVersionsLinkagesResponse schema from the OpenAPI specification
type GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BetaGroupsResponse represents the BetaGroupsResponse schema from the OpenAPI specification
type BetaGroupsResponse struct {
	Data []BetaGroup `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// AppStoreReviewAttachmentCreateRequest represents the AppStoreReviewAttachmentCreateRequest schema from the OpenAPI specification
type AppStoreReviewAttachmentCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaBuildLocalizationUpdateRequest represents the BetaBuildLocalizationUpdateRequest schema from the OpenAPI specification
type BetaBuildLocalizationUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// InAppPurchasesResponse represents the InAppPurchasesResponse schema from the OpenAPI specification
type InAppPurchasesResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []InAppPurchase `json:"data"`
	Links PagedDocumentLinks `json:"links"`
}

// AppStoreVersion represents the AppStoreVersion schema from the OpenAPI specification
type AppStoreVersion struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// DeviceUpdateRequest represents the DeviceUpdateRequest schema from the OpenAPI specification
type DeviceUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// BetaAppReviewSubmission represents the BetaAppReviewSubmission schema from the OpenAPI specification
type BetaAppReviewSubmission struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// BetaTesterCreateRequest represents the BetaTesterCreateRequest schema from the OpenAPI specification
type BetaTesterCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// AppInfoLocalizationResponse represents the AppInfoLocalizationResponse schema from the OpenAPI specification
type AppInfoLocalizationResponse struct {
	Data AppInfoLocalization `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppStoreVersionUpdateRequest represents the AppStoreVersionUpdateRequest schema from the OpenAPI specification
type AppStoreVersionUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// UploadOperationHeader represents the UploadOperationHeader schema from the OpenAPI specification
type UploadOperationHeader struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// UsersResponse represents the UsersResponse schema from the OpenAPI specification
type UsersResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []User `json:"data"`
	Included []App `json:"included,omitempty"`
}

// DiagnosticSignaturesResponse represents the DiagnosticSignaturesResponse schema from the OpenAPI specification
type DiagnosticSignaturesResponse struct {
	Data []DiagnosticSignature `json:"data"`
	Included []DiagnosticLog `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BuildAppEncryptionDeclarationLinkageResponse represents the BuildAppEncryptionDeclarationLinkageResponse schema from the OpenAPI specification
type BuildAppEncryptionDeclarationLinkageResponse struct {
	Data map[string]interface{} `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppCategory represents the AppCategory schema from the OpenAPI specification
type AppCategory struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
}

// DeviceResponse represents the DeviceResponse schema from the OpenAPI specification
type DeviceResponse struct {
	Data Device `json:"data"`
	Links DocumentLinks `json:"links"`
}

// GameCenterEnabledVersion represents the GameCenterEnabledVersion schema from the OpenAPI specification
type GameCenterEnabledVersion struct {
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
}

// BetaAppLocalizationUpdateRequest represents the BetaAppLocalizationUpdateRequest schema from the OpenAPI specification
type BetaAppLocalizationUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// DiagnosticLogsResponse represents the DiagnosticLogsResponse schema from the OpenAPI specification
type DiagnosticLogsResponse struct {
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
	Data []DiagnosticLog `json:"data"`
}

// AppScreenshotSetAppScreenshotsLinkagesResponse represents the AppScreenshotSetAppScreenshotsLinkagesResponse schema from the OpenAPI specification
type AppScreenshotSetAppScreenshotsLinkagesResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// BuildIndividualTestersLinkagesRequest represents the BuildIndividualTestersLinkagesRequest schema from the OpenAPI specification
type BuildIndividualTestersLinkagesRequest struct {
	Data []map[string]interface{} `json:"data"`
}

// BetaTesterResponse represents the BetaTesterResponse schema from the OpenAPI specification
type BetaTesterResponse struct {
	Data BetaTester `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// BetaAppReviewSubmissionCreateRequest represents the BetaAppReviewSubmissionCreateRequest schema from the OpenAPI specification
type BetaAppReviewSubmissionCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// DevicesResponse represents the DevicesResponse schema from the OpenAPI specification
type DevicesResponse struct {
	Data []Device `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta PagingInformation `json:"meta,omitempty"`
}

// RoutingAppCoverageUpdateRequest represents the RoutingAppCoverageUpdateRequest schema from the OpenAPI specification
type RoutingAppCoverageUpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// DiagnosticSignature represents the DiagnosticSignature schema from the OpenAPI specification
type DiagnosticSignature struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// AppStoreVersionCreateRequest represents the AppStoreVersionCreateRequest schema from the OpenAPI specification
type AppStoreVersionCreateRequest struct {
	Data map[string]interface{} `json:"data"`
}

// Build represents the Build schema from the OpenAPI specification
type Build struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type"`
}

// AgeRatingDeclaration represents the AgeRatingDeclaration schema from the OpenAPI specification
type AgeRatingDeclaration struct {
	Id string `json:"id"`
	Links ResourceLinks `json:"links"`
	TypeField string `json:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// AppStoreVersionsResponse represents the AppStoreVersionsResponse schema from the OpenAPI specification
type AppStoreVersionsResponse struct {
	Meta PagingInformation `json:"meta,omitempty"`
	Data []AppStoreVersion `json:"data"`
	Included []interface{} `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
}
