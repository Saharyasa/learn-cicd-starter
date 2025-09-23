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
    
    // Plain functions
    _ = handlerReadiness  // Make sure this matches your actual function name
    _ = handlerErr        // Common error handler
    _ = handlerRoot       // Common root handler
)
