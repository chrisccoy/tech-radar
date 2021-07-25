aws s3 cp ./techradar.json s3://ccoy-test/techradar.json
aws s3api put-object-acl --bucket ccoy-test --key techradar.json --acl public-read
