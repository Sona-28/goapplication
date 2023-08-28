# PROTOC

1. The dataformats are defined in the form of messages. 
2. Every member in the message will be having a index associated with it.
3. Messages can take a variety of datatypes - strings, integers, booleans, arrays and other message types
4. Arrays are represented with a keyword called repeated

* Protobuf files will never implement any logic. They will only talk about message exchanging formats and supported methods. 
    1. We will create a server.
    2. We will give the implementation to the rpc method.
    3. We will deploy a server on a port
    4. We will create a client and we will call the server method from the client.