## **Bet Settlement Engine Documentation**

### **Overview**

The **Bet Settlement Engine** is a Go-based service that handles the settlement of bets, user account management, and bet transaction processing. It interacts with an underlying data store, processes bet settlements, handles user operations, and provides APIs to manage events, bets, bet transactions, and users.

This documentation provides an overview of the architecture, the API routes, the Docker setup, and development guidelines.

---

### **Key Features**

- **Bet Handling**: Handles the creation, settlement, and cancellation of bets.
- **User Management**: Manages user data, including balance and user actions.
- **Bet Transaction**: Manages the creation and handling of bet transactions.
- **Event Handling**: Handles event data, including creating and fetching events related to bets.
- **API Endpoints**: Exposes various endpoints to interact with the system, including the creation of users, bets, and transactions.

---

### **API Routes**

#### **1. Bet Routes**

- **POST /bet**: Create a new bet.
   - Request body: `{"betID": string, "amount": float, "userID": string, "status": string}`
   - Response: `201 Created`

- **GET /bet/{betID}**: Get details of a specific bet.
   - Path param: `betID`
   - Response: `200 OK` with bet details.

- **PUT /bet/{betID}/settle**: Settle a specific bet.
   - Path param: `betID`
   - Request body: `{"result": string, "amountWon": float}`
   - Response: `200 OK`

- **PUT /bet/{betID}/cancel**: Cancel a specific bet.
   - Path param: `betID`
   - Response: `200 OK`

#### **2. Bet Transaction Routes**

- **POST /bet_transaction**: Create a new bet transaction.
   - Request body: `{"betID": string, "transactionAmount": float, "transactionType": string}`
   - Response: `201 Created`

#### **3. User Routes**

- **POST /users**: Create a new user.
   - Request body: `{"userID": string, "username": string, "email": string, "balance": float}`
   - Response: `201 Created`

- **GET /users**: Fetch all users.
   - Response: `200 OK` with list of users.

#### **4. Event Routes**

- **POST /events**: Create a new event.
   - Request body: `{"eventID": string, "eventName": string, "eventDate": string}`
   - Response: `201 Created`

- **GET /events/{event_id}**: Fetch details of an event.
   - Path param: `event_id`
   - Response: `200 OK`

- **GET /events**: Get all events.
   - Response: `200 OK` with list of events.

---

#### ** Environment Variables**

The `.env` file contains the environment variables required to configure the application, such as database credentials, API keys, etc. Example:

```env
SERVER_PORT=8081
SERVER_MODE=development
RATE_LIMIT_REQUESTS_PER_MINUTE=60
LOG_LEVEL=debug
BET_MAX_AMOUNT=1000
BET_MIN_AMOUNT=10
```

---

### **Testing**

Unit tests are provided to ensure that the core functionality works as expected.

#### **Sample Test Cases:**

- **TestUserHandler_Success**: Verifies that a new user is created successfully and returns a `201 Created` status.
- **TestUserHandler_UserAlreadyExists**: Verifies that an attempt to create a user with an existing `userID` returns a `409 Conflict` status.
- **TestUserHandler_InvalidEmail**: Verifies that a user with an invalid email address returns a `400 Bad Request`.
- **TestUserHandler_NegativeBalance**: Verifies that a user with a negative balance returns a `400 Bad Request`.

---

### **Conclusion**

The Bet Settlement Engine is designed to handle bet creation, settlement, and transaction processing in a robust and scalable manner. With easy-to-use APIs, environment configuration via Docker, and clear unit tests, this project provides an efficient platform for managing bets and user data.