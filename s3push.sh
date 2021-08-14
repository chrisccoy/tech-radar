aws s3 cp ./techradar.json s3://ccoy-tech-radar/techradar.json
aws s3api put-object-acl --bucket ccoy-tech-radar --key techradar.json --acl public-read
