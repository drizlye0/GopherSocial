package main

import "net/http"

// CheckHealth godoc
//
//	@Summary		Check server health
//	@Description	Check server status
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) checkHealth(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}

}
