package environment

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	ICMS                 float64
	PIS                  float64
	COFINS               float64
	CONT_SOCIAL          float64
	IRRF                 float64
	PROFIT_MARGIN        float64
	ADM_EXPENSES         float64
	INVESTMENT_36_MONTHS float64
	LABOUR_COST          float64
	RAW_MATERIAL_COST    float64

	APPLICATION_PORT    string
	APPLICATION_ADDRESS string

	CORS_URL string

	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DB_NAME  string
}

func (e *single) Setup() {

	e.APPLICATION_PORT = getenv("APPLICATION_PORT", "8080")
	e.APPLICATION_ADDRESS = getenv("APPLICATION_ADDRESS", "localhost")

	e.HOST = getenv("HOST", "localhost")
	e.PORT, _ = strconv.Atoi(getenv("PORT", "3306"))
	e.USER = getenv("USER", "root")
	e.PASSWORD = getenv("PASSWORD", "root")
	e.DB_NAME = getenv("DB_NAME", "db")

	e.CORS_URL = getenv("CORS_URL", "*")
}

func init() {
	env := GetInstance()
	env.Setup()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetenvFloat(key, fallback string) float64 {
	value := getenv(key, fallback)
	valueFloat, _ := strconv.ParseFloat(value, 32)
	return valueFloat
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
			singleInstance.Setup()
		} else {
			fmt.Println("Single instance already created.")
		}
	}

	return singleInstance
}
