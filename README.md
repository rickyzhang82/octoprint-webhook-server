Octoprint Web Hook Server
=========================

A simple web hook server to handle Octoprint event.

How
===

1. To build, run `make`.
1. To run app, run `./app`. It create a web hook at `http://server:12000/done?delay=xxx`.
1. It runs command defined `doneCommand` in `handler/done.go` with `xxx` minutes later from now.