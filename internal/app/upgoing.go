package app

import (
	"fmt"
	"time"

	"github.com/doddy-s/miren/internal/utils"
)

func Upgoing() error {
	username, err := utils.Whoami()
	if err != nil {
		return err
	}

	targetFiles := []utils.MirenDirEntry{}

	gitFiles, err := utils.ListFiles("/home/"+username+"/", "*.git*")
	if err != nil {
		return err
	}

	sshFiles, err := utils.ListFiles("/home/"+username+"/.ssh", "")
	if err != nil {
		return err
	}

	targetFiles = append(targetFiles, gitFiles...)
	targetFiles = append(targetFiles, sshFiles...)

	gitFilesSize, err := utils.CountTotalSize(gitFiles)
	if err != nil {
		return err
	}

	fmt.Println("Git size: ", utils.FormatBytes(gitFilesSize))

	sshFilesSize, err := utils.CountTotalSize(sshFiles)
	if err != nil {
		return err
	}

	fmt.Println("SSH size: ", utils.FormatBytes(sshFilesSize))

	isContinuing := utils.AskConfirmation("Proceed with backup?")
	if !isContinuing {
		return nil
	}

	zipFullPath := "/home/" + username + "/miren-backup-" + time.Now().UTC().Format(time.RFC3339) + ".zip"

	err = utils.ZipFilesPreserveStructure(
		zipFullPath,
		targetFiles,
		"/",
		utils.PrintProgressBar,
	)
	if err != nil {
		return err
	}

	fmt.Println("Zip Path: ", zipFullPath)

	return nil
}
