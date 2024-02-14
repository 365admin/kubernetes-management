// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Database Discovery
---
*/
package cmds

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/365admin/kubernetes-management/execution"
	"github.com/365admin/kubernetes-management/schemas"
	"github.com/365admin/kubernetes-management/utils"
)

func BackupDiscoverPost(ctx context.Context, args []string) (*schemas.Databaseservices, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "kubernetes-management", "40 backup", "10 database discover.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}

	resultingFile := path.Join(utils.WorkDir("kubernetes-management"), "databaseservices.json")
	data, err := os.ReadFile(resultingFile)
	if err != nil {
		return nil, err
	}
	resultObject := schemas.Databaseservices{}
	err = json.Unmarshal(data, &resultObject)
	return &resultObject, nil

}