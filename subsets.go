package subsets

import (
    "bufio"
    "fmt"
    "log"
    "unicode"
    "strings"
    "bytes"
    "github.com/JeremyOne/runeTree"
)

//check if every item in array 'B' is contained in array 'A' using a brute-force method
func ArrayIsSubset_BruteForce(A []string, B []string, returnEarly bool) bool{

    foundAll := true
    for _, itemB := range B {
        foundB := false

        for _, itemA := range A {

            if itemA == itemB{
                foundB = true
                break
            }
        }

        if(returnEarly && foundB == false){
            //if we did not find, and we want to return ASAP
            return false
        } else if(foundB == false) {
            //if we are going to return late
            foundAll = false
        }
    }

    return foundAll
}

//check if every item in array 'B' is contained in array 'A' using a hash map
func ArrayIsSubset_HashMap(A []string, B []string, returnEarly bool, reportSize bool) bool{

    mapA := make(map[string]bool)
    foundAll := true

    //insert all words into a hash map
    for _, itemA := range A {
        //use the word as the key, this will eliminate duplicates
        mapA[itemA] = true
    }

    if(reportSize){
        fmt.Printf("- Hashmap is %d elements long\n", len(mapA))
    }

    for _, itemB := range B {
        _, foundItem := mapA[itemB]
        
        if(foundItem == false && returnEarly){
            //if we did not find, and we want to return ASAP
            return false
        } else if foundItem == false {
            //if we are going to return late
            foundAll = false
        }
    }

    return foundAll;
}

//check if every item in array 'B' is contained in array 'A' using a RuneTree (search tree)
func ArrayIsSubset_RuneTree(A []string, B []string, returnEarly bool) bool{

    tree := &runeTree.RuneTree{'-', &[]runeTree.RuneTree{}}

    foundAll := true

    for _, itemA := range A {
        //add all items to the search tree
        runeTree.RuneTreeInsert(tree, itemA, false)
    }

    var foundItem bool
    for _, itemB := range B {
        foundItem = runeTree.RuneTreeInsert(tree, itemB, true)

        if(foundItem == false && returnEarly){  
            //if we did not find, and we want to return ASAP
            return false
        } else if foundItem == false {
            //if we are going to return late
            foundAll = false
        }
    }

    return foundAll
}

//Reads from a scanner and returns a list of lower-case words with no punctuation
func ReadScannerWords(scanner *bufio.Scanner) []string{
    output := []string{}

    //create a word buffer
    currentWordBuffer := bytes.NewBuffer(nil)
    currentWordWriter := bufio.NewWriter(currentWordBuffer)
    currentWordLength := 0;

    //scan line by line
    //this implementation ignores punctuation, numbers and repeated whitespace
    for scanner.Scan() {
        
        currentLine := strings.ToLower(scanner.Text())
        
        for _, c := range currentLine {
            if unicode.IsLower(c) {
                currentWordWriter.WriteRune(c)
                currentWordLength++;
            } else if unicode.IsSpace(c) {

                //end of the word, add to output and reset
                if(currentWordLength > 0){
                    currentWordWriter.Flush()
                    output = append(output, string(currentWordBuffer.Bytes()))
                    currentWordLength = 0;
                    currentWordBuffer.Reset()
                }
            }
        }

        //write the last word
        if(currentWordLength > 0){
            currentWordWriter.Flush()
            output = append(output, string(currentWordBuffer.Bytes()))
            currentWordLength = 0;
            currentWordBuffer.Reset()
        }

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return output
}