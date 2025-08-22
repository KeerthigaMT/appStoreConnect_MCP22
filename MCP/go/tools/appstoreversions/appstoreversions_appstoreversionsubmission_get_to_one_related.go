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

func Appstoreversions_appstoreversionsubmission_get_to_one_relatedHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["fields[appStoreVersions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersions]=%v", val))
		}
		if val, ok := args["fields[appStoreVersionSubmissions]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[appStoreVersionSubmissions]=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/appStoreVersions/%s/appStoreVersionSubmission%s", cfg.BaseURL, queryString)
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
		var result models.AppStoreVersionSubmissionResponse
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

func CreateAppstoreversions_appstoreversionsubmission_get_to_one_relatedTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_appStoreVersions_id_appStoreVersionSubmission",
		mcp.WithDescription(""),
		mcp.WithArray("fields[appStoreVersions]", mcp.Description("the fields to include for returned resources of type appStoreVersions")),
		mcp.WithArray("fields[appStoreVersionSubmissions]", mcp.Description("the fields to include for returned resources of type appStoreVersionSubmissions")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Appstoreversions_appstoreversionsubmission_get_to_one_relatedHandler(cfg),
	}
}
