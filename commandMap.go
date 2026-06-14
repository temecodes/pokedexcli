package main


import(
	"fmt"
	"net/http"
	"encoding/json"
)

type AreaName struct{
	Results []Location `json:"results"`
}

type Location struct{
	Name string `json:"name"`
}	

func commandMap() error{

	apiURL := "https://pokeapi.co/api/v2/location-area"
	res, err := http.Get(apiURL)
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
	return nil
}
