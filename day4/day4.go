package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*  passport expected fields:
byr - birth year
iyr - issue year
eyr - expiration year
hgt - height
hcl - hair color
ecl - eye color
pid - passport id
cid - country id (not required)
*/
type passport struct {
	byr, iyr, eyr, hgt string
	hcl, ecl, pid, cid string
}

var requiredFields = []string{
	"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
}

func main() {
	abs, err := filepath.Abs("./passports.txt")
	file, err := os.Open(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if len(data) < 1 {
		log.Fatal("data not parsed")
	}

	var passports []passport
	var validPassports int
	fieldsMap := make(map[string]string)

	for i, s := range data {
		//fmt.Println(s)
		fields := strings.Split(s, " ")

		for _, f := range fields {
			split := strings.Split(f, ":")
			if len(split) > 1 {
				fieldsMap[split[0]] = split[1]
				//fmt.Printf("var: %v, %v\n", split[0], split[1])
			}
		}

		if s == "" || i == len(data)-1 {
			if !validatePassport(fieldsMap) {
				//fmt.Printf("invalid passport! %v\n", fieldsMap)
				fieldsMap = make(map[string]string)
				continue
			}
			//fmt.Printf("new passport!\n")
			p := passport{
				byr: fieldsMap["byr"],
				iyr: fieldsMap["iyr"],
				eyr: fieldsMap["eyr"],
				hgt: fieldsMap["hgt"],
				hcl: fieldsMap["hcl"],
				ecl: fieldsMap["ecl"],
				pid: fieldsMap["pid"],
				cid: fieldsMap["cid"],
			}
			passports = append(passports, p)
			fieldsMap = make(map[string]string)
			validPassports++
			continue
		}
	}

	for _, p := range passports {
		fmt.Printf("%v\n", p)
	}

	fmt.Printf("valid passports: %v", validPassports)
}

func validatePassport(fieldsMap map[string]string) (validPassport bool) {
	validPassport = true
	for _, s := range requiredFields {
		if fieldsMap[s] == "" {
			validPassport = false
			//fmt.Printf("failed check! %v\n", s)
			break
		}
	}
	return
}
