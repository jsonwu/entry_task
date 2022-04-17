mkdir -p output/bin output/conf

cp script/bootstrap.sh output
chmod +x output/bootstrap.sh

go build -o output/bin/entry_task
