#Installation

Install this littul utility with:

`go get -u github.com/RobMeades/coap-echo-server`

There is command-line help.  An example command line that leaves the echo server running on port 1000 might be:

`nohup coap-echo-server -p 1000 > udp.log &`

If no port is specified the default COAP port of 5683 is used.