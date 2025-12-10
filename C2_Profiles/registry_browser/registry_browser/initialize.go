package registry_browser

import (
	"fmt"

	"github.com/MythicMeta/MythicContainer/custombrowserstructs"
	"github.com/MythicMeta/MythicContainer/logging"
	"github.com/MythicMeta/MythicContainer/mythicrpc"
)

const version = "0.0.1"

func Initialize() {
	registry := custombrowserstructs.CustomBrowserDefinition{
		Name:          "registry_browser",
		Description:   fmt.Sprintf("Basic Windows registry browser"),
		Author:        "@its_a_feature_",
		SemVer:        version,
		Type:          custombrowserstructs.CUSTOMBROWSER_TYPE_FILE,
		PathSeparator: "\\",
		Columns: []custombrowserstructs.CustomBrowserTableColumn{
			{
				Key:       "type",
				Name:      "Type",
				Type:      "string",
				Width:     100,
				FillWidth: false,
			},
			{
				Key:       "value",
				Name:      "Value",
				Type:      "string",
				FillWidth: true,
			},
		},
		RowActions: []custombrowserstructs.CustomBrowserRowAction{
			{
				Name:           "Set Value",
				SupportsFile:   true,
				SupportsFolder: true,
				OpenDialog:     true,
				UIFeature:      "registry_browser:set_value",
				Icon:           "fa-pen-to-square",
				Color:          "warning",
			},
			{
				Name:           "Create Key",
				SupportsFile:   false,
				SupportsFolder: true,
				OpenDialog:     true,
				UIFeature:      "registry_browser:create_key",
				Icon:           "fa-folder-plus",
				Color:          "success",
			},
		},
		DefaultVisibleColumns:      []string{"Type", "Value"},
		IndicatePartialListingInUI: true,
		ShowCurrentPathAboveTable:  true,
		ExportFunction: func(message custombrowserstructs.ExportFunctionMessage) custombrowserstructs.ExportFunctionMessageResponse {
			response, err := mythicrpc.SendMythicRPCCustomBrowserSearch(mythicrpc.MythicRPCCustomBrowserSearchMessage{
				OperationID:            &message.OperationID,
				GetAllMatchingChildren: false,
				SearchCustomBrowser: mythicrpc.MythicRPCCustomBrowserSearchData{
					TreeType: "registry_browser",
					Host:     &message.Host,
					FullPath: &message.Path,
				},
			})
			if err != nil {
				logging.LogError(err, "failed to send custom browser search")
			} else if !response.Success {
				logging.LogInfo("failed to search", "error", response.Error)
			} else {
				logging.LogInfo("found custom browser data", "response", response)
			}
			return custombrowserstructs.ExportFunctionMessageResponse{
				Success: true,
			}
		},
	}
	custombrowserstructs.AllCustomBrowserData.Get(registry.Name).AddCustomBrowserDefinition(registry)
}
