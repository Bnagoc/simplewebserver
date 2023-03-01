package api

// Full API Handler initialization file

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simplewebserver/internal/app/middleware"
	"simplewebserver/internal/app/models"
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

// Helper for generate messages
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHandlers(writer http.ResponseWriter) {
	writer.Header().Set("Content-type", "application/json")
}

// Return an actual list of all articles
func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	// Headers initialization
	initHandlers(writer)
	// Logging the beginning of request processing
	api.logger.Info("Get All Articles GET /api/v1/articles")
	// Return anything from DB
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles with accessing to DB",
			IsError:    true,
		}

		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("Post Article POST /api/v1/articles")

	var article models.Article

	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new article: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles with accessing to DB",
			IsError:    true,
		}

		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("Get articles by ID /api/v1/articles{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid id",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing DB table (articles). Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing DB",
			IsError:    true,
		}

		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find article with this ID")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with this ID does not exist in DB",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("Delete Article by Id DELETE /api/v1/articles/{id}")

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid id",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing DB table (articles). Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing DB",
			IsError:    true,
		}

		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find article with this ID")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with this ID does not exist in DB",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles with deleting element from DB table (articles) by id. Err: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing DB",
			IsError:    true,
		}

		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with ID %d deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("Post User Register POST /api/vi/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// Find user with this login
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing DB table (articles). Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing DB",
			IsError:    true,
		}

		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// Does not register if user already exists
	if ok {
		api.logger.Info("User with that ID already exists", err)
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login already exists in DB",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// Append User into DP
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles while accessing DB table (articles). Err: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing DB",
			IsError:    true,
		}

		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login: %s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHandlers(writer)
	api.logger.Info("Post to Auth POST api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	if err != nil {
		api.logger.Info("Invalid json recieved from client", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	userInDB, ok, err := api.storage.User().FindByLogin(userFromJson.Login)
	if err != nil {
		api.logger.Info("Can not make user search in DB: ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles with accessing to DB",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("User with that login does not exists", err)
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exists in DB. Need register first",
			IsError:    true,
		}

		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if userInDB.Password != userFromJson.Password {
		api.logger.Info("Invalid credentials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// Now we can give token to user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["amdin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Come back later.",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
	}

	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
