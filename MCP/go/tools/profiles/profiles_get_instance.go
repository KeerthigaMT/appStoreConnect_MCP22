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

func Profiles_get_instanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["fields[profiles]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[profiles]=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[certificates]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[certificates]=%v", val))
		}
		if val, ok := args["fields[devices]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[devices]=%v", val))
		}
		if val, ok := args["fields[bundleIds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[bundleIds]=%v", val))
		}
		if val, ok := args["limit[certificates]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[certificates]=%v", val))
		}
		if val, ok := args["limit[devices]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[devices]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/profiles/%s%s", cfg.BaseURL, queryString)
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
		var result models.ProfileResponse
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

func CreateProfiles_get_instanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_profiles_id",
		mcp.WithDescription(""),
		mcp.WithArray("fields[profiles]", mcp.Description("the fields to include for returned resources of type profiles")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[certificates]", mcp.Description("the fields to include for returned resources of type certificates")),
		mcp.WithArray("fields[devices]", mcp.Description("the fields to include for returned resources of type devices")),
		mcp.WithArray("fields[bundleIds]", mcp.Description("the fields to include for returned resources of type bundleIds")),
		mcp.WithNumber("limit[certificates]", mcp.Description("maximum number of related certificates returned (when they are included)")),
		mcp.WithNumber("limit[devices]", mcp.Description("maximum number of related devices returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Profiles_get_instanceHandler(cfg),
	}
}
