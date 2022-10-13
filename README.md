# topip
IPv4 grep and top usage summary for log files.

## Build / install
~~~
git clone https://github.com/yaroslaff/topip
cd topip
go build
cp topip /usr/local/bin 
~~~

## Top mode (default)
Default top mode counts occurences of each IP address and print top N (10) results.

~~~
topip /var/log/mail.log
~~~

or via stdin, something like:
~~~
grep "SASL LOGIN authentication failed:" /var/log/mail.log | topip
~~~

##  Grep mode

Grep mode activated with `-g` or `-i` keys. `-g` print whole lines which has any IP address, `-i` print only IP addresses.
