# topip
IPv4 grep and top usage summary for log files.

## Build / install
~~~
git clone https://github.com/yaroslaff/topip
cd topip
go build
cp topip /usr/local/bin 
~~~

##  grep IPs

~~~
topip < /var/log/mail.log
142.98.11.27
159.60.83.126
211.79.60.219
~~~

or `topip -i /var/log/mail.log`.

If you want to see full lines use -f (only lines with IP addresses are printed):
~~~
topip -f < /var/log/mail.log
Oct  2 00:16:06 mx postfix/smtpd[32199]: lost connection after AUTH from unknown[208.79.20.219]
Oct  2 00:16:06 mx postfix/smtpd[32199]: disconnect from unknown[208.79.20.219] ehlo=1 auth=0/1 commands=1/2
~~~

## Top

Use `-t NUM` to top TOP NUM IPv4 addresses:
~~~
$ ./topip -t 5 < /var/log/mail.log 
      12 52.28.217.103
      27 35.52.135.98
      68 127.0.0.1
     100 18.132.33.18
     116 13.110.205.132
~~~

