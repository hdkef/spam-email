package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func spamTarget(target *string, total *int, delay *int, text *string) {

	if *target == "" {
		panic("NO TARGET, PLS BUILD NEW SCRIPT")
	}

	fmt.Println("OK. pls wait")

	for {
		go func() {
			fmt.Println("endpoint hit ", *target, *total, *delay, *text)

			counter := 0

			for {

				if counter >= *total {
					break
				}

				code := rand.Int31()

				emailMe := os.Getenv("EMAIL")
				pswdMe := os.Getenv("PSWD")
				host := os.Getenv("HOST")
				emailTo := []string{*target}
				port := os.Getenv("SMTPPORT")
				addr := fmt.Sprintf("%s:%s", host, port)
				auth := smtp.PlainAuth("", emailMe, pswdMe, host)
				msgString := "Subject: Hello brother\n\n" + *text + fmt.Sprint(code)
				msgBytes := []byte(msgString)

				err := smtp.SendMail(addr, auth, emailMe, emailTo, msgBytes)

				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("counter : ", counter)

				time.Sleep(time.Duration(*delay) * time.Second)

				counter++
			}

			fmt.Println("job completed")
		}()
	}

}

func main() {

	target := flag.String("target", "", "email address of target")
	total := flag.Int("total", 1000, "total of email will be sent")
	delay := flag.Int("delay", 2, "delay for each email in second")
	text := flag.String("text", "sorry just dabbling", "the text of email")

	flag.Parse()

	spamTarget(target, total, delay, text)
}
