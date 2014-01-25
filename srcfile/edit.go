package srcfile

import (
	"bufio"
	"os"
	"strings"
)

//original sentence and translated sentence
type coupleOfSent struct {
	Org   string
	Trans string
}

//connect two list to one map
func MakeMap(orgList []string, transList []string) (map[int]coupleOfSent, int) {
	tempMap := make(map[int]coupleOfSent)
	if len(orgList) < len(transList) {
		return tempMap, 4
		//errcode 4 is LengthError
	} else if len(orgList) == len(transList) {
		for i := 0; i < len(orgList); i++ {
			tempMap[i] = coupleOfSent{orgList[i], transList[i]}
		}
	} else {
		for i := 0; i < len(transList); i++ {
			tempMap[i] = coupleOfSent{orgList[i], transList[i]}
		}
		for j := len(transList); j < len(orgList); j++ {
			tempMap[j] = coupleOfSent{orgList[j], ""}
		}
	}
	return tempMap, 0
}

func ReadTxt(fileAddr string) []string {
	//Open file
	file, _ := os.Open(fileAddr)
	defer file.Close()
	//Read file
	tempList := []string{}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		tempList = append(tempList, reader.Text())
	}
	return tempList
}

func ReadHtml(fileAddr string) []string {
	return []string{}
}

func TextList(orgFileAddr string) ([]string, int) {
	//Check for file existence
	if _, err := os.Stat(orgFileAddr); err != nil {
		return []string{}, 1
		//errcode 1 is noFileError
	}
	readTextList := []string{}
	//Catch the extension of file
	ext := strings.Split(orgFileAddr, ".")
	str := ext[len(ext)-1]
	//Distinguish by extension
	switch str {
	case "txt":
		readTextList = ReadTxt(orgFileAddr)
	case "html":
		readTextList = ReadHtml(orgFileAddr)
	default:
		return []string{}, 2
		//errcode 2 is no-extension-added error
	}
	//return list
	return readTextList, 0
}

func TransList(transFileAddr string) ([]string, int) {
	fileInfo, err := os.Stat(transFileAddr)
	if err != nil {
		file, _ := os.Create(transFileAddr)
		file.Close()
		return []string{}, 1
	} else if fileInfo.IsDir() {
		return []string{}, 3
		//errcode 3 is ItsNotFileError
	}
	//fileread
	return ReadTxt(transFileAddr), 0
}
