package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/blang/semver"

	log "github.com/sirupsen/logrus"
)

func main() {
	if cap(os.Args) < 2 {
		log.Fatalf("invalid argument specified. pass major, minor, patch or init")
		os.Exit(1)
	}

	callArgument := os.Args[1]
	if callArgument == "init" {
		err := initializeVersion()
		if err != nil {
			log.Fatalf("Error initializing version %v", err)
			os.Exit(1)
		}
		os.Exit(0)
		return
	}

	version, err := getLatestGitTag()
	if err != nil {
		log.Fatalf("Error getting latest git tag %v", err)
		os.Exit(1)
	}
	bumpedVersion, err := bumpVersion(version, callArgument)
	if err != nil {
		log.Fatalf("Error bumping version %v", err)
		os.Exit(1)
	}
	err = pushGitTag(bumpedVersion)
	if err != nil {
		log.Fatalf("Error pushing git tag %v", err)
		os.Exit(1)
	}
	log.Infof("Successfully bumped to version %v", bumpedVersion)
	os.Exit(0)
}

func initializeVersion() error {
	initVersion, err := semver.Make("0.0.1")
	if err != nil {
		return err
	}
	err = pushGitTag(&initVersion)
	if err != nil {
		return err
	}
	return nil
}

func getLatestGitTag() (*semver.Version, error) {
	output, err := runCommand(`git tag --sort=committerdate | grep -o 'v[0-9]*\.[0-9]*\.[0-9]*' | tail -n1`)
	if err != nil {
		return nil, err
	}
	versionString := strings.TrimSpace(strings.TrimPrefix(output, "v"))
	version, err := semver.Make(versionString)
	if err != nil {
		return nil, err
	}
	return &version, nil
}

func bumpVersion(version *semver.Version, whichToBump string) (*semver.Version, error) {
	if whichToBump == "major" {
		version.Major += 1
		version.Minor = 0
		version.Patch = 0
	} else if whichToBump == "minor" {
		version.Minor += 1
		version.Patch = 0
	} else if whichToBump == "patch" {
		version.Patch += 1
	} else {
		return nil, fmt.Errorf("invalid argument %v", whichToBump)
	}
	return version, nil
}

func runCommand(cmd string) (string, error) {
	outputBytes, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	outputString := fmt.Sprintf("%s", outputBytes)
	return outputString, nil
}

func pushGitTag(version *semver.Version) error {
	tag := fmt.Sprintf(`v%v`, version.String())
	createTagCommand := fmt.Sprintf(`git tag -a %v -m 'pushing %v'`, tag, tag)
	pushTagCommand := fmt.Sprintf(`git push origin %v`, tag)

	_, err := runCommand(createTagCommand)
	if err != nil {
		return fmt.Errorf("command `%v` failed %v", createTagCommand, err)
	}
	log.Infof("ran command `%v`", createTagCommand)

	_, err = runCommand(pushTagCommand)
	if err != nil {
		return fmt.Errorf("command `%v` failed %v", pushTagCommand, err)
	}
	log.Infof("ran command `%v`", pushTagCommand)
	return nil
}
