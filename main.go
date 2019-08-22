package main

import (
	"fmt"
	"os"

	"github.com/mikloslorinczi/go-test/client"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Go-test")
	res, err := client.Call("https://reqres.in/api/users")

	if err != nil {
		fmt.Println("Testing Request Error:", err)
		os.Exit(1)
	}

	fmt.Println("Testing Request Results:\n", res)

	if err := spammGitHub(20); err != nil {
		log.WithError(err).Fatal("Cannot query GitHub")
	}
}

func spammGitHub(rounds int) error {
	for i := 0; i < rounds; i++ {
		res, err := client.Call("https://api.github.com/repos/bitrise-io/bitrise.io/contents/system_reports")
		if err != nil {
			return errors.Wrap(err, "Failed to request GitHub")
		}
		log.WithField("Round", i+1).WithField("Response", res).Info("Succesfull request")
	}
	return nil
}
