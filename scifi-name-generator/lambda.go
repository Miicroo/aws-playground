package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"
    "strings"
    "github.com/aws/aws-lambda-go/lambda"
)

type SciFiCity struct {
    Name string `json:"name"`
}

func HandleRequest() (string, error) {
    rand.Seed(time.Now().UTC().UnixNano())
    city := &SciFiCity{GenerateSciFiName(4, 11)}
    return fmt.Sprintf(toJson(city)), nil
}

func main() {
    lambda.Start(HandleRequest)
}

func toJson(data interface{}) string {
    jsonStr, _ := json.Marshal(data)
    return string(jsonStr)
}

func GenerateSciFiName(min int, max int) string {
    alphabet := GetFrequencyAlphabet()

    var name string = ""
    nameLength := Random(min, max)

    for i := 0; i<nameLength; i++ {
    	character := alphabet[Random(0, len(alphabet) - 1)]
    	name += string(character)
    }

    return strings.Title(strings.ToLower(name))
}

func GetFrequencyAlphabet() string {
    letterFrequencies := map[string]int{
	"E": 5688,
	"A": 4331,
	"R": 3864,
	"I": 3845,
	"O": 3651,
	"T": 3543,
	"N": 3392,
	"S": 2923,
	"L": 2798,
	"C": 2313,
	"U": 1851,
	"D": 1725,
	"P": 1614,
	"M": 1536,
	"H": 1531,
	"G": 1259,
	"B": 1056,
	"F": 924,
	"Y": 906,
	"W": 657,
	"K": 561,
	"V": 513,
	"X": 148,
	"Z": 139,
	"J": 100,
	"Q": 100,
    }

    var alphabet string = ""

    for letter, frequency := range letterFrequencies {
	alphabet += strings.Repeat(letter, frequency)
    }

    return alphabet
}

func Random(min int, max int) int {
    return rand.Intn(max + 1 - min) + min
}
