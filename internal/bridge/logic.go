package bridge

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/mdns"
)

//maybe rename this file later

// client is a global http client to be used for all http requests
var client = http.Client{
	Timeout: 25 * time.Second,
}

func Find_bridges() tea.Msg {
	var brdgs []Bridge

	// 1- find via mdns
	brdgs, err := lookup_bridge_mdns(brdgs)
	if err != nil {
		log.Println(err)
	}

	// 2- find via querying meetdicovery.com if mdns doesnt find anything
	if len(brdgs) < 1 {
		brdgs, err = lookup_bridge_meethue(brdgs)
		if err != nil {
			return NoBridgeFoundMsg(ErrMsg{err})
		}
	}
	// 3- return the list of bridges
	return BridgeFoundMsg(brdgs[0])
}

func lookup_bridge_mdns(br []Bridge) ([]Bridge, error) {
	entriesCh := make(chan *mdns.ServiceEntry, 5)
	go func() {
		for entry := range entriesCh {
			brdg := Bridge{ID: entry.InfoFields[0], Ip_addr: entry.AddrV4.String(), Port: entry.Port, Info: entry.InfoFields}
			br = append(br, brdg)
		}
	}()
	queryParams := mdns.DefaultParams("_hue._tcp")
	queryParams.Logger = log.New(io.Discard, "", 0)
	queryParams.Entries = entriesCh
	if err := mdns.Query(queryParams); err != nil {
		//deal with error (dont return we have to try other methods first)
		return nil, err
	}
	return br, nil
}
func lookup_bridge_meethue(br []Bridge) ([]Bridge, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://discovery.meethue.com/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&br)
	if err != nil {
		return nil, err
	}
	return br, nil
}
