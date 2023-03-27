# Project Description
In this project, we will develop an async worker to fetch videos data from YouTube and store it in our database. On top of this data, we want to develop some APIs to fetch OR search videos

# Requirements
1. Our service should call the YouTube API in background (async) repeatedly in some time interval (say 10 sec), fetch the latest videos for a predefined query and store the data of videos (specifically - Video Title, Video Description, Publishing Datetime, thumbnails URLs and any other necessary fields) in a database with proper indexes
2. Admin of the service should be able to see which videos are stored in the database in paginated form and in descending order of published datetime
3. We should be able to search specific videos from the stored videos using their title and description
4. Dockerize the project
5. The project should be scalable and optimised


## Video Data Format in YouTube's API
```
{
  "kind": "youtube#searchResult",
  "etag": etag, //Used to identify if a resource is changed
  "id": {
    "kind": string,
    "videoId": string,
    "channelId": string,
    "playlistId": string
  },
  "snippet": {
    "publishedAt": datetime,
    "channelId": string,
    "title": string,
    "description": string,
    "thumbnails": {
      (key): { // key is one of [default, medium, high, standard, maxres]
        "url": string,
        "width": unsigned integer,
        "height": unsigned integer
      }
    },
    "channelTitle": string,
    "liveBroadcastContent": string
  }
}
```

# Database (PostgreSQL)
### Videos Table
```
CREATE TABLE videos (
id TEXT PRIMARY KEY, // Stores Video ID here
title TEXT NOT NULL,
description TEXT NOT NULL,
published_at DATETIME NOT NULL
);
CREATE INDEX videos_title_description_index ON videos (title, description);
```

# API Endpoints
1. `/videos/all?page={page_number}` GET - Returns a list of all stored videos in paginated form
2. `/videos/search?q={search_query}&page={page_number}` GET - Returns a list of videos having partially/fully matching title OR description with the search_query from the database

# Worker Details
Worker keeps requesting data on YouTube's API `GET https://www.googleapis.com/youtube/v3/search` with following other details
Query Parameters:
1. part - `[snippet]` //**required**
2. maxResults - `[0, 50]` //_default is 25_
3. order - `[date, rating, relevance, title, videoCount, viewCount]`. //_We will use `date`_
4. pageToken - Keep updating from the YouTube's response
5. publishedAfter - RFC 3339 formatted date-time value (1970-01-01T00:00:00Z)
6. q - Query parameter to search for. //_We can use `official` as it might be most frequently used_
7. type - `[channel, playlist, video]` //_We will use `video`_

## NOTE
Send API Key in the `key` query parameter to get the data