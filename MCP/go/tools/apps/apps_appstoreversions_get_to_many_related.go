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

func Apps_appstoreversions_get_to_many_relatedHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter[appStoreState]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[appStoreState]=%v", val))
		}
		if val, ok := args["filter[platform]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[platform]=%v", val))
		}
		if val, ok := args["filter[versionString]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[versionString]=%v", val))
		}
		if val, ok := args["filter[id]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[id]=%v", val))
		}
		if val, ok := args["fields[idfaDeclarations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[idfaDeclarations]=%v", val))
		}
		if val, ok := args["fields[appStoreVersionLocalizations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersionLocalizations]=%v", val))
		}
		if val, ok := args["fields[routingAppCoverages]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[routingAppCoverages]=%v", val))
		}
		if val, ok := args["fields[appStoreVersionPhasedReleases]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersionPhasedReleases]=%v", val))
		}
		if val, ok := args["fields[ageRatingDeclarations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[ageRatingDeclarations]=%v", val))
		}
		if val, ok := args["fields[appStoreReviewDetails]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreReviewDetails]=%v", val))
		}
		if val, ok := args["fields[appStoreVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersions]=%v", val))
		}
		if val, ok := args["fields[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[builds]=%v", val))
		}
		if val, ok := args["fields[appStoreVersionSubmissions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersionSubmissions]=%v", val))
		}
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/apps/%s/appStoreVersions%s", cfg.BaseURL, queryString)
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
		var result models.AppStoreVersionsResponse
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

func CreateApps_appstoreversions_get_to_many_relatedTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_apps_id_appStoreVersions",
		mcp.WithDescription(""),
		mcp.WithArray("filter[appStoreState]", mcp.Description("filter by attribute 'appStoreState'")),
		mcp.WithArray("filter[platform]", mcp.Description("filter by attribute 'platform'")),
		mcp.WithArray("filter[versionString]", mcp.Description("filter by attribute 'versionString'")),
		mcp.WithArray("filter[id]", mcp.Description("filter by id(s)")),
		mcp.WithArray("fields[idfaDeclarations]", mcp.Description("the fields to include for returned resources of type idfaDeclarations")),
		mcp.WithArray("fields[appStoreVersionLocalizations]", mcp.Description("the fields to include for returned resources of type appStoreVersionLocalizations")),
		mcp.WithArray("fields[routingAppCoverages]", mcp.Description("the fields to include for returned resources of type routingAppCoverages")),
		mcp.WithArray("fields[appStoreVersionPhasedReleases]", mcp.Description("the fields to include for returned resources of type appStoreVersionPhasedReleases")),
		mcp.WithArray("fields[ageRatingDeclarations]", mcp.Description("the fields to include for returned resources of type ageRatingDeclarations")),
		mcp.WithArray("fields[appStoreReviewDetails]", mcp.Description("the fields to include for returned resources of type appStoreReviewDetails")),
		mcp.WithArray("fields[appStoreVersions]", mcp.Description("the fields to include for returned resources of type appStoreVersions")),
		mcp.WithArray("fields[builds]", mcp.Description("the fields to include for returned resources of type builds")),
		mcp.WithArray("fields[appStoreVersionSubmissions]", mcp.Description("the fields to include for returned resources of type appStoreVersionSubmissions")),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithNumber("limit", mcp.Description("maximum resources per page")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apps_appstoreversions_get_to_many_relatedHandler(cfg),
	}
}
