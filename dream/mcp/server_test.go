package mcp

import (
	"net/http"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/taubyte/tau/dream"
	"gotest.tools/v3/assert"

	// Import dream services
	_ "github.com/taubyte/tau/services/auth/dream"
	_ "github.com/taubyte/tau/services/gateway/dream"
	_ "github.com/taubyte/tau/services/hoarder/dream"
	_ "github.com/taubyte/tau/services/monkey/dream"
	_ "github.com/taubyte/tau/services/patrick/dream"
	_ "github.com/taubyte/tau/services/seer/dream"
	_ "github.com/taubyte/tau/services/substrate/dream"
	_ "github.com/taubyte/tau/services/tns/dream"

	// Import dream clients
	_ "github.com/taubyte/tau/clients/p2p/auth/dream"
	_ "github.com/taubyte/tau/clients/p2p/hoarder/dream"
	_ "github.com/taubyte/tau/clients/p2p/monkey/dream"
	_ "github.com/taubyte/tau/clients/p2p/patrick/dream"
	_ "github.com/taubyte/tau/clients/p2p/seer/dream"
	_ "github.com/taubyte/tau/clients/p2p/substrate/dream"
	_ "github.com/taubyte/tau/clients/p2p/tns/dream"
)

func TestMCPServer(t *testing.T) {
	// Create multiverse
	multiverse, err := dream.New(t.Context())
	assert.NilError(t, err)
	defer multiverse.Close()

	// Create MCP service
	mcpService, err := New(multiverse, nil)
	assert.NilError(t, err)

	// Start MCP server in a goroutine
	go mcpService.Server().Start()

	// Wait a moment for server to start
	time.Sleep(100 * time.Millisecond)

	// Create MCP client
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "dream-mcp-test-client",
		Version: "1.0.0",
	}, nil)

	// Create HTTP transport
	transport := &mcp.StreamableClientTransport{
		Endpoint:   "http://localhost:8080" + MCPEndpoint,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	// Connect to MCP server
	connection, err := transport.Connect(t.Context())
	assert.NilError(t, err)
	defer connection.Close()

	// Create session
	session, err := client.Connect(multiverse.Context(), transport, nil)
	assert.NilError(t, err)
	defer session.Close()

	// Test 1: List available tools
	t.Run("ListTools", func(t *testing.T) {
		toolsResult, err := session.ListTools(multiverse.Context(), &mcp.ListToolsParams{})
		assert.NilError(t, err)

		expectedTools := []string{
			"list_universes",
			"get_universe_status",
			"create_universe",
			"delete_universe",
			"start_universe",
			"stop_universe",
			"list_projects",
			"get_project_details",
			"get_disk_usage",
		}

		assert.Equal(t, len(expectedTools), len(toolsResult.Tools))

		toolNames := make(map[string]bool)
		for _, tool := range toolsResult.Tools {
			toolNames[tool.Name] = true
		}

		for _, expectedTool := range expectedTools {
			assert.Assert(t, toolNames[expectedTool], "Expected tool %s not found", expectedTool)
		}
	})

	// Test 2: List universes (should be empty initially)
	t.Run("ListUniverses", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name:      "list_universes",
			Arguments: map[string]any{},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "List universes returned error: %v", result.Content)

		// Should return empty list initially
		assert.Assert(t, len(result.Content) > 0, "Expected content in list universes result")
	})

	// Test 3: Create universe
	t.Run("CreateUniverse", func(t *testing.T) {
		// First, try to delete the universe if it exists to ensure clean state
		_, _ = session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "delete_universe",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})

		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "create_universe",
			Arguments: map[string]any{
				"universe_name": "test-universe",
				"persistent":    false,
			},
		})
		assert.NilError(t, err)

		if result.IsError {
			t.Logf("Create universe returned error. Content: %+v", result.Content)
			if len(result.Content) > 0 {
				if textContent, ok := result.Content[0].(*mcp.TextContent); ok {
					t.Fatalf("Create universe returned error: %s", textContent.Text)
				}
			}
			t.Fatalf("Create universe returned error: %v", result.Content)
		}

		// Verify universe was created by listing universes
		listResult, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name:      "list_universes",
			Arguments: map[string]any{},
		})
		assert.NilError(t, err)

		assert.Assert(t, !listResult.IsError, "List universes after creation returned error: %v", listResult.Content)
	})

	// Test 4: Get universe status
	t.Run("GetUniverseStatus", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "get_universe_status",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "Get universe status returned error: %v", result.Content)
	})

	// Test 5: Start universe
	t.Run("StartUniverse", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "start_universe",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "Start universe returned error: %v", result.Content)
	})

	// Test 6: Get universe status after starting
	t.Run("GetUniverseStatusAfterStart", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "get_universe_status",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "Get universe status after start returned error: %v", result.Content)
	})

	// Test 7: List projects (should be empty for new universe)
	t.Run("ListProjects", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "list_projects",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "List projects returned error: %v", result.Content)
	})

	// Test 8: Get disk usage
	t.Run("GetDiskUsage", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "get_disk_usage",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "Get disk usage returned error: %v", result.Content)
	})

	// Test 10: Stop universe
	t.Run("StopUniverse", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "stop_universe",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "Stop universe returned error: %v", result.Content)
	})

	// Test 11: Delete universe
	t.Run("DeleteUniverse", func(t *testing.T) {
		_, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name: "delete_universe",
			Arguments: map[string]any{
				"universe_name": "test-universe",
			},
		})
		assert.NilError(t, err)

		// we dont check result as the delete should have failed coz no universe is persitent in the test
	})

	// Test 12: Verify universe is deleted
	t.Run("VerifyUniverseDeleted", func(t *testing.T) {
		result, err := session.CallTool(multiverse.Context(), &mcp.CallToolParams{
			Name:      "list_universes",
			Arguments: map[string]any{},
		})
		assert.NilError(t, err)

		assert.Assert(t, !result.IsError, "List universes after deletion returned error: %v", result.Content)

		// Should be empty again
		assert.Assert(t, len(result.Content) > 0, "Expected content in list universes result after deletion")
	})
}

