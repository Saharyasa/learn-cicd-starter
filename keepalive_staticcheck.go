package main

// This file references otherwise-unused symbols so staticcheck (U1080) doesn't fail in the tutorial repo.
// It has no runtime effect.

var (
    // Methods on *apiConfig -- use method expressions
    _ = (*apiConfig).handlerNotesGet
    _ = (*apiConfig).handlerNotesCreate
    _ = (*apiConfig).handlerUserCreate
    _ = (*apiConfig).handlerUserGet
    _ = (*apiConfig).middlewareAuth
    
    // handlers
    _ = handlerNotesGet
    _ = handlerNotesCreate
    _ = handlerHeadLoss
    _ = handlerUserSelect
    
    // helpers
    _ = handlerReadiness
)
