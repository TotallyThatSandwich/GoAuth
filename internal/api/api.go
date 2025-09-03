package api

import (
	"context"
	"net/http"
	"log/slog"
	"fmt"

	"encoding/json"

	"github.com/TotallyThatSandwich/GoAuth/internal/repository"
	"github.com/TotallyThatSandwich/GoAuth/internal/sqlc"
)

type APIServer struct {
	addr string
}

func New(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run(ctx context.Context, repo *repository.UserRepository, rootAddr string) error {
	router := http.NewServeMux()

	router.HandleFunc(fmt.Sprintf("GET %s/getUser", rootAddr), func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		hashedPass := r.URL.Query().Get("hashedPass")

		if username == "" || hashedPass == "" {
			http.Error(w, "Missing query parameters", http.StatusBadRequest)
			return
		}

		user, err := repo.GetUser(ctx, username, hashedPass)
		
		if err == nil {
			userJSON, err := json.Marshal(user)
			if err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write(userJSON)
			} else {
				w.Write([]byte(fmt.Sprintf("An error occurred while retreving user, username: %s hashedPass: %s", username, hashedPass)))
			}
		} else {
			w.Write([]byte(fmt.Sprintf("could not find user, username: %s hashedPass: %s", username, hashedPass)))
		}
		
	})

	router.HandleFunc(fmt.Sprintf("POST %s/createUser", rootAddr), func(w http.ResponseWriter, r *http.Request) {
		
		username := r.URL.Query().Get("username")
		hashedPass := r.URL.Query().Get("hashedPass")

		if username == "" || hashedPass == "" {
			http.Error(w, "Missing query parameters", http.StatusBadRequest)
			return
		}

		user, err := repo.CreateUser(ctx, username, hashedPass)
		
		if err == nil {
			userJSON, err := json.Marshal(user)
			if err == nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write(userJSON)
				repo.Commit(ctx)
			} else {
				w.Write([]byte(fmt.Sprintf("An error occurred while creating user, username: %s hashedPass: %s", username, hashedPass)))
			}
		} else {
			w.Write([]byte(fmt.Sprintf("An error occurred while creating user, username: %s hashedPass: %s", username, hashedPass)))
		}	
	})
	

	router.HandleFunc(fmt.Sprintf("DELETE %s/deleteUser", rootAddr), func(w http.ResponseWriter, r *http.Request) {
		
		var user sqlc.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err == nil {
			err := repo.DeleteUser(ctx, user)
		
			if err == nil {
				w.Write([]byte(fmt.Sprintf("Deleted User")))
				repo.Commit(ctx)
			} else {
				w.Write([]byte(fmt.Sprintf("An error occurred while deleting user, username: %s hashedPass: %s", user.Username, user.HashedPassword)))
			}
		}
	})

	router.HandleFunc(fmt.Sprintf("PUT %s/updateUser", rootAddr), func(w http.ResponseWriter, r *http.Request) {

		var user sqlc.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err == nil {
			user, err := repo.UpdateUser(ctx, user)
			if err == nil {
				userJSON, err := json.Marshal(user)
				if err == nil {
					w.Header().Set("Content-Type", "application/json")
					w.Write(userJSON)
					repo.Commit(ctx)
				}
				repo.Commit(ctx)
			} else {
				w.Write([]byte(fmt.Sprintf("An error occurred while updating user, username: %s hashedPass: %s", user.Username, user.HashedPassword)))
			}
		}
	})



	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	slog.Info("API server started.", "address", s.addr)

	return server.ListenAndServe()
}
