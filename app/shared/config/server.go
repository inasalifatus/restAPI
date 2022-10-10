package config

type server struct {
	PORTHTTP string
}

var Server = server{
	PORTHTTP: ":8080",
}
