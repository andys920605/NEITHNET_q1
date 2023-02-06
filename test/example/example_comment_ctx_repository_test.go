package example

import (
	"context"
	"fmt"
	"log"
	"q1/database"
	"q1/infras"
	"q1/infras/configs"
	"q1/infras/logger"
	"q1/repository/postgredb"
	"q1/utils"
	"time"

	models_rep "q1/models/repository"

	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/xorcare/pointer"
)

var (
	cfg       *configs.Config
	db        *gorm.DB
	apiLogger *logger.ApiLogger
)

func ExampleCommentRep_Insert() {
	initCommentRep()
	repo := postgredb.NewCommentRep(db)
	want := CreateRandomComment()
	err := repo.Insert(context.Background(), want)
	fmt.Println(err)
	// Output:
	// <nil>
}

func ExampleCommentRep_Find() {
	initCommentRep()
	repo := postgredb.NewCommentRep(db)
	want := CreateRandomComment()
	repo.Insert(context.Background(), want)
	got, _ := repo.Find(context.Background(), want.Uuid)
	got.CreateAt = toTaiwanTime(got.CreateAt)
	got.UpdateAt = toTaiwanTime(got.UpdateAt)
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Printf("ExampleCredentialsCtxRepository_Find() mismatch (-want +got):\n%s", diff)
	}
	fmt.Println(true)
	// Output:
	// true
}

func ExampleCommentRep_Updates() {
	initCommentRep()
	repo := postgredb.NewCommentRep(db)
	want := CreateRandomComment()
	repo.Insert(context.Background(), want)
	want.Comment = utils.RandomString(50)
	err := repo.Updates(context.Background(), want)
	fmt.Println(err)
	// Output:
	// <nil>
}
func ExampleCommentRep_Delete() {
	initCommentRep()
	repo := postgredb.NewCommentRep(db)
	want := CreateRandomComment()
	repo.Insert(context.Background(), want)
	repo.Delete(context.Background(), want.Uuid)
	_, got := repo.Find(context.Background(), want.Uuid)
	fmt.Println(got)
	// Output:
	// sql: no rows in result set
}

func initCommentRep() {
	utils.ConfigPath = "example"
	cfgFile, err := configs.LoadConfig(utils.GetConfigPath())
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, _ = configs.ParseConfig(cfgFile)
	apiLogger = logger.NewApiLogger(cfg)
	options := &infras.Options{
		Ctx:    context.Background(),
		Info:   nil,
		Config: cfg,
		Logger: apiLogger,
	}
	db, _ = database.NewDb(options)
}

func CreateRandomComment() *models_rep.Comment {
	return &models_rep.Comment{
		Uuid:     uuid.NewV4().String(),
		ParentId: uuid.NewV4().String(),
		Comment:  utils.RandomString(50),
		Author:   utils.RandomString(10),
		Favorite: false,
		CreateAt: utils.ToLocalDate(pointer.Time(time.Now())),
		UpdateAt: utils.ToLocalDate(pointer.Time(time.Now())),
	}
}

func toTaiwanTime(t *time.Time) *time.Time {
	var localLocation *time.Location
	localLocation, _ = time.LoadLocation("Asia/Taipei")
	res := t.In(localLocation)
	return utils.ToLocalDate(&res)
}
