package api

//func GetClassInfo(c *gin.Context) {
//	var class ClassesStruct
//	if err := c.ShouldBindJSON(&class); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}

//if class.Password == "" {
//	log.Println(user.Email + "Password not undefined. Creating new password ...")
//	user.Password = core.PasswordGenerate(14)
//}

//	output := methods.CreateUser(user.Email, user.Password)
//	c.Data(http.StatusOK, "application/json; charset=utf-8", output)
//
//	emailResponse := methods.SendEmail(user.Email, user.Password)
//	c.Data(http.StatusOK, "application/json; charset=utf-8", emailResponse)
//}
