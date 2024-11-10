# DeveloperBin

DeveloperBin is a minimalistic code-sharing snippet site, created as a weekend project to explore Golang templating and frontend styling. This site allows users to post and view code snippets without authentication, making it an ideal tool for quick code sharing.

You can also view it at [developerbin.com](http://developerbin.com).

## Features

- **Share Code Snippets**: Users can create new code snippets with a title, description, language, and code content.
- **Syntax Highlighting**: Each snippet preview is styled using the Atom One Dark theme for readability and consistency.
- **Frontend Styling with TailwindCSS and Highlight.js**: Clean, responsive design with syntax highlighting and a code editor.
- **CodeMirror Editor**: Enables a smooth code-writing experience with real-time syntax highlighting on the "New Post" page.

## Tech Stack

- **Backend**: Golang with Fiber framework for handling HTTP requests and HTML templating.
- **Frontend**: TailwindCSS for styling, Highlight.js for code snippet display, and CodeMirror for the code editor.
- **Database**: SQLite (or other lightweight databases, if configured) for storing posts.

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/developerbin.git
   cd developerbin
   ```

2. **Install Dependencies**:
   Ensure you have Go installed, then initialize the Go modules:

   ```bash
   go mod download
   ```

3. **Run the Application**:
   Start the application locally:

   ```bash
   go run main.go
   ```

4. **View in Browser**:
   Open [http://localhost:3000](http://localhost:3000) in your browser to access DeveloperBin.

## Project Structure

- **`main.go`**: Sets up routing, database, and template functions.
- **`templates/`**: HTML templates for pages, including:
  - `index.html`: Lists all posts with a preview of each code snippet.
  - `show.html`: Displays the full code snippet for a specific post.
  - `new.html`: Form to create a new code snippet.
- **`static/`**: CSS files and JavaScript dependencies for syntax highlighting and styling.

## Usage

1. **Home Page**: Lists all code snippets with a brief description and a truncated code preview.
2. **Create New Post**: Click "Create New Post" to add a new code snippet with a title, description, language, and code content.
3. **View Full Snippet**: Click "Read More" on any snippet preview to view the full code snippet on a separate page.

## Future Improvements

- **Authentication**: Optionally add user accounts for snippet ownership.
- **Edit and Delete Snippets**: Add functionality to manage posted snippets.
- **Search and Filter**: Enable searching by programming language or title for easier navigation.

## License

This project is open source and available under the MIT License. Feel free to use, modify, and distribute as needed.
