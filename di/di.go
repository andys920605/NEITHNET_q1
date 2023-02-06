package di

import (
	"context"
	"q1/database"
	"q1/infras"
	"q1/infras/logger"
	"q1/models/commons"
	rep "q1/repository/postgredb"
	"q1/router"
	svc "q1/service"

	"q1/app"
)

// create multipay server
// @result server and error
func CreateQ1Server(ctx context.Context, info *commons.SystemInfo) (*app.Q1Server, error) {
	config := infras.ProvideConfig()
	apiLogger := logger.NewApiLogger(config)
	options := &infras.Options{
		Ctx:    ctx,
		Info:   info,
		Config: config,
		Logger: apiLogger,
	}
	db, _ := database.NewDb(options)
	// Repo
	commentRep := rep.NewCommentRep(db)
	// Svc
	commentSvc := svc.NewCommentSvc(commentRep)
	// Router
	router := router.NewRouter(commentSvc)
	q1Server := app.NewQ1Server(options, router)
	return q1Server, nil
}
