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

func Betatesters_get_instanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["fields[betaTesters]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaTesters]=%v", val))
		}
		if val, ok := args["include"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include=%v", val))
		}
		if val, ok := args["fields[betaGroups]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[betaGroups]=%v", val))
		}
		if val, ok := args["fields[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[builds]=%v", val))
		}
		if val, ok := args["fields[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fields[apps]=%v", val))
		}
		if val, ok := args["limit[apps]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[apps]=%v", val))
		}
		if val, ok := args["limit[betaGroups]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[betaGroups]=%v", val))
		}
		if val, ok := args["limit[builds]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit[builds]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/betaTesters/%s%s", cfg.BaseURL, queryString)
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
		var result models.BetaTesterResponse
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

func CreateBetatesters_get_instanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_betaTesters_id",
		mcp.WithDescription(""),
		mcp.WithArray("fields[betaTesters]", mcp.Description("the fields to include for returned resources of type betaTesters")),
		mcp.WithArray("include", mcp.Description("comma-separated list of relationships to include")),
		mcp.WithArray("fields[betaGroups]", mcp.Description("the fields to include for returned resources of type betaGroups")),
		mcp.WithArray("fields[builds]", mcp.Description("the fields to include for returned resources of type builds")),
		mcp.WithArray("fields[apps]", mcp.Description("the fields to include for returned resources of type apps")),
		mcp.WithNumber("limit[apps]", mcp.Description("maximum number of related apps returned (when they are included)")),
		mcp.WithNumber("limit[betaGroups]", mcp.Description("maximum number of related betaGroups returned (when they are included)")),
		mcp.WithNumber("limit[builds]", mcp.Description("maximum number of related builds returned (when they are included)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Betatesters_get_instanceHandler(cfg),
	}
}
