## BroFor

**Browser Forensic framework** - special util for extract data from database Google Chrome and Firefox for incident response purpouse.
Worker in two modes:
1) **Forensic mode**. Extracts data and writes data in csv/excel/json file for direct forensic process
2) **Threat Hunting mode**. This mode works like a daemon. Every 5 sec util take data from browser's database and send data in syslog type to remote log server