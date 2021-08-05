package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main()  {
	var sdkAddr = "git.medlinker.com/foundations/med-common-sdk/"
	src := "/Users/med/work/git/medlinker/med-common/app/service/chronic-patient-service/api/assistant_clinic.pb.go"
	//src := "/Users/med/work/git/medlinker/med-common/app/service/med-xuser/grpc/doctor.pb.go"
	dest := "./aa.pb.go"
	fsrc, err := os.Open(src)
	if err != nil {
		return
	}
	defer fsrc.Close()
	fdest, err := os.Create(dest)
	if err != nil {
		return
	}
	defer fdest.Close()
	fileScanner := bufio.NewScanner(fsrc)
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		var reg = regexp.MustCompile(`([\S\s]*)med-common(/app/service/|/library/)([\w-_]*)/api([\S\s]*)`)
		//var reg = regexp.MustCompile(`([\S\s]*)med-common/app/service/([\w-_]*)/api([\S]*)`)
		if reg.MatchString(text) {
			for _, i := range reg.FindAllStringSubmatch(text, -1) {
				if len(i) > 4 {
					fmt.Printf("FindAllStringSubmatch (%+v)\n", i)
					text = i[1] + sdkAddr +i[3]+ i[4]
				}
			}
		}
		fdest.WriteString(text+"\n")
	}
	return
}
