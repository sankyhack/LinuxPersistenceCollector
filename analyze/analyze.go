package analyze

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// In Analyze function we have specify RegEx that search for keywords , IP addresses
// Keywords like "tmp" "ExecStart" etc, you can specify your own keywords line 45
func Analyze(path string) error {
	err := filepath.Walk(path, func(serv string, info fs.FileInfo, err error) error {

		if err != nil {
			fmt.Println("error recursing dir", err)
		}

		if path == serv {
			return nil

		}
		file_obj, err := os.Open(serv)

		if err != nil {
			fmt.Println("error opening file")
		}

		matchwords, err := os.OpenFile("MatchingKeywords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {
			fmt.Println("error opening file")
		}

		defer matchwords.Close()

		scanner := bufio.NewScanner(file_obj)

		for scanner.Scan() {
			line := scanner.Text()
			//keyword, err := regexp.Compile(`(tmp|sudo|http|www)`)
			keyword, err := regexp.Compile(`(http|www|tmp|ExecStart)`) // here you can add your own keywords and if you dont want just add any random word
			if err != nil {
				fmt.Println("Error Processing Regex")
			}
			//ip, err := regexp.Compile(`([0-9]{1,3}\.){3}[0-9]{1,3}`)
			ip, err := regexp.Compile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`) // this is to match IP address

			if err != nil {
				fmt.Println("error in Regex")
			}
			if keyword.MatchString(line) || ip.MatchString(line) {
				//keywordmatch := keyword.FindAllString(line, -1)
				//ipmatch := ip.FindAllString(line, -1)
				if keyword.MatchString(line) {
					matchwords.WriteString(serv + "\n")
					matchwords.WriteString(line + "\n")

					fmt.Printf("Matching keyword found in file -> %s -> matching line is -> %s \n", serv, line)
				} else {

					matchwords.WriteString(serv + "\n")
					matchwords.WriteString(line + "\n")

					fmt.Printf("IP found in file -> %s -> IP is %s \n", serv, line)
				}

			}

		}

		return err
	})
	return err

}

func FetchIPDomain(ipdomainpath string) {

	peroutput, err := os.Open(ipdomainpath)
	if err != nil {
		fmt.Println("error opening matchingline path")
	}

	matchwords, err := os.OpenFile("IP_Domain_Extract.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Println("error opening file")
	}

	defer matchwords.Close()

	//scanner := bufio.NewScanner(file_obj)
	scanner := bufio.NewScanner(peroutput)

	for scanner.Scan() {
		line := scanner.Text()
		//keyword, err := regexp.Compile(`(tmp|sudo|http|www)`)
		keyword, err := regexp.Compile(`(https?|ftp|file)://\S+`) // here you can add your own keywords and if you dont want just add any random word
		if err != nil {
			fmt.Println("Error Processing Regex")
		}
		//ip, err := regexp.Compile(`([0-9]{1,3}\.){3}[0-9]{1,3}`)
		ip, err := regexp.Compile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`) // this is to match IP address

		if err != nil {
			fmt.Println("error in Regex")
		}
		if keyword.MatchString(line) || ip.MatchString(line) {
			//keywordmatch := keyword.FindAllString(line, -1)
			//ipmatch := ip.FindAllString(line, -1)
			if keyword.MatchString(line) {
				ipdomainmatch_url := keyword.FindAllString(line, -1)
				matchwords.WriteString(strings.Join(ipdomainmatch_url, "") + "\n")
				//	matchwords.WriteString(line + "\n")

				//	fmt.Printf("Matching keyword found in file -> %s -> matching line is -> %s \n", serv, line)
			} else {
				ipdomainmatch_ip := ip.FindAllString(line, -1)
				matchwords.WriteString(strings.Join(ipdomainmatch_ip, "") + "\n")
				//	matchwords.WriteString(line + "\n")

				//fmt.Printf("IP found in file -> %s -> IP is %s \n", serv, line)
			}

		}

	}

}
