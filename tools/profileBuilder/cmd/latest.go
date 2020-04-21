package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/samjegal/fincloud-sdk-for-go/tools/internal/dirs"
	"github.com/samjegal/fincloud-sdk-for-go/tools/profileBuilder/model"
	"github.com/spf13/cobra"
)

const (
	previewLongName    = "preview"
	previewShortName   = "p"
	previewDescription = "Include preview API versions."
)

const (
	rootLongName    = "root"
	rootShortName   = "r"
	rootDescription = "The location of the API Version folders which should be considered for `latest`."
)

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Reflects on the available packages, choosing the most recent ones.",
	Long: `Scans through the availabe API Versions, and chooses only the most 
recent functionality.

By default, this command ignores API versions that are in preview.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logWriter := ioutil.Discard
		if verboseFlag {
			logWriter = os.Stdout
		}

		outputLog := log.New(logWriter, "[STATUS] ", 0)
		errLog := log.New(os.Stderr, "[ERROR] ", 0)

		if !filepath.IsAbs(outputRootDir) {
			abs, err := filepath.Abs(outputRootDir)
			if err != nil {
				errLog.Fatalf("failed to convert to absolute path: %v", err)
			}
			outputRootDir = abs
		}
		fmt.Printf("Executes profileBuilder in %s\n", outputRootDir)
		outputLog.Printf("Output-Location set to: %s", outputRootDir)

		includePreview, err := cmd.Flags().GetBool(previewLongName)
		if err != nil {
			errLog.Fatalf("failed to get preview flag: %v", err)
		}
		if includePreview {
			outputLog.Println("Using preview versions.")
		}

		if clearOutputFlag {
			if err := dirs.DeleteChildDirs(outputRootDir); err != nil {
				errLog.Fatalf("Unsable to clear output-folder: %v", err)
			}
		}
		rootDir, err := cmd.Flags().GetString(rootLongName)
		if err != nil {
			errLog.Fatalf("failed to get root dir: %v", err)
		}
		listDef, err := model.GetLatestPackages(rootDir, includePreview, outputLog)
		model.BuildProfile(listDef, profileName, outputRootDir, outputLog, errLog, false, modulesFlag)
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
	latestCmd.Flags().BoolP(previewLongName, previewShortName, false, previewDescription)
	latestCmd.Flags().StringP(rootLongName, rootShortName, "", rootDescription)
	latestCmd.MarkFlagRequired(rootLongName)
}
