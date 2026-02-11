package main

import (
	"database/sql"
	"github.com/Kazimlyc/blog-aggregator/internal/config"
	"github.com/Kazimlyc/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerFeed)
	cmds.register("feeds", handlerFeeds)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{Name: cmdName, Args: cmdArgs}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	if err := cmds.run(programState, cmd); err != nil {
		log.Fatal(err)
	}

}
