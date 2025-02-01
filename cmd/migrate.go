package cmd

import (
	"os"

	"api-service/model"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DoMigrateUpCMD() cli.Command {

	return cli.Command{
		Name:  "migrate-up",
		Usage: "Run migration up with specific database source address",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "source",
				Value: "",
			},
		},
		Action: func(c *cli.Context) error {
			sourceDBArg := c.String("source")

			db, err := gorm.Open(postgres.Open(sourceDBArg), &gorm.Config{})

			if err != nil {
				logrus.Fatalf("[migrate-up] Failed to connect database source %s \n", err.Error())
				os.Exit(1)
			}

			err = db.AutoMigrate(
				&model.Course{},
				&model.User{},
				&model.SubMaterial{},
				&model.Material{},
			)

			if err != nil {
				logrus.Fatalf("[migrate-up] Failed to run migration because %s \n", err.Error())
				os.Exit(1)
			}

			logrus.Info("[migrate-up] Successfuly migrate up database...")

			return nil
		},
	}
}
