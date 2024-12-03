# **Wombat: A Social Network for Developers**

Wombat is a social platform designed for developers to connect, share knowledge, and showcase their projects. Built with **Golang** on the backend and **React** on the frontend, Wombat delivers high performance, scalability, and an engaging developer-centric experience.

---

## **Features**

- 🛠️ **Profile Management**: Showcase your professional profile with links to GitHub, LinkedIn, and personal projects.  
- 🚀 **Project Showcasing**: Share your open-source or personal projects with the community.  
- 💬 **Tech Discussions**: Ask questions, share expertise, and participate in vibrant discussions.  
- 🌐 **Communities**: Join or create communities centered on programming languages, frameworks, or general tech topics.  
- 🤝 **Collaborations**: Discover like-minded developers and collaborate on projects.  
- 🎨 **Feed Customization**: Tailor your feed to focus on the topics that matter to you.  
- 🌙 **Dark Mode**: Enjoy a developer-friendly dark mode for enhanced comfort.

---

## **Tech Stack**

### **Backend**
- **Golang**: High-performance backend services.
- **MongoDB**: Scalable and reliable NoSQL database.
- **Redis**: Real-time caching and data management.
- **JWT Authentication**: Secure and modern authentication.

### **Frontend**
- **React**: Modern and dynamic user interfaces.
- **React Query**: Simplified state management and efficient data fetching.
- **Axios**: Robust API interaction.
- **Key Feature**: All frontend components were built from scratch by our team. We did not use any pre-made components from libraries or frameworks, ensuring complete customization and control over the design and functionality of the application.

---

## **Getting Started**

Follow these steps to set up Wombat locally on your machine.

### **Prerequisites**
- **Node.js** and **npm** (for the frontend)
- **Golang** (for the backend)
- **MongoDB** and **Redis** instances running locally or on the cloud

### **Installation**

#### Clone the Repository
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/wombat.git
   cd wombat
   ```

#### Backend Setup
2. Navigate to the backend directory:
   ```bash
   cd backend
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Configure environment variables:
   ```bash
   cp .env.example .env
   ```
   Edit the `.env` file with your configurations (e.g., MongoDB URI, Redis host).
5. Start the backend server:
   ```bash
   go run app.go
   ```
   The backend API will be available at: `http://localhost:8080`

#### Frontend Setup
6. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
7. Install dependencies:
   ```bash
   npm install
   ```
8. Start the development server:
   ```bash
   npm start
   ```
   The application will be available at: `http://localhost:5173`

---

## **Usage**

Once the servers are running:
1. Open your browser and navigate to: `http://localhost:5173` to access Wombat.
2. Ensure the backend API is running at: `http://localhost:8080` if you want to use it with another database.

---

## **Contributing**

We welcome contributions from the community! To contribute:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Make your changes and commit:
   ```bash
   git commit -m "Add your commit message"
   ```
4. Push your changes:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Create a Pull Request.
