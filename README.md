## BroFor

**Browser Forensic framework** - special util for extract data from database Google Chrome, Opera and Firefox for incident response purpose.
Worker in two modes:
1) **Forensic mode**. Extracts data and writes data in csv/excel/json file for direct forensic processes.
2) **Live mode**. This mode needs if you run this util on system from which you want to extract data. Additionally, you can use flag -hash and calculate hashes of downloaded files.

Extracted data from browser's database you can save local in csv, json, excel type or sending data in JSON type to remote collector. If you don't choose any output process util print all information to stdout.

Supported browsers: 
* chrome (Google )
* firefox (Firefox)
* opera (Opera)

Supported output:
* csv
* xls 
* console
* json
* remote

Examples:
```bash
# Run autosearch dbs browser Firefox and output to stdout
go run cmd\brofor\main.go -live -b firefox 

# Run autosearch dbs browser Firefox and output to excel file
go run cmd\brofor\main.go -live -b firefox -o xls

```

Plans:
* processor for checking domain names in TI feed lists
* add support browsers: MS Edge, Yandex Browser
* add tests