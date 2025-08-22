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

func Bundleids_get_collectionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter[identifier]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[identifier]=%v", val))
		}
		if val, ok := args["filter[name]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[name]=%v", val))
		}
		if val, ok := args["filter[platform]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[platform]=%v", val))
		}
		if val, ok := args["filter[seedId]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[seedId]=%v", val))
		}
		if val, ok := args["filter[id]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[id]=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["fields[bundleIds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[bundleIds]=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[bundleIdCapabilities]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[bundleIdCapabilities]=%v", val))
		}
		if val, ok := args["fields[profiles]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[profiles]=%v", val))
		}
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["limit[bundleIdCapabilities]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[bundleIdCapabilities]=%v", val))
		}
		if val, ok := args["limit[profiles]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[profiles]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/bundleIds%s", cfg.BaseURL, queryString)
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
		var result models.BundleIdsResponse
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

func CreateBundleids_get_collectionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_bundleIds",
		mcp.WithDescription(""),
		mcp.WithArray("filter[identifier]", mcp.Description("filter by attribute 'identifier'")),
		mcp.WithArray("filter[name]", mcp.Description("filter by attribute 'name'")),
		mcp.WithArray("filter[platform]", mcp.Description("filter by attribute 'platform'")),
		mcp.WithArray("filter[seedId]", mcp.Description("filter by attribute 'seedId'")),
		mcp.WithArray("filter[id]", mcp.Description("filter by id(s)")),
		mcp.WithArray("sort", mcp.Description("comma-separated list of sort expressions; resources will be sorted as specified")),
		mcp.WithArray("fields[bundleIds]", mcp.Description("the fields to include for returned resources of type bundleIds")),
		mcp.WithNumber("limit", mcp.Description("maximum resources per page")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[bundleIdCapabilities]", mcp.Description("the fields to include for returned resources of type bundleIdCapabilities")),
		mcp.WithArray("fields[profiles]", mcp.Description("the fields to include for returned resources of type profiles")),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithNumber("limit[bundleIdCapabilities]", mcp.Description("maximum number of related bundleIdCapabilities returned (when they are included)")),
		mcp.WithNumber("limit[profiles]", mcp.Description("maximum number of related profiles returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Bundleids_get_collectionHandler(cfg),
	}
}
