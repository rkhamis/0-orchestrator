package tools

import (
	"net/http"

	ays "github.com/g8os/grid/api/ays-client"
)

var (
	ayscl *ays.AtYourServiceAPI
)

func SetAYSClient(client *ays.AtYourServiceAPI) {
	ayscl = client
}

// ExecuteBlueprint runs ays operations needed to run blueprints
func ExecuteBlueprint(repoName, blueprintName string, blueprint map[string]interface{}) error {

	if err := createBlueprint(repoName, blueprintName, blueprint); err != nil {
		return err
	}

	if err := executeBlueprint(blueprintName, repoName); err != nil {
		archiveBlueprint(blueprintName, repoName)
		return err
	}

	if err := runRepo(repoName); err != nil {
		archiveBlueprint(blueprintName, repoName)
		return err
	}

	return archiveBlueprint(blueprintName, repoName)
}

func createBlueprint(repoName string, name string, bp map[string]interface{}) error {
	blueprint := ays.Blueprint{
		Content: bp,
		Name:    name,
	}

	_, resp, err := ayscl.Ays.CreateBlueprint(repoName, blueprint, nil, nil)
	if err != nil {
		return NewHTTPError(resp, err.Error())
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		return NewHTTPError(resp, resp.Status)
	}

	return nil
}

func executeBlueprint(blueprintName string, repoName string) error {

	resp, err := ayscl.Ays.ExecuteBlueprint(blueprintName, repoName, nil, nil)
	if err != nil {
		return NewHTTPError(resp, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return NewHTTPError(resp, resp.Status)
	}
	return nil
}

func runRepo(repoName string) error {

	_, resp, err := ayscl.Ays.CreateRun(repoName, nil, nil)
	if err != nil {
		return NewHTTPError(resp, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return NewHTTPError(resp, resp.Status)
	}
	return nil
}

func archiveBlueprint(blueprintName string, repoName string) error {

	resp, err := ayscl.Ays.ArchiveBlueprint(blueprintName, repoName, nil, nil)
	if err != nil {
		return NewHTTPError(resp, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return NewHTTPError(resp, resp.Status)
	}
	return nil
}
