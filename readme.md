# TCP Chat (client/server)

This project is a simple TCP server implemented in Go. It listens for incoming TCP connections on port 8888. For each connection, it creates a new client and starts a separate goroutine to handle reading input from that client.

If there's an error starting the server or accepting a connection, it logs the error.

The server continues to accept and process new connections until it is manually stopped.

# Preview

![preview](https://i.ibb.co/y0Yccdj/preview.png)