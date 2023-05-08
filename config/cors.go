package config

func AllowOrigins() string {

	return "http://localhost:3000, http://127.0.0.1:3000"
	
}

func AllowHeaders() string {
	
	return "Origin, Content-Type, Accept, Authorization, X-Requested-With, Cookie"
	
}

func AllowMethods() string {

	return "GET,POST,PUT,DELETE"
		
}
