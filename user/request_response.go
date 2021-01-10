package user

type request struct {
	ID        int64
	Firstname string
	Lastname  string
	Surname   string
	Age       int32
	Gender    string
}

type response struct {
	User    User
	Message string
}
