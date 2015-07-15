package fixtures

import (
	"github.com/fdsolutions/logan/args"
)

var (
	defaultFlags = map[string]bool{"--help": false, "--version": false, "--sudo": false}

	UserInputExamples = []struct {
		In       string
		Expected args.Arg
	}{
		{"--sudo start:server",
			args.Arg{"start:server",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				nil}},
		{"start:server --sudo ", // Test overlaping
			args.Arg{"start:server",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				nil}},
		{"show:help",
			args.Arg{"show:help",
				defaultFlags,
				nil}},
		{"install:pkg:ubuntu PKG_NAME='apache'",
			args.Arg{"install:pkg:ubuntu",
				defaultFlags,
				map[string]string{"PKG_NAME": "apache"}}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			args.Arg{"connect:database:mysql",
				defaultFlags,
				map[string]string{"DATABASE_NAME": "mysqldb"}}},
		{"--sudo connect:database:mysql DATABASE_NAME='mysqldb' USER='root' PASSWORD='root' VERSION=1.0.1",
			args.Arg{"connect:database:mysql",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				map[string]string{"DATABASE_NAME": "mysqldb",
					"USER":     "root",
					"PASSWORD": "root",
					"VERSION":  "1.0.1"}}},
	}
)
