# rate-limit
Rate Limit Notification Service


## Overview
- The application receives requests for notifiying clients and rate limits them by each **notification type** and **email recipient**.
- This means that for any combination of **notification type** and **email recipient** the notifications will be sent by a **configured frequency** 
- I solved this problem completly in code by using **channels** and **go routines**: 
   - there is a channel for each combination of **notification type** and **email recipient**
   - there as a **go routine** for each of these channels that receives the notifications and send email for the recipients, by the configured frequency (keeping in mind that the frequency the inverse of the rate: for a configured frequency of, lets say, 10 seconds, the rate is 0.1/s. I just thought would easier to configure the frequency instead of rate)
- Observations: 
    - this solution solves the problem if the load **if not too high** and we can cope with only one instance of the application
    - for a more scalable solution, we should use message system, for allowing persistency and concurrency control among various instances
    - i tried to solve the problem of "faning out" the rate limit notifications, so the notification model it self may need a lot more information
    - i did not had time for actualy implement the "email sending" function, but i left an interface to be implemented using any smtp service of choice


## Running the Application

# 1. Running with go run
- First run the dependencies with:
`sudo sh compose.sh` 
(this will up the firestore database depency)
- Second run the scrip:
`sh run_go.sh` 


# 2. Running with do docker
- Run the script:
`sh run_docker.sh`
- This will get the application and dependencies up and running, both conteinarized
- This option is good if the client does not have go installed, but i will not allow the client to acess the firestore emulator for checking the application functionality

## Testing the Application
- Run the application with option 1
- There is a single endpoint the application will expose:

```
POST http://localhost:8182/notification-service/notifications

{

    "type": "STATUS | NEWS | MARKETING",   (these are the three basic types of notifications the system allows - for now)
    "content": "some content",             (the content of the notification)
    "email": "some@mail.com"               (the recipent email)
}

```

- Posting a notification to the endpoint will:
  a- Create a notification with status "PENDING" 
  b- After the configured FREQUENCY, for that type of notification (check the env file), the notification will be "sent". This means it status will be changed to "SENT" and the email for the recipent will be sent
  c- If a another notification of **same type** for the **same email** recipient is posted **before the last notification is sent**, this will be rejected (set to status "REJECTED")

- Posting various notificaitons of same type and email recipient
  a- Post repeatadaly a notification of **same type** and **email recipient**
  b- After this, one should notice that there will be a few "SENT" notifications, and a lot of "REJECTED"
  c- The few "SENT" notifications should have the updated_at differing by approximately the configured frequency

- Checking the notification status
  a.Access the firetore emulator in  *http://localhost:4000* and click the
  b.After the first notification is posted, a collection of "notifications" will apear
  c.For checking the notifications, use the filter "status = SENT" in the collection
