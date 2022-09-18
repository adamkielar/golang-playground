module app/hello-2

go 1.18

replace app/greetings => ../greetings

require app/greetings v0.0.0-00010101000000-000000000000
