# What is it?
It's a tiny webserver sitting on port 1988 that allows you to start and stop wireshark captures.

# Why?
Windows network drivers often batch up packets and deliver them to wireshark in a way that looses timing precision.

It might look like a bunch of packets have arrived at the same time when actually they have arrived at different times that a very close to each other.

When trying to analyse network packet timings (specifically for AVB related packets), it is essential that as much of the timing information is preserved.

Tools costing thousands are normally used but I got good enough results with an RPi4 and an ETAP-2003.

# Setup
+ Install Google Go.
+ Compile the code with `go build wireshark_server.go`
+ Run the server with `./wireshark_server`
+ Navigate to the network name or IP address of your PI e.g. `http://myPiName:1988/`

You will see options to start & stop captures, adjust access permissions and also shutdown the Pi.

My work flow uses the server to create captures in a samba share that is accessible from windows. You may wish to extend the code to support directly downloading the capture from the web server.

