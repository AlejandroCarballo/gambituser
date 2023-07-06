git add .
git commit -m "ultimo commit"
git push
go build main.go
rm -i main.zip
zip -r main.zip main