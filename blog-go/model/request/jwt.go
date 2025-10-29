package request

import (
	"blog-go/model/appTypes"

	"github.com/gofrs/uuid"
	jwt "github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims defines custom JWT claims.
// It embeds BaseClaims and standard JWT registered claims.
type JwtCustomClaims struct {
	BaseClaims           // Basic user claims: user ID, UUID, and role ID.
	jwt.RegisteredClaims // Standard JWT registered claims (e.g. exp, iss, sub).
}

// JwtCustomRefreshClaims defines custom claims for refresh tokens.
// It includes the user ID and standard registered claims.
type JwtCustomRefreshClaims struct {
	UserID               uint // User ID used for refresh token authentication.
	jwt.RegisteredClaims      // Standard JWT registered claims.
}

// BaseClaims stores basic user information used as part of the JWT claims.
type BaseClaims struct {
	UserID uint            // User ID, unique identifier for the user.
	UUID   uuid.UUID       // User UUID, globally unique identifier.
	RoleID appTypes.RoleID // User role ID, representing permission level.
}