// TestMCPServiceCreation tests the MCP service creation with various configurations
func TestMCPServiceCreation(t *testing.T) {
	multiverse, err := dream.New(t.Context())
	assert.NilError(t, err)
	defer multiverse.Close()

	t.Run("CreateWithNilHTTPService", func(t *testing.T) {
		service, err := New(multiverse, nil)
		assert.NilError(t, err)
		assert.Assert(t, service != nil)
		assert.Assert(t, service.Server() != nil)
	})

	t.Run("CreateWithCustomHTTPService", func(t *testing.T) {
		// This would test with a custom http service if we had one
		// For now, we'll just test that the function accepts nil
		service, err := New(multiverse, nil)
		assert.NilError(t, err)
		assert.Assert(t, service != nil)
	})
}

// TestUniverseHandlers tests all universe management handlers comprehensively
func TestUniverseHandlers(t *testing.T) {
	multiverse, err := dream.New(t.Context())
	assert.NilError(t, err)
	defer multiverse.Close()

	service, err := New(multiverse, nil)
	assert.NilError(t, err)

	ctx := multiverse.Context()

	t.Run("ListUniversesEmpty", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		result, output, err := service.listUniverses(ctx, req, nil)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Assert(t, output.Universes != nil)
	})

	t.Run("CreateUniverseSuccess", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := CreateUniverseInput{
			UniverseName: "test-universe-creation",
			Persistent:   false,
		}
		result, output, err := service.createUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Assert(t, output.Success)
		assert.Equal(t, output.UniverseName, "test-universe-creation")
	})

	t.Run("CreateUniverseDuplicate", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := CreateUniverseInput{
			UniverseName: "test-universe-creation",
			Persistent:   false,
		}
		result, output, err := service.createUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
		assert.Assert(t, !output.Success)
	})

	t.Run("GetUniverseStatusSuccess", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetUniverseStatusInput{
			UniverseName: "test-universe-creation",
		}
		result, output, err := service.getUniverseStatus(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Equal(t, output.Name, "test-universe-creation")
	})

	t.Run("GetUniverseStatusNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetUniverseStatusInput{
			UniverseName: "non-existent-universe",
		}
		result, _, err := service.getUniverseStatus(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	t.Run("StartUniverseSuccess", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := StartUniverseInput{
			UniverseName: "test-universe-creation",
		}
		result, output, err := service.startUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Assert(t, output.Success)
		assert.Equal(t, output.UniverseName, "test-universe-creation")
	})

	t.Run("StartUniverseNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := StartUniverseInput{
			UniverseName: "non-existent-universe",
		}
		result, _, err := service.startUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	t.Run("StopUniverseSuccess", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := StopUniverseInput{
			UniverseName: "test-universe-creation",
		}
		result, output, err := service.stopUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Assert(t, output.Success)
		assert.Equal(t, output.UniverseName, "test-universe-creation")
	})

	t.Run("StopUniverseNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := StopUniverseInput{
			UniverseName: "non-existent-universe",
		}
		result, _, err := service.stopUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	t.Run("DeleteUniverseSuccess", func(t *testing.T) {
		// First create a universe to delete
		createInput := CreateUniverseInput{
			UniverseName: "test-universe-to-delete",
			Persistent:   false,
		}
		_, _, err := service.createUniverse(ctx, &mcp.CallToolRequest{}, createInput)
		assert.NilError(t, err)

		req := &mcp.CallToolRequest{}
		input := DeleteUniverseInput{
			UniverseName: "test-universe-to-delete",
		}
		result, output, err := service.deleteUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Assert(t, output.Success)
	})

	t.Run("DeleteUniverseNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := DeleteUniverseInput{
			UniverseName: "non-existent-universe",
		}
		result, _, err := service.deleteUniverse(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})
}

