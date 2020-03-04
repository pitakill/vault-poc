package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

type wrapper struct {
	client     *api.Client
	username   string
	password   string
	authorizer string
	token      string
}

const (
	url      = "https://vault.pitakill.net:8200"
	username = "authorizer"
	password = "helloworld"

	approleLogin       = "auth/approle/login"
	lmsPrograms        = "v1/data/lms/capture/programs"
	lmsProgramsPublish = "v1/data/lms/capture/programs/publish"
)

func main() {
	w, err := newWrapper(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Login as authorizer to get the permissions to query Vault as a defined role
	if err := w.loginWithUserPassword(); err != nil {
		log.Fatal(err)
	}

	// Roles to test
	roles := []string{
		"qa",       // quality assurance
		"insdes",   // instructional designer
		"insdessr", // instructional designer sr
		"admin",    // adminstrator
	}

	for _, role := range roles {
		// Login to Vault as role
		if err := w.loginAs(role); err != nil {
			log.Fatal(err)
		}

		cantDelete := fmt.Sprintf("The role %q DOES NOT HAVE permissions to DELETE programs on the LMS\n", role)
		canDelete := fmt.Sprintf("The role %q HAS permissions to DELETE programs on the LMS\n", role)
		cantWrite := fmt.Sprintf("The role %q DOES NOT HAVE permissions to MODIFY programs on the LMS\n", role)
		canWrite := fmt.Sprintf("The role %q HAS permissions to MODIFY programs on the LMS\n", role)
		cantRead := fmt.Sprintf("The role %q DOES NOT HAVE permissions to READ programs on the LMS\n", role)
		canRead := fmt.Sprintf("The role %q HAS permissions to READ programs on the LMS\n", role)
		cantPublish := fmt.Sprintf("The role %q DOES NOT HAVE permissions to PUBLISH programs on the LMS\n", role)
		canPublish := fmt.Sprintf("The role %q HAS permissions to PUBLISH programs on the LMS\n", role)

		fmt.Printf("\n")

		// Verify if can write
		if ok := w.canWrite(lmsPrograms); !ok {
			fmt.Printf(cantWrite)
		} else {
			fmt.Printf(canWrite)
		}

		// Verify if can read
		if ok := w.canRead(lmsPrograms); !ok {
			fmt.Printf(cantRead)
		} else {
			fmt.Printf(canRead)
		}

		// Verify if can delete
		if ok := w.canDelete(lmsPrograms); !ok {
			fmt.Printf(cantDelete)
		} else {
			fmt.Printf(canDelete)
		}

		// Verify if can publish
		if ok := w.canRead(lmsProgramsPublish); !ok {
			fmt.Printf(cantPublish)
		} else {
			fmt.Printf(canPublish)
		}
	}
}
