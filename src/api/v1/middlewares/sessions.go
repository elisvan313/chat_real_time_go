package middlewares

// func ValidateSession(store sessions.Store) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		sessionID := session.Get("sessionID")
// 		if sessionID == nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesión no válida"})
// 			c.Abort()
// 			return
// 		}

// 		// Verificar la sesión en la base de datos
// 		dbSessionID := sessionID.(string)
// 		//db := store.(*postgres.Store).DB

// 		var valid bool
// 		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM sessions WHERE session_id = $1)", dbSessionID).Scan(&valid)
// 		if err != nil || !valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesión no válida"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()

// 	}
// }
