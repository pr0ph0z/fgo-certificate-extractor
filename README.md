# FGO Certificate Extractor

## Description
As what the title says, this repository is a tool to extract the certificate from FGO file data, to be exact the `54cc790bf952ea710ed7e8be08049531` file (both NA and JP). What's this for? ¯\_(ツ)_/¯

## Origin
I translated [this snippet](https://dotnetfiddle.net/ug7C0x) to Go and made it a CLI tool.

## Usage
1. Clone this repository
2. Run `go build` to build the executable file
3. Run `./fgo-cert-extractor -f <path-to-certificate-file>` to extract the certificate
4. Or alternatively, you can put the content of the file directly as the second argument with `./fgo-cert-extractor <certificate-file-content>`
5. Or alternatively, you can pipe the content of the file with `cat <path-to-certificate-file> | ./fgo-cert-extractor`