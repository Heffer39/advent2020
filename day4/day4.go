package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
		fields := strings.Split(s, " ")
		for _, f := range fields {
			split := strings.Split(f, ":")
			if len(split) > 1 {
				fieldsMap[split[0]] = split[1]
			}
		}

		if s == "" || i == len(data)-1 {
			if !validatePassport(fieldsMap) {
				fieldsMap = make(map[string]string)
				continue
			}
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

	fmt.Printf("valid passports: %v", validPassports)
}

func validatePassport(fieldsMap map[string]string) (validPassport bool) {
	validPassport = true
	for _, s := range requiredFields {
		val := fieldsMap[s]
		if val == "" {
			return false
		}
		switch s {
		case "byr":
			birthYear, _ := strconv.Atoi(val)
			if len(val) != 4 || birthYear < 1920 || birthYear > 2002 {
				fmt.Printf("invalid byr: %v\n", val)
				return false
			}
		case "iyr":
			issueYear, _ := strconv.Atoi(val)
			if len(val) != 4 || issueYear < 2010 || issueYear > 2020 {
				fmt.Printf("invalid iyr: %v\n", val)
				return false
			}
		case "eyr":
			expirationYear, _ := strconv.Atoi(val)
			if len(val) != 4 || expirationYear < 2020 || expirationYear > 2030 {
				fmt.Printf("invalid eyr: %v\n", val)
				return false
			}
		case "hgt":
			validPassport = false
			if strings.Contains(val, "in") {
				num, _ := strconv.Atoi(strings.Split(val, "in")[0])
				if num >= 59 && num <= 76 {
					validPassport = true
				}
			}
			if strings.Contains(val, "cm") {
				num, _ := strconv.Atoi(strings.Split(val, "cm")[0])
				if num >= 150 && num <= 193 {
					validPassport = true
				}
			}
			if !validPassport {
				fmt.Printf("invalid hgt: %v\n", val)
				return false
			}
		case "hcl":
			if len(val) != 7 || val[0] != '#' || !regexp.MustCompile(
				`^[a-f0-9]+$`).MatchString(strings.SplitN(val, "#", 2)[1]) {
				fmt.Printf("invalid hcl: %v\n", val)
				return false
			}
		case "ecl":
			validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			validPassport = false
			for _, c := range validColors {
				if c == val {
					validPassport = true
					break
				}
			}
			if validPassport == false {
				fmt.Printf("invalid ecl: %v\n", val)
				return false
			}
		case "pid":
			if len(val) != 9 || !regexp.MustCompile(`^[0-9]+$`).MatchString(val) {
				fmt.Printf("invalid pid: %v\n", val)
				return false
			}
		}
	}
	return
}
