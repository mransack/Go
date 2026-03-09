module mransack/hello

go 1.25.0

require (
	mransack/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/example/hello v0.0.0-20250915201037-7f05d217867b // indirect
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace mransack/greetings => ../greetings
