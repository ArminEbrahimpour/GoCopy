# GoCopy
#### Simple application to save the output of your terminal command into your clipboard 
### Usage:
```
git clone https://github.com/ArminEbrahimpour/GoCopy

cd GoCopy/cmd

go build -o gocopy

# copying to the /usr/local/bin 
sudo cp gocopy /usr/local/bin

# example
ls / | gocopy -c

## go anywhere you want and press ctrl+v and you see the clipboard saved the 

```
### notes:
#### the process will close after 2 minutes itself or if you don't need the application you can always ctl+c it or you could simply run it in background in case you need the terminal 
```
ls / | gocopy -c &
```
