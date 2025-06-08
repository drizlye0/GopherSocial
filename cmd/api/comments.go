package main

import (
	"net/http"

	"github.com/drizlye0/GopherSocial/internal/store"
)

type createCommentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
}

// CreateComment godoc
//
//	@Summary		Create Comment
//	@Description	Create user comment in post
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		createCommentPayload	true	"Create Comment Payload"
//	@Success		201		{object}	store.Comment
//	@Failure		400		{object}	error	"Comment payload missing"
//	@Failure		404		{object}	error	"Post not found"
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id}/comment [post]
func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)
	var payload createCommentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// TODO: refactor when implements auth
	userID := 2

	comment := &store.Comment{
		PostID:  post.ID,
		UserID:  int64(userID),
		Content: payload.Content,
	}

	ctx := r.Context()
	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
