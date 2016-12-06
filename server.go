/* Entry point for COAP echo server running over UDP.
 *
 * This code based on the UDP echo server code also in this repo
 * and on the examples here:
 * https://github.com/dustin/go-coap
 *
 * A COAP echo is constituted by a confirmable message,
 * which this server will confirm, sending back the content it received.
 */
 
package main

import (
    "net"
    "fmt"
    "os"
    "flag"
    "log"
    "github.com/dustin/go-coap"
)

//--------------------------------------------------------------------
// Types
//--------------------------------------------------------------------

//--------------------------------------------------------------------
// Variables
//--------------------------------------------------------------------

var numPackets int

// Command-line flags
var pPort = flag.String ("p", "5683", "the UDP port to listen on.")
var Usage = func() {
    fmt.Fprintf(os.Stderr, "\n%s: run the COAP echo server.  Usage:\n", os.Args[0])
        flag.PrintDefaults()
    }

//--------------------------------------------------------------------
// Functions
//--------------------------------------------------------------------

func coapHandler (l *net.UDPConn, pAddress *net.UDPAddr, pMessage *coap.Message) *coap.Message {
    var pResponse *coap.Message
    
	numPackets++;
	log.Printf("%d: %v <-> %v: %#v", numPackets, pAddress, pMessage.Path(), pMessage)
	if pMessage.IsConfirmable() {
		pResponse = &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: pMessage.MessageID,
			Token:     pMessage.Token,
			Payload:   pMessage.Payload,
		}
		pResponse.SetOption(coap.ContentFormat, coap.TextPlain)

		log.Printf("Transmitting %#v", pResponse)
	}
	
	return pResponse
}

// Entry point
func main() {

    // Deal with the command-line parameters
    flag.Parse()
    
    // Set up logging
    log.SetFlags(log.LstdFlags)
    
    // Say what we're doing
    fmt.Printf("Responding to COAP packets (e.g. confirmable messages) received on port %s.\n", *pPort)
    
    // Run the server
    err := coap.ListenAndServe("udp", ":" + *pPort, coap.FuncHandler(coapHandler))
        
    if err != nil {
        fmt.Printf("Couldn't start COAP echo server on port %s (%s).\n", *pPort, err.Error())
    }            
}

// End Of File
