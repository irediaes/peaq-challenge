# Describe security risks with unpacking buffers.

## Buffer Overflow
This happens when the client/server is sending data that is out of bound with the size of the data required. Unpacking such buffer in a non type-safe language could result the program writing the extra data to other variables in the stack; thereby making the entire program vulnerable for an attack. In the instance where the server instigate the overflow, this can result in private information in the stack being written to public variables available to users.

```
# code block

func main() {
    // public information saved on the server
    harmlessData := [4]byte{'A', 'A', 'A', 'A'}
    // however, this could be critical private information such as a private key
    secret := [17]byte{'l', '3', '3', 't', '-', 'h', '4', 'x', 'x', '0', 'r', '-', 'w', '1', 'n', 's', '!'}

    // (accidentally) cast harmless buffer into a new buffer type of wrong size
    var dangerousData = (*[8+17]byte)(unsafe.Pointer(&harmlessData[0]))

    // print (misused) buffer
    fmt.Println(string((*dangerousData)[:]))
}

```
[CODE SOURCE::](https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga)

## Data Corruption
Especially with public gRPC APIs, buffers could be recieved in an unexpected format not documented; probably based on the fact that some changes was made during protobuf code generation. Corrupted buffer is a huge security threat to any gRPC server with a poor workflow in handling buffers.

## Denial of Service Attack
In an insecured or possibly secured(subjective based on the developer's experience on security) gRPc implementation where priviledges are granted to clients to call functions on the server, a DoS attack can be exploited when the opportunity arises. An attacker can send corrupted buffers that might trigger a funtion to launch a DoS or possibly send buffers that set the procedure into an infinite loop during an unpacking process.



