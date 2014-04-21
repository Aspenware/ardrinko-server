# Arduinko-Server

A golang server that sits on the network and listens for raw telemmetry from
Arduinko (via UDP). Exposes a restful API that offers this information in a
friendlier format.

## Compiling/Running

Install Go. Run:

    go get

And then:

    go build

Edit config.ini, then:

    ./ardrinko-server
