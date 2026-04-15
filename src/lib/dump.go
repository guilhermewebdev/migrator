package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DumpSchema(driver string, dsn string, outputFile string) error {
	var cmd *exec.Cmd

	switch driver {
	case "mysql":
		// DSN format: user:password@tcp(host:port)/dbname
		parts := strings.Split(dsn, "@")
		if len(parts) < 2 {
			return fmt.Errorf("invalid MySQL DSN")
		}
		credentials := strings.Split(parts[0], ":")
		user := credentials[0]
		password := ""
		if len(credentials) > 1 {
			password = credentials[1]
		}
		
		connection := strings.Split(parts[1], "/")
		if len(connection) < 2 {
			return fmt.Errorf("invalid MySQL DSN")
		}
		dbname := connection[1]
		if strings.Contains(dbname, "?") {
			dbname = strings.Split(dbname, "?")[0]
		}
		hostport := strings.TrimSuffix(strings.TrimPrefix(connection[0], "tcp("), ")")
		hostparts := strings.Split(hostport, ":")
		host := hostparts[0]
		port := "3306"
		if len(hostparts) > 1 {
			port = hostparts[1]
		}

		args := []string{
			"-h", host,
			"-P", port,
			"-u", user,
			"--no-data",
			"--skip-ssl",
			"--no-tablespaces",
			dbname,
		}
		if password != "" {
			args = append(args, "-p"+password)
		}
		cmd = exec.Command("mysqldump", args...)

	case "postgres":
		// DSN format: postgres://user:pass@host:port/db?sslmode=disable
		cmd = exec.Command("pg_dump", "--schema-only", dsn)

	case "sqlite", "sqlite3":
		// DSN format: file path
		cmd = exec.Command("sqlite3", dsn, ".schema")

	default:
		return fmt.Errorf("driver %s not supported for schema dump", driver)
	}

	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("dump failed: %s", string(exitErr.Stderr))
		}
		return err
	}

	return os.WriteFile(outputFile, output, 0644)
}
