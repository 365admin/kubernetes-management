// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Database Restore
---
*/
package endpoints

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/swaggest/usecase"

	"github.com/365admin/kubernetes-management/execution"
	"github.com/365admin/kubernetes-management/schemas"
	"github.com/365admin/kubernetes-management/utils"
)

func RestoreViewPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *schemas.Backupcontent) error {

		_, err := execution.ExecutePowerShell("john", "*", "kubernetes-management", "50-restore", "view.ps1", "")
		if err != nil {
			return err
		}

		resultingFile := path.Join(utils.WorkDir("kubernetes-management"), "backupcontent.json")
		data, err := os.ReadFile(resultingFile)
		if err != nil {
			return err
		}
		resultObject := schemas.Backupcontent{}
		err = json.Unmarshal(data, &resultObject)
		*output = resultObject
		return err

	})
	u.SetTitle("Database Restore")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("50-restore")
	return u
}