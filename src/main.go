package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func findKeyWord(word string, doc []string) (wordCound int){
    for _, v := range doc { 
        if v == word {
            fmt.Println(v)
            wordCound += 1
        }
    } 
    return wordCound
}

func parsString(str io.Reader) (parsedString []string) {
    scanner := bufio.NewScanner(str)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        parsedString = append(parsedString, scanner.Text())
    }
    
    return parsedString
}

func main() {    
    resume, err := os.Open("../resume/resume")
    if err != nil {
        log.Fatal(err)
    }
    defer resume.Close()

    jobDescription, err := os.Open("../jobDescription/jobDescription ") 
    if err != nil {
        log.Fatal(err)
    }
    defer jobDescription.Close()

    parsedResume := parsString(resume)
    parsedJobDes := parsString(jobDescription)

    var wordCount int

    for _, v := range parsedJobDes {
        wordCount += findKeyWord(v, parsedResume) 
    }

    mached := fmt.Sprintf("%d",(wordCount/len(parsedJobDes))*100)

    fmt.Println(mached+"%")
}
