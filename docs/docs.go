// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "hynduf@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/card/copy": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Copy Card To Deck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Copy Card To Deck",
                "parameters": [
                    {
                        "description": "Copy Card To Deck Request",
                        "name": "copy_card_to_deck_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CopyCardToDeckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CopyCardToDeckResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/card/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create New Card",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Create New Card",
                "parameters": [
                    {
                        "description": "Create Card Request",
                        "name": "create_card_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/card/review": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Review Cards",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Update Review Cards",
                "parameters": [
                    {
                        "description": "Update Review Cards Request",
                        "name": "update_review_cards_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateReviewCardsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateReviewCardsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/card/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Card Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Update Card Details",
                "parameters": [
                    {
                        "description": "Update Card Request",
                        "name": "update_card_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/copy": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Copy Deck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Copy Deck",
                "parameters": [
                    {
                        "description": "Copy Deck Request",
                        "name": "copy_deck_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CopyDeckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CopyDeckResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create New Deck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Create New Deck",
                "parameters": [
                    {
                        "description": "Create Deck Request",
                        "name": "create_deck_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateDeckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateDeckResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/review-cards": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Deck With Review Cards Of Logged In User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Get Deck With Review Cards Of Logged In User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.DeckWithReviewCards"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Deck Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Update Deck Details",
                "parameters": [
                    {
                        "description": "Update Deck Request",
                        "name": "update_deck_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateDeckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateDeckResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/signup": {
            "post": {
                "description": "Sign Up",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SignupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update User Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update User Details",
                "parameters": [
                    {
                        "description": "Update User Request",
                        "name": "update_user_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CopyCardToDeckRequest": {
            "type": "object",
            "required": [
                "card_id",
                "deck_id"
            ],
            "properties": {
                "card_id": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                }
            }
        },
        "dto.CopyCardToDeckResponse": {
            "type": "object",
            "properties": {
                "card": {
                    "$ref": "#/definitions/entity.Card"
                }
            }
        },
        "dto.CopyDeckRequest": {
            "type": "object",
            "required": [
                "deck_id"
            ],
            "properties": {
                "deck_id": {
                    "type": "string"
                }
            }
        },
        "dto.CopyDeckResponse": {
            "type": "object",
            "properties": {
                "deck": {
                    "$ref": "#/definitions/entity.DeckWithCards"
                },
                "deck_review": {
                    "$ref": "#/definitions/entity.DeckWithReviewCards"
                }
            }
        },
        "dto.CreateCardRequest": {
            "type": "object",
            "required": [
                "answer",
                "deck_id",
                "question",
                "wrong_answers"
            ],
            "properties": {
                "answer": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                },
                "question_img_label": {
                    "type": "string"
                },
                "question_img_url": {
                    "type": "string"
                },
                "wrong_answers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.CreateCardResponse": {
            "type": "object",
            "properties": {
                "card": {
                    "$ref": "#/definitions/entity.Card"
                }
            }
        },
        "dto.CreateDeckRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "description_img_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "total_cards": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateDeckResponse": {
            "type": "object",
            "properties": {
                "deck": {
                    "$ref": "#/definitions/entity.Deck"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.SignupResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateCardRequest": {
            "type": "object",
            "required": [
                "card_id"
            ],
            "properties": {
                "answer": {
                    "type": "string"
                },
                "card_id": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                },
                "question_img_label": {
                    "type": "string"
                },
                "question_img_url": {
                    "type": "string"
                },
                "wrong_answers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.UpdateCardResponse": {
            "type": "object",
            "properties": {
                "card": {
                    "$ref": "#/definitions/entity.Card"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "dto.UpdateDeckRequest": {
            "type": "object",
            "required": [
                "deck_id"
            ],
            "properties": {
                "cur_new_cards": {
                    "type": "integer"
                },
                "cur_review_cards": {
                    "type": "integer"
                },
                "deck_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "description_img_url": {
                    "type": "string"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "is_public": {
                    "type": "boolean"
                },
                "last_review": {
                    "type": "string"
                },
                "max_new_cards": {
                    "type": "integer"
                },
                "max_review_cards": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "total_learned_cards": {
                    "type": "integer"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateDeckResponse": {
            "type": "object",
            "properties": {
                "deck": {
                    "$ref": "#/definitions/entity.Deck"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "dto.UpdateReviewCardsRequest": {
            "type": "object",
            "required": [
                "card_ids",
                "deck_id",
                "is_correct"
            ],
            "properties": {
                "card_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "deck_id": {
                    "type": "string"
                },
                "is_correct": {
                    "type": "array",
                    "items": {
                        "type": "boolean"
                    }
                },
                "total_xp": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateReviewCardsResponse": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Card"
                    }
                },
                "num_blue_cards": {
                    "type": "integer"
                },
                "num_green_cards": {
                    "type": "integer"
                },
                "num_red_cards": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/entity.User"
                }
            }
        },
        "dto.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/entity.User"
                }
            }
        },
        "entity.Card": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "card_type": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "last_review": {
                    "type": "string"
                },
                "next_review": {
                    "type": "string"
                },
                "num_reviews": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                },
                "question_img_label": {
                    "type": "string"
                },
                "question_img_url": {
                    "type": "string"
                },
                "sm2_ef": {
                    "type": "number"
                },
                "sm2_i": {
                    "type": "integer"
                },
                "sm2_n": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "wrong_answers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entity.Deck": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "cur_new_cards": {
                    "type": "integer"
                },
                "cur_review_cards": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "description_img_url": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "is_public": {
                    "type": "boolean"
                },
                "last_review": {
                    "type": "string"
                },
                "max_new_cards": {
                    "type": "integer"
                },
                "max_review_cards": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "total_cards": {
                    "type": "integer"
                },
                "total_learned_cards": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "entity.DeckWithCards": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Card"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "cur_new_cards": {
                    "type": "integer"
                },
                "cur_review_cards": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "description_img_url": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "is_public": {
                    "type": "boolean"
                },
                "last_review": {
                    "type": "string"
                },
                "max_new_cards": {
                    "type": "integer"
                },
                "max_review_cards": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "num_blue_cards": {
                    "type": "integer"
                },
                "num_green_cards": {
                    "type": "integer"
                },
                "num_red_cards": {
                    "type": "integer"
                },
                "position": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "total_cards": {
                    "type": "integer"
                },
                "total_learned_cards": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "entity.DeckWithReviewCards": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Card"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "cur_new_cards": {
                    "type": "integer"
                },
                "cur_review_cards": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "description_img_url": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "is_public": {
                    "type": "boolean"
                },
                "last_review": {
                    "type": "string"
                },
                "max_new_cards": {
                    "type": "integer"
                },
                "max_review_cards": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "num_blue_cards": {
                    "type": "integer"
                },
                "num_green_cards": {
                    "type": "integer"
                },
                "num_red_cards": {
                    "type": "integer"
                },
                "position": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "total_cards": {
                    "type": "integer"
                },
                "total_learned_cards": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "last_streak": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "streak": {
                    "type": "integer"
                },
                "xp": {
                    "type": "integer"
                },
                "xp_to_level_up": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "VietCard Backend API",
	Description:      "Backend server for VietCard application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
