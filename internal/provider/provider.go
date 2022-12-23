package provider

func Init(dbConnStr string, jwtSecret string) {
	DBInit(dbConnStr)
	InitJWT(jwtSecret)
}
