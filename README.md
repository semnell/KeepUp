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
edit the provided config.yaml to reflect your stack, afterwards start the project uing the Makefile and check the /metrics endpoint (defaults to http://127.0.0.1/metrics)

## License
This tool is licensed under the MIT-license. Do whatever you like with it. Go monitor the world. Or don't.

## Is it any good?
I hope so.
