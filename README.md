# YouTube API Experiment

## Setup
1. Clone the repository: `git clone https://github.com/jhitesh/youtube_api_experiment.git`
2. Install the dependencies: `go mod download`
3. Set up the PostgreSQL Database and fill details in `database_credentials.yaml` file
4. Create the relevent table in the database using following command:
```
CREATE TABLE videos (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    published_at timestamp NOT NULL
);
CREATE INDEX videos_title_description_index ON videos (title, description);
```
5. Enter the number of videos that you want to get per page from the APIs in file `app_constants.yaml`
6. Provide API Key and the `query` for YouTube API in file `query_params_values.yaml`
7. Enter the time interval values to fetch data from YouTube in file `youtube_fetch_worker_values.yaml`
8. Start the server `go run main/main.go`

### NOTE
All the details about the API Endpoints and the system can be found [here](project_description.md)