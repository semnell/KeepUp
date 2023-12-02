# KeepUp
A deadly simple uptime monitoring tool. 
Designed to be lean and mean and distributed to hundreds of workers each handling requests.


If you like KeepUp, please consider starring it.

### Roadmap
- Setup docker 
- Setup helm

### Features
- Deadly simple config language
- Promethues metrics endpoint for easy integration into any monitoring system
- Only external dependency is Faktory, a single binary job queue
- Rich logging

## Basic usage
- Copy the .env.example to .env
- run either `make run` or make `run-server-standalone` + `make run-worker-standalone`
- Check metrics at 127.0.0.1:8080/metrics and the Faktory server at 127.0.0.1:7420

## License
This tool is licensed under the MIT-license. Do whatever you like with it. Go monitor the world. Or don't.

## Is it any good?
I hope so.
