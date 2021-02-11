package main

import (
    "context"
    "net/http"
    "encoding/json"
    "strings"
    "os"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
)

type Notification struct {
    Type string `json:"type"`
    Title string `json:"title"`
    Body string `json:"body"`
}

func handler(ctx context.Context, snsEvent events.SNSEvent) {
    for _, record := range snsEvent.Records {
        snsRecord := record.SNS
        
        notification := &Notification{
            Title: snsRecord.Subject,
            Body:  snsRecord.Message,
            Type: "note",
        }
        
        push(*notification)
    }
}

func main() {
    lambda.Start(handler)
}

func push(notification Notification) int {
    access_token := os.Getenv("PUSHBULLET_ACCESS_TOKEN")

    body := to_json(notification)
    req, _ := http.NewRequest("POST", "https://api.pushbullet.com/v2/pushes", strings.NewReader(body))
    req.Header.Add("Access-Token", access_token)
    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{}
    resp, _ := client.Do(req)
    defer resp.Body.Close()

    return resp.StatusCode
}

func to_json(data interface{}) string {
    jsonStr, _ := json.Marshal(data)
    return string(jsonStr)
}