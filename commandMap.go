package main


import(
	"fmt"
	"net/http"
	"encoding/json"
)

type AreaName struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []Location `json:"results"`
}

type Location struct{
	Name string `json:"name"`
	Url string `json:"url"`
}	

func commandMap(cfg *config) error{
	
	if cfg.nextLoc == nil{
		fmt.Println("You are on the last page")
		return nil
	}

	res, err := http.Get(*cfg.nextLoc)
	if err != nil{
		return fmt.Errorf("request failed %w", err)
	}

	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK{
		return fmt.Errorf("Non-OK HTTP status: %s", res.Status)
	}
	var info AreaName

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&info); err != nil{
		return fmt.Errorf("Decoder didnt run correctly %w", err)
	}

	infoOfInfo := info.Results

	for _, i := range infoOfInfo{
		fmt.Println(i.Name)
	}
	cfg.nextLoc = info.Next
	cfg.prevLoc = info.Previous
	return nil
}
