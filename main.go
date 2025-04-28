package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"per/analyze"
	"per/copy"
	"strings"
)

func main() {

	os.Mkdir("Collector", 0777) // Creating Directory to Copy Persistence Entries
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting directory")
	}
	collectingdir := filepath.Join(cwd, "Collector") // Creates directory to store files

	dirlist, err := os.Open("persistencelist.txt") // this file contains all the files and directories where malwares can hide, taken from below path
	//https://github.com/elastic/detection-rules/blob/main/rules/integrations/fim/persistence_suspicious_file_modifications.toml
	if err != nil {
		fmt.Println("error accessing file list", err)
	}

	scanner := bufio.NewScanner(dirlist)

	// Iterating throuh Persistence Entries (Files & Folders)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("Collecting -->", line)

		checkdir, err := os.Stat(line)

		if err != nil {
			fmt.Printf("file/folder does not exists or perm issue %s \n", err)
			continue
		}

		if checkdir.IsDir() {

			output := copy.CopyDir(line, collectingdir)
			if output != nil {
				fmt.Println("error coping dir", output)
			}
		} else {

			temp := strings.Split(line, "/")
			fname := strings.Join(temp, "") //Here we are creaeting path inside "Collection" directory.

			finalpath := filepath.Join(collectingdir, fname)
			error := copy.CopyFile(line, finalpath)
			if error != nil {
				fmt.Println("error coping file", error)
			}
		}
	}

	// Here we are collecting files inside different profiles under home directory

	profiles := exec.Command("ls", "/home/")
	stdop, err := profiles.Output()

	if err != nil {
		fmt.Println("err", err)
	}

	var profileiles []string = []string{".profile", ".bashrc", ".bash_profile", ".bash_login", ".bash_logout", ".zprofile", ".zshrc", ".cshrc", ".login", ".logout", ".kshrc",
		".config/autostart-scripts/", ".local/share/autostart/", ".kde4/share/autostart/", ".kde/share/autostart/", ".kde4/Autostart/",
		".kde/Autostart/", ".config/autostart/"}

	for _, l := range strings.Split(string(stdop), "\n") {
		//	fmt.Println("Inside Home", l)
		for _, prof := range profileiles {

			filename := fmt.Sprintf("%s", prof)
			basefile := filepath.Join("/home/", l, filename)

			file_or_dir, err := os.Stat(basefile)

			if err != nil {
				//fmt.Println("not a file or directory")
				continue
			}

			if file_or_dir.IsDir() {
				dir_pro_op := copy.CopyDir(basefile, collectingdir)

				if dir_pro_op != nil {
					fmt.Println("Error processing directory")
				}
			} else {

				oput := copy.CopyFile(basefile, filepath.Join(collectingdir, l+filename))
				if oput != nil {
					fmt.Println("file not present or permission issues", oput)
				}

				if err != nil {
					fmt.Println(oput)
				}
			}

		}

	}

	analyzeop := analyze.Analyze(collectingdir) //Here we are simply searching for strings and ip in persistence entries like "tmp" or "sudo" or IP address
	if analyzeop != nil {
		fmt.Println("error in analyzer")
	}

	ipdomainpath := "/home/sanket/Desktop/GoProject/Persistence/MatchingKeywords.txt"
	// MatchingKeywords file contains lines matching suspicious keywords like "tmp" or "ExecStart" which helps in analysis
	//"tmp" is for temporary path from where malwares execute ,  "ExecStart" shows binary path of services.
	//Refer line 45 in Analyze Function
	analyze.FetchIPDomain(ipdomainpath) // This function will search for IP and URL in MatchingKeywords file

}
