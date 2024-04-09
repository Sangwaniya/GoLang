package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"

	"github.com/google/generative-ai-go/genai"
)

const (
	apiEndpoint        = "https://api.openai.com/v1/chat/completions"
	msg         string = "Heyyy...."
	// []interface{}{map[string]interface{}{"role": "system", "content": "Hi can you tell me what is the factorial of 10?"}}
)

func main() {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("apiKey")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.0-pro")
	model.SetTemperature(0.1) // to reduce the variablity in answer
	response, err := model.GenerateContent(ctx, genai.Text(msg))

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}
	printResponse(response)

	
}
func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("-------------------------------")
    // var data map[string]interface{}
	// err = json.Unmarshal(response, &data)
	// if err != nil {
	//     fmt.Println("Error while decoding JSON response:", err)
	//     return
	// }

	// Extract the content from the JSON response
	// content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

}
