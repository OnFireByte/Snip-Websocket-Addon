# Snip-Websocket-Addon

A very small program that build with Go. Its purpose is to serve now-playing track's data via web-sockets, work with (and require) [Snip](https://github.com/dlrudie/Snip) as its dependency (Like the name suggests).

This project also includes a simple webpage that can be used as an overlay in live streaming.

# How to use

1. move the executable file to Snip's folder
2. In Snip setting, enable "Save infomation separately"
3. Run the executable file
4. The websocket server can be accessed via localhost:5487/ws, the track's artwork also can be accessed via localhost:5487/artwork
5. (optional) You can get simple now-playing overlay by adding browser layer to your live stream application, then point to the html file.
