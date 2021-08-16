# Tech Radar

This is a sample project for how one might generate a dataset to be used by the Backstage tech radar [plugin](https://github.com/backstage/backstage/tree/master/plugins/tech-radar).
The Tech Radar expects a JSON object matching the TechRadar type defined [here](pkg/types/types.go)

This project is just a sample for demponstration purposes. There was no effort put into design and scaffolding of this project.

##Building

From the root directory run:
`make`

## Running
```
GITLAB_TOKEN=<your git lab token> GITLAB_URL=<gitlab instance> bin/darwin_amd64/languages
```
This will generate a json file that can be pushed to s3

## s3
The Tech Radar defined [here](https://git.ecd.axway.org/jdavanne/backstage/-/blob/master/packages/app/src/lib/AxwayTechClient.ts) is implemented to acess an s3 bucket for rendering. The Client code simply fetches the object over https, 
so we need to make sure the object is present prior to rendering the page.

This assumes that you have the aws cli installed and have properly authenticated using aws-vault.

**s3create.sh** shows a simple creation of an s3 bucket

**s3push.sh** will upload the json object to the bucket and set the acl's allowing for public read access





