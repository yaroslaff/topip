# topip
IPv4 grep and top usage summary for log files.

## Install option 1: precompiled version
~~~
wget -O topip https://github.com/yaroslaff/topip/releases/latest/download/topip-`arch`
chmod +x topip
mv topip /usr/local/bin
~~~

## Install option 2: 
If you have golang installed, use this command
~~~
go install github.com/yaroslaff/topip@latest
~~~
It will install topip inside $GOPATH (often it will be `~/go/bin/topip`)

## Install option 3: Build from sources
~~~
git clone https://github.com/yaroslaff/topip
cd topip
go build .
cp topip /usr/local/bin 
~~~

## Top mode (default)
Default top mode counts occurences of each IP address and print top N (10) results.

~~~
$ topip /var/log/mail.log
...
    8405 92.155.149.266
    8496 35.142.45.298
   66340 127.0.0.1
~~~

or via stdin, something like:
~~~
grep "SASL LOGIN authentication failed:" /var/log/mail.log | topip
~~~

use `-t N` to override default number of lines (10).

##  Grep mode

Grep mode activated with `-g` or `-i` keys. `-g` print whole lines which has any IP address, `-i` print only IP addresses.

Just grep all lines when we have IP:
~~~
# topip -g /var/log/mail.log
Oct  9 05:48:25 mx postfix/smtpd[2166]: connect from unknown[191.211.100.228]
Oct  9 05:48:26 mx postfix/smtpd[2162]: connect from unknown[103.129.202.216]
Oct  9 05:48:30 mx postfix/smtpd[2166]: warning: unknown[191.211.100.228]: SASL LOGIN authentication failed: UGFzc3dvcmQ6
....
~~~

or print only IPs:
~~~
# topip -i /var/log/mail.log
191.211.100.228
103.129.202.216
191.211.100.228
~~~