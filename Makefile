test:
	mc cp spread/resources/*.sjs tw-test/template/spreadjs/
prod:
	mc cp spread/resources/*.sjs tw-prod/template/spreadjs/
ssl:
	mkcert -key-file key.pem -cert-file cert.pem *.tvoc.site localhost 127.0.0.1 ::1
