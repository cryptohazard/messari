package main

import "fmt"
import "messari"
import "log"

func main() {
	fmt.Println("\t***gomessari starting***")

	client := messari.Client{}
	/*markets, err := client.Markets()
	if err != nil {
		log.Fatal("unexpected error\n", err)
	}
	fmt.Println(markets.String())*/

	/*assets, err := client.Assets()
	if err != nil {
		log.Fatal("unexpected error\n", err)
	}
	fmt.Println(assets.String())*/

	news, err := client.News()
	if err != nil {
		log.Fatal("unexpected error\n", err)
	}
	fmt.Println(news.String())

}
