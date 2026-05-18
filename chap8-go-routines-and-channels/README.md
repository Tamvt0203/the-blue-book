- Exercis e 8.1: Modify clock2 to accept a port number, and write a program, clockwall, that
acts as a client of several clock ser vers at once, reading the times from each one and displaying
the results in a table, akin to the wall of clocks seen in some bu iness offices. (clock2 & clockwall)
- Exercis e 8.3: In netcat3, the interface value conn has the concrete typ e *net.TCPConn, which
represents a TCP connection. A TCP connetion consists of two halves that may be closed
independently using its CloseRead and CloseWrite methods. Modify the main goroutine of
netcat3 to close only the write half of the connection so that the program will continue to
print the final echoes from the reverb1 server even after the standard input has been closed. (netcat3)