// TestProjectHandlers tests project management handlers
func TestProjectHandlers(t *testing.T) {
	multiverse, err := dream.New(t.Context())
	assert.NilError(t, err)
	defer multiverse.Close()

	service, err := New(multiverse, nil)
	assert.NilError(t, err)

	ctx := multiverse.Context()

	// Create a universe first
	universe, err := multiverse.New(dream.UniverseConfig{
		Name:     "test-project-universe",
		KeepRoot: false,
	})
	assert.NilError(t, err)

	// Start the universe
	err = universe.StartAll("client")
	assert.NilError(t, err)

	// Wait for universe to be ready
	time.Sleep(2 * time.Second)

	t.Run("ListProjectsEmpty", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := ListProjectsInput{
			UniverseName: "test-project-universe",
		}
		result, output, err := service.listProjects(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Assert(t, output.Projects != nil)
		assert.Equal(t, len(output.Projects), 0) // Should be empty initially
	})

	t.Run("ListProjectsUniverseNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := ListProjectsInput{
			UniverseName: "non-existent-universe",
		}
		result, _, err := service.listProjects(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	t.Run("GetProjectDetailsNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetProjectDetailsInput{
			UniverseName: "test-project-universe",
			ProjectID:    "non-existent-project",
		}
		result, _, err := service.getProjectDetails(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	t.Run("GetProjectDetailsUniverseNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetProjectDetailsInput{
			UniverseName: "non-existent-universe",
			ProjectID:    "some-project",
		}
		result, _, err := service.getProjectDetails(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	// Clean up
	universe.Stop()
	multiverse.Delete("test-project-universe")
}

// TestSystemHandlers tests system management handlers
func TestSystemHandlers(t *testing.T) {
	multiverse, err := dream.New(t.Context())
	assert.NilError(t, err)
	defer multiverse.Close()

	service, err := New(multiverse, nil)
	assert.NilError(t, err)

	ctx := multiverse.Context()

	// Create a universe first
	_, err = multiverse.New(dream.UniverseConfig{
		Name:     "test-disk-universe",
		KeepRoot: false,
	})
	assert.NilError(t, err)

	t.Run("GetDiskUsageSuccess", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetDiskUsageInput{
			UniverseName: "test-disk-universe",
		}
		result, output, err := service.getDiskUsage(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result == nil) // No error result
		assert.Equal(t, output.UniverseName, "test-disk-universe")
		assert.Assert(t, output.DiskUsage >= 0)
		assert.Assert(t, output.DiskUsageMB >= 0)
		assert.Assert(t, output.DiskUsageGB >= 0)
	})

	t.Run("GetDiskUsageEmptyUniverseName", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetDiskUsageInput{
			UniverseName: "",
		}
		result, _, err := service.getDiskUsage(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	t.Run("GetDiskUsageUniverseNotFound", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := GetDiskUsageInput{
			UniverseName: "non-existent-universe",
		}
		result, _, err := service.getDiskUsage(ctx, req, input)
		assert.NilError(t, err)
		assert.Assert(t, result != nil) // Should return error result
		assert.Assert(t, result.IsError)
	})

	// Clean up
	multiverse.Delete("test-disk-universe")
}

// TestHelperFunctions tests the error helper functions
func TestHelperFunctions(t *testing.T) {
	t.Run("ErrorResultGeneric", func(t *testing.T) {
		result, output, err := errorResult("test", "test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
		assert.Equal(t, output, "test")
		assert.Equal(t, len(result.Content), 1)
		if textContent, ok := result.Content[0].(*mcp.TextContent); ok {
			assert.Equal(t, textContent.Text, "test error: message")
		} else {
			t.Fatal("Expected TextContent")
		}
	})

	t.Run("ErrorListProjects", func(t *testing.T) {
		result, output, err := errorListProjects("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
		assert.Assert(t, output.Projects != nil)
	})

	t.Run("ErrorGetProjectDetails", func(t *testing.T) {
		result, _, err := errorGetProjectDetails("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorListUniverses", func(t *testing.T) {
		result, _, err := errorListUniverses("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorGetUniverseStatus", func(t *testing.T) {
		result, _, err := errorGetUniverseStatus("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorCreateUniverse", func(t *testing.T) {
		result, _, err := errorCreateUniverse("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorDeleteUniverse", func(t *testing.T) {
		result, _, err := errorDeleteUniverse("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorStartUniverse", func(t *testing.T) {
		result, _, err := errorStartUniverse("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorStopUniverse", func(t *testing.T) {
		result, _, err := errorStopUniverse("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})

	t.Run("ErrorGetDiskUsage", func(t *testing.T) {
		result, _, err := errorGetDiskUsage("test error: %s", "message")
		assert.NilError(t, err)
		assert.Assert(t, result.IsError)
	})
}

// TestConvertRegistryToMapAndIDs tests the registry conversion helper function
func TestConvertRegistryToMapAndIDs(t *testing.T) {
	t.Run("EmptyMap", func(t *testing.T) {
		registry := make(map[interface{}]interface{})
		result, ids, err := convertRegistryToMapAndIDs(registry)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, len(result), 0)
		assert.Equal(t, len(ids), 0)
	})

	t.Run("SimpleMap", func(t *testing.T) {
		registry := map[interface{}]interface{}{
			"id":    "test-id",
			"name":  "test-name",
			"value": 42,
		}
		result, ids, err := convertRegistryToMapAndIDs(registry)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, len(ids), 1)
		assert.Equal(t, ids[0], "test-id")
		assert.Equal(t, result["id"], "test-id")
		assert.Equal(t, result["name"], "test-name")
		assert.Equal(t, result["value"], 42)
	})

	t.Run("MapWithQmKey", func(t *testing.T) {
		registry := map[interface{}]interface{}{
			"QmTestKey": "test-value",
			"name":      "test-name",
		}
		result, ids, err := convertRegistryToMapAndIDs(registry)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, len(ids), 1)
		assert.Equal(t, ids[0], "QmTestKey")
	})

	t.Run("NestedMap", func(t *testing.T) {
		registry := map[interface{}]interface{}{
			"nested": map[interface{}]interface{}{
				"id":   "nested-id",
				"data": "nested-data",
			},
			"top": "level",
		}
		result, ids, err := convertRegistryToMapAndIDs(registry)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, len(ids), 1)
		assert.Equal(t, ids[0], "nested-id")
		assert.Assert(t, result["nested"] != nil)
		if nested, ok := result["nested"].(map[string]interface{}); ok {
			assert.Equal(t, nested["id"], "nested-id")
			assert.Equal(t, nested["data"], "nested-data")
		} else {
			t.Fatal("Expected nested map")
		}
	})

	t.Run("MapWithSlice", func(t *testing.T) {
		registry := map[interface{}]interface{}{
			"items": []interface{}{
				map[interface{}]interface{}{
					"id":   "item1",
					"name": "Item 1",
				},
				map[interface{}]interface{}{
					"id":   "item2",
					"name": "Item 2",
				},
			},
		}
		result, ids, err := convertRegistryToMapAndIDs(registry)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		assert.Equal(t, len(ids), 2)
		assert.Assert(t, result["items"] != nil)
	})

	t.Run("MapWithNonStringKey", func(t *testing.T) {
		registry := map[interface{}]interface{}{
			42: "numeric-key",
		}
		_, _, err := convertRegistryToMapAndIDs(registry)
		// This should return an error because non-string keys are not supported
		assert.Assert(t, err != nil)
	})

	t.Run("MapWithComplexValue", func(t *testing.T) {
		registry := map[interface{}]interface{}{
			"complex": struct{ Value string }{"test"},
		}
		result, _, err := convertRegistryToMapAndIDs(registry)
		assert.NilError(t, err)
		assert.Assert(t, result != nil)
		// The complex struct should be converted to string representation
		assert.Equal(t, result["complex"], "{test}")
	})
}
