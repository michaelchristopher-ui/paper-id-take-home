package main

import (
	"context"
	"flag"
	"fmt"
	disbursementHandlerPkg "paperid-entry-task/handler/http/impl/disbursementhandler"
	"paperid-entry-task/internal/conf"
	"paperid-entry-task/internal/pkg/adapters/accountrepo"
	accountRepoPkg "paperid-entry-task/internal/pkg/repository/account"
	journalRepoPkg "paperid-entry-task/internal/pkg/repository/journal"
	disbursementSvcPkg "paperid-entry-task/internal/pkg/service/disbursement"
	"paperid-entry-task/internal/pkg/transport"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	//Initialize server
	cfgPath := flag.String("configpath", "config.yaml", "path to config file")
	flag.Parse()

	err := conf.Init(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	//Initialize database
	sqliteDB, err := gorm.Open(sqlite.Open(conf.GetConfig().SQLite.DBPath), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//Initialize Repos
	journalRepo := journalRepoPkg.New(sqliteDB)
	accountRepo := accountRepoPkg.New(sqliteDB)

	//Initialize Dummy Values
	accountRepo.Create(context.Background(), nil, accountrepo.Create{
		ID:          1,
		AccountName: "1",
		Balance:     3000,
		IsActive:    true,
	})
	accountRepo.Create(context.Background(), nil, accountrepo.Create{
		ID:          2,
		AccountName: "2",
		Balance:     3000,
		IsActive:    true,
	})

	//Initialize Services
	disbursementSvc := disbursementSvcPkg.New(journalRepo, accountRepo)

	//Initialize server
	srv := transport.NewServer()

	//Initialize APIs
	disbursementHandlerPkg.API(srv.GetEcho(), disbursementSvc)

	//Start the server
	srv.StartServer()
}
