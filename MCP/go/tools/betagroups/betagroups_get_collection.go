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

func Betagroups_get_collectionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["filter[isInternalGroup]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[isInternalGroup]=%v", val))
		}
		if val, ok := args["filter[name]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[name]=%v", val))
		}
		if val, ok := args["filter[publicLink]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[publicLink]=%v", val))
		}
		if val, ok := args["filter[publicLinkEnabled]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[publicLinkEnabled]=%v", val))
		}
		if val, ok := args["filter[publicLinkLimitEnabled]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[publicLinkLimitEnabled]=%v", val))
		}
		if val, ok := args["filter[app]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[app]=%v", val))
		}
		if val, ok := args["filter[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[builds]=%v", val))
		}
		if val, ok := args["filter[id]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[id]=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["fields[betaGroups]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaGroups]=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[builds]=%v", val))
		}
		if val, ok := args["fields[betaTesters]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaTesters]=%v", val))
		}
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["limit[betaTesters]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[betaTesters]=%v", val))
		}
		if val, ok := args["limit[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[builds]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/betaGroups%s", cfg.BaseURL, queryString)
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
		var result models.BetaGroupsResponse
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

func CreateBetagroups_get_collectionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_betaGroups",
		mcp.WithDescription(""),
		mcp.WithArray("filter[isInternalGroup]", mcp.Description("filter by attribute 'isInternalGroup'")),
		mcp.WithArray("filter[name]", mcp.Description("filter by attribute 'name'")),
		mcp.WithArray("filter[publicLink]", mcp.Description("filter by attribute 'publicLink'")),
		mcp.WithArray("filter[publicLinkEnabled]", mcp.Description("filter by attribute 'publicLinkEnabled'")),
		mcp.WithArray("filter[publicLinkLimitEnabled]", mcp.Description("filter by attribute 'publicLinkLimitEnabled'")),
		mcp.WithArray("filter[app]", mcp.Description("filter by id(s) of related 'app'")),
		mcp.WithArray("filter[builds]", mcp.Description("filter by id(s) of related 'builds'")),
		mcp.WithArray("filter[id]", mcp.Description("filter by id(s)")),
		mcp.WithArray("sort", mcp.Description("comma-separated list of sort expressions; resources will be sorted as specified")),
		mcp.WithArray("fields[betaGroups]", mcp.Description("the fields to include for returned resources of type betaGroups")),
		mcp.WithNumber("limit", mcp.Description("maximum resources per page")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[builds]", mcp.Description("the fields to include for returned resources of type builds")),
		mcp.WithArray("fields[betaTesters]", mcp.Description("the fields to include for returned resources of type betaTesters")),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithNumber("limit[betaTesters]", mcp.Description("maximum number of related betaTesters returned (when they are included)")),
		mcp.WithNumber("limit[builds]", mcp.Description("maximum number of related builds returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Betagroups_get_collectionHandler(cfg),
	}
}
