module hello

go 1.20

replace greetings => ../greetings

replace greetings/greetings => ../greetings

require greetings/greetings v0.0.0-00010101000000-000000000000
