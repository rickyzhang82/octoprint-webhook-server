Octoprint Web Hook Server
=========================

A simple web hook server to handle Octoprint event.

How
===

1. To build, run `make`.
1. To run app, run `./octoprint-webhook-server.o`. It create a web hook server with URL `http://server:12000/done?delay=xxx`.
1. It runs command defined `doneCommand` in `handler/done.go` with `xxx` minutes later from now.
1. To override the listening port, set environment variable `WEBHOOK_SEREVER_PORT`.
1. To override the done command, set environment variable `DONE_COMMAND`.