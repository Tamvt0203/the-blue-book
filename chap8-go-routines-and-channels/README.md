- Exercise 8.1: Modify clock2 to accept a port number, and write a program, clockwall, that
acts as a client of several clock ser vers at once, reading the times from each one and displaying
the results in a table, akin to the wall of clocks seen in some buiness offices. (clock2 & clockwall)
- Exercise 8.3: In netcat3, the interface value conn has the concrete typ e *net.TCPConn, which
represents a TCP connection. A TCP connetion consists of two halves that may be closed
independently using its CloseRead and CloseWrite methods. Modify the main goroutine of
netcat3 to close only the write half of the connection so that the program will continue to
print the final echoes from the reverb1 server even after the standard input has been closed. (netcat3)
- Exercise 8.4: Modify the reverb2 server to use a sync.WaitGroup per connection to count
the number of active echo goroutines. When it falls to zero, clos e the write half of the TCP
connection as described in Exercise 8.3. Verify that your modified netcat3 client from that
exercise waits for the final echoes of multiple con current shouts, even after the standard input
has been closed. (reverb2)
- Exercise 8.6: Add depth-limiting to the concurrent crawler. That is, if the user sets -depth=3,
then only URLs reachable by at most three links will be fetched. (crawl4)
- Exercise 8.8: Use a select statement, add a timeout to the echo server from section 8.3 so that it disconnects any clients that shouts nothing within 10 seconds. (reverb3 & netcat3)
