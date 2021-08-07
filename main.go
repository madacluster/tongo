package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/common-nighthawk/go-figure"
	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

// Config form metri
type Config struct {
	Props struct {
		PageProps struct {
			SerializedInitialState string `json:"serializedInitialState"`
		} `json:"pageProps"`
	} `JSON:"props"`
}

// State from metri
type State struct {
	Questions []struct {
		ID      string `json:"id"`
		Choices []struct {
			ID int `json:"id"`
		} `json:"choices"`
	} `json:"questions"`
	Pace struct {
		Presenter struct {
			PresenterID string `json:"presenterId"`
		} `json:"presenter"`
	} `json:"pace"`
}

// Votes object
type Votes map[string][2]int

// ResponseVotes object
type ResponseVotes struct {
	QuestionType string `json:"question_type"`
	Vote         Votes  `json:"vote"`
}

func main() {
	var loop int
	var value int
	var url string

	var rootCmd = &cobra.Command{
		Use:   "tongo",
		Short: "Vote several time to menti.com/",
		Long:  `Vote several time to menti.com/`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			myFigure := figure.NewColorFigure("Tongo", "", "green", true)
			myFigure.Print()
			log.Printf("This POOL=%s will be vote %d times with a %d\n", url, loop, value)
			vote(loop, value, url)
		},
	}

	defaultLoop, _ := strconv.Atoi(getEnv("TONGO_LOOP", "1"))
	defaultValue, _ := strconv.Atoi(getEnv("TONGO_VALUE", "1"))
	defaultHTTP := getEnv("TONGO_MENTI_URL", "")
	rootCmd.Flags().IntVarP(&loop, "loop", "l", defaultLoop, "times to echo the input")
	rootCmd.Flags().IntVarP(&value, "value", "v", defaultValue, "times to echo the input")
	rootCmd.Flags().StringVarP(&url, "url", "u", defaultHTTP, "url (required) Ex: https://www.menti.com/1ct2pwd8ba")
	if defaultHTTP == "" {
		rootCmd.MarkPersistentFlagRequired("url")
	}
	rootCmd.Execute()
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func vote(loop, value int, url string) {
	c := colly.NewCollector()
	c.OnHTML("script", func(e *colly.HTMLElement) {
		presenterID, votes, err := getPresenterIDAndVotes(e.Text, value)
		if err == nil {
			var wg sync.WaitGroup
			for i := 0; i < loop; i++ {
				wg.Add(1)
				go hackTheVote(presenterID, url, votes, &wg, value, i)
			}
			wg.Wait()
			log.Printf("TONGAZO HAS BEEN FINISHED\n")
		}
	})
	c.OnError(func(r *colly.Response, e error) {
		log.Panic(e)
	})
	c.Visit(url)
}

func hackTheVote(presenterID, url string, votes Votes, wg *sync.WaitGroup, value, id int) {
	defer wg.Done()
	identifier, err := getIdentifier(url)
	if err != nil {
		log.Panic(err)
	}
	requestBody := ResponseVotes{
		QuestionType: "scales",
		Vote:         votes,
	}
	jsonStr, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "https://www.menti.com/core/votes/"+presenterID, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("origin", "https://menti.com")
	req.Header.Set("referer", url)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	req.Header.Set("x-identifier", identifier)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("HAHAHAHAHA LOOKS LIKE ERROR, LOOKS WHAT YOU DID ┐('～`;)┌")
		log.Println(string(body))
	} else {
		log.Printf("INDEX=%d IDENTIFIER=%s POOL=%s PRESENTER=%s VOTE=%d\n", id, identifier, url, presenterID, value)
	}

}

func getIdentifier(url string) (string, error) {
	jsonStr := []byte(`{}`)
	req, err := http.NewRequest("POST", "https://www.menti.com/core/identifiers", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("origin", "https://menti.com")
	req.Header.Set("referer", url)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("yo I can't get the Identifier (⊙_⊙)")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response["identifier"], nil
}

func getPresenterIDAndVotes(text string, value int) (string, Votes, error) {
	var grades Config
	var props State
	err := json.Unmarshal([]byte(text), &grades)
	if err == nil {
		propsRaw := grades.Props.PageProps.SerializedInitialState
		err := json.Unmarshal([]byte(propsRaw), &props)
		if err == nil {
			return props.Pace.Presenter.PresenterID, getChoices(props, value), nil
		}
	}
	return "", nil, err
}

func getChoices(props State, value int) Votes {
	var vote = Votes{}
	for _, question := range props.Questions {
		for _, choice := range question.Choices {
			vote[strconv.Itoa(choice.ID)] = [2]int{value, 1}
		}
	}
	return vote
}
