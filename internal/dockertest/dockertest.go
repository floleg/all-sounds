package dockertest

import (
	"allsounds/pkg/config"
	"allsounds/pkg/db"
	"log"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/gorm"
)

const (
	dbName   = "test"
	dbPasswd = "test"
)

func SetupDockerDb() (*gorm.DB, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	runDockerOpt := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14",
		Env:        []string{"POSTGRES_PASSWORD=" + dbPasswd, "POSTGRES_DB=" + dbName}}

	// set AutoRemove to true so that stopped container goes away by itself
	// don't restart container
	fnConfig := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.NeverRestart()
	}

	resource, err := pool.RunWithOptions(runDockerOpt, fnConfig)
	if err != nil {
		panic(err)
	}

	// declare clean up function to release resource
	fnCleanup := func() {
		err := resource.Close()
		if err != nil {
			panic(err)
		}
	}

	config := config.Config{DBHost: "localhost", DBPort: resource.GetPort("5432/tcp"), DBName: dbName, DBUSer: "postgres", DBPassword: dbPasswd}

	// retry until db server is ready
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		db.Init(&config)
		gdb, err := db.GetDB().DB()
		if err != nil {
			return err
		}
		return gdb.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err != nil {
		panic(err)
	}

	// container is ready, return *gorm.Db for testing
	return db.GetDB(), fnCleanup
}
