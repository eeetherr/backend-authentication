package test

func RunDevTest() {
	//db := database.GetDB()
	//
	//log.Println("⏳ Running AutoMigrate...")
	//if err := db.AutoMigrate(&users.User{}); err != nil {
	//	log.Fatal("❌ Migration failed:", err)
	//}
	//log.Println("✅ AutoMigrate completed")
	//
	//userRepo := repositories.NewUserRepository(db)
	//
	//// Create user
	//newUser := &users.User{
	//	Email:    "test@example.com",
	//	Password: "hashedpassword",
	//	Name:     "John",
	//}
	//
	//err := userRepo.CreateUser(newUser)
	//if err != nil {
	//	log.Fatal("Error creating user:", err)
	//}
	//log.Println("User created with ID:", newUser.ID)
	//
	//// Fetch user
	//user, err := userRepo.GetUserByEmailAndPhoneNumber("test@example.com")
	//if err != nil {
	//	log.Fatal("Error getting user:", err)
	//}
	//if user != nil {
	//	log.Printf("User found: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)
	//} else {
	//	log.Println("User not found")
	//}
}
