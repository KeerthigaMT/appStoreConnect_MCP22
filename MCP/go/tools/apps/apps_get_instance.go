package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/app-store-connect-api/mcp-server/config"
	"github.com/app-store-connect-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Apps_get_instanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[betaGroups]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaGroups]=%v", val))
		}
		if val, ok := args["fields[perfPowerMetrics]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[perfPowerMetrics]=%v", val))
		}
		if val, ok := args["fields[appInfos]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appInfos]=%v", val))
		}
		if val, ok := args["fields[appPreOrders]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appPreOrders]=%v", val))
		}
		if val, ok := args["fields[preReleaseVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[preReleaseVersions]=%v", val))
		}
		if val, ok := args["fields[appPrices]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appPrices]=%v", val))
		}
		if val, ok := args["fields[inAppPurchases]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[inAppPurchases]=%v", val))
		}
		if val, ok := args["fields[betaAppReviewDetails]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaAppReviewDetails]=%v", val))
		}
		if val, ok := args["fields[territories]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[territories]=%v", val))
		}
		if val, ok := args["fields[gameCenterEnabledVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[gameCenterEnabledVersions]=%v", val))
		}
		if val, ok := args["fields[appStoreVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersions]=%v", val))
		}
		if val, ok := args["fields[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[builds]=%v", val))
		}
		if val, ok := args["fields[betaAppLocalizations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaAppLocalizations]=%v", val))
		}
		if val, ok := args["fields[betaLicenseAgreements]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaLicenseAgreements]=%v", val))
		}
		if val, ok := args["fields[endUserLicenseAgreements]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[endUserLicenseAgreements]=%v", val))
		}
		if val, ok := args["limit[appInfos]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[appInfos]=%v", val))
		}
		if val, ok := args["limit[appStoreVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[appStoreVersions]=%v", val))
		}
		if val, ok := args["limit[availableTerritories]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[availableTerritories]=%v", val))
		}
		if val, ok := args["limit[betaAppLocalizations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[betaAppLocalizations]=%v", val))
		}
		if val, ok := args["limit[betaGroups]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[betaGroups]=%v", val))
		}
		if val, ok := args["limit[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[builds]=%v", val))
		}
		if val, ok := args["limit[gameCenterEnabledVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[gameCenterEnabledVersions]=%v", val))
		}
		if val, ok := args["limit[inAppPurchases]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[inAppPurchases]=%v", val))
		}
		if val, ok := args["limit[preReleaseVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[preReleaseVersions]=%v", val))
		}
		if val, ok := args["limit[prices]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[prices]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/apps/%s%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.AppResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateApps_get_instanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_apps_id",
		mcp.WithDescription(""),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[betaGroups]", mcp.Description("the fields to include for returned resources of type betaGroups")),
		mcp.WithArray("fields[perfPowerMetrics]", mcp.Description("the fields to include for returned resources of type perfPowerMetrics")),
		mcp.WithArray("fields[appInfos]", mcp.Description("the fields to include for returned resources of type appInfos")),
		mcp.WithArray("fields[appPreOrders]", mcp.Description("the fields to include for returned resources of type appPreOrders")),
		mcp.WithArray("fields[preReleaseVersions]", mcp.Description("the fields to include for returned resources of type preReleaseVersions")),
		mcp.WithArray("fields[appPrices]", mcp.Description("the fields to include for returned resources of type appPrices")),
		mcp.WithArray("fields[inAppPurchases]", mcp.Description("the fields to include for returned resources of type inAppPurchases")),
		mcp.WithArray("fields[betaAppReviewDetails]", mcp.Description("the fields to include for returned resources of type betaAppReviewDetails")),
		mcp.WithArray("fields[territories]", mcp.Description("the fields to include for returned resources of type territories")),
		mcp.WithArray("fields[gameCenterEnabledVersions]", mcp.Description("the fields to include for returned resources of type gameCenterEnabledVersions")),
		mcp.WithArray("fields[appStoreVersions]", mcp.Description("the fields to include for returned resources of type appStoreVersions")),
		mcp.WithArray("fields[builds]", mcp.Description("the fields to include for returned resources of type builds")),
		mcp.WithArray("fields[betaAppLocalizations]", mcp.Description("the fields to include for returned resources of type betaAppLocalizations")),
		mcp.WithArray("fields[betaLicenseAgreements]", mcp.Description("the fields to include for returned resources of type betaLicenseAgreements")),
		mcp.WithArray("fields[endUserLicenseAgreements]", mcp.Description("the fields to include for returned resources of type endUserLicenseAgreements")),
		mcp.WithNumber("limit[appInfos]", mcp.Description("maximum number of related appInfos returned (when they are included)")),
		mcp.WithNumber("limit[appStoreVersions]", mcp.Description("maximum number of related appStoreVersions returned (when they are included)")),
		mcp.WithNumber("limit[availableTerritories]", mcp.Description("maximum number of related availableTerritories returned (when they are included)")),
		mcp.WithNumber("limit[betaAppLocalizations]", mcp.Description("maximum number of related betaAppLocalizations returned (when they are included)")),
		mcp.WithNumber("limit[betaGroups]", mcp.Description("maximum number of related betaGroups returned (when they are included)")),
		mcp.WithNumber("limit[builds]", mcp.Description("maximum number of related builds returned (when they are included)")),
		mcp.WithNumber("limit[gameCenterEnabledVersions]", mcp.Description("maximum number of related gameCenterEnabledVersions returned (when they are included)")),
		mcp.WithNumber("limit[inAppPurchases]", mcp.Description("maximum number of related inAppPurchases returned (when they are included)")),
		mcp.WithNumber("limit[preReleaseVersions]", mcp.Description("maximum number of related preReleaseVersions returned (when they are included)")),
		mcp.WithNumber("limit[prices]", mcp.Description("maximum number of related prices returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apps_get_instanceHandler(cfg),
	}
}
