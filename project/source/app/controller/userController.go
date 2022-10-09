package controller

//func (c *Controller) SignUp(ctx *gin.Context) bool {
//
//	var body struct {
//		Username string
//		Password string
//	}
//
//	if ctx.BindJSON(&body) != nil {
//		fmt.Errorf("400", gin.H{
//			"error": "Failed to read body",
//		})
//		return false
//	}
//
//	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
//	if err != nil {
//		fmt.Errorf("400", gin.H{
//			"error": "Failed to hash password",
//		})
//		return false
//	}
//
//	user := entity.User{0, body.Username, string(hash)}
//	return c.Repository.CreateUser(user)
//}
//
//func (c *Controller) Login(ctx *gin.Context) string {
//	var body struct {
//		Username string
//		Password string
//	}
//
//	if ctx.BindJSON(&body) != nil {
//		fmt.Errorf("400", gin.H{
//			"error": "Failed to read body",
//		})
//		return ""
//	}
//	user := entity.User{0, body.Username, body.Password}
//	user = c.Repository.FindUserID(user)
//
//	if user.UserID == 0 {
//		return ""
//	}
//
//	user = c.Repository.FindUserPassword(user)
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
//	if err != nil {
//		fmt.Errorf("400", gin.H{
//			"error": "Not right password",
//		})
//		return ""
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"sub": user.UserID,
//		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
//	})
//
//	tokenString, err := token.SignedString([]byte("sdff3234sr2134rewt3t2sra"))
//	ctx.SetSameSite(http.SameSiteLaxMode)
//	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
//	if err != nil {
//		fmt.Errorf("400", gin.H{
//			"error": "Not right password",
//		})
//		return ""
//	}
//
//	return tokenString
//}
