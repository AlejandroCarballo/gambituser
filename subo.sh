git add .
git commit -m "ultimo commit"
git push
set GOOS=linux
Set GOARCH=amd64
go build main.go
rm main.zip
zip main.zip main