# Snack CRUD Application

This is a proof of concept project demonstrating a CRUD (Create, Read, Update, Delete) application for managing a list of snacks. The project is built using Go for the backend, HTMX for handling AJAX requests, and UIKit for the frontend styling.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Author](#author)

## Features

- Add new snacks
- View a list of snacks
- Update snack details
- Delete snacks
- Real-time updates without page reloads using HTMX
- Responsive design using UIKit

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.16 or later)
- Node.js and npm (for frontend dependencies)
- Git

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/snack-crud-app.git
    cd snack-crud-app
    ```

2. Install Go dependencies:

    ```bash
    go mod tidy
    ```

3. Install frontend dependencies:

    ```bash
    npm install
    ```

### Running the Application

1. Start the Go server:

    ```bash
    go run main.go
    ```

2. Open your web browser and navigate to `http://localhost:8080`.

## Usage

### Adding a Snack

1. Click the "Add Snack" button.
2. Fill in the snack details in the form that appears.
3. Click "Submit" to add the snack to the list.

### Viewing Snacks

- The home page displays a list of all snacks.

### Updating a Snack

1. Click the "Edit" button next to the snack you want to update.
2. Modify the snack details in the form that appears.
3. Click "Save" to update the snack.

### Deleting a Snack

- Click the "Delete" button next to the snack you want to remove.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you would like to contribute to this project.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

This project was created by JB Tellez jumping off from [excellent tutorial](https://dev.to/calvinmclean/how-to-build-a-web-application-with-htmx-and-go-3183) by Calvin McLean.