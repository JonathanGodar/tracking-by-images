package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/JonathanGodar/go-web-gin/models"
	"github.com/JonathanGodar/go-web-gin/server/myoto"
	_ "github.com/lib/pq"
	"github.com/pacedotdev/oto/otohttp"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Could not initialize logger")
	}
	defer logger.Sync()

	otoServer := otohttp.NewServer()
	db, err := createDatabase()
	if err != nil {
		logger.Fatal("Could not initilize database", zap.Error(err))
	}

	RegisterServices(otoServer, db, logger)
	serverAddr := ":8080"

	logger.Debug("Binding TrackerHandler")
	http.Handle("/tracker/", CORSMiddleware(trackerHandler(db, logger)))

	logger.Debug("Binding Oto endpoint")
	http.Handle("/oto/", CORSMiddleware(authenticationMiddleware(otoServer)))

	logger.Info("Starting server listen", zap.String("address", serverAddr))
	logger.Fatal("Server exited", zap.Error(http.ListenAndServe(":8080", nil)))
}

func trackerHandler(db *sql.DB, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		results := r.URL.Query()["id"]

		if len(results) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}

		id := results[0]
		log.Println(id)
		tracker, err := models.Trackers(qm.Where("id = ?", id)).One(r.Context(), db)
		if err != nil {
			logger.Info("Could not find tracker", zap.String("trackerId", id), zap.Error(err))
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}

		if tracker.IsActive {
			tracker.TimesAccessed++
			_, err = tracker.Update(r.Context(), db, boil.Whitelist(models.TrackerColumns.TimesAccessed))
			if err != nil {
				logger.DPanic("Could not update timesAccessed on tracker",
					zap.String("trackerId", id), zap.Error(err),
				)
			}
		}

		imgbytes, err := ioutil.ReadFile("./server/static/transparent-pixel.png")
		DieIf(err)
		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Type", "image/png")
		w.Header().Add("Content-Length", fmt.Sprint(len(imgbytes)))
		w.Write(imgbytes)
	})
}

func createDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", `dbname=go-testing host=localhost user=postgres password=postgres sslmode=disable`)

	return db, err
}

func RegisterServices(server *otohttp.Server, db *sql.DB, logger *zap.Logger) {
	userServiceInstance := new(userService)
	trackerServiceInstance := new(trackerService)

	trackerServiceInstance.Initialize(db, logger)
	userServiceInstance.Initialize(db, logger)

	myoto.RegisterTrackerService(server, trackerServiceInstance)
	myoto.RegisterUserService(server, userServiceInstance)
}

func DieIf(err error) {
	if err != nil {
		log.Fatalf("%f", err)
	}
}
