package main

// References otherwise-unused symbols so staticcheck (U1000) doesn't fail in the tutorial repo.
// No runtime effect.

var (
	// Methods on *apiConfig — use method expressions
	_ = (*apiConfig).handlerNotesGet
	_ = (*apiConfig).handlerNotesCreate
	_ = (*apiConfig).handlerUsersCreate
	_ = (*apiConfig).handlerUsersGet
	_ = (*apiConfig).middlewareAuth

	// Plain functions
	_ = handleReadiness
	_ = generateRandomSHA256Hash
	_ = respondWithError
	_ = respondWithJSON
	_ = writeJSON

	// Types / zero values
	_ = authMiddleware{}

	// Model mapping helpers
	_ = databaseUserToUser
	_ = databaseNoteToNote
	_ = databasePostsToPosts
)
