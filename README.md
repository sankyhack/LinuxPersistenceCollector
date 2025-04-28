# LinuxPersistenceCollector

Linux Persistence Collector is all about collecting and analyzing all the malware persistence entries in Linux OS.
It will collect all the entries and store it in directory named "Collector" created in same directory from where tool will be run.
Run this tool with root priviledges to avoid permission issues.
persistencelist.txt file contains all the files and folders which malwares can use to maintain persistence

After collecting files, Analyze function will look for keywords like "tmp" "http" "ExecStart" "www" in those files and store in file "MatchingKeywords.txt". 
You can add your own keywords refer anayze.go file (line 45 ) 
"tmp" -> Executable using tmp path
"ExecStart" -> Services files in linux uses "ExecStart" to point binary path ( malware can maintain persistence via services )
Analyze function also look for IP address pattern in those files.

One more file will get created named "IP_Domain_Extract.txt" this will have all the IP and Domain for easier analysis.
You will see known Domains / Private IP addresses as well, as some files have those as part of comments (ignore those).

So far this tool includes above mentioned things only.
As a future improvement I will add command line flags. and Any other suggestion from you :) 

#IR #LinuxPersistence #SOC #persistence #ThreatHunting

This tool is inspired by blog Persistence series by Ruben Groenewoud and PANIX script
Thanks Ruben for amazing series !!!
https://github.com/Aegrah/PANIX
