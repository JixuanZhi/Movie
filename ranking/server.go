
package main

import (
    "fmt"
    "net/http"
    "flag"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"sort"
)

type Result struct {
	Name  string
	Score float32
}

type Recommender struct {
	ElasticURL string
}

type elasticMatch struct {
	MovieName string `json:"genre"`
}

type elasticQuery struct {
	Match elasticMatch `json:"match"`
}

type elasticQueryBody struct {
	//Source []string `json:"_source"`
	Query elasticQuery `json:"query"`
}

type elasticTotal struct {
	Value int `json:"value"`
	Relation string `json:"relation"`
}

type elasticSource struct {
	Name  string   `json:"movie_name"`
	Genre []string `json:"genre"`
	Score string  `json:"rating"`
}

type elasticResponseShards struct {
	Total int `json:"total"`
	Successful int `json:"successful"`
	Skipped int `json:"skipped"`
	Failed int `json:"failed"`
}

type elasticHit struct {
	Index string `json:"_index"`
	Type string `json:"_type"`
	ID string `json:"_id"`
	Score float32 `json:"_score"`
	Source elasticSource `json:"_source"`
}

type elasticResponseHits struct {
	Total elasticTotal `json:"total"`
	MaxScore float32 `json:"max_score"`
	Hits []elasticHit `json:"hits"`
}

type elasticResponse struct {
	Took int `json:"took"`
	TimeOut bool `json:"time_out"`
	Shards elasticResponseShards `json:"_shards"`
	Hits elasticResponseHits `json:"hits"`
}

func (reco *Recommender) handler() http.Handler{
    handler := http.NewServeMux()
    //handler.HandleFunc("/hello", reco.hello)
    //handler.HandleFunc("/headers", reco.headers)
    handler.HandleFunc("/autocomplete", reco.autocomplete)
    return handler
}

func (reco *Recommender) autocomplete(w http.ResponseWriter, req *http.Request){
	search, ok := req.URL.Query()["search"]
	if !ok || len(search) < 1 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"error\": \"Qeury must contain a param q.\"}")
		return
	}
	query := search[0]

    elasticQuery := elasticQueryBody {
		Query: elasticQuery {
			Match: elasticMatch {
				MovieName: query,
			},
		},
	}



	elasticQueryBody, err := json.Marshal(elasticQuery)
	if (err != nil) {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"Fail to build elastic request.\"}")
		return
	}
    
	// Query elastic for text match.
	//fmt.Fprintf(w, "test: %s\n", reco.ElasticURL)
    requestURL := reco.ElasticURL + "/customer/_search?pretty"

	client := &http.Client{}


	req, err = http.NewRequest("GET", requestURL, bytes.NewBuffer(elasticQueryBody))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"Elastic new request call failed with error "+ err.Error() + "\"}")
		return 
	}


	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	elasticRes, err := client.Do(req)

	//elasticRes, err := http.Post(requestURL, "application/json", bytes.NewBuffer(elasticQueryBody))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"Elastic request call failed with error "+ err.Error() + "\"}")
		return 
	}




    elasticResBytes, err := ioutil.ReadAll(elasticRes.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"error\": \"Cannot read elastic response.\"}")
		return
	}
	if elasticRes.StatusCode != 200 {
		w.WriteHeader(500)
		fmt.Fprintf(w, 
			"{\"error\": \"Elastic request call rejeted with status "+
			elasticRes.Status +
			". Response is\n" + string(elasticResBytes)+ "\n\"}")
	}
   // fmt.Fprintf(w, "re: %s\n", elasticResBytes)

	var elasticResParsed elasticResponse
	err = json.Unmarshal(elasticResBytes, &elasticResParsed)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("error:", err)
		fmt.Fprintf(w, "{\"error\": \"Cannot parse elastic response.\"}")
		return
	}

	rankedMovies := make([]Result, len(elasticResParsed.Hits.Hits))
	for idx, hit := range elasticResParsed.Hits.Hits {
		rankedMovies[idx].Name = hit.Source.Name
		scoreFloat, err := strconv.ParseFloat(hit.Source.Score, 64)
		if err != nil {
			fmt.Println(scoreFloat, err)
		}
		rankedMovies[idx].Score = float32(scoreFloat)
	}
	sort.Slice(rankedMovies, func(i, j int) bool {
		return rankedMovies[i].Score > rankedMovies[j].Score
	})
	responseContent, err := json.Marshal(rankedMovies)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error: Serialize ranked result failed due to "+err.Error())
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseContent)

}

func (reco *Recommender) hello(w http.ResponseWriter, req *http.Request) {
    requestURL := "http://google.com"
    res, err := http.Get(requestURL)
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        fmt.Fprintf(w, "call google failed: %s\n", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    fmt.Printf("client: status code: %d\n", res.StatusCode)
    fmt.Fprintf(w, "%s\n", res)
}

func  (reco *Recommender) headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    elasticURL := flag.String("elastic_url", "", "The url for elastic search service.")
    flag.Parse()
	reco := Recommender{ElasticURL: *elasticURL}
    http.ListenAndServe(":80", reco.handler())
}