git add .
git commit -m "ultimo commit"
git push
go build main.go
rm main.zip
zip -r main.zip -i main