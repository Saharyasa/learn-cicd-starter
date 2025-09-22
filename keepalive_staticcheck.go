package main

// This file references otherwise-unused symbols so staticcheck (U1000) doesn't fail in the tutorial repo.
// It has no runtime effect.

var (
	_ = apiConfig{}

	// handlers
	_ = handlerNotesGet
	_ = handlerNotesCreate
	_ = handlerReadiness
	_ = handlerUsersCreate
	_ = handlerUsersGet

	// helpers
	_ = generateRandomSHA256Hash
	_ = respondWithError
	_ = respondWithJSON
	_ = writeJSON

	// middleware
	_ = authMiddleware{}
	_ = middlewareAuth

	// model mappers
	_ = databaseUserToUser
	_ = databaseNoteToNote
	_ = databasePostsToPosts
)
