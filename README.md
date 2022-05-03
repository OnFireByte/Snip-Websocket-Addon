# Snip-Websocket-Addon

A very small program that build with Go. Its purpose is to serve now-playing track's data via web-sockets, work with (and require) [Snip](https://github.com/dlrudie/Snip) as its dependency (Like the name suggests).

This project also includes a simple webpage that can be used as an overlay in live streaming.

# How to use

1. move the executable file to Snip's folder
2. In Snip setting, enable "Save infomation separately"
3. Run the executable file
4. The websocket server can be accessed via localhost:5487/ws, the track's artwork also can be accessed via localhost:5487/artwork
5. (optional) You can get simple now-playing overlay by adding browser layer to your live stream application, then point to the html file.

# How to compile

1. Install 2goarray
    ```bash
    go install github.com/cratonica/2goarray
    ```
2. Go to /server/icon and run bat/sh file
    ### Windows (Powershell)
    ```powershell
    .\server\icon\make_icon.bat
    ```
    ### Linux
    ```bash
    .\server\icon\make_icon.sh
    ```
3. Install go-winres and run (Optional, require for changing icon)
    ```bash
    go install github.com/tc-hib/go-winres@latest
    go-winres make
    ```
4. compile go file

    ```bash
    # without console *recomended*
    go build -ldflags -H=windowsgui

    # with console
    go build
    ```
