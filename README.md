# interview


Can not create unit test case as we are refreshing on-memory data after every 60 sec. Which will be no longer be valid after each minute.


Post /transactions
{
"amount":"2",
"timestamp":"2022-11-09T16:50:00.312Z"  
}

(Time should be in UTC time zone, To get latest UTC time : https://www.google.com/search?q=current+utc+timing&oq=current+utc+timing&aqs=chrome..69i57.4783j0j4&sourceid=chrome&ie=UTF-8)

GET /statistics

DELETE  /transactions
