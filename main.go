package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NopparootSuree/Hexagonal-Architechture-GO/handler"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/logs"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/repository"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/service"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customer", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id}", customerHandler.GetCustomer).Methods(http.MethodGet)

	logs.Info("Server start in port : " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dsn := fmt.Sprintf("mongodb+srv://%v:%v@mongodb.xftca8p.mongodb.net/%v", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.database"))
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		logs.Error(err)
	}

	return db
}
