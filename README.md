# TCPChat
This project is simple form of Unix utility netcat which enables the creation of TCP and UDP connections.
Current version supports only TCP connections.

User cn create a server specifying the port.

The project is written in Golang.

<h3>USAGE: ./TCPChat $port</h3>

In order to run the server the port should be speified. If no port is given server  runs on a default port 2525.

When client joins the server the linux logo is shown and asks for the name on a client-side:
<pre>
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: 
</pre>

As soon as user inserts the name, user recieves all previous messages and will be able to send and recieve messages.
