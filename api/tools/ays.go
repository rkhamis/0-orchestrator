package tools

import (
	"encoding/json"
	"fmt"
	"net/http"

	ays "github.com/g8os/grid/api/ays-client"
	client "github.com/g8os/grid/ays-client"
)

var (
	ayscl *ays.AtYourServiceAPI
)

func SetAYSClient(client ays.AtYourServiceAPI) {
	ayscl = client
}

// ExecuteBlueprint runs ays operations needed to run blueprints
func ExecuteBlueprint(w http.ResponseWriter, repoName, blueprintName string, blueprint map[string]interface{}) error {

	if err := createBlueprint(w, repoName, blueprintName, blueprint); err != nil {
		return err
	}

	if err := executeBlueprint(w, blueprintName, repoName); err != nil {
		archiveBlueprint(w, blueprintName, repoName)
		return err
	}

	if err := runRepo(w, repoName); err != nil {
		archiveBlueprint(w, blueprintName, repoName)
		return err
	}

	return archiveBlueprint(w, blueprintName, repoName)
}

func createRepo(w http.ResponseWriter, repoName string) error {
	// Create ays repo
	var repo client.AysRepositoryPostReqBody
	repo.Name = repoName
	repo.Git_url = "https://github.com/g8os/test"

	_, resp, err := ayscl.Ays.CreateRepository(repo, nil, nil)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return err
	}
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		w.WriteHeader(resp.StatusCode)
		return err
	}
	return nil
}

func createBlueprint(w http.ResponseWriter, repoName string, name string, bp map[string]interface{}) error {
	data, err := json.Marshal(bp)
	if err != nil {
		return err
	}

	blueprint := client.Blueprint{
		Content: data,
		Name:    name,
	}

	_, resp, err := ayscl.Ays.CreateBlueprint(repoName, blueprint, nil, nil)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return err
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		w.WriteHeader(resp.StatusCode)
		return fmt.Errorf("bad response code %+v", resp.StatusCode)
	}

	return nil
}

func executeBlueprint(w http.ResponseWriter, blueprintName string, repoName string) error {

	resp, err := ayscl.Ays.ExecuteBlueprint(blueprintName, repoName, nil, nil)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return err
	}
	return nil
}

func runRepo(w http.ResponseWriter, repoName string) error {

	_, resp, err := ayscl.Ays.CreateRun(repoName, nil, nil)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return err
	}
	return nil
}

func archiveBlueprint(w http.ResponseWriter, blueprintName string, repoName string) error {

	resp, err := ayscl.Ays.ArchiveBlueprint(blueprintName, repoName, nil, nil)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return err
	}
	return nil
}
