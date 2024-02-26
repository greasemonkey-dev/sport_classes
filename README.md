
# Sport Class Scheduling System
## Exercise

This project aims to develop a system for scheduling sport classes, focusing on clean architecture, extensibility, and robust error handling.

### Key Objectives:

* User-centric: Allow users to search for instructors based on specific sports and book classes with available instructors.
* RESTful API: Implement a REST API for efficient data exchange using HTTP requests.
* Clean Architecture: Adhere to the principles of clean architecture for better maintainability and testability.
* Extensible Design: Design the system to accommodate future functionality additions seamlessly.
* Robust Testing: Ensure code quality and reliability through comprehensive unit and integration tests.
* Error Handling: Implement proper error handling mechanisms to provide informative responses to unexpected situations.
* Logging: Integrate logging capabilities for system monitoring and debugging purposes.
## Entities:

* Instructor: Represents a sports club instructor with the following attributes:
* name: String representing the instructor's name (e.g., "Eli Copter").
* sports: Array of strings listing the sports the instructor teaches (e.g., ["weight lifting", "yoga"]).
* availableDates: Array of objects representing available time slots, where each object has:
* from: Long integer representing the start time in milliseconds since epoch (inclusive).
* to: Long integer representing the end time in milliseconds since epoch (inclusive).
* score: Float value between 0 and 10.0 representing the instructor's rating.
## Dates:

* Dates are stored as milliseconds since epoch (inclusive) for consistency and accurate comparisons.
## Scores:

* Instructor scores range from 0 to 10.0 to represent their overall rating.
## Data Source:

* Initially, create an instructor list in a JSON file to simulate data storage.
Design the system to support large-scale data handling (millions of instructors, thousands of sports and availabilities per instructor) in memory, considering future migration to persistent storage options like databases.
## Server Configuration:

Allow setting the server port number through a separate configuration file for flexible deployment.
## REST API Endpoints:

GET /instructs?sport=<SPORT>&date=<DATE>&minScore=<SCORE>

### Retrieves a list of instructors matching the specified criteria:
* sport: Required string parameter specifying the desired sport.
* date: Required long integer parameter representing the desired date (milliseconds since epoch).
* minScore: Optional float parameter representing the minimum acceptable instructor score (default: 0).
## Response:
* Array of instructor names ordered by score (highest to lowest).
* Empty array if no suitable instructors are found.
* HTTP status code:
* 200 (OK) for successful retrieval.
* 400 (BAD REQUEST) for invalid parameters (e.g., missing sport, invalid date format).


### Mock Data Source:

Initially use mock data (e.g., a JSON file) as a data source, but ensure code is adaptable to connect with real data sources like databases or external APIs.
###Code Structure:

Implement the system adhering to clean architecture principles, separating concerns and promoting maintainability.
Utilize unit and integration tests to verify code functionality and ensure high quality.
Integrate logging mechanisms for system monitoring and debugging purposes.
This exercise provides a comprehensive framework for developing the sport class scheduling system, emphasizing clean architecture, user experience, and robust technical practices.
