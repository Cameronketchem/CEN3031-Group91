package routes

import (
	"database/sql"
	"github.com/Cameronketchem/CEN3031-Group91/server/auth"
	"github.com/Cameronketchem/CEN3031-Group91/server/blockchain"
	"github.com/Cameronketchem/CEN3031-Group91/server/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authentication is accomplished by requesting a random nonce assosciated
// with a user address. The user signs this nonce using their browser crypto
// extension (i.e metamask). If the signature can be verified, then user
// has proven that they own their wallet address and can be logged in.

// Returns the nonce assosciated with an address, or an error if the
// address is not registered.
func getNonce(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	addr := c.Params.ByName("addr")
	user, err := store.QueryUserByAddr(addr)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server failed to fetch response"})
		return
	}

	// User must be created
	if err == sql.ErrNoRows {
		c.JSON(http.StatusForbidden, gin.H{"error": "No account with this address"})
		return
	}

	c.JSON(http.StatusOK, user.Nonce)
}

// isLoggedIn is middleware for checking whether the current session is logged in.
func isLoggedIn(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	if c.GetHeader("auth") == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No auth header"})
		return
	}

	jwt := c.GetHeader("auth")
	address, userId, err := auth.VerifyJWT(c.Value("jwt-secret").(string), jwt)

	// Return unauthorized error.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not logged in"})
		return
	}

	// Pass claims to next middleware.
	c.Set("jwt-address", address)
	c.Set("jwt-userid", userId)
}

// SignUp route input schema
type signupInput struct {
	Addr       string `json:"addr" binding:"required"`
	ProfilePic string `json:"profile_pic_url" binding:"required"`
	Bio        string `json:"bio" binding:"required"`
	Signature  string `json:"sig" binding:"required"`
}

// Creates a new user
func postSignUp(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Validate input
	var input signupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure wallet address doesn't exist
	_, err := store.QueryUserByAddr(input.Addr)
	if err != sql.ErrNoRows {
		c.JSON(http.StatusForbidden, gin.H{"error": "User with that address already exists"})
		return
	}

	// Verify signature
	if err := auth.VerifySignature(input.Addr, input.Signature, "ADDR:"+input.Addr); err != nil {
		// VerifySignature err response is human-friendly
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Generate new account nonce
	nonce, err := auth.Nonce(20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server failed to fetch response"})
		return
	}

	// Insert User
	user := data.User{0, input.Addr, nonce, input.ProfilePic, input.Bio}
	store.InsertUser(user)
	c.JSON(http.StatusOK, gin.H{"success": "Account created"})
}

// SignIn route input schema
type signInInput struct {
	Addr      string `json:"addr" binding:"required"`
	Signature string `json:"sig" binding:"required"`
}

// Authenticates a user
func postLogIn(store *data.Store, cont *blockchain.Executor, c *gin.Context) {
	// Validate input
	var input signInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve nonce
	user, err := store.QueryUserByAddr(input.Addr)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server failed to fetch response"})
		return
	}

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find user with that address"})
		return
	}

	// Verify signature
	if err := auth.VerifySignature(input.Addr, input.Signature, user.Nonce); err != nil {
		// VerifySignature err response is human-friendly
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// --- Login: successful ---

	// Change user nonce
	newNonce, err := auth.Nonce(20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server failed to fetch response"})
		return
	}

	store.UpdateUserNonce(user.UserID, newNonce)

	// Create JWT
	jwt, err := auth.NewJWT(c.Value("jwt-secret").(string), input.Addr, user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server failed to fetch response"})
		return
	}

	c.Writer.Header().Set("auth", jwt)
	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
