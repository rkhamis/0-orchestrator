package main

import (
	"encoding/json"
	"fmt"
	"github.com/g8os/go-client"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"strings"
)

func WriteError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err.Error())

	return
}

func Url(r *http.Request, p ...string) string {
	vars := mux.Vars(r)

	tail := path.Join(p...)
	return strings.TrimRight(fmt.Sprintf("/core0/%s/%s", vars["id"], tail), "/")
}

func ResultUrl(r *http.Request, job client.Job) string {
	return Url(r, "command", string(job))
}

func Options(cmd Command) []client.Option {
	var opts []client.Option
	if cmd.Id != "" {
		opts = append(opts, client.ID(cmd.Id))
	}

	if cmd.MaxRestart != 0 {
		opts = append(opts, client.MaxRestart(cmd.MaxRestart))
	}

	if cmd.MaxTime != 0 {
		opts = append(opts, client.MaxTime(cmd.MaxTime))
	}

	if cmd.Queue != "" {
		opts = append(opts, client.Queue(cmd.Queue))
	}

	if cmd.RecurringPeriod != 0 {
		opts = append(opts, client.RecurringPeriod(cmd.RecurringPeriod))
	}

	if cmd.StatsInterval != 0 {
		//TODO:
	}

	if cmd.Tags != "" {
		opts = append(opts, client.Tags(cmd.Tags))
	}

	return opts
}
