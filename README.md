# Konsultanku-Backend
Konsultanku use Firebase and GCP technology.
## How Our MVP Works?
> "To run this application, we use MySQL in our MVP. In my development, I use Xampp to run MySQL and edit it in phpMyAdmin"
1. **MSE See All Solution from Student**
   - Flutter gets all data comments from (http://localhost:8080/api/v1/mse/comments).
   - Server sends data in JSON form to Flutter.

2. **Mse Send Offer**
   - Flutter hits the endpoint (http://localhost:8080/api/v1/send-offer/{id}) with the POST method.
   - Server makes a new row in table collaboration.

3. **Student Get Notification**
   - Flutter will serve some notifications for students who have joined a team through (http://localhost:8080/api/v1/team/offers).
   - Server will send JSON data with the name of MSE innit.

4. **Student Decisions**
   - Flutter will hit (http://localhost:8080/api/v1/team/decision) with the PUT method to update the row if they accept the offer.
   - If collaboration happens, the server will allow some rules.