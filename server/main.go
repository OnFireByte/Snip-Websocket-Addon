package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"icon"

	"github.com/fsnotify/fsnotify"
	"github.com/getlantern/systray"
	"github.com/gorilla/websocket"
)

const(
    port = ":5487"
)

type track struct {
    Name string `json:"name"`
    Artist string `json:"artist"`
    Album string `json:"album"`
    Image string `json:"image"`
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 128,
    WriteBufferSize: 512,
    CheckOrigin : func(r *http.Request) bool { return true },

}

func sentTrackData(conn *websocket.Conn, data track) {    
    err := conn.WriteJSON(data)
    if err != nil {
        fmt.Println("Can't send")
    } else {
        fmt.Println("Sending")
    }
    
}

func getTrackData() (track,error) {
    track_warped,err1 := os.ReadFile("Snip_Track.txt")
    artist_warped,err2 := os.ReadFile("Snip_Artist.txt")
    album_warped,err3 := os.ReadFile("Snip_Album.txt")
    image := fmt.Sprintf("localhost%s/artwork",port)

    if err1 != nil{
        return track{},err1
    }
    if err2 != nil{
        return track{},err2
    }
    if err3 != nil{
        return track{},err3
    }

    return track{
        Name: string(track_warped),
        Artist: string(artist_warped),
        Album: string(album_warped),
        Image: image,
    },nil
}

func wsEndpoint(w http.ResponseWriter, r *http.Request){

    ws,err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }
    defer ws.Close()
    log.Println("Client connected")

    track,err := getTrackData()
    prepTrackName := track.Name
    if err != nil {
        log.Fatal(err)
    }
    sentTrackData(ws,track)

    watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)


	//
	go func() {
		for {
			select {
			// watch for events
			case event,ok := <-watcher.Events:
				
                if !ok {
                    continue
                }
                if event.Op&fsnotify.Write != fsnotify.Write {
                    continue
                }
                track, err := getTrackData()
                if err != nil {
                    log.Println(err)
                }
                if prepTrackName == track.Name {
                    continue
                }

                prepTrackName = track.Name
                fmt.Println("EVENT!",track)
                sentTrackData(ws,track)  
				
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("Snip_TrackId.txt"); err != nil {
		fmt.Println("ERROR", err)
	}

    <-done
}

func setupRoutes() {

    http.HandleFunc("/artwork", func (w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "Snip_Artwork.jpg")})
    http.HandleFunc("/ws", wsEndpoint)
}

func main() {
    fmt.Println("Starting systray...")
    go systray.Run(onReady, onExit)
    fmt.Println("Starting server...")
    setupRoutes()
    log.Fatal(http.ListenAndServe(port, nil))
    
}

func onReady() {
    systray.SetTemplateIcon(icon.Data, icon.Data)
    systray.SetIcon(icon.Data)
    systray.SetTitle("Now-Playing-Server")
    systray.SetTooltip("Now-Playing-Server")
    mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
    go func() {
        for {
            select {
                case <- mQuit.ClickedCh:
                    systray.Quit()
                    return
            }
        }
    }()
}

func onExit() {
    os.Exit(3)
}
