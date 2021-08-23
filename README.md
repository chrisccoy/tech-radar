# Tech Radar

This is a sample project for how one might generate a dataset to be used by the Backstage tech radar [plugin](https://github.com/backstage/backstage/tree/master/plugins/tech-radar).
The Tech Radar expects a JSON object matching the TechRadar type defined [here](pkg/types/types.go)


This project is just a sample for demonstration purposes. There was no effort put into design and scaffolding of this project.

## Building

From the root directory run:
`make`

## Running
```
GITLAB_TOKEN=<your git lab token> GITLAB_URL=<gitlab instance> bin/<OS-Type>/languages
# Make will build an binary based on your OS - running:  find ./bin -name 'languages' should show the location for the binary on your machine
```
This command will crawl through a gitlab instance and collect software language information across projects. Once complete a json file is produced that can be pushed to s3 for use by the Tech Radar component in our Backstage application.

## s3
The Tech Radar defined [here](https://git.ecd.axway.org/jdavanne/backstage/-/blob/master/packages/app/src/lib/AxwayTechClient.ts) is implemented to acess an s3 bucket for rendering. The Client code simply fetches the object over https, 
so we need to make sure the object is present prior to rendering the page.

This assumes that you have the aws cli installed and have properly authenticated using aws-vault.

**s3create.sh** shows a simple creation of an s3 bucket

**s3push.sh** will upload the json object to the bucket and set the ACL's allowing for public read access





