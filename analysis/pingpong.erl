-module(pingpong).

-export([start/2, ping/2, pong/1]).

ping(0, Pong_PID) ->
    Pong_PID ! finished,
	io:format("Pong PID ~20lp~n", [Pong_PID]);

ping(N, Pong_PID) ->
    Pong_PID ! {ping, self()},
    receive
		finished ->
			true;
        pong ->
			ping(N + 1, Pong_PID)
	end.

pong(N) ->
    receive
        finished ->
            io:format("Pong count ~20lp~n", [N]);
        {ping, Ping_PID} ->
            Ping_PID ! pong,
			pong(N+1)
	end.

start_ping(Nsec, 0, Pong_PID) ->
	true;
	
start_ping(Nsec, Npings, Pong_PID) ->
	Ping_PID = spawn(tut15, ping, [1, Pong_PID]),
	timer:send_after(Nsec * 1000, Ping_PID, finished),
	start_ping(Nsec, Npings - 1, Pong_PID).

start(Nsec, Npings) ->
    Pong_PID = spawn(tut15, pong, [0]),
    start_ping(Nsec, Npings, Pong_PID),
	timer:send_after(Nsec * 1000,Pong_PID,finished),
	io:format("Ping PID ~20lp~n", [Pong_PID]).