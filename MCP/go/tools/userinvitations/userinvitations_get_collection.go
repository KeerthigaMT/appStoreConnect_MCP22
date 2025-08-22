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

func Userinvitations_get_collectionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter[email]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[email]=%v", val))
		}
		if val, ok := args["filter[roles]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[roles]=%v", val))
		}
		if val, ok := args["filter[visibleApps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[visibleApps]=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["fields[userInvitations]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[userInvitations]=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["limit[visibleApps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[visibleApps]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/userInvitations%s", cfg.BaseURL, queryString)
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
		var result models.UserInvitationsResponse
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

func CreateUserinvitations_get_collectionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_userInvitations",
		mcp.WithDescription(""),
		mcp.WithArray("filter[email]", mcp.Description("filter by attribute 'email'")),
		mcp.WithArray("filter[roles]", mcp.Description("filter by attribute 'roles'")),
		mcp.WithArray("filter[visibleApps]", mcp.Description("filter by id(s) of related 'visibleApps'")),
		mcp.WithArray("sort", mcp.Description("comma-separated list of sort expressions; resources will be sorted as specified")),
		mcp.WithArray("fields[userInvitations]", mcp.Description("the fields to include for returned resources of type userInvitations")),
		mcp.WithNumber("limit", mcp.Description("maximum resources per page")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithNumber("limit[visibleApps]", mcp.Description("maximum number of related visibleApps returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Userinvitations_get_collectionHandler(cfg),
	}
}
