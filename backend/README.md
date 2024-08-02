# Backend API Challenge

## Objective
Develop a RESTful API for a simple project management application. This API will allow users to manage projects, tasks, and comments on tasks. The main goal of this challenge is to test your ability to design and implement a robust, scalable, and secure API.

## Requirements

### 1. User Authentication
- Implement user registration and login functionality.
- Use JWT (JSON Web Token) for authentication.
- Protect routes to ensure only authenticated users can access them.

### 2. Projects
- Users should be able to create, read, update, and delete projects.
- Each project should have:
  - An ID (automatically generated)
  - A title
  - A description
  - A list of tasks (initially empty)

### 3. Tasks
- Users should be able to add, update, delete, and view tasks within a project.
- Each task should have:
  - An ID (automatically generated)
  - A title
  - A description
  - A status (e.g., "To Do", "In Progress", "Done")
  - A due date
  - A list of comments (initially empty)

### 4. Comments
- Users should be able to add, update, delete, and view comments on tasks.
- Each comment should have:
  - An ID (automatically generated)
  - A content
  - A timestamp of when the comment was created

## Bonus Points
- **Testing:** Write unit tests for your endpoints.
- **Documentation:** Use Swagger or similar to document your API.
- **Error Handling:** Implement comprehensive error handling.
- **Performance:** Optimize your API for performance and scalability.
- **Security:** Implement security best practices (e.g., input validation, rate limiting).

## Technical Requirements
- Use Node.js with Express.js for the server.
- Use a relational database (e.g., PostgreSQL, MySQL) or a NoSQL database (e.g., MongoDB).
- Use an ORM (e.g., Sequelize, TypeORM) if using a relational database.
- Follow RESTful principles in your API design.

## Submission
- Provide a GitHub repository link with your project code.
- Include a README file with instructions on how to set up and run your project locally.
- Ensure your project can be easily tested by providing necessary setup scripts or Docker configurations.

## Evaluation Criteria
- Code quality and organization
- Adherence to RESTful principles
- Security and performance considerations
- Completeness and correctness of the functionality
- Quality of the documentation
