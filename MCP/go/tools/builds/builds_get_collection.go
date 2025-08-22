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

func Builds_get_collectionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter[betaAppReviewSubmission.betaReviewState]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[betaAppReviewSubmission.betaReviewState]=%v", val))
		}
		if val, ok := args["filter[expired]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[expired]=%v", val))
		}
		if val, ok := args["filter[preReleaseVersion.platform]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[preReleaseVersion.platform]=%v", val))
		}
		if val, ok := args["filter[preReleaseVersion.version]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[preReleaseVersion.version]=%v", val))
		}
		if val, ok := args["filter[processingState]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[processingState]=%v", val))
		}
		if val, ok := args["filter[usesNonExemptEncryption]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[usesNonExemptEncryption]=%v", val))
		}
		if val, ok := args["filter[version]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[version]=%v", val))
		}
		if val, ok := args["filter[app]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[app]=%v", val))
		}
		if val, ok := args["filter[appStoreVersion]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[appStoreVersion]=%v", val))
		}
		if val, ok := args["filter[betaGroups]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[betaGroups]=%v", val))
		}
		if val, ok := args["filter[preReleaseVersion]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[preReleaseVersion]=%v", val))
		}
		if val, ok := args["filter[id]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[id]=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["fields[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[builds]=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[appEncryptionDeclarations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appEncryptionDeclarations]=%v", val))
		}
		if val, ok := args["fields[betaAppReviewSubmissions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaAppReviewSubmissions]=%v", val))
		}
		if val, ok := args["fields[buildBetaDetails]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[buildBetaDetails]=%v", val))
		}
		if val, ok := args["fields[buildIcons]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[buildIcons]=%v", val))
		}
		if val, ok := args["fields[perfPowerMetrics]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[perfPowerMetrics]=%v", val))
		}
		if val, ok := args["fields[preReleaseVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[preReleaseVersions]=%v", val))
		}
		if val, ok := args["fields[appStoreVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersions]=%v", val))
		}
		if val, ok := args["fields[diagnosticSignatures]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[diagnosticSignatures]=%v", val))
		}
		if val, ok := args["fields[betaTesters]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaTesters]=%v", val))
		}
		if val, ok := args["fields[betaBuildLocalizations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaBuildLocalizations]=%v", val))
		}
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["limit[betaBuildLocalizations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[betaBuildLocalizations]=%v", val))
		}
		if val, ok := args["limit[icons]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[icons]=%v", val))
		}
		if val, ok := args["limit[individualTesters]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[individualTesters]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/builds%s", cfg.BaseURL, queryString)
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
		var result models.BuildsResponse
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

func CreateBuilds_get_collectionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_builds",
		mcp.WithDescription(""),
		mcp.WithArray("filter[betaAppReviewSubmission.betaReviewState]", mcp.Description("filter by attribute 'betaAppReviewSubmission.betaReviewState'")),
		mcp.WithArray("filter[expired]", mcp.Description("filter by attribute 'expired'")),
		mcp.WithArray("filter[preReleaseVersion.platform]", mcp.Description("filter by attribute 'preReleaseVersion.platform'")),
		mcp.WithArray("filter[preReleaseVersion.version]", mcp.Description("filter by attribute 'preReleaseVersion.version'")),
		mcp.WithArray("filter[processingState]", mcp.Description("filter by attribute 'processingState'")),
		mcp.WithArray("filter[usesNonExemptEncryption]", mcp.Description("filter by attribute 'usesNonExemptEncryption'")),
		mcp.WithArray("filter[version]", mcp.Description("filter by attribute 'version'")),
		mcp.WithArray("filter[app]", mcp.Description("filter by id(s) of related 'app'")),
		mcp.WithArray("filter[appStoreVersion]", mcp.Description("filter by id(s) of related 'appStoreVersion'")),
		mcp.WithArray("filter[betaGroups]", mcp.Description("filter by id(s) of related 'betaGroups'")),
		mcp.WithArray("filter[preReleaseVersion]", mcp.Description("filter by id(s) of related 'preReleaseVersion'")),
		mcp.WithArray("filter[id]", mcp.Description("filter by id(s)")),
		mcp.WithArray("sort", mcp.Description("comma-separated list of sort expressions; resources will be sorted as specified")),
		mcp.WithArray("fields[builds]", mcp.Description("the fields to include for returned resources of type builds")),
		mcp.WithNumber("limit", mcp.Description("maximum resources per page")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[appEncryptionDeclarations]", mcp.Description("the fields to include for returned resources of type appEncryptionDeclarations")),
		mcp.WithArray("fields[betaAppReviewSubmissions]", mcp.Description("the fields to include for returned resources of type betaAppReviewSubmissions")),
		mcp.WithArray("fields[buildBetaDetails]", mcp.Description("the fields to include for returned resources of type buildBetaDetails")),
		mcp.WithArray("fields[buildIcons]", mcp.Description("the fields to include for returned resources of type buildIcons")),
		mcp.WithArray("fields[perfPowerMetrics]", mcp.Description("the fields to include for returned resources of type perfPowerMetrics")),
		mcp.WithArray("fields[preReleaseVersions]", mcp.Description("the fields to include for returned resources of type preReleaseVersions")),
		mcp.WithArray("fields[appStoreVersions]", mcp.Description("the fields to include for returned resources of type appStoreVersions")),
		mcp.WithArray("fields[diagnosticSignatures]", mcp.Description("the fields to include for returned resources of type diagnosticSignatures")),
		mcp.WithArray("fields[betaTesters]", mcp.Description("the fields to include for returned resources of type betaTesters")),
		mcp.WithArray("fields[betaBuildLocalizations]", mcp.Description("the fields to include for returned resources of type betaBuildLocalizations")),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithNumber("limit[betaBuildLocalizations]", mcp.Description("maximum number of related betaBuildLocalizations returned (when they are included)")),
		mcp.WithNumber("limit[icons]", mcp.Description("maximum number of related icons returned (when they are included)")),
		mcp.WithNumber("limit[individualTesters]", mcp.Description("maximum number of related individualTesters returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Builds_get_collectionHandler(cfg),
	}
}
