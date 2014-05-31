package main

func read_password() string {
	print("Password: ")
	passwd := stdin_read_password()

	print("Again: ")
	pwd_again := stdin_read_password()

	if passwd != pwd_again {
		panic("Passwords don't match")
	}

	return passwd
}
