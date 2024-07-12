package app_rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pr1/internal/rest/bootstrap"
	"pr1/internal/rest/config"
	"pr1/internal/rest/repository/database"
	"syscall"

	"pr1/internal/rest/app/service"
)

func Run(cfg *config.Config) error {

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	bicycleService := service.NewService(database.NewDatabase(db))

	router := http.NewServeMux()
	router.HandleFunc("POST /bicycles", bicycleService.Create)
	router.HandleFunc("GET /bicycles/{id}", bicycleService.Get)
	router.HandleFunc("GET /bicycles", bicycleService.GetAll)
	router.HandleFunc("PUT /bicycles/{id}", bicycleService.Update)
	router.HandleFunc("DELETE /bicycles/{id}", bicycleService.Delete)

	srv := http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("run server: http://localhost%s", cfg.Port)
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("error when listen and serve: ", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	sig := <-ch
	log.Printf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done\n")
	return srv.Shutdown(context.Background())
}
