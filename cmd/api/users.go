package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/drizlye0/GopherSocial/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

var userCtx userKey = "user"

// GetUser godoc
//
//	@Summary		Fetches a user profile
//	@Description	Fetches a user profile by ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	store.User
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{id} [get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

// FollowUser godoc
//
//	@Summary		Follow a user
//	@Description	Follow a user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"User ID"
//	@Success		204	{string}	string	"User Followed"
//	@Failure		400	{object}	error	"User payload missing"
//	@Failure		404	{object}	error	"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/follow/{id} [put]
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followUser := getUserFromCtx(r)

	// TODO: Revert back when implement auth
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, followUser.ID, payload.UserID); err != nil {
		switch {
		case errors.Is(err, store.ErrConflict):
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

// UnfollowUser godoc
//
//	@Summary		Unfollow a user
//	@Description	Unfollow a user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"User ID"
//	@Success		204	{string}	string	"User unfollowed"
//	@Failure		400	{object}	error	"User payload missing"
//	@Failure		404	{object}	error	"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/unfollow/{id} [put]
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowUser := getUserFromCtx(r)

	// TODO: Revert back when implement auth
	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, unfollowUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}
func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "userID")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
				return
			default:
				app.internalServerError(w, r, err)
				return
			}
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	post, _ := r.Context().Value(userCtx).(*store.User)
	return post
}
