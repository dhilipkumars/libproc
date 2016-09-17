##Ping Pong example
A simple implementation of Ping and Pong in erlang as well as go-lang

###Pong
Is the server that responds to any Ping request

###Ping
Is the client which sends a PING request to the PONG server and waits for a 'PONG' response, usually there is more than one PING served by the PONG server

###Implementation
Both PONG and PING are go-routines in golang and actors in erlang implementation,
Channels are used in go-routine and mailboxes are used in erlang implementation. 

###Run
We supply two arguments 
1) Nsec - How long the program should run
2) Nping - how many ping clients should be used.

####Erlang

 `pingpong::start (10, 2)` :- Run for 10 seconds using 2 ping clients

```
9> pingpong:start(10,2).
Ping PID <0.109.0>
ok
Pong count 5246067
10> pingpong:start(10,4).
Ping PID <0.113.0>
ok
Pong count 6014161
11> pingpong:start(10,8).
Ping PID <0.119.0>
ok
Pong count 7657618
12> pingpong:start(10,4).
Ping PID <0.129.0>
ok
Pong count 5920362
13> pingpong:start(10,16).
Ping PID <0.135.0>
ok
Pong count 7806158
14> pingpong:start(10,32). 
Ping PID <0.153.0>

```

####go-lang

`pingpong -NSec 10 -Pings 2` - Run for 10 secs using 2 ping clients

```
D:\>pingpong.exe -NSec 10 -Pings 2
2016/09/17 13:54:29 Pong Count 7695050

D:\>pingpong.exe -NSec 10 -Pings 4
2016/09/17 13:55:13 Pong Count 8009643

D:\>pingpong.exe -NSec 10 -Pings 8
2016/09/17 13:56:02 Pong Count 8104391

D:\>pingpong.exe -NSec 10 -Pings 16
2016/09/17 13:59:02 Pong Count 8215920

D:\>pingpong.exe -NSec 10 -Pings 32
2016/09/17 13:59:38 Pong Count 8542576
```
